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

type ClusterInstanceObservation struct {

	// Amazon Resource Name (ARN) of cluster instance
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The region-unique, immutable identifier for the DB instance.
	DbiResourceID *string `json:"dbiResourceId,omitempty" tf:"dbi_resource_id,omitempty"`

	// The DNS address for this instance. May not be writable
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The database engine version
	EngineVersionActual *string `json:"engineVersionActual,omitempty" tf:"engine_version_actual,omitempty"`

	// The Instance identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN for the KMS encryption key if one is set to the cluster.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The database port
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// – Boolean indicating if this instance is writable. False indicates this instance is a read replica.
	Writer *bool `json:"writer,omitempty" tf:"writer,omitempty"`
}

type ClusterInstanceParameters struct {

	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default isfalse.
	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default true.
	// +kubebuilder:validation:Optional
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// The EC2 Availability Zone that the DB instance is created in. See docs about the details.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The identifier of the CA certificate for the DB instance.
	// +kubebuilder:validation:Optional
	CACertIdentifier *string `json:"caCertIdentifier,omitempty" tf:"ca_cert_identifier,omitempty"`

	// The identifier of the aws_rds_cluster in which to launch this instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// Reference to a Cluster in rds to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierRef *v1.Reference `json:"clusterIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster in rds to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierSelector *v1.Selector `json:"clusterIdentifierSelector,omitempty" tf:"-"`

	// defined tags from the DB instance to snapshots of the DB instance. Default false.
	// +kubebuilder:validation:Optional
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`

	// The name of the DB parameter group to associate with this instance.
	// +kubebuilder:validation:Optional
	DBParameterGroupName *string `json:"dbParameterGroupName,omitempty" tf:"db_parameter_group_name,omitempty"`

	// A DB subnet group to associate with this DB instance. NOTE: This must match the db_subnet_group_name of the attached aws_rds_cluster.
	// +crossplane:generate:reference:type=SubnetGroup
	// +kubebuilder:validation:Optional
	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name,omitempty"`

	// Reference to a SubnetGroup to populate dbSubnetGroupName.
	// +kubebuilder:validation:Optional
	DBSubnetGroupNameRef *v1.Reference `json:"dbSubnetGroupNameRef,omitempty" tf:"-"`

	// Selector for a SubnetGroup to populate dbSubnetGroupName.
	// +kubebuilder:validation:Optional
	DBSubnetGroupNameSelector *v1.Selector `json:"dbSubnetGroupNameSelector,omitempty" tf:"-"`

	// The name of the database engine to be used for the RDS instance. Defaults to aurora. Valid Values: aurora, aurora-mysql, aurora-postgresql.
	// For information on the difference between the available Aurora MySQL engines
	// see Comparison between Aurora MySQL 1 and Aurora MySQL 2
	// in the Amazon RDS User Guide.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The database engine version.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The instance class to use. For details on CPU
	// and memory, see Scaling Aurora DB Instances. Aurora uses db.* instance classes/types. Please see AWS Documentation for currently available instance classes and complete details.
	// +kubebuilder:validation:Required
	InstanceClass *string `json:"instanceClass" tf:"instance_class,omitempty"`

	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
	// +kubebuilder:validation:Optional
	MonitoringInterval *float64 `json:"monitoringInterval,omitempty" tf:"monitoring_interval,omitempty"`

	// The ARN for the IAM role that permits RDS to send
	// enhanced monitoring metrics to CloudWatch Logs. You can find more information on the AWS Documentation
	// what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	MonitoringRoleArn *string `json:"monitoringRoleArn,omitempty" tf:"monitoring_role_arn,omitempty"`

	// Reference to a Role in iam to populate monitoringRoleArn.
	// +kubebuilder:validation:Optional
	MonitoringRoleArnRef *v1.Reference `json:"monitoringRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate monitoringRoleArn.
	// +kubebuilder:validation:Optional
	MonitoringRoleArnSelector *v1.Selector `json:"monitoringRoleArnSelector,omitempty" tf:"-"`

	// Specifies whether Performance Insights is enabled or not.
	// +kubebuilder:validation:Optional
	PerformanceInsightsEnabled *bool `json:"performanceInsightsEnabled,omitempty" tf:"performance_insights_enabled,omitempty"`

	// ARN for the KMS key to encrypt Performance Insights data. When specifying performance_insights_kms_key_id, performance_insights_enabled needs to be set to true.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	PerformanceInsightsKMSKeyID *string `json:"performanceInsightsKmsKeyId,omitempty" tf:"performance_insights_kms_key_id,omitempty"`

	// Reference to a Key in kms to populate performanceInsightsKmsKeyId.
	// +kubebuilder:validation:Optional
	PerformanceInsightsKMSKeyIDRef *v1.Reference `json:"performanceInsightsKmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate performanceInsightsKmsKeyId.
	// +kubebuilder:validation:Optional
	PerformanceInsightsKMSKeyIDSelector *v1.Selector `json:"performanceInsightsKmsKeyIdSelector,omitempty" tf:"-"`

	// Amount of time in days to retain Performance Insights data. Either 7 (7 days) or 731 (2 years). When specifying performance_insights_retention_period, performance_insights_enabled needs to be set to true. Defaults to '7'.
	// +kubebuilder:validation:Optional
	PerformanceInsightsRetentionPeriod *float64 `json:"performanceInsightsRetentionPeriod,omitempty" tf:"performance_insights_retention_period,omitempty"`

	// The daily time range during which automated backups are created if automated backups are enabled.
	// Eg: "04:00-09:00"
	// +kubebuilder:validation:Optional
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`

	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	// +kubebuilder:validation:Optional
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
	// +kubebuilder:validation:Optional
	PromotionTier *float64 `json:"promotionTier,omitempty" tf:"promotion_tier,omitempty"`

	// Bool to control if instance is publicly accessible.
	// Default false. See the documentation on Creating DB Instances for more
	// details on controlling this property.
	// +kubebuilder:validation:Optional
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A map of tags to assign to the instance. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ClusterInstanceSpec defines the desired state of ClusterInstance
type ClusterInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterInstanceParameters `json:"forProvider"`
}

// ClusterInstanceStatus defines the observed state of ClusterInstance.
type ClusterInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterInstance is the Schema for the ClusterInstances API. Provides an RDS Cluster Resource Instance
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ClusterInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterInstanceSpec   `json:"spec"`
	Status            ClusterInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterInstanceList contains a list of ClusterInstances
type ClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterInstance `json:"items"`
}

// Repository type metadata.
var (
	ClusterInstance_Kind             = "ClusterInstance"
	ClusterInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterInstance_Kind}.String()
	ClusterInstance_KindAPIVersion   = ClusterInstance_Kind + "." + CRDGroupVersion.String()
	ClusterInstance_GroupVersionKind = CRDGroupVersion.WithKind(ClusterInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterInstance{}, &ClusterInstanceList{})
}
