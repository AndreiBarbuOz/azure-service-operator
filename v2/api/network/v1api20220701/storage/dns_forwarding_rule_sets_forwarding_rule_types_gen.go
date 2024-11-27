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

// +kubebuilder:rbac:groups=network.azure.com,resources=dnsforwardingrulesetsforwardingrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={dnsforwardingrulesetsforwardingrules/status,dnsforwardingrulesetsforwardingrules/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220701.DnsForwardingRuleSetsForwardingRule
// Generator information:
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/forwardingRules/{forwardingRuleName}
type DnsForwardingRuleSetsForwardingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsForwardingRuleSetsForwardingRule_Spec   `json:"spec,omitempty"`
	Status            DnsForwardingRuleSetsForwardingRule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DnsForwardingRuleSetsForwardingRule{}

// GetConditions returns the conditions of the resource
func (rule *DnsForwardingRuleSetsForwardingRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *DnsForwardingRuleSetsForwardingRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ configmaps.Exporter = &DnsForwardingRuleSetsForwardingRule{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (rule *DnsForwardingRuleSetsForwardingRule) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if rule.Spec.OperatorSpec == nil {
		return nil
	}
	return rule.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &DnsForwardingRuleSetsForwardingRule{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (rule *DnsForwardingRuleSetsForwardingRule) SecretDestinationExpressions() []*core.DestinationExpression {
	if rule.Spec.OperatorSpec == nil {
		return nil
	}
	return rule.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &DnsForwardingRuleSetsForwardingRule{}

// AzureName returns the Azure name of the resource
func (rule *DnsForwardingRuleSetsForwardingRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (rule DnsForwardingRuleSetsForwardingRule) GetAPIVersion() string {
	return "2022-07-01"
}

// GetResourceScope returns the scope of the resource
func (rule *DnsForwardingRuleSetsForwardingRule) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (rule *DnsForwardingRuleSetsForwardingRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *DnsForwardingRuleSetsForwardingRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (rule *DnsForwardingRuleSetsForwardingRule) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsForwardingRulesets/forwardingRules"
func (rule *DnsForwardingRuleSetsForwardingRule) GetType() string {
	return "Microsoft.Network/dnsForwardingRulesets/forwardingRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *DnsForwardingRuleSetsForwardingRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DnsForwardingRuleSetsForwardingRule_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (rule *DnsForwardingRuleSetsForwardingRule) Owner() *genruntime.ResourceReference {
	if rule.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return rule.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (rule *DnsForwardingRuleSetsForwardingRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DnsForwardingRuleSetsForwardingRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st DnsForwardingRuleSetsForwardingRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// Hub marks that this DnsForwardingRuleSetsForwardingRule is the hub type for conversion
func (rule *DnsForwardingRuleSetsForwardingRule) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *DnsForwardingRuleSetsForwardingRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion,
		Kind:    "DnsForwardingRuleSetsForwardingRule",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220701.DnsForwardingRuleSetsForwardingRule
// Generator information:
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/forwardingRules/{forwardingRuleName}
type DnsForwardingRuleSetsForwardingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DnsForwardingRuleSetsForwardingRule `json:"items"`
}

// Storage version of v1api20220701.DnsForwardingRuleSetsForwardingRule_Spec
type DnsForwardingRuleSetsForwardingRule_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName           string                                           `json:"azureName,omitempty"`
	DomainName          *string                                          `json:"domainName,omitempty"`
	ForwardingRuleState *string                                          `json:"forwardingRuleState,omitempty"`
	Metadata            map[string]string                                `json:"metadata,omitempty"`
	OperatorSpec        *DnsForwardingRuleSetsForwardingRuleOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion     string                                           `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/DnsForwardingRuleset resource
	Owner            *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"DnsForwardingRuleset"`
	PropertyBag      genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	TargetDnsServers []TargetDnsServer                  `json:"targetDnsServers,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DnsForwardingRuleSetsForwardingRule_Spec{}

// ConvertSpecFrom populates our DnsForwardingRuleSetsForwardingRule_Spec from the provided source
func (rule *DnsForwardingRuleSetsForwardingRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == rule {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(rule)
}

// ConvertSpecTo populates the provided destination from our DnsForwardingRuleSetsForwardingRule_Spec
func (rule *DnsForwardingRuleSetsForwardingRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == rule {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(rule)
}

// Storage version of v1api20220701.DnsForwardingRuleSetsForwardingRule_STATUS
type DnsForwardingRuleSetsForwardingRule_STATUS struct {
	Conditions          []conditions.Condition   `json:"conditions,omitempty"`
	DomainName          *string                  `json:"domainName,omitempty"`
	Etag                *string                  `json:"etag,omitempty"`
	ForwardingRuleState *string                  `json:"forwardingRuleState,omitempty"`
	Id                  *string                  `json:"id,omitempty"`
	Metadata            map[string]string        `json:"metadata,omitempty"`
	Name                *string                  `json:"name,omitempty"`
	PropertyBag         genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
	ProvisioningState   *string                  `json:"provisioningState,omitempty"`
	SystemData          *SystemData_STATUS       `json:"systemData,omitempty"`
	TargetDnsServers    []TargetDnsServer_STATUS `json:"targetDnsServers,omitempty"`
	Type                *string                  `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DnsForwardingRuleSetsForwardingRule_STATUS{}

// ConvertStatusFrom populates our DnsForwardingRuleSetsForwardingRule_STATUS from the provided source
func (rule *DnsForwardingRuleSetsForwardingRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == rule {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(rule)
}

// ConvertStatusTo populates the provided destination from our DnsForwardingRuleSetsForwardingRule_STATUS
func (rule *DnsForwardingRuleSetsForwardingRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == rule {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(rule)
}

// Storage version of v1api20220701.DnsForwardingRuleSetsForwardingRuleOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type DnsForwardingRuleSetsForwardingRuleOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20220701.SystemData_STATUS
// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220701.TargetDnsServer
// Describes a server to forward the DNS queries to.
type TargetDnsServer struct {
	IpAddress           *string                        `json:"ipAddress,omitempty" optionalConfigMapPair:"IpAddress"`
	IpAddressFromConfig *genruntime.ConfigMapReference `json:"ipAddressFromConfig,omitempty" optionalConfigMapPair:"IpAddress"`
	Port                *int                           `json:"port,omitempty"`
	PropertyBag         genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220701.TargetDnsServer_STATUS
// Describes a server to forward the DNS queries to.
type TargetDnsServer_STATUS struct {
	IpAddress   *string                `json:"ipAddress,omitempty"`
	Port        *int                   `json:"port,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DnsForwardingRuleSetsForwardingRule{}, &DnsForwardingRuleSetsForwardingRuleList{})
}
