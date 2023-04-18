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

type AliasObservation struct {

	// The Amazon Resource Name (ARN) identifying your Lambda function alias.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of the alias.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Lambda Function name or ARN.
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// Lambda function version for which you are creating the alias. Pattern: (\$LATEST|[0-9]+).
	FunctionVersion *string `json:"functionVersion,omitempty" tf:"function_version,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN to be used for invoking Lambda Function from API Gateway - to be used in aws_api_gateway_integration's uri
	InvokeArn *string `json:"invokeArn,omitempty" tf:"invoke_arn,omitempty"`

	// The Lambda alias' route configuration settings. Fields documented below
	RoutingConfig []RoutingConfigObservation `json:"routingConfig,omitempty" tf:"routing_config,omitempty"`
}

type AliasParameters struct {

	// Description of the alias.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Lambda Function name or ARN.
	// +crossplane:generate:reference:type=Function
	// +kubebuilder:validation:Optional
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// Reference to a Function to populate functionName.
	// +kubebuilder:validation:Optional
	FunctionNameRef *v1.Reference `json:"functionNameRef,omitempty" tf:"-"`

	// Selector for a Function to populate functionName.
	// +kubebuilder:validation:Optional
	FunctionNameSelector *v1.Selector `json:"functionNameSelector,omitempty" tf:"-"`

	// Lambda function version for which you are creating the alias. Pattern: (\$LATEST|[0-9]+).
	// +kubebuilder:validation:Optional
	FunctionVersion *string `json:"functionVersion,omitempty" tf:"function_version,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Lambda alias' route configuration settings. Fields documented below
	// +kubebuilder:validation:Optional
	RoutingConfig []RoutingConfigParameters `json:"routingConfig,omitempty" tf:"routing_config,omitempty"`
}

type RoutingConfigObservation struct {

	// A map that defines the proportion of events that should be sent to different versions of a lambda function.
	AdditionalVersionWeights map[string]*float64 `json:"additionalVersionWeights,omitempty" tf:"additional_version_weights,omitempty"`
}

type RoutingConfigParameters struct {

	// A map that defines the proportion of events that should be sent to different versions of a lambda function.
	// +kubebuilder:validation:Optional
	AdditionalVersionWeights map[string]*float64 `json:"additionalVersionWeights,omitempty" tf:"additional_version_weights,omitempty"`
}

// AliasSpec defines the desired state of Alias
type AliasSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AliasParameters `json:"forProvider"`
}

// AliasStatus defines the observed state of Alias.
type AliasStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Alias is the Schema for the Aliass API. Creates a Lambda function alias.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Alias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.functionVersion)",message="functionVersion is a required parameter"
	Spec   AliasSpec   `json:"spec"`
	Status AliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AliasList contains a list of Aliass
type AliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alias `json:"items"`
}

// Repository type metadata.
var (
	Alias_Kind             = "Alias"
	Alias_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Alias_Kind}.String()
	Alias_KindAPIVersion   = Alias_Kind + "." + CRDGroupVersion.String()
	Alias_GroupVersionKind = CRDGroupVersion.WithKind(Alias_Kind)
)

func init() {
	SchemeBuilder.Register(&Alias{}, &AliasList{})
}
