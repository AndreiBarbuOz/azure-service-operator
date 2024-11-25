// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type PrivateEndpointsPrivateDnsZoneGroup_Spec struct {
	// Name: Name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name string `json:"name,omitempty"`

	// Properties: Properties of the private dns zone group.
	Properties *PrivateDnsZoneGroupPropertiesFormat `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &PrivateEndpointsPrivateDnsZoneGroup_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-03-01"
func (group PrivateEndpointsPrivateDnsZoneGroup_Spec) GetAPIVersion() string {
	return "2024-03-01"
}

// GetName returns the Name of the resource
func (group *PrivateEndpointsPrivateDnsZoneGroup_Spec) GetName() string {
	return group.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/privateEndpoints/privateDnsZoneGroups"
func (group *PrivateEndpointsPrivateDnsZoneGroup_Spec) GetType() string {
	return "Microsoft.Network/privateEndpoints/privateDnsZoneGroups"
}

// Properties of the private dns zone group.
type PrivateDnsZoneGroupPropertiesFormat struct {
	// PrivateDnsZoneConfigs: A collection of private dns zone configurations of the private dns zone group.
	PrivateDnsZoneConfigs []PrivateDnsZoneConfig `json:"privateDnsZoneConfigs,omitempty"`
}

// PrivateDnsZoneConfig resource.
type PrivateDnsZoneConfig struct {
	// Name: Name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the private dns zone configuration.
	Properties *PrivateDnsZonePropertiesFormat `json:"properties,omitempty"`
}

// Properties of the private dns zone configuration resource.
type PrivateDnsZonePropertiesFormat struct {
	PrivateDnsZoneId *string `json:"privateDnsZoneId,omitempty"`
}