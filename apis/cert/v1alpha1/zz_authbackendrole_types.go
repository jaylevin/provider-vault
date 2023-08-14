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

type AuthBackendRoleInitParameters struct {
	AllowedCommonNames []*string `json:"allowedCommonNames,omitempty" tf:"allowed_common_names,omitempty"`

	AllowedDNSSans []*string `json:"allowedDnsSans,omitempty" tf:"allowed_dns_sans,omitempty"`

	AllowedEmailSans []*string `json:"allowedEmailSans,omitempty" tf:"allowed_email_sans,omitempty"`

	AllowedNames []*string `json:"allowedNames,omitempty" tf:"allowed_names,omitempty"`

	AllowedOrganizationUnits []*string `json:"allowedOrganizationUnits,omitempty" tf:"allowed_organization_units,omitempty"`

	AllowedOrganizationalUnits []*string `json:"allowedOrganizationalUnits,omitempty" tf:"allowed_organizational_units,omitempty"`

	AllowedURISans []*string `json:"allowedUriSans,omitempty" tf:"allowed_uri_sans,omitempty"`

	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	RequiredExtensions []*string `json:"requiredExtensions,omitempty" tf:"required_extensions,omitempty"`

	// Specifies the blocks of IP addresses which are allowed to use the generated token
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime of the generated token
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// Generated Token's Period
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// Generated Token's Policies
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The initial ttl of the token to generate in seconds
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token to generate, service or batch
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

type AuthBackendRoleObservation struct {
	AllowedCommonNames []*string `json:"allowedCommonNames,omitempty" tf:"allowed_common_names,omitempty"`

	AllowedDNSSans []*string `json:"allowedDnsSans,omitempty" tf:"allowed_dns_sans,omitempty"`

	AllowedEmailSans []*string `json:"allowedEmailSans,omitempty" tf:"allowed_email_sans,omitempty"`

	AllowedNames []*string `json:"allowedNames,omitempty" tf:"allowed_names,omitempty"`

	AllowedOrganizationUnits []*string `json:"allowedOrganizationUnits,omitempty" tf:"allowed_organization_units,omitempty"`

	AllowedOrganizationalUnits []*string `json:"allowedOrganizationalUnits,omitempty" tf:"allowed_organizational_units,omitempty"`

	AllowedURISans []*string `json:"allowedUriSans,omitempty" tf:"allowed_uri_sans,omitempty"`

	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	RequiredExtensions []*string `json:"requiredExtensions,omitempty" tf:"required_extensions,omitempty"`

	// Specifies the blocks of IP addresses which are allowed to use the generated token
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime of the generated token
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// Generated Token's Period
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// Generated Token's Policies
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The initial ttl of the token to generate in seconds
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token to generate, service or batch
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

type AuthBackendRoleParameters struct {

	// +kubebuilder:validation:Optional
	AllowedCommonNames []*string `json:"allowedCommonNames,omitempty" tf:"allowed_common_names,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedDNSSans []*string `json:"allowedDnsSans,omitempty" tf:"allowed_dns_sans,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedEmailSans []*string `json:"allowedEmailSans,omitempty" tf:"allowed_email_sans,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedNames []*string `json:"allowedNames,omitempty" tf:"allowed_names,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedOrganizationUnits []*string `json:"allowedOrganizationUnits,omitempty" tf:"allowed_organization_units,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedOrganizationalUnits []*string `json:"allowedOrganizationalUnits,omitempty" tf:"allowed_organizational_units,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedURISans []*string `json:"allowedUriSans,omitempty" tf:"allowed_uri_sans,omitempty"`

	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// +kubebuilder:validation:Optional
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	RequiredExtensions []*string `json:"requiredExtensions,omitempty" tf:"required_extensions,omitempty"`

	// Specifies the blocks of IP addresses which are allowed to use the generated token
	// +kubebuilder:validation:Optional
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// Generated Token's Explicit Maximum TTL in seconds
	// +kubebuilder:validation:Optional
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime of the generated token
	// +kubebuilder:validation:Optional
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	// +kubebuilder:validation:Optional
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	// +kubebuilder:validation:Optional
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// Generated Token's Period
	// +kubebuilder:validation:Optional
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// Generated Token's Policies
	// +kubebuilder:validation:Optional
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The initial ttl of the token to generate in seconds
	// +kubebuilder:validation:Optional
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token to generate, service or batch
	// +kubebuilder:validation:Optional
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

// AuthBackendRoleSpec defines the desired state of AuthBackendRole
type AuthBackendRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendRoleParameters `json:"forProvider"`
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
	InitProvider AuthBackendRoleInitParameters `json:"initProvider,omitempty"`
}

// AuthBackendRoleStatus defines the observed state of AuthBackendRole.
type AuthBackendRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendRole is the Schema for the AuthBackendRoles API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackendRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.certificate) || has(self.initProvider.certificate)",message="certificate is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   AuthBackendRoleSpec   `json:"spec"`
	Status AuthBackendRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendRoleList contains a list of AuthBackendRoles
type AuthBackendRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackendRole `json:"items"`
}

// Repository type metadata.
var (
	AuthBackendRole_Kind             = "AuthBackendRole"
	AuthBackendRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackendRole_Kind}.String()
	AuthBackendRole_KindAPIVersion   = AuthBackendRole_Kind + "." + CRDGroupVersion.String()
	AuthBackendRole_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackendRole_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackendRole{}, &AuthBackendRoleList{})
}
