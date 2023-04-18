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

type UserPoolUICustomizationObservation struct {

	// The CSS values in the UI customization, provided as a String. At least one of css or image_file is required.
	CSS *string `json:"css,omitempty" tf:"css,omitempty"`

	// The CSS version number.
	CSSVersion *string `json:"cssVersion,omitempty" tf:"css_version,omitempty"`

	// The client ID for the client app. Defaults to ALL. If ALL is specified, the css and/or image_file settings will be used for every client that has no UI customization set previously.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The creation date in RFC3339 format for the UI customization.
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The uploaded logo image for the UI customization, provided as a base64-encoded String. Drift detection is not possible for this argument. At least one of css or image_file is required.
	ImageFile *string `json:"imageFile,omitempty" tf:"image_file,omitempty"`

	// The logo image URL for the UI customization.
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// The last-modified date in RFC3339 format for the UI customization.
	LastModifiedDate *string `json:"lastModifiedDate,omitempty" tf:"last_modified_date,omitempty"`

	// The user pool ID for the user pool.
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`
}

type UserPoolUICustomizationParameters struct {

	// The CSS values in the UI customization, provided as a String. At least one of css or image_file is required.
	// +kubebuilder:validation:Optional
	CSS *string `json:"css,omitempty" tf:"css,omitempty"`

	// The client ID for the client app. Defaults to ALL. If ALL is specified, the css and/or image_file settings will be used for every client that has no UI customization set previously.
	// +crossplane:generate:reference:type=UserPoolClient
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a UserPoolClient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a UserPoolClient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// The uploaded logo image for the UI customization, provided as a base64-encoded String. Drift detection is not possible for this argument. At least one of css or image_file is required.
	// +kubebuilder:validation:Optional
	ImageFile *string `json:"imageFile,omitempty" tf:"image_file,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The user pool ID for the user pool.
	// +crossplane:generate:reference:type=UserPool
	// +kubebuilder:validation:Optional
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// Reference to a UserPool to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDRef *v1.Reference `json:"userPoolIdRef,omitempty" tf:"-"`

	// Selector for a UserPool to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDSelector *v1.Selector `json:"userPoolIdSelector,omitempty" tf:"-"`
}

// UserPoolUICustomizationSpec defines the desired state of UserPoolUICustomization
type UserPoolUICustomizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserPoolUICustomizationParameters `json:"forProvider"`
}

// UserPoolUICustomizationStatus defines the observed state of UserPoolUICustomization.
type UserPoolUICustomizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserPoolUICustomizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserPoolUICustomization is the Schema for the UserPoolUICustomizations API. Provides a Cognito User Pool UI Customization resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserPoolUICustomization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserPoolUICustomizationSpec   `json:"spec"`
	Status            UserPoolUICustomizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserPoolUICustomizationList contains a list of UserPoolUICustomizations
type UserPoolUICustomizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserPoolUICustomization `json:"items"`
}

// Repository type metadata.
var (
	UserPoolUICustomization_Kind             = "UserPoolUICustomization"
	UserPoolUICustomization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserPoolUICustomization_Kind}.String()
	UserPoolUICustomization_KindAPIVersion   = UserPoolUICustomization_Kind + "." + CRDGroupVersion.String()
	UserPoolUICustomization_GroupVersionKind = CRDGroupVersion.WithKind(UserPoolUICustomization_Kind)
)

func init() {
	SchemeBuilder.Register(&UserPoolUICustomization{}, &UserPoolUICustomizationList{})
}
