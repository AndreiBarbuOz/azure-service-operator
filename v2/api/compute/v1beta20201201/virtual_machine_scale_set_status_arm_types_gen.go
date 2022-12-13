// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

// Describes a Virtual Machine Scale Set.
type VirtualMachineScaleSet_STATUS_ARM struct {
	// ExtendedLocation: The extended location of the Virtual Machine Scale Set.
	ExtendedLocation *ExtendedLocation_STATUS_ARM `json:"extendedLocation,omitempty"`

	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Identity: The identity of the virtual machine scale set, if configured.
	Identity *VirtualMachineScaleSetIdentity_STATUS_ARM `json:"identity,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`

	// Name: Resource name
	Name *string `json:"name,omitempty"`

	// Plan: Specifies information about the marketplace image used to create the virtual machine. This element is only used
	// for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic
	// use.  In the Azure portal, find the marketplace image that you want to use and then click Want to deploy
	// programmatically, Get Started ->. Enter any required information and then click Save.
	Plan *Plan_STATUS_ARM `json:"plan,omitempty"`

	// Properties: Describes the properties of a Virtual Machine Scale Set.
	Properties *VirtualMachineScaleSetProperties_STATUS_ARM `json:"properties,omitempty"`

	// Sku: The virtual machine scale set sku.
	Sku *Sku_STATUS_ARM `json:"sku,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`

	// Zones: The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set
	Zones []string `json:"zones,omitempty"`
}

// The complex type of the extended location.
type ExtendedLocation_STATUS_ARM struct {
	// Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	// Type: The type of the extended location.
	Type *ExtendedLocationType_STATUS `json:"type,omitempty"`
}

// Specifies information about the marketplace image used to create the virtual machine. This element is only used for
// marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.
// In the Azure portal, find the marketplace image that you want to use and then click Want to deploy programmatically,
// Get Started ->. Enter any required information and then click Save.
type Plan_STATUS_ARM struct {
	// Name: The plan ID.
	Name *string `json:"name,omitempty"`

	// Product: Specifies the product of the image from the marketplace. This is the same value as Offer under the
	// imageReference element.
	Product *string `json:"product,omitempty"`

	// PromotionCode: The promotion code.
	PromotionCode *string `json:"promotionCode,omitempty"`

	// Publisher: The publisher ID.
	Publisher *string `json:"publisher,omitempty"`
}

// Describes a virtual machine scale set sku. NOTE: If the new VM SKU is not supported on the hardware the scale set is
// currently on, you need to deallocate the VMs in the scale set before you modify the SKU name.
type Sku_STATUS_ARM struct {
	// Capacity: Specifies the number of virtual machines in the scale set.
	Capacity *int `json:"capacity,omitempty"`

	// Name: The sku name.
	Name *string `json:"name,omitempty"`

	// Tier: Specifies the tier of virtual machines in a scale set.
	// Possible Values:
	// Standard
	// Basic
	Tier *string `json:"tier,omitempty"`
}

// Identity for the virtual machine scale set.
type VirtualMachineScaleSetIdentity_STATUS_ARM struct {
	// PrincipalId: The principal id of virtual machine scale set identity. This property will only be provided for a system
	// assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant id associated with the virtual machine scale set. This property will only be provided for a system
	// assigned identity.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The type of identity used for the virtual machine scale set. The type 'SystemAssigned, UserAssigned' includes both
	// an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from
	// the virtual machine scale set.
	Type *VirtualMachineScaleSetIdentity_Type_STATUS `json:"type,omitempty"`
}

