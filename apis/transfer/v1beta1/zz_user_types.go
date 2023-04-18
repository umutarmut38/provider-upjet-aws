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

type HomeDirectoryMappingsObservation struct {

	// Represents an entry and a target.
	Entry *string `json:"entry,omitempty" tf:"entry,omitempty"`

	// Represents the map target.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`
}

type HomeDirectoryMappingsParameters struct {

	// Represents an entry and a target.
	// +kubebuilder:validation:Required
	Entry *string `json:"entry" tf:"entry,omitempty"`

	// Represents the map target.
	// +kubebuilder:validation:Required
	Target *string `json:"target" tf:"target,omitempty"`
}

type PosixProfileObservation struct {

	// The POSIX group ID used for all EFS operations by this user.
	GID *float64 `json:"gid,omitempty" tf:"gid,omitempty"`

	// The secondary POSIX group IDs used for all EFS operations by this user.
	SecondaryGids []*float64 `json:"secondaryGids,omitempty" tf:"secondary_gids,omitempty"`

	// The POSIX user ID used for all EFS operations by this user.
	UID *float64 `json:"uid,omitempty" tf:"uid,omitempty"`
}

type PosixProfileParameters struct {

	// The POSIX group ID used for all EFS operations by this user.
	// +kubebuilder:validation:Required
	GID *float64 `json:"gid" tf:"gid,omitempty"`

	// The secondary POSIX group IDs used for all EFS operations by this user.
	// +kubebuilder:validation:Optional
	SecondaryGids []*float64 `json:"secondaryGids,omitempty" tf:"secondary_gids,omitempty"`

	// The POSIX user ID used for all EFS operations by this user.
	// +kubebuilder:validation:Required
	UID *float64 `json:"uid" tf:"uid,omitempty"`
}

type UserObservation struct {

	// Amazon Resource Name (ARN) of Transfer User
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The landing directory (folder) for a user when they log in to the server using their SFTP client.  It should begin with a /.  The first item in the path is the name of the home bucket (accessible as ${Transfer:HomeBucket} in the policy) and the rest is the home directory (accessible as ${Transfer:HomeDirectory} in the policy). For example, /example-bucket-1234/username would set the home bucket to example-bucket-1234 and the home directory to username.
	HomeDirectory *string `json:"homeDirectory,omitempty" tf:"home_directory,omitempty"`

	// Logical directory mappings that specify what S3 paths and keys should be visible to your user and how you want to make them visible. See Home Directory Mappings below.
	HomeDirectoryMappings []HomeDirectoryMappingsObservation `json:"homeDirectoryMappings,omitempty" tf:"home_directory_mappings,omitempty"`

	// The type of landing directory (folder) you mapped for your users' home directory. Valid values are PATH and LOGICAL.
	HomeDirectoryType *string `json:"homeDirectoryType,omitempty" tf:"home_directory_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include ${Transfer:UserName}, ${Transfer:HomeDirectory}, and ${Transfer:HomeBucket}.  These are evaluated on-the-fly when navigating the bucket.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Specifies the full POSIX identity, including user ID (Uid), group ID (Gid), and any secondary groups IDs (SecondaryGids), that controls your users' access to your Amazon EFS file systems. See Posix Profile below.
	PosixProfile []PosixProfileObservation `json:"posixProfile,omitempty" tf:"posix_profile,omitempty"`

	// Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The Server ID of the Transfer Server (e.g., s-12345678)
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type UserParameters struct {

	// The landing directory (folder) for a user when they log in to the server using their SFTP client.  It should begin with a /.  The first item in the path is the name of the home bucket (accessible as ${Transfer:HomeBucket} in the policy) and the rest is the home directory (accessible as ${Transfer:HomeDirectory} in the policy). For example, /example-bucket-1234/username would set the home bucket to example-bucket-1234 and the home directory to username.
	// +kubebuilder:validation:Optional
	HomeDirectory *string `json:"homeDirectory,omitempty" tf:"home_directory,omitempty"`

	// Logical directory mappings that specify what S3 paths and keys should be visible to your user and how you want to make them visible. See Home Directory Mappings below.
	// +kubebuilder:validation:Optional
	HomeDirectoryMappings []HomeDirectoryMappingsParameters `json:"homeDirectoryMappings,omitempty" tf:"home_directory_mappings,omitempty"`

	// The type of landing directory (folder) you mapped for your users' home directory. Valid values are PATH and LOGICAL.
	// +kubebuilder:validation:Optional
	HomeDirectoryType *string `json:"homeDirectoryType,omitempty" tf:"home_directory_type,omitempty"`

	// An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include ${Transfer:UserName}, ${Transfer:HomeDirectory}, and ${Transfer:HomeBucket}.  These are evaluated on-the-fly when navigating the bucket.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Specifies the full POSIX identity, including user ID (Uid), group ID (Gid), and any secondary groups IDs (SecondaryGids), that controls your users' access to your Amazon EFS file systems. See Posix Profile below.
	// +kubebuilder:validation:Optional
	PosixProfile []PosixProfileParameters `json:"posixProfile,omitempty" tf:"posix_profile,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a Role in iam to populate role.
	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate role.
	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`

	// The Server ID of the Transfer Server (e.g., s-12345678)
	// +crossplane:generate:reference:type=Server
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a Server to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a Server to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// User is the Schema for the Users API. Provides a AWS Transfer User resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserSpec   `json:"spec"`
	Status            UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
