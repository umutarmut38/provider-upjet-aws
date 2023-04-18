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

type KeyGroupObservation struct {

	// A comment to describe the key group..
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The identifier for this version of the key group.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// The identifier for the key group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of the identifiers of the public keys in the key group.
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`

	// A name to identify the key group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type KeyGroupParameters struct {

	// A comment to describe the key group..
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// References to PublicKey to populate items.
	// +kubebuilder:validation:Optional
	ItemRefs []v1.Reference `json:"itemRefs,omitempty" tf:"-"`

	// Selector for a list of PublicKey to populate items.
	// +kubebuilder:validation:Optional
	ItemSelector *v1.Selector `json:"itemSelector,omitempty" tf:"-"`

	// A list of the identifiers of the public keys in the key group.
	// +crossplane:generate:reference:type=PublicKey
	// +crossplane:generate:reference:refFieldName=ItemRefs
	// +crossplane:generate:reference:selectorFieldName=ItemSelector
	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`

	// A name to identify the key group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// KeyGroupSpec defines the desired state of KeyGroup
type KeyGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyGroupParameters `json:"forProvider"`
}

// KeyGroupStatus defines the observed state of KeyGroup.
type KeyGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KeyGroup is the Schema for the KeyGroups API. Provides a CloudFront key group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type KeyGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   KeyGroupSpec   `json:"spec"`
	Status KeyGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyGroupList contains a list of KeyGroups
type KeyGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyGroup `json:"items"`
}

// Repository type metadata.
var (
	KeyGroup_Kind             = "KeyGroup"
	KeyGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KeyGroup_Kind}.String()
	KeyGroup_KindAPIVersion   = KeyGroup_Kind + "." + CRDGroupVersion.String()
	KeyGroup_GroupVersionKind = CRDGroupVersion.WithKind(KeyGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&KeyGroup{}, &KeyGroupList{})
}