// Describes the properties of a Virtual Machine Scale Set.
type VirtualMachineScaleSetProperties_STATUS_ARM struct {
	// AdditionalCapabilities: Specifies additional capabilities enabled or disabled on the Virtual Machines in the Virtual
	// Machine Scale Set. For instance: whether the Virtual Machines have the capability to support attaching managed data
	// disks with UltraSSD_LRS storage account type.
	AdditionalCapabilities *AdditionalCapabilities_STATUS_ARM `json:"additionalCapabilities,omitempty"`

	// AutomaticRepairsPolicy: Policy for automatic repairs.
	AutomaticRepairsPolicy *AutomaticRepairsPolicy_STATUS_ARM `json:"automaticRepairsPolicy,omitempty"`

	// DoNotRunExtensionsOnOverprovisionedVMs: When Overprovision is enabled, extensions are launched only on the requested
	// number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra
	// overprovisioned VMs.
	DoNotRunExtensionsOnOverprovisionedVMs *bool `json:"doNotRunExtensionsOnOverprovisionedVMs,omitempty"`

	// HostGroup: Specifies information about the dedicated host group that the virtual machine scale set resides in.
	// Minimum api-version: 2020-06-01.
	HostGroup *SubResource_STATUS_ARM `json:"hostGroup,omitempty"`

	// OrchestrationMode: Specifies the orchestration mode for the virtual machine scale set.
	OrchestrationMode *OrchestrationMode_STATUS `json:"orchestrationMode,omitempty"`

	// Overprovision: Specifies whether the Virtual Machine Scale Set should be overprovisioned.
	Overprovision *bool `json:"overprovision,omitempty"`

	// PlatformFaultDomainCount: Fault Domain count for each placement group.
	PlatformFaultDomainCount *int `json:"platformFaultDomainCount,omitempty"`

	// ProvisioningState: The provisioning state, which only appears in the response.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// ProximityPlacementGroup: Specifies information about the proximity placement group that the virtual machine scale set
	// should be assigned to.
	// Minimum api-version: 2018-04-01.
	ProximityPlacementGroup *SubResource_STATUS_ARM `json:"proximityPlacementGroup,omitempty"`

	// ScaleInPolicy: Specifies the scale-in policy that decides which virtual machines are chosen for removal when a Virtual
	// Machine Scale Set is scaled-in.
	ScaleInPolicy *ScaleInPolicy_STATUS_ARM `json:"scaleInPolicy,omitempty"`

	// SinglePlacementGroup: When true this limits the scale set to a single placement group, of max size 100 virtual machines.
	// NOTE: If singlePlacementGroup is true, it may be modified to false. However, if singlePlacementGroup is false, it may
	// not be modified to true.
	SinglePlacementGroup *bool `json:"singlePlacementGroup,omitempty"`

	// UniqueId: Specifies the ID which uniquely identifies a Virtual Machine Scale Set.
	UniqueId *string `json:"uniqueId,omitempty"`

	// UpgradePolicy: The upgrade policy.
	UpgradePolicy *UpgradePolicy_STATUS_ARM `json:"upgradePolicy,omitempty"`

	// VirtualMachineProfile: The virtual machine profile.
	VirtualMachineProfile *VirtualMachineScaleSetVMProfile_STATUS_ARM `json:"virtualMachineProfile,omitempty"`

	// ZoneBalance: Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.
	ZoneBalance *bool `json:"zoneBalance,omitempty"`
}

// Enables or disables a capability on the virtual machine or virtual machine scale set.
type AdditionalCapabilities_STATUS_ARM struct {
	// UltraSSDEnabled: The flag that enables or disables a capability to have one or more managed data disks with UltraSSD_LRS
	// storage account type on the VM or VMSS. Managed disks with storage account type UltraSSD_LRS can be added to a virtual
	// machine or virtual machine scale set only if this property is enabled.
	UltraSSDEnabled *bool `json:"ultraSSDEnabled,omitempty"`
}

// Specifies the configuration parameters for automatic repairs on the virtual machine scale set.
type AutomaticRepairsPolicy_STATUS_ARM struct {
	// Enabled: Specifies whether automatic repairs should be enabled on the virtual machine scale set. The default value is
	// false.
	Enabled *bool `json:"enabled,omitempty"`

	// GracePeriod: The amount of time for which automatic repairs are suspended due to a state change on VM. The grace time
	// starts after the state change has completed. This helps avoid premature or accidental repairs. The time duration should
	// be specified in ISO 8601 format. The minimum allowed grace period is 30 minutes (PT30M), which is also the default
	// value. The maximum allowed grace period is 90 minutes (PT90M).
	GracePeriod *string `json:"gracePeriod,omitempty"`
}

// The type of extendedLocation.
type ExtendedLocationType_STATUS string

const ExtendedLocationType_STATUS_EdgeZone = ExtendedLocationType_STATUS("EdgeZone")

// Describes a scale-in policy for a virtual machine scale set.
type ScaleInPolicy_STATUS_ARM struct {
	// Rules: The rules to be followed when scaling-in a virtual machine scale set.
	// Possible values are:
	// Default When a virtual machine scale set is scaled in, the scale set will first be balanced across zones if it is a
	// zonal scale set. Then, it will be balanced across Fault Domains as far as possible. Within each Fault Domain, the
	// virtual machines chosen for removal will be the newest ones that are not protected from scale-in.
	// OldestVM When a virtual machine scale set is being scaled-in, the oldest virtual machines that are not protected from
	// scale-in will be chosen for removal. For zonal virtual machine scale sets, the scale set will first be balanced across
	// zones. Within each zone, the oldest virtual machines that are not protected will be chosen for removal.
	// NewestVM When a virtual machine scale set is being scaled-in, the newest virtual machines that are not protected from
	// scale-in will be chosen for removal. For zonal virtual machine scale sets, the scale set will first be balanced across
	// zones. Within each zone, the newest virtual machines that are not protected will be chosen for removal.
	Rules []ScaleInPolicy_Rules_STATUS `json:"rules,omitempty"`
}

