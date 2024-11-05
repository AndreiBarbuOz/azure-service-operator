// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type TrafficManagerProfile_STATUS struct {
	// Id: Fully qualified resource Id for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The Azure Region where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The properties of the Traffic Manager profile.
	Properties *ProfileProperties_STATUS `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type *string `json:"type,omitempty"`
}

// Class representing the Traffic Manager profile properties.
type ProfileProperties_STATUS struct {
	// AllowedEndpointRecordTypes: The list of allowed endpoint record types.
	AllowedEndpointRecordTypes []AllowedEndpointRecordType_STATUS `json:"allowedEndpointRecordTypes,omitempty"`

	// DnsConfig: The DNS settings of the Traffic Manager profile.
	DnsConfig *DnsConfig_STATUS `json:"dnsConfig,omitempty"`

	// Endpoints: The list of endpoints in the Traffic Manager profile.
	Endpoints []Endpoint_STATUS `json:"endpoints,omitempty"`

	// MaxReturn: Maximum number of endpoints to be returned for MultiValue routing type.
	MaxReturn *int `json:"maxReturn,omitempty"`

	// MonitorConfig: The endpoint monitoring settings of the Traffic Manager profile.
	MonitorConfig *MonitorConfig_STATUS `json:"monitorConfig,omitempty"`

	// ProfileStatus: The status of the Traffic Manager profile.
	ProfileStatus *ProfileProperties_ProfileStatus_STATUS `json:"profileStatus,omitempty"`

	// TrafficRoutingMethod: The traffic routing method of the Traffic Manager profile.
	TrafficRoutingMethod *ProfileProperties_TrafficRoutingMethod_STATUS `json:"trafficRoutingMethod,omitempty"`

	// TrafficViewEnrollmentStatus: Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile.
	// Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
	TrafficViewEnrollmentStatus *ProfileProperties_TrafficViewEnrollmentStatus_STATUS `json:"trafficViewEnrollmentStatus,omitempty"`
}

// The allowed type DNS record types for this profile.
type AllowedEndpointRecordType_STATUS string

const (
	AllowedEndpointRecordType_STATUS_Any         = AllowedEndpointRecordType_STATUS("Any")
	AllowedEndpointRecordType_STATUS_DomainName  = AllowedEndpointRecordType_STATUS("DomainName")
	AllowedEndpointRecordType_STATUS_IPv4Address = AllowedEndpointRecordType_STATUS("IPv4Address")
	AllowedEndpointRecordType_STATUS_IPv6Address = AllowedEndpointRecordType_STATUS("IPv6Address")
)

// Mapping from string to AllowedEndpointRecordType_STATUS
var allowedEndpointRecordType_STATUS_Values = map[string]AllowedEndpointRecordType_STATUS{
	"any":         AllowedEndpointRecordType_STATUS_Any,
	"domainname":  AllowedEndpointRecordType_STATUS_DomainName,
	"ipv4address": AllowedEndpointRecordType_STATUS_IPv4Address,
	"ipv6address": AllowedEndpointRecordType_STATUS_IPv6Address,
}

// Class containing DNS settings in a Traffic Manager profile.
type DnsConfig_STATUS struct {
	// Fqdn: The fully-qualified domain name (FQDN) of the Traffic Manager profile. This is formed from the concatenation of
	// the RelativeName with the DNS domain used by Azure Traffic Manager.
	Fqdn *string `json:"fqdn,omitempty"`

	// RelativeName: The relative DNS name provided by this Traffic Manager profile. This value is combined with the DNS domain
	// name used by Azure Traffic Manager to form the fully-qualified domain name (FQDN) of the profile.
	RelativeName *string `json:"relativeName,omitempty"`

	// Ttl: The DNS Time-To-Live (TTL), in seconds. This informs the local DNS resolvers and DNS clients how long to cache DNS
	// responses provided by this Traffic Manager profile.
	Ttl *int `json:"ttl,omitempty"`
}

// Class representing a Traffic Manager endpoint.
type Endpoint_STATUS struct {
	// Id: Fully qualified resource Id for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	Id *string `json:"id,omitempty"`
}

// Class containing endpoint monitoring settings in a Traffic Manager profile.
type MonitorConfig_STATUS struct {
	// CustomHeaders: List of custom headers.
	CustomHeaders []MonitorConfig_CustomHeaders_STATUS `json:"customHeaders,omitempty"`

	// ExpectedStatusCodeRanges: List of expected status code ranges.
	ExpectedStatusCodeRanges []MonitorConfig_ExpectedStatusCodeRanges_STATUS `json:"expectedStatusCodeRanges,omitempty"`

	// IntervalInSeconds: The monitor interval for endpoints in this profile. This is the interval at which Traffic Manager
	// will check the health of each endpoint in this profile.
	IntervalInSeconds *int `json:"intervalInSeconds,omitempty"`

	// Path: The path relative to the endpoint domain name used to probe for endpoint health.
	Path *string `json:"path,omitempty"`

	// Port: The TCP port used to probe for endpoint health.
	Port *int `json:"port,omitempty"`

	// ProfileMonitorStatus: The profile-level monitoring status of the Traffic Manager profile.
	ProfileMonitorStatus *MonitorConfig_ProfileMonitorStatus_STATUS `json:"profileMonitorStatus,omitempty"`

	// Protocol: The protocol (HTTP, HTTPS or TCP) used to probe for endpoint health.
	Protocol *MonitorConfig_Protocol_STATUS `json:"protocol,omitempty"`

	// TimeoutInSeconds: The monitor timeout for endpoints in this profile. This is the time that Traffic Manager allows
	// endpoints in this profile to response to the health check.
	TimeoutInSeconds *int `json:"timeoutInSeconds,omitempty"`

	// ToleratedNumberOfFailures: The number of consecutive failed health check that Traffic Manager tolerates before declaring
	// an endpoint in this profile Degraded after the next failed health check.
	ToleratedNumberOfFailures *int `json:"toleratedNumberOfFailures,omitempty"`
}

type ProfileProperties_ProfileStatus_STATUS string

const (
	ProfileProperties_ProfileStatus_STATUS_Disabled = ProfileProperties_ProfileStatus_STATUS("Disabled")
	ProfileProperties_ProfileStatus_STATUS_Enabled  = ProfileProperties_ProfileStatus_STATUS("Enabled")
)

// Mapping from string to ProfileProperties_ProfileStatus_STATUS
var profileProperties_ProfileStatus_STATUS_Values = map[string]ProfileProperties_ProfileStatus_STATUS{
	"disabled": ProfileProperties_ProfileStatus_STATUS_Disabled,
	"enabled":  ProfileProperties_ProfileStatus_STATUS_Enabled,
}

type ProfileProperties_TrafficRoutingMethod_STATUS string

const (
	ProfileProperties_TrafficRoutingMethod_STATUS_Geographic  = ProfileProperties_TrafficRoutingMethod_STATUS("Geographic")
	ProfileProperties_TrafficRoutingMethod_STATUS_MultiValue  = ProfileProperties_TrafficRoutingMethod_STATUS("MultiValue")
	ProfileProperties_TrafficRoutingMethod_STATUS_Performance = ProfileProperties_TrafficRoutingMethod_STATUS("Performance")
	ProfileProperties_TrafficRoutingMethod_STATUS_Priority    = ProfileProperties_TrafficRoutingMethod_STATUS("Priority")
	ProfileProperties_TrafficRoutingMethod_STATUS_Subnet      = ProfileProperties_TrafficRoutingMethod_STATUS("Subnet")
	ProfileProperties_TrafficRoutingMethod_STATUS_Weighted    = ProfileProperties_TrafficRoutingMethod_STATUS("Weighted")
)

// Mapping from string to ProfileProperties_TrafficRoutingMethod_STATUS
var profileProperties_TrafficRoutingMethod_STATUS_Values = map[string]ProfileProperties_TrafficRoutingMethod_STATUS{
	"geographic":  ProfileProperties_TrafficRoutingMethod_STATUS_Geographic,
	"multivalue":  ProfileProperties_TrafficRoutingMethod_STATUS_MultiValue,
	"performance": ProfileProperties_TrafficRoutingMethod_STATUS_Performance,
	"priority":    ProfileProperties_TrafficRoutingMethod_STATUS_Priority,
	"subnet":      ProfileProperties_TrafficRoutingMethod_STATUS_Subnet,
	"weighted":    ProfileProperties_TrafficRoutingMethod_STATUS_Weighted,
}

type ProfileProperties_TrafficViewEnrollmentStatus_STATUS string

const (
	ProfileProperties_TrafficViewEnrollmentStatus_STATUS_Disabled = ProfileProperties_TrafficViewEnrollmentStatus_STATUS("Disabled")
	ProfileProperties_TrafficViewEnrollmentStatus_STATUS_Enabled  = ProfileProperties_TrafficViewEnrollmentStatus_STATUS("Enabled")
)

// Mapping from string to ProfileProperties_TrafficViewEnrollmentStatus_STATUS
var profileProperties_TrafficViewEnrollmentStatus_STATUS_Values = map[string]ProfileProperties_TrafficViewEnrollmentStatus_STATUS{
	"disabled": ProfileProperties_TrafficViewEnrollmentStatus_STATUS_Disabled,
	"enabled":  ProfileProperties_TrafficViewEnrollmentStatus_STATUS_Enabled,
}

type MonitorConfig_CustomHeaders_STATUS struct {
	// Name: Header name.
	Name *string `json:"name,omitempty"`

	// Value: Header value.
	Value *string `json:"value,omitempty"`
}

type MonitorConfig_ExpectedStatusCodeRanges_STATUS struct {
	// Max: Max status code.
	Max *int `json:"max,omitempty"`

	// Min: Min status code.
	Min *int `json:"min,omitempty"`
}

type MonitorConfig_ProfileMonitorStatus_STATUS string

const (
	MonitorConfig_ProfileMonitorStatus_STATUS_CheckingEndpoints = MonitorConfig_ProfileMonitorStatus_STATUS("CheckingEndpoints")
	MonitorConfig_ProfileMonitorStatus_STATUS_Degraded          = MonitorConfig_ProfileMonitorStatus_STATUS("Degraded")
	MonitorConfig_ProfileMonitorStatus_STATUS_Disabled          = MonitorConfig_ProfileMonitorStatus_STATUS("Disabled")
	MonitorConfig_ProfileMonitorStatus_STATUS_Inactive          = MonitorConfig_ProfileMonitorStatus_STATUS("Inactive")
	MonitorConfig_ProfileMonitorStatus_STATUS_Online            = MonitorConfig_ProfileMonitorStatus_STATUS("Online")
)

// Mapping from string to MonitorConfig_ProfileMonitorStatus_STATUS
var monitorConfig_ProfileMonitorStatus_STATUS_Values = map[string]MonitorConfig_ProfileMonitorStatus_STATUS{
	"checkingendpoints": MonitorConfig_ProfileMonitorStatus_STATUS_CheckingEndpoints,
	"degraded":          MonitorConfig_ProfileMonitorStatus_STATUS_Degraded,
	"disabled":          MonitorConfig_ProfileMonitorStatus_STATUS_Disabled,
	"inactive":          MonitorConfig_ProfileMonitorStatus_STATUS_Inactive,
	"online":            MonitorConfig_ProfileMonitorStatus_STATUS_Online,
}

type MonitorConfig_Protocol_STATUS string

const (
	MonitorConfig_Protocol_STATUS_HTTP  = MonitorConfig_Protocol_STATUS("HTTP")
	MonitorConfig_Protocol_STATUS_HTTPS = MonitorConfig_Protocol_STATUS("HTTPS")
	MonitorConfig_Protocol_STATUS_TCP   = MonitorConfig_Protocol_STATUS("TCP")
)

// Mapping from string to MonitorConfig_Protocol_STATUS
var monitorConfig_Protocol_STATUS_Values = map[string]MonitorConfig_Protocol_STATUS{
	"http":  MonitorConfig_Protocol_STATUS_HTTP,
	"https": MonitorConfig_Protocol_STATUS_HTTPS,
	"tcp":   MonitorConfig_Protocol_STATUS_TCP,
}