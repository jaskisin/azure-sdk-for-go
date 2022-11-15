//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package kubernetesconfiguration

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/kubernetesconfiguration/mgmt/2022-07-01/kubernetesconfiguration"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AKSIdentityType = original.AKSIdentityType

const (
	SystemAssigned AKSIdentityType = original.SystemAssigned
	UserAssigned   AKSIdentityType = original.UserAssigned
)

type ComplianceStateType = original.ComplianceStateType

const (
	Compliant    ComplianceStateType = original.Compliant
	Failed       ComplianceStateType = original.Failed
	Installed    ComplianceStateType = original.Installed
	Noncompliant ComplianceStateType = original.Noncompliant
	Pending      ComplianceStateType = original.Pending
)

type CreatedByType = original.CreatedByType

const (
	Application     CreatedByType = original.Application
	Key             CreatedByType = original.Key
	ManagedIdentity CreatedByType = original.ManagedIdentity
	User            CreatedByType = original.User
)

type FluxComplianceState = original.FluxComplianceState

const (
	FluxComplianceStateCompliant    FluxComplianceState = original.FluxComplianceStateCompliant
	FluxComplianceStateNonCompliant FluxComplianceState = original.FluxComplianceStateNonCompliant
	FluxComplianceStatePending      FluxComplianceState = original.FluxComplianceStatePending
	FluxComplianceStateSuspended    FluxComplianceState = original.FluxComplianceStateSuspended
	FluxComplianceStateUnknown      FluxComplianceState = original.FluxComplianceStateUnknown
)

type KustomizationValidationType = original.KustomizationValidationType

const (
	Client KustomizationValidationType = original.Client
	None   KustomizationValidationType = original.None
	Server KustomizationValidationType = original.Server
)

type LevelType = original.LevelType

const (
	Error       LevelType = original.Error
	Information LevelType = original.Information
	Warning     LevelType = original.Warning
)

type MessageLevelType = original.MessageLevelType

const (
	MessageLevelTypeError       MessageLevelType = original.MessageLevelTypeError
	MessageLevelTypeInformation MessageLevelType = original.MessageLevelTypeInformation
	MessageLevelTypeWarning     MessageLevelType = original.MessageLevelTypeWarning
)

type OperatorScopeType = original.OperatorScopeType

const (
	Cluster   OperatorScopeType = original.Cluster
	Namespace OperatorScopeType = original.Namespace
)

type OperatorType = original.OperatorType

const (
	Flux OperatorType = original.Flux
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCanceled  ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating  ProvisioningState = original.ProvisioningStateUpdating
)

type ProvisioningStateType = original.ProvisioningStateType

const (
	ProvisioningStateTypeAccepted  ProvisioningStateType = original.ProvisioningStateTypeAccepted
	ProvisioningStateTypeDeleting  ProvisioningStateType = original.ProvisioningStateTypeDeleting
	ProvisioningStateTypeFailed    ProvisioningStateType = original.ProvisioningStateTypeFailed
	ProvisioningStateTypeRunning   ProvisioningStateType = original.ProvisioningStateTypeRunning
	ProvisioningStateTypeSucceeded ProvisioningStateType = original.ProvisioningStateTypeSucceeded
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
)

type ScopeType = original.ScopeType

const (
	ScopeTypeCluster   ScopeType = original.ScopeTypeCluster
	ScopeTypeNamespace ScopeType = original.ScopeTypeNamespace
)

type SkuTier = original.SkuTier

const (
	Basic    SkuTier = original.Basic
	Free     SkuTier = original.Free
	Premium  SkuTier = original.Premium
	Standard SkuTier = original.Standard
)

type SourceKindType = original.SourceKindType

const (
	AzureBlob     SourceKindType = original.AzureBlob
	Bucket        SourceKindType = original.Bucket
	GitRepository SourceKindType = original.GitRepository
)

