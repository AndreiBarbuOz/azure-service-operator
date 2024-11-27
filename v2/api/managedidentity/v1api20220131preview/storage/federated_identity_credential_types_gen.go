// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220131preview.FederatedIdentityCredential
// Generator information:
// - Generated from: /msi/resource-manager/Microsoft.ManagedIdentity/preview/2022-01-31-preview/ManagedIdentity.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}/federatedIdentityCredentials/{federatedIdentityCredentialResourceName}
type FederatedIdentityCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FederatedIdentityCredential_Spec   `json:"spec,omitempty"`
	Status            FederatedIdentityCredential_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FederatedIdentityCredential{}

// GetConditions returns the conditions of the resource
func (credential *FederatedIdentityCredential) GetConditions() conditions.Conditions {
	return credential.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (credential *FederatedIdentityCredential) SetConditions(conditions conditions.Conditions) {
	credential.Status.Conditions = conditions
}

var _ conversion.Convertible = &FederatedIdentityCredential{}

// ConvertFrom populates our FederatedIdentityCredential from the provided hub FederatedIdentityCredential
func (credential *FederatedIdentityCredential) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.FederatedIdentityCredential)
	if !ok {
		return fmt.Errorf("expected managedidentity/v1api20230131/storage/FederatedIdentityCredential but received %T instead", hub)
	}

	return credential.AssignProperties_From_FederatedIdentityCredential(source)
}

// ConvertTo populates the provided hub FederatedIdentityCredential from our FederatedIdentityCredential
func (credential *FederatedIdentityCredential) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.FederatedIdentityCredential)
	if !ok {
		return fmt.Errorf("expected managedidentity/v1api20230131/storage/FederatedIdentityCredential but received %T instead", hub)
	}

	return credential.AssignProperties_To_FederatedIdentityCredential(destination)
}

