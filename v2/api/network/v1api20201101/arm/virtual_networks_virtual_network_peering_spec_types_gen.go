// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type VirtualNetworksVirtualNetworkPeering_Spec struct {
	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name string `json:"name,omitempty"`

	// Properties: Properties of the virtual network peering.
	Properties *VirtualNetworkPeeringPropertiesFormat `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetworksVirtualNetworkPeering_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (peering VirtualNetworksVirtualNetworkPeering_Spec) GetAPIVersion() string {
	return "2020-11-01"
}

// GetName returns the Name of the resource
func (peering *VirtualNetworksVirtualNetworkPeering_Spec) GetName() string {
	return peering.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks/virtualNetworkPeerings"
func (peering *VirtualNetworksVirtualNetworkPeering_Spec) GetType() string {
	return "Microsoft.Network/virtualNetworks/virtualNetworkPeerings"
}

// Properties of the virtual network peering.
type VirtualNetworkPeeringPropertiesFormat struct {
	// AllowForwardedTraffic: Whether the forwarded traffic from the VMs in the local virtual network will be
	// allowed/disallowed in remote virtual network.
	AllowForwardedTraffic *bool `json:"allowForwardedTraffic,omitempty"`

	// AllowGatewayTransit: If gateway links can be used in remote virtual networking to link to this virtual network.
	AllowGatewayTransit *bool `json:"allowGatewayTransit,omitempty"`

	// AllowVirtualNetworkAccess: Whether the VMs in the local virtual network space would be able to access the VMs in remote
	// virtual network space.
	AllowVirtualNetworkAccess *bool `json:"allowVirtualNetworkAccess,omitempty"`

	// DoNotVerifyRemoteGateways: If we need to verify the provisioning state of the remote gateway.
	DoNotVerifyRemoteGateways *bool `json:"doNotVerifyRemoteGateways,omitempty"`

	// PeeringState: The status of the virtual network peering.
	PeeringState *VirtualNetworkPeeringPropertiesFormat_PeeringState `json:"peeringState,omitempty"`

	// RemoteAddressSpace: The reference to the remote virtual network address space.
	RemoteAddressSpace *AddressSpace `json:"remoteAddressSpace,omitempty"`

	// RemoteBgpCommunities: The reference to the remote virtual network's Bgp Communities.
	RemoteBgpCommunities *VirtualNetworkBgpCommunities `json:"remoteBgpCommunities,omitempty"`

	// RemoteVirtualNetwork: The reference to the remote virtual network. The remote virtual network can be in the same or
	// different region (preview). See here to register for the preview and learn more
	// (https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering).
	RemoteVirtualNetwork *SubResource `json:"remoteVirtualNetwork,omitempty"`

	// UseRemoteGateways: If remote gateways can be used on this virtual network. If the flag is set to true, and
	// allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for
	// transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a
	// gateway.
	UseRemoteGateways *bool `json:"useRemoteGateways,omitempty"`
}

// +kubebuilder:validation:Enum={"Connected","Disconnected","Initiated"}
type VirtualNetworkPeeringPropertiesFormat_PeeringState string

const (
	VirtualNetworkPeeringPropertiesFormat_PeeringState_Connected    = VirtualNetworkPeeringPropertiesFormat_PeeringState("Connected")
	VirtualNetworkPeeringPropertiesFormat_PeeringState_Disconnected = VirtualNetworkPeeringPropertiesFormat_PeeringState("Disconnected")
	VirtualNetworkPeeringPropertiesFormat_PeeringState_Initiated    = VirtualNetworkPeeringPropertiesFormat_PeeringState("Initiated")
)

// Mapping from string to VirtualNetworkPeeringPropertiesFormat_PeeringState
var virtualNetworkPeeringPropertiesFormat_PeeringState_Values = map[string]VirtualNetworkPeeringPropertiesFormat_PeeringState{
	"connected":    VirtualNetworkPeeringPropertiesFormat_PeeringState_Connected,
	"disconnected": VirtualNetworkPeeringPropertiesFormat_PeeringState_Disconnected,
	"initiated":    VirtualNetworkPeeringPropertiesFormat_PeeringState_Initiated,
}