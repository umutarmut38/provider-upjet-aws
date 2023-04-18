//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FairSharePolicyObservation) DeepCopyInto(out *FairSharePolicyObservation) {
	*out = *in
	if in.ComputeReservation != nil {
		in, out := &in.ComputeReservation, &out.ComputeReservation
		*out = new(float64)
		**out = **in
	}
	if in.ShareDecaySeconds != nil {
		in, out := &in.ShareDecaySeconds, &out.ShareDecaySeconds
		*out = new(float64)
		**out = **in
	}
	if in.ShareDistribution != nil {
		in, out := &in.ShareDistribution, &out.ShareDistribution
		*out = make([]ShareDistributionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FairSharePolicyObservation.
func (in *FairSharePolicyObservation) DeepCopy() *FairSharePolicyObservation {
	if in == nil {
		return nil
	}
	out := new(FairSharePolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FairSharePolicyParameters) DeepCopyInto(out *FairSharePolicyParameters) {
	*out = *in
	if in.ComputeReservation != nil {
		in, out := &in.ComputeReservation, &out.ComputeReservation
		*out = new(float64)
		**out = **in
	}
	if in.ShareDecaySeconds != nil {
		in, out := &in.ShareDecaySeconds, &out.ShareDecaySeconds
		*out = new(float64)
		**out = **in
	}
	if in.ShareDistribution != nil {
		in, out := &in.ShareDistribution, &out.ShareDistribution
		*out = make([]ShareDistributionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FairSharePolicyParameters.
func (in *FairSharePolicyParameters) DeepCopy() *FairSharePolicyParameters {
	if in == nil {
		return nil
	}
	out := new(FairSharePolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicy) DeepCopyInto(out *SchedulingPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicy.
func (in *SchedulingPolicy) DeepCopy() *SchedulingPolicy {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SchedulingPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicyList) DeepCopyInto(out *SchedulingPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SchedulingPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicyList.
func (in *SchedulingPolicyList) DeepCopy() *SchedulingPolicyList {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SchedulingPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicyObservation) DeepCopyInto(out *SchedulingPolicyObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.FairSharePolicy != nil {
		in, out := &in.FairSharePolicy, &out.FairSharePolicy
		*out = make([]FairSharePolicyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicyObservation.
func (in *SchedulingPolicyObservation) DeepCopy() *SchedulingPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicyParameters) DeepCopyInto(out *SchedulingPolicyParameters) {
	*out = *in
	if in.FairSharePolicy != nil {
		in, out := &in.FairSharePolicy, &out.FairSharePolicy
		*out = make([]FairSharePolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicyParameters.
func (in *SchedulingPolicyParameters) DeepCopy() *SchedulingPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicySpec) DeepCopyInto(out *SchedulingPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicySpec.
func (in *SchedulingPolicySpec) DeepCopy() *SchedulingPolicySpec {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicyStatus) DeepCopyInto(out *SchedulingPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicyStatus.
func (in *SchedulingPolicyStatus) DeepCopy() *SchedulingPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareDistributionObservation) DeepCopyInto(out *ShareDistributionObservation) {
	*out = *in
	if in.ShareIdentifier != nil {
		in, out := &in.ShareIdentifier, &out.ShareIdentifier
		*out = new(string)
		**out = **in
	}
	if in.WeightFactor != nil {
		in, out := &in.WeightFactor, &out.WeightFactor
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareDistributionObservation.
func (in *ShareDistributionObservation) DeepCopy() *ShareDistributionObservation {
	if in == nil {
		return nil
	}
	out := new(ShareDistributionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareDistributionParameters) DeepCopyInto(out *ShareDistributionParameters) {
	*out = *in
	if in.ShareIdentifier != nil {
		in, out := &in.ShareIdentifier, &out.ShareIdentifier
		*out = new(string)
		**out = **in
	}
	if in.WeightFactor != nil {
		in, out := &in.WeightFactor, &out.WeightFactor
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareDistributionParameters.
func (in *ShareDistributionParameters) DeepCopy() *ShareDistributionParameters {
	if in == nil {
		return nil
	}
	out := new(ShareDistributionParameters)
	in.DeepCopyInto(out)
	return out
}
