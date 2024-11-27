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

// +kubebuilder:rbac:groups=eventhub.azure.com,resources=namespaceseventhubsconsumergroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eventhub.azure.com,resources={namespaceseventhubsconsumergroups/status,namespaceseventhubsconsumergroups/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20211101.NamespacesEventhubsConsumerGroup
// Generator information:
// - Generated from: /eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/consumergroups.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}
type NamespacesEventhubsConsumerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesEventhubsConsumerGroup_Spec   `json:"spec,omitempty"`
	Status            NamespacesEventhubsConsumerGroup_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesEventhubsConsumerGroup{}

// GetConditions returns the conditions of the resource
func (group *NamespacesEventhubsConsumerGroup) GetConditions() conditions.Conditions {
	return group.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (group *NamespacesEventhubsConsumerGroup) SetConditions(conditions conditions.Conditions) {
	group.Status.Conditions = conditions
}

var _ configmaps.Exporter = &NamespacesEventhubsConsumerGroup{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (group *NamespacesEventhubsConsumerGroup) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if group.Spec.OperatorSpec == nil {
		return nil
	}
	return group.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &NamespacesEventhubsConsumerGroup{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (group *NamespacesEventhubsConsumerGroup) SecretDestinationExpressions() []*core.DestinationExpression {
	if group.Spec.OperatorSpec == nil {
		return nil
	}
	return group.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &NamespacesEventhubsConsumerGroup{}

// AzureName returns the Azure name of the resource
func (group *NamespacesEventhubsConsumerGroup) AzureName() string {
	return group.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (group NamespacesEventhubsConsumerGroup) GetAPIVersion() string {
	return "2021-11-01"
}

// GetResourceScope returns the scope of the resource
func (group *NamespacesEventhubsConsumerGroup) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (group *NamespacesEventhubsConsumerGroup) GetSpec() genruntime.ConvertibleSpec {
	return &group.Spec
}

// GetStatus returns the status of this resource
func (group *NamespacesEventhubsConsumerGroup) GetStatus() genruntime.ConvertibleStatus {
	return &group.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (group *NamespacesEventhubsConsumerGroup) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs/consumergroups"
func (group *NamespacesEventhubsConsumerGroup) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs/consumergroups"
}

// NewEmptyStatus returns a new empty (blank) status
func (group *NamespacesEventhubsConsumerGroup) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &NamespacesEventhubsConsumerGroup_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (group *NamespacesEventhubsConsumerGroup) Owner() *genruntime.ResourceReference {
	if group.Spec.Owner == nil {
		return nil
	}

	ownerGroup, ownerKind := genruntime.LookupOwnerGroupKind(group.Spec)
	return group.Spec.Owner.AsResourceReference(ownerGroup, ownerKind)
}

// SetStatus sets the status of this resource
func (group *NamespacesEventhubsConsumerGroup) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*NamespacesEventhubsConsumerGroup_STATUS); ok {
		group.Status = *st
		return nil
	}

	// Convert status to required version
	var st NamespacesEventhubsConsumerGroup_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	group.Status = st
	return nil
}

// Hub marks that this NamespacesEventhubsConsumerGroup is the hub type for conversion
func (group *NamespacesEventhubsConsumerGroup) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (group *NamespacesEventhubsConsumerGroup) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: group.Spec.OriginalVersion,
		Kind:    "NamespacesEventhubsConsumerGroup",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20211101.NamespacesEventhubsConsumerGroup
// Generator information:
// - Generated from: /eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/consumergroups.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}
type NamespacesEventhubsConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesEventhubsConsumerGroup `json:"items"`
}

// Storage version of v1api20211101.NamespacesEventhubsConsumerGroup_Spec
type NamespacesEventhubsConsumerGroup_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                                        `json:"azureName,omitempty"`
	OperatorSpec    *NamespacesEventhubsConsumerGroupOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventhub.azure.com/NamespacesEventhub resource
	Owner        *genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner,omitempty" kind:"NamespacesEventhub"`
	PropertyBag  genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	UserMetadata *string                            `json:"userMetadata,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NamespacesEventhubsConsumerGroup_Spec{}

// ConvertSpecFrom populates our NamespacesEventhubsConsumerGroup_Spec from the provided source
func (group *NamespacesEventhubsConsumerGroup_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == group {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(group)
}

// ConvertSpecTo populates the provided destination from our NamespacesEventhubsConsumerGroup_Spec
func (group *NamespacesEventhubsConsumerGroup_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == group {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(group)
}

// Storage version of v1api20211101.NamespacesEventhubsConsumerGroup_STATUS
type NamespacesEventhubsConsumerGroup_STATUS struct {
	Conditions   []conditions.Condition `json:"conditions,omitempty"`
	CreatedAt    *string                `json:"createdAt,omitempty"`
	Id           *string                `json:"id,omitempty"`
	Location     *string                `json:"location,omitempty"`
	Name         *string                `json:"name,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SystemData   *SystemData_STATUS     `json:"systemData,omitempty"`
	Type         *string                `json:"type,omitempty"`
	UpdatedAt    *string                `json:"updatedAt,omitempty"`
	UserMetadata *string                `json:"userMetadata,omitempty"`
}

var _ genruntime.ConvertibleStatus = &NamespacesEventhubsConsumerGroup_STATUS{}

// ConvertStatusFrom populates our NamespacesEventhubsConsumerGroup_STATUS from the provided source
func (group *NamespacesEventhubsConsumerGroup_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == group {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(group)
}

// ConvertStatusTo populates the provided destination from our NamespacesEventhubsConsumerGroup_STATUS
func (group *NamespacesEventhubsConsumerGroup_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == group {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(group)
}

// Storage version of v1api20211101.NamespacesEventhubsConsumerGroupOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type NamespacesEventhubsConsumerGroupOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NamespacesEventhubsConsumerGroup{}, &NamespacesEventhubsConsumerGroupList{})
}
