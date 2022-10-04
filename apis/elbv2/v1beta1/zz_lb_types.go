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

type AccessLogsObservation struct {
}

type AccessLogsParameters struct {

	// The S3 bucket name to store the logs in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Boolean to enable / disable access_logs. Defaults to false, even when bucket is specified.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The S3 bucket prefix. Logs are stored in the root if not configured.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type LBObservation struct {

	// The ARN of the load balancer (matches id).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ARN suffix for use with CloudWatch Metrics.
	ArnSuffix *string `json:"arnSuffix,omitempty" tf:"arn_suffix,omitempty"`

	// The DNS name of the load balancer.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// The ARN of the load balancer (matches arn).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A subnet mapping block as documented below.
	// +kubebuilder:validation:Optional
	SubnetMapping []SubnetMappingObservation `json:"subnetMapping,omitempty" tf:"subnet_mapping,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ARN of the load balancer (matches arn).
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type LBParameters struct {

	// An Access Logs block. Access Logs documented below.
	// +kubebuilder:validation:Optional
	AccessLogs []AccessLogsParameters `json:"accessLogs,omitempty" tf:"access_logs,omitempty"`

	// The ID of the customer owned ipv4 pool to use for this load balancer.
	// +kubebuilder:validation:Optional
	CustomerOwnedIPv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool,omitempty"`

	// Determines how the load balancer handles requests that might pose a security risk to an application due to HTTP desync. Valid values are monitor, defensive (default), strictest.
	// +kubebuilder:validation:Optional
	DesyncMitigationMode *string `json:"desyncMitigationMode,omitempty" tf:"desync_mitigation_mode,omitempty"`

	// Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer (true) or routed to targets (false). The default is false. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens. Only valid for Load Balancers of type application.
	// +kubebuilder:validation:Optional
	DropInvalidHeaderFields *bool `json:"dropInvalidHeaderFields,omitempty" tf:"drop_invalid_header_fields,omitempty"`

	// If true, cross-zone load balancing of the load balancer will be enabled.
	// This is a network load balancer feature. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableCrossZoneLoadBalancing *bool `json:"enableCrossZoneLoadBalancing,omitempty" tf:"enable_cross_zone_load_balancing,omitempty"`

	// If true, deletion of the load balancer will be disabled via
	// the AWS API. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableDeletionProtection *bool `json:"enableDeletionProtection,omitempty" tf:"enable_deletion_protection,omitempty"`

	// Indicates whether HTTP/2 is enabled in application load balancers. Defaults to true.
	// +kubebuilder:validation:Optional
	EnableHttp2 *bool `json:"enableHttp2,omitempty" tf:"enable_http2,omitempty"`

	// Indicates whether to allow a WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableWafFailOpen *bool `json:"enableWafFailOpen,omitempty" tf:"enable_waf_fail_open,omitempty"`

	// The type of IP addresses used by the subnets for your load balancer. The possible values are ipv4 and dualstack
	// +kubebuilder:validation:Optional
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type application. Default: 60.
	// +kubebuilder:validation:Optional
	IdleTimeout *float64 `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`

	// If true, the LB will be internal.
	// +kubebuilder:validation:Optional
	Internal *bool `json:"internal,omitempty" tf:"internal,omitempty"`

	// The type of load balancer to create. Possible values are application, gateway, or network. The default value is application.
	// +kubebuilder:validation:Optional
	LoadBalancerType *string `json:"loadBalancerType,omitempty" tf:"load_balancer_type,omitempty"`

	// The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// References to SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupRefs []v1.Reference `json:"securityGroupRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupSelector *v1.Selector `json:"securityGroupSelector,omitempty" tf:"-"`

	// A list of security group IDs to assign to the LB. Only valid for Load Balancers of type application.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupSelector
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// A subnet mapping block as documented below.
	// +kubebuilder:validation:Optional
	SubnetMapping []SubnetMappingParameters `json:"subnetMapping,omitempty" tf:"subnet_mapping,omitempty"`

	// References to Subnet in ec2 to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetRefs []v1.Reference `json:"subnetRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetSelector *v1.Selector `json:"subnetSelector,omitempty" tf:"-"`

	// A list of subnet IDs to attach to the LB. Subnets
	// cannot be updated for Load Balancers of type network. Changing this value
	// for load balancers of type network will force a recreation of the resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetSelector
	// +kubebuilder:validation:Optional
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SubnetMappingObservation struct {

	// ID of the Outpost containing the load balancer.
	OutpostID *string `json:"outpostId,omitempty" tf:"outpost_id,omitempty"`
}

type SubnetMappingParameters struct {

	// The allocation ID of the Elastic IP address.
	// +kubebuilder:validation:Optional
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// An ipv6 address within the subnet to assign to the internet-facing load balancer.
	// +kubebuilder:validation:Optional
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`

	// A private ipv4 address within the subnet to assign to the internal-facing load balancer.
	// +kubebuilder:validation:Optional
	PrivateIPv4Address *string `json:"privateIpv4Address,omitempty" tf:"private_ipv4_address,omitempty"`

	// The id of the subnet of which to attach to the load balancer. You can specify only one subnet per Availability Zone.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// LBSpec defines the desired state of LB
type LBSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBParameters `json:"forProvider"`
}

// LBStatus defines the observed state of LB.
type LBStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LB is the Schema for the LBs API. Provides a Load Balancer resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LBSpec   `json:"spec"`
	Status            LBStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBList contains a list of LBs
type LBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LB `json:"items"`
}

// Repository type metadata.
var (
	LB_Kind             = "LB"
	LB_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LB_Kind}.String()
	LB_KindAPIVersion   = LB_Kind + "." + CRDGroupVersion.String()
	LB_GroupVersionKind = CRDGroupVersion.WithKind(LB_Kind)
)

func init() {
	SchemeBuilder.Register(&LB{}, &LBList{})
}
