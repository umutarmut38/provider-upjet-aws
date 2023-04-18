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

type SSHKeyObservation struct {

	// (Requirement) The public key portion of an SSH key pair.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Requirement) The Server ID of the Transfer Server (e.g., s-12345678)
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// (Requirement) The name of the user account that is assigned to one or more servers.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type SSHKeyParameters struct {

	// (Requirement) The public key portion of an SSH key pair.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// (Requirement) The Server ID of the Transfer Server (e.g., s-12345678)
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/transfer/v1beta1.Server
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a Server in transfer to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a Server in transfer to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// (Requirement) The name of the user account that is assigned to one or more servers.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/transfer/v1beta1.User
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// Reference to a User in transfer to populate userName.
	// +kubebuilder:validation:Optional
	UserNameRef *v1.Reference `json:"userNameRef,omitempty" tf:"-"`

	// Selector for a User in transfer to populate userName.
	// +kubebuilder:validation:Optional
	UserNameSelector *v1.Selector `json:"userNameSelector,omitempty" tf:"-"`
}

// SSHKeySpec defines the desired state of SSHKey
type SSHKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SSHKeyParameters `json:"forProvider"`
}

// SSHKeyStatus defines the observed state of SSHKey.
type SSHKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SSHKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SSHKey is the Schema for the SSHKeys API. Provides a AWS Transfer SSH Public Key resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SSHKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.body)",message="body is a required parameter"
	Spec   SSHKeySpec   `json:"spec"`
	Status SSHKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SSHKeyList contains a list of SSHKeys
type SSHKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSHKey `json:"items"`
}

// Repository type metadata.
var (
	SSHKey_Kind             = "SSHKey"
	SSHKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SSHKey_Kind}.String()
	SSHKey_KindAPIVersion   = SSHKey_Kind + "." + CRDGroupVersion.String()
	SSHKey_GroupVersionKind = CRDGroupVersion.WithKind(SSHKey_Kind)
)

func init() {
	SchemeBuilder.Register(&SSHKey{}, &SSHKeyList{})
}
