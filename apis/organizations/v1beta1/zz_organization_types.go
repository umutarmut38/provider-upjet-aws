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

type AccountsObservation struct {

	// ARN of the account
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Email of the account
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Identifier of the account
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the account
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Current status of the account
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type AccountsParameters struct {
}

type NonMasterAccountsObservation struct {

	// ARN of the account
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Email of the account
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Identifier of the account
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the account
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Current status of the account
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type NonMasterAccountsParameters struct {
}

type OrganizationObservation struct {

	// List of organization accounts including the master account. For a list excluding the master account, see the non_master_accounts attribute. All elements have these attributes:
	Accounts []AccountsObservation `json:"accounts,omitempty" tf:"accounts,omitempty"`

	// ARN of the account
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have feature_set set to ALL. Some services do not support enablement via this endpoint, see warning in aws docs.
	AwsServiceAccessPrincipals []*string `json:"awsServiceAccessPrincipals,omitempty" tf:"aws_service_access_principals,omitempty"`

	// List of Organizations policy types to enable in the Organization Root. Organization must have feature_set set to ALL. For additional information about valid policy types (e.g., AISERVICES_OPT_OUT_POLICY, BACKUP_POLICY, SERVICE_CONTROL_POLICY, and TAG_POLICY), see the AWS Organizations API Reference.
	EnabledPolicyTypes []*string `json:"enabledPolicyTypes,omitempty" tf:"enabled_policy_types,omitempty"`

	// Specify "ALL" (default) or "CONSOLIDATED_BILLING".
	FeatureSet *string `json:"featureSet,omitempty" tf:"feature_set,omitempty"`

	// Identifier of the account
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of the master account
	MasterAccountArn *string `json:"masterAccountArn,omitempty" tf:"master_account_arn,omitempty"`

	// Email address of the master account
	MasterAccountEmail *string `json:"masterAccountEmail,omitempty" tf:"master_account_email,omitempty"`

	// Identifier of the master account
	MasterAccountID *string `json:"masterAccountId,omitempty" tf:"master_account_id,omitempty"`

	// List of organization accounts excluding the master account. For a list including the master account, see the accounts attribute. All elements have these attributes:
	NonMasterAccounts []NonMasterAccountsObservation `json:"nonMasterAccounts,omitempty" tf:"non_master_accounts,omitempty"`

	// List of organization roots. All elements have these attributes:
	Roots []RootsObservation `json:"roots,omitempty" tf:"roots,omitempty"`
}

type OrganizationParameters struct {

	// List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have feature_set set to ALL. Some services do not support enablement via this endpoint, see warning in aws docs.
	// +kubebuilder:validation:Optional
	AwsServiceAccessPrincipals []*string `json:"awsServiceAccessPrincipals,omitempty" tf:"aws_service_access_principals,omitempty"`

	// List of Organizations policy types to enable in the Organization Root. Organization must have feature_set set to ALL. For additional information about valid policy types (e.g., AISERVICES_OPT_OUT_POLICY, BACKUP_POLICY, SERVICE_CONTROL_POLICY, and TAG_POLICY), see the AWS Organizations API Reference.
	// +kubebuilder:validation:Optional
	EnabledPolicyTypes []*string `json:"enabledPolicyTypes,omitempty" tf:"enabled_policy_types,omitempty"`

	// Specify "ALL" (default) or "CONSOLIDATED_BILLING".
	// +kubebuilder:validation:Optional
	FeatureSet *string `json:"featureSet,omitempty" tf:"feature_set,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type PolicyTypesObservation struct {

	// Current status of the account
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PolicyTypesParameters struct {
}

type RootsObservation struct {

	// ARN of the account
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Identifier of the account
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the account
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// List of policy types enabled for this root. All elements have these attributes:
	PolicyTypes []PolicyTypesObservation `json:"policyTypes,omitempty" tf:"policy_types,omitempty"`
}

type RootsParameters struct {
}

// OrganizationSpec defines the desired state of Organization
type OrganizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationParameters `json:"forProvider"`
}

// OrganizationStatus defines the observed state of Organization.
type OrganizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Organization is the Schema for the Organizations API. Provides a resource to create an organization.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Organization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationSpec   `json:"spec"`
	Status            OrganizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationList contains a list of Organizations
type OrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Organization `json:"items"`
}

// Repository type metadata.
var (
	Organization_Kind             = "Organization"
	Organization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Organization_Kind}.String()
	Organization_KindAPIVersion   = Organization_Kind + "." + CRDGroupVersion.String()
	Organization_GroupVersionKind = CRDGroupVersion.WithKind(Organization_Kind)
)

func init() {
	SchemeBuilder.Register(&Organization{}, &OrganizationList{})
}
