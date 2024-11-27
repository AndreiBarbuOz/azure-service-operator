// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101/storage"
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
// Storage version of v1api20210101preview.NamespacesTopic
// Generator information:
// - Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/preview/2021-01-01-preview/topics.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}
type NamespacesTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesTopic_Spec   `json:"spec,omitempty"`
	Status            NamespacesTopic_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesTopic{}

// GetConditions returns the conditions of the resource
func (topic *NamespacesTopic) GetConditions() conditions.Conditions {
	return topic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (topic *NamespacesTopic) SetConditions(conditions conditions.Conditions) {
	topic.Status.Conditions = conditions
}

var _ conversion.Convertible = &NamespacesTopic{}

// ConvertFrom populates our NamespacesTopic from the provided hub NamespacesTopic
func (topic *NamespacesTopic) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.NamespacesTopic)
	if !ok {
		return fmt.Errorf("expected servicebus/v1api20211101/storage/NamespacesTopic but received %T instead", hub)
	}

	return topic.AssignProperties_From_NamespacesTopic(source)
}

// ConvertTo populates the provided hub NamespacesTopic from our NamespacesTopic
func (topic *NamespacesTopic) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.NamespacesTopic)
	if !ok {
		return fmt.Errorf("expected servicebus/v1api20211101/storage/NamespacesTopic but received %T instead", hub)
	}

	return topic.AssignProperties_To_NamespacesTopic(destination)
}