var _ configmaps.Exporter = &FederatedIdentityCredential{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (credential *FederatedIdentityCredential) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if credential.Spec.OperatorSpec == nil {
		return nil
	}
	return credential.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &FederatedIdentityCredential{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (credential *FederatedIdentityCredential) SecretDestinationExpressions() []*core.DestinationExpression {
	if credential.Spec.OperatorSpec == nil {
		return nil
	}
	return credential.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &FederatedIdentityCredential{}

// AzureName returns the Azure name of the resource
func (credential *FederatedIdentityCredential) AzureName() string {
	return credential.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-01-31-preview"
func (credential FederatedIdentityCredential) GetAPIVersion() string {
	return "2022-01-31-preview"
}

// GetResourceScope returns the scope of the resource
func (credential *FederatedIdentityCredential) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (credential *FederatedIdentityCredential) GetSpec() genruntime.ConvertibleSpec {
	return &credential.Spec
}

// GetStatus returns the status of this resource
func (credential *FederatedIdentityCredential) GetStatus() genruntime.ConvertibleStatus {
	return &credential.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (credential *FederatedIdentityCredential) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"
func (credential *FederatedIdentityCredential) GetType() string {
	return "Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"
}

// NewEmptyStatus returns a new empty (blank) status
func (credential *FederatedIdentityCredential) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FederatedIdentityCredential_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (credential *FederatedIdentityCredential) Owner() *genruntime.ResourceReference {
	if credential.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(credential.Spec)
	return credential.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (credential *FederatedIdentityCredential) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FederatedIdentityCredential_STATUS); ok {
		credential.Status = *st
		return nil
	}

	// Convert status to required version
	var st FederatedIdentityCredential_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	credential.Status = st
	return nil
}

// AssignProperties_From_FederatedIdentityCredential populates our FederatedIdentityCredential from the provided source FederatedIdentityCredential
func (credential *FederatedIdentityCredential) AssignProperties_From_FederatedIdentityCredential(source *storage.FederatedIdentityCredential) error {

	// ObjectMeta
	credential.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FederatedIdentityCredential_Spec
	err := spec.AssignProperties_From_FederatedIdentityCredential_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_FederatedIdentityCredential_Spec() to populate field Spec")
	}
	credential.Spec = spec

	// Status
	var status FederatedIdentityCredential_STATUS
	err = status.AssignProperties_From_FederatedIdentityCredential_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_FederatedIdentityCredential_STATUS() to populate field Status")
	}
	credential.Status = status

	// Invoke the augmentConversionForFederatedIdentityCredential interface (if implemented) to customize the conversion
	var credentialAsAny any = credential
	if augmentedCredential, ok := credentialAsAny.(augmentConversionForFederatedIdentityCredential); ok {
		err := augmentedCredential.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FederatedIdentityCredential populates the provided destination FederatedIdentityCredential from our FederatedIdentityCredential
func (credential *FederatedIdentityCredential) AssignProperties_To_FederatedIdentityCredential(destination *storage.FederatedIdentityCredential) error {

	// ObjectMeta
	destination.ObjectMeta = *credential.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.FederatedIdentityCredential_Spec
	err := credential.Spec.AssignProperties_To_FederatedIdentityCredential_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_FederatedIdentityCredential_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.FederatedIdentityCredential_STATUS
	err = credential.Status.AssignProperties_To_FederatedIdentityCredential_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_FederatedIdentityCredential_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForFederatedIdentityCredential interface (if implemented) to customize the conversion
	var credentialAsAny any = credential
	if augmentedCredential, ok := credentialAsAny.(augmentConversionForFederatedIdentityCredential); ok {
		err := augmentedCredential.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (credential *FederatedIdentityCredential) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: credential.Spec.OriginalVersion,
		Kind:    "FederatedIdentityCredential",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220131preview.FederatedIdentityCredential
// Generator information:
// - Generated from: /msi/resource-manager/Microsoft.ManagedIdentity/preview/2022-01-31-preview/ManagedIdentity.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}/federatedIdentityCredentials/{federatedIdentityCredentialResourceName}
type FederatedIdentityCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedIdentityCredential `json:"items"`
}

// Storage version of v1api20220131preview.APIVersion
// +kubebuilder:validation:Enum={"2022-01-31-preview"}
type APIVersion string

const APIVersion_Value = APIVersion("2022-01-31-preview")

type augmentConversionForFederatedIdentityCredential interface {
	AssignPropertiesFrom(src *storage.FederatedIdentityCredential) error
	AssignPropertiesTo(dst *storage.FederatedIdentityCredential) error
}

// Storage version of v1api20220131preview.FederatedIdentityCredential_Spec
type FederatedIdentityCredential_Spec struct {
	Audiences []string `json:"audiences,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string                                   `json:"azureName,omitempty"`
	Issuer           *string                                  `json:"issuer,omitempty" optionalConfigMapPair:"Issuer"`
	IssuerFromConfig *genruntime.ConfigMapReference           `json:"issuerFromConfig,omitempty" optionalConfigMapPair:"Issuer"`
	OperatorSpec     *FederatedIdentityCredentialOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion  string                                   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a managedidentity.azure.com/UserAssignedIdentity resource
	Owner             *genruntime.KnownResourceReference `group:"managedidentity.azure.com" json:"owner,omitempty" kind:"UserAssignedIdentity"`
	PropertyBag       genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Subject           *string                            `json:"subject,omitempty" optionalConfigMapPair:"Subject"`
	SubjectFromConfig *genruntime.ConfigMapReference     `json:"subjectFromConfig,omitempty" optionalConfigMapPair:"Subject"`
}

var _ genruntime.ConvertibleSpec = &FederatedIdentityCredential_Spec{}

// ConvertSpecFrom populates our FederatedIdentityCredential_Spec from the provided source
func (credential *FederatedIdentityCredential_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.FederatedIdentityCredential_Spec)
	if ok {
		// Populate our instance from source
		return credential.AssignProperties_From_FederatedIdentityCredential_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.FederatedIdentityCredential_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = credential.AssignProperties_From_FederatedIdentityCredential_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FederatedIdentityCredential_Spec
func (credential *FederatedIdentityCredential_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.FederatedIdentityCredential_Spec)
	if ok {
		// Populate destination from our instance
		return credential.AssignProperties_To_FederatedIdentityCredential_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.FederatedIdentityCredential_Spec{}
	err := credential.AssignProperties_To_FederatedIdentityCredential_Spec(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_FederatedIdentityCredential_Spec populates our FederatedIdentityCredential_Spec from the provided source FederatedIdentityCredential_Spec
func (credential *FederatedIdentityCredential_Spec) AssignProperties_From_FederatedIdentityCredential_Spec(source *storage.FederatedIdentityCredential_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Audiences
	credential.Audiences = genruntime.CloneSliceOfString(source.Audiences)

	// AzureName
	credential.AzureName = source.AzureName

	// Issuer
	credential.Issuer = genruntime.ClonePointerToString(source.Issuer)

	// IssuerFromConfig
	if source.IssuerFromConfig != nil {
		issuerFromConfig := source.IssuerFromConfig.Copy()
		credential.IssuerFromConfig = &issuerFromConfig
	} else {
		credential.IssuerFromConfig = nil
	}

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec FederatedIdentityCredentialOperatorSpec
		err := operatorSpec.AssignProperties_From_FederatedIdentityCredentialOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_FederatedIdentityCredentialOperatorSpec() to populate field OperatorSpec")
		}
		credential.OperatorSpec = &operatorSpec
	} else {
		credential.OperatorSpec = nil
	}

	// OriginalVersion
	credential.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		credential.Owner = &owner
	} else {
		credential.Owner = nil
	}

	// Subject
	credential.Subject = genruntime.ClonePointerToString(source.Subject)

	// SubjectFromConfig
	if source.SubjectFromConfig != nil {
		subjectFromConfig := source.SubjectFromConfig.Copy()
		credential.SubjectFromConfig = &subjectFromConfig
	} else {
		credential.SubjectFromConfig = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		credential.PropertyBag = propertyBag
	} else {
		credential.PropertyBag = nil
	}

	// Invoke the augmentConversionForFederatedIdentityCredential_Spec interface (if implemented) to customize the conversion
	var credentialAsAny any = credential
	if augmentedCredential, ok := credentialAsAny.(augmentConversionForFederatedIdentityCredential_Spec); ok {
		err := augmentedCredential.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FederatedIdentityCredential_Spec populates the provided destination FederatedIdentityCredential_Spec from our FederatedIdentityCredential_Spec
func (credential *FederatedIdentityCredential_Spec) AssignProperties_To_FederatedIdentityCredential_Spec(destination *storage.FederatedIdentityCredential_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(credential.PropertyBag)

	// Audiences
	destination.Audiences = genruntime.CloneSliceOfString(credential.Audiences)

	// AzureName
	destination.AzureName = credential.AzureName

	// Issuer
	destination.Issuer = genruntime.ClonePointerToString(credential.Issuer)

	// IssuerFromConfig
	if credential.IssuerFromConfig != nil {
		issuerFromConfig := credential.IssuerFromConfig.Copy()
		destination.IssuerFromConfig = &issuerFromConfig
	} else {
		destination.IssuerFromConfig = nil
	}

	// OperatorSpec
	if credential.OperatorSpec != nil {
		var operatorSpec storage.FederatedIdentityCredentialOperatorSpec
		err := credential.OperatorSpec.AssignProperties_To_FederatedIdentityCredentialOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_FederatedIdentityCredentialOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = credential.OriginalVersion

	// Owner
	if credential.Owner != nil {
		owner := credential.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Subject
	destination.Subject = genruntime.ClonePointerToString(credential.Subject)

	// SubjectFromConfig
	if credential.SubjectFromConfig != nil {
		subjectFromConfig := credential.SubjectFromConfig.Copy()
		destination.SubjectFromConfig = &subjectFromConfig
	} else {
		destination.SubjectFromConfig = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForFederatedIdentityCredential_Spec interface (if implemented) to customize the conversion
	var credentialAsAny any = credential
	if augmentedCredential, ok := credentialAsAny.(augmentConversionForFederatedIdentityCredential_Spec); ok {
		err := augmentedCredential.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20220131preview.FederatedIdentityCredential_STATUS
type FederatedIdentityCredential_STATUS struct {
	Audiences   []string               `json:"audiences,omitempty"`
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Issuer      *string                `json:"issuer,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Subject     *string                `json:"subject,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FederatedIdentityCredential_STATUS{}

// ConvertStatusFrom populates our FederatedIdentityCredential_STATUS from the provided source
func (credential *FederatedIdentityCredential_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.FederatedIdentityCredential_STATUS)
	if ok {
		// Populate our instance from source
		return credential.AssignProperties_From_FederatedIdentityCredential_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.FederatedIdentityCredential_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = credential.AssignProperties_From_FederatedIdentityCredential_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our FederatedIdentityCredential_STATUS
func (credential *FederatedIdentityCredential_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.FederatedIdentityCredential_STATUS)
	if ok {
		// Populate destination from our instance
		return credential.AssignProperties_To_FederatedIdentityCredential_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.FederatedIdentityCredential_STATUS{}
	err := credential.AssignProperties_To_FederatedIdentityCredential_STATUS(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_FederatedIdentityCredential_STATUS populates our FederatedIdentityCredential_STATUS from the provided source FederatedIdentityCredential_STATUS
func (credential *FederatedIdentityCredential_STATUS) AssignProperties_From_FederatedIdentityCredential_STATUS(source *storage.FederatedIdentityCredential_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Audiences
	credential.Audiences = genruntime.CloneSliceOfString(source.Audiences)

	// Conditions
	credential.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	credential.Id = genruntime.ClonePointerToString(source.Id)

	// Issuer
	credential.Issuer = genruntime.ClonePointerToString(source.Issuer)

	// Name
	credential.Name = genruntime.ClonePointerToString(source.Name)

	// Subject
	credential.Subject = genruntime.ClonePointerToString(source.Subject)

	// SystemData
	if source.SystemData != nil {
		propertyBag.Add("SystemData", *source.SystemData)
	} else {
		propertyBag.Remove("SystemData")
	}

	// Type
	credential.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		credential.PropertyBag = propertyBag
	} else {
		credential.PropertyBag = nil
	}

	// Invoke the augmentConversionForFederatedIdentityCredential_STATUS interface (if implemented) to customize the conversion
	var credentialAsAny any = credential
	if augmentedCredential, ok := credentialAsAny.(augmentConversionForFederatedIdentityCredential_STATUS); ok {
		err := augmentedCredential.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FederatedIdentityCredential_STATUS populates the provided destination FederatedIdentityCredential_STATUS from our FederatedIdentityCredential_STATUS
func (credential *FederatedIdentityCredential_STATUS) AssignProperties_To_FederatedIdentityCredential_STATUS(destination *storage.FederatedIdentityCredential_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(credential.PropertyBag)

	// Audiences
	destination.Audiences = genruntime.CloneSliceOfString(credential.Audiences)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(credential.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(credential.Id)

	// Issuer
	destination.Issuer = genruntime.ClonePointerToString(credential.Issuer)

	// Name
	destination.Name = genruntime.ClonePointerToString(credential.Name)

	// Subject
	destination.Subject = genruntime.ClonePointerToString(credential.Subject)

	// SystemData
	if propertyBag.Contains("SystemData") {
		var systemDatum storage.SystemData_STATUS
		err := propertyBag.Pull("SystemData", &systemDatum)
		if err != nil {
			return eris.Wrap(err, "pulling 'SystemData' from propertyBag")
		}

		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(credential.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForFederatedIdentityCredential_STATUS interface (if implemented) to customize the conversion
	var credentialAsAny any = credential
	if augmentedCredential, ok := credentialAsAny.(augmentConversionForFederatedIdentityCredential_STATUS); ok {
		err := augmentedCredential.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForFederatedIdentityCredential_Spec interface {
	AssignPropertiesFrom(src *storage.FederatedIdentityCredential_Spec) error
	AssignPropertiesTo(dst *storage.FederatedIdentityCredential_Spec) error
}

type augmentConversionForFederatedIdentityCredential_STATUS interface {
	AssignPropertiesFrom(src *storage.FederatedIdentityCredential_STATUS) error
	AssignPropertiesTo(dst *storage.FederatedIdentityCredential_STATUS) error
}

// Storage version of v1api20220131preview.FederatedIdentityCredentialOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type FederatedIdentityCredentialOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_FederatedIdentityCredentialOperatorSpec populates our FederatedIdentityCredentialOperatorSpec from the provided source FederatedIdentityCredentialOperatorSpec
func (operator *FederatedIdentityCredentialOperatorSpec) AssignProperties_From_FederatedIdentityCredentialOperatorSpec(source *storage.FederatedIdentityCredentialOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ConfigMapExpressions
	if source.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(source.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range source.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		operator.ConfigMapExpressions = configMapExpressionList
	} else {
		operator.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if source.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(source.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range source.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		operator.SecretExpressions = secretExpressionList
	} else {
		operator.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		operator.PropertyBag = propertyBag
	} else {
		operator.PropertyBag = nil
	}

	// Invoke the augmentConversionForFederatedIdentityCredentialOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForFederatedIdentityCredentialOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FederatedIdentityCredentialOperatorSpec populates the provided destination FederatedIdentityCredentialOperatorSpec from our FederatedIdentityCredentialOperatorSpec
func (operator *FederatedIdentityCredentialOperatorSpec) AssignProperties_To_FederatedIdentityCredentialOperatorSpec(destination *storage.FederatedIdentityCredentialOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(operator.PropertyBag)

	// ConfigMapExpressions
	if operator.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(operator.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range operator.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		destination.ConfigMapExpressions = configMapExpressionList
	} else {
		destination.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if operator.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(operator.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range operator.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		destination.SecretExpressions = secretExpressionList
	} else {
		destination.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForFederatedIdentityCredentialOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForFederatedIdentityCredentialOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForFederatedIdentityCredentialOperatorSpec interface {
	AssignPropertiesFrom(src *storage.FederatedIdentityCredentialOperatorSpec) error
	AssignPropertiesTo(dst *storage.FederatedIdentityCredentialOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&FederatedIdentityCredential{}, &FederatedIdentityCredentialList{})
}
