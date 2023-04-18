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

type DestinationObservation struct {

	// A Region to replicate to.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The account ID of the destination registry to replicate to.
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`
}

type DestinationParameters struct {

	// A Region to replicate to.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The account ID of the destination registry to replicate to.
	// +kubebuilder:validation:Required
	RegistryID *string `json:"registryId" tf:"registry_id,omitempty"`
}

type ReplicationConfigurationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The account ID of the destination registry to replicate to.
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// Replication configuration for a registry. See Replication Configuration.
	ReplicationConfiguration []ReplicationConfigurationReplicationConfigurationObservation `json:"replicationConfiguration,omitempty" tf:"replication_configuration,omitempty"`
}

type ReplicationConfigurationParameters struct {

	// A Region to replicate to.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Replication configuration for a registry. See Replication Configuration.
	// +kubebuilder:validation:Optional
	ReplicationConfiguration []ReplicationConfigurationReplicationConfigurationParameters `json:"replicationConfiguration,omitempty" tf:"replication_configuration,omitempty"`
}

type ReplicationConfigurationReplicationConfigurationObservation struct {

	// The replication rules for a replication configuration. A maximum of 10 are allowed per replication_configuration. See Rule
	Rule []ReplicationConfigurationRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type ReplicationConfigurationReplicationConfigurationParameters struct {

	// The replication rules for a replication configuration. A maximum of 10 are allowed per replication_configuration. See Rule
	// +kubebuilder:validation:Required
	Rule []ReplicationConfigurationRuleParameters `json:"rule" tf:"rule,omitempty"`
}

type ReplicationConfigurationRuleObservation struct {

	// the details of a replication destination. A maximum of 25 are allowed per rule. See Destination.
	Destination []DestinationObservation `json:"destination,omitempty" tf:"destination,omitempty"`

	// filters for a replication rule. See Repository Filter.
	RepositoryFilter []RuleRepositoryFilterObservation `json:"repositoryFilter,omitempty" tf:"repository_filter,omitempty"`
}

type ReplicationConfigurationRuleParameters struct {

	// the details of a replication destination. A maximum of 25 are allowed per rule. See Destination.
	// +kubebuilder:validation:Required
	Destination []DestinationParameters `json:"destination" tf:"destination,omitempty"`

	// filters for a replication rule. See Repository Filter.
	// +kubebuilder:validation:Optional
	RepositoryFilter []RuleRepositoryFilterParameters `json:"repositoryFilter,omitempty" tf:"repository_filter,omitempty"`
}

type RuleRepositoryFilterObservation struct {

	// The repository filter details.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// The repository filter type. The only supported value is PREFIX_MATCH, which is a repository name prefix specified with the filter parameter.
	FilterType *string `json:"filterType,omitempty" tf:"filter_type,omitempty"`
}

type RuleRepositoryFilterParameters struct {

	// The repository filter details.
	// +kubebuilder:validation:Required
	Filter *string `json:"filter" tf:"filter,omitempty"`

	// The repository filter type. The only supported value is PREFIX_MATCH, which is a repository name prefix specified with the filter parameter.
	// +kubebuilder:validation:Required
	FilterType *string `json:"filterType" tf:"filter_type,omitempty"`
}

// ReplicationConfigurationSpec defines the desired state of ReplicationConfiguration
type ReplicationConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReplicationConfigurationParameters `json:"forProvider"`
}

// ReplicationConfigurationStatus defines the observed state of ReplicationConfiguration.
type ReplicationConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReplicationConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationConfiguration is the Schema for the ReplicationConfigurations API. Provides an Elastic Container Registry Replication Configuration.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ReplicationConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReplicationConfigurationSpec   `json:"spec"`
	Status            ReplicationConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationConfigurationList contains a list of ReplicationConfigurations
type ReplicationConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicationConfiguration `json:"items"`
}

// Repository type metadata.
var (
	ReplicationConfiguration_Kind             = "ReplicationConfiguration"
	ReplicationConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReplicationConfiguration_Kind}.String()
	ReplicationConfiguration_KindAPIVersion   = ReplicationConfiguration_Kind + "." + CRDGroupVersion.String()
	ReplicationConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(ReplicationConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&ReplicationConfiguration{}, &ReplicationConfigurationList{})
}
