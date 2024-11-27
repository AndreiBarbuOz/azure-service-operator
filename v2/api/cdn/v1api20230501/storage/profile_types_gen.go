// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=cdn.azure.com,resources=profiles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cdn.azure.com,resources={profiles/status,profiles/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230501.Profile
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/cdn.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}
type Profile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Profile_Spec   `json:"spec,omitempty"`
	Status            Profile_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Profile{}

// GetConditions returns the conditions of the resource
func (profile *Profile) GetConditions() conditions.Conditions {
	return profile.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (profile *Profile) SetConditions(conditions conditions.Conditions) {
	profile.Status.Conditions = conditions
}

var _ configmaps.Exporter = &Profile{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (profile *Profile) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if profile.Spec.OperatorSpec == nil {
		return nil
	}
	return profile.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &Profile{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (profile *Profile) SecretDestinationExpressions() []*core.DestinationExpression {
	if profile.Spec.OperatorSpec == nil {
		return nil
	}
	return profile.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &Profile{}

// AzureName returns the Azure name of the resource
func (profile *Profile) AzureName() string {
	return profile.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (profile Profile) GetAPIVersion() string {
	return "2023-05-01"
}

// GetResourceScope returns the scope of the resource
func (profile *Profile) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (profile *Profile) GetSpec() genruntime.ConvertibleSpec {
	return &profile.Spec
}

// GetStatus returns the status of this resource
func (profile *Profile) GetStatus() genruntime.ConvertibleStatus {
	return &profile.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (profile *Profile) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles"
func (profile *Profile) GetType() string {
	return "Microsoft.Cdn/profiles"
}

// NewEmptyStatus returns a new empty (blank) status
func (profile *Profile) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Profile_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (profile *Profile) Owner() *genruntime.ResourceReference {
	if profile.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(profile.Spec)
	return profile.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (profile *Profile) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Profile_STATUS); ok {
		profile.Status = *st
		return nil
	}

	// Convert status to required version
	var st Profile_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	profile.Status = st
	return nil
}

// Hub marks that this Profile is the hub type for conversion
func (profile *Profile) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (profile *Profile) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: profile.Spec.OriginalVersion,
		Kind:    "Profile",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230501.Profile
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/cdn.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}
type ProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Profile `json:"items"`
}

// Storage version of v1api20230501.Profile_Spec
type Profile_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                    string                  `json:"azureName,omitempty"`
	Identity                     *ManagedServiceIdentity `json:"identity,omitempty"`
	Location                     *string                 `json:"location,omitempty"`
	OperatorSpec                 *ProfileOperatorSpec    `json:"operatorSpec,omitempty"`
	OriginResponseTimeoutSeconds *int                    `json:"originResponseTimeoutSeconds,omitempty"`
	OriginalVersion              string                  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Sku         *Sku                               `json:"sku,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Profile_Spec{}

// ConvertSpecFrom populates our Profile_Spec from the provided source
func (profile *Profile_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == profile {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(profile)
}

// ConvertSpecTo populates the provided destination from our Profile_Spec
func (profile *Profile_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == profile {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(profile)
}

// Storage version of v1api20230501.Profile_STATUS
// A profile is a logical grouping of endpoints that share the same settings.
type Profile_STATUS struct {
	Conditions                   []conditions.Condition         `json:"conditions,omitempty"`
	ExtendedProperties           map[string]string              `json:"extendedProperties,omitempty"`
	FrontDoorId                  *string                        `json:"frontDoorId,omitempty"`
	Id                           *string                        `json:"id,omitempty"`
	Identity                     *ManagedServiceIdentity_STATUS `json:"identity,omitempty"`
	Kind                         *string                        `json:"kind,omitempty"`
	Location                     *string                        `json:"location,omitempty"`
	Name                         *string                        `json:"name,omitempty"`
	OriginResponseTimeoutSeconds *int                           `json:"originResponseTimeoutSeconds,omitempty"`
	PropertyBag                  genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	ProvisioningState            *string                        `json:"provisioningState,omitempty"`
	ResourceState                *string                        `json:"resourceState,omitempty"`
	Sku                          *Sku_STATUS                    `json:"sku,omitempty"`
	SystemData                   *SystemData_STATUS             `json:"systemData,omitempty"`
	Tags                         map[string]string              `json:"tags,omitempty"`
	Type                         *string                        `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Profile_STATUS{}

// ConvertStatusFrom populates our Profile_STATUS from the provided source
func (profile *Profile_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == profile {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(profile)
}

// ConvertStatusTo populates the provided destination from our Profile_STATUS
func (profile *Profile_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == profile {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(profile)
}

// Storage version of v1api20230501.ManagedServiceIdentity
// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	PropertyBag            genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Type                   *string                       `json:"type,omitempty"`
	UserAssignedIdentities []UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20230501.ManagedServiceIdentity_STATUS
// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity_STATUS struct {
	PrincipalId            *string                                `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	TenantId               *string                                `json:"tenantId,omitempty"`
	Type                   *string                                `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20230501.ProfileOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ProfileOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20230501.Sku
// Standard_Verizon = The SKU name for a Standard Verizon CDN profile.
// Premium_Verizon = The SKU name for a Premium Verizon
// CDN profile.
// Custom_Verizon = The SKU name for a Custom Verizon CDN profile.
// Standard_Akamai = The SKU name for an
// Akamai CDN profile.
// Standard_ChinaCdn = The SKU name for a China CDN profile for VOD, Web and download scenarios using
// GB based billing model.
// Standard_Microsoft = The SKU name for a Standard Microsoft CDN profile.
// Standard_AzureFrontDoor
// =  The SKU name for an Azure Front Door Standard profile.
// Premium_AzureFrontDoor = The SKU name for an Azure Front Door
// Premium profile.
// Standard_955BandWidth_ChinaCdn = The SKU name for a China CDN profile for VOD, Web and download
// scenarios using 95-5 peak bandwidth billing model.
// Standard_AvgBandWidth_ChinaCdn = The SKU name for a China CDN profile
// for VOD, Web and download scenarios using monthly average peak bandwidth billing model.
// StandardPlus_ChinaCdn = The SKU
// name for a China CDN profile for live-streaming using GB based billing model.
// StandardPlus_955BandWidth_ChinaCdn = The
// SKU name for a China CDN live-streaming profile using 95-5 peak bandwidth billing
// model.
// StandardPlus_AvgBandWidth_ChinaCdn = The SKU name for a China CDN live-streaming profile using monthly average
// peak bandwidth billing model.
type Sku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230501.Sku_STATUS
// Standard_Verizon = The SKU name for a Standard Verizon CDN profile.
// Premium_Verizon = The SKU name for a Premium Verizon
// CDN profile.
// Custom_Verizon = The SKU name for a Custom Verizon CDN profile.
// Standard_Akamai = The SKU name for an
// Akamai CDN profile.
// Standard_ChinaCdn = The SKU name for a China CDN profile for VOD, Web and download scenarios using
// GB based billing model.
// Standard_Microsoft = The SKU name for a Standard Microsoft CDN profile.
// Standard_AzureFrontDoor
// =  The SKU name for an Azure Front Door Standard profile.
// Premium_AzureFrontDoor = The SKU name for an Azure Front Door
// Premium profile.
// Standard_955BandWidth_ChinaCdn = The SKU name for a China CDN profile for VOD, Web and download
// scenarios using 95-5 peak bandwidth billing model.
// Standard_AvgBandWidth_ChinaCdn = The SKU name for a China CDN profile
// for VOD, Web and download scenarios using monthly average peak bandwidth billing model.
// StandardPlus_ChinaCdn = The SKU
// name for a China CDN profile for live-streaming using GB based billing model.
// StandardPlus_955BandWidth_ChinaCdn = The
// SKU name for a China CDN live-streaming profile using 95-5 peak bandwidth billing
// model.
// StandardPlus_AvgBandWidth_ChinaCdn = The SKU name for a China CDN live-streaming profile using monthly average
// peak bandwidth billing model.
type Sku_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230501.UserAssignedIdentity_STATUS
// User assigned identity properties
type UserAssignedIdentity_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230501.UserAssignedIdentityDetails
// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Reference   genruntime.ResourceReference `armReference:"Reference" json:"reference,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Profile{}, &ProfileList{})
}
