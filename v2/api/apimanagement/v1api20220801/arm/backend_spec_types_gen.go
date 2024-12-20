// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Backend_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Backend entity contract properties.
	Properties *BackendContractProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Backend_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (backend Backend_Spec) GetAPIVersion() string {
	return "2022-08-01"
}

// GetName returns the Name of the resource
func (backend *Backend_Spec) GetName() string {
	return backend.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/backends"
func (backend *Backend_Spec) GetType() string {
	return "Microsoft.ApiManagement/service/backends"
}

// Parameters supplied to the Create Backend operation.
type BackendContractProperties struct {
	// Credentials: Backend Credentials Contract Properties
	Credentials *BackendCredentialsContract `json:"credentials,omitempty"`

	// Description: Backend Description.
	Description *string `json:"description,omitempty"`

	// Properties: Backend Properties contract
	Properties *BackendProperties `json:"properties,omitempty"`

	// Protocol: Backend communication protocol.
	Protocol *BackendContractProperties_Protocol `json:"protocol,omitempty"`

	// Proxy: Backend gateway Contract Properties
	Proxy      *BackendProxyContract `json:"proxy,omitempty"`
	ResourceId *string               `json:"resourceId,omitempty"`

	// Title: Backend Title.
	Title *string `json:"title,omitempty"`

	// Tls: Backend TLS Properties
	Tls *BackendTlsProperties `json:"tls,omitempty"`

	// Url: Runtime Url of the Backend.
	Url *string `json:"url,omitempty"`
}

// +kubebuilder:validation:Enum={"http","soap"}
type BackendContractProperties_Protocol string

const (
	BackendContractProperties_Protocol_Http = BackendContractProperties_Protocol("http")
	BackendContractProperties_Protocol_Soap = BackendContractProperties_Protocol("soap")
)

// Mapping from string to BackendContractProperties_Protocol
var backendContractProperties_Protocol_Values = map[string]BackendContractProperties_Protocol{
	"http": BackendContractProperties_Protocol_Http,
	"soap": BackendContractProperties_Protocol_Soap,
}

// Details of the Credentials used to connect to Backend.
type BackendCredentialsContract struct {
	// Authorization: Authorization header authentication
	Authorization *BackendAuthorizationHeaderCredentials `json:"authorization,omitempty"`

	// Certificate: List of Client Certificate Thumbprints. Will be ignored if certificatesIds are provided.
	Certificate []string `json:"certificate,omitempty"`

	// CertificateIds: List of Client Certificate Ids.
	CertificateIds []string `json:"certificateIds,omitempty"`

	// Header: Header Parameter description.
	Header map[string][]string `json:"header,omitempty"`

	// Query: Query Parameter description.
	Query map[string][]string `json:"query,omitempty"`
}

// Properties specific to the Backend Type.
type BackendProperties struct {
	// ServiceFabricCluster: Backend Service Fabric Cluster Properties
	ServiceFabricCluster *BackendServiceFabricClusterProperties `json:"serviceFabricCluster,omitempty"`
}

// Details of the Backend WebProxy Server to use in the Request to Backend.
type BackendProxyContract struct {
	// Password: Password to connect to the WebProxy Server
	Password *string `json:"password,omitempty"`

	// Url: WebProxy Server AbsoluteUri property which includes the entire URI stored in the Uri instance, including all
	// fragments and query strings.
	Url *string `json:"url,omitempty"`

	// Username: Username to connect to the WebProxy server
	Username *string `json:"username,omitempty"`
}

// Properties controlling TLS Certificate Validation.
type BackendTlsProperties struct {
	// ValidateCertificateChain: Flag indicating whether SSL certificate chain validation should be done when using self-signed
	// certificates for this backend host.
	ValidateCertificateChain *bool `json:"validateCertificateChain,omitempty"`

	// ValidateCertificateName: Flag indicating whether SSL certificate name validation should be done when using self-signed
	// certificates for this backend host.
	ValidateCertificateName *bool `json:"validateCertificateName,omitempty"`
}

// Authorization header information.
type BackendAuthorizationHeaderCredentials struct {
	// Parameter: Authentication Parameter value.
	Parameter *string `json:"parameter,omitempty"`

	// Scheme: Authentication Scheme name.
	Scheme *string `json:"scheme,omitempty"`
}

// Properties of the Service Fabric Type Backend.
type BackendServiceFabricClusterProperties struct {
	// ClientCertificateId: The client certificate id for the management endpoint.
	ClientCertificateId *string `json:"clientCertificateId,omitempty"`

	// ClientCertificatethumbprint: The client certificate thumbprint for the management endpoint. Will be ignored if
	// certificatesIds are provided
	ClientCertificatethumbprint *string `json:"clientCertificatethumbprint,omitempty"`

	// ManagementEndpoints: The cluster management endpoint.
	ManagementEndpoints []string `json:"managementEndpoints,omitempty"`

	// MaxPartitionResolutionRetries: Maximum number of retries while attempting resolve the partition.
	MaxPartitionResolutionRetries *int `json:"maxPartitionResolutionRetries,omitempty"`

	// ServerCertificateThumbprints: Thumbprints of certificates cluster management service uses for tls communication
	ServerCertificateThumbprints []string `json:"serverCertificateThumbprints,omitempty"`

	// ServerX509Names: Server X509 Certificate Names Collection
	ServerX509Names []X509CertificateName `json:"serverX509Names,omitempty"`
}

// Properties of server X509Names.
type X509CertificateName struct {
	// IssuerCertificateThumbprint: Thumbprint for the Issuer of the Certificate.
	IssuerCertificateThumbprint *string `json:"issuerCertificateThumbprint,omitempty"`

	// Name: Common Name of the Certificate.
	Name *string `json:"name,omitempty"`
}
