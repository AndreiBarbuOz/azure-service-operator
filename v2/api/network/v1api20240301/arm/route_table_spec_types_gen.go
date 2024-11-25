// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RouteTable_Spec struct {
	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the route table.
	Properties *RouteTablePropertiesFormat `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RouteTable_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-03-01"
func (table RouteTable_Spec) GetAPIVersion() string {
	return "2024-03-01"
}

// GetName returns the Name of the resource
func (table *RouteTable_Spec) GetName() string {
	return table.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/routeTables"
func (table *RouteTable_Spec) GetType() string {
	return "Microsoft.Network/routeTables"
}

// Route Table resource.
type RouteTablePropertiesFormat struct {
	// DisableBgpRoutePropagation: Whether to disable the routes learned by BGP on that route table. True means disable.
	DisableBgpRoutePropagation *bool `json:"disableBgpRoutePropagation,omitempty"`

	// Routes: Collection of routes contained within a route table.
	Routes []Route `json:"routes,omitempty"`
}

// Route resource.
type Route struct {
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the route.
	Properties *RoutePropertiesFormat `json:"properties,omitempty"`

	// Type: The type of the resource.
	Type *string `json:"type,omitempty"`
}