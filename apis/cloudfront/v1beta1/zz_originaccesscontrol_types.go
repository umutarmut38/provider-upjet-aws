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

type OriginAccessControlObservation struct {

	// The description of the Origin Access Control. It may be empty.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The current version of this Origin Access Control.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// The unique identifier of this Origin Access Control.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A name that identifies the Origin Access Control.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of origin that this Origin Access Control is for. The only valid value is s3.
	OriginAccessControlOriginType *string `json:"originAccessControlOriginType,omitempty" tf:"origin_access_control_origin_type,omitempty"`

	// Specifies which requests CloudFront signs. Specify always for the most common use case. Allowed values: always, never, no-override.
	SigningBehavior *string `json:"signingBehavior,omitempty" tf:"signing_behavior,omitempty"`

	// Determines how CloudFront signs (authenticates) requests. Valid values: sigv4.
	SigningProtocol *string `json:"signingProtocol,omitempty" tf:"signing_protocol,omitempty"`
}

type OriginAccessControlParameters struct {

	// The description of the Origin Access Control. It may be empty.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A name that identifies the Origin Access Control.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of origin that this Origin Access Control is for. The only valid value is s3.
	// +kubebuilder:validation:Optional
	OriginAccessControlOriginType *string `json:"originAccessControlOriginType,omitempty" tf:"origin_access_control_origin_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies which requests CloudFront signs. Specify always for the most common use case. Allowed values: always, never, no-override.
	// +kubebuilder:validation:Optional
	SigningBehavior *string `json:"signingBehavior,omitempty" tf:"signing_behavior,omitempty"`

	// Determines how CloudFront signs (authenticates) requests. Valid values: sigv4.
	// +kubebuilder:validation:Optional
	SigningProtocol *string `json:"signingProtocol,omitempty" tf:"signing_protocol,omitempty"`
}

// OriginAccessControlSpec defines the desired state of OriginAccessControl
type OriginAccessControlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OriginAccessControlParameters `json:"forProvider"`
}

// OriginAccessControlStatus defines the observed state of OriginAccessControl.
type OriginAccessControlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OriginAccessControlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OriginAccessControl is the Schema for the OriginAccessControls API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type OriginAccessControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.originAccessControlOriginType)",message="originAccessControlOriginType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.signingBehavior)",message="signingBehavior is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.signingProtocol)",message="signingProtocol is a required parameter"
	Spec   OriginAccessControlSpec   `json:"spec"`
	Status OriginAccessControlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OriginAccessControlList contains a list of OriginAccessControls
type OriginAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OriginAccessControl `json:"items"`
}

// Repository type metadata.
var (
	OriginAccessControl_Kind             = "OriginAccessControl"
	OriginAccessControl_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OriginAccessControl_Kind}.String()
	OriginAccessControl_KindAPIVersion   = OriginAccessControl_Kind + "." + CRDGroupVersion.String()
	OriginAccessControl_GroupVersionKind = CRDGroupVersion.WithKind(OriginAccessControl_Kind)
)

func init() {
	SchemeBuilder.Register(&OriginAccessControl{}, &OriginAccessControlList{})
}
