// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Profile_Spec_ARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: Name of the Azure Front Door Standard or Azure Front Door Premium or CDN profile which is unique within the
	// resource group.
	Name string `json:"name,omitempty"`

	// Properties: The JSON object that contains the properties required to create a profile.
	Properties *ProfileProperties_ARM `json:"properties,omitempty"`

	// Sku: Standard_Verizon = The SKU name for a Standard Verizon CDN profile.
	// Premium_Verizon = The SKU name for a Premium Verizon CDN profile.
	// Custom_Verizon = The SKU name for a Custom Verizon CDN profile.
	// Standard_Akamai = The SKU name for an Akamai CDN profile.
	// Standard_ChinaCdn = The SKU name for a China CDN profile for VOD, Web and download scenarios using GB based billing
	// model.
	// Standard_Microsoft = The SKU name for a Standard Microsoft CDN profile.
	// Standard_AzureFrontDoor =  The SKU name for an Azure Front Door Standard profile.
	// Premium_AzureFrontDoor = The SKU name for an Azure Front Door Premium profile.
	// Standard_955BandWidth_ChinaCdn = The SKU name for a China CDN profile for VOD, Web and download scenarios using 95-5
	// peak bandwidth billing model.
	// Standard_AvgBandWidth_ChinaCdn = The SKU name for a China CDN profile for VOD, Web and download scenarios using monthly
	// average peak bandwidth billing model.
	// StandardPlus_ChinaCdn = The SKU name for a China CDN profile for live-streaming using GB based billing model.
	// StandardPlus_955BandWidth_ChinaCdn = The SKU name for a China CDN live-streaming profile using 95-5 peak bandwidth
	// billing model.
	// StandardPlus_AvgBandWidth_ChinaCdn = The SKU name for a China CDN live-streaming profile using monthly average peak
	// bandwidth billing model.
	Sku *Sku_ARM `json:"sku,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Profile_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (profile Profile_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (profile *Profile_Spec_ARM) GetName() string {
	return profile.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles"
func (profile *Profile_Spec_ARM) GetType() string {
	return "Microsoft.Cdn/profiles"
}

// Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.Cdn.json#/definitions/ProfileProperties
type ProfileProperties_ARM struct {
	// OriginResponseTimeoutSeconds: Send and receive timeout on forwarding request to the origin. When timeout is reached, the
	// request fails and returns.
	OriginResponseTimeoutSeconds *int `json:"originResponseTimeoutSeconds,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.Cdn.json#/definitions/Sku
type Sku_ARM struct {
	// Name: Name of the pricing tier.
	Name *Sku_Name `json:"name,omitempty"`
}

// +kubebuilder:validation:Enum={"Custom_Verizon","Premium_AzureFrontDoor","Premium_Verizon","StandardPlus_955BandWidth_ChinaCdn","StandardPlus_AvgBandWidth_ChinaCdn","StandardPlus_ChinaCdn","Standard_955BandWidth_ChinaCdn","Standard_Akamai","Standard_AvgBandWidth_ChinaCdn","Standard_AzureFrontDoor","Standard_ChinaCdn","Standard_Microsoft","Standard_Verizon"}
type Sku_Name string

const (
	Sku_Name_Custom_Verizon                     = Sku_Name("Custom_Verizon")
	Sku_Name_Premium_AzureFrontDoor             = Sku_Name("Premium_AzureFrontDoor")
	Sku_Name_Premium_Verizon                    = Sku_Name("Premium_Verizon")
	Sku_Name_StandardPlus_955BandWidth_ChinaCdn = Sku_Name("StandardPlus_955BandWidth_ChinaCdn")
	Sku_Name_StandardPlus_AvgBandWidth_ChinaCdn = Sku_Name("StandardPlus_AvgBandWidth_ChinaCdn")
	Sku_Name_StandardPlus_ChinaCdn              = Sku_Name("StandardPlus_ChinaCdn")
	Sku_Name_Standard_955BandWidth_ChinaCdn     = Sku_Name("Standard_955BandWidth_ChinaCdn")
	Sku_Name_Standard_Akamai                    = Sku_Name("Standard_Akamai")
	Sku_Name_Standard_AvgBandWidth_ChinaCdn     = Sku_Name("Standard_AvgBandWidth_ChinaCdn")
	Sku_Name_Standard_AzureFrontDoor            = Sku_Name("Standard_AzureFrontDoor")
	Sku_Name_Standard_ChinaCdn                  = Sku_Name("Standard_ChinaCdn")
	Sku_Name_Standard_Microsoft                 = Sku_Name("Standard_Microsoft")
	Sku_Name_Standard_Verizon                   = Sku_Name("Standard_Verizon")
)