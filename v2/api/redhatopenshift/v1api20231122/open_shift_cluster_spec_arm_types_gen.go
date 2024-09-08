// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20231122

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type OpenShiftCluster_Spec_ARM struct {
	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: The cluster properties.
	Properties *OpenShiftClusterProperties_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &OpenShiftCluster_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-11-22"
func (cluster OpenShiftCluster_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (cluster *OpenShiftCluster_Spec_ARM) GetName() string {
	return cluster.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.RedHatOpenShift/openShiftClusters"
func (cluster *OpenShiftCluster_Spec_ARM) GetType() string {
	return "Microsoft.RedHatOpenShift/openShiftClusters"
}

// OpenShiftClusterProperties represents an OpenShift cluster's properties.
type OpenShiftClusterProperties_ARM struct {
	// ApiserverProfile: The cluster API server profile.
	ApiserverProfile *APIServerProfile_ARM `json:"apiserverProfile,omitempty"`

	// ClusterProfile: The cluster profile.
	ClusterProfile *ClusterProfile_ARM `json:"clusterProfile,omitempty"`

	// IngressProfiles: The cluster ingress profiles.
	IngressProfiles []IngressProfile_ARM `json:"ingressProfiles,omitempty"`

	// MasterProfile: The cluster master profile.
	MasterProfile *MasterProfile_ARM `json:"masterProfile,omitempty"`

	// NetworkProfile: The cluster network profile.
	NetworkProfile *NetworkProfile_ARM `json:"networkProfile,omitempty"`

	// ServicePrincipalProfile: The cluster service principal profile.
	ServicePrincipalProfile *ServicePrincipalProfile_ARM `json:"servicePrincipalProfile,omitempty"`

	// WorkerProfiles: The cluster worker profiles.
	WorkerProfiles []WorkerProfile_ARM `json:"workerProfiles,omitempty"`
}

// APIServerProfile represents an API server profile.
type APIServerProfile_ARM struct {
	// Visibility: API server visibility.
	Visibility *Visibility_ARM `json:"visibility,omitempty"`
}

// ClusterProfile represents a cluster profile.
type ClusterProfile_ARM struct {
	// Domain: The domain for the cluster.
	Domain *string `json:"domain,omitempty"`

	// FipsValidatedModules: If FIPS validated crypto modules are used
	FipsValidatedModules *FipsValidatedModules_ARM `json:"fipsValidatedModules,omitempty"`

	// PullSecret: The pull secret for the cluster.
	PullSecret *string `json:"pullSecret,omitempty"`

	// ResourceGroupId: The ID of the cluster resource group.
	ResourceGroupId *string `json:"resourceGroupId,omitempty"`

	// Version: The version of the cluster.
	Version *string `json:"version,omitempty"`
}

// IngressProfile represents an ingress profile.
type IngressProfile_ARM struct {
	// Name: The ingress profile name.
	Name *string `json:"name,omitempty"`

	// Visibility: Ingress visibility.
	Visibility *Visibility_ARM `json:"visibility,omitempty"`
}

// MasterProfile represents a master profile.
type MasterProfile_ARM struct {
	DiskEncryptionSetId *string `json:"diskEncryptionSetId,omitempty"`

	// EncryptionAtHost: Whether master virtual machines are encrypted at host.
	EncryptionAtHost *EncryptionAtHost_ARM `json:"encryptionAtHost,omitempty"`
	SubnetId         *string               `json:"subnetId,omitempty"`

	// VmSize: The size of the master VMs.
	VmSize *string `json:"vmSize,omitempty"`
}

// NetworkProfile represents a network profile.
type NetworkProfile_ARM struct {
	// LoadBalancerProfile: The cluster load balancer profile.
	LoadBalancerProfile *LoadBalancerProfile_ARM `json:"loadBalancerProfile,omitempty"`

	// OutboundType: The OutboundType used for egress traffic.
	OutboundType *OutboundType_ARM `json:"outboundType,omitempty"`

	// PodCidr: The CIDR used for OpenShift/Kubernetes Pods.
	PodCidr *string `json:"podCidr,omitempty"`

	// PreconfiguredNSG: Specifies whether subnets are pre-attached with an NSG
	PreconfiguredNSG *PreconfiguredNSG_ARM `json:"preconfiguredNSG,omitempty"`

	// ServiceCidr: The CIDR used for OpenShift/Kubernetes Services.
	ServiceCidr *string `json:"serviceCidr,omitempty"`
}

// ServicePrincipalProfile represents a service principal profile.
type ServicePrincipalProfile_ARM struct {
	// ClientId: The client ID used for the cluster.
	ClientId *string `json:"clientId,omitempty" optionalConfigMapPair:"ClientId"`

	// ClientSecret: The client secret used for the cluster.
	ClientSecret *string `json:"clientSecret,omitempty"`
}

// WorkerProfile represents a worker profile.
type WorkerProfile_ARM struct {
	// Count: The number of worker VMs.
	Count               *int    `json:"count,omitempty"`
	DiskEncryptionSetId *string `json:"diskEncryptionSetId,omitempty"`

	// DiskSizeGB: The disk size of the worker VMs.
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// EncryptionAtHost: Whether master virtual machines are encrypted at host.
	EncryptionAtHost *EncryptionAtHost_ARM `json:"encryptionAtHost,omitempty"`

	// Name: The worker profile name.
	Name     *string `json:"name,omitempty"`
	SubnetId *string `json:"subnetId,omitempty"`

	// VmSize: The size of the worker VMs.
	VmSize *string `json:"vmSize,omitempty"`
}

// EncryptionAtHost represents encryption at host state
// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type EncryptionAtHost_ARM string

const (
	EncryptionAtHost_ARM_Disabled = EncryptionAtHost_ARM("Disabled")
	EncryptionAtHost_ARM_Enabled  = EncryptionAtHost_ARM("Enabled")
)

// Mapping from string to EncryptionAtHost_ARM
var encryptionAtHost_ARM_Values = map[string]EncryptionAtHost_ARM{
	"disabled": EncryptionAtHost_ARM_Disabled,
	"enabled":  EncryptionAtHost_ARM_Enabled,
}

// FipsValidatedModules determines if FIPS is used.
// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type FipsValidatedModules_ARM string

const (
	FipsValidatedModules_ARM_Disabled = FipsValidatedModules_ARM("Disabled")
	FipsValidatedModules_ARM_Enabled  = FipsValidatedModules_ARM("Enabled")
)

// Mapping from string to FipsValidatedModules_ARM
var fipsValidatedModules_ARM_Values = map[string]FipsValidatedModules_ARM{
	"disabled": FipsValidatedModules_ARM_Disabled,
	"enabled":  FipsValidatedModules_ARM_Enabled,
}

// LoadBalancerProfile represents the profile of the cluster public load balancer.
type LoadBalancerProfile_ARM struct {
	// ManagedOutboundIps: The desired managed outbound IPs for the cluster public load balancer.
	ManagedOutboundIps *ManagedOutboundIPs_ARM `json:"managedOutboundIps,omitempty"`
}

// The outbound routing strategy used to provide your cluster egress to the internet.
// +kubebuilder:validation:Enum={"Loadbalancer","UserDefinedRouting"}
type OutboundType_ARM string

const (
	OutboundType_ARM_Loadbalancer       = OutboundType_ARM("Loadbalancer")
	OutboundType_ARM_UserDefinedRouting = OutboundType_ARM("UserDefinedRouting")
)

// Mapping from string to OutboundType_ARM
var outboundType_ARM_Values = map[string]OutboundType_ARM{
	"loadbalancer":       OutboundType_ARM_Loadbalancer,
	"userdefinedrouting": OutboundType_ARM_UserDefinedRouting,
}

// PreconfiguredNSG represents whether customers want to use their own NSG attached to the subnets
// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type PreconfiguredNSG_ARM string

const (
	PreconfiguredNSG_ARM_Disabled = PreconfiguredNSG_ARM("Disabled")
	PreconfiguredNSG_ARM_Enabled  = PreconfiguredNSG_ARM("Enabled")
)

// Mapping from string to PreconfiguredNSG_ARM
var preconfiguredNSG_ARM_Values = map[string]PreconfiguredNSG_ARM{
	"disabled": PreconfiguredNSG_ARM_Disabled,
	"enabled":  PreconfiguredNSG_ARM_Enabled,
}

// Visibility represents visibility.
// +kubebuilder:validation:Enum={"Private","Public"}
type Visibility_ARM string

const (
	Visibility_ARM_Private = Visibility_ARM("Private")
	Visibility_ARM_Public  = Visibility_ARM("Public")
)

// Mapping from string to Visibility_ARM
var visibility_ARM_Values = map[string]Visibility_ARM{
	"private": Visibility_ARM_Private,
	"public":  Visibility_ARM_Public,
}

// ManagedOutboundIPs represents the desired managed outbound IPs for the cluster public load balancer.
type ManagedOutboundIPs_ARM struct {
	// Count: Count represents the desired number of IPv4 outbound IPs created and managed by Azure for the cluster public load
	// balancer.  Allowed values are in the range of 1 - 20.  The default value is 1.
	Count *int `json:"count,omitempty"`
}