package kustoapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/kusto/mgmt/2019-11-09/kusto"
)

// ClustersClientAPI contains the set of methods on the ClustersClient type.
type ClustersClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, clusterName kusto.ClusterCheckNameRequest) (result kusto.CheckNameResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters kusto.Cluster) (result kusto.ClustersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.ClustersDeleteFuture, err error)
	DetachFollowerDatabases(ctx context.Context, resourceGroupName string, clusterName string, followerDatabaseToRemove kusto.FollowerDatabaseDefinition) (result kusto.ClustersDetachFollowerDatabasesFuture, err error)
	DiagnoseVirtualNetwork(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.ClustersDiagnoseVirtualNetworkFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.Cluster, err error)
	List(ctx context.Context) (result kusto.ClusterListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result kusto.ClusterListResult, err error)
	ListFollowerDatabases(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.FollowerDatabaseListResult, err error)
	ListSkus(ctx context.Context) (result kusto.SkuDescriptionList, err error)
	ListSkusByResource(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.ListResourceSkusResult, err error)
	Start(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.ClustersStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.ClustersStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, parameters kusto.ClusterUpdate) (result kusto.ClustersUpdateFuture, err error)
}

var _ ClustersClientAPI = (*kusto.ClustersClient)(nil)

// ClusterPrincipalAssignmentsClientAPI contains the set of methods on the ClusterPrincipalAssignmentsClient type.
type ClusterPrincipalAssignmentsClientAPI interface {
	CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName kusto.ClusterPrincipalAssignmentCheckNameRequest) (result kusto.CheckNameResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string, parameters kusto.ClusterPrincipalAssignment) (result kusto.ClusterPrincipalAssignmentsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string) (result kusto.ClusterPrincipalAssignmentsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string) (result kusto.ClusterPrincipalAssignment, err error)
	List(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.ClusterPrincipalAssignmentListResult, err error)
}

var _ ClusterPrincipalAssignmentsClientAPI = (*kusto.ClusterPrincipalAssignmentsClient)(nil)

// DatabasesClientAPI contains the set of methods on the DatabasesClient type.
type DatabasesClientAPI interface {
	AddPrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToAdd kusto.DatabasePrincipalListRequest) (result kusto.DatabasePrincipalListResult, err error)
	CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, resourceName kusto.CheckNameRequest) (result kusto.CheckNameResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters kusto.BasicDatabase) (result kusto.DatabasesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result kusto.DatabasesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result kusto.DatabaseModel, err error)
	ListByCluster(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.DatabaseListResult, err error)
	ListPrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result kusto.DatabasePrincipalListResult, err error)
	RemovePrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToRemove kusto.DatabasePrincipalListRequest) (result kusto.DatabasePrincipalListResult, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters kusto.BasicDatabase) (result kusto.DatabasesUpdateFuture, err error)
}

var _ DatabasesClientAPI = (*kusto.DatabasesClient)(nil)

// DatabasePrincipalAssignmentsClientAPI contains the set of methods on the DatabasePrincipalAssignmentsClient type.
type DatabasePrincipalAssignmentsClientAPI interface {
	CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, principalAssignmentName kusto.DatabasePrincipalAssignmentCheckNameRequest) (result kusto.CheckNameResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, principalAssignmentName string, parameters kusto.DatabasePrincipalAssignment) (result kusto.DatabasePrincipalAssignmentsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, principalAssignmentName string) (result kusto.DatabasePrincipalAssignmentsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, principalAssignmentName string) (result kusto.DatabasePrincipalAssignment, err error)
	List(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result kusto.DatabasePrincipalAssignmentListResult, err error)
}

var _ DatabasePrincipalAssignmentsClientAPI = (*kusto.DatabasePrincipalAssignmentsClient)(nil)

// AttachedDatabaseConfigurationsClientAPI contains the set of methods on the AttachedDatabaseConfigurationsClient type.
type AttachedDatabaseConfigurationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, attachedDatabaseConfigurationName string, parameters kusto.AttachedDatabaseConfiguration) (result kusto.AttachedDatabaseConfigurationsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, attachedDatabaseConfigurationName string) (result kusto.AttachedDatabaseConfigurationsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, attachedDatabaseConfigurationName string) (result kusto.AttachedDatabaseConfiguration, err error)
	ListByCluster(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.AttachedDatabaseConfigurationListResult, err error)
}

var _ AttachedDatabaseConfigurationsClientAPI = (*kusto.AttachedDatabaseConfigurationsClient)(nil)

// DataConnectionsClientAPI contains the set of methods on the DataConnectionsClient type.
type DataConnectionsClientAPI interface {
	CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName kusto.DataConnectionCheckNameRequest) (result kusto.CheckNameResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters kusto.BasicDataConnection) (result kusto.DataConnectionsCreateOrUpdateFuture, err error)
	DataConnectionValidationMethod(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters kusto.DataConnectionValidation) (result kusto.DataConnectionValidationListResult, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string) (result kusto.DataConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string) (result kusto.DataConnectionModel, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result kusto.DataConnectionListResult, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters kusto.BasicDataConnection) (result kusto.DataConnectionsUpdateFuture, err error)
}

var _ DataConnectionsClientAPI = (*kusto.DataConnectionsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result kusto.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result kusto.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*kusto.OperationsClient)(nil)