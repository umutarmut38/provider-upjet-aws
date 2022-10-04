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

type DocumentationVersionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DocumentationVersionParameters struct {

	// The description of the API documentation version.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the associated Rest API
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`

	// The version identifier of the API documentation snapshot.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

// DocumentationVersionSpec defines the desired state of DocumentationVersion
type DocumentationVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DocumentationVersionParameters `json:"forProvider"`
}

// DocumentationVersionStatus defines the observed state of DocumentationVersion.
type DocumentationVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DocumentationVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentationVersion is the Schema for the DocumentationVersions API. Provides a resource to manage an API Gateway Documentation Version.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DocumentationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocumentationVersionSpec   `json:"spec"`
	Status            DocumentationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentationVersionList contains a list of DocumentationVersions
type DocumentationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DocumentationVersion `json:"items"`
}

// Repository type metadata.
var (
	DocumentationVersion_Kind             = "DocumentationVersion"
	DocumentationVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DocumentationVersion_Kind}.String()
	DocumentationVersion_KindAPIVersion   = DocumentationVersion_Kind + "." + CRDGroupVersion.String()
	DocumentationVersion_GroupVersionKind = CRDGroupVersion.WithKind(DocumentationVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&DocumentationVersion{}, &DocumentationVersionList{})
}
