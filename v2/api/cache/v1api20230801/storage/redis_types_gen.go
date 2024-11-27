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

// +kubebuilder:rbac:groups=cache.azure.com,resources=redis,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.azure.com,resources={redis/status,redis/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230801.Redis
// Generator information:
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2023-08-01/redis.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}
type Redis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Redis_Spec   `json:"spec,omitempty"`
	Status            Redis_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Redis{}

// GetConditions returns the conditions of the resource
func (redis *Redis) GetConditions() conditions.Conditions {
	return redis.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (redis *Redis) SetConditions(conditions conditions.Conditions) {
	redis.Status.Conditions = conditions
}

var _ configmaps.Exporter = &Redis{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (redis *Redis) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if redis.Spec.OperatorSpec == nil {
		return nil
	}
	return redis.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &Redis{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (redis *Redis) SecretDestinationExpressions() []*core.DestinationExpression {
	if redis.Spec.OperatorSpec == nil {
		return nil
	}
	return redis.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &Redis{}

// AzureName returns the Azure name of the resource
func (redis *Redis) AzureName() string {
	return redis.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-08-01"
func (redis Redis) GetAPIVersion() string {
	return "2023-08-01"
}

// GetResourceScope returns the scope of the resource
func (redis *Redis) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (redis *Redis) GetSpec() genruntime.ConvertibleSpec {
	return &redis.Spec
}

// GetStatus returns the status of this resource
func (redis *Redis) GetStatus() genruntime.ConvertibleStatus {
	return &redis.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (redis *Redis) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis"
func (redis *Redis) GetType() string {
	return "Microsoft.Cache/redis"
}

// NewEmptyStatus returns a new empty (blank) status
func (redis *Redis) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Redis_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (redis *Redis) Owner() *genruntime.ResourceReference {
	if redis.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(redis.Spec)
	return redis.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (redis *Redis) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Redis_STATUS); ok {
		redis.Status = *st
		return nil
	}

	// Convert status to required version
	var st Redis_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	redis.Status = st
	return nil
}

// Hub marks that this Redis is the hub type for conversion
func (redis *Redis) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (redis *Redis) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: redis.Spec.OriginalVersion,
		Kind:    "Redis",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230801.Redis
// Generator information:
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2023-08-01/redis.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}
type RedisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Redis `json:"items"`
}

// Storage version of v1api20230801.APIVersion
// +kubebuilder:validation:Enum={"2023-08-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2023-08-01")

// Storage version of v1api20230801.Redis_Spec
type Redis_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName         string                  `json:"azureName,omitempty"`
	EnableNonSslPort  *bool                   `json:"enableNonSslPort,omitempty"`
	Identity          *ManagedServiceIdentity `json:"identity,omitempty"`
	Location          *string                 `json:"location,omitempty"`
	MinimumTlsVersion *string                 `json:"minimumTlsVersion,omitempty"`
	OperatorSpec      *RedisOperatorSpec      `json:"operatorSpec,omitempty"`
	OriginalVersion   string                  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner               *genruntime.KnownResourceReference        `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag         genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                                   `json:"publicNetworkAccess,omitempty"`
	RedisConfiguration  *RedisCreateProperties_RedisConfiguration `json:"redisConfiguration,omitempty"`
	RedisVersion        *string                                   `json:"redisVersion,omitempty"`
	ReplicasPerMaster   *int                                      `json:"replicasPerMaster,omitempty"`
	ReplicasPerPrimary  *int                                      `json:"replicasPerPrimary,omitempty"`
	ShardCount          *int                                      `json:"shardCount,omitempty"`
	Sku                 *Sku                                      `json:"sku,omitempty"`
	StaticIP            *string                                   `json:"staticIP,omitempty"`

	// SubnetReference: The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
	SubnetReference *genruntime.ResourceReference `armReference:"SubnetId" json:"subnetReference,omitempty"`
	Tags            map[string]string             `json:"tags,omitempty"`
	TenantSettings  map[string]string             `json:"tenantSettings,omitempty"`
	UpdateChannel   *string                       `json:"updateChannel,omitempty"`
	Zones           []string                      `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Redis_Spec{}

// ConvertSpecFrom populates our Redis_Spec from the provided source
func (redis *Redis_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == redis {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(redis)
}

// ConvertSpecTo populates the provided destination from our Redis_Spec
func (redis *Redis_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == redis {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(redis)
}

// Storage version of v1api20230801.Redis_STATUS
type Redis_STATUS struct {
	Conditions                 []conditions.Condition                     `json:"conditions,omitempty"`
	EnableNonSslPort           *bool                                      `json:"enableNonSslPort,omitempty"`
	HostName                   *string                                    `json:"hostName,omitempty"`
	Id                         *string                                    `json:"id,omitempty"`
	Identity                   *ManagedServiceIdentity_STATUS             `json:"identity,omitempty"`
	Instances                  []RedisInstanceDetails_STATUS              `json:"instances,omitempty"`
	LinkedServers              []RedisLinkedServer_STATUS                 `json:"linkedServers,omitempty"`
	Location                   *string                                    `json:"location,omitempty"`
	MinimumTlsVersion          *string                                    `json:"minimumTlsVersion,omitempty"`
	Name                       *string                                    `json:"name,omitempty"`
	Port                       *int                                       `json:"port,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS         `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                     `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                                    `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *string                                    `json:"publicNetworkAccess,omitempty"`
	RedisConfiguration         *RedisProperties_RedisConfiguration_STATUS `json:"redisConfiguration,omitempty"`
	RedisVersion               *string                                    `json:"redisVersion,omitempty"`
	ReplicasPerMaster          *int                                       `json:"replicasPerMaster,omitempty"`
	ReplicasPerPrimary         *int                                       `json:"replicasPerPrimary,omitempty"`
	ShardCount                 *int                                       `json:"shardCount,omitempty"`
	Sku                        *Sku_STATUS                                `json:"sku,omitempty"`
	SslPort                    *int                                       `json:"sslPort,omitempty"`
	StaticIP                   *string                                    `json:"staticIP,omitempty"`
	SubnetId                   *string                                    `json:"subnetId,omitempty"`
	Tags                       map[string]string                          `json:"tags,omitempty"`
	TenantSettings             map[string]string                          `json:"tenantSettings,omitempty"`
	Type                       *string                                    `json:"type,omitempty"`
	UpdateChannel              *string                                    `json:"updateChannel,omitempty"`
	Zones                      []string                                   `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Redis_STATUS{}

// ConvertStatusFrom populates our Redis_STATUS from the provided source
func (redis *Redis_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == redis {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(redis)
}

// ConvertStatusTo populates the provided destination from our Redis_STATUS
func (redis *Redis_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == redis {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(redis)
}

// Storage version of v1api20230801.ManagedServiceIdentity
// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	PropertyBag            genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Type                   *string                       `json:"type,omitempty"`
	UserAssignedIdentities []UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20230801.ManagedServiceIdentity_STATUS
// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity_STATUS struct {
	PrincipalId            *string                                `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	TenantId               *string                                `json:"tenantId,omitempty"`
	Type                   *string                                `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20230801.PrivateEndpointConnection_STATUS
// The Private Endpoint Connection resource.
type PrivateEndpointConnection_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230801.RedisCreateProperties_RedisConfiguration
type RedisCreateProperties_RedisConfiguration struct {
	AadEnabled                         *string                `json:"aad-enabled,omitempty"`
	AofBackupEnabled                   *string                `json:"aof-backup-enabled,omitempty"`
	AofStorageConnectionString0        *string                `json:"aof-storage-connection-string-0,omitempty"`
	AofStorageConnectionString1        *string                `json:"aof-storage-connection-string-1,omitempty"`
	Authnotrequired                    *string                `json:"authnotrequired,omitempty"`
	MaxfragmentationmemoryReserved     *string                `json:"maxfragmentationmemory-reserved,omitempty"`
	MaxmemoryDelta                     *string                `json:"maxmemory-delta,omitempty"`
	MaxmemoryPolicy                    *string                `json:"maxmemory-policy,omitempty"`
	MaxmemoryReserved                  *string                `json:"maxmemory-reserved,omitempty"`
	NotifyKeyspaceEvents               *string                `json:"notify-keyspace-events,omitempty"`
	PreferredDataPersistenceAuthMethod *string                `json:"preferred-data-persistence-auth-method,omitempty"`
	PropertyBag                        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RdbBackupEnabled                   *string                `json:"rdb-backup-enabled,omitempty"`
	RdbBackupFrequency                 *string                `json:"rdb-backup-frequency,omitempty"`
	RdbBackupMaxSnapshotCount          *string                `json:"rdb-backup-max-snapshot-count,omitempty"`
	RdbStorageConnectionString         *string                `json:"rdb-storage-connection-string,omitempty"`
	StorageSubscriptionId              *string                `json:"storage-subscription-id,omitempty"`
}

// Storage version of v1api20230801.RedisInstanceDetails_STATUS
// Details of single instance of redis.
type RedisInstanceDetails_STATUS struct {
	IsMaster    *bool                  `json:"isMaster,omitempty"`
	IsPrimary   *bool                  `json:"isPrimary,omitempty"`
	NonSslPort  *int                   `json:"nonSslPort,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ShardId     *int                   `json:"shardId,omitempty"`
	SslPort     *int                   `json:"sslPort,omitempty"`
	Zone        *string                `json:"zone,omitempty"`
}

// Storage version of v1api20230801.RedisLinkedServer_STATUS
// Linked server Id
type RedisLinkedServer_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230801.RedisOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type RedisOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
	Secrets              *RedisOperatorSecrets         `json:"secrets,omitempty"`
}

// Storage version of v1api20230801.RedisProperties_RedisConfiguration_STATUS
type RedisProperties_RedisConfiguration_STATUS struct {
	AadEnabled                         *string                `json:"aad-enabled,omitempty"`
	AofBackupEnabled                   *string                `json:"aof-backup-enabled,omitempty"`
	AofStorageConnectionString0        *string                `json:"aof-storage-connection-string-0,omitempty"`
	AofStorageConnectionString1        *string                `json:"aof-storage-connection-string-1,omitempty"`
	Authnotrequired                    *string                `json:"authnotrequired,omitempty"`
	Maxclients                         *string                `json:"maxclients,omitempty"`
	MaxfragmentationmemoryReserved     *string                `json:"maxfragmentationmemory-reserved,omitempty"`
	MaxmemoryDelta                     *string                `json:"maxmemory-delta,omitempty"`
	MaxmemoryPolicy                    *string                `json:"maxmemory-policy,omitempty"`
	MaxmemoryReserved                  *string                `json:"maxmemory-reserved,omitempty"`
	NotifyKeyspaceEvents               *string                `json:"notify-keyspace-events,omitempty"`
	PreferredDataArchiveAuthMethod     *string                `json:"preferred-data-archive-auth-method,omitempty"`
	PreferredDataPersistenceAuthMethod *string                `json:"preferred-data-persistence-auth-method,omitempty"`
	PropertyBag                        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RdbBackupEnabled                   *string                `json:"rdb-backup-enabled,omitempty"`
	RdbBackupFrequency                 *string                `json:"rdb-backup-frequency,omitempty"`
	RdbBackupMaxSnapshotCount          *string                `json:"rdb-backup-max-snapshot-count,omitempty"`
	RdbStorageConnectionString         *string                `json:"rdb-storage-connection-string,omitempty"`
	StorageSubscriptionId              *string                `json:"storage-subscription-id,omitempty"`
	ZonalConfiguration                 *string                `json:"zonal-configuration,omitempty"`
}

// Storage version of v1api20230801.Sku
// SKU parameters supplied to the create Redis operation.
type Sku struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230801.Sku_STATUS
// SKU parameters supplied to the create Redis operation.
type Sku_STATUS struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230801.RedisOperatorSecrets
type RedisOperatorSecrets struct {
	HostName     *genruntime.SecretDestination `json:"hostName,omitempty"`
	Port         *genruntime.SecretDestination `json:"port,omitempty"`
	PrimaryKey   *genruntime.SecretDestination `json:"primaryKey,omitempty"`
	PropertyBag  genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SSLPort      *genruntime.SecretDestination `json:"sslPort,omitempty"`
	SecondaryKey *genruntime.SecretDestination `json:"secondaryKey,omitempty"`
}

// Storage version of v1api20230801.UserAssignedIdentity_STATUS
// User assigned identity properties
type UserAssignedIdentity_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230801.UserAssignedIdentityDetails
// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Reference   genruntime.ResourceReference `armReference:"Reference" json:"reference,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Redis{}, &RedisList{})
}
