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

type AccountTakeoverRiskConfigurationObservation struct {
}

type AccountTakeoverRiskConfigurationParameters struct {

	// Account takeover risk configuration actions. See details below.
	// +kubebuilder:validation:Required
	Actions []ActionsParameters `json:"actions" tf:"actions,omitempty"`

	// The notify configuration used to construct email notifications. See details below.
	// +kubebuilder:validation:Required
	NotifyConfiguration []NotifyConfigurationParameters `json:"notifyConfiguration" tf:"notify_configuration,omitempty"`
}

type ActionsObservation struct {
}

type ActionsParameters struct {

	// Action to take for a high risk. See action block below.
	// +kubebuilder:validation:Optional
	HighAction []HighActionParameters `json:"highAction,omitempty" tf:"high_action,omitempty"`

	// Action to take for a low risk. See action block below.
	// +kubebuilder:validation:Optional
	LowAction []LowActionParameters `json:"lowAction,omitempty" tf:"low_action,omitempty"`

	// Action to take for a medium risk. See action block below.
	// +kubebuilder:validation:Optional
	MediumAction []MediumActionParameters `json:"mediumAction,omitempty" tf:"medium_action,omitempty"`
}

type BlockEmailObservation struct {
}

type BlockEmailParameters struct {

	// The email HTML body.
	// +kubebuilder:validation:Required
	HTMLBody *string `json:"htmlBody" tf:"html_body,omitempty"`

	// The email subject.
	// +kubebuilder:validation:Required
	Subject *string `json:"subject" tf:"subject,omitempty"`

	// The email text body.
	// +kubebuilder:validation:Required
	TextBody *string `json:"textBody" tf:"text_body,omitempty"`
}

type CompromisedCredentialsRiskConfigurationActionsObservation struct {
}

type CompromisedCredentialsRiskConfigurationActionsParameters struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	// +kubebuilder:validation:Required
	EventAction *string `json:"eventAction" tf:"event_action,omitempty"`
}

type CompromisedCredentialsRiskConfigurationObservation struct {
}

type CompromisedCredentialsRiskConfigurationParameters struct {

	// The compromised credentials risk configuration actions. See details below.
	// +kubebuilder:validation:Required
	Actions []CompromisedCredentialsRiskConfigurationActionsParameters `json:"actions" tf:"actions,omitempty"`

	// Perform the action for these events. The default is to perform all events if no event filter is specified. Valid values are SIGN_IN, PASSWORD_CHANGE, and SIGN_UP.
	// +kubebuilder:validation:Optional
	EventFilter []*string `json:"eventFilter,omitempty" tf:"event_filter,omitempty"`
}

type HighActionObservation struct {
}

type HighActionParameters struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	// +kubebuilder:validation:Required
	EventAction *string `json:"eventAction" tf:"event_action,omitempty"`

	// Whether to send a notification.
	// +kubebuilder:validation:Required
	Notify *bool `json:"notify" tf:"notify,omitempty"`
}

type LowActionObservation struct {
}

type LowActionParameters struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	// +kubebuilder:validation:Required
	EventAction *string `json:"eventAction" tf:"event_action,omitempty"`

	// Whether to send a notification.
	// +kubebuilder:validation:Required
	Notify *bool `json:"notify" tf:"notify,omitempty"`
}

type MediumActionObservation struct {
}

type MediumActionParameters struct {

	// The action to take in response to the account takeover action. Valid values are BLOCK, MFA_IF_CONFIGURED, MFA_REQUIRED and NO_ACTION.
	// +kubebuilder:validation:Required
	EventAction *string `json:"eventAction" tf:"event_action,omitempty"`

	// Whether to send a notification.
	// +kubebuilder:validation:Required
	Notify *bool `json:"notify" tf:"notify,omitempty"`
}

type MfaEmailObservation struct {
}

type MfaEmailParameters struct {

	// The email HTML body.
	// +kubebuilder:validation:Required
	HTMLBody *string `json:"htmlBody" tf:"html_body,omitempty"`

	// The email subject.
	// +kubebuilder:validation:Required
	Subject *string `json:"subject" tf:"subject,omitempty"`

	// The email text body.
	// +kubebuilder:validation:Required
	TextBody *string `json:"textBody" tf:"text_body,omitempty"`
}

type NoActionEmailObservation struct {
}

type NoActionEmailParameters struct {

	// The email HTML body.
	// +kubebuilder:validation:Required
	HTMLBody *string `json:"htmlBody" tf:"html_body,omitempty"`

	// The email subject.
	// +kubebuilder:validation:Required
	Subject *string `json:"subject" tf:"subject,omitempty"`

	// The email text body.
	// +kubebuilder:validation:Required
	TextBody *string `json:"textBody" tf:"text_body,omitempty"`
}