type SubResource_STATUS_ARM struct {
	// Id: Resource Id
	Id *string `json:"id,omitempty"`
}

// Describes an upgrade policy - automatic, manual, or rolling.
type UpgradePolicy_STATUS_ARM struct {
	// AutomaticOSUpgradePolicy: Configuration parameters used for performing automatic OS Upgrade.
	AutomaticOSUpgradePolicy *AutomaticOSUpgradePolicy_STATUS_ARM `json:"automaticOSUpgradePolicy,omitempty"`

	// Mode: Specifies the mode of an upgrade to virtual machines in the scale set.
	// Possible values are:
	// Manual - You  control the application of updates to virtual machines in the scale set. You do this by using the
	// manualUpgrade action.
	// Automatic - All virtual machines in the scale set are  automatically updated at the same time.
	Mode *UpgradePolicy_Mode_STATUS `json:"mode,omitempty"`

	// RollingUpgradePolicy: The configuration parameters used while performing a rolling upgrade.
	RollingUpgradePolicy *RollingUpgradePolicy_STATUS_ARM `json:"rollingUpgradePolicy,omitempty"`
}

type VirtualMachineScaleSetIdentity_Type_STATUS string

const (
	VirtualMachineScaleSetIdentity_Type_STATUS_None                       = VirtualMachineScaleSetIdentity_Type_STATUS("None")
	VirtualMachineScaleSetIdentity_Type_STATUS_SystemAssigned             = VirtualMachineScaleSetIdentity_Type_STATUS("SystemAssigned")
	VirtualMachineScaleSetIdentity_Type_STATUS_SystemAssignedUserAssigned = VirtualMachineScaleSetIdentity_Type_STATUS("SystemAssigned, UserAssigned")
	VirtualMachineScaleSetIdentity_Type_STATUS_UserAssigned               = VirtualMachineScaleSetIdentity_Type_STATUS("UserAssigned")
)

// Describes a virtual machine scale set virtual machine profile.
type VirtualMachineScaleSetVMProfile_STATUS_ARM struct {
	// BillingProfile: Specifies the billing related details of a Azure Spot VMSS.
	// Minimum api-version: 2019-03-01.
	BillingProfile *BillingProfile_STATUS_ARM `json:"billingProfile,omitempty"`

	// DiagnosticsProfile: Specifies the boot diagnostic settings state.
	// Minimum api-version: 2015-06-15.
	DiagnosticsProfile *DiagnosticsProfile_STATUS_ARM `json:"diagnosticsProfile,omitempty"`

	// EvictionPolicy: Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set.
	// For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01.
	// For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is
	// 2017-10-30-preview.
	EvictionPolicy *EvictionPolicy_STATUS `json:"evictionPolicy,omitempty"`

	// ExtensionProfile: Specifies a collection of settings for extensions installed on virtual machines in the scale set.
	ExtensionProfile *VirtualMachineScaleSetExtensionProfile_STATUS_ARM `json:"extensionProfile,omitempty"`

	// LicenseType: Specifies that the image or disk that is being used was licensed on-premises.
	// Possible values for Windows Server operating system are:
	// Windows_Client
	// Windows_Server
	// Possible values for Linux Server operating system are:
	// RHEL_BYOS (for RHEL)
	// SLES_BYOS (for SUSE)
	// For more information, see [Azure Hybrid Use Benefit for Windows
	// Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing)
	// [Azure Hybrid Use Benefit for Linux
	// Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux)
	// Minimum api-version: 2015-06-15
	LicenseType *string `json:"licenseType,omitempty"`

	// NetworkProfile: Specifies properties of the network interfaces of the virtual machines in the scale set.
	NetworkProfile *VirtualMachineScaleSetNetworkProfile_STATUS_ARM `json:"networkProfile,omitempty"`

	// OsProfile: Specifies the operating system settings for the virtual machines in the scale set.
	OsProfile *VirtualMachineScaleSetOSProfile_STATUS_ARM `json:"osProfile,omitempty"`

	// Priority: Specifies the priority for the virtual machines in the scale set.
	// Minimum api-version: 2017-10-30-preview
	Priority *Priority_STATUS `json:"priority,omitempty"`

	// ScheduledEventsProfile: Specifies Scheduled Event related configurations.
	ScheduledEventsProfile *ScheduledEventsProfile_STATUS_ARM `json:"scheduledEventsProfile,omitempty"`

	// SecurityProfile: Specifies the Security related profile settings for the virtual machines in the scale set.
	SecurityProfile *SecurityProfile_STATUS_ARM `json:"securityProfile,omitempty"`

	// StorageProfile: Specifies the storage settings for the virtual machine disks.
	StorageProfile *VirtualMachineScaleSetStorageProfile_STATUS_ARM `json:"storageProfile,omitempty"`
}

