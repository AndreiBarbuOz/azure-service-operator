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

// +kubebuilder:rbac:groups=apimanagement.azure.com,resources=products,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apimanagement.azure.com,resources={products/status,products/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220801.Product
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimproducts.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}
type Product struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Product_Spec   `json:"spec,omitempty"`
	Status            Product_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Product{}

// GetConditions returns the conditions of the resource
func (product *Product) GetConditions() conditions.Conditions {
	return product.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (product *Product) SetConditions(conditions conditions.Conditions) {
	product.Status.Conditions = conditions
}

var _ configmaps.Exporter = &Product{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (product *Product) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if product.Spec.OperatorSpec == nil {
		return nil
	}
	return product.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &Product{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (product *Product) SecretDestinationExpressions() []*core.DestinationExpression {
	if product.Spec.OperatorSpec == nil {
		return nil
	}
	return product.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &Product{}

// AzureName returns the Azure name of the resource
func (product *Product) AzureName() string {
	return product.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (product Product) GetAPIVersion() string {
	return "2022-08-01"
}

// GetResourceScope returns the scope of the resource
func (product *Product) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (product *Product) GetSpec() genruntime.ConvertibleSpec {
	return &product.Spec
}

// GetStatus returns the status of this resource
func (product *Product) GetStatus() genruntime.ConvertibleStatus {
	return &product.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (product *Product) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationHead,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/products"
func (product *Product) GetType() string {
	return "Microsoft.ApiManagement/service/products"
}

// NewEmptyStatus returns a new empty (blank) status
func (product *Product) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Product_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (product *Product) Owner() *genruntime.ResourceReference {
	if product.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(product.Spec)
	return product.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (product *Product) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Product_STATUS); ok {
		product.Status = *st
		return nil
	}

	// Convert status to required version
	var st Product_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	product.Status = st
	return nil
}

// Hub marks that this Product is the hub type for conversion
func (product *Product) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (product *Product) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: product.Spec.OriginalVersion,
		Kind:    "Product",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220801.Product
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimproducts.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}
type ProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Product `json:"items"`
}

// Storage version of v1api20220801.Product_Spec
type Product_Spec struct {
	ApprovalRequired *bool `json:"approvalRequired,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string               `json:"azureName,omitempty"`
	Description     *string              `json:"description,omitempty"`
	DisplayName     *string              `json:"displayName,omitempty"`
	OperatorSpec    *ProductOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner                *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`
	PropertyBag          genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	State                *string                            `json:"state,omitempty"`
	SubscriptionRequired *bool                              `json:"subscriptionRequired,omitempty"`
	SubscriptionsLimit   *int                               `json:"subscriptionsLimit,omitempty"`
	Terms                *string                            `json:"terms,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Product_Spec{}

// ConvertSpecFrom populates our Product_Spec from the provided source
func (product *Product_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == product {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(product)
}

// ConvertSpecTo populates the provided destination from our Product_Spec
func (product *Product_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == product {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(product)
}

// Storage version of v1api20220801.Product_STATUS
type Product_STATUS struct {
	ApprovalRequired     *bool                  `json:"approvalRequired,omitempty"`
	Conditions           []conditions.Condition `json:"conditions,omitempty"`
	Description          *string                `json:"description,omitempty"`
	DisplayName          *string                `json:"displayName,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	State                *string                `json:"state,omitempty"`
	SubscriptionRequired *bool                  `json:"subscriptionRequired,omitempty"`
	SubscriptionsLimit   *int                   `json:"subscriptionsLimit,omitempty"`
	Terms                *string                `json:"terms,omitempty"`
	Type                 *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Product_STATUS{}

// ConvertStatusFrom populates our Product_STATUS from the provided source
func (product *Product_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == product {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(product)
}

// ConvertStatusTo populates the provided destination from our Product_STATUS
func (product *Product_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == product {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(product)
}

// Storage version of v1api20220801.ProductOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ProductOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Product{}, &ProductList{})
}
