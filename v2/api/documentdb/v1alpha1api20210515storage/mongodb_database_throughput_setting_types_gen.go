// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=mongodbdatabasethroughputsettings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={mongodbdatabasethroughputsettings/status,mongodbdatabasethroughputsettings/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210515.MongodbDatabaseThroughputSetting
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_throughputSettings
type MongodbDatabaseThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsMongodbDatabasesThroughputSettings_Spec `json:"spec,omitempty"`
	Status            ThroughputSettingsGetResults_Status                     `json:"status,omitempty"`
}

var _ conditions.Conditioner = &MongodbDatabaseThroughputSetting{}

// GetConditions returns the conditions of the resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) GetConditions() conditions.Conditions {
	return mongodbDatabaseThroughputSetting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) SetConditions(conditions conditions.Conditions) {
	mongodbDatabaseThroughputSetting.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &MongodbDatabaseThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (mongodbDatabaseThroughputSetting MongodbDatabaseThroughputSetting) GetAPIVersion() string {
	return "2021-05-15"
}

// GetResourceKind returns the kind of the resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &mongodbDatabaseThroughputSetting.Spec
}

// GetStatus returns the status of this resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &mongodbDatabaseThroughputSetting.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/throughputSettings"
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ThroughputSettingsGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(mongodbDatabaseThroughputSetting.Spec)
	return &genruntime.ResourceReference{
		Group:     group,
		Kind:      kind,
		Namespace: mongodbDatabaseThroughputSetting.Namespace,
		Name:      mongodbDatabaseThroughputSetting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ThroughputSettingsGetResults_Status); ok {
		mongodbDatabaseThroughputSetting.Status = *st
		return nil
	}

	// Convert status to required version
	var st ThroughputSettingsGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	mongodbDatabaseThroughputSetting.Status = st
	return nil
}

// Hub marks that this MongodbDatabaseThroughputSetting is the hub type for conversion
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: mongodbDatabaseThroughputSetting.Spec.OriginalVersion,
		Kind:    "MongodbDatabaseThroughputSetting",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210515.MongodbDatabaseThroughputSetting
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_throughputSettings
type MongodbDatabaseThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabaseThroughputSetting `json:"items"`
}

//Storage version of v1alpha1api20210515.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
type DatabaseAccountsMongodbDatabasesThroughputSettings_Spec struct {
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner" kind:"MongodbDatabase"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsResource       `json:"resource,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsMongodbDatabasesThroughputSettings_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsMongodbDatabasesThroughputSettings_Spec from the provided source
func (databaseAccountsMongodbDatabasesThroughputSettingsSpec *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == databaseAccountsMongodbDatabasesThroughputSettingsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(databaseAccountsMongodbDatabasesThroughputSettingsSpec)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
func (databaseAccountsMongodbDatabasesThroughputSettingsSpec *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == databaseAccountsMongodbDatabasesThroughputSettingsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(databaseAccountsMongodbDatabasesThroughputSettingsSpec)
}

func init() {
	SchemeBuilder.Register(&MongodbDatabaseThroughputSetting{}, &MongodbDatabaseThroughputSettingList{})
}