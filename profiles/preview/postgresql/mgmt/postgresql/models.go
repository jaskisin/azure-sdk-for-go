//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package postgresql

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CreateMode = original.CreateMode

const (
	CreateModeDefault                   CreateMode = original.CreateModeDefault
	CreateModeGeoRestore                CreateMode = original.CreateModeGeoRestore
	CreateModePointInTimeRestore        CreateMode = original.CreateModePointInTimeRestore
	CreateModeReplica                   CreateMode = original.CreateModeReplica
	CreateModeServerPropertiesForCreate CreateMode = original.CreateModeServerPropertiesForCreate
)

type GeoRedundantBackup = original.GeoRedundantBackup

const (
	Disabled GeoRedundantBackup = original.Disabled
	Enabled  GeoRedundantBackup = original.Enabled
)

type IdentityType = original.IdentityType

const (
	SystemAssigned IdentityType = original.SystemAssigned
)

type InfrastructureEncryption = original.InfrastructureEncryption

const (
	InfrastructureEncryptionDisabled InfrastructureEncryption = original.InfrastructureEncryptionDisabled
	InfrastructureEncryptionEnabled  InfrastructureEncryption = original.InfrastructureEncryptionEnabled
)

type MinimalTLSVersionEnum = original.MinimalTLSVersionEnum

const (
	TLS10                  MinimalTLSVersionEnum = original.TLS10
	TLS11                  MinimalTLSVersionEnum = original.TLS11
	TLS12                  MinimalTLSVersionEnum = original.TLS12
	TLSEnforcementDisabled MinimalTLSVersionEnum = original.TLSEnforcementDisabled
)

type OperationOrigin = original.OperationOrigin

const (
	NotSpecified OperationOrigin = original.NotSpecified
	System       OperationOrigin = original.System
	User         OperationOrigin = original.User
)

type PrivateEndpointProvisioningState = original.PrivateEndpointProvisioningState

const (
	Approving PrivateEndpointProvisioningState = original.Approving
	Dropping  PrivateEndpointProvisioningState = original.Dropping
	Failed    PrivateEndpointProvisioningState = original.Failed
	Ready     PrivateEndpointProvisioningState = original.Ready
	Rejecting PrivateEndpointProvisioningState = original.Rejecting
)

type PrivateLinkServiceConnectionStateActionsRequire = original.PrivateLinkServiceConnectionStateActionsRequire

const (
	None PrivateLinkServiceConnectionStateActionsRequire = original.None
)

type PrivateLinkServiceConnectionStateStatus = original.PrivateLinkServiceConnectionStateStatus

const (
	Approved     PrivateLinkServiceConnectionStateStatus = original.Approved
	Disconnected PrivateLinkServiceConnectionStateStatus = original.Disconnected
	Pending      PrivateLinkServiceConnectionStateStatus = original.Pending
	Rejected     PrivateLinkServiceConnectionStateStatus = original.Rejected
)

type PublicNetworkAccessEnum = original.PublicNetworkAccessEnum

const (
	PublicNetworkAccessEnumDisabled PublicNetworkAccessEnum = original.PublicNetworkAccessEnumDisabled
	PublicNetworkAccessEnumEnabled  PublicNetworkAccessEnum = original.PublicNetworkAccessEnumEnabled
)

type ServerSecurityAlertPolicyState = original.ServerSecurityAlertPolicyState

const (
	ServerSecurityAlertPolicyStateDisabled ServerSecurityAlertPolicyState = original.ServerSecurityAlertPolicyStateDisabled
	ServerSecurityAlertPolicyStateEnabled  ServerSecurityAlertPolicyState = original.ServerSecurityAlertPolicyStateEnabled
)

type ServerState = original.ServerState

const (
	ServerStateDisabled     ServerState = original.ServerStateDisabled
	ServerStateDropping     ServerState = original.ServerStateDropping
	ServerStateInaccessible ServerState = original.ServerStateInaccessible
	ServerStateReady        ServerState = original.ServerStateReady
)

type ServerVersion = original.ServerVersion

const (
	NineFullStopFive    ServerVersion = original.NineFullStopFive
	NineFullStopSix     ServerVersion = original.NineFullStopSix
	OneOne              ServerVersion = original.OneOne
	OneZero             ServerVersion = original.OneZero
	OneZeroFullStopTwo  ServerVersion = original.OneZeroFullStopTwo
	OneZeroFullStopZero ServerVersion = original.OneZeroFullStopZero
)

