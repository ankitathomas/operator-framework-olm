// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InstallPlanReference).DeepCopyInto(out.(*InstallPlanReference))
			return nil
		}, InType: reflect.TypeOf(&InstallPlanReference{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Subscription).DeepCopyInto(out.(*Subscription))
			return nil
		}, InType: reflect.TypeOf(&Subscription{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SubscriptionList).DeepCopyInto(out.(*SubscriptionList))
			return nil
		}, InType: reflect.TypeOf(&SubscriptionList{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SubscriptionSpec).DeepCopyInto(out.(*SubscriptionSpec))
			return nil
		}, InType: reflect.TypeOf(&SubscriptionSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SubscriptionStatus).DeepCopyInto(out.(*SubscriptionStatus))
			return nil
		}, InType: reflect.TypeOf(&SubscriptionStatus{})},
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallPlanReference) DeepCopyInto(out *InstallPlanReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallPlanReference.
func (in *InstallPlanReference) DeepCopy() *InstallPlanReference {
	if in == nil {
		return nil
	}
	out := new(InstallPlanReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subscription) DeepCopyInto(out *Subscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		if *in == nil {
			*out = nil
		} else {
			*out = new(SubscriptionSpec)
			**out = **in
		}
	}
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subscription.
func (in *Subscription) DeepCopy() *Subscription {
	if in == nil {
		return nil
	}
	out := new(Subscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Subscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionList) DeepCopyInto(out *SubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Subscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionList.
func (in *SubscriptionList) DeepCopy() *SubscriptionList {
	if in == nil {
		return nil
	}
	out := new(SubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionSpec) DeepCopyInto(out *SubscriptionSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionSpec.
func (in *SubscriptionSpec) DeepCopy() *SubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(SubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionStatus) DeepCopyInto(out *SubscriptionStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Install != nil {
		in, out := &in.Install, &out.Install
		if *in == nil {
			*out = nil
		} else {
			*out = new(InstallPlanReference)
			**out = **in
		}
	}
	in.LastUpdated.DeepCopyInto(&out.LastUpdated)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionStatus.
func (in *SubscriptionStatus) DeepCopy() *SubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(SubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}
