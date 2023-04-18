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

type AccessKeyObservation struct {

	// Date and time in RFC3339 format that the access key was created.
	CreateDate *string `json:"createDate,omitempty" tf:"create_date,omitempty"`

	// Encrypted secret, base64 encoded, if pgp_key was specified. This attribute is not available for imported resources.
	EncryptedSecret *string `json:"encryptedSecret,omitempty" tf:"encrypted_secret,omitempty"`

	// Encrypted SES SMTP password, base64 encoded, if pgp_key was specified. This attribute is not available for imported resources.
	EncryptedSesSMTPPasswordV4 *string `json:"encryptedSesSmtpPasswordV4,omitempty" tf:"encrypted_ses_smtp_password_v4,omitempty"`

	// Access key ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Fingerprint of the PGP key used to encrypt the secret. This attribute is not available for imported resources.
	KeyFingerprint *string `json:"keyFingerprint,omitempty" tf:"key_fingerprint,omitempty"`

	// Either a base-64 encoded PGP public key, or a keybase username in the form keybase:some_person_that_exists, for use in the encrypted_secret output attribute. If providing a base-64 encoded PGP public key, make sure to provide the "raw" version and not the "armored" one (e.g. avoid passing the -a option to gpg --export).
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// Access key status to apply. Defaults to Active. Valid values are Active and Inactive.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// IAM user to associate with this access key.
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type AccessKeyParameters struct {

	// Either a base-64 encoded PGP public key, or a keybase username in the form keybase:some_person_that_exists, for use in the encrypted_secret output attribute. If providing a base-64 encoded PGP public key, make sure to provide the "raw" version and not the "armored" one (e.g. avoid passing the -a option to gpg --export).
	// +kubebuilder:validation:Optional
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// Access key status to apply. Defaults to Active. Valid values are Active and Inactive.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// IAM user to associate with this access key.
	// +crossplane:generate:reference:type=User
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`

	// Reference to a User to populate user.
	// +kubebuilder:validation:Optional
	UserRef *v1.Reference `json:"userRef,omitempty" tf:"-"`

	// Selector for a User to populate user.
	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`
}

// AccessKeySpec defines the desired state of AccessKey
type AccessKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessKeyParameters `json:"forProvider"`
}

// AccessKeyStatus defines the observed state of AccessKey.
type AccessKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccessKey is the Schema for the AccessKeys API. Provides an IAM access key. This is a set of credentials that allow API requests to be made as an IAM user.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AccessKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccessKeySpec   `json:"spec"`
	Status            AccessKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessKeyList contains a list of AccessKeys
type AccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessKey `json:"items"`
}

// Repository type metadata.
var (
	AccessKey_Kind             = "AccessKey"
	AccessKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessKey_Kind}.String()
	AccessKey_KindAPIVersion   = AccessKey_Kind + "." + CRDGroupVersion.String()
	AccessKey_GroupVersionKind = CRDGroupVersion.WithKind(AccessKey_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessKey{}, &AccessKeyList{})
}
