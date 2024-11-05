// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type ServersFirewallRule_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *ServerFirewallRuleProperties_STATUS `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// The properties of a server firewall rule.
type ServerFirewallRuleProperties_STATUS struct {
	// EndIpAddress: The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to
	// startIpAddress. Use value '0.0.0.0' for all Azure-internal IP addresses.
	EndIpAddress *string `json:"endIpAddress,omitempty"`

	// StartIpAddress: The start IP address of the firewall rule. Must be IPv4 format. Use value '0.0.0.0' for all
	// Azure-internal IP addresses.
	StartIpAddress *string `json:"startIpAddress,omitempty"`
}