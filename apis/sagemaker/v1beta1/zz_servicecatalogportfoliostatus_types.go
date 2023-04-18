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

type ServicecatalogPortfolioStatusObservation struct {

	// The AWS Region the Servicecatalog portfolio status resides in.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether Service Catalog is enabled or disabled in SageMaker. Valid values are Enabled and Disabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ServicecatalogPortfolioStatusParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Whether Service Catalog is enabled or disabled in SageMaker. Valid values are Enabled and Disabled.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// ServicecatalogPortfolioStatusSpec defines the desired state of ServicecatalogPortfolioStatus
type ServicecatalogPortfolioStatusSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServicecatalogPortfolioStatusParameters `json:"forProvider"`
}

// ServicecatalogPortfolioStatusStatus defines the observed state of ServicecatalogPortfolioStatus.
type ServicecatalogPortfolioStatusStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServicecatalogPortfolioStatusObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogPortfolioStatus is the Schema for the ServicecatalogPortfolioStatuss API. Manages status of Service Catalog in SageMaker. Service Catalog is used to create SageMaker projects.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ServicecatalogPortfolioStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.status)",message="status is a required parameter"
	Spec   ServicecatalogPortfolioStatusSpec   `json:"spec"`
	Status ServicecatalogPortfolioStatusStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogPortfolioStatusList contains a list of ServicecatalogPortfolioStatuss
type ServicecatalogPortfolioStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicecatalogPortfolioStatus `json:"items"`
}

// Repository type metadata.
var (
	ServicecatalogPortfolioStatus_Kind             = "ServicecatalogPortfolioStatus"
	ServicecatalogPortfolioStatus_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServicecatalogPortfolioStatus_Kind}.String()
	ServicecatalogPortfolioStatus_KindAPIVersion   = ServicecatalogPortfolioStatus_Kind + "." + CRDGroupVersion.String()
	ServicecatalogPortfolioStatus_GroupVersionKind = CRDGroupVersion.WithKind(ServicecatalogPortfolioStatus_Kind)
)

func init() {
	SchemeBuilder.Register(&ServicecatalogPortfolioStatus{}, &ServicecatalogPortfolioStatusList{})
}