type AzureBlobDefinition = original.AzureBlobDefinition
type AzureBlobPatchDefinition = original.AzureBlobPatchDefinition
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type BucketDefinition = original.BucketDefinition
type BucketPatchDefinition = original.BucketPatchDefinition
type ComplianceStatus = original.ComplianceStatus
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type Extension = original.Extension
type ExtensionProperties = original.ExtensionProperties
type ExtensionPropertiesAksAssignedIdentity = original.ExtensionPropertiesAksAssignedIdentity
type ExtensionStatus = original.ExtensionStatus
type ExtensionsClient = original.ExtensionsClient
type ExtensionsCreateFuture = original.ExtensionsCreateFuture
type ExtensionsDeleteFuture = original.ExtensionsDeleteFuture
type ExtensionsList = original.ExtensionsList
type ExtensionsListIterator = original.ExtensionsListIterator
type ExtensionsListPage = original.ExtensionsListPage
type ExtensionsUpdateFuture = original.ExtensionsUpdateFuture
type FluxConfigOperationStatusClient = original.FluxConfigOperationStatusClient
type FluxConfiguration = original.FluxConfiguration
type FluxConfigurationPatch = original.FluxConfigurationPatch
type FluxConfigurationPatchProperties = original.FluxConfigurationPatchProperties
type FluxConfigurationProperties = original.FluxConfigurationProperties
type FluxConfigurationsClient = original.FluxConfigurationsClient
type FluxConfigurationsCreateOrUpdateFuture = original.FluxConfigurationsCreateOrUpdateFuture
type FluxConfigurationsDeleteFuture = original.FluxConfigurationsDeleteFuture
type FluxConfigurationsList = original.FluxConfigurationsList
type FluxConfigurationsListIterator = original.FluxConfigurationsListIterator
type FluxConfigurationsListPage = original.FluxConfigurationsListPage
type FluxConfigurationsUpdateFuture = original.FluxConfigurationsUpdateFuture
type GitRepositoryDefinition = original.GitRepositoryDefinition
type GitRepositoryPatchDefinition = original.GitRepositoryPatchDefinition
type HelmOperatorProperties = original.HelmOperatorProperties
type HelmReleasePropertiesDefinition = original.HelmReleasePropertiesDefinition
type Identity = original.Identity
type KustomizationDefinition = original.KustomizationDefinition
type KustomizationPatchDefinition = original.KustomizationPatchDefinition
type ManagedIdentityDefinition = original.ManagedIdentityDefinition
type ManagedIdentityPatchDefinition = original.ManagedIdentityPatchDefinition
type ObjectReferenceDefinition = original.ObjectReferenceDefinition
type ObjectStatusConditionDefinition = original.ObjectStatusConditionDefinition
type ObjectStatusDefinition = original.ObjectStatusDefinition
type OperationStatusClient = original.OperationStatusClient
type OperationStatusList = original.OperationStatusList
type OperationStatusListIterator = original.OperationStatusListIterator
type OperationStatusListPage = original.OperationStatusListPage
type OperationStatusResult = original.OperationStatusResult
type OperationsClient = original.OperationsClient
type PatchExtension = original.PatchExtension
type PatchExtensionProperties = original.PatchExtensionProperties
type Plan = original.Plan
type ProxyResource = original.ProxyResource
type RepositoryRefDefinition = original.RepositoryRefDefinition
type Resource = original.Resource
type ResourceModelWithAllowedPropertySet = original.ResourceModelWithAllowedPropertySet
type ResourceModelWithAllowedPropertySetIdentity = original.ResourceModelWithAllowedPropertySetIdentity
type ResourceModelWithAllowedPropertySetPlan = original.ResourceModelWithAllowedPropertySetPlan
type ResourceModelWithAllowedPropertySetSku = original.ResourceModelWithAllowedPropertySetSku
type ResourceProviderOperation = original.ResourceProviderOperation
type ResourceProviderOperationDisplay = original.ResourceProviderOperationDisplay
type ResourceProviderOperationList = original.ResourceProviderOperationList
type ResourceProviderOperationListIterator = original.ResourceProviderOperationListIterator
type ResourceProviderOperationListPage = original.ResourceProviderOperationListPage
type Scope = original.Scope
type ScopeCluster = original.ScopeCluster
type ScopeNamespace = original.ScopeNamespace
type ServicePrincipalDefinition = original.ServicePrincipalDefinition
type ServicePrincipalPatchDefinition = original.ServicePrincipalPatchDefinition
type Sku = original.Sku
type SourceControlConfiguration = original.SourceControlConfiguration
type SourceControlConfigurationList = original.SourceControlConfigurationList
type SourceControlConfigurationListIterator = original.SourceControlConfigurationListIterator
type SourceControlConfigurationListPage = original.SourceControlConfigurationListPage
type SourceControlConfigurationProperties = original.SourceControlConfigurationProperties
type SourceControlConfigurationsClient = original.SourceControlConfigurationsClient
type SourceControlConfigurationsDeleteFuture = original.SourceControlConfigurationsDeleteFuture
type SystemData = original.SystemData
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewExtensionsClient(subscriptionID string) ExtensionsClient {
	return original.NewExtensionsClient(subscriptionID)
}
func NewExtensionsClientWithBaseURI(baseURI string, subscriptionID string) ExtensionsClient {
	return original.NewExtensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewExtensionsListIterator(page ExtensionsListPage) ExtensionsListIterator {
	return original.NewExtensionsListIterator(page)
}
func NewExtensionsListPage(cur ExtensionsList, getNextPage func(context.Context, ExtensionsList) (ExtensionsList, error)) ExtensionsListPage {
	return original.NewExtensionsListPage(cur, getNextPage)
}
func NewFluxConfigOperationStatusClient(subscriptionID string) FluxConfigOperationStatusClient {
	return original.NewFluxConfigOperationStatusClient(subscriptionID)
}
func NewFluxConfigOperationStatusClientWithBaseURI(baseURI string, subscriptionID string) FluxConfigOperationStatusClient {
	return original.NewFluxConfigOperationStatusClientWithBaseURI(baseURI, subscriptionID)
}
func NewFluxConfigurationsClient(subscriptionID string) FluxConfigurationsClient {
	return original.NewFluxConfigurationsClient(subscriptionID)
}
func NewFluxConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) FluxConfigurationsClient {
	return original.NewFluxConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFluxConfigurationsListIterator(page FluxConfigurationsListPage) FluxConfigurationsListIterator {
	return original.NewFluxConfigurationsListIterator(page)
}
func NewFluxConfigurationsListPage(cur FluxConfigurationsList, getNextPage func(context.Context, FluxConfigurationsList) (FluxConfigurationsList, error)) FluxConfigurationsListPage {
	return original.NewFluxConfigurationsListPage(cur, getNextPage)
}
func NewOperationStatusClient(subscriptionID string) OperationStatusClient {
	return original.NewOperationStatusClient(subscriptionID)
}
func NewOperationStatusClientWithBaseURI(baseURI string, subscriptionID string) OperationStatusClient {
	return original.NewOperationStatusClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationStatusListIterator(page OperationStatusListPage) OperationStatusListIterator {
	return original.NewOperationStatusListIterator(page)
}
func NewOperationStatusListPage(cur OperationStatusList, getNextPage func(context.Context, OperationStatusList) (OperationStatusList, error)) OperationStatusListPage {
	return original.NewOperationStatusListPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceProviderOperationListIterator(page ResourceProviderOperationListPage) ResourceProviderOperationListIterator {
	return original.NewResourceProviderOperationListIterator(page)
}
func NewResourceProviderOperationListPage(cur ResourceProviderOperationList, getNextPage func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)) ResourceProviderOperationListPage {
	return original.NewResourceProviderOperationListPage(cur, getNextPage)
}
func NewSourceControlConfigurationListIterator(page SourceControlConfigurationListPage) SourceControlConfigurationListIterator {
	return original.NewSourceControlConfigurationListIterator(page)
}
func NewSourceControlConfigurationListPage(cur SourceControlConfigurationList, getNextPage func(context.Context, SourceControlConfigurationList) (SourceControlConfigurationList, error)) SourceControlConfigurationListPage {
	return original.NewSourceControlConfigurationListPage(cur, getNextPage)
}
func NewSourceControlConfigurationsClient(subscriptionID string) SourceControlConfigurationsClient {
	return original.NewSourceControlConfigurationsClient(subscriptionID)
}
func NewSourceControlConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) SourceControlConfigurationsClient {
	return original.NewSourceControlConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAKSIdentityTypeValues() []AKSIdentityType {
	return original.PossibleAKSIdentityTypeValues()
}
func PossibleComplianceStateTypeValues() []ComplianceStateType {
	return original.PossibleComplianceStateTypeValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleFluxComplianceStateValues() []FluxComplianceState {
	return original.PossibleFluxComplianceStateValues()
}
func PossibleKustomizationValidationTypeValues() []KustomizationValidationType {
	return original.PossibleKustomizationValidationTypeValues()
}
func PossibleLevelTypeValues() []LevelType {
	return original.PossibleLevelTypeValues()
}
func PossibleMessageLevelTypeValues() []MessageLevelType {
	return original.PossibleMessageLevelTypeValues()
}
func PossibleOperatorScopeTypeValues() []OperatorScopeType {
	return original.PossibleOperatorScopeTypeValues()
}
func PossibleOperatorTypeValues() []OperatorType {
	return original.PossibleOperatorTypeValues()
}
func PossibleProvisioningStateTypeValues() []ProvisioningStateType {
	return original.PossibleProvisioningStateTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleScopeTypeValues() []ScopeType {
	return original.PossibleScopeTypeValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleSourceKindTypeValues() []SourceKindType {
	return original.PossibleSourceKindTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}