// The configuration parameters used for performing automatic OS upgrade.
type AutomaticOSUpgradePolicy_STATUS_ARM struct {
	// DisableAutomaticRollback: Whether OS image rollback feature should be disabled. Default value is false.
	DisableAutomaticRollback *bool `json:"disableAutomaticRollback,omitempty"`

	// EnableAutomaticOSUpgrade: Indicates whether OS upgrades should automatically be applied to scale set instances in a
	// rolling fashion when a newer version of the OS image becomes available. Default value is false.
	// If this is set to true for Windows based scale sets,
	// [enableAutomaticUpdates](https://docs.microsoft.com/dotnet/api/microsoft.azure.management.compute.models.windowsconfiguration.enableautomaticupdates?view=azure-dotnet)
	// is automatically set to false and cannot be set to true.
	EnableAutomaticOSUpgrade *bool `json:"enableAutomaticOSUpgrade,omitempty"`
}

// The configuration parameters used while performing a rolling upgrade.
type RollingUpgradePolicy_STATUS_ARM struct {
	// EnableCrossZoneUpgrade: Allow VMSS to ignore AZ boundaries when constructing upgrade batches. Take into consideration
	// the Update Domain and maxBatchInstancePercent to determine the batch size.
	EnableCrossZoneUpgrade *bool `json:"enableCrossZoneUpgrade,omitempty"`

	// MaxBatchInstancePercent: The maximum percent of total virtual machine instances that will be upgraded simultaneously by
	// the rolling upgrade in one batch. As this is a maximum, unhealthy instances in previous or future batches can cause the
	// percentage of instances in a batch to decrease to ensure higher reliability. The default value for this parameter is 20%.
	MaxBatchInstancePercent *int `json:"maxBatchInstancePercent,omitempty"`

	// MaxUnhealthyInstancePercent: The maximum percentage of the total virtual machine instances in the scale set that can be
	// simultaneously unhealthy, either as a result of being upgraded, or by being found in an unhealthy state by the virtual
	// machine health checks before the rolling upgrade aborts. This constraint will be checked prior to starting any batch.
	// The default value for this parameter is 20%.
	MaxUnhealthyInstancePercent *int `json:"maxUnhealthyInstancePercent,omitempty"`

	// MaxUnhealthyUpgradedInstancePercent: The maximum percentage of upgraded virtual machine instances that can be found to
	// be in an unhealthy state. This check will happen after each batch is upgraded. If this percentage is ever exceeded, the
	// rolling update aborts. The default value for this parameter is 20%.
	MaxUnhealthyUpgradedInstancePercent *int `json:"maxUnhealthyUpgradedInstancePercent,omitempty"`

	// PauseTimeBetweenBatches: The wait time between completing the update for all virtual machines in one batch and starting
	// the next batch. The time duration should be specified in ISO 8601 format. The default value is 0 seconds (PT0S).
	PauseTimeBetweenBatches *string `json:"pauseTimeBetweenBatches,omitempty"`

	// PrioritizeUnhealthyInstances: Upgrade all unhealthy instances in a scale set before any healthy instances.
	PrioritizeUnhealthyInstances *bool `json:"prioritizeUnhealthyInstances,omitempty"`
}

type ScheduledEventsProfile_STATUS_ARM struct {
	// TerminateNotificationProfile: Specifies Terminate Scheduled Event related configurations.
	TerminateNotificationProfile *TerminateNotificationProfile_STATUS_ARM `json:"terminateNotificationProfile,omitempty"`
}

// Describes a virtual machine scale set extension profile.
type VirtualMachineScaleSetExtensionProfile_STATUS_ARM struct {
	// Extensions: The virtual machine scale set child extension resources.
	Extensions []VirtualMachineScaleSetExtension_STATUS_ARM `json:"extensions,omitempty"`

	// ExtensionsTimeBudget: Specifies the time alloted for all extensions to start. The time duration should be between 15
	// minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes
	// (PT1H30M).
	// Minimum api-version: 2020-06-01
	ExtensionsTimeBudget *string `json:"extensionsTimeBudget,omitempty"`
}

