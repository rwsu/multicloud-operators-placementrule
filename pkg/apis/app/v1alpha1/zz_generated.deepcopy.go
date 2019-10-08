// +build !ignore_autogenerated

// Copyright 2019 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConditionFilter) DeepCopyInto(out *ClusterConditionFilter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConditionFilter.
func (in *ClusterConditionFilter) DeepCopy() *ClusterConditionFilter {
	if in == nil {
		return nil
	}
	out := new(ClusterConditionFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericClusterReference) DeepCopyInto(out *GenericClusterReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericClusterReference.
func (in *GenericClusterReference) DeepCopy() *GenericClusterReference {
	if in == nil {
		return nil
	}
	out := new(GenericClusterReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericPlacementFields) DeepCopyInto(out *GenericPlacementFields) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]GenericClusterReference, len(*in))
		copy(*out, *in)
	}
	if in.ClusterSelector != nil {
		in, out := &in.ClusterSelector, &out.ClusterSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericPlacementFields.
func (in *GenericPlacementFields) DeepCopy() *GenericPlacementFields {
	if in == nil {
		return nil
	}
	out := new(GenericPlacementFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Placement) DeepCopyInto(out *Placement) {
	*out = *in
	in.GenericPlacementFields.DeepCopyInto(&out.GenericPlacementFields)
	if in.PlacementRef != nil {
		in, out := &in.PlacementRef, &out.PlacementRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Placement.
func (in *Placement) DeepCopy() *Placement {
	if in == nil {
		return nil
	}
	out := new(Placement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementDecision) DeepCopyInto(out *PlacementDecision) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementDecision.
func (in *PlacementDecision) DeepCopy() *PlacementDecision {
	if in == nil {
		return nil
	}
	out := new(PlacementDecision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementRule) DeepCopyInto(out *PlacementRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementRule.
func (in *PlacementRule) DeepCopy() *PlacementRule {
	if in == nil {
		return nil
	}
	out := new(PlacementRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlacementRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementRuleList) DeepCopyInto(out *PlacementRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PlacementRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementRuleList.
func (in *PlacementRuleList) DeepCopy() *PlacementRuleList {
	if in == nil {
		return nil
	}
	out := new(PlacementRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlacementRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementRuleSpec) DeepCopyInto(out *PlacementRuleSpec) {
	*out = *in
	if in.ClusterReplicas != nil {
		in, out := &in.ClusterReplicas, &out.ClusterReplicas
		*out = new(int32)
		**out = **in
	}
	in.GenericPlacementFields.DeepCopyInto(&out.GenericPlacementFields)
	if in.ClusterConditions != nil {
		in, out := &in.ClusterConditions, &out.ClusterConditions
		*out = make([]ClusterConditionFilter, len(*in))
		copy(*out, *in)
	}
	if in.ResourceHint != nil {
		in, out := &in.ResourceHint, &out.ResourceHint
		*out = new(ResourceHint)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementRuleSpec.
func (in *PlacementRuleSpec) DeepCopy() *PlacementRuleSpec {
	if in == nil {
		return nil
	}
	out := new(PlacementRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementRuleStatus) DeepCopyInto(out *PlacementRuleStatus) {
	*out = *in
	if in.Decisions != nil {
		in, out := &in.Decisions, &out.Decisions
		*out = make([]PlacementDecision, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementRuleStatus.
func (in *PlacementRuleStatus) DeepCopy() *PlacementRuleStatus {
	if in == nil {
		return nil
	}
	out := new(PlacementRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceHint) DeepCopyInto(out *ResourceHint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceHint.
func (in *ResourceHint) DeepCopy() *ResourceHint {
	if in == nil {
		return nil
	}
	out := new(ResourceHint)
	in.DeepCopyInto(out)
	return out
}
