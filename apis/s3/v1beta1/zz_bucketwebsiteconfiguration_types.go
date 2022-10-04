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

type BucketWebsiteConfigurationObservation struct {

	// The bucket or bucket and expected_bucket_owner separated by a comma (,) if the latter is provided.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The domain of the website endpoint. This is used to create Route 53 alias records.
	WebsiteDomain *string `json:"websiteDomain,omitempty" tf:"website_domain,omitempty"`

	// The website endpoint.
	WebsiteEndpoint *string `json:"websiteEndpoint,omitempty" tf:"website_endpoint,omitempty"`
}

type BucketWebsiteConfigurationParameters struct {

	// The name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The name of the error document for the website detailed below.
	// +kubebuilder:validation:Optional
	ErrorDocument []ErrorDocumentParameters `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// The account ID of the expected bucket owner.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// The name of the index document for the website detailed below.
	// +kubebuilder:validation:Optional
	IndexDocument []IndexDocumentParameters `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// The redirect behavior for every request to this bucket's website endpoint detailed below. Conflicts with error_document, index_document, and routing_rule.
	// +kubebuilder:validation:Optional
	RedirectAllRequestsTo []RedirectAllRequestsToParameters `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// List of rules that define when a redirect is applied and the redirect behavior detailed below.
	// +kubebuilder:validation:Optional
	RoutingRule []RoutingRuleParameters `json:"routingRule,omitempty" tf:"routing_rule,omitempty"`

	// A json array containing routing rules
	// describing redirect behavior and when redirects are applied. Use this parameter when your routing rules contain empty String values ("") as seen in the example above.
	// +kubebuilder:validation:Optional
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type ConditionObservation struct {
}

type ConditionParameters struct {

	// The HTTP error code when the redirect is applied. If specified with key_prefix_equals, then both must be true for the redirect to be applied.
	// +kubebuilder:validation:Optional
	HTTPErrorCodeReturnedEquals *string `json:"httpErrorCodeReturnedEquals,omitempty" tf:"http_error_code_returned_equals,omitempty"`

	// The object key name prefix when the redirect is applied. If specified with http_error_code_returned_equals, then both must be true for the redirect to be applied.
	// +kubebuilder:validation:Optional
	KeyPrefixEquals *string `json:"keyPrefixEquals,omitempty" tf:"key_prefix_equals,omitempty"`
}

type ErrorDocumentObservation struct {
}

type ErrorDocumentParameters struct {

	// The object key name to use when a 4XX class error occurs.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`
}

type IndexDocumentObservation struct {
}

type IndexDocumentParameters struct {

	// A suffix that is appended to a request that is for a directory on the website endpoint.
	// For example, if the suffix is index.html and you make a request to samplebucket/images/, the data that is returned will be for the object with the key name images/index.html.
	// The suffix must not be empty and must not include a slash character.
	// +kubebuilder:validation:Required
	Suffix *string `json:"suffix" tf:"suffix,omitempty"`
}

type RedirectAllRequestsToObservation struct {
}

type RedirectAllRequestsToParameters struct {

	// Name of the host where requests are redirected.
	// +kubebuilder:validation:Required
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request. Valid values: http, https.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type RedirectObservation struct {
}

type RedirectParameters struct {

	// The HTTP redirect code to use on the response.
	// +kubebuilder:validation:Optional
	HTTPRedirectCode *string `json:"httpRedirectCode,omitempty" tf:"http_redirect_code,omitempty"`

	// Name of the host where requests are redirected.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request. Valid values: http, https.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The object key prefix to use in the redirect request. For example, to redirect requests for all pages with prefix docs/ (objects in the docs/ folder) to documents/, you can set a condition block with key_prefix_equals set to docs/ and in the redirect set replace_key_prefix_with to /documents.
	// +kubebuilder:validation:Optional
	ReplaceKeyPrefixWith *string `json:"replaceKeyPrefixWith,omitempty" tf:"replace_key_prefix_with,omitempty"`

	// The specific object key to use in the redirect request. For example, redirect request to error.html.
	// +kubebuilder:validation:Optional
	ReplaceKeyWith *string `json:"replaceKeyWith,omitempty" tf:"replace_key_with,omitempty"`
}

type RoutingRuleObservation struct {
}

type RoutingRuleParameters struct {

	// A configuration block for describing a condition that must be met for the specified redirect to apply detailed below.
	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// A configuration block for redirect information detailed below.
	// +kubebuilder:validation:Required
	Redirect []RedirectParameters `json:"redirect" tf:"redirect,omitempty"`
}

// BucketWebsiteConfigurationSpec defines the desired state of BucketWebsiteConfiguration
type BucketWebsiteConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketWebsiteConfigurationParameters `json:"forProvider"`
}

// BucketWebsiteConfigurationStatus defines the observed state of BucketWebsiteConfiguration.
type BucketWebsiteConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketWebsiteConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketWebsiteConfiguration is the Schema for the BucketWebsiteConfigurations API. Provides an S3 bucket website configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketWebsiteConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketWebsiteConfigurationSpec   `json:"spec"`
	Status            BucketWebsiteConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketWebsiteConfigurationList contains a list of BucketWebsiteConfigurations
type BucketWebsiteConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketWebsiteConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketWebsiteConfiguration_Kind             = "BucketWebsiteConfiguration"
	BucketWebsiteConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketWebsiteConfiguration_Kind}.String()
	BucketWebsiteConfiguration_KindAPIVersion   = BucketWebsiteConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketWebsiteConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketWebsiteConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketWebsiteConfiguration{}, &BucketWebsiteConfigurationList{})
}