type SkuTier = original.SkuTier

const (
	Basic           SkuTier = original.Basic
	GeneralPurpose  SkuTier = original.GeneralPurpose
	MemoryOptimized SkuTier = original.MemoryOptimized
)

type SslEnforcementEnum = original.SslEnforcementEnum

const (
	SslEnforcementEnumDisabled SslEnforcementEnum = original.SslEnforcementEnumDisabled
	SslEnforcementEnumEnabled  SslEnforcementEnum = original.SslEnforcementEnumEnabled
)

type StorageAutogrow = original.StorageAutogrow

const (
	StorageAutogrowDisabled StorageAutogrow = original.StorageAutogrowDisabled
	StorageAutogrowEnabled  StorageAutogrow = original.StorageAutogrowEnabled
)

type VirtualNetworkRuleState = original.VirtualNetworkRuleState

const (
	VirtualNetworkRuleStateDeleting     VirtualNetworkRuleState = original.VirtualNetworkRuleStateDeleting
	VirtualNetworkRuleStateInitializing VirtualNetworkRuleState = original.VirtualNetworkRuleStateInitializing
	VirtualNetworkRuleStateInProgress   VirtualNetworkRuleState = original.VirtualNetworkRuleStateInProgress
	VirtualNetworkRuleStateReady        VirtualNetworkRuleState = original.VirtualNetworkRuleStateReady
	VirtualNetworkRuleStateUnknown      VirtualNetworkRuleState = original.VirtualNetworkRuleStateUnknown
)

