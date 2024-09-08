// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240302

// Snapshot resource.
type Snapshot_STATUS_ARM struct {
	// ExtendedLocation: The extended location where the snapshot will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocation_STATUS_ARM `json:"extendedLocation,omitempty"`

	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`

	// ManagedBy: Unused. Always Null.
	ManagedBy *string `json:"managedBy,omitempty"`

	// Name: Resource name
	Name *string `json:"name,omitempty"`

	// Properties: Snapshot resource properties.
	Properties *SnapshotProperties_STATUS_ARM `json:"properties,omitempty"`

	// Sku: The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for
	// incremental  snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot
	Sku *SnapshotSku_STATUS_ARM `json:"sku,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`
}

// Snapshot resource properties.
type SnapshotProperties_STATUS_ARM struct {
	// CompletionPercent: Percentage complete for the background copy when a resource is created via the CopyStart operation.
	CompletionPercent *float64 `json:"completionPercent,omitempty"`

	// CopyCompletionError: Indicates the error details if the background copy of a resource created via the CopyStart
	// operation fails.
	CopyCompletionError *CopyCompletionError_STATUS_ARM `json:"copyCompletionError,omitempty"`

	// CreationData: Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData *CreationData_STATUS_ARM `json:"creationData,omitempty"`

	// DataAccessAuthMode: Additional authentication requirements when exporting or uploading to a disk or snapshot.
	DataAccessAuthMode *DataAccessAuthMode_STATUS_ARM `json:"dataAccessAuthMode,omitempty"`

	// DiskAccessId: ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId *string `json:"diskAccessId,omitempty"`

	// DiskSizeBytes: The size of the disk in bytes. This field is read only.
	DiskSizeBytes *int `json:"diskSizeBytes,omitempty"`

	// DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
	// create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
	// allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// DiskState: The state of the snapshot.
	DiskState *DiskState_STATUS_ARM `json:"diskState,omitempty"`

	// Encryption: Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption *Encryption_STATUS_ARM `json:"encryption,omitempty"`

	// EncryptionSettingsCollection: Encryption settings collection used be Azure Disk Encryption, can contain multiple
	// encryption settings per disk or snapshot.
	EncryptionSettingsCollection *EncryptionSettingsCollection_STATUS_ARM `json:"encryptionSettingsCollection,omitempty"`

	// HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *SnapshotProperties_HyperVGeneration_STATUS_ARM `json:"hyperVGeneration,omitempty"`

	// Incremental: Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full
	// snapshots and can be diffed.
	Incremental *bool `json:"incremental,omitempty"`

	// IncrementalSnapshotFamilyId: Incremental snapshots for a disk share an incremental snapshot family id. The Get Page
	// Range Diff API can only be called on incremental snapshots with the same family id.
	IncrementalSnapshotFamilyId *string `json:"incrementalSnapshotFamilyId,omitempty"`

	// NetworkAccessPolicy: Policy for accessing the disk via network.
	NetworkAccessPolicy *NetworkAccessPolicy_STATUS_ARM `json:"networkAccessPolicy,omitempty"`

	// OsType: The Operating System type.
	OsType *SnapshotProperties_OsType_STATUS_ARM `json:"osType,omitempty"`

	// ProvisioningState: The disk provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Policy for controlling export on the disk.
	PublicNetworkAccess *PublicNetworkAccess_STATUS_ARM `json:"publicNetworkAccess,omitempty"`

	// PurchasePlan: Purchase plan information for the image from which the source disk for the snapshot was originally created.
	PurchasePlan *PurchasePlan_STATUS_ARM `json:"purchasePlan,omitempty"`

	// SecurityProfile: Contains the security related information for the resource.
	SecurityProfile *DiskSecurityProfile_STATUS_ARM `json:"securityProfile,omitempty"`

	// SupportedCapabilities: List of supported capabilities for the image from which the source disk from the snapshot was
	// originally created.
	SupportedCapabilities *SupportedCapabilities_STATUS_ARM `json:"supportedCapabilities,omitempty"`

	// SupportsHibernation: Indicates the OS on a snapshot supports hibernation.
	SupportsHibernation *bool `json:"supportsHibernation,omitempty"`

	// TimeCreated: The time when the snapshot was created.
	TimeCreated *string `json:"timeCreated,omitempty"`

	// UniqueId: Unique Guid identifying the resource.
	UniqueId *string `json:"uniqueId,omitempty"`
}

// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for incremental
// snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot
type SnapshotSku_STATUS_ARM struct {
	// Name: The sku name.
	Name *SnapshotSku_Name_STATUS_ARM `json:"name,omitempty"`

	// Tier: The sku tier.
	Tier *string `json:"tier,omitempty"`
}

// Indicates the error details if the background copy of a resource created via the CopyStart operation fails.
type CopyCompletionError_STATUS_ARM struct {
	// ErrorCode: Indicates the error code if the background copy of a resource created via the CopyStart operation fails.
	ErrorCode *CopyCompletionError_ErrorCode_STATUS_ARM `json:"errorCode,omitempty"`

	// ErrorMessage: Indicates the error message if the background copy of a resource created via the CopyStart operation fails.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

type SnapshotProperties_HyperVGeneration_STATUS_ARM string

const (
	SnapshotProperties_HyperVGeneration_STATUS_ARM_V1 = SnapshotProperties_HyperVGeneration_STATUS_ARM("V1")
	SnapshotProperties_HyperVGeneration_STATUS_ARM_V2 = SnapshotProperties_HyperVGeneration_STATUS_ARM("V2")
)

// Mapping from string to SnapshotProperties_HyperVGeneration_STATUS_ARM
var snapshotProperties_HyperVGeneration_STATUS_ARM_Values = map[string]SnapshotProperties_HyperVGeneration_STATUS_ARM{
	"v1": SnapshotProperties_HyperVGeneration_STATUS_ARM_V1,
	"v2": SnapshotProperties_HyperVGeneration_STATUS_ARM_V2,
}

type SnapshotProperties_OsType_STATUS_ARM string

const (
	SnapshotProperties_OsType_STATUS_ARM_Linux   = SnapshotProperties_OsType_STATUS_ARM("Linux")
	SnapshotProperties_OsType_STATUS_ARM_Windows = SnapshotProperties_OsType_STATUS_ARM("Windows")
)

// Mapping from string to SnapshotProperties_OsType_STATUS_ARM
var snapshotProperties_OsType_STATUS_ARM_Values = map[string]SnapshotProperties_OsType_STATUS_ARM{
	"linux":   SnapshotProperties_OsType_STATUS_ARM_Linux,
	"windows": SnapshotProperties_OsType_STATUS_ARM_Windows,
}

type SnapshotSku_Name_STATUS_ARM string

const (
	SnapshotSku_Name_STATUS_ARM_Premium_LRS  = SnapshotSku_Name_STATUS_ARM("Premium_LRS")
	SnapshotSku_Name_STATUS_ARM_Standard_LRS = SnapshotSku_Name_STATUS_ARM("Standard_LRS")
	SnapshotSku_Name_STATUS_ARM_Standard_ZRS = SnapshotSku_Name_STATUS_ARM("Standard_ZRS")
)

// Mapping from string to SnapshotSku_Name_STATUS_ARM
var snapshotSku_Name_STATUS_ARM_Values = map[string]SnapshotSku_Name_STATUS_ARM{
	"premium_lrs":  SnapshotSku_Name_STATUS_ARM_Premium_LRS,
	"standard_lrs": SnapshotSku_Name_STATUS_ARM_Standard_LRS,
	"standard_zrs": SnapshotSku_Name_STATUS_ARM_Standard_ZRS,
}

type CopyCompletionError_ErrorCode_STATUS_ARM string

const CopyCompletionError_ErrorCode_STATUS_ARM_CopySourceNotFound = CopyCompletionError_ErrorCode_STATUS_ARM("CopySourceNotFound")

// Mapping from string to CopyCompletionError_ErrorCode_STATUS_ARM
var copyCompletionError_ErrorCode_STATUS_ARM_Values = map[string]CopyCompletionError_ErrorCode_STATUS_ARM{
	"copysourcenotfound": CopyCompletionError_ErrorCode_STATUS_ARM_CopySourceNotFound,
}