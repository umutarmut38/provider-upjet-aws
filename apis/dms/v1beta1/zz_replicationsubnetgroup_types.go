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

type ReplicationSubnetGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ReplicationSubnetGroupArn *string `json:"replicationSubnetGroupArn,omitempty" tf:"replication_subnet_group_arn,omitempty"`

	// Description for the subnet group.
	ReplicationSubnetGroupDescription *string `json:"replicationSubnetGroupDescription,omitempty" tf:"replication_subnet_group_description,omitempty"`

	// List of at least 2 EC2 subnet IDs for the subnet group. The subnets must cover at least 2 availability zones.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ID of the VPC the subnet group is in.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type ReplicationSubnetGroupParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Description for the subnet group.
	// +kubebuilder:validation:Optional
	ReplicationSubnetGroupDescription *string `json:"replicationSubnetGroupDescription,omitempty" tf:"replication_subnet_group_description,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// List of at least 2 EC2 subnet IDs for the subnet group. The subnets must cover at least 2 availability zones.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ReplicationSubnetGroupSpec defines the desired state of ReplicationSubnetGroup
type ReplicationSubnetGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReplicationSubnetGroupParameters `json:"forProvider"`
}

// ReplicationSubnetGroupStatus defines the observed state of ReplicationSubnetGroup.
type ReplicationSubnetGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReplicationSubnetGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationSubnetGroup is the Schema for the ReplicationSubnetGroups API. Provides a DMS (Data Migration Service) subnet group resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ReplicationSubnetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.replicationSubnetGroupDescription)",message="replicationSubnetGroupDescription is a required parameter"
	Spec   ReplicationSubnetGroupSpec   `json:"spec"`
	Status ReplicationSubnetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationSubnetGroupList contains a list of ReplicationSubnetGroups
type ReplicationSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicationSubnetGroup `json:"items"`
}

// Repository type metadata.
var (
	ReplicationSubnetGroup_Kind             = "ReplicationSubnetGroup"
	ReplicationSubnetGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReplicationSubnetGroup_Kind}.String()
	ReplicationSubnetGroup_KindAPIVersion   = ReplicationSubnetGroup_Kind + "." + CRDGroupVersion.String()
	ReplicationSubnetGroup_GroupVersionKind = CRDGroupVersion.WithKind(ReplicationSubnetGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ReplicationSubnetGroup{}, &ReplicationSubnetGroupList{})
}
