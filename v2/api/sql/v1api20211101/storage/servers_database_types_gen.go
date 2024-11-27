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

// +kubebuilder:rbac:groups=sql.azure.com,resources=serversdatabases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sql.azure.com,resources={serversdatabases/status,serversdatabases/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20211101.ServersDatabase
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/Databases.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}
type ServersDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServersDatabase_Spec   `json:"spec,omitempty"`
	Status            ServersDatabase_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersDatabase{}

// GetConditions returns the conditions of the resource
func (database *ServersDatabase) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *ServersDatabase) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ configmaps.Exporter = &ServersDatabase{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (database *ServersDatabase) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if database.Spec.OperatorSpec == nil {
		return nil
	}
	return database.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &ServersDatabase{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (database *ServersDatabase) SecretDestinationExpressions() []*core.DestinationExpression {
	if database.Spec.OperatorSpec == nil {
		return nil
	}
	return database.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &ServersDatabase{}

// AzureName returns the Azure name of the resource
func (database *ServersDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (database ServersDatabase) GetAPIVersion() string {
	return "2021-11-01"
}

// GetResourceScope returns the scope of the resource
func (database *ServersDatabase) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (database *ServersDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *ServersDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (database *ServersDatabase) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases"
func (database *ServersDatabase) GetType() string {
	return "Microsoft.Sql/servers/databases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *ServersDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ServersDatabase_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (database *ServersDatabase) Owner() *genruntime.ResourceReference {
	if database.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return database.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (database *ServersDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ServersDatabase_STATUS); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st ServersDatabase_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// Hub marks that this ServersDatabase is the hub type for conversion
func (database *ServersDatabase) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *ServersDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion,
		Kind:    "ServersDatabase",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20211101.ServersDatabase
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/Databases.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}
type ServersDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersDatabase `json:"items"`
}

// Storage version of v1api20211101.ServersDatabase_Spec
type ServersDatabase_Spec struct {
	AutoPauseDelay *int `json:"autoPauseDelay,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string  `json:"azureName,omitempty"`
	CatalogCollation *string `json:"catalogCollation,omitempty"`
	Collation        *string `json:"collation,omitempty"`
	CreateMode       *string `json:"createMode,omitempty"`

	// ElasticPoolReference: The resource identifier of the elastic pool containing this database.
	ElasticPoolReference         *genruntime.ResourceReference `armReference:"ElasticPoolId" json:"elasticPoolReference,omitempty"`
	FederatedClientId            *string                       `json:"federatedClientId,omitempty"`
	HighAvailabilityReplicaCount *int                          `json:"highAvailabilityReplicaCount,omitempty"`
	Identity                     *DatabaseIdentity             `json:"identity,omitempty"`
	IsLedgerOn                   *bool                         `json:"isLedgerOn,omitempty"`
	LicenseType                  *string                       `json:"licenseType,omitempty"`
	Location                     *string                       `json:"location,omitempty"`

	// LongTermRetentionBackupResourceReference: The resource identifier of the long term retention backup associated with
	// create operation of this database.
	LongTermRetentionBackupResourceReference *genruntime.ResourceReference `armReference:"LongTermRetentionBackupResourceId" json:"longTermRetentionBackupResourceReference,omitempty"`
	MaintenanceConfigurationId               *string                       `json:"maintenanceConfigurationId,omitempty"`
	MaxSizeBytes                             *int                          `json:"maxSizeBytes,omitempty"`
	MinCapacity                              *float64                      `json:"minCapacity,omitempty"`
	OperatorSpec                             *ServersDatabaseOperatorSpec  `json:"operatorSpec,omitempty"`
	OriginalVersion                          string                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/Server resource
	Owner       *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"Server"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ReadScale   *string                            `json:"readScale,omitempty"`

	// RecoverableDatabaseReference: The resource identifier of the recoverable database associated with create operation of
	// this database.
	RecoverableDatabaseReference *genruntime.ResourceReference `armReference:"RecoverableDatabaseId" json:"recoverableDatabaseReference,omitempty"`

	// RecoveryServicesRecoveryPointReference: The resource identifier of the recovery point associated with create operation
	// of this database.
	RecoveryServicesRecoveryPointReference *genruntime.ResourceReference `armReference:"RecoveryServicesRecoveryPointId" json:"recoveryServicesRecoveryPointReference,omitempty"`
	RequestedBackupStorageRedundancy       *string                       `json:"requestedBackupStorageRedundancy,omitempty"`

	// RestorableDroppedDatabaseReference: The resource identifier of the restorable dropped database associated with create
	// operation of this database.
	RestorableDroppedDatabaseReference *genruntime.ResourceReference `armReference:"RestorableDroppedDatabaseId" json:"restorableDroppedDatabaseReference,omitempty"`
	RestorePointInTime                 *string                       `json:"restorePointInTime,omitempty"`
	SampleName                         *string                       `json:"sampleName,omitempty"`
	SecondaryType                      *string                       `json:"secondaryType,omitempty"`
	Sku                                *Sku                          `json:"sku,omitempty"`
	SourceDatabaseDeletionDate         *string                       `json:"sourceDatabaseDeletionDate,omitempty"`

	// SourceDatabaseReference: The resource identifier of the source database associated with create operation of this
	// database.
	SourceDatabaseReference *genruntime.ResourceReference `armReference:"SourceDatabaseId" json:"sourceDatabaseReference,omitempty"`

	// SourceResourceReference: The resource identifier of the source associated with the create operation of this database.
	// This property is only supported for DataWarehouse edition and allows to restore across subscriptions.
	// When sourceResourceId is specified, sourceDatabaseId, recoverableDatabaseId, restorableDroppedDatabaseId and
	// sourceDatabaseDeletionDate must not be specified and CreateMode must be PointInTimeRestore, Restore or Recover.
	// When createMode is PointInTimeRestore, sourceResourceId must be the resource ID of the existing database or existing sql
	// pool, and restorePointInTime must be specified.
	// When createMode is Restore, sourceResourceId must be the resource ID of restorable dropped database or restorable
	// dropped sql pool.
	// When createMode is Recover, sourceResourceId must be the resource ID of recoverable database or recoverable sql pool.
	// When source subscription belongs to a different tenant than target subscription, “x-ms-authorization-auxiliary”
	// header must contain authentication token for the source tenant. For more details about
	// “x-ms-authorization-auxiliary” header see
	// https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/authenticate-multi-tenant
	SourceResourceReference *genruntime.ResourceReference `armReference:"SourceResourceId" json:"sourceResourceReference,omitempty"`
	Tags                    map[string]string             `json:"tags,omitempty"`
	ZoneRedundant           *bool                         `json:"zoneRedundant,omitempty"`
}

var _ genruntime.ConvertibleSpec = &ServersDatabase_Spec{}

// ConvertSpecFrom populates our ServersDatabase_Spec from the provided source
func (database *ServersDatabase_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == database {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(database)
}

// ConvertSpecTo populates the provided destination from our ServersDatabase_Spec
func (database *ServersDatabase_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == database {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(database)
}

// Storage version of v1api20211101.ServersDatabase_STATUS
type ServersDatabase_STATUS struct {
	AutoPauseDelay                    *int                     `json:"autoPauseDelay,omitempty"`
	CatalogCollation                  *string                  `json:"catalogCollation,omitempty"`
	Collation                         *string                  `json:"collation,omitempty"`
	Conditions                        []conditions.Condition   `json:"conditions,omitempty"`
	CreateMode                        *string                  `json:"createMode,omitempty"`
	CreationDate                      *string                  `json:"creationDate,omitempty"`
	CurrentBackupStorageRedundancy    *string                  `json:"currentBackupStorageRedundancy,omitempty"`
	CurrentServiceObjectiveName       *string                  `json:"currentServiceObjectiveName,omitempty"`
	CurrentSku                        *Sku_STATUS              `json:"currentSku,omitempty"`
	DatabaseId                        *string                  `json:"databaseId,omitempty"`
	DefaultSecondaryLocation          *string                  `json:"defaultSecondaryLocation,omitempty"`
	EarliestRestoreDate               *string                  `json:"earliestRestoreDate,omitempty"`
	ElasticPoolId                     *string                  `json:"elasticPoolId,omitempty"`
	FailoverGroupId                   *string                  `json:"failoverGroupId,omitempty"`
	FederatedClientId                 *string                  `json:"federatedClientId,omitempty"`
	HighAvailabilityReplicaCount      *int                     `json:"highAvailabilityReplicaCount,omitempty"`
	Id                                *string                  `json:"id,omitempty"`
	Identity                          *DatabaseIdentity_STATUS `json:"identity,omitempty"`
	IsInfraEncryptionEnabled          *bool                    `json:"isInfraEncryptionEnabled,omitempty"`
	IsLedgerOn                        *bool                    `json:"isLedgerOn,omitempty"`
	Kind                              *string                  `json:"kind,omitempty"`
	LicenseType                       *string                  `json:"licenseType,omitempty"`
	Location                          *string                  `json:"location,omitempty"`
	LongTermRetentionBackupResourceId *string                  `json:"longTermRetentionBackupResourceId,omitempty"`
	MaintenanceConfigurationId        *string                  `json:"maintenanceConfigurationId,omitempty"`
	ManagedBy                         *string                  `json:"managedBy,omitempty"`
	MaxLogSizeBytes                   *int                     `json:"maxLogSizeBytes,omitempty"`
	MaxSizeBytes                      *int                     `json:"maxSizeBytes,omitempty"`
	MinCapacity                       *float64                 `json:"minCapacity,omitempty"`
	Name                              *string                  `json:"name,omitempty"`
	PausedDate                        *string                  `json:"pausedDate,omitempty"`
	PropertyBag                       genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
	ReadScale                         *string                  `json:"readScale,omitempty"`
	RecoverableDatabaseId             *string                  `json:"recoverableDatabaseId,omitempty"`
	RecoveryServicesRecoveryPointId   *string                  `json:"recoveryServicesRecoveryPointId,omitempty"`
	RequestedBackupStorageRedundancy  *string                  `json:"requestedBackupStorageRedundancy,omitempty"`
	RequestedServiceObjectiveName     *string                  `json:"requestedServiceObjectiveName,omitempty"`
	RestorableDroppedDatabaseId       *string                  `json:"restorableDroppedDatabaseId,omitempty"`
	RestorePointInTime                *string                  `json:"restorePointInTime,omitempty"`
	ResumedDate                       *string                  `json:"resumedDate,omitempty"`
	SampleName                        *string                  `json:"sampleName,omitempty"`
	SecondaryType                     *string                  `json:"secondaryType,omitempty"`
	Sku                               *Sku_STATUS              `json:"sku,omitempty"`
	SourceDatabaseDeletionDate        *string                  `json:"sourceDatabaseDeletionDate,omitempty"`
	SourceDatabaseId                  *string                  `json:"sourceDatabaseId,omitempty"`
	SourceResourceId                  *string                  `json:"sourceResourceId,omitempty"`
	Status                            *string                  `json:"status,omitempty"`
	Tags                              map[string]string        `json:"tags,omitempty"`
	Type                              *string                  `json:"type,omitempty"`
	ZoneRedundant                     *bool                    `json:"zoneRedundant,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ServersDatabase_STATUS{}

// ConvertStatusFrom populates our ServersDatabase_STATUS from the provided source
func (database *ServersDatabase_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == database {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(database)
}

// ConvertStatusTo populates the provided destination from our ServersDatabase_STATUS
func (database *ServersDatabase_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == database {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(database)
}

// Storage version of v1api20211101.DatabaseIdentity
// Azure Active Directory identity configuration for a resource.
type DatabaseIdentity struct {
	PropertyBag            genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Type                   *string                       `json:"type,omitempty"`
	UserAssignedIdentities []UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20211101.DatabaseIdentity_STATUS
// Azure Active Directory identity configuration for a resource.
type DatabaseIdentity_STATUS struct {
	PropertyBag            genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	TenantId               *string                                `json:"tenantId,omitempty"`
	Type                   *string                                `json:"type,omitempty"`
	UserAssignedIdentities map[string]DatabaseUserIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20211101.ServersDatabaseOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ServersDatabaseOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20211101.Sku
// An ARM Resource SKU.
type Sku struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Size        *string                `json:"size,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1api20211101.Sku_STATUS
// An ARM Resource SKU.
type Sku_STATUS struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Size        *string                `json:"size,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1api20211101.DatabaseUserIdentity_STATUS
// Azure Active Directory identity configuration for a resource.
type DatabaseUserIdentity_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ServersDatabase{}, &ServersDatabaseList{})
}
