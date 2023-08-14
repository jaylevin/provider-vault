/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SecretRoleInitParameters struct {

	// Whitelist entry in CIDR notation to be added for the API key.
	// Whitelist entry in CIDR notation to be added for the API key
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`

	// IP address to be added to the whitelist for the API key.
	// IP address to be added to the whitelist for the API key
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// The maximum allowed lifetime of credentials issued using this role.
	// The maximum allowed lifetime of credentials issued using this role
	MaxTTL *string `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// Path where the MongoDB Atlas Secrets Engine is mounted.
	// Path where MongoDB Atlas secret backend is mounted
	Mount *string `json:"mount,omitempty" tf:"mount,omitempty"`

	// The name of the role.
	// Name of the role
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Unique identifier for the organization to which the target API Key belongs.
	// Required if project_id is not set.
	// ID for the organization to which the target API Key belongs
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Unique identifier for the project to which the target API Key belongs.
	// Required if organization_id is not set.
	// ID for the project to which the target API Key belongs
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Roles assigned when an org API key is assigned to a project API key.
	// Roles assigned when an org API key is assigned to a project API key
	ProjectRoles []*string `json:"projectRoles,omitempty" tf:"project_roles,omitempty"`

	// List of roles that the API Key needs to have.
	// List of roles that the API Key needs to have
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// Duration in seconds after which the issued credential should expire.
	// Duration in seconds after which the issued credential should expire
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type SecretRoleObservation struct {

	// Whitelist entry in CIDR notation to be added for the API key.
	// Whitelist entry in CIDR notation to be added for the API key
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP address to be added to the whitelist for the API key.
	// IP address to be added to the whitelist for the API key
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// The maximum allowed lifetime of credentials issued using this role.
	// The maximum allowed lifetime of credentials issued using this role
	MaxTTL *string `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// Path where the MongoDB Atlas Secrets Engine is mounted.
	// Path where MongoDB Atlas secret backend is mounted
	Mount *string `json:"mount,omitempty" tf:"mount,omitempty"`

	// The name of the role.
	// Name of the role
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Unique identifier for the organization to which the target API Key belongs.
	// Required if project_id is not set.
	// ID for the organization to which the target API Key belongs
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Unique identifier for the project to which the target API Key belongs.
	// Required if organization_id is not set.
	// ID for the project to which the target API Key belongs
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Roles assigned when an org API key is assigned to a project API key.
	// Roles assigned when an org API key is assigned to a project API key
	ProjectRoles []*string `json:"projectRoles,omitempty" tf:"project_roles,omitempty"`

	// List of roles that the API Key needs to have.
	// List of roles that the API Key needs to have
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// Duration in seconds after which the issued credential should expire.
	// Duration in seconds after which the issued credential should expire
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type SecretRoleParameters struct {

	// Whitelist entry in CIDR notation to be added for the API key.
	// Whitelist entry in CIDR notation to be added for the API key
	// +kubebuilder:validation:Optional
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`

	// IP address to be added to the whitelist for the API key.
	// IP address to be added to the whitelist for the API key
	// +kubebuilder:validation:Optional
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// The maximum allowed lifetime of credentials issued using this role.
	// The maximum allowed lifetime of credentials issued using this role
	// +kubebuilder:validation:Optional
	MaxTTL *string `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// Path where the MongoDB Atlas Secrets Engine is mounted.
	// Path where MongoDB Atlas secret backend is mounted
	// +kubebuilder:validation:Optional
	Mount *string `json:"mount,omitempty" tf:"mount,omitempty"`

	// The name of the role.
	// Name of the role
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Unique identifier for the organization to which the target API Key belongs.
	// Required if project_id is not set.
	// ID for the organization to which the target API Key belongs
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Unique identifier for the project to which the target API Key belongs.
	// Required if organization_id is not set.
	// ID for the project to which the target API Key belongs
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Roles assigned when an org API key is assigned to a project API key.
	// Roles assigned when an org API key is assigned to a project API key
	// +kubebuilder:validation:Optional
	ProjectRoles []*string `json:"projectRoles,omitempty" tf:"project_roles,omitempty"`

	// List of roles that the API Key needs to have.
	// List of roles that the API Key needs to have
	// +kubebuilder:validation:Optional
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// Duration in seconds after which the issued credential should expire.
	// Duration in seconds after which the issued credential should expire
	// +kubebuilder:validation:Optional
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

// SecretRoleSpec defines the desired state of SecretRole
type SecretRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretRoleParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SecretRoleInitParameters `json:"initProvider,omitempty"`
}

// SecretRoleStatus defines the observed state of SecretRole.
type SecretRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretRole is the Schema for the SecretRoles API. Creates a role for the MongoDB Atlas Secret Engine in Vault.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mount) || has(self.initProvider.mount)",message="mount is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roles) || has(self.initProvider.roles)",message="roles is a required parameter"
	Spec   SecretRoleSpec   `json:"spec"`
	Status SecretRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretRoleList contains a list of SecretRoles
type SecretRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretRole `json:"items"`
}

// Repository type metadata.
var (
	SecretRole_Kind             = "SecretRole"
	SecretRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretRole_Kind}.String()
	SecretRole_KindAPIVersion   = SecretRole_Kind + "." + CRDGroupVersion.String()
	SecretRole_GroupVersionKind = CRDGroupVersion.WithKind(SecretRole_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretRole{}, &SecretRoleList{})
}
