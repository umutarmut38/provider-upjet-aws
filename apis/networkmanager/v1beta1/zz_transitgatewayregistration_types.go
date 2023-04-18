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

type TransitGatewayRegistrationObservation struct {

	// The ID of the Global Network to register to.
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of the Transit Gateway to register.
	TransitGatewayArn *string `json:"transitGatewayArn,omitempty" tf:"transit_gateway_arn,omitempty"`
}

type TransitGatewayRegistrationParameters struct {

	// The ID of the Global Network to register to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.GlobalNetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	// Reference to a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDRef *v1.Reference `json:"globalNetworkIdRef,omitempty" tf:"-"`

	// Selector for a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDSelector *v1.Selector `json:"globalNetworkIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ARN of the Transit Gateway to register.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.TransitGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	TransitGatewayArn *string `json:"transitGatewayArn,omitempty" tf:"transit_gateway_arn,omitempty"`

	// Reference to a TransitGateway in ec2 to populate transitGatewayArn.
	// +kubebuilder:validation:Optional
	TransitGatewayArnRef *v1.Reference `json:"transitGatewayArnRef,omitempty" tf:"-"`

	// Selector for a TransitGateway in ec2 to populate transitGatewayArn.
	// +kubebuilder:validation:Optional
	TransitGatewayArnSelector *v1.Selector `json:"transitGatewayArnSelector,omitempty" tf:"-"`
}

// TransitGatewayRegistrationSpec defines the desired state of TransitGatewayRegistration
type TransitGatewayRegistrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayRegistrationParameters `json:"forProvider"`
}

// TransitGatewayRegistrationStatus defines the observed state of TransitGatewayRegistration.
type TransitGatewayRegistrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayRegistrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRegistration is the Schema for the TransitGatewayRegistrations API. Registers a transit gateway to a global network.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayRegistration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayRegistrationSpec   `json:"spec"`
	Status            TransitGatewayRegistrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRegistrationList contains a list of TransitGatewayRegistrations
type TransitGatewayRegistrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayRegistration `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayRegistration_Kind             = "TransitGatewayRegistration"
	TransitGatewayRegistration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayRegistration_Kind}.String()
	TransitGatewayRegistration_KindAPIVersion   = TransitGatewayRegistration_Kind + "." + CRDGroupVersion.String()
	TransitGatewayRegistration_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayRegistration_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayRegistration{}, &TransitGatewayRegistrationList{})
}
