// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=servicebus.azure.com,resources=namespacestopicssubscriptionsrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=servicebus.azure.com,resources={namespacestopicssubscriptionsrules/status,namespacestopicssubscriptionsrules/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20211101.NamespacesTopicsSubscriptionsRule
// Generator information:
// - Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/Rules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}
type NamespacesTopicsSubscriptionsRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Namespaces_Topics_Subscriptions_Rule_Spec   `json:"spec,omitempty"`
	Status            Namespaces_Topics_Subscriptions_Rule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesTopicsSubscriptionsRule{}

// GetConditions returns the conditions of the resource
func (rule *NamespacesTopicsSubscriptionsRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *NamespacesTopicsSubscriptionsRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &NamespacesTopicsSubscriptionsRule{}

// AzureName returns the Azure name of the resource
func (rule *NamespacesTopicsSubscriptionsRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule NamespacesTopicsSubscriptionsRule) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (rule *NamespacesTopicsSubscriptionsRule) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (rule *NamespacesTopicsSubscriptionsRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *NamespacesTopicsSubscriptionsRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/topics/subscriptions/rules"
func (rule *NamespacesTopicsSubscriptionsRule) GetType() string {
	return "Microsoft.ServiceBus/namespaces/topics/subscriptions/rules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *NamespacesTopicsSubscriptionsRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Namespaces_Topics_Subscriptions_Rule_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (rule *NamespacesTopicsSubscriptionsRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  rule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (rule *NamespacesTopicsSubscriptionsRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Namespaces_Topics_Subscriptions_Rule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st Namespaces_Topics_Subscriptions_Rule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// Hub marks that this NamespacesTopicsSubscriptionsRule is the hub type for conversion
func (rule *NamespacesTopicsSubscriptionsRule) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *NamespacesTopicsSubscriptionsRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion,
		Kind:    "NamespacesTopicsSubscriptionsRule",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20211101.NamespacesTopicsSubscriptionsRule
// Generator information:
// - Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/Rules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}
type NamespacesTopicsSubscriptionsRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesTopicsSubscriptionsRule `json:"items"`
}

// Storage version of v1api20211101.Namespaces_Topics_Subscriptions_Rule_Spec
type Namespaces_Topics_Subscriptions_Rule_Spec struct {
	Action *Action `json:"action,omitempty"`

	// +kubebuilder:validation:MaxLength=50
	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName         string             `json:"azureName,omitempty"`
	CorrelationFilter *CorrelationFilter `json:"correlationFilter,omitempty"`
	FilterType        *string            `json:"filterType,omitempty"`
	OriginalVersion   string             `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a servicebus.azure.com/NamespacesTopicsSubscription resource
	Owner       *genruntime.KnownResourceReference `group:"servicebus.azure.com" json:"owner,omitempty" kind:"NamespacesTopicsSubscription"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	SqlFilter   *SqlFilter                         `json:"sqlFilter,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Namespaces_Topics_Subscriptions_Rule_Spec{}

// ConvertSpecFrom populates our Namespaces_Topics_Subscriptions_Rule_Spec from the provided source
func (rule *Namespaces_Topics_Subscriptions_Rule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == rule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(rule)
}

// ConvertSpecTo populates the provided destination from our Namespaces_Topics_Subscriptions_Rule_Spec
func (rule *Namespaces_Topics_Subscriptions_Rule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == rule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(rule)
}

// Storage version of v1api20211101.Namespaces_Topics_Subscriptions_Rule_STATUS
type Namespaces_Topics_Subscriptions_Rule_STATUS struct {
	Action            *Action_STATUS            `json:"action,omitempty"`
	Conditions        []conditions.Condition    `json:"conditions,omitempty"`
	CorrelationFilter *CorrelationFilter_STATUS `json:"correlationFilter,omitempty"`
	FilterType        *string                   `json:"filterType,omitempty"`
	Id                *string                   `json:"id,omitempty"`
	Location          *string                   `json:"location,omitempty"`
	Name              *string                   `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	SqlFilter         *SqlFilter_STATUS         `json:"sqlFilter,omitempty"`
	SystemData        *SystemData_STATUS        `json:"systemData,omitempty"`
	Type              *string                   `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Namespaces_Topics_Subscriptions_Rule_STATUS{}

// ConvertStatusFrom populates our Namespaces_Topics_Subscriptions_Rule_STATUS from the provided source
func (rule *Namespaces_Topics_Subscriptions_Rule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == rule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(rule)
}

// ConvertStatusTo populates the provided destination from our Namespaces_Topics_Subscriptions_Rule_STATUS
func (rule *Namespaces_Topics_Subscriptions_Rule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == rule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(rule)
}

// Storage version of v1api20211101.Action
// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter
// expression.
type Action struct {
	CompatibilityLevel    *int                   `json:"compatibilityLevel,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequiresPreprocessing *bool                  `json:"requiresPreprocessing,omitempty"`
	SqlExpression         *string                `json:"sqlExpression,omitempty"`
}

// Storage version of v1api20211101.Action_STATUS
// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter
// expression.
type Action_STATUS struct {
	CompatibilityLevel    *int                   `json:"compatibilityLevel,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequiresPreprocessing *bool                  `json:"requiresPreprocessing,omitempty"`
	SqlExpression         *string                `json:"sqlExpression,omitempty"`
}

// Storage version of v1api20211101.CorrelationFilter
// Represents the correlation filter expression.
type CorrelationFilter struct {
	ContentType           *string                `json:"contentType,omitempty"`
	CorrelationId         *string                `json:"correlationId,omitempty"`
	Label                 *string                `json:"label,omitempty"`
	MessageId             *string                `json:"messageId,omitempty"`
	Properties            map[string]string      `json:"properties,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ReplyTo               *string                `json:"replyTo,omitempty"`
	ReplyToSessionId      *string                `json:"replyToSessionId,omitempty"`
	RequiresPreprocessing *bool                  `json:"requiresPreprocessing,omitempty"`
	SessionId             *string                `json:"sessionId,omitempty"`
	To                    *string                `json:"to,omitempty"`
}

// Storage version of v1api20211101.CorrelationFilter_STATUS
// Represents the correlation filter expression.
type CorrelationFilter_STATUS struct {
	ContentType           *string                `json:"contentType,omitempty"`
	CorrelationId         *string                `json:"correlationId,omitempty"`
	Label                 *string                `json:"label,omitempty"`
	MessageId             *string                `json:"messageId,omitempty"`
	Properties            map[string]string      `json:"properties,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ReplyTo               *string                `json:"replyTo,omitempty"`
	ReplyToSessionId      *string                `json:"replyToSessionId,omitempty"`
	RequiresPreprocessing *bool                  `json:"requiresPreprocessing,omitempty"`
	SessionId             *string                `json:"sessionId,omitempty"`
	To                    *string                `json:"to,omitempty"`
}

// Storage version of v1api20211101.SqlFilter
// Represents a filter which is a composition of an expression and an action that is executed in the pub/sub pipeline.
type SqlFilter struct {
	CompatibilityLevel    *int                   `json:"compatibilityLevel,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequiresPreprocessing *bool                  `json:"requiresPreprocessing,omitempty"`
	SqlExpression         *string                `json:"sqlExpression,omitempty"`
}

// Storage version of v1api20211101.SqlFilter_STATUS
// Represents a filter which is a composition of an expression and an action that is executed in the pub/sub pipeline.
type SqlFilter_STATUS struct {
	CompatibilityLevel    *int                   `json:"compatibilityLevel,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequiresPreprocessing *bool                  `json:"requiresPreprocessing,omitempty"`
	SqlExpression         *string                `json:"sqlExpression,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NamespacesTopicsSubscriptionsRule{}, &NamespacesTopicsSubscriptionsRuleList{})
}
