//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package redhatopenshift

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/redhatopenshift/mgmt/2022-04-01/redhatopenshift"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type EncryptionAtHost = original.EncryptionAtHost

const (
	EncryptionAtHostDisabled EncryptionAtHost = original.EncryptionAtHostDisabled
	EncryptionAtHostEnabled  EncryptionAtHost = original.EncryptionAtHostEnabled
)

type FipsValidatedModules = original.FipsValidatedModules

const (
	FipsValidatedModulesDisabled FipsValidatedModules = original.FipsValidatedModulesDisabled
	FipsValidatedModulesEnabled  FipsValidatedModules = original.FipsValidatedModulesEnabled
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAdminUpdating ProvisioningState = original.ProvisioningStateAdminUpdating
	ProvisioningStateCreating      ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting      ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed        ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateSucceeded     ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating      ProvisioningState = original.ProvisioningStateUpdating
)

type Visibility = original.Visibility

const (
	VisibilityPrivate Visibility = original.VisibilityPrivate
	VisibilityPublic  Visibility = original.VisibilityPublic
)

type APIServerProfile = original.APIServerProfile
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type ClusterProfile = original.ClusterProfile
type ConsoleProfile = original.ConsoleProfile
type Display = original.Display
type IngressProfile = original.IngressProfile
type MasterProfile = original.MasterProfile
type NetworkProfile = original.NetworkProfile
type OpenShiftCluster = original.OpenShiftCluster
type OpenShiftClusterAdminKubeconfig = original.OpenShiftClusterAdminKubeconfig
type OpenShiftClusterCredentials = original.OpenShiftClusterCredentials
type OpenShiftClusterList = original.OpenShiftClusterList
type OpenShiftClusterListIterator = original.OpenShiftClusterListIterator
type OpenShiftClusterListPage = original.OpenShiftClusterListPage
type OpenShiftClusterProperties = original.OpenShiftClusterProperties
type OpenShiftClusterUpdate = original.OpenShiftClusterUpdate
type OpenShiftClustersClient = original.OpenShiftClustersClient
type OpenShiftClustersCreateOrUpdateFuture = original.OpenShiftClustersCreateOrUpdateFuture
type OpenShiftClustersDeleteFuture = original.OpenShiftClustersDeleteFuture
type OpenShiftClustersUpdateFuture = original.OpenShiftClustersUpdateFuture
type Operation = original.Operation
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ServicePrincipalProfile = original.ServicePrincipalProfile
type SystemData = original.SystemData
type TrackedResource = original.TrackedResource
type WorkerProfile = original.WorkerProfile

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewOpenShiftClusterListIterator(page OpenShiftClusterListPage) OpenShiftClusterListIterator {
	return original.NewOpenShiftClusterListIterator(page)
}
func NewOpenShiftClusterListPage(cur OpenShiftClusterList, getNextPage func(context.Context, OpenShiftClusterList) (OpenShiftClusterList, error)) OpenShiftClusterListPage {
	return original.NewOpenShiftClusterListPage(cur, getNextPage)
}
func NewOpenShiftClustersClient(subscriptionID string) OpenShiftClustersClient {
	return original.NewOpenShiftClustersClient(subscriptionID)
}
func NewOpenShiftClustersClientWithBaseURI(baseURI string, subscriptionID string) OpenShiftClustersClient {
	return original.NewOpenShiftClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(cur OperationList, getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleEncryptionAtHostValues() []EncryptionAtHost {
	return original.PossibleEncryptionAtHostValues()
}
func PossibleFipsValidatedModulesValues() []FipsValidatedModules {
	return original.PossibleFipsValidatedModulesValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleVisibilityValues() []Visibility {
	return original.PossibleVisibilityValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}