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

	// List of managed key registry entry names that the mount in question is allowed to access
	AllowedManagedKeys []*string `json:"allowedManagedKeys,omitempty" tf:"allowed_managed_keys,omitempty"`

	// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
	AuditNonHMACRequestKeys []*string `json:"auditNonHmacRequestKeys,omitempty" tf:"audit_non_hmac_request_keys,omitempty"`

	// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
	AuditNonHMACResponseKeys []*string `json:"auditNonHmacResponseKeys,omitempty" tf:"audit_non_hmac_response_keys,omitempty"`

	// Default lease duration for tokens and secrets in seconds
	DefaultLeaseTTLSeconds *float64 `json:"defaultLeaseTtlSeconds,omitempty" tf:"default_lease_ttl_seconds,omitempty"`

	// Human-friendly description of the mount
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Disable defaulting to the local CA certificate and
	// service account JWT when Vault is running in a Kubernetes pod.
	// Disable defaulting to the local CA certificate and service account JWT when running in a Kubernetes pod.
	DisableLocalCAJwt *bool `json:"disableLocalCaJwt,omitempty" tf:"disable_local_ca_jwt,omitempty"`

	// Enable the secrets engine to access Vault's external entropy source
	ExternalEntropyAccess *bool `json:"externalEntropyAccess,omitempty" tf:"external_entropy_access,omitempty"`

	// A PEM-encoded CA certificate used by the
	// secrets engine to verify the Kubernetes API server certificate. Defaults to the local
	// pod’s CA if Vault is running in Kubernetes. Otherwise, defaults to the root CA set where
	// Vault is running.
	// A PEM-encoded CA certificate used by the secret engine to verify the Kubernetes API server certificate. Defaults to the local pod’s CA if found, or otherwise the host's root CA set.
	KubernetesCACert *string `json:"kubernetesCaCert,omitempty" tf:"kubernetes_ca_cert,omitempty"`

	// The Kubernetes API URL to connect to. Required if the
	// standard pod environment variables KUBERNETES_SERVICE_HOST or KUBERNETES_SERVICE_PORT
	// are not set on the host that Vault is running on.
	// The Kubernetes API URL to connect to.
	KubernetesHost *string `json:"kubernetesHost,omitempty" tf:"kubernetes_host,omitempty"`

	// Local mount flag that can be explicitly set to true to enforce local mount in HA environment
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// Maximum possible lease duration for tokens and secrets in seconds
	MaxLeaseTTLSeconds *float64 `json:"maxLeaseTtlSeconds,omitempty" tf:"max_lease_ttl_seconds,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Specifies mount type specific options that are passed to the backend
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`

	// Where the secret backend will be mounted
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
	SealWrap *bool `json:"sealWrap,omitempty" tf:"seal_wrap,omitempty"`
}

