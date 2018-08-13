// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSProvisioner) DeepCopyInto(out *EFSProvisioner) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSProvisioner.
func (in *EFSProvisioner) DeepCopy() *EFSProvisioner {
	if in == nil {
		return nil
	}
	out := new(EFSProvisioner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EFSProvisioner) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSProvisionerList) DeepCopyInto(out *EFSProvisionerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EFSProvisioner, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSProvisionerList.
func (in *EFSProvisionerList) DeepCopy() *EFSProvisionerList {
	if in == nil {
		return nil
	}
	out := new(EFSProvisionerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EFSProvisionerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSProvisionerSpec) DeepCopyInto(out *EFSProvisionerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSProvisionerSpec.
func (in *EFSProvisionerSpec) DeepCopy() *EFSProvisionerSpec {
	if in == nil {
		return nil
	}
	out := new(EFSProvisionerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSProvisionerStatus) DeepCopyInto(out *EFSProvisionerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSProvisionerStatus.
func (in *EFSProvisionerStatus) DeepCopy() *EFSProvisionerStatus {
	if in == nil {
		return nil
	}
	out := new(EFSProvisionerStatus)
	in.DeepCopyInto(out)
	return out
}
