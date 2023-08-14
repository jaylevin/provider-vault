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

type LeaseCountInitParameters struct {

	// The maximum number of leases to be allowed by the quota
	// rule. The max_leases must be positive.
	// The maximum number of leases to be allowed by the quota rule. The max_leases must be positive.
	MaxLeases *float64 `json:"maxLeases,omitempty" tf:"max_leases,omitempty"`

	// Name of the rate limit quota
	// The name of the quota.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example namespace1/ adds a quota to a full namespace,
	// namespace1/auth/userpass adds a quota to userpass in namespace1.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// auth/userpass to namespace1/auth/userpass moves this quota from being a global mount quota to
	// a namespace specific mount quota. Note, namespaces are supported in Enterprise only.
	// Path of the mount or namespace to apply the quota. A blank path configures a global lease count quota.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type LeaseCountObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The maximum number of leases to be allowed by the quota
	// rule. The max_leases must be positive.
	// The maximum number of leases to be allowed by the quota rule. The max_leases must be positive.
	MaxLeases *float64 `json:"maxLeases,omitempty" tf:"max_leases,omitempty"`

	// Name of the rate limit quota
	// The name of the quota.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example namespace1/ adds a quota to a full namespace,
	// namespace1/auth/userpass adds a quota to userpass in namespace1.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// auth/userpass to namespace1/auth/userpass moves this quota from being a global mount quota to
	// a namespace specific mount quota. Note, namespaces are supported in Enterprise only.
	// Path of the mount or namespace to apply the quota. A blank path configures a global lease count quota.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type LeaseCountParameters struct {

	// The maximum number of leases to be allowed by the quota
	// rule. The max_leases must be positive.
	// The maximum number of leases to be allowed by the quota rule. The max_leases must be positive.
	// +kubebuilder:validation:Optional
	MaxLeases *float64 `json:"maxLeases,omitempty" tf:"max_leases,omitempty"`

	// Name of the rate limit quota
	// The name of the quota.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example namespace1/ adds a quota to a full namespace,
	// namespace1/auth/userpass adds a quota to userpass in namespace1.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// auth/userpass to namespace1/auth/userpass moves this quota from being a global mount quota to
	// a namespace specific mount quota. Note, namespaces are supported in Enterprise only.
	// Path of the mount or namespace to apply the quota. A blank path configures a global lease count quota.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

// LeaseCountSpec defines the desired state of LeaseCount
type LeaseCountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LeaseCountParameters `json:"forProvider"`
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
	InitProvider LeaseCountInitParameters `json:"initProvider,omitempty"`
}

// LeaseCountStatus defines the observed state of LeaseCount.
type LeaseCountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LeaseCountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LeaseCount is the Schema for the LeaseCounts API. Manage Lease Count Quota
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type LeaseCount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.maxLeases) || has(self.initProvider.maxLeases)",message="maxLeases is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   LeaseCountSpec   `json:"spec"`
	Status LeaseCountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LeaseCountList contains a list of LeaseCounts
type LeaseCountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LeaseCount `json:"items"`
}

// Repository type metadata.
var (
	LeaseCount_Kind             = "LeaseCount"
	LeaseCount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LeaseCount_Kind}.String()
	LeaseCount_KindAPIVersion   = LeaseCount_Kind + "." + CRDGroupVersion.String()
	LeaseCount_GroupVersionKind = CRDGroupVersion.WithKind(LeaseCount_Kind)
)

func init() {
	SchemeBuilder.Register(&LeaseCount{}, &LeaseCountList{})
}