// Describes a virtual machine scale set network profile.
type VirtualMachineScaleSetNetworkProfile_STATUS_ARM struct {
	// HealthProbe: A reference to a load balancer probe used to determine the health of an instance in the virtual machine
	// scale set. The reference will be in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes/{probeName}'.
	HealthProbe *ApiEntityReference_STATUS_ARM `json:"healthProbe,omitempty"`

	// NetworkInterfaceConfigurations: The list of network configurations.
	NetworkInterfaceConfigurations []VirtualMachineScaleSetNetworkConfiguration_STATUS_ARM `json:"networkInterfaceConfigurations,omitempty"`
}

// Describes a virtual machine scale set OS profile.
type VirtualMachineScaleSetOSProfile_STATUS_ARM struct {
	// AdminUsername: Specifies the name of the administrator account.
	// Windows-only restriction: Cannot end in "."
	// Disallowed values: "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123",
	// "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server",
	// "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5".
	// Minimum-length (Linux): 1  character
	// Max-length (Linux): 64 characters
	// Max-length (Windows): 20 characters
	// <li> For root access to the Linux VM, see [Using root privileges on Linux virtual machines in
	// Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
	// <li> For a list of built-in system users on Linux that should not be used in this field, see [Selecting User Names for
	// Linux on
	// Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
	AdminUsername *string `json:"adminUsername,omitempty"`

	// ComputerNamePrefix: Specifies the computer name prefix for all of the virtual machines in the scale set. Computer name
	// prefixes must be 1 to 15 characters long.
	ComputerNamePrefix *string `json:"computerNamePrefix,omitempty"`

	// CustomData: Specifies a base-64 encoded string of custom data. The base-64 encoded string is decoded to a binary array
	// that is saved as a file on the Virtual Machine. The maximum length of the binary array is 65535 bytes.
	// For using cloud-init for your VM, see [Using cloud-init to customize a Linux VM during
	// creation](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-cloud-init?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
	CustomData *string `json:"customData,omitempty"`

	// LinuxConfiguration: Specifies the Linux operating system settings on the virtual machine.
	// For a list of supported Linux distributions, see [Linux on Azure-Endorsed
	// Distributions](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-endorsed-distros?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
	// For running non-endorsed distributions, see [Information for Non-Endorsed
	// Distributions](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-create-upload-generic?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json).
	LinuxConfiguration *LinuxConfiguration_STATUS_ARM `json:"linuxConfiguration,omitempty"`

	// Secrets: Specifies set of certificates that should be installed onto the virtual machines in the scale set.
	Secrets []VaultSecretGroup_STATUS_ARM `json:"secrets,omitempty"`

	// WindowsConfiguration: Specifies Windows operating system settings on the virtual machine.
	WindowsConfiguration *WindowsConfiguration_STATUS_ARM `json:"windowsConfiguration,omitempty"`
}

// Describes a virtual machine scale set storage profile.
type VirtualMachineScaleSetStorageProfile_STATUS_ARM struct {
	// DataDisks: Specifies the parameters that are used to add data disks to the virtual machines in the scale set.
	// For more information about disks, see [About disks and VHDs for Azure virtual
	// machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
	DataDisks []VirtualMachineScaleSetDataDisk_STATUS_ARM `json:"dataDisks,omitempty"`

	// ImageReference: Specifies information about the image to use. You can specify information about platform images,
	// marketplace images, or virtual machine images. This element is required when you want to use a platform image,
	// marketplace image, or virtual machine image, but is not used in other creation operations.
	ImageReference *ImageReference_STATUS_ARM `json:"imageReference,omitempty"`

	// OsDisk: Specifies information about the operating system disk used by the virtual machines in the scale set.
	// For more information about disks, see [About disks and VHDs for Azure virtual
	// machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
	OsDisk *VirtualMachineScaleSetOSDisk_STATUS_ARM `json:"osDisk,omitempty"`
}

// The API entity reference.
type ApiEntityReference_STATUS_ARM struct {
	// Id: The ARM resource id in the form of /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/...
	Id *string `json:"id,omitempty"`
}

type TerminateNotificationProfile_STATUS_ARM struct {
	// Enable: Specifies whether the Terminate Scheduled event is enabled or disabled.
	Enable *bool `json:"enable,omitempty"`

	// NotBeforeTimeout: Configurable length of time a Virtual Machine being deleted will have to potentially approve the
	// Terminate Scheduled Event before the event is auto approved (timed out). The configuration must be specified in ISO 8601
	// format, the default value is 5 minutes (PT5M)
	NotBeforeTimeout *string `json:"notBeforeTimeout,omitempty"`
}

