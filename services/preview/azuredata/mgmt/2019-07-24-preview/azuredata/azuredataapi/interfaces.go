package azuredataapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/azuredata/mgmt/2019-07-24-preview/azuredata"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result azuredata.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result azuredata.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*azuredata.OperationsClient)(nil)

// SQLServerRegistrationsClientAPI contains the set of methods on the SQLServerRegistrationsClient type.
type SQLServerRegistrationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, parameters azuredata.SQLServerRegistration) (result azuredata.SQLServerRegistration, err error)
	Delete(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string) (result azuredata.SQLServerRegistration, err error)
	List(ctx context.Context) (result azuredata.SQLServerRegistrationListResultPage, err error)
	ListComplete(ctx context.Context) (result azuredata.SQLServerRegistrationListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result azuredata.SQLServerRegistrationListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result azuredata.SQLServerRegistrationListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, parameters azuredata.SQLServerRegistrationUpdate) (result azuredata.SQLServerRegistration, err error)
}

var _ SQLServerRegistrationsClientAPI = (*azuredata.SQLServerRegistrationsClient)(nil)

// SQLServersClientAPI contains the set of methods on the SQLServersClient type.
type SQLServersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, SQLServerName string, parameters azuredata.SQLServer) (result azuredata.SQLServer, err error)
	Delete(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, SQLServerName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, SQLServerName string, expand string) (result azuredata.SQLServer, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, expand string) (result azuredata.SQLServerListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, expand string) (result azuredata.SQLServerListResultIterator, err error)
}

var _ SQLServersClientAPI = (*azuredata.SQLServersClient)(nil)

// SQLManagedInstancesClientAPI contains the set of methods on the SQLManagedInstancesClient type.
type SQLManagedInstancesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, SQLManagedInstanceName string, parameters azuredata.SQLManagedInstance) (result azuredata.SQLManagedInstance, err error)
	Delete(ctx context.Context, resourceGroupName string, SQLManagedInstanceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, SQLManagedInstanceName string) (result azuredata.SQLManagedInstance, err error)
	List(ctx context.Context) (result azuredata.SQLManagedInstanceListResultPage, err error)
	ListComplete(ctx context.Context) (result azuredata.SQLManagedInstanceListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result azuredata.SQLManagedInstanceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result azuredata.SQLManagedInstanceListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, SQLManagedInstanceName string, parameters azuredata.SQLManagedInstanceUpdate) (result azuredata.SQLManagedInstance, err error)
}

var _ SQLManagedInstancesClientAPI = (*azuredata.SQLManagedInstancesClient)(nil)

// SQLServerInstancesClientAPI contains the set of methods on the SQLServerInstancesClient type.
type SQLServerInstancesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, SQLServerInstanceName string, parameters azuredata.SQLServerInstance) (result azuredata.SQLServerInstance, err error)
	Delete(ctx context.Context, resourceGroupName string, SQLServerInstanceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, SQLServerInstanceName string) (result azuredata.SQLServerInstance, err error)
	List(ctx context.Context) (result azuredata.SQLServerInstanceListResultPage, err error)
	ListComplete(ctx context.Context) (result azuredata.SQLServerInstanceListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result azuredata.SQLServerInstanceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result azuredata.SQLServerInstanceListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, SQLServerInstanceName string, parameters azuredata.SQLServerInstanceUpdate) (result azuredata.SQLServerInstance, err error)
}

var _ SQLServerInstancesClientAPI = (*azuredata.SQLServerInstancesClient)(nil)

// PostgresInstancesClientAPI contains the set of methods on the PostgresInstancesClient type.
type PostgresInstancesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, postgresInstanceName string) (result azuredata.PostgresInstance, err error)
	Delete(ctx context.Context, resourceGroupName string, postgresInstanceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, postgresInstanceName string) (result azuredata.PostgresInstance, err error)
	List(ctx context.Context) (result azuredata.PostgresInstanceListResultPage, err error)
	ListComplete(ctx context.Context) (result azuredata.PostgresInstanceListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result azuredata.PostgresInstanceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result azuredata.PostgresInstanceListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, postgresInstanceName string, parameters azuredata.PostgresInstanceUpdate) (result azuredata.PostgresInstance, err error)
}

var _ PostgresInstancesClientAPI = (*azuredata.PostgresInstancesClient)(nil)

// DataControllersClientAPI contains the set of methods on the DataControllersClient type.
type DataControllersClientAPI interface {
	DeleteDataController(ctx context.Context, resourceGroupName string, dataControllerName string) (result autorest.Response, err error)
	GetDataController(ctx context.Context, resourceGroupName string, dataControllerName string) (result azuredata.DataControllerResource, err error)
	ListInGroup(ctx context.Context, resourceGroupName string) (result azuredata.PageOfDataControllerResourcePage, err error)
	ListInGroupComplete(ctx context.Context, resourceGroupName string) (result azuredata.PageOfDataControllerResourceIterator, err error)
	ListInSubscription(ctx context.Context) (result azuredata.PageOfDataControllerResourcePage, err error)
	ListInSubscriptionComplete(ctx context.Context) (result azuredata.PageOfDataControllerResourceIterator, err error)
	PatchDataController(ctx context.Context, resourceGroupName string, dataControllerName string, dataControllerResource azuredata.DataControllerUpdate) (result azuredata.DataControllerResource, err error)
	PutDataController(ctx context.Context, resourceGroupName string, dataControllerResource azuredata.DataControllerResource, dataControllerName string) (result azuredata.DataControllerResource, err error)
}

var _ DataControllersClientAPI = (*azuredata.DataControllersClient)(nil)