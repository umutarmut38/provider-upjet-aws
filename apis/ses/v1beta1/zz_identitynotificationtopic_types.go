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

type IdentityNotificationTopicObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
	Identity *string `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether SES should include original email headers in SNS notifications of this type. false by default.
	IncludeOriginalHeaders *bool `json:"includeOriginalHeaders,omitempty" tf:"include_original_headers,omitempty"`

	// The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: Bounce, Complaint or Delivery.
	NotificationType *string `json:"notificationType,omitempty" tf:"notification_type,omitempty"`

	// The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to "" (an empty string) to disable publishing.
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type IdentityNotificationTopicParameters struct {

	// The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ses/v1beta1.DomainIdentity
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("domain",false)
	// +kubebuilder:validation:Optional
	Identity *string `json:"identity,omitempty" tf:"identity,omitempty"`

	// Reference to a DomainIdentity in ses to populate identity.
	// +kubebuilder:validation:Optional
	IdentityRef *v1.Reference `json:"identityRef,omitempty" tf:"-"`

	// Selector for a DomainIdentity in ses to populate identity.
	// +kubebuilder:validation:Optional
	IdentitySelector *v1.Selector `json:"identitySelector,omitempty" tf:"-"`

	// Whether SES should include original email headers in SNS notifications of this type. false by default.
	// +kubebuilder:validation:Optional
	IncludeOriginalHeaders *bool `json:"includeOriginalHeaders,omitempty" tf:"include_original_headers,omitempty"`

	// The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: Bounce, Complaint or Delivery.
	// +kubebuilder:validation:Optional
	NotificationType *string `json:"notificationType,omitempty" tf:"notification_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to "" (an empty string) to disable publishing.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`

	// Reference to a Topic in sns to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnRef *v1.Reference `json:"topicArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnSelector *v1.Selector `json:"topicArnSelector,omitempty" tf:"-"`
}

// IdentityNotificationTopicSpec defines the desired state of IdentityNotificationTopic
type IdentityNotificationTopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdentityNotificationTopicParameters `json:"forProvider"`
}

// IdentityNotificationTopicStatus defines the observed state of IdentityNotificationTopic.
type IdentityNotificationTopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdentityNotificationTopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityNotificationTopic is the Schema for the IdentityNotificationTopics API. Setting AWS SES Identity Notification Topic
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IdentityNotificationTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.notificationType)",message="notificationType is a required parameter"
	Spec   IdentityNotificationTopicSpec   `json:"spec"`
	Status IdentityNotificationTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityNotificationTopicList contains a list of IdentityNotificationTopics
type IdentityNotificationTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityNotificationTopic `json:"items"`
}

// Repository type metadata.
var (
	IdentityNotificationTopic_Kind             = "IdentityNotificationTopic"
	IdentityNotificationTopic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityNotificationTopic_Kind}.String()
	IdentityNotificationTopic_KindAPIVersion   = IdentityNotificationTopic_Kind + "." + CRDGroupVersion.String()
	IdentityNotificationTopic_GroupVersionKind = CRDGroupVersion.WithKind(IdentityNotificationTopic_Kind)
)

func init() {
	SchemeBuilder.Register(&IdentityNotificationTopic{}, &IdentityNotificationTopicList{})
}
