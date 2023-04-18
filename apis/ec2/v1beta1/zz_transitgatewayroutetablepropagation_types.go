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

type TransitGatewayRouteTablePropagationObservation struct {

	// EC2 Transit Gateway Route Table identifier combined with EC2 Transit Gateway Attachment identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Identifier of the resource
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Type of the resource
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableID *string `json:"transitGatewayRouteTableId,omitempty" tf:"transit_gateway_route_table_id,omitempty"`
}

type TransitGatewayRouteTablePropagationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Identifier of EC2 Transit Gateway Attachment.
	// +crossplane:generate:reference:type=TransitGatewayVPCAttachment
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// Reference to a TransitGatewayVPCAttachment to populate transitGatewayAttachmentId.
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDRef *v1.Reference `json:"transitGatewayAttachmentIdRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayVPCAttachment to populate transitGatewayAttachmentId.
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDSelector *v1.Selector `json:"transitGatewayAttachmentIdSelector,omitempty" tf:"-"`

	// Identifier of EC2 Transit Gateway Route Table.
	// +crossplane:generate:reference:type=TransitGatewayRouteTable
	// +kubebuilder:validation:Optional
	TransitGatewayRouteTableID *string `json:"transitGatewayRouteTableId,omitempty" tf:"transit_gateway_route_table_id,omitempty"`

	// Reference to a TransitGatewayRouteTable to populate transitGatewayRouteTableId.
	// +kubebuilder:validation:Optional
	TransitGatewayRouteTableIDRef *v1.Reference `json:"transitGatewayRouteTableIdRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayRouteTable to populate transitGatewayRouteTableId.
	// +kubebuilder:validation:Optional
	TransitGatewayRouteTableIDSelector *v1.Selector `json:"transitGatewayRouteTableIdSelector,omitempty" tf:"-"`
}

// TransitGatewayRouteTablePropagationSpec defines the desired state of TransitGatewayRouteTablePropagation
type TransitGatewayRouteTablePropagationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayRouteTablePropagationParameters `json:"forProvider"`
}

// TransitGatewayRouteTablePropagationStatus defines the observed state of TransitGatewayRouteTablePropagation.
type TransitGatewayRouteTablePropagationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayRouteTablePropagationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRouteTablePropagation is the Schema for the TransitGatewayRouteTablePropagations API. Manages an EC2 Transit Gateway Route Table propagation
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayRouteTablePropagation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayRouteTablePropagationSpec   `json:"spec"`
	Status            TransitGatewayRouteTablePropagationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRouteTablePropagationList contains a list of TransitGatewayRouteTablePropagations
type TransitGatewayRouteTablePropagationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayRouteTablePropagation `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayRouteTablePropagation_Kind             = "TransitGatewayRouteTablePropagation"
	TransitGatewayRouteTablePropagation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayRouteTablePropagation_Kind}.String()
	TransitGatewayRouteTablePropagation_KindAPIVersion   = TransitGatewayRouteTablePropagation_Kind + "." + CRDGroupVersion.String()
	TransitGatewayRouteTablePropagation_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayRouteTablePropagation_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayRouteTablePropagation{}, &TransitGatewayRouteTablePropagationList{})
}
