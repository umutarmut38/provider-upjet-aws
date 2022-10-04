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

type CapacityProviderStrategyObservation struct {
}

type CapacityProviderStrategyParameters struct {

	// Number of tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined.
	// +kubebuilder:validation:Optional
	Base *float64 `json:"base,omitempty" tf:"base,omitempty"`

	// Short name of the capacity provider.
	// +kubebuilder:validation:Required
	CapacityProvider *string `json:"capacityProvider" tf:"capacity_provider,omitempty"`

	// Relative percentage of the total number of launched tasks that should use the specified capacity provider.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type DeploymentCircuitBreakerObservation struct {
}

type DeploymentCircuitBreakerParameters struct {

	// Whether to enable the deployment circuit breaker logic for the service.
	// +kubebuilder:validation:Required
	Enable *bool `json:"enable" tf:"enable,omitempty"`

	// Whether to enable Amazon ECS to roll back the service if a service deployment fails. If rollback is enabled, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.
	// +kubebuilder:validation:Required
	Rollback *bool `json:"rollback" tf:"rollback,omitempty"`
}

type DeploymentControllerObservation struct {
}

type DeploymentControllerParameters struct {

	// Type of deployment controller. Valid values: CODE_DEPLOY, ECS, EXTERNAL. Default: ECS.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LoadBalancerObservation struct {
}

type LoadBalancerParameters struct {

	// Name of the container to associate with the load balancer (as it appears in a container definition).
	// +kubebuilder:validation:Required
	ContainerName *string `json:"containerName" tf:"container_name,omitempty"`

	// Port on the container to associate with the load balancer.
	// +kubebuilder:validation:Required
	ContainerPort *float64 `json:"containerPort" tf:"container_port,omitempty"`

	// Name of the ELB (Classic) to associate with the service.
	// +kubebuilder:validation:Optional
	ELBName *string `json:"elbName,omitempty" tf:"elb_name,omitempty"`

	// ARN of the Load Balancer target group to associate with the service.
	// +kubebuilder:validation:Optional
	TargetGroupArn *string `json:"targetGroupArn,omitempty" tf:"target_group_arn,omitempty"`
}

type NetworkConfigurationObservation struct {
}

