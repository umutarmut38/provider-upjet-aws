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

type AnalyticsConfigurationObservation struct {
}

type AnalyticsConfigurationParameters struct {

	// Application ARN for an Amazon Pinpoint application. Conflicts with external_id and role_arn.
	// +kubebuilder:validation:Optional
	ApplicationArn *string `json:"applicationArn,omitempty" tf:"application_arn,omitempty"`

	// Application ID for an Amazon Pinpoint application.
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// ID for the Analytics Configuration. Conflicts with application_arn.
	// +kubebuilder:validation:Optional
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// ARN of an IAM role that authorizes Amazon Cognito to publish events to Amazon Pinpoint analytics. Conflicts with application_arn.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// If set to true, Amazon Cognito will include user data in the events it publishes to Amazon Pinpoint analytics.
	// +kubebuilder:validation:Optional
	UserDataShared *bool `json:"userDataShared,omitempty" tf:"user_data_shared,omitempty"`
}

type TokenValidityUnitsObservation struct {
}

type TokenValidityUnitsParameters struct {

	// Time unit in for the value in access_token_validity, defaults to hours.
	// +kubebuilder:validation:Optional
	AccessToken *string `json:"accessToken,omitempty" tf:"access_token,omitempty"`

	// Time unit in for the value in id_token_validity, defaults to hours.
	// +kubebuilder:validation:Optional
	IDToken *string `json:"idToken,omitempty" tf:"id_token,omitempty"`

	// Time unit in for the value in refresh_token_validity, defaults to days.
	// +kubebuilder:validation:Optional
	RefreshToken *string `json:"refreshToken,omitempty" tf:"refresh_token,omitempty"`
}