type NotifyConfigurationObservation struct {
}

type NotifyConfigurationParameters struct {

	// Email template used when a detected risk event is blocked. See notify email type below.
	// +kubebuilder:validation:Optional
	BlockEmail []BlockEmailParameters `json:"blockEmail,omitempty" tf:"block_email,omitempty"`

	// The email address that is sending the email. The address must be either individually verified with Amazon Simple Email Service, or from a domain that has been verified with Amazon SES.
	// +kubebuilder:validation:Optional
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// The multi-factor authentication (MFA) email template used when MFA is challenged as part of a detected risk. See notify email type below.
	// +kubebuilder:validation:Optional
	MfaEmail []MfaEmailParameters `json:"mfaEmail,omitempty" tf:"mfa_email,omitempty"`

	// The email template used when a detected risk event is allowed. See notify email type below.
	// +kubebuilder:validation:Optional
	NoActionEmail []NoActionEmailParameters `json:"noActionEmail,omitempty" tf:"no_action_email,omitempty"`

	// The destination to which the receiver of an email should reply to.
	// +kubebuilder:validation:Optional
	ReplyTo *string `json:"replyTo,omitempty" tf:"reply_to,omitempty"`

	// The Amazon Resource Name (ARN) of the identity that is associated with the sending authorization policy. This identity permits Amazon Cognito to send for the email address specified in the From parameter.
	// +kubebuilder:validation:Required
	SourceArn *string `json:"sourceArn" tf:"source_arn,omitempty"`
}

type RiskConfigurationObservation struct {

	// The user pool ID. or The user pool ID and Client Id separated by a : if the configuration is client specific.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RiskConfigurationParameters struct {

	// The account takeover risk configuration. See details below.
	// +kubebuilder:validation:Optional
	AccountTakeoverRiskConfiguration []AccountTakeoverRiskConfigurationParameters `json:"accountTakeoverRiskConfiguration,omitempty" tf:"account_takeover_risk_configuration,omitempty"`

	// The app client ID. When the client ID is not provided, the same risk configuration is applied to all the clients in the User Pool.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The compromised credentials risk configuration. See details below.
	// +kubebuilder:validation:Optional
	CompromisedCredentialsRiskConfiguration []CompromisedCredentialsRiskConfigurationParameters `json:"compromisedCredentialsRiskConfiguration,omitempty" tf:"compromised_credentials_risk_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The configuration to override the risk decision. See details below.
	// +kubebuilder:validation:Optional
	RiskExceptionConfiguration []RiskExceptionConfigurationParameters `json:"riskExceptionConfiguration,omitempty" tf:"risk_exception_configuration,omitempty"`

	// The user pool ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPool
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// Reference to a UserPool in cognitoidp to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDRef *v1.Reference `json:"userPoolIdRef,omitempty" tf:"-"`

	// Selector for a UserPool in cognitoidp to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDSelector *v1.Selector `json:"userPoolIdSelector,omitempty" tf:"-"`
}

type RiskExceptionConfigurationObservation struct {
}

type RiskExceptionConfigurationParameters struct {

	// Overrides the risk decision to always block the pre-authentication requests. The IP range is in CIDR notation, a compact representation of an IP address and its routing prefix.
	// +kubebuilder:validation:Optional
	BlockedIPRangeList []*string `json:"blockedIpRangeList,omitempty" tf:"blocked_ip_range_list,omitempty"`

	// Risk detection isn't performed on the IP addresses in this range list. The IP range is in CIDR notation.
	// +kubebuilder:validation:Optional
	SkippedIPRangeList []*string `json:"skippedIpRangeList,omitempty" tf:"skipped_ip_range_list,omitempty"`
}

// RiskConfigurationSpec defines the desired state of RiskConfiguration
type RiskConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RiskConfigurationParameters `json:"forProvider"`
}

// RiskConfigurationStatus defines the observed state of RiskConfiguration.
type RiskConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RiskConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RiskConfiguration is the Schema for the RiskConfigurations API. Provides a Cognito Risk Configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RiskConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RiskConfigurationSpec   `json:"spec"`
	Status            RiskConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RiskConfigurationList contains a list of RiskConfigurations
type RiskConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RiskConfiguration `json:"items"`
}

// Repository type metadata.
var (
	RiskConfiguration_Kind             = "RiskConfiguration"
	RiskConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RiskConfiguration_Kind}.String()
	RiskConfiguration_KindAPIVersion   = RiskConfiguration_Kind + "." + CRDGroupVersion.String()
	RiskConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(RiskConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&RiskConfiguration{}, &RiskConfigurationList{})
}