type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type BasicServerPropertiesForCreate = original.BasicServerPropertiesForCreate
type CheckNameAvailabilityClient = original.CheckNameAvailabilityClient
type CloudError = original.CloudError
type Configuration = original.Configuration
type ConfigurationListResult = original.ConfigurationListResult
type ConfigurationProperties = original.ConfigurationProperties
type ConfigurationsClient = original.ConfigurationsClient
type ConfigurationsCreateOrUpdateFuture = original.ConfigurationsCreateOrUpdateFuture
type Database = original.Database
type DatabaseListResult = original.DatabaseListResult
type DatabaseProperties = original.DatabaseProperties
type DatabasesClient = original.DatabasesClient
type DatabasesCreateOrUpdateFuture = original.DatabasesCreateOrUpdateFuture
type DatabasesDeleteFuture = original.DatabasesDeleteFuture
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorResponse = original.ErrorResponse
type FirewallRule = original.FirewallRule
type FirewallRuleListResult = original.FirewallRuleListResult
type FirewallRuleProperties = original.FirewallRuleProperties
type FirewallRulesClient = original.FirewallRulesClient
type FirewallRulesCreateOrUpdateFuture = original.FirewallRulesCreateOrUpdateFuture
type FirewallRulesDeleteFuture = original.FirewallRulesDeleteFuture
type LocationBasedPerformanceTierClient = original.LocationBasedPerformanceTierClient
type LogFile = original.LogFile
type LogFileListResult = original.LogFileListResult
type LogFileProperties = original.LogFileProperties
type LogFilesClient = original.LogFilesClient
type NameAvailability = original.NameAvailability
type NameAvailabilityRequest = original.NameAvailabilityRequest
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type PerformanceTierListResult = original.PerformanceTierListResult
type PerformanceTierProperties = original.PerformanceTierProperties
type PerformanceTierServiceLevelObjectives = original.PerformanceTierServiceLevelObjectives
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionListResultIterator = original.PrivateEndpointConnectionListResultIterator
type PrivateEndpointConnectionListResultPage = original.PrivateEndpointConnectionListResultPage
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsCreateOrUpdateFuture = original.PrivateEndpointConnectionsCreateOrUpdateFuture
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateEndpointConnectionsUpdateTagsFuture = original.PrivateEndpointConnectionsUpdateTagsFuture
type PrivateEndpointProperty = original.PrivateEndpointProperty
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceListResultIterator = original.PrivateLinkResourceListResultIterator
type PrivateLinkResourceListResultPage = original.PrivateLinkResourceListResultPage
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionStateProperty = original.PrivateLinkServiceConnectionStateProperty
type ProxyResource = original.ProxyResource
type RecoverableServerProperties = original.RecoverableServerProperties
type RecoverableServerResource = original.RecoverableServerResource
type RecoverableServersClient = original.RecoverableServersClient
type ReplicasClient = original.ReplicasClient
type Resource = original.Resource
type ResourceIdentity = original.ResourceIdentity
type SecurityAlertPolicyProperties = original.SecurityAlertPolicyProperties
type Server = original.Server
type ServerAdministratorProperties = original.ServerAdministratorProperties
type ServerAdministratorResource = original.ServerAdministratorResource
type ServerAdministratorResourceListResult = original.ServerAdministratorResourceListResult
type ServerAdministratorsClient = original.ServerAdministratorsClient
type ServerAdministratorsCreateOrUpdateFuture = original.ServerAdministratorsCreateOrUpdateFuture
type ServerAdministratorsDeleteFuture = original.ServerAdministratorsDeleteFuture
type ServerBasedPerformanceTierClient = original.ServerBasedPerformanceTierClient
type ServerForCreate = original.ServerForCreate
type ServerKey = original.ServerKey
type ServerKeyListResult = original.ServerKeyListResult
type ServerKeyListResultIterator = original.ServerKeyListResultIterator
type ServerKeyListResultPage = original.ServerKeyListResultPage
type ServerKeyProperties = original.ServerKeyProperties
type ServerKeysClient = original.ServerKeysClient
type ServerKeysCreateOrUpdateFuture = original.ServerKeysCreateOrUpdateFuture
type ServerKeysDeleteFuture = original.ServerKeysDeleteFuture
type ServerListResult = original.ServerListResult
type ServerParametersClient = original.ServerParametersClient
type ServerParametersListUpdateConfigurationsFuture = original.ServerParametersListUpdateConfigurationsFuture
type ServerPrivateEndpointConnection = original.ServerPrivateEndpointConnection
type ServerPrivateEndpointConnectionProperties = original.ServerPrivateEndpointConnectionProperties
type ServerPrivateLinkServiceConnectionStateProperty = original.ServerPrivateLinkServiceConnectionStateProperty
type ServerProperties = original.ServerProperties
type ServerPropertiesForCreate = original.ServerPropertiesForCreate
type ServerPropertiesForDefaultCreate = original.ServerPropertiesForDefaultCreate
type ServerPropertiesForGeoRestore = original.ServerPropertiesForGeoRestore
type ServerPropertiesForReplica = original.ServerPropertiesForReplica
type ServerPropertiesForRestore = original.ServerPropertiesForRestore
type ServerSecurityAlertPoliciesClient = original.ServerSecurityAlertPoliciesClient
type ServerSecurityAlertPoliciesCreateOrUpdateFuture = original.ServerSecurityAlertPoliciesCreateOrUpdateFuture
type ServerSecurityAlertPolicy = original.ServerSecurityAlertPolicy
type ServerSecurityAlertPolicyListResult = original.ServerSecurityAlertPolicyListResult
type ServerSecurityAlertPolicyListResultIterator = original.ServerSecurityAlertPolicyListResultIterator
type ServerSecurityAlertPolicyListResultPage = original.ServerSecurityAlertPolicyListResultPage
type ServerUpdateParameters = original.ServerUpdateParameters
type ServerUpdateParametersProperties = original.ServerUpdateParametersProperties
type ServersClient = original.ServersClient
type ServersCreateFuture = original.ServersCreateFuture
type ServersDeleteFuture = original.ServersDeleteFuture
type ServersRestartFuture = original.ServersRestartFuture
type ServersUpdateFuture = original.ServersUpdateFuture
type Sku = original.Sku
type StorageProfile = original.StorageProfile
type TagsObject = original.TagsObject
type TrackedResource = original.TrackedResource
type VirtualNetworkRule = original.VirtualNetworkRule
type VirtualNetworkRuleListResult = original.VirtualNetworkRuleListResult
type VirtualNetworkRuleListResultIterator = original.VirtualNetworkRuleListResultIterator
type VirtualNetworkRuleListResultPage = original.VirtualNetworkRuleListResultPage
type VirtualNetworkRuleProperties = original.VirtualNetworkRuleProperties
type VirtualNetworkRulesClient = original.VirtualNetworkRulesClient
type VirtualNetworkRulesCreateOrUpdateFuture = original.VirtualNetworkRulesCreateOrUpdateFuture
type VirtualNetworkRulesDeleteFuture = original.VirtualNetworkRulesDeleteFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewCheckNameAvailabilityClient(subscriptionID string) CheckNameAvailabilityClient {
	return original.NewCheckNameAvailabilityClient(subscriptionID)
}
func NewCheckNameAvailabilityClientWithBaseURI(baseURI string, subscriptionID string) CheckNameAvailabilityClient {
	return original.NewCheckNameAvailabilityClientWithBaseURI(baseURI, subscriptionID)
}
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClient(subscriptionID)
}
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabasesClient(subscriptionID string) DatabasesClient {
	return original.NewDatabasesClient(subscriptionID)
}
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
	return original.NewDatabasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClient(subscriptionID)
}
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationBasedPerformanceTierClient(subscriptionID string) LocationBasedPerformanceTierClient {
	return original.NewLocationBasedPerformanceTierClient(subscriptionID)
}
func NewLocationBasedPerformanceTierClientWithBaseURI(baseURI string, subscriptionID string) LocationBasedPerformanceTierClient {
	return original.NewLocationBasedPerformanceTierClientWithBaseURI(baseURI, subscriptionID)
}
func NewLogFilesClient(subscriptionID string) LogFilesClient {
	return original.NewLogFilesClient(subscriptionID)
}
func NewLogFilesClientWithBaseURI(baseURI string, subscriptionID string) LogFilesClient {
	return original.NewLogFilesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionListResultIterator(page PrivateEndpointConnectionListResultPage) PrivateEndpointConnectionListResultIterator {
	return original.NewPrivateEndpointConnectionListResultIterator(page)
}
func NewPrivateEndpointConnectionListResultPage(cur PrivateEndpointConnectionListResult, getNextPage func(context.Context, PrivateEndpointConnectionListResult) (PrivateEndpointConnectionListResult, error)) PrivateEndpointConnectionListResultPage {
	return original.NewPrivateEndpointConnectionListResultPage(cur, getNextPage)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourceListResultIterator(page PrivateLinkResourceListResultPage) PrivateLinkResourceListResultIterator {
	return original.NewPrivateLinkResourceListResultIterator(page)
}
func NewPrivateLinkResourceListResultPage(cur PrivateLinkResourceListResult, getNextPage func(context.Context, PrivateLinkResourceListResult) (PrivateLinkResourceListResult, error)) PrivateLinkResourceListResultPage {
	return original.NewPrivateLinkResourceListResultPage(cur, getNextPage)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRecoverableServersClient(subscriptionID string) RecoverableServersClient {
	return original.NewRecoverableServersClient(subscriptionID)
}
func NewRecoverableServersClientWithBaseURI(baseURI string, subscriptionID string) RecoverableServersClient {
	return original.NewRecoverableServersClientWithBaseURI(baseURI, subscriptionID)
}
func NewReplicasClient(subscriptionID string) ReplicasClient {
	return original.NewReplicasClient(subscriptionID)
}
func NewReplicasClientWithBaseURI(baseURI string, subscriptionID string) ReplicasClient {
	return original.NewReplicasClientWithBaseURI(baseURI, subscriptionID)
}
func NewServerAdministratorsClient(subscriptionID string) ServerAdministratorsClient {
	return original.NewServerAdministratorsClient(subscriptionID)
}
func NewServerAdministratorsClientWithBaseURI(baseURI string, subscriptionID string) ServerAdministratorsClient {
	return original.NewServerAdministratorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServerBasedPerformanceTierClient(subscriptionID string) ServerBasedPerformanceTierClient {
	return original.NewServerBasedPerformanceTierClient(subscriptionID)
}
func NewServerBasedPerformanceTierClientWithBaseURI(baseURI string, subscriptionID string) ServerBasedPerformanceTierClient {
	return original.NewServerBasedPerformanceTierClientWithBaseURI(baseURI, subscriptionID)
}
func NewServerKeyListResultIterator(page ServerKeyListResultPage) ServerKeyListResultIterator {
	return original.NewServerKeyListResultIterator(page)
}
func NewServerKeyListResultPage(cur ServerKeyListResult, getNextPage func(context.Context, ServerKeyListResult) (ServerKeyListResult, error)) ServerKeyListResultPage {
	return original.NewServerKeyListResultPage(cur, getNextPage)
}
func NewServerKeysClient(subscriptionID string) ServerKeysClient {
	return original.NewServerKeysClient(subscriptionID)
}
func NewServerKeysClientWithBaseURI(baseURI string, subscriptionID string) ServerKeysClient {
	return original.NewServerKeysClientWithBaseURI(baseURI, subscriptionID)
}
func NewServerParametersClient(subscriptionID string) ServerParametersClient {
	return original.NewServerParametersClient(subscriptionID)
}
func NewServerParametersClientWithBaseURI(baseURI string, subscriptionID string) ServerParametersClient {
	return original.NewServerParametersClientWithBaseURI(baseURI, subscriptionID)
}
func NewServerSecurityAlertPoliciesClient(subscriptionID string) ServerSecurityAlertPoliciesClient {
	return original.NewServerSecurityAlertPoliciesClient(subscriptionID)
}
func NewServerSecurityAlertPoliciesClientWithBaseURI(baseURI string, subscriptionID string) ServerSecurityAlertPoliciesClient {
	return original.NewServerSecurityAlertPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewServerSecurityAlertPolicyListResultIterator(page ServerSecurityAlertPolicyListResultPage) ServerSecurityAlertPolicyListResultIterator {
	return original.NewServerSecurityAlertPolicyListResultIterator(page)
}
func NewServerSecurityAlertPolicyListResultPage(cur ServerSecurityAlertPolicyListResult, getNextPage func(context.Context, ServerSecurityAlertPolicyListResult) (ServerSecurityAlertPolicyListResult, error)) ServerSecurityAlertPolicyListResultPage {
	return original.NewServerSecurityAlertPolicyListResultPage(cur, getNextPage)
}
func NewServersClient(subscriptionID string) ServersClient {
	return original.NewServersClient(subscriptionID)
}
func NewServersClientWithBaseURI(baseURI string, subscriptionID string) ServersClient {
	return original.NewServersClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualNetworkRuleListResultIterator(page VirtualNetworkRuleListResultPage) VirtualNetworkRuleListResultIterator {
	return original.NewVirtualNetworkRuleListResultIterator(page)
}
func NewVirtualNetworkRuleListResultPage(cur VirtualNetworkRuleListResult, getNextPage func(context.Context, VirtualNetworkRuleListResult) (VirtualNetworkRuleListResult, error)) VirtualNetworkRuleListResultPage {
	return original.NewVirtualNetworkRuleListResultPage(cur, getNextPage)
}
func NewVirtualNetworkRulesClient(subscriptionID string) VirtualNetworkRulesClient {
	return original.NewVirtualNetworkRulesClient(subscriptionID)
}
func NewVirtualNetworkRulesClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkRulesClient {
	return original.NewVirtualNetworkRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreateModeValues() []CreateMode {
	return original.PossibleCreateModeValues()
}
func PossibleGeoRedundantBackupValues() []GeoRedundantBackup {
	return original.PossibleGeoRedundantBackupValues()
}
func PossibleIdentityTypeValues() []IdentityType {
	return original.PossibleIdentityTypeValues()
}
func PossibleInfrastructureEncryptionValues() []InfrastructureEncryption {
	return original.PossibleInfrastructureEncryptionValues()
}
func PossibleMinimalTLSVersionEnumValues() []MinimalTLSVersionEnum {
	return original.PossibleMinimalTLSVersionEnumValues()
}
func PossibleOperationOriginValues() []OperationOrigin {
	return original.PossibleOperationOriginValues()
}
func PossiblePrivateEndpointProvisioningStateValues() []PrivateEndpointProvisioningState {
	return original.PossiblePrivateEndpointProvisioningStateValues()
}
func PossiblePrivateLinkServiceConnectionStateActionsRequireValues() []PrivateLinkServiceConnectionStateActionsRequire {
	return original.PossiblePrivateLinkServiceConnectionStateActionsRequireValues()
}
func PossiblePrivateLinkServiceConnectionStateStatusValues() []PrivateLinkServiceConnectionStateStatus {
	return original.PossiblePrivateLinkServiceConnectionStateStatusValues()
}
func PossiblePublicNetworkAccessEnumValues() []PublicNetworkAccessEnum {
	return original.PossiblePublicNetworkAccessEnumValues()
}
func PossibleServerSecurityAlertPolicyStateValues() []ServerSecurityAlertPolicyState {
	return original.PossibleServerSecurityAlertPolicyStateValues()
}
func PossibleServerStateValues() []ServerState {
	return original.PossibleServerStateValues()
}
func PossibleServerVersionValues() []ServerVersion {
	return original.PossibleServerVersionValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleSslEnforcementEnumValues() []SslEnforcementEnum {
	return original.PossibleSslEnforcementEnumValues()
}
func PossibleStorageAutogrowValues() []StorageAutogrow {
	return original.PossibleStorageAutogrowValues()
}
func PossibleVirtualNetworkRuleStateValues() []VirtualNetworkRuleState {
	return original.PossibleVirtualNetworkRuleStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}