type UserPoolClientObservation struct {

	// ID of the user pool client.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type UserPoolClientParameters struct {

	// Time limit, between 5 minutes and 1 day, after which the access token is no longer valid and cannot be used. This value will be overridden if you have entered a value in token_validity_units.
	// +kubebuilder:validation:Optional
	AccessTokenValidity *float64 `json:"accessTokenValidity,omitempty" tf:"access_token_validity,omitempty"`

	// List of allowed OAuth flows (code, implicit, client_credentials).
	// +kubebuilder:validation:Optional
	AllowedOauthFlows []*string `json:"allowedOauthFlows,omitempty" tf:"allowed_oauth_flows,omitempty"`

	// Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
	// +kubebuilder:validation:Optional
	AllowedOauthFlowsUserPoolClient *bool `json:"allowedOauthFlowsUserPoolClient,omitempty" tf:"allowed_oauth_flows_user_pool_client,omitempty"`

	// List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
	// +kubebuilder:validation:Optional
	AllowedOauthScopes []*string `json:"allowedOauthScopes,omitempty" tf:"allowed_oauth_scopes,omitempty"`

	// Configuration block for Amazon Pinpoint analytics for collecting metrics for this user pool. Detailed below.
	// +kubebuilder:validation:Optional
	AnalyticsConfiguration []AnalyticsConfigurationParameters `json:"analyticsConfiguration,omitempty" tf:"analytics_configuration,omitempty"`

	// List of allowed callback URLs for the identity providers.
	// +kubebuilder:validation:Optional
	CallbackUrls []*string `json:"callbackUrls,omitempty" tf:"callback_urls,omitempty"`

	// Default redirect URI. Must be in the list of callback URLs.
	// +kubebuilder:validation:Optional
	DefaultRedirectURI *string `json:"defaultRedirectUri,omitempty" tf:"default_redirect_uri,omitempty"`

	// Enables or disables token revocation.
	// +kubebuilder:validation:Optional
	EnableTokenRevocation *bool `json:"enableTokenRevocation,omitempty" tf:"enable_token_revocation,omitempty"`

	// List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY, USER_PASSWORD_AUTH, ALLOW_ADMIN_USER_PASSWORD_AUTH, ALLOW_CUSTOM_AUTH, ALLOW_USER_PASSWORD_AUTH, ALLOW_USER_SRP_AUTH, ALLOW_REFRESH_TOKEN_AUTH).
	// +kubebuilder:validation:Optional
	ExplicitAuthFlows []*string `json:"explicitAuthFlows,omitempty" tf:"explicit_auth_flows,omitempty"`

	// Should an application secret be generated.
	// +kubebuilder:validation:Optional
	GenerateSecret *bool `json:"generateSecret,omitempty" tf:"generate_secret,omitempty"`

	// Time limit, between 5 minutes and 1 day, after which the ID token is no longer valid and cannot be used. This value will be overridden if you have entered a value in token_validity_units.
	// +kubebuilder:validation:Optional
	IDTokenValidity *float64 `json:"idTokenValidity,omitempty" tf:"id_token_validity,omitempty"`

	// List of allowed logout URLs for the identity providers.
	// +kubebuilder:validation:Optional
	LogoutUrls []*string `json:"logoutUrls,omitempty" tf:"logout_urls,omitempty"`

	// Name of the application client.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Choose which errors and responses are returned by Cognito APIs during authentication, account confirmation, and password recovery when the user does not exist in the user pool. When set to ENABLED and the user does not exist, authentication returns an error indicating either the username or password was incorrect, and account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to LEGACY, those APIs will return a UserNotFoundException exception if the user does not exist in the user pool.
	// +kubebuilder:validation:Optional
	PreventUserExistenceErrors *string `json:"preventUserExistenceErrors,omitempty" tf:"prevent_user_existence_errors,omitempty"`

	// List of user pool attributes the application client can read from.
	// +kubebuilder:validation:Optional
	ReadAttributes []*string `json:"readAttributes,omitempty" tf:"read_attributes,omitempty"`

	// Time limit in days refresh tokens are valid for.
	// +kubebuilder:validation:Optional
	RefreshTokenValidity *float64 `json:"refreshTokenValidity,omitempty" tf:"refresh_token_validity,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// List of provider names for the identity providers that are supported on this client. Uses the provider_name attribute of aws_cognito_identity_provider resource(s), or the equivalent string(s).
	// +kubebuilder:validation:Optional
	SupportedIdentityProviders []*string `json:"supportedIdentityProviders,omitempty" tf:"supported_identity_providers,omitempty"`

	// Configuration block for units in which the validity times are represented in. Detailed below.
	// +kubebuilder:validation:Optional
	TokenValidityUnits []TokenValidityUnitsParameters `json:"tokenValidityUnits,omitempty" tf:"token_validity_units,omitempty"`

	// User pool the client belongs to.
	// +crossplane:generate:reference:type=UserPool
	// +kubebuilder:validation:Optional
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// Reference to a UserPool to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDRef *v1.Reference `json:"userPoolIdRef,omitempty" tf:"-"`

	// Selector for a UserPool to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDSelector *v1.Selector `json:"userPoolIdSelector,omitempty" tf:"-"`

	// List of user pool attributes the application client can write to.
	// +kubebuilder:validation:Optional
	WriteAttributes []*string `json:"writeAttributes,omitempty" tf:"write_attributes,omitempty"`
}

// UserPoolClientSpec defines the desired state of UserPoolClient
type UserPoolClientSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserPoolClientParameters `json:"forProvider"`
}

// UserPoolClientStatus defines the observed state of UserPoolClient.
type UserPoolClientStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserPoolClientObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserPoolClient is the Schema for the UserPoolClients API. Provides a Cognito User Pool Client resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserPoolClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserPoolClientSpec   `json:"spec"`
	Status            UserPoolClientStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserPoolClientList contains a list of UserPoolClients
type UserPoolClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserPoolClient `json:"items"`
}

// Repository type metadata.
var (
	UserPoolClient_Kind             = "UserPoolClient"
	UserPoolClient_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserPoolClient_Kind}.String()
	UserPoolClient_KindAPIVersion   = UserPoolClient_Kind + "." + CRDGroupVersion.String()
	UserPoolClient_GroupVersionKind = CRDGroupVersion.WithKind(UserPoolClient_Kind)
)

func init() {
	SchemeBuilder.Register(&UserPoolClient{}, &UserPoolClientList{})
}
