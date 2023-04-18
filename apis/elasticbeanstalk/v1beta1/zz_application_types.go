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

type ApplicationObservation struct {
	AppversionLifecycle []AppversionLifecycleObservation `json:"appversionLifecycle,omitempty" tf:"appversion_lifecycle,omitempty"`

	// The ARN assigned by AWS for this Elastic Beanstalk Application.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Short description of the application
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ApplicationParameters struct {

	// +kubebuilder:validation:Optional
	AppversionLifecycle []AppversionLifecycleParameters `json:"appversionLifecycle,omitempty" tf:"appversion_lifecycle,omitempty"`

	// Short description of the application
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AppversionLifecycleObservation struct {

	// Set to true to delete a version's source bundle from S3 when the application version is deleted.
	DeleteSourceFromS3 *bool `json:"deleteSourceFromS3,omitempty" tf:"delete_source_from_s3,omitempty"`

	// The number of days to retain an application version ('max_age_in_days' and 'max_count' cannot be enabled simultaneously.).
	MaxAgeInDays *float64 `json:"maxAgeInDays,omitempty" tf:"max_age_in_days,omitempty"`

	// The maximum number of application versions to retain ('max_age_in_days' and 'max_count' cannot be enabled simultaneously.).
	MaxCount *float64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`

	// The ARN of an IAM service role under which the application version is deleted.  Elastic Beanstalk must have permission to assume this role.
	ServiceRole *string `json:"serviceRole,omitempty" tf:"service_role,omitempty"`
}

type AppversionLifecycleParameters struct {

	// Set to true to delete a version's source bundle from S3 when the application version is deleted.
	// +kubebuilder:validation:Optional
	DeleteSourceFromS3 *bool `json:"deleteSourceFromS3,omitempty" tf:"delete_source_from_s3,omitempty"`

	// The number of days to retain an application version ('max_age_in_days' and 'max_count' cannot be enabled simultaneously.).
	// +kubebuilder:validation:Optional
	MaxAgeInDays *float64 `json:"maxAgeInDays,omitempty" tf:"max_age_in_days,omitempty"`

	// The maximum number of application versions to retain ('max_age_in_days' and 'max_count' cannot be enabled simultaneously.).
	// +kubebuilder:validation:Optional
	MaxCount *float64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`

	// The ARN of an IAM service role under which the application version is deleted.  Elastic Beanstalk must have permission to assume this role.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ServiceRole *string `json:"serviceRole,omitempty" tf:"service_role,omitempty"`

	// Reference to a Role in iam to populate serviceRole.
	// +kubebuilder:validation:Optional
	ServiceRoleRef *v1.Reference `json:"serviceRoleRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate serviceRole.
	// +kubebuilder:validation:Optional
	ServiceRoleSelector *v1.Selector `json:"serviceRoleSelector,omitempty" tf:"-"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Application is the Schema for the Applications API. Provides an Elastic Beanstalk Application Resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSpec   `json:"spec"`
	Status            ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