// Describes a virtual machine scale set data disk.
type VirtualMachineScaleSetDataDisk_STATUS_ARM struct {
	// Caching: Specifies the caching requirements.
	// Possible values are:
	// None
	// ReadOnly
	// ReadWrite
	// Default: None for Standard storage. ReadOnly for Premium storage
	Caching *Caching_STATUS `json:"caching,omitempty"`

	// CreateOption: The create option.
	CreateOption *CreateOption_STATUS `json:"createOption,omitempty"`

	// DiskIOPSReadWrite: Specifies the Read-Write IOPS for the managed disk. Should be used only when StorageAccountType is
	// UltraSSD_LRS. If not specified, a default value would be assigned based on diskSizeGB.
	DiskIOPSReadWrite *int `json:"diskIOPSReadWrite,omitempty"`

	// DiskMBpsReadWrite: Specifies the bandwidth in MB per second for the managed disk. Should be used only when
	// StorageAccountType is UltraSSD_LRS. If not specified, a default value would be assigned based on diskSizeGB.
	DiskMBpsReadWrite *int `json:"diskMBpsReadWrite,omitempty"`

	// DiskSizeGB: Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the
	// disk in a virtual machine image.
	// This value cannot be larger than 1023 GB
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
	// therefore must be unique for each data disk attached to a VM.
	Lun *int `json:"lun,omitempty"`

	// ManagedDisk: The managed disk parameters.
	ManagedDisk *VirtualMachineScaleSetManagedDiskParameters_STATUS_ARM `json:"managedDisk,omitempty"`

	// Name: The disk name.
	Name *string `json:"name,omitempty"`

	// WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty"`
}

// Describes a Virtual Machine Scale Set Extension.
type VirtualMachineScaleSetExtension_STATUS_ARM struct {
	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Name: The name of the extension.
	Name *string `json:"name,omitempty"`

	// Properties: Describes the properties of a Virtual Machine Scale Set Extension.
	Properties *VirtualMachineScaleSetExtensionProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`
}

// Describes a virtual machine scale set network profile's network configurations.
type VirtualMachineScaleSetNetworkConfiguration_STATUS_ARM struct {
	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Name: The network configuration name.
	Name *string `json:"name,omitempty"`

	// Properties: Describes a virtual machine scale set network profile's IP configuration.
	Properties *VirtualMachineScaleSetNetworkConfigurationProperties_STATUS_ARM `json:"properties,omitempty"`
}

// Describes a virtual machine scale set operating system disk.
type VirtualMachineScaleSetOSDisk_STATUS_ARM struct {
	// Caching: Specifies the caching requirements.
	// Possible values are:
	// None
	// ReadOnly
	// ReadWrite
	// Default: None for Standard storage. ReadOnly for Premium storage
	Caching *Caching_STATUS `json:"caching,omitempty"`

	// CreateOption: Specifies how the virtual machines in the scale set should be created.
	// The only allowed value is: FromImage \u2013 This value is used when you are using an image to create the virtual
	// machine. If you are using a platform image, you also use the imageReference element described above. If you are using a
	// marketplace image, you  also use the plan element previously described.
	CreateOption *CreateOption_STATUS `json:"createOption,omitempty"`

	// DiffDiskSettings: Specifies the ephemeral disk Settings for the operating system disk used by the virtual machine scale
	// set.
	DiffDiskSettings *DiffDiskSettings_STATUS_ARM `json:"diffDiskSettings,omitempty"`

	// DiskSizeGB: Specifies the size of the operating system disk in gigabytes. This element can be used to overwrite the size
	// of the disk in a virtual machine image.
	// This value cannot be larger than 1023 GB
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// Image: Specifies information about the unmanaged user image to base the scale set on.
	Image *VirtualHardDisk_STATUS_ARM `json:"image,omitempty"`

	// ManagedDisk: The managed disk parameters.
	ManagedDisk *VirtualMachineScaleSetManagedDiskParameters_STATUS_ARM `json:"managedDisk,omitempty"`

	// Name: The disk name.
	Name *string `json:"name,omitempty"`

	// OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from
	// user-image or a specialized VHD.
	// Possible values are:
	// Windows
	// Linux
	OsType *VirtualMachineScaleSetOSDisk_OsType_STATUS `json:"osType,omitempty"`

	// VhdContainers: Specifies the container urls that are used to store operating system disks for the scale set.
	VhdContainers []string `json:"vhdContainers,omitempty"`

	// WriteAcceleratorEnabled: Specifies whether writeAccelerator should be enabled or disabled on the disk.
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty"`
}

