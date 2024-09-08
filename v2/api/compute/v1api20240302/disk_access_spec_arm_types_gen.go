// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240302

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DiskAccess_Spec_ARM struct {
	// ExtendedLocation: The extended location where the disk access will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocation_ARM `json:"extendedLocation,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DiskAccess_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-03-02"
func (access DiskAccess_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (access *DiskAccess_Spec_ARM) GetName() string {
	return access.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/diskAccesses"
func (access *DiskAccess_Spec_ARM) GetType() string {
	return "Microsoft.Compute/diskAccesses"
}

// The complex type of the extended location.
type ExtendedLocation_ARM struct {
	// Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	// Type: The type of the extended location.
	Type *ExtendedLocationType_ARM `json:"type,omitempty"`
}

// The type of extendedLocation.
// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocationType_ARM string

const ExtendedLocationType_ARM_EdgeZone = ExtendedLocationType_ARM("EdgeZone")

// Mapping from string to ExtendedLocationType_ARM
var extendedLocationType_ARM_Values = map[string]ExtendedLocationType_ARM{
	"edgezone": ExtendedLocationType_ARM_EdgeZone,
}