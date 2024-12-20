// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

type VirtualMachinesExtension_STATUS struct {
	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`

	// Name: Resource name
	Name *string `json:"name,omitempty"`

	// Properties: Describes the properties of a Virtual Machine Extension.
	Properties *VirtualMachineExtensionProperties_STATUS `json:"properties,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`
}

// Describes the properties of a Virtual Machine Extension.
type VirtualMachineExtensionProperties_STATUS struct {
	// AutoUpgradeMinorVersion: Indicates whether the extension should use a newer minor version if one is available at
	// deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this
	// property set to true.
	AutoUpgradeMinorVersion *bool `json:"autoUpgradeMinorVersion,omitempty"`

	// EnableAutomaticUpgrade: Indicates whether the extension should be automatically upgraded by the platform if there is a
	// newer version of the extension available.
	EnableAutomaticUpgrade *bool `json:"enableAutomaticUpgrade,omitempty"`

	// ForceUpdateTag: How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag *string `json:"forceUpdateTag,omitempty"`

	// InstanceView: The virtual machine extension instance view.
	InstanceView *VirtualMachineExtensionInstanceView_STATUS `json:"instanceView,omitempty"`

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

// The instance view of a virtual machine extension.
type VirtualMachineExtensionInstanceView_STATUS struct {
	// Name: The virtual machine extension name.
	Name *string `json:"name,omitempty"`

	// Statuses: The resource status information.
	Statuses []InstanceViewStatus_STATUS `json:"statuses,omitempty"`

	// Substatuses: The resource status information.
	Substatuses []InstanceViewStatus_STATUS `json:"substatuses,omitempty"`

	// Type: Specifies the type of the extension; an example is "CustomScriptExtension".
	Type *string `json:"type,omitempty"`

	// TypeHandlerVersion: Specifies the version of the script handler.
	TypeHandlerVersion *string `json:"typeHandlerVersion,omitempty"`
}
