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

type TransitGatewayPeeringAttachmentAccepterObservation struct {

	// EC2 Transit Gateway Attachment identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Identifier of the AWS account that owns the EC2 TGW peering.
	PeerAccountID *string `json:"peerAccountId,omitempty" tf:"peer_account_id,omitempty"`

	PeerRegion *string `json:"peerRegion,omitempty" tf:"peer_region,omitempty"`

	// Identifier of EC2 Transit Gateway to peer with.
	PeerTransitGatewayID *string `json:"peerTransitGatewayId,omitempty" tf:"peer_transit_gateway_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ID of the EC2 Transit Gateway Peering Attachment to manage.
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// Identifier of EC2 Transit Gateway.
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`
}

type TransitGatewayPeeringAttachmentAccepterParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the EC2 Transit Gateway Peering Attachment to manage.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.TransitGatewayPeeringAttachment
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// Reference to a TransitGatewayPeeringAttachment in ec2 to populate transitGatewayAttachmentId.
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDRef *v1.Reference `json:"transitGatewayAttachmentIdRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayPeeringAttachment in ec2 to populate transitGatewayAttachmentId.
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDSelector *v1.Selector `json:"transitGatewayAttachmentIdSelector,omitempty" tf:"-"`
}

// TransitGatewayPeeringAttachmentAccepterSpec defines the desired state of TransitGatewayPeeringAttachmentAccepter
type TransitGatewayPeeringAttachmentAccepterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayPeeringAttachmentAccepterParameters `json:"forProvider"`
}

// TransitGatewayPeeringAttachmentAccepterStatus defines the observed state of TransitGatewayPeeringAttachmentAccepter.
type TransitGatewayPeeringAttachmentAccepterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayPeeringAttachmentAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayPeeringAttachmentAccepter is the Schema for the TransitGatewayPeeringAttachmentAccepters API. Manages the accepter's side of an EC2 Transit Gateway peering Attachment
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayPeeringAttachmentAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayPeeringAttachmentAccepterSpec   `json:"spec"`
	Status            TransitGatewayPeeringAttachmentAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayPeeringAttachmentAccepterList contains a list of TransitGatewayPeeringAttachmentAccepters
type TransitGatewayPeeringAttachmentAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayPeeringAttachmentAccepter `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayPeeringAttachmentAccepter_Kind             = "TransitGatewayPeeringAttachmentAccepter"
	TransitGatewayPeeringAttachmentAccepter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayPeeringAttachmentAccepter_Kind}.String()
	TransitGatewayPeeringAttachmentAccepter_KindAPIVersion   = TransitGatewayPeeringAttachmentAccepter_Kind + "." + CRDGroupVersion.String()
	TransitGatewayPeeringAttachmentAccepter_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayPeeringAttachmentAccepter_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayPeeringAttachmentAccepter{}, &TransitGatewayPeeringAttachmentAccepterList{})
}
