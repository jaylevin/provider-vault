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

type SecretBackendInitParameters struct {

	// The default TTL for credentials
	// issued by this backend. Defaults to '0'.
	// Default lease duration for secrets in seconds
	DefaultLeaseTTLSeconds *float64 `json:"defaultLeaseTtlSeconds,omitempty" tf:"default_lease_ttl_seconds,omitempty"`

	// A human-friendly description for this backend.
	// Human-friendly description of the mount for the backend.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	// See here for more info on Mount Migration
	// If set, opts out of mount migration on path updates.
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	// Boolean flag that can be explicitly set to true to enforce local mount in HA environment
	// Local mount flag that can be explicitly set to true to enforce local mount in HA environment
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// The maximum TTL that can be requested
	// for credentials issued by this backend. Defaults to '0'.
	// Maximum possible lease duration for secrets in seconds
	MaxLeaseTTLSeconds *float64 `json:"maxLeaseTtlSeconds,omitempty" tf:"max_lease_ttl_seconds,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The unique path this backend should be mounted at. Must
	// not begin or end with a /. Defaults to gcp.
	// Path to mount the backend at.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type SecretBackendObservation struct {

	// The default TTL for credentials
	// issued by this backend. Defaults to '0'.
	// Default lease duration for secrets in seconds
	DefaultLeaseTTLSeconds *float64 `json:"defaultLeaseTtlSeconds,omitempty" tf:"default_lease_ttl_seconds,omitempty"`

	// A human-friendly description for this backend.
	// Human-friendly description of the mount for the backend.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	// See here for more info on Mount Migration
	// If set, opts out of mount migration on path updates.
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Boolean flag that can be explicitly set to true to enforce local mount in HA environment
	// Local mount flag that can be explicitly set to true to enforce local mount in HA environment
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// The maximum TTL that can be requested
	// for credentials issued by this backend. Defaults to '0'.
	// Maximum possible lease duration for secrets in seconds
	MaxLeaseTTLSeconds *float64 `json:"maxLeaseTtlSeconds,omitempty" tf:"max_lease_ttl_seconds,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The unique path this backend should be mounted at. Must
	// not begin or end with a /. Defaults to gcp.
	// Path to mount the backend at.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type SecretBackendParameters struct {

	// The GCP service account credentials in JSON format.
	// JSON-encoded credentials to use to connect to GCP
	// +kubebuilder:validation:Optional
	CredentialsSecretRef *v1.SecretKeySelector `json:"credentialsSecretRef,omitempty" tf:"-"`

	// The default TTL for credentials
	// issued by this backend. Defaults to '0'.
	// Default lease duration for secrets in seconds
	// +kubebuilder:validation:Optional
	DefaultLeaseTTLSeconds *float64 `json:"defaultLeaseTtlSeconds,omitempty" tf:"default_lease_ttl_seconds,omitempty"`

	// A human-friendly description for this backend.
	// Human-friendly description of the mount for the backend.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	// See here for more info on Mount Migration
	// If set, opts out of mount migration on path updates.
	// +kubebuilder:validation:Optional
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	// Boolean flag that can be explicitly set to true to enforce local mount in HA environment
	// Local mount flag that can be explicitly set to true to enforce local mount in HA environment
	// +kubebuilder:validation:Optional
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// The maximum TTL that can be requested
	// for credentials issued by this backend. Defaults to '0'.
	// Maximum possible lease duration for secrets in seconds
	// +kubebuilder:validation:Optional
	MaxLeaseTTLSeconds *float64 `json:"maxLeaseTtlSeconds,omitempty" tf:"max_lease_ttl_seconds,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The unique path this backend should be mounted at. Must
	// not begin or end with a /. Defaults to gcp.
	// Path to mount the backend at.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

// SecretBackendSpec defines the desired state of SecretBackend
type SecretBackendSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendParameters `json:"forProvider"`
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
	InitProvider SecretBackendInitParameters `json:"initProvider,omitempty"`
}

// SecretBackendStatus defines the observed state of SecretBackend.
type SecretBackendStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackend is the Schema for the SecretBackends API. Creates an GCP secret backend for Vault.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretBackendSpec   `json:"spec"`
	Status            SecretBackendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendList contains a list of SecretBackends
type SecretBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackend `json:"items"`
}

// Repository type metadata.
var (
	SecretBackend_Kind             = "SecretBackend"
	SecretBackend_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackend_Kind}.String()
	SecretBackend_KindAPIVersion   = SecretBackend_Kind + "." + CRDGroupVersion.String()
	SecretBackend_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackend_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackend{}, &SecretBackendList{})
}
