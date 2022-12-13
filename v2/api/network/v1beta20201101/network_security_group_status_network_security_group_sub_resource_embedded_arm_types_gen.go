// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

// NetworkSecurityGroup resource.
type NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the network security group.
	Properties *NetworkSecurityGroupPropertiesFormat_STATUS_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Network Security Group resource.
type NetworkSecurityGroupPropertiesFormat_STATUS_ARM struct {
	// DefaultSecurityRules: The default security rules of network security group.
	DefaultSecurityRules []SecurityRule_STATUS_ARM `json:"defaultSecurityRules,omitempty"`

	// FlowLogs: A collection of references to flow log resources.
	FlowLogs []FlowLog_STATUS_ARM `json:"flowLogs,omitempty"`

	// NetworkInterfaces: A collection of references to network interfaces.
	NetworkInterfaces []NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM `json:"networkInterfaces,omitempty"`

	// ProvisioningState: The provisioning state of the network security group resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ResourceGuid: The resource GUID property of the network security group resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`

	// SecurityRules: A collection of security rules of the network security group.
	SecurityRules []SecurityRule_STATUS_ARM `json:"securityRules,omitempty"`

	// Subnets: A collection of references to subnets.
	Subnets []Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM `json:"subnets,omitempty"`
}

// A flow log resource.
type FlowLog_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// A network interface in a resource group.
type NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Network security rule.
type SecurityRule_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Subnet in a virtual network resource.
type Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}