type NetworkConfigurationParameters struct {

	// Assign a public IP address to the ENI (Fargate launch type only). Valid values are true or false. Default false.
	// +kubebuilder:validation:Optional
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupRefs []v1.Reference `json:"securityGroupRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupSelector *v1.Selector `json:"securityGroupSelector,omitempty" tf:"-"`

	// Security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupSelector
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// References to Subnet in ec2 to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetRefs []v1.Reference `json:"subnetRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetSelector *v1.Selector `json:"subnetSelector,omitempty" tf:"-"`

	// Subnets associated with the task or service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetSelector
	// +kubebuilder:validation:Optional
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`
}

type OrderedPlacementStrategyObservation struct {
}

type OrderedPlacementStrategyParameters struct {

	// For the spread placement strategy, valid values are instanceId (or host,
	// which has the same effect), or any platform or custom attribute that is applied to a container instance.
	// For the binpack type, valid values are memory and cpu. For the random type, this attribute is not
	// needed. For more information, see Placement Strategy.
	// +kubebuilder:validation:Optional
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// Type of placement strategy. Must be one of: binpack, random, or spread
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type PlacementConstraintsObservation struct {
}

type PlacementConstraintsParameters struct {

	// Cluster Query Language expression to apply to the constraint. Does not need to be specified for the distinctInstance type. For more information, see Cluster Query Language in the Amazon EC2 Container Service Developer Guide.
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Type of constraint. The only valid values at this time are memberOf and distinctInstance.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ServiceObservation struct {

	// ARN that identifies the service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ServiceParameters struct {

	// Capacity provider strategies to use for the service. Can be one or more. These can be updated without destroying and recreating the service only if force_new_deployment = true and not changing from 0 capacity_provider_strategy blocks to greater than 0, or vice versa. See below.
	// +kubebuilder:validation:Optional
	CapacityProviderStrategy []CapacityProviderStrategyParameters `json:"capacityProviderStrategy,omitempty" tf:"capacity_provider_strategy,omitempty"`

	// ARN of an ECS cluster.
	// +crossplane:generate:reference:type=Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// Reference to a Cluster to populate cluster.
	// +kubebuilder:validation:Optional
	ClusterRef *v1.Reference `json:"clusterRef,omitempty" tf:"-"`

	// Selector for a Cluster to populate cluster.
	// +kubebuilder:validation:Optional
	ClusterSelector *v1.Selector `json:"clusterSelector,omitempty" tf:"-"`

	// Configuration block for deployment circuit breaker. See below.
	// +kubebuilder:validation:Optional
	DeploymentCircuitBreaker []DeploymentCircuitBreakerParameters `json:"deploymentCircuitBreaker,omitempty" tf:"deployment_circuit_breaker,omitempty"`

	// Configuration block for deployment controller configuration. See below.
	// +kubebuilder:validation:Optional
	DeploymentController []DeploymentControllerParameters `json:"deploymentController,omitempty" tf:"deployment_controller,omitempty"`

	// Upper limit (as a percentage of the service's desiredCount) of the number of running tasks that can be running in a service during a deployment. Not valid when using the DAEMON scheduling strategy.
	// +kubebuilder:validation:Optional
	DeploymentMaximumPercent *float64 `json:"deploymentMaximumPercent,omitempty" tf:"deployment_maximum_percent,omitempty"`

	// Lower limit (as a percentage of the service's desiredCount) of the number of running tasks that must remain running and healthy in a service during a deployment.
	// +kubebuilder:validation:Optional
	DeploymentMinimumHealthyPercent *float64 `json:"deploymentMinimumHealthyPercent,omitempty" tf:"deployment_minimum_healthy_percent,omitempty"`

	// Number of instances of the task definition to place and keep running. Defaults to 0. Do not specify if using the DAEMON scheduling strategy.
	// +kubebuilder:validation:Optional
	DesiredCount *float64 `json:"desiredCount,omitempty" tf:"desired_count,omitempty"`

	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	// +kubebuilder:validation:Optional
	EnableEcsManagedTags *bool `json:"enableEcsManagedTags,omitempty" tf:"enable_ecs_managed_tags,omitempty"`

	// Specifies whether to enable Amazon ECS Exec for the tasks within the service.
	// +kubebuilder:validation:Optional
	EnableExecuteCommand *bool `json:"enableExecuteCommand,omitempty" tf:"enable_execute_command,omitempty"`

	// Enable to force a new task deployment of the service. This can be used to update tasks to use a newer Docker image with same image/tag combination (e.g., myimage:latest), roll Fargate tasks onto a newer platform version, or immediately deploy ordered_placement_strategy and placement_constraints updates.
	// +kubebuilder:validation:Optional
	ForceNewDeployment *bool `json:"forceNewDeployment,omitempty" tf:"force_new_deployment,omitempty"`

	// Seconds to ignore failing load balancer health checks on newly instantiated tasks to prevent premature shutdown, up to 2147483647. Only valid for services configured to use load balancers.
	// +kubebuilder:validation:Optional
	HealthCheckGracePeriodSeconds *float64 `json:"healthCheckGracePeriodSeconds,omitempty" tf:"health_check_grace_period_seconds,omitempty"`

	// ARN of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf. This parameter is required if you are using a load balancer with your service, but only if your task definition does not use the awsvpc network mode. If using awsvpc network mode, do not specify this role. If your account has already created the Amazon ECS service-linked role, that role is used by default for your service unless you specify a role here.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	IAMRole *string `json:"iamRole,omitempty" tf:"iam_role,omitempty"`

	// Reference to a Role in iam to populate iamRole.
	// +kubebuilder:validation:Optional
	IAMRoleRef *v1.Reference `json:"iamRoleRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate iamRole.
	// +kubebuilder:validation:Optional
	IAMRoleSelector *v1.Selector `json:"iamRoleSelector,omitempty" tf:"-"`

	// Launch type on which to run your service. The valid values are EC2, FARGATE, and EXTERNAL. Defaults to EC2.
	// +kubebuilder:validation:Optional
	LaunchType *string `json:"launchType,omitempty" tf:"launch_type,omitempty"`

	// Configuration block for load balancers. See below.
	// +kubebuilder:validation:Optional
	LoadBalancer []LoadBalancerParameters `json:"loadBalancer,omitempty" tf:"load_balancer,omitempty"`

	// Network configuration for the service. This parameter is required for task definitions that use the awsvpc network mode to receive their own Elastic Network Interface, and it is not supported for other network modes. See below.
	// +kubebuilder:validation:Optional
	NetworkConfiguration []NetworkConfigurationParameters `json:"networkConfiguration,omitempty" tf:"network_configuration,omitempty"`

	// Service level strategy rules that are taken into consideration during task placement. List from top to bottom in order of precedence. Updates to this configuration will take effect next task deployment unless force_new_deployment is enabled. The maximum number of ordered_placement_strategy blocks is 5. See below.
	// +kubebuilder:validation:Optional
	OrderedPlacementStrategy []OrderedPlacementStrategyParameters `json:"orderedPlacementStrategy,omitempty" tf:"ordered_placement_strategy,omitempty"`

	// Rules that are taken into consideration during task placement. Updates to this configuration will take effect next task deployment unless force_new_deployment is enabled. Maximum number of placement_constraints is 10. See below.
	// +kubebuilder:validation:Optional
	PlacementConstraints []PlacementConstraintsParameters `json:"placementConstraints,omitempty" tf:"placement_constraints,omitempty"`

	// Platform version on which to run your service. Only applicable for launch_type set to FARGATE. Defaults to LATEST. More information about Fargate platform versions can be found in the AWS ECS User Guide.
	// +kubebuilder:validation:Optional
	PlatformVersion *string `json:"platformVersion,omitempty" tf:"platform_version,omitempty"`

	// Specifies whether to propagate the tags from the task definition or the service to the tasks. The valid values are SERVICE and TASK_DEFINITION.
	// +kubebuilder:validation:Optional
	PropagateTags *string `json:"propagateTags,omitempty" tf:"propagate_tags,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Scheduling strategy to use for the service. The valid values are REPLICA and DAEMON. Defaults to REPLICA. Note that Tasks using the Fargate launch type or the .
	// +kubebuilder:validation:Optional
	SchedulingStrategy *string `json:"schedulingStrategy,omitempty" tf:"scheduling_strategy,omitempty"`

	// Service discovery registries for the service. The maximum number of service_registries blocks is 1. See below.
	// +kubebuilder:validation:Optional
	ServiceRegistries []ServiceRegistriesParameters `json:"serviceRegistries,omitempty" tf:"service_registries,omitempty"`

	// Key-value map of resource tags. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Family and revision (family:revision) or full ARN of the task definition that you want to run in your service. Required unless using the EXTERNAL deployment controller. If a revision is not specified, the latest ACTIVE revision is used.
	// +kubebuilder:validation:Optional
	TaskDefinition *string `json:"taskDefinition,omitempty" tf:"task_definition,omitempty"`

	// Default false.
	// +kubebuilder:validation:Optional
	WaitForSteadyState *bool `json:"waitForSteadyState,omitempty" tf:"wait_for_steady_state,omitempty"`
}

type ServiceRegistriesObservation struct {
}

type ServiceRegistriesParameters struct {

	// Container name value, already specified in the task definition, to be used for your service discovery service.
	// +kubebuilder:validation:Optional
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// Port value, already specified in the task definition, to be used for your service discovery service.
	// +kubebuilder:validation:Optional
	ContainerPort *float64 `json:"containerPort,omitempty" tf:"container_port,omitempty"`

	// Port value used if your Service Discovery service specified an SRV record.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// ARN of the Service Registry. The currently supported service registry is Amazon Route 53 Auto Naming Service(aws_service_discovery_service). For more information, see Service
	// +kubebuilder:validation:Required
	RegistryArn *string `json:"registryArn" tf:"registry_arn,omitempty"`
}

// ServiceSpec defines the desired state of Service
type ServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceParameters `json:"forProvider"`
}

// ServiceStatus defines the observed state of Service.
type ServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Service is the Schema for the Services API. Provides an ECS service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec   `json:"spec"`
	Status            ServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceList contains a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

// Repository type metadata.
var (
	Service_Kind             = "Service"
	Service_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Service_Kind}.String()
	Service_KindAPIVersion   = Service_Kind + "." + CRDGroupVersion.String()
	Service_GroupVersionKind = CRDGroupVersion.WithKind(Service_Kind)
)

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}
