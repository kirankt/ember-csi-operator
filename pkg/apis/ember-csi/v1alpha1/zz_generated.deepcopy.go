// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmberCSI) DeepCopyInto(out *EmberCSI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmberCSI.
func (in *EmberCSI) DeepCopy() *EmberCSI {
	if in == nil {
		return nil
	}
	out := new(EmberCSI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmberCSI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmberCSIConfig) DeepCopyInto(out *EmberCSIConfig) {
	*out = *in
	in.EnvVars.DeepCopyInto(&out.EnvVars)
	out.SysFiles = in.SysFiles
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmberCSIConfig.
func (in *EmberCSIConfig) DeepCopy() *EmberCSIConfig {
	if in == nil {
		return nil
	}
	out := new(EmberCSIConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmberCSIList) DeepCopyInto(out *EmberCSIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EmberCSI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmberCSIList.
func (in *EmberCSIList) DeepCopy() *EmberCSIList {
	if in == nil {
		return nil
	}
	out := new(EmberCSIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmberCSIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmberCSISpec) DeepCopyInto(out *EmberCSISpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Topologies != nil {
		in, out := &in.Topologies, &out.Topologies
		*out = make([]Topologies, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmberCSISpec.
func (in *EmberCSISpec) DeepCopy() *EmberCSISpec {
	if in == nil {
		return nil
	}
	out := new(EmberCSISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmberCSIStatus) DeepCopyInto(out *EmberCSIStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmberCSIStatus.
func (in *EmberCSIStatus) DeepCopy() *EmberCSIStatus {
	if in == nil {
		return nil
	}
	out := new(EmberCSIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVars) DeepCopyInto(out *EnvVars) {
	*out = *in
	if in.EnvSecrets != nil {
		in, out := &in.EnvSecrets, &out.EnvSecrets
		*out = make([]Secrets, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVars.
func (in *EnvVars) DeepCopy() *EnvVars {
	if in == nil {
		return nil
	}
	out := new(EnvVars)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secrets) DeepCopyInto(out *Secrets) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secrets.
func (in *Secrets) DeepCopy() *Secrets {
	if in == nil {
		return nil
	}
	out := new(Secrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Topologies) DeepCopyInto(out *Topologies) {
	*out = *in
	if in.Topology != nil {
		in, out := &in.Topology, &out.Topology
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]v1.NodeSelectorRequirement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Topologies.
func (in *Topologies) DeepCopy() *Topologies {
	if in == nil {
		return nil
	}
	out := new(Topologies)
	in.DeepCopyInto(out)
	return out
}