var _ configmaps.Exporter = &NamespacesTopic{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (topic *NamespacesTopic) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if topic.Spec.OperatorSpec == nil {
		return nil
	}
	return topic.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &NamespacesTopic{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (topic *NamespacesTopic) SecretDestinationExpressions() []*core.DestinationExpression {
	if topic.Spec.OperatorSpec == nil {
		return nil
	}
	return topic.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &NamespacesTopic{}

// AzureName returns the Azure name of the resource
func (topic *NamespacesTopic) AzureName() string {
	return topic.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (topic NamespacesTopic) GetAPIVersion() string {
	return "2021-01-01-preview"
}

// GetResourceScope returns the scope of the resource
func (topic *NamespacesTopic) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (topic *NamespacesTopic) GetSpec() genruntime.ConvertibleSpec {
	return &topic.Spec
}

// GetStatus returns the status of this resource
func (topic *NamespacesTopic) GetStatus() genruntime.ConvertibleStatus {
	return &topic.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (topic *NamespacesTopic) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/topics"
func (topic *NamespacesTopic) GetType() string {
	return "Microsoft.ServiceBus/namespaces/topics"
}

// NewEmptyStatus returns a new empty (blank) status
func (topic *NamespacesTopic) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &NamespacesTopic_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (topic *NamespacesTopic) Owner() *genruntime.ResourceReference {
	if topic.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(topic.Spec)
	return topic.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (topic *NamespacesTopic) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*NamespacesTopic_STATUS); ok {
		topic.Status = *st
		return nil
	}

	// Convert status to required version
	var st NamespacesTopic_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	topic.Status = st
	return nil
}

// AssignProperties_From_NamespacesTopic populates our NamespacesTopic from the provided source NamespacesTopic
func (topic *NamespacesTopic) AssignProperties_From_NamespacesTopic(source *storage.NamespacesTopic) error {

	// ObjectMeta
	topic.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec NamespacesTopic_Spec
	err := spec.AssignProperties_From_NamespacesTopic_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_NamespacesTopic_Spec() to populate field Spec")
	}
	topic.Spec = spec

	// Status
	var status NamespacesTopic_STATUS
	err = status.AssignProperties_From_NamespacesTopic_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_NamespacesTopic_STATUS() to populate field Status")
	}
	topic.Status = status

	// Invoke the augmentConversionForNamespacesTopic interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForNamespacesTopic); ok {
		err := augmentedTopic.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_NamespacesTopic populates the provided destination NamespacesTopic from our NamespacesTopic
func (topic *NamespacesTopic) AssignProperties_To_NamespacesTopic(destination *storage.NamespacesTopic) error {

	// ObjectMeta
	destination.ObjectMeta = *topic.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.NamespacesTopic_Spec
	err := topic.Spec.AssignProperties_To_NamespacesTopic_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_NamespacesTopic_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.NamespacesTopic_STATUS
	err = topic.Status.AssignProperties_To_NamespacesTopic_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_NamespacesTopic_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForNamespacesTopic interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForNamespacesTopic); ok {
		err := augmentedTopic.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (topic *NamespacesTopic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: topic.Spec.OriginalVersion,
		Kind:    "NamespacesTopic",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210101preview.NamespacesTopic
// Generator information:
// - Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/preview/2021-01-01-preview/topics.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}
type NamespacesTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesTopic `json:"items"`
}

type augmentConversionForNamespacesTopic interface {
	AssignPropertiesFrom(src *storage.NamespacesTopic) error
	AssignPropertiesTo(dst *storage.NamespacesTopic) error
}

// Storage version of v1api20210101preview.NamespacesTopic_Spec
type NamespacesTopic_Spec struct {
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                           string                       `json:"azureName,omitempty"`
	DefaultMessageTimeToLive            *string                      `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string                      `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                        `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                        `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                        `json:"enablePartitioning,omitempty"`
	MaxSizeInMegabytes                  *int                         `json:"maxSizeInMegabytes,omitempty"`
	OperatorSpec                        *NamespacesTopicOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion                     string                       `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a servicebus.azure.com/Namespace resource
	Owner                      *genruntime.KnownResourceReference `group:"servicebus.azure.com" json:"owner,omitempty" kind:"Namespace"`
	PropertyBag                genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection *bool                              `json:"requiresDuplicateDetection,omitempty"`
	SupportOrdering            *bool                              `json:"supportOrdering,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NamespacesTopic_Spec{}

// ConvertSpecFrom populates our NamespacesTopic_Spec from the provided source
func (topic *NamespacesTopic_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.NamespacesTopic_Spec)
	if ok {
		// Populate our instance from source
		return topic.AssignProperties_From_NamespacesTopic_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.NamespacesTopic_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = topic.AssignProperties_From_NamespacesTopic_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our NamespacesTopic_Spec
func (topic *NamespacesTopic_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.NamespacesTopic_Spec)
	if ok {
		// Populate destination from our instance
		return topic.AssignProperties_To_NamespacesTopic_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.NamespacesTopic_Spec{}
	err := topic.AssignProperties_To_NamespacesTopic_Spec(dst)
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

// AssignProperties_From_NamespacesTopic_Spec populates our NamespacesTopic_Spec from the provided source NamespacesTopic_Spec
func (topic *NamespacesTopic_Spec) AssignProperties_From_NamespacesTopic_Spec(source *storage.NamespacesTopic_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AutoDeleteOnIdle
	topic.AutoDeleteOnIdle = genruntime.ClonePointerToString(source.AutoDeleteOnIdle)

	// AzureName
	topic.AzureName = source.AzureName

	// DefaultMessageTimeToLive
	topic.DefaultMessageTimeToLive = genruntime.ClonePointerToString(source.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	topic.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(source.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if source.EnableBatchedOperations != nil {
		enableBatchedOperation := *source.EnableBatchedOperations
		topic.EnableBatchedOperations = &enableBatchedOperation
	} else {
		topic.EnableBatchedOperations = nil
	}

	// EnableExpress
	if source.EnableExpress != nil {
		enableExpress := *source.EnableExpress
		topic.EnableExpress = &enableExpress
	} else {
		topic.EnableExpress = nil
	}

	// EnablePartitioning
	if source.EnablePartitioning != nil {
		enablePartitioning := *source.EnablePartitioning
		topic.EnablePartitioning = &enablePartitioning
	} else {
		topic.EnablePartitioning = nil
	}

	// MaxMessageSizeInKilobytes
	if source.MaxMessageSizeInKilobytes != nil {
		propertyBag.Add("MaxMessageSizeInKilobytes", *source.MaxMessageSizeInKilobytes)
	} else {
		propertyBag.Remove("MaxMessageSizeInKilobytes")
	}

	// MaxSizeInMegabytes
	topic.MaxSizeInMegabytes = genruntime.ClonePointerToInt(source.MaxSizeInMegabytes)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec NamespacesTopicOperatorSpec
		err := operatorSpec.AssignProperties_From_NamespacesTopicOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_NamespacesTopicOperatorSpec() to populate field OperatorSpec")
		}
		topic.OperatorSpec = &operatorSpec
	} else {
		topic.OperatorSpec = nil
	}

	// OriginalVersion
	topic.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		topic.Owner = &owner
	} else {
		topic.Owner = nil
	}

	// RequiresDuplicateDetection
	if source.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *source.RequiresDuplicateDetection
		topic.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		topic.RequiresDuplicateDetection = nil
	}

	// SupportOrdering
	if source.SupportOrdering != nil {
		supportOrdering := *source.SupportOrdering
		topic.SupportOrdering = &supportOrdering
	} else {
		topic.SupportOrdering = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		topic.PropertyBag = propertyBag
	} else {
		topic.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespacesTopic_Spec interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForNamespacesTopic_Spec); ok {
		err := augmentedTopic.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_NamespacesTopic_Spec populates the provided destination NamespacesTopic_Spec from our NamespacesTopic_Spec
func (topic *NamespacesTopic_Spec) AssignProperties_To_NamespacesTopic_Spec(destination *storage.NamespacesTopic_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(topic.PropertyBag)

	// AutoDeleteOnIdle
	destination.AutoDeleteOnIdle = genruntime.ClonePointerToString(topic.AutoDeleteOnIdle)

	// AzureName
	destination.AzureName = topic.AzureName

	// DefaultMessageTimeToLive
	destination.DefaultMessageTimeToLive = genruntime.ClonePointerToString(topic.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	destination.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(topic.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if topic.EnableBatchedOperations != nil {
		enableBatchedOperation := *topic.EnableBatchedOperations
		destination.EnableBatchedOperations = &enableBatchedOperation
	} else {
		destination.EnableBatchedOperations = nil
	}

	// EnableExpress
	if topic.EnableExpress != nil {
		enableExpress := *topic.EnableExpress
		destination.EnableExpress = &enableExpress
	} else {
		destination.EnableExpress = nil
	}

	// EnablePartitioning
	if topic.EnablePartitioning != nil {
		enablePartitioning := *topic.EnablePartitioning
		destination.EnablePartitioning = &enablePartitioning
	} else {
		destination.EnablePartitioning = nil
	}

	// MaxMessageSizeInKilobytes
	if propertyBag.Contains("MaxMessageSizeInKilobytes") {
		var maxMessageSizeInKilobyte int
		err := propertyBag.Pull("MaxMessageSizeInKilobytes", &maxMessageSizeInKilobyte)
		if err != nil {
			return eris.Wrap(err, "pulling 'MaxMessageSizeInKilobytes' from propertyBag")
		}

		destination.MaxMessageSizeInKilobytes = &maxMessageSizeInKilobyte
	} else {
		destination.MaxMessageSizeInKilobytes = nil
	}

	// MaxSizeInMegabytes
	destination.MaxSizeInMegabytes = genruntime.ClonePointerToInt(topic.MaxSizeInMegabytes)

	// OperatorSpec
	if topic.OperatorSpec != nil {
		var operatorSpec storage.NamespacesTopicOperatorSpec
		err := topic.OperatorSpec.AssignProperties_To_NamespacesTopicOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_NamespacesTopicOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = topic.OriginalVersion

	// Owner
	if topic.Owner != nil {
		owner := topic.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// RequiresDuplicateDetection
	if topic.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *topic.RequiresDuplicateDetection
		destination.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		destination.RequiresDuplicateDetection = nil
	}

	// SupportOrdering
	if topic.SupportOrdering != nil {
		supportOrdering := *topic.SupportOrdering
		destination.SupportOrdering = &supportOrdering
	} else {
		destination.SupportOrdering = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespacesTopic_Spec interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForNamespacesTopic_Spec); ok {
		err := augmentedTopic.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20210101preview.NamespacesTopic_STATUS
type NamespacesTopic_STATUS struct {
	AccessedAt                          *string                     `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                    *string                     `json:"autoDeleteOnIdle,omitempty"`
	Conditions                          []conditions.Condition      `json:"conditions,omitempty"`
	CountDetails                        *MessageCountDetails_STATUS `json:"countDetails,omitempty"`
	CreatedAt                           *string                     `json:"createdAt,omitempty"`
	DefaultMessageTimeToLive            *string                     `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string                     `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                       `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                       `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                       `json:"enablePartitioning,omitempty"`
	Id                                  *string                     `json:"id,omitempty"`
	MaxSizeInMegabytes                  *int                        `json:"maxSizeInMegabytes,omitempty"`
	Name                                *string                     `json:"name,omitempty"`
	PropertyBag                         genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection          *bool                       `json:"requiresDuplicateDetection,omitempty"`
	SizeInBytes                         *int                        `json:"sizeInBytes,omitempty"`
	Status                              *string                     `json:"status,omitempty"`
	SubscriptionCount                   *int                        `json:"subscriptionCount,omitempty"`
	SupportOrdering                     *bool                       `json:"supportOrdering,omitempty"`
	SystemData                          *SystemData_STATUS          `json:"systemData,omitempty"`
	Type                                *string                     `json:"type,omitempty"`
	UpdatedAt                           *string                     `json:"updatedAt,omitempty"`
}

var _ genruntime.ConvertibleStatus = &NamespacesTopic_STATUS{}

// ConvertStatusFrom populates our NamespacesTopic_STATUS from the provided source
func (topic *NamespacesTopic_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.NamespacesTopic_STATUS)
	if ok {
		// Populate our instance from source
		return topic.AssignProperties_From_NamespacesTopic_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.NamespacesTopic_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = topic.AssignProperties_From_NamespacesTopic_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our NamespacesTopic_STATUS
func (topic *NamespacesTopic_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.NamespacesTopic_STATUS)
	if ok {
		// Populate destination from our instance
		return topic.AssignProperties_To_NamespacesTopic_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.NamespacesTopic_STATUS{}
	err := topic.AssignProperties_To_NamespacesTopic_STATUS(dst)
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

// AssignProperties_From_NamespacesTopic_STATUS populates our NamespacesTopic_STATUS from the provided source NamespacesTopic_STATUS
func (topic *NamespacesTopic_STATUS) AssignProperties_From_NamespacesTopic_STATUS(source *storage.NamespacesTopic_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AccessedAt
	topic.AccessedAt = genruntime.ClonePointerToString(source.AccessedAt)

	// AutoDeleteOnIdle
	topic.AutoDeleteOnIdle = genruntime.ClonePointerToString(source.AutoDeleteOnIdle)

	// Conditions
	topic.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CountDetails
	if source.CountDetails != nil {
		var countDetail MessageCountDetails_STATUS
		err := countDetail.AssignProperties_From_MessageCountDetails_STATUS(source.CountDetails)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_MessageCountDetails_STATUS() to populate field CountDetails")
		}
		topic.CountDetails = &countDetail
	} else {
		topic.CountDetails = nil
	}

	// CreatedAt
	topic.CreatedAt = genruntime.ClonePointerToString(source.CreatedAt)

	// DefaultMessageTimeToLive
	topic.DefaultMessageTimeToLive = genruntime.ClonePointerToString(source.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	topic.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(source.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if source.EnableBatchedOperations != nil {
		enableBatchedOperation := *source.EnableBatchedOperations
		topic.EnableBatchedOperations = &enableBatchedOperation
	} else {
		topic.EnableBatchedOperations = nil
	}

	// EnableExpress
	if source.EnableExpress != nil {
		enableExpress := *source.EnableExpress
		topic.EnableExpress = &enableExpress
	} else {
		topic.EnableExpress = nil
	}

	// EnablePartitioning
	if source.EnablePartitioning != nil {
		enablePartitioning := *source.EnablePartitioning
		topic.EnablePartitioning = &enablePartitioning
	} else {
		topic.EnablePartitioning = nil
	}

	// Id
	topic.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	if source.Location != nil {
		propertyBag.Add("Location", *source.Location)
	} else {
		propertyBag.Remove("Location")
	}

	// MaxMessageSizeInKilobytes
	if source.MaxMessageSizeInKilobytes != nil {
		propertyBag.Add("MaxMessageSizeInKilobytes", *source.MaxMessageSizeInKilobytes)
	} else {
		propertyBag.Remove("MaxMessageSizeInKilobytes")
	}

	// MaxSizeInMegabytes
	topic.MaxSizeInMegabytes = genruntime.ClonePointerToInt(source.MaxSizeInMegabytes)

	// Name
	topic.Name = genruntime.ClonePointerToString(source.Name)

	// RequiresDuplicateDetection
	if source.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *source.RequiresDuplicateDetection
		topic.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		topic.RequiresDuplicateDetection = nil
	}

	// SizeInBytes
	topic.SizeInBytes = genruntime.ClonePointerToInt(source.SizeInBytes)

	// Status
	topic.Status = genruntime.ClonePointerToString(source.Status)

	// SubscriptionCount
	topic.SubscriptionCount = genruntime.ClonePointerToInt(source.SubscriptionCount)

	// SupportOrdering
	if source.SupportOrdering != nil {
		supportOrdering := *source.SupportOrdering
		topic.SupportOrdering = &supportOrdering
	} else {
		topic.SupportOrdering = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		topic.SystemData = &systemDatum
	} else {
		topic.SystemData = nil
	}

	// Type
	topic.Type = genruntime.ClonePointerToString(source.Type)

	// UpdatedAt
	topic.UpdatedAt = genruntime.ClonePointerToString(source.UpdatedAt)

	// Update the property bag
	if len(propertyBag) > 0 {
		topic.PropertyBag = propertyBag
	} else {
		topic.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespacesTopic_STATUS interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForNamespacesTopic_STATUS); ok {
		err := augmentedTopic.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_NamespacesTopic_STATUS populates the provided destination NamespacesTopic_STATUS from our NamespacesTopic_STATUS
func (topic *NamespacesTopic_STATUS) AssignProperties_To_NamespacesTopic_STATUS(destination *storage.NamespacesTopic_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(topic.PropertyBag)

	// AccessedAt
	destination.AccessedAt = genruntime.ClonePointerToString(topic.AccessedAt)

	// AutoDeleteOnIdle
	destination.AutoDeleteOnIdle = genruntime.ClonePointerToString(topic.AutoDeleteOnIdle)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(topic.Conditions)

	// CountDetails
	if topic.CountDetails != nil {
		var countDetail storage.MessageCountDetails_STATUS
		err := topic.CountDetails.AssignProperties_To_MessageCountDetails_STATUS(&countDetail)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_MessageCountDetails_STATUS() to populate field CountDetails")
		}
		destination.CountDetails = &countDetail
	} else {
		destination.CountDetails = nil
	}

	// CreatedAt
	destination.CreatedAt = genruntime.ClonePointerToString(topic.CreatedAt)

	// DefaultMessageTimeToLive
	destination.DefaultMessageTimeToLive = genruntime.ClonePointerToString(topic.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	destination.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(topic.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if topic.EnableBatchedOperations != nil {
		enableBatchedOperation := *topic.EnableBatchedOperations
		destination.EnableBatchedOperations = &enableBatchedOperation
	} else {
		destination.EnableBatchedOperations = nil
	}

	// EnableExpress
	if topic.EnableExpress != nil {
		enableExpress := *topic.EnableExpress
		destination.EnableExpress = &enableExpress
	} else {
		destination.EnableExpress = nil
	}

	// EnablePartitioning
	if topic.EnablePartitioning != nil {
		enablePartitioning := *topic.EnablePartitioning
		destination.EnablePartitioning = &enablePartitioning
	} else {
		destination.EnablePartitioning = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(topic.Id)

	// Location
	if propertyBag.Contains("Location") {
		var location string
		err := propertyBag.Pull("Location", &location)
		if err != nil {
			return eris.Wrap(err, "pulling 'Location' from propertyBag")
		}

		destination.Location = &location
	} else {
		destination.Location = nil
	}

	// MaxMessageSizeInKilobytes
	if propertyBag.Contains("MaxMessageSizeInKilobytes") {
		var maxMessageSizeInKilobyte int
		err := propertyBag.Pull("MaxMessageSizeInKilobytes", &maxMessageSizeInKilobyte)
		if err != nil {
			return eris.Wrap(err, "pulling 'MaxMessageSizeInKilobytes' from propertyBag")
		}

		destination.MaxMessageSizeInKilobytes = &maxMessageSizeInKilobyte
	} else {
		destination.MaxMessageSizeInKilobytes = nil
	}

	// MaxSizeInMegabytes
	destination.MaxSizeInMegabytes = genruntime.ClonePointerToInt(topic.MaxSizeInMegabytes)

	// Name
	destination.Name = genruntime.ClonePointerToString(topic.Name)

	// RequiresDuplicateDetection
	if topic.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *topic.RequiresDuplicateDetection
		destination.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		destination.RequiresDuplicateDetection = nil
	}

	// SizeInBytes
	destination.SizeInBytes = genruntime.ClonePointerToInt(topic.SizeInBytes)

	// Status
	destination.Status = genruntime.ClonePointerToString(topic.Status)

	// SubscriptionCount
	destination.SubscriptionCount = genruntime.ClonePointerToInt(topic.SubscriptionCount)

	// SupportOrdering
	if topic.SupportOrdering != nil {
		supportOrdering := *topic.SupportOrdering
		destination.SupportOrdering = &supportOrdering
	} else {
		destination.SupportOrdering = nil
	}

	// SystemData
	if topic.SystemData != nil {
		var systemDatum storage.SystemData_STATUS
		err := topic.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(topic.Type)

	// UpdatedAt
	destination.UpdatedAt = genruntime.ClonePointerToString(topic.UpdatedAt)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespacesTopic_STATUS interface (if implemented) to customize the conversion
	var topicAsAny any = topic
	if augmentedTopic, ok := topicAsAny.(augmentConversionForNamespacesTopic_STATUS); ok {
		err := augmentedTopic.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForNamespacesTopic_Spec interface {
	AssignPropertiesFrom(src *storage.NamespacesTopic_Spec) error
	AssignPropertiesTo(dst *storage.NamespacesTopic_Spec) error
}

type augmentConversionForNamespacesTopic_STATUS interface {
	AssignPropertiesFrom(src *storage.NamespacesTopic_STATUS) error
	AssignPropertiesTo(dst *storage.NamespacesTopic_STATUS) error
}

// Storage version of v1api20210101preview.NamespacesTopicOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type NamespacesTopicOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_NamespacesTopicOperatorSpec populates our NamespacesTopicOperatorSpec from the provided source NamespacesTopicOperatorSpec
func (operator *NamespacesTopicOperatorSpec) AssignProperties_From_NamespacesTopicOperatorSpec(source *storage.NamespacesTopicOperatorSpec) error {
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

	// Invoke the augmentConversionForNamespacesTopicOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForNamespacesTopicOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_NamespacesTopicOperatorSpec populates the provided destination NamespacesTopicOperatorSpec from our NamespacesTopicOperatorSpec
func (operator *NamespacesTopicOperatorSpec) AssignProperties_To_NamespacesTopicOperatorSpec(destination *storage.NamespacesTopicOperatorSpec) error {
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

	// Invoke the augmentConversionForNamespacesTopicOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForNamespacesTopicOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForNamespacesTopicOperatorSpec interface {
	AssignPropertiesFrom(src *storage.NamespacesTopicOperatorSpec) error
	AssignPropertiesTo(dst *storage.NamespacesTopicOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&NamespacesTopic{}, &NamespacesTopicList{})
}
