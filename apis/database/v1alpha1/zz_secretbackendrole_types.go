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

type SecretBackendRoleInitParameters struct {

	// The unique name of the Vault mount to configure.
	// The path of the Database Secret Backend the role belongs to.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// The database statements to execute when
	// creating a user.
	// Database statements to execute to create and configure a user.
	CreationStatements []*string `json:"creationStatements,omitempty" tf:"creation_statements,omitempty"`

	// The unique name of the database connection to use for
	// the role.
	// Database connection to use for this role.
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// The default number of seconds for leases for this
	// role.
	// Default TTL for leases associated with this role, in seconds.
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// The maximum number of seconds for leases for this
	// role.
	// Maximum TTL for leases associated with this role, in seconds.
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// A unique name to give the role.
	// Unique name for the role.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The database statements to execute when
	// renewing a user.
	// Database statements to execute to renew a user.
	RenewStatements []*string `json:"renewStatements,omitempty" tf:"renew_statements,omitempty"`

	// The database statements to execute when
	// revoking a user.
	// Database statements to execute to revoke a user.
	RevocationStatements []*string `json:"revocationStatements,omitempty" tf:"revocation_statements,omitempty"`

	// The database statements to execute when
	// rolling back creation due to an error.
	// Database statements to execute to rollback a create operation in the event of an error.
	RollbackStatements []*string `json:"rollbackStatements,omitempty" tf:"rollback_statements,omitempty"`
}

type SecretBackendRoleObservation struct {

	// The unique name of the Vault mount to configure.
	// The path of the Database Secret Backend the role belongs to.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// The database statements to execute when
	// creating a user.
	// Database statements to execute to create and configure a user.
	CreationStatements []*string `json:"creationStatements,omitempty" tf:"creation_statements,omitempty"`

	// The unique name of the database connection to use for
	// the role.
	// Database connection to use for this role.
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// The default number of seconds for leases for this
	// role.
	// Default TTL for leases associated with this role, in seconds.
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The maximum number of seconds for leases for this
	// role.
	// Maximum TTL for leases associated with this role, in seconds.
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// A unique name to give the role.
	// Unique name for the role.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The database statements to execute when
	// renewing a user.
	// Database statements to execute to renew a user.
	RenewStatements []*string `json:"renewStatements,omitempty" tf:"renew_statements,omitempty"`

	// The database statements to execute when
	// revoking a user.
	// Database statements to execute to revoke a user.
	RevocationStatements []*string `json:"revocationStatements,omitempty" tf:"revocation_statements,omitempty"`

	// The database statements to execute when
	// rolling back creation due to an error.
	// Database statements to execute to rollback a create operation in the event of an error.
	RollbackStatements []*string `json:"rollbackStatements,omitempty" tf:"rollback_statements,omitempty"`
}

type SecretBackendRoleParameters struct {

	// The unique name of the Vault mount to configure.
	// The path of the Database Secret Backend the role belongs to.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// The database statements to execute when
	// creating a user.
	// Database statements to execute to create and configure a user.
	// +kubebuilder:validation:Optional
	CreationStatements []*string `json:"creationStatements,omitempty" tf:"creation_statements,omitempty"`

	// The unique name of the database connection to use for
	// the role.
	// Database connection to use for this role.
	// +kubebuilder:validation:Optional
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// The default number of seconds for leases for this
	// role.
	// Default TTL for leases associated with this role, in seconds.
	// +kubebuilder:validation:Optional
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// The maximum number of seconds for leases for this
	// role.
	// Maximum TTL for leases associated with this role, in seconds.
	// +kubebuilder:validation:Optional
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// A unique name to give the role.
	// Unique name for the role.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The database statements to execute when
	// renewing a user.
	// Database statements to execute to renew a user.
	// +kubebuilder:validation:Optional
	RenewStatements []*string `json:"renewStatements,omitempty" tf:"renew_statements,omitempty"`

	// The database statements to execute when
	// revoking a user.
	// Database statements to execute to revoke a user.
	// +kubebuilder:validation:Optional
	RevocationStatements []*string `json:"revocationStatements,omitempty" tf:"revocation_statements,omitempty"`

	// The database statements to execute when
	// rolling back creation due to an error.
	// Database statements to execute to rollback a create operation in the event of an error.
	// +kubebuilder:validation:Optional
	RollbackStatements []*string `json:"rollbackStatements,omitempty" tf:"rollback_statements,omitempty"`
}

// SecretBackendRoleSpec defines the desired state of SecretBackendRole
type SecretBackendRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendRoleParameters `json:"forProvider"`
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
	InitProvider SecretBackendRoleInitParameters `json:"initProvider,omitempty"`
}

// SecretBackendRoleStatus defines the observed state of SecretBackendRole.
type SecretBackendRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendRole is the Schema for the SecretBackendRoles API. Configures a database secret backend role for Vault.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackendRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.backend) || has(self.initProvider.backend)",message="backend is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.creationStatements) || has(self.initProvider.creationStatements)",message="creationStatements is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dbName) || has(self.initProvider.dbName)",message="dbName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   SecretBackendRoleSpec   `json:"spec"`
	Status SecretBackendRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendRoleList contains a list of SecretBackendRoles
type SecretBackendRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackendRole `json:"items"`
}

// Repository type metadata.
var (
	SecretBackendRole_Kind             = "SecretBackendRole"
	SecretBackendRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackendRole_Kind}.String()
	SecretBackendRole_KindAPIVersion   = SecretBackendRole_Kind + "." + CRDGroupVersion.String()
	SecretBackendRole_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackendRole_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackendRole{}, &SecretBackendRoleList{})
}
