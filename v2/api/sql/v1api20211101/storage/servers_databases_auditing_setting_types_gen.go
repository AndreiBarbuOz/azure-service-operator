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

// +kubebuilder:rbac:groups=sql.azure.com,resources=serversdatabasesauditingsettings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sql.azure.com,resources={serversdatabasesauditingsettings/status,serversdatabasesauditingsettings/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20211101.ServersDatabasesAuditingSetting
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/BlobAuditing.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/auditingSettings/default
type ServersDatabasesAuditingSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServersDatabasesAuditingSetting_Spec   `json:"spec,omitempty"`
	Status            ServersDatabasesAuditingSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersDatabasesAuditingSetting{}

// GetConditions returns the conditions of the resource
func (setting *ServersDatabasesAuditingSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *ServersDatabasesAuditingSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ configmaps.Exporter = &ServersDatabasesAuditingSetting{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (setting *ServersDatabasesAuditingSetting) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if setting.Spec.OperatorSpec == nil {
		return nil
	}
	return setting.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &ServersDatabasesAuditingSetting{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (setting *ServersDatabasesAuditingSetting) SecretDestinationExpressions() []*core.DestinationExpression {
	if setting.Spec.OperatorSpec == nil {
		return nil
	}
	return setting.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &ServersDatabasesAuditingSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (setting *ServersDatabasesAuditingSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (setting ServersDatabasesAuditingSetting) GetAPIVersion() string {
	return "2021-11-01"
}

// GetResourceScope returns the scope of the resource
func (setting *ServersDatabasesAuditingSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *ServersDatabasesAuditingSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *ServersDatabasesAuditingSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (setting *ServersDatabasesAuditingSetting) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/auditingSettings"
func (setting *ServersDatabasesAuditingSetting) GetType() string {
	return "Microsoft.Sql/servers/databases/auditingSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *ServersDatabasesAuditingSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ServersDatabasesAuditingSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (setting *ServersDatabasesAuditingSetting) Owner() *genruntime.ResourceReference {
	if setting.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return setting.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (setting *ServersDatabasesAuditingSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ServersDatabasesAuditingSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st ServersDatabasesAuditingSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// Hub marks that this ServersDatabasesAuditingSetting is the hub type for conversion
func (setting *ServersDatabasesAuditingSetting) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *ServersDatabasesAuditingSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion,
		Kind:    "ServersDatabasesAuditingSetting",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20211101.ServersDatabasesAuditingSetting
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/BlobAuditing.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/auditingSettings/default
type ServersDatabasesAuditingSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersDatabasesAuditingSetting `json:"items"`
}

// Storage version of v1api20211101.ServersDatabasesAuditingSetting_Spec
type ServersDatabasesAuditingSetting_Spec struct {
	AuditActionsAndGroups       []string                                     `json:"auditActionsAndGroups,omitempty"`
	IsAzureMonitorTargetEnabled *bool                                        `json:"isAzureMonitorTargetEnabled,omitempty"`
	IsManagedIdentityInUse      *bool                                        `json:"isManagedIdentityInUse,omitempty"`
	IsStorageSecondaryKeyInUse  *bool                                        `json:"isStorageSecondaryKeyInUse,omitempty"`
	OperatorSpec                *ServersDatabasesAuditingSettingOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion             string                                       `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/ServersDatabase resource
	Owner                        *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"ServersDatabase"`
	PropertyBag                  genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	QueueDelayMs                 *int                               `json:"queueDelayMs,omitempty"`
	RetentionDays                *int                               `json:"retentionDays,omitempty"`
	State                        *string                            `json:"state,omitempty"`
	StorageAccountAccessKey      *genruntime.SecretReference        `json:"storageAccountAccessKey,omitempty"`
	StorageAccountSubscriptionId *string                            `json:"storageAccountSubscriptionId,omitempty"`
	StorageEndpoint              *string                            `json:"storageEndpoint,omitempty"`
}

var _ genruntime.ConvertibleSpec = &ServersDatabasesAuditingSetting_Spec{}

// ConvertSpecFrom populates our ServersDatabasesAuditingSetting_Spec from the provided source
func (setting *ServersDatabasesAuditingSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == setting {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(setting)
}

// ConvertSpecTo populates the provided destination from our ServersDatabasesAuditingSetting_Spec
func (setting *ServersDatabasesAuditingSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == setting {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(setting)
}

// Storage version of v1api20211101.ServersDatabasesAuditingSetting_STATUS
type ServersDatabasesAuditingSetting_STATUS struct {
	AuditActionsAndGroups        []string               `json:"auditActionsAndGroups,omitempty"`
	Conditions                   []conditions.Condition `json:"conditions,omitempty"`
	Id                           *string                `json:"id,omitempty"`
	IsAzureMonitorTargetEnabled  *bool                  `json:"isAzureMonitorTargetEnabled,omitempty"`
	IsManagedIdentityInUse       *bool                  `json:"isManagedIdentityInUse,omitempty"`
	IsStorageSecondaryKeyInUse   *bool                  `json:"isStorageSecondaryKeyInUse,omitempty"`
	Kind                         *string                `json:"kind,omitempty"`
	Name                         *string                `json:"name,omitempty"`
	PropertyBag                  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	QueueDelayMs                 *int                   `json:"queueDelayMs,omitempty"`
	RetentionDays                *int                   `json:"retentionDays,omitempty"`
	State                        *string                `json:"state,omitempty"`
	StorageAccountSubscriptionId *string                `json:"storageAccountSubscriptionId,omitempty"`
	StorageEndpoint              *string                `json:"storageEndpoint,omitempty"`
	Type                         *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ServersDatabasesAuditingSetting_STATUS{}

// ConvertStatusFrom populates our ServersDatabasesAuditingSetting_STATUS from the provided source
func (setting *ServersDatabasesAuditingSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == setting {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(setting)
}

// ConvertStatusTo populates the provided destination from our ServersDatabasesAuditingSetting_STATUS
func (setting *ServersDatabasesAuditingSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == setting {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(setting)
}

// Storage version of v1api20211101.ServersDatabasesAuditingSettingOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ServersDatabasesAuditingSettingOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ServersDatabasesAuditingSetting{}, &ServersDatabasesAuditingSettingList{})
}