// Describes the properties of a Virtual Machine Scale Set Extension.
type VirtualMachineScaleSetExtensionProperties_STATUS_ARM struct {
	// AutoUpgradeMinorVersion: Indicates whether the extension should use a newer minor version if one is available at
	// deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this
	// property set to true.
	AutoUpgradeMinorVersion *bool `json:"autoUpgradeMinorVersion,omitempty"`

	// EnableAutomaticUpgrade: Indicates whether the extension should be automatically upgraded by the platform if there is a
	// newer version of the extension available.
	EnableAutomaticUpgrade *bool `json:"enableAutomaticUpgrade,omitempty"`

	// ForceUpdateTag: If a value is provided and is different from the previous value, the extension handler will be forced to
	// update even if the extension configuration has not changed.
	ForceUpdateTag *string `json:"forceUpdateTag,omitempty"`

	// ProtectedSettings: The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected
	// settings at all.
	ProtectedSettings map[string]v1.JSON `json:"protectedSettings,omitempty"`

	// ProvisionAfterExtensions: Collection of extension names after which this extension needs to be provisioned.
	ProvisionAfterExtensions []string `json:"provisionAfterExtensions,omitempty"`

	// ProvisioningState: The provisioning state, which only appears in the response.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// Publisher: The name of the extension handler publisher.
	Publisher *string `json:"publisher,omitempty"`

	// Settings: Json formatted public settings for the extension.
	Settings map[string]v1.JSON `json:"settings,omitempty"`

	// Type: Specifies the type of the extension; an example is "CustomScriptExtension".
	Type *string `json:"type,omitempty"`

	// TypeHandlerVersion: Specifies the version of the script handler.
	TypeHandlerVersion *string `json:"typeHandlerVersion,omitempty"`
}

// Describes the parameters of a ScaleSet managed disk.
type VirtualMachineScaleSetManagedDiskParameters_STATUS_ARM struct {
	// DiskEncryptionSet: Specifies the customer managed disk encryption set resource id for the managed disk.
	DiskEncryptionSet *SubResource_STATUS_ARM `json:"diskEncryptionSet,omitempty"`

	// StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
	// data disks, it cannot be used with OS Disk.
	StorageAccountType *StorageAccountType_STATUS `json:"storageAccountType,omitempty"`
}

// Describes a virtual machine scale set network profile's IP configuration.
type VirtualMachineScaleSetNetworkConfigurationProperties_STATUS_ARM struct {
	// DnsSettings: The dns settings to be applied on the network interfaces.
	DnsSettings *VirtualMachineScaleSetNetworkConfigurationDnsSettings_STATUS_ARM `json:"dnsSettings,omitempty"`

	// EnableAcceleratedNetworking: Specifies whether the network interface is accelerated networking-enabled.
	EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty"`

	// EnableFpga: Specifies whether the network interface is FPGA networking-enabled.
	EnableFpga *bool `json:"enableFpga,omitempty"`

	// EnableIPForwarding: Whether IP forwarding enabled on this NIC.
	EnableIPForwarding *bool `json:"enableIPForwarding,omitempty"`

	// IpConfigurations: Specifies the IP configurations of the network interface.
	IpConfigurations []VirtualMachineScaleSetIPConfiguration_STATUS_ARM `json:"ipConfigurations,omitempty"`

	// NetworkSecurityGroup: The network security group.
	NetworkSecurityGroup *SubResource_STATUS_ARM `json:"networkSecurityGroup,omitempty"`

	// Primary: Specifies the primary network interface in case the virtual machine has more than 1 network interface.
	Primary *bool `json:"primary,omitempty"`
}

// Describes a virtual machine scale set network profile's IP configuration.
type VirtualMachineScaleSetIPConfiguration_STATUS_ARM struct {
	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Name: The IP configuration name.
	Name *string `json:"name,omitempty"`

	// Properties: Describes a virtual machine scale set network profile's IP configuration properties.
	Properties *VirtualMachineScaleSetIPConfigurationProperties_STATUS_ARM `json:"properties,omitempty"`
}

// Describes a virtual machines scale sets network configuration's DNS settings.
type VirtualMachineScaleSetNetworkConfigurationDnsSettings_STATUS_ARM struct {
	// DnsServers: List of DNS servers IP addresses
	DnsServers []string `json:"dnsServers,omitempty"`
}