type SecretBackendObservation struct {

	// Accessor of the mount
	Accessor *string `json:"accessor,omitempty" tf:"accessor,omitempty"`

	// List of managed key registry entry names that the mount in question is allowed to access
	AllowedManagedKeys []*string `json:"allowedManagedKeys,omitempty" tf:"allowed_managed_keys,omitempty"`

	// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
	AuditNonHMACRequestKeys []*string `json:"auditNonHmacRequestKeys,omitempty" tf:"audit_non_hmac_request_keys,omitempty"`

	// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
	AuditNonHMACResponseKeys []*string `json:"auditNonHmacResponseKeys,omitempty" tf:"audit_non_hmac_response_keys,omitempty"`

	// Default lease duration for tokens and secrets in seconds
	DefaultLeaseTTLSeconds *float64 `json:"defaultLeaseTtlSeconds,omitempty" tf:"default_lease_ttl_seconds,omitempty"`

	// Human-friendly description of the mount
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Disable defaulting to the local CA certificate and
	// service account JWT when Vault is running in a Kubernetes pod.
	// Disable defaulting to the local CA certificate and service account JWT when running in a Kubernetes pod.
	DisableLocalCAJwt *bool `json:"disableLocalCaJwt,omitempty" tf:"disable_local_ca_jwt,omitempty"`

	// Enable the secrets engine to access Vault's external entropy source
	ExternalEntropyAccess *bool `json:"externalEntropyAccess,omitempty" tf:"external_entropy_access,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A PEM-encoded CA certificate used by the
	// secrets engine to verify the Kubernetes API server certificate. Defaults to the local
	// pod’s CA if Vault is running in Kubernetes. Otherwise, defaults to the root CA set where
	// Vault is running.
	// A PEM-encoded CA certificate used by the secret engine to verify the Kubernetes API server certificate. Defaults to the local pod’s CA if found, or otherwise the host's root CA set.
	KubernetesCACert *string `json:"kubernetesCaCert,omitempty" tf:"kubernetes_ca_cert,omitempty"`

	// The Kubernetes API URL to connect to. Required if the
	// standard pod environment variables KUBERNETES_SERVICE_HOST or KUBERNETES_SERVICE_PORT
	// are not set on the host that Vault is running on.
	// The Kubernetes API URL to connect to.
	KubernetesHost *string `json:"kubernetesHost,omitempty" tf:"kubernetes_host,omitempty"`

	// Local mount flag that can be explicitly set to true to enforce local mount in HA environment
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// Maximum possible lease duration for tokens and secrets in seconds
	MaxLeaseTTLSeconds *float64 `json:"maxLeaseTtlSeconds,omitempty" tf:"max_lease_ttl_seconds,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Specifies mount type specific options that are passed to the backend
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`

	// Where the secret backend will be mounted
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
	SealWrap *bool `json:"sealWrap,omitempty" tf:"seal_wrap,omitempty"`
}

type SecretBackendParameters struct {

	// List of managed key registry entry names that the mount in question is allowed to access
	// +kubebuilder:validation:Optional
	AllowedManagedKeys []*string `json:"allowedManagedKeys,omitempty" tf:"allowed_managed_keys,omitempty"`

	// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
	// +kubebuilder:validation:Optional
	AuditNonHMACRequestKeys []*string `json:"auditNonHmacRequestKeys,omitempty" tf:"audit_non_hmac_request_keys,omitempty"`

	// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
	// +kubebuilder:validation:Optional
	AuditNonHMACResponseKeys []*string `json:"auditNonHmacResponseKeys,omitempty" tf:"audit_non_hmac_response_keys,omitempty"`

	// Default lease duration for tokens and secrets in seconds
	// +kubebuilder:validation:Optional
	DefaultLeaseTTLSeconds *float64 `json:"defaultLeaseTtlSeconds,omitempty" tf:"default_lease_ttl_seconds,omitempty"`

	// Human-friendly description of the mount
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Disable defaulting to the local CA certificate and
	// service account JWT when Vault is running in a Kubernetes pod.
	// Disable defaulting to the local CA certificate and service account JWT when running in a Kubernetes pod.
	// +kubebuilder:validation:Optional
	DisableLocalCAJwt *bool `json:"disableLocalCaJwt,omitempty" tf:"disable_local_ca_jwt,omitempty"`

	// Enable the secrets engine to access Vault's external entropy source
	// +kubebuilder:validation:Optional
	ExternalEntropyAccess *bool `json:"externalEntropyAccess,omitempty" tf:"external_entropy_access,omitempty"`

	// A PEM-encoded CA certificate used by the
	// secrets engine to verify the Kubernetes API server certificate. Defaults to the local
	// pod’s CA if Vault is running in Kubernetes. Otherwise, defaults to the root CA set where
	// Vault is running.
	// A PEM-encoded CA certificate used by the secret engine to verify the Kubernetes API server certificate. Defaults to the local pod’s CA if found, or otherwise the host's root CA set.
	// +kubebuilder:validation:Optional
	KubernetesCACert *string `json:"kubernetesCaCert,omitempty" tf:"kubernetes_ca_cert,omitempty"`

	// The Kubernetes API URL to connect to. Required if the
	// standard pod environment variables KUBERNETES_SERVICE_HOST or KUBERNETES_SERVICE_PORT
	// are not set on the host that Vault is running on.
	// The Kubernetes API URL to connect to.
	// +kubebuilder:validation:Optional
	KubernetesHost *string `json:"kubernetesHost,omitempty" tf:"kubernetes_host,omitempty"`

	// Local mount flag that can be explicitly set to true to enforce local mount in HA environment
	// +kubebuilder:validation:Optional
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// Maximum possible lease duration for tokens and secrets in seconds
	// +kubebuilder:validation:Optional
	MaxLeaseTTLSeconds *float64 `json:"maxLeaseTtlSeconds,omitempty" tf:"max_lease_ttl_seconds,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Specifies mount type specific options that are passed to the backend
	// +kubebuilder:validation:Optional
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`

	// Where the secret backend will be mounted
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
	// +kubebuilder:validation:Optional
	SealWrap *bool `json:"sealWrap,omitempty" tf:"seal_wrap,omitempty"`

	// The JSON web token of the service account used by the
	// secrets engine to manage Kubernetes credentials. Defaults to the local pod’s JWT if Vault
	// is running in Kubernetes.
	// The JSON web token of the service account used by the secrets engine to manage Kubernetes credentials. Defaults to the local pod’s JWT if found.
	// +kubebuilder:validation:Optional
	ServiceAccountJwtSecretRef *v1.SecretKeySelector `json:"serviceAccountJwtSecretRef,omitempty" tf:"-"`
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

// SecretBackend is the Schema for the SecretBackends API. Creates a Kubernetes Secrets Engine in Vault.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.path) || has(self.initProvider.path)",message="path is a required parameter"
	Spec   SecretBackendSpec   `json:"spec"`
	Status SecretBackendStatus `json:"status,omitempty"`
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
