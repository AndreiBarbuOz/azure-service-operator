// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101preview

import (
	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:rbac:groups=sql.azure.com,resources=servers_databases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sql.azure.com,resources={servers_databases/status,servers_databases/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/Databases_legacy.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}
type Servers_Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_Database_Spec   `json:"spec,omitempty"`
	Status            Servers_Database_STATUS `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/Databases_legacy.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}
type Servers_DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Servers_Database `json:"items"`
}

type Servers_Database_Spec struct {
	v1alpha1.ResourceSpec `json:",inline,omitempty"`
	ForProvider           Servers_DatabaseParameters `json:"forProvider,omitempty"`
}

type Servers_Database_STATUS struct {
	v1alpha1.ResourceStatus `json:",inline,omitempty"`
	AtProvider              ServersObservation `json:"atProvider,omitempty"`
}

type Servers_DatabaseParameters struct {
	// AutoPauseDelay: Time in minutes after which database is automatically paused. A value of -1 means that automatic pause
	// is disabled
	AutoPauseDelay *int `json:"autoPauseDelay,omitempty"`

	// CatalogCollation: Collation of the metadata catalog.
	CatalogCollation *DatabaseProperties_CatalogCollation `json:"catalogCollation,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`

	// CreateMode: Specifies the mode of database creation.
	// Default: regular database creation.
	// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the
	// source database.
	// Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the
	// resource ID of the existing primary database.
	// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId
	// must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable
	// database resource ID to restore.
	// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If
	// sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise
	// sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored.
	// restorePointInTime may also be specified to restore from an earlier point in time.
	// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault.
	// recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	// Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode *DatabaseProperties_CreateMode `json:"createMode,omitempty"`

	// ElasticPoolId: The resource identifier of the elastic pool containing this database.
	ElasticPoolId *string `json:"elasticPoolId,omitempty"`

	// HighAvailabilityReplicaCount: The number of secondary replicas associated with the database that are used to provide
	// high availability.
	HighAvailabilityReplicaCount *int `json:"highAvailabilityReplicaCount,omitempty"`

	// LicenseType: The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you
	// have a license and are eligible for the Azure Hybrid Benefit.
	LicenseType *DatabaseProperties_LicenseType `json:"licenseType,omitempty"`

	// +kubebuilder:validation:Required
	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// LongTermRetentionBackupResourceId: The resource identifier of the long term retention backup associated with create
	// operation of this database.
	LongTermRetentionBackupResourceId *string `json:"longTermRetentionBackupResourceId,omitempty"`

	// MaintenanceConfigurationId: Maintenance configuration id assigned to the database. This configuration defines the period
	// when the maintenance updates will occur.
	MaintenanceConfigurationId *string `json:"maintenanceConfigurationId,omitempty"`

	// MaxSizeBytes: The max size of the database expressed in bytes.
	MaxSizeBytes *int `json:"maxSizeBytes,omitempty"`

	// MinCapacity: Minimal capacity that database will always have allocated, if not paused
	MinCapacity *float64 `json:"minCapacity,omitempty"`
	Name        string   `json:"name,omitempty"`

	// ReadScale: The state of read-only routing. If enabled, connections that have application intent set to readonly in their
	// connection string may be routed to a readonly secondary replica in the same region.
	ReadScale *DatabaseProperties_ReadScale `json:"readScale,omitempty"`

	// RecoverableDatabaseId: The resource identifier of the recoverable database associated with create operation of this
	// database.
	RecoverableDatabaseId *string `json:"recoverableDatabaseId,omitempty"`

	// RecoveryServicesRecoveryPointId: The resource identifier of the recovery point associated with create operation of this
	// database.
	RecoveryServicesRecoveryPointId *string `json:"recoveryServicesRecoveryPointId,omitempty"`

	// RequestedBackupStorageRedundancy: The storage account type to be used to store backups for this database.
	RequestedBackupStorageRedundancy *DatabaseProperties_RequestedBackupStorageRedundancy `json:"requestedBackupStorageRedundancy,omitempty"`
	ResourceGroupName                string                                               `json:"resourceGroupName,omitempty"`
	ResourceGroupNameRef             *v1alpha1.Reference                                  `json:"resourceGroupNameRef,omitempty"`
	ResourceGroupNameSelector        *v1alpha1.Selector                                   `json:"resourceGroupNameSelector,omitempty"`

	// RestorableDroppedDatabaseId: The resource identifier of the restorable dropped database associated with create operation
	// of this database.
	RestorableDroppedDatabaseId *string `json:"restorableDroppedDatabaseId,omitempty"`

	// RestorePointInTime: Specifies the point in time (ISO8601 format) of the source database that will be restored to create
	// the new database.
	RestorePointInTime *string `json:"restorePointInTime,omitempty"`

	// SampleName: The name of the sample schema to apply when creating this database.
	SampleName *DatabaseProperties_SampleName `json:"sampleName,omitempty"`

	// SecondaryType: The secondary type of the database if it is a secondary.  Valid values are Geo and Named.
	SecondaryType      *DatabaseProperties_SecondaryType `json:"secondaryType,omitempty"`
	ServerName         string                            `json:"serverName,omitempty"`
	ServerNameRef      *v1alpha1.Reference               `json:"serverNameRef,omitempty"`
	ServerNameSelector *v1alpha1.Selector                `json:"serverNameSelector,omitempty"`

	// Sku: The database SKU.
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition,
	// family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation`
	// REST API or one of the following commands:
	// ```azurecli
	// az sql db list-editions -l <location> -o table
	// ````
	// ```powershell
	// Get-AzSqlServerServiceObjective -Location <location>
	// ````
	Sku *Sku `json:"sku,omitempty"`

	// SourceDatabaseDeletionDate: Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate *string `json:"sourceDatabaseDeletionDate,omitempty"`

	// SourceDatabaseId: The resource identifier of the source database associated with create operation of this database.
	SourceDatabaseId *string `json:"sourceDatabaseId,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// ZoneRedundant: Whether or not this database is zone redundant, which means the replicas of this database will be spread
	// across multiple availability zones.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

type ServersObservation struct {
	// AutoPauseDelay: Time in minutes after which database is automatically paused. A value of -1 means that automatic pause
	// is disabled
	AutoPauseDelay *int `json:"autoPauseDelay,omitempty"`

	// CatalogCollation: Collation of the metadata catalog.
	CatalogCollation *DatabaseProperties_CatalogCollation_STATUS `json:"catalogCollation,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`

	// CreateMode: Specifies the mode of database creation.
	// Default: regular database creation.
	// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the
	// source database.
	// Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the
	// resource ID of the existing primary database.
	// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId
	// must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable
	// database resource ID to restore.
	// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If
	// sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise
	// sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored.
	// restorePointInTime may also be specified to restore from an earlier point in time.
	// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault.
	// recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	// Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode *DatabaseProperties_CreateMode_STATUS `json:"createMode,omitempty"`

	// CreationDate: The creation date of the database (ISO8601 format).
	CreationDate *string `json:"creationDate,omitempty"`

	// CurrentBackupStorageRedundancy: The storage account type used to store backups for this database.
	CurrentBackupStorageRedundancy *DatabaseProperties_CurrentBackupStorageRedundancy_STATUS `json:"currentBackupStorageRedundancy,omitempty"`

	// CurrentServiceObjectiveName: The current service level objective name of the database.
	CurrentServiceObjectiveName *string `json:"currentServiceObjectiveName,omitempty"`

	// CurrentSku: The name and tier of the SKU.
	CurrentSku *Sku_STATUS `json:"currentSku,omitempty"`

	// DatabaseId: The ID of the database.
	DatabaseId *string `json:"databaseId,omitempty"`

	// DefaultSecondaryLocation: The default secondary region for this database.
	DefaultSecondaryLocation *string `json:"defaultSecondaryLocation,omitempty"`

	// EarliestRestoreDate: This records the earliest start date and time that restore is available for this database (ISO8601
	// format).
	EarliestRestoreDate *string `json:"earliestRestoreDate,omitempty"`

	// ElasticPoolId: The resource identifier of the elastic pool containing this database.
	ElasticPoolId *string `json:"elasticPoolId,omitempty"`

	// FailoverGroupId: Failover Group resource identifier that this database belongs to.
	FailoverGroupId *string `json:"failoverGroupId,omitempty"`

	// HighAvailabilityReplicaCount: The number of secondary replicas associated with the database that are used to provide
	// high availability.
	HighAvailabilityReplicaCount *int `json:"highAvailabilityReplicaCount,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Kind: Kind of database. This is metadata used for the Azure portal experience.
	Kind *string `json:"kind,omitempty"`

	// LicenseType: The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you
	// have a license and are eligible for the Azure Hybrid Benefit.
	LicenseType *DatabaseProperties_LicenseType_STATUS `json:"licenseType,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// LongTermRetentionBackupResourceId: The resource identifier of the long term retention backup associated with create
	// operation of this database.
	LongTermRetentionBackupResourceId *string `json:"longTermRetentionBackupResourceId,omitempty"`

	// MaintenanceConfigurationId: Maintenance configuration id assigned to the database. This configuration defines the period
	// when the maintenance updates will occur.
	MaintenanceConfigurationId *string `json:"maintenanceConfigurationId,omitempty"`

	// ManagedBy: Resource that manages the database.
	ManagedBy *string `json:"managedBy,omitempty"`

	// MaxLogSizeBytes: The max log size for this database.
	MaxLogSizeBytes *int `json:"maxLogSizeBytes,omitempty"`

	// MaxSizeBytes: The max size of the database expressed in bytes.
	MaxSizeBytes *int `json:"maxSizeBytes,omitempty"`

	// MinCapacity: Minimal capacity that database will always have allocated, if not paused
	MinCapacity *float64 `json:"minCapacity,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// PausedDate: The date when database was paused by user configuration or action(ISO8601 format). Null if the database is
	// ready.
	PausedDate *string `json:"pausedDate,omitempty"`

	// ReadScale: The state of read-only routing. If enabled, connections that have application intent set to readonly in their
	// connection string may be routed to a readonly secondary replica in the same region.
	ReadScale *DatabaseProperties_ReadScale_STATUS `json:"readScale,omitempty"`

	// RecoverableDatabaseId: The resource identifier of the recoverable database associated with create operation of this
	// database.
	RecoverableDatabaseId *string `json:"recoverableDatabaseId,omitempty"`

	// RecoveryServicesRecoveryPointId: The resource identifier of the recovery point associated with create operation of this
	// database.
	RecoveryServicesRecoveryPointId *string `json:"recoveryServicesRecoveryPointId,omitempty"`

	// RequestedBackupStorageRedundancy: The storage account type to be used to store backups for this database.
	RequestedBackupStorageRedundancy *DatabaseProperties_RequestedBackupStorageRedundancy_STATUS `json:"requestedBackupStorageRedundancy,omitempty"`

	// RequestedServiceObjectiveName: The requested service level objective name of the database.
	RequestedServiceObjectiveName *string `json:"requestedServiceObjectiveName,omitempty"`

	// RestorableDroppedDatabaseId: The resource identifier of the restorable dropped database associated with create operation
	// of this database.
	RestorableDroppedDatabaseId *string `json:"restorableDroppedDatabaseId,omitempty"`

	// RestorePointInTime: Specifies the point in time (ISO8601 format) of the source database that will be restored to create
	// the new database.
	RestorePointInTime *string `json:"restorePointInTime,omitempty"`

	// ResumedDate: The date when database was resumed by user action or database login (ISO8601 format). Null if the database
	// is paused.
	ResumedDate *string `json:"resumedDate,omitempty"`

	// SampleName: The name of the sample schema to apply when creating this database.
	SampleName *DatabaseProperties_SampleName_STATUS `json:"sampleName,omitempty"`

	// SecondaryType: The secondary type of the database if it is a secondary.  Valid values are Geo and Named.
	SecondaryType *DatabaseProperties_SecondaryType_STATUS `json:"secondaryType,omitempty"`

	// Sku: The database SKU.
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition,
	// family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation`
	// REST API or one of the following commands:
	// ```azurecli
	// az sql db list-editions -l <location> -o table
	// ````
	// ```powershell
	// Get-AzSqlServerServiceObjective -Location <location>
	// ````
	Sku *Sku_STATUS `json:"sku,omitempty"`

	// SourceDatabaseDeletionDate: Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate *string `json:"sourceDatabaseDeletionDate,omitempty"`

	// SourceDatabaseId: The resource identifier of the source database associated with create operation of this database.
	SourceDatabaseId *string `json:"sourceDatabaseId,omitempty"`

	// Status: The status of the database.
	Status *DatabaseProperties_Status_STATUS `json:"status,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// ZoneRedundant: Whether or not this database is zone redundant, which means the replicas of this database will be spread
	// across multiple availability zones.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

// +kubebuilder:validation:Enum={"DATABASE_DEFAULT","SQL_Latin1_General_CP1_CI_AS"}
type DatabaseProperties_CatalogCollation string

const (
	DatabaseProperties_CatalogCollation_DATABASE_DEFAULT             = DatabaseProperties_CatalogCollation("DATABASE_DEFAULT")
	DatabaseProperties_CatalogCollation_SQL_Latin1_General_CP1_CI_AS = DatabaseProperties_CatalogCollation("SQL_Latin1_General_CP1_CI_AS")
)

type DatabaseProperties_CatalogCollation_STATUS string

const (
	DatabaseProperties_CatalogCollation_STATUS_DATABASE_DEFAULT             = DatabaseProperties_CatalogCollation_STATUS("DATABASE_DEFAULT")
	DatabaseProperties_CatalogCollation_STATUS_SQL_Latin1_General_CP1_CI_AS = DatabaseProperties_CatalogCollation_STATUS("SQL_Latin1_General_CP1_CI_AS")
)

// +kubebuilder:validation:Enum={"Copy","Default","OnlineSecondary","PointInTimeRestore","Recovery","Restore","RestoreExternalBackup","RestoreExternalBackupSecondary","RestoreLongTermRetentionBackup","Secondary"}
type DatabaseProperties_CreateMode string

const (
	DatabaseProperties_CreateMode_Copy                           = DatabaseProperties_CreateMode("Copy")
	DatabaseProperties_CreateMode_Default                        = DatabaseProperties_CreateMode("Default")
	DatabaseProperties_CreateMode_OnlineSecondary                = DatabaseProperties_CreateMode("OnlineSecondary")
	DatabaseProperties_CreateMode_PointInTimeRestore             = DatabaseProperties_CreateMode("PointInTimeRestore")
	DatabaseProperties_CreateMode_Recovery                       = DatabaseProperties_CreateMode("Recovery")
	DatabaseProperties_CreateMode_Restore                        = DatabaseProperties_CreateMode("Restore")
	DatabaseProperties_CreateMode_RestoreExternalBackup          = DatabaseProperties_CreateMode("RestoreExternalBackup")
	DatabaseProperties_CreateMode_RestoreExternalBackupSecondary = DatabaseProperties_CreateMode("RestoreExternalBackupSecondary")
	DatabaseProperties_CreateMode_RestoreLongTermRetentionBackup = DatabaseProperties_CreateMode("RestoreLongTermRetentionBackup")
	DatabaseProperties_CreateMode_Secondary                      = DatabaseProperties_CreateMode("Secondary")
)

type DatabaseProperties_CreateMode_STATUS string

const (
	DatabaseProperties_CreateMode_STATUS_Copy                           = DatabaseProperties_CreateMode_STATUS("Copy")
	DatabaseProperties_CreateMode_STATUS_Default                        = DatabaseProperties_CreateMode_STATUS("Default")
	DatabaseProperties_CreateMode_STATUS_OnlineSecondary                = DatabaseProperties_CreateMode_STATUS("OnlineSecondary")
	DatabaseProperties_CreateMode_STATUS_PointInTimeRestore             = DatabaseProperties_CreateMode_STATUS("PointInTimeRestore")
	DatabaseProperties_CreateMode_STATUS_Recovery                       = DatabaseProperties_CreateMode_STATUS("Recovery")
	DatabaseProperties_CreateMode_STATUS_Restore                        = DatabaseProperties_CreateMode_STATUS("Restore")
	DatabaseProperties_CreateMode_STATUS_RestoreExternalBackup          = DatabaseProperties_CreateMode_STATUS("RestoreExternalBackup")
	DatabaseProperties_CreateMode_STATUS_RestoreExternalBackupSecondary = DatabaseProperties_CreateMode_STATUS("RestoreExternalBackupSecondary")
	DatabaseProperties_CreateMode_STATUS_RestoreLongTermRetentionBackup = DatabaseProperties_CreateMode_STATUS("RestoreLongTermRetentionBackup")
	DatabaseProperties_CreateMode_STATUS_Secondary                      = DatabaseProperties_CreateMode_STATUS("Secondary")
)

type DatabaseProperties_CurrentBackupStorageRedundancy_STATUS string

const (
	DatabaseProperties_CurrentBackupStorageRedundancy_STATUS_Geo   = DatabaseProperties_CurrentBackupStorageRedundancy_STATUS("Geo")
	DatabaseProperties_CurrentBackupStorageRedundancy_STATUS_Local = DatabaseProperties_CurrentBackupStorageRedundancy_STATUS("Local")
	DatabaseProperties_CurrentBackupStorageRedundancy_STATUS_Zone  = DatabaseProperties_CurrentBackupStorageRedundancy_STATUS("Zone")
)

// +kubebuilder:validation:Enum={"BasePrice","LicenseIncluded"}
type DatabaseProperties_LicenseType string

const (
	DatabaseProperties_LicenseType_BasePrice       = DatabaseProperties_LicenseType("BasePrice")
	DatabaseProperties_LicenseType_LicenseIncluded = DatabaseProperties_LicenseType("LicenseIncluded")
)

type DatabaseProperties_LicenseType_STATUS string

const (
	DatabaseProperties_LicenseType_STATUS_BasePrice       = DatabaseProperties_LicenseType_STATUS("BasePrice")
	DatabaseProperties_LicenseType_STATUS_LicenseIncluded = DatabaseProperties_LicenseType_STATUS("LicenseIncluded")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type DatabaseProperties_ReadScale string

const (
	DatabaseProperties_ReadScale_Disabled = DatabaseProperties_ReadScale("Disabled")
	DatabaseProperties_ReadScale_Enabled  = DatabaseProperties_ReadScale("Enabled")
)

type DatabaseProperties_ReadScale_STATUS string

const (
	DatabaseProperties_ReadScale_STATUS_Disabled = DatabaseProperties_ReadScale_STATUS("Disabled")
	DatabaseProperties_ReadScale_STATUS_Enabled  = DatabaseProperties_ReadScale_STATUS("Enabled")
)

// +kubebuilder:validation:Enum={"Geo","Local","Zone"}
type DatabaseProperties_RequestedBackupStorageRedundancy string

const (
	DatabaseProperties_RequestedBackupStorageRedundancy_Geo   = DatabaseProperties_RequestedBackupStorageRedundancy("Geo")
	DatabaseProperties_RequestedBackupStorageRedundancy_Local = DatabaseProperties_RequestedBackupStorageRedundancy("Local")
	DatabaseProperties_RequestedBackupStorageRedundancy_Zone  = DatabaseProperties_RequestedBackupStorageRedundancy("Zone")
)

type DatabaseProperties_RequestedBackupStorageRedundancy_STATUS string

const (
	DatabaseProperties_RequestedBackupStorageRedundancy_STATUS_Geo   = DatabaseProperties_RequestedBackupStorageRedundancy_STATUS("Geo")
	DatabaseProperties_RequestedBackupStorageRedundancy_STATUS_Local = DatabaseProperties_RequestedBackupStorageRedundancy_STATUS("Local")
	DatabaseProperties_RequestedBackupStorageRedundancy_STATUS_Zone  = DatabaseProperties_RequestedBackupStorageRedundancy_STATUS("Zone")
)

// +kubebuilder:validation:Enum={"AdventureWorksLT","WideWorldImportersFull","WideWorldImportersStd"}
type DatabaseProperties_SampleName string

const (
	DatabaseProperties_SampleName_AdventureWorksLT       = DatabaseProperties_SampleName("AdventureWorksLT")
	DatabaseProperties_SampleName_WideWorldImportersFull = DatabaseProperties_SampleName("WideWorldImportersFull")
	DatabaseProperties_SampleName_WideWorldImportersStd  = DatabaseProperties_SampleName("WideWorldImportersStd")
)

type DatabaseProperties_SampleName_STATUS string

const (
	DatabaseProperties_SampleName_STATUS_AdventureWorksLT       = DatabaseProperties_SampleName_STATUS("AdventureWorksLT")
	DatabaseProperties_SampleName_STATUS_WideWorldImportersFull = DatabaseProperties_SampleName_STATUS("WideWorldImportersFull")
	DatabaseProperties_SampleName_STATUS_WideWorldImportersStd  = DatabaseProperties_SampleName_STATUS("WideWorldImportersStd")
)

// +kubebuilder:validation:Enum={"Geo","Named"}
type DatabaseProperties_SecondaryType string

const (
	DatabaseProperties_SecondaryType_Geo   = DatabaseProperties_SecondaryType("Geo")
	DatabaseProperties_SecondaryType_Named = DatabaseProperties_SecondaryType("Named")
)

type DatabaseProperties_SecondaryType_STATUS string

const (
	DatabaseProperties_SecondaryType_STATUS_Geo   = DatabaseProperties_SecondaryType_STATUS("Geo")
	DatabaseProperties_SecondaryType_STATUS_Named = DatabaseProperties_SecondaryType_STATUS("Named")
)

type DatabaseProperties_Status_STATUS string

const (
	DatabaseProperties_Status_STATUS_AutoClosed                        = DatabaseProperties_Status_STATUS("AutoClosed")
	DatabaseProperties_Status_STATUS_Copying                           = DatabaseProperties_Status_STATUS("Copying")
	DatabaseProperties_Status_STATUS_Creating                          = DatabaseProperties_Status_STATUS("Creating")
	DatabaseProperties_Status_STATUS_Disabled                          = DatabaseProperties_Status_STATUS("Disabled")
	DatabaseProperties_Status_STATUS_EmergencyMode                     = DatabaseProperties_Status_STATUS("EmergencyMode")
	DatabaseProperties_Status_STATUS_Inaccessible                      = DatabaseProperties_Status_STATUS("Inaccessible")
	DatabaseProperties_Status_STATUS_Offline                           = DatabaseProperties_Status_STATUS("Offline")
	DatabaseProperties_Status_STATUS_OfflineChangingDwPerformanceTiers = DatabaseProperties_Status_STATUS("OfflineChangingDwPerformanceTiers")
	DatabaseProperties_Status_STATUS_OfflineSecondary                  = DatabaseProperties_Status_STATUS("OfflineSecondary")
	DatabaseProperties_Status_STATUS_Online                            = DatabaseProperties_Status_STATUS("Online")
	DatabaseProperties_Status_STATUS_OnlineChangingDwPerformanceTiers  = DatabaseProperties_Status_STATUS("OnlineChangingDwPerformanceTiers")
	DatabaseProperties_Status_STATUS_Paused                            = DatabaseProperties_Status_STATUS("Paused")
	DatabaseProperties_Status_STATUS_Pausing                           = DatabaseProperties_Status_STATUS("Pausing")
	DatabaseProperties_Status_STATUS_Recovering                        = DatabaseProperties_Status_STATUS("Recovering")
	DatabaseProperties_Status_STATUS_RecoveryPending                   = DatabaseProperties_Status_STATUS("RecoveryPending")
	DatabaseProperties_Status_STATUS_Restoring                         = DatabaseProperties_Status_STATUS("Restoring")
	DatabaseProperties_Status_STATUS_Resuming                          = DatabaseProperties_Status_STATUS("Resuming")
	DatabaseProperties_Status_STATUS_Scaling                           = DatabaseProperties_Status_STATUS("Scaling")
	DatabaseProperties_Status_STATUS_Shutdown                          = DatabaseProperties_Status_STATUS("Shutdown")
	DatabaseProperties_Status_STATUS_Standby                           = DatabaseProperties_Status_STATUS("Standby")
	DatabaseProperties_Status_STATUS_Suspect                           = DatabaseProperties_Status_STATUS("Suspect")
)

// An ARM Resource SKU.
type Sku struct {
	// Capacity: Capacity of the particular SKU.
	Capacity *int `json:"capacity,omitempty"`

	// Family: If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `json:"family,omitempty"`

	// +kubebuilder:validation:Required
	// Name: The name of the SKU, typically, a letter + Number code, e.g. P3.
	Name *string `json:"name,omitempty"`

	// Size: Size of the particular SKU
	Size *string `json:"size,omitempty"`

	// Tier: The tier or edition of the particular SKU, e.g. Basic, Premium.
	Tier *string `json:"tier,omitempty"`
}

// An ARM Resource SKU.
type Sku_STATUS struct {
	// Capacity: Capacity of the particular SKU.
	Capacity *int `json:"capacity,omitempty"`

	// Family: If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `json:"family,omitempty"`

	// Name: The name of the SKU, typically, a letter + Number code, e.g. P3.
	Name *string `json:"name,omitempty"`

	// Size: Size of the particular SKU
	Size *string `json:"size,omitempty"`

	// Tier: The tier or edition of the particular SKU, e.g. Basic, Premium.
	Tier *string `json:"tier,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Servers_Database{}, &Servers_DatabaseList{})
}
