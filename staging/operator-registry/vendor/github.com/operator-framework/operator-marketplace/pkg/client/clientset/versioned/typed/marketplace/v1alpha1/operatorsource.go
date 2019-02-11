/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/operator-framework/operator-marketplace/pkg/apis/marketplace/v1alpha1"
	scheme "github.com/operator-framework/operator-marketplace/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OperatorSourcesGetter has a method to return a OperatorSourceInterface.
// A group's client should implement this interface.
type OperatorSourcesGetter interface {
	OperatorSources(namespace string) OperatorSourceInterface
}

// OperatorSourceInterface has methods to work with OperatorSource resources.
type OperatorSourceInterface interface {
	Create(*v1alpha1.OperatorSource) (*v1alpha1.OperatorSource, error)
	Update(*v1alpha1.OperatorSource) (*v1alpha1.OperatorSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.OperatorSource, error)
	List(opts v1.ListOptions) (*v1alpha1.OperatorSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OperatorSource, err error)
	OperatorSourceExpansion
}

// operatorSources implements OperatorSourceInterface
type operatorSources struct {
	client rest.Interface
	ns     string
}

// newOperatorSources returns a OperatorSources
func newOperatorSources(c *MarketplaceV1alpha1Client, namespace string) *operatorSources {
	return &operatorSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the operatorSource, and returns the corresponding operatorSource object, and an error if there is any.
func (c *operatorSources) Get(name string, options v1.GetOptions) (result *v1alpha1.OperatorSource, err error) {
	result = &v1alpha1.OperatorSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("operatorsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OperatorSources that match those selectors.
func (c *operatorSources) List(opts v1.ListOptions) (result *v1alpha1.OperatorSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OperatorSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("operatorsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested operatorSources.
func (c *operatorSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("operatorsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a operatorSource and creates it.  Returns the server's representation of the operatorSource, and an error, if there is any.
func (c *operatorSources) Create(operatorSource *v1alpha1.OperatorSource) (result *v1alpha1.OperatorSource, err error) {
	result = &v1alpha1.OperatorSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("operatorsources").
		Body(operatorSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a operatorSource and updates it. Returns the server's representation of the operatorSource, and an error, if there is any.
func (c *operatorSources) Update(operatorSource *v1alpha1.OperatorSource) (result *v1alpha1.OperatorSource, err error) {
	result = &v1alpha1.OperatorSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("operatorsources").
		Name(operatorSource.Name).
		Body(operatorSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the operatorSource and deletes it. Returns an error if one occurs.
func (c *operatorSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("operatorsources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *operatorSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("operatorsources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched operatorSource.
func (c *operatorSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OperatorSource, err error) {
	result = &v1alpha1.OperatorSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("operatorsources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
