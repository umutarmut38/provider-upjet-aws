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

type TransitGatewayRouteObservation struct {

	// EC2 Transit Gateway Route Table identifier combined with destination
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TransitGatewayRouteParameters struct {

	// Indicates whether to drop traffic that matches this route (default to false).
	// +kubebuilder:validation:Optional
	Blackhole *bool `json:"blackhole,omitempty" tf:"blackhole,omitempty"`

	// IPv4 or IPv6 RFC1924 CIDR used for destination matches. Routing decisions are based on the most specific match.
	// +kubebuilder:validation:Required
	DestinationCidrBlock *string `json:"destinationCidrBlock" tf:"destination_cidr_block,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Identifier of EC2 Transit Gateway Attachment .
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
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.TransitGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("association_default_route_table_id",true)
	// +kubebuilder:validation:Optional
	TransitGatewayRouteTableID *string `json:"transitGatewayRouteTableId,omitempty" tf:"transit_gateway_route_table_id,omitempty"`

	// Reference to a TransitGateway in ec2 to populate transitGatewayRouteTableId.
	// +kubebuilder:validation:Optional
	TransitGatewayRouteTableIDRef *v1.Reference `json:"transitGatewayRouteTableIdRef,omitempty" tf:"-"`

	// Selector for a TransitGateway in ec2 to populate transitGatewayRouteTableId.
	// +kubebuilder:validation:Optional
	TransitGatewayRouteTableIDSelector *v1.Selector `json:"transitGatewayRouteTableIdSelector,omitempty" tf:"-"`
}

// TransitGatewayRouteSpec defines the desired state of TransitGatewayRoute
type TransitGatewayRouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayRouteParameters `json:"forProvider"`
}

// TransitGatewayRouteStatus defines the observed state of TransitGatewayRoute.
type TransitGatewayRouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRoute is the Schema for the TransitGatewayRoutes API. Manages an EC2 Transit Gateway Route
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayRouteSpec   `json:"spec"`
	Status            TransitGatewayRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRouteList contains a list of TransitGatewayRoutes
type TransitGatewayRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayRoute `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayRoute_Kind             = "TransitGatewayRoute"
	TransitGatewayRoute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayRoute_Kind}.String()
	TransitGatewayRoute_KindAPIVersion   = TransitGatewayRoute_Kind + "." + CRDGroupVersion.String()
	TransitGatewayRoute_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayRoute_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayRoute{}, &TransitGatewayRouteList{})
}
