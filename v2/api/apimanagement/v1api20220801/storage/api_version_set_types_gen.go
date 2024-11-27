// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=apimanagement.azure.com,resources=apiversionsets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apimanagement.azure.com,resources={apiversionsets/status,apiversionsets/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220801.ApiVersionSet
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimapiversionsets.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}
type ApiVersionSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiVersionSet_Spec   `json:"spec,omitempty"`
	Status            ApiVersionSet_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ApiVersionSet{}

// GetConditions returns the conditions of the resource
func (versionSet *ApiVersionSet) GetConditions() conditions.Conditions {
	return versionSet.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (versionSet *ApiVersionSet) SetConditions(conditions conditions.Conditions) {
	versionSet.Status.Conditions = conditions
}

var _ configmaps.Exporter = &ApiVersionSet{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (versionSet *ApiVersionSet) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if versionSet.Spec.OperatorSpec == nil {
		return nil
	}
	return versionSet.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &ApiVersionSet{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (versionSet *ApiVersionSet) SecretDestinationExpressions() []*core.DestinationExpression {
	if versionSet.Spec.OperatorSpec == nil {
		return nil
	}
	return versionSet.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &ApiVersionSet{}

// AzureName returns the Azure name of the resource
func (versionSet *ApiVersionSet) AzureName() string {
	return versionSet.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (versionSet ApiVersionSet) GetAPIVersion() string {
	return "2022-08-01"
}

// GetResourceScope returns the scope of the resource
func (versionSet *ApiVersionSet) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (versionSet *ApiVersionSet) GetSpec() genruntime.ConvertibleSpec {
	return &versionSet.Spec
}

// GetStatus returns the status of this resource
func (versionSet *ApiVersionSet) GetStatus() genruntime.ConvertibleStatus {
	return &versionSet.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (versionSet *ApiVersionSet) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationHead,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/apiVersionSets"
func (versionSet *ApiVersionSet) GetType() string {
	return "Microsoft.ApiManagement/service/apiVersionSets"
}

// NewEmptyStatus returns a new empty (blank) status
func (versionSet *ApiVersionSet) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ApiVersionSet_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (versionSet *ApiVersionSet) Owner() *genruntime.ResourceReference {
	if versionSet.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(versionSet.Spec)
	return versionSet.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (versionSet *ApiVersionSet) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ApiVersionSet_STATUS); ok {
		versionSet.Status = *st
		return nil
	}

	// Convert status to required version
	var st ApiVersionSet_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	versionSet.Status = st
	return nil
}

// Hub marks that this ApiVersionSet is the hub type for conversion
func (versionSet *ApiVersionSet) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (versionSet *ApiVersionSet) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: versionSet.Spec.OriginalVersion,
		Kind:    "ApiVersionSet",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220801.ApiVersionSet
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimapiversionsets.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}
type ApiVersionSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiVersionSet `json:"items"`
}

// Storage version of v1api20220801.ApiVersionSet_Spec
type ApiVersionSet_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                     `json:"azureName,omitempty"`
	Description     *string                    `json:"description,omitempty"`
	DisplayName     *string                    `json:"displayName,omitempty"`
	OperatorSpec    *ApiVersionSetOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                     `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner             *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`
	PropertyBag       genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	VersionHeaderName *string                            `json:"versionHeaderName,omitempty"`
	VersionQueryName  *string                            `json:"versionQueryName,omitempty"`
	VersioningScheme  *string                            `json:"versioningScheme,omitempty"`
}

var _ genruntime.ConvertibleSpec = &ApiVersionSet_Spec{}

// ConvertSpecFrom populates our ApiVersionSet_Spec from the provided source
func (versionSet *ApiVersionSet_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == versionSet {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(versionSet)
}

// ConvertSpecTo populates the provided destination from our ApiVersionSet_Spec
func (versionSet *ApiVersionSet_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == versionSet {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(versionSet)
}

// Storage version of v1api20220801.ApiVersionSet_STATUS
type ApiVersionSet_STATUS struct {
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	Description       *string                `json:"description,omitempty"`
	DisplayName       *string                `json:"displayName,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type              *string                `json:"type,omitempty"`
	VersionHeaderName *string                `json:"versionHeaderName,omitempty"`
	VersionQueryName  *string                `json:"versionQueryName,omitempty"`
	VersioningScheme  *string                `json:"versioningScheme,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ApiVersionSet_STATUS{}

// ConvertStatusFrom populates our ApiVersionSet_STATUS from the provided source
func (versionSet *ApiVersionSet_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == versionSet {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(versionSet)
}

// ConvertStatusTo populates the provided destination from our ApiVersionSet_STATUS
func (versionSet *ApiVersionSet_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == versionSet {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(versionSet)
}

// Storage version of v1api20220801.ApiVersionSetOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ApiVersionSetOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ApiVersionSet{}, &ApiVersionSetList{})
}
