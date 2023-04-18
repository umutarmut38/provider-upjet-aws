/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClusterParameterGroupObservation struct {

	// The ARN of the neptune cluster parameter group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The description of the neptune cluster parameter group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The family of the neptune cluster parameter group.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// The neptune cluster parameter group name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of neptune parameters to apply.
	Parameter []ParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ClusterParameterGroupParameters struct {

	// The description of the neptune cluster parameter group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The family of the neptune cluster parameter group.
	// +kubebuilder:validation:Optional
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// A list of neptune parameters to apply.
	// +kubebuilder:validation:Optional
	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ParameterObservation struct {

	// Valid values are immediate and pending-reboot. Defaults to pending-reboot.
	ApplyMethod *string `json:"applyMethod,omitempty" tf:"apply_method,omitempty"`

	// The name of the neptune cluster parameter group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the neptune parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParameterParameters struct {

	// Valid values are immediate and pending-reboot. Defaults to pending-reboot.
	// +kubebuilder:validation:Optional
	ApplyMethod *string `json:"applyMethod,omitempty" tf:"apply_method,omitempty"`

	// The name of the neptune cluster parameter group.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The value of the neptune parameter.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// ClusterParameterGroupSpec defines the desired state of ClusterParameterGroup
type ClusterParameterGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameterGroupParameters `json:"forProvider"`
}

// ClusterParameterGroupStatus defines the observed state of ClusterParameterGroup.
type ClusterParameterGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterParameterGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterParameterGroup is the Schema for the ClusterParameterGroups API. Manages a Neptune Cluster Parameter Group
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ClusterParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.family)",message="family is a required parameter"
	Spec   ClusterParameterGroupSpec   `json:"spec"`
	Status ClusterParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterParameterGroupList contains a list of ClusterParameterGroups
type ClusterParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterParameterGroup `json:"items"`
}

// Repository type metadata.
var (
	ClusterParameterGroup_Kind             = "ClusterParameterGroup"
	ClusterParameterGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterParameterGroup_Kind}.String()
	ClusterParameterGroup_KindAPIVersion   = ClusterParameterGroup_Kind + "." + CRDGroupVersion.String()
	ClusterParameterGroup_GroupVersionKind = CRDGroupVersion.WithKind(ClusterParameterGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterParameterGroup{}, &ClusterParameterGroupList{})
}