// Describes a virtual machine scale set network profile's IP configuration properties.
type VirtualMachineScaleSetIPConfigurationProperties_STATUS_ARM struct {
	// ApplicationGatewayBackendAddressPools: Specifies an array of references to backend address pools of application
	// gateways. A scale set can reference backend address pools of multiple application gateways. Multiple scale sets cannot
	// use the same application gateway.
	ApplicationGatewayBackendAddressPools []SubResource_STATUS_ARM `json:"applicationGatewayBackendAddressPools,omitempty"`

	// ApplicationSecurityGroups: Specifies an array of references to application security group.
	ApplicationSecurityGroups []SubResource_STATUS_ARM `json:"applicationSecurityGroups,omitempty"`

	// LoadBalancerBackendAddressPools: Specifies an array of references to backend address pools of load balancers. A scale
	// set can reference backend address pools of one public and one internal load balancer. Multiple scale sets cannot use the
	// same basic sku load balancer.
	LoadBalancerBackendAddressPools []SubResource_STATUS_ARM `json:"loadBalancerBackendAddressPools,omitempty"`

	// LoadBalancerInboundNatPools: Specifies an array of references to inbound Nat pools of the load balancers. A scale set
	// can reference inbound nat pools of one public and one internal load balancer. Multiple scale sets cannot use the same
	// basic sku load balancer.
	LoadBalancerInboundNatPools []SubResource_STATUS_ARM `json:"loadBalancerInboundNatPools,omitempty"`

	// Primary: Specifies the primary network interface in case the virtual machine has more than 1 network interface.
	Primary *bool `json:"primary,omitempty"`

	// PrivateIPAddressVersion: Available from Api-Version 2017-03-30 onwards, it represents whether the specific
	// ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.  Possible values are: 'IPv4' and 'IPv6'.
	PrivateIPAddressVersion *VirtualMachineScaleSetIPConfigurationProperties_PrivateIPAddressVersion_STATUS `json:"privateIPAddressVersion,omitempty"`

	// PublicIPAddressConfiguration: The publicIPAddressConfiguration.
	PublicIPAddressConfiguration *VirtualMachineScaleSetPublicIPAddressConfiguration_STATUS_ARM `json:"publicIPAddressConfiguration,omitempty"`

	// Subnet: Specifies the identifier of the subnet.
	Subnet *ApiEntityReference_STATUS_ARM `json:"subnet,omitempty"`
}

// Describes a virtual machines scale set IP Configuration's PublicIPAddress configuration
type VirtualMachineScaleSetPublicIPAddressConfiguration_STATUS_ARM struct {
	// Name: The publicIP address configuration name.
	Name *string `json:"name,omitempty"`

	// Properties: Describes a virtual machines scale set IP Configuration's PublicIPAddress configuration
	Properties *VirtualMachineScaleSetPublicIPAddressConfigurationProperties_STATUS_ARM `json:"properties,omitempty"`
}

// Describes a virtual machines scale set IP Configuration's PublicIPAddress configuration
type VirtualMachineScaleSetPublicIPAddressConfigurationProperties_STATUS_ARM struct {
	// DnsSettings: The dns settings to be applied on the publicIP addresses .
	DnsSettings *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_STATUS_ARM `json:"dnsSettings,omitempty"`

	// IdleTimeoutInMinutes: The idle timeout of the public IP address.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	// IpTags: The list of IP tags associated with the public IP address.
	IpTags []VirtualMachineScaleSetIpTag_STATUS_ARM `json:"ipTags,omitempty"`

	// PublicIPAddressVersion: Available from Api-Version 2019-07-01 onwards, it represents whether the specific
	// ipconfiguration is IPv4 or IPv6. Default is taken as IPv4. Possible values are: 'IPv4' and 'IPv6'.
	PublicIPAddressVersion *VirtualMachineScaleSetPublicIPAddressConfigurationProperties_PublicIPAddressVersion_STATUS `json:"publicIPAddressVersion,omitempty"`

	// PublicIPPrefix: The PublicIPPrefix from which to allocate publicIP addresses.
	PublicIPPrefix *SubResource_STATUS_ARM `json:"publicIPPrefix,omitempty"`
}

// Contains the IP tag associated with the public IP address.
type VirtualMachineScaleSetIpTag_STATUS_ARM struct {
	// IpTagType: IP tag type. Example: FirstPartyUsage.
	IpTagType *string `json:"ipTagType,omitempty"`

	// Tag: IP tag associated with the public IP. Example: SQL, Storage etc.
	Tag *string `json:"tag,omitempty"`
}

// Describes a virtual machines scale sets network configuration's DNS settings.
type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_STATUS_ARM struct {
	// DomainNameLabel: The Domain name label.The concatenation of the domain name label and vm index will be the domain name
	// labels of the PublicIPAddress resources that will be created
	DomainNameLabel *string `json:"domainNameLabel,omitempty"`
}
