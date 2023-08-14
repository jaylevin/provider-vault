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

type AuthBackendGroupInitParameters struct {
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	Groupname *string `json:"groupname,omitempty" tf:"groupname,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type AuthBackendGroupObservation struct {
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	Groupname *string `json:"groupname,omitempty" tf:"groupname,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type AuthBackendGroupParameters struct {

	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// +kubebuilder:validation:Optional
	Groupname *string `json:"groupname,omitempty" tf:"groupname,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

// AuthBackendGroupSpec defines the desired state of AuthBackendGroup
type AuthBackendGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendGroupParameters `json:"forProvider"`
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
	InitProvider AuthBackendGroupInitParameters `json:"initProvider,omitempty"`
}

// AuthBackendGroupStatus defines the observed state of AuthBackendGroup.
type AuthBackendGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendGroup is the Schema for the AuthBackendGroups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackendGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupname) || has(self.initProvider.groupname)",message="groupname is a required parameter"
	Spec   AuthBackendGroupSpec   `json:"spec"`
	Status AuthBackendGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendGroupList contains a list of AuthBackendGroups
type AuthBackendGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackendGroup `json:"items"`
}

// Repository type metadata.
var (
	AuthBackendGroup_Kind             = "AuthBackendGroup"
	AuthBackendGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackendGroup_Kind}.String()
	AuthBackendGroup_KindAPIVersion   = AuthBackendGroup_Kind + "." + CRDGroupVersion.String()
	AuthBackendGroup_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackendGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackendGroup{}, &AuthBackendGroupList{})
}
