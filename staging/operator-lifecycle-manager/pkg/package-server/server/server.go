package server

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/provider"

	"github.com/spf13/cobra"

	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	//"k8s.io/sample-apiserver/pkg/admission/plugin/banflunder"
	//"k8s.io/sample-apiserver/pkg/admission/packageinitializer"

	"github.com/operator-framework/operator-lifecycle-manager/pkg/api/client"
	"github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/informers/externalversions"
	"github.com/operator-framework/operator-lifecycle-manager/pkg/lib/queueinformer"
	"github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/apiserver"
	genericpackagemanifests "github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/apiserver/generic"
)

// NewCommandStartPackageServer provides a CLI handler for 'start master' command
// with a default PackageServerOptions.
func NewCommandStartPackageServer(defaults *PackageServerOptions, stopCh <-chan struct{}) *cobra.Command {
	o := *defaults
	cmd := &cobra.Command{
		Short: "Launch a package API server",
		Long:  "Launch a package API server",
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Run(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()

	// flags.BoolVar(&o.InsecureKubeletTLS, "kubelet-insecure-tls", o.InsecureKubeletTLS, "Do not verify CA of serving certificates presented by Kubelets.  For testing purposes only.")
	flags.StringVar(&o.Kubeconfig, "kubeconfig", o.Kubeconfig, "The path to the kubeconfig used to connect to the Kubernetes API server and the Kubelets (defaults to in-cluster config)")
	flags.StringSliceVar(&o.WatchedNamespaces, "watched-namespaces", o.WatchedNamespaces, "The resolution at which metrics-server will retain metrics.")

	o.SecureServing.AddFlags(flags)
	o.Authentication.AddFlags(flags)
	o.Authorization.AddFlags(flags)
	o.Features.AddFlags(flags)

	return cmd
}

type PackageServerOptions struct {
	// RecommendedOptions *genericoptions.RecommendedOptions - EtcdOptions
	SecureServing  *genericoptions.SecureServingOptionsWithLoopback
	Authentication *genericoptions.DelegatingAuthenticationOptions
	Authorization  *genericoptions.DelegatingAuthorizationOptions
	Features       *genericoptions.FeatureOptions

	WakeupInterval    time.Duration
	WatchedNamespaces []string

	Kubeconfig string

	// Only to be used to for testing
	DisableAuthForTesting bool

	SharedInformerFactory informers.SharedInformerFactory
	StdOut                io.Writer
	StdErr                io.Writer
}

func NewPackageServerOptions(out, errOut io.Writer) *PackageServerOptions {
	o := &PackageServerOptions{

		SecureServing:  genericoptions.WithLoopback(genericoptions.NewSecureServingOptions()),
		Authentication: genericoptions.NewDelegatingAuthenticationOptions(),
		Authorization:  genericoptions.NewDelegatingAuthorizationOptions(),
		Features:       genericoptions.NewFeatureOptions(),

		WatchedNamespaces: []string{"local"},
		WakeupInterval:    5 * time.Second,

		StdOut: out,
		StdErr: errOut,
	}

	return o
}

func (o *PackageServerOptions) Complete() error {
	return nil
}

func (o *PackageServerOptions) Config() (*apiserver.Config, error) {
	if err := o.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	serverConfig := genericapiserver.NewConfig(genericpackagemanifests.Codecs)
	if err := o.SecureServing.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	if !o.DisableAuthForTesting {
		if err := o.Authentication.ApplyTo(&serverConfig.Authentication, serverConfig.SecureServing, nil); err != nil {
			return nil, err
		}
		if err := o.Authorization.ApplyTo(&serverConfig.Authorization); err != nil {
			return nil, err
		}
	}

	return &apiserver.Config{
		GenericConfig:  serverConfig,
		ProviderConfig: genericpackagemanifests.ProviderConfig{},
	}, nil
}

func (o *PackageServerOptions) Run(stopCh <-chan struct{}) error {
	// grab the config for the API server
	config, err := o.Config()
	if err != nil {
		return err
	}
	config.GenericConfig.EnableMetrics = true

	// set up the client config
	var clientConfig *rest.Config
	if len(o.Kubeconfig) > 0 {
		loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: o.Kubeconfig}
		loader := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{})

		clientConfig, err = loader.ClientConfig()
	} else {
		clientConfig, err = rest.InClusterConfig()
	}
	if err != nil {
		return fmt.Errorf("unable to construct lister client config: %v", err)
	}

	// set up the informers
	kubeClient, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return fmt.Errorf("unable to construct lister client: %v", err)
	}

	// create a new client for OLM types (CRs)
	crClient, err := client.NewClient(o.Kubeconfig)
	if err != nil {
		return err
	}

	// Create an informer for each catalog namespace
	catsrcSharedIndexInformers := []cache.SharedIndexInformer{}
	for _, namespace := range o.WatchedNamespaces {
		nsInformerFactory := externalversions.NewSharedInformerFactoryWithOptions(crClient, o.WakeupInterval, externalversions.WithNamespace(namespace))
		catsrcSharedIndexInformers = append(catsrcSharedIndexInformers, nsInformerFactory.Operators().V1alpha1().CatalogSources().Informer())
	}

	// Create a new queueinformer-based operator.
	queueOperator, err := queueinformer.NewOperator(o.Kubeconfig)
	if err != nil {
		return err
	}

	sourceProvider := provider.NewInMemoryProvider(catsrcSharedIndexInformers, queueOperator)

	// inject the providers into the config
	config.ProviderConfig.Provider = sourceProvider

	// we should never need to resync, since we're not worried about missing events,
	// and resync is actually for regular interval-based reconciliation these days,
	// so set the default resync interval to 0
	informerFactory := informers.NewSharedInformerFactory(kubeClient, 0)

	// complete the config to get an API server
	server, err := config.Complete(informerFactory).New()
	if err != nil {
		return err
	}

	// run the source provider's informers
	go sourceProvider.Run(stopCh)

	// add health checks
	// server.AddHealthzChecks(healthz.NamedCheck("healthz", mgr.CheckHealth))

	// run everything (the apiserver runs the shared informer factory for us)
	// mgr.RunUntil(stopCh)
	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}