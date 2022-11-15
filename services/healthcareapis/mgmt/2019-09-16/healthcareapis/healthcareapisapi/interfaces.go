package healthcareapisapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/healthcareapis/mgmt/2019-09-16/healthcareapis"
)

// ServicesClientAPI contains the set of methods on the ServicesClient type.
type ServicesClientAPI interface {
	CheckNameAvailability(ctx context.Context, checkNameAvailabilityInputs healthcareapis.CheckNameAvailabilityParameters) (result healthcareapis.ServicesNameAvailabilityInfo, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, serviceDescription healthcareapis.ServicesDescription) (result healthcareapis.ServicesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result healthcareapis.ServicesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result healthcareapis.ServicesDescription, err error)
	List(ctx context.Context) (result healthcareapis.ServicesDescriptionListResultPage, err error)
	ListComplete(ctx context.Context) (result healthcareapis.ServicesDescriptionListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result healthcareapis.ServicesDescriptionListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result healthcareapis.ServicesDescriptionListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription healthcareapis.ServicesPatchDescription) (result healthcareapis.ServicesUpdateFuture, err error)
}

var _ ServicesClientAPI = (*healthcareapis.ServicesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result healthcareapis.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result healthcareapis.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*healthcareapis.OperationsClient)(nil)

// OperationResultsClientAPI contains the set of methods on the OperationResultsClient type.
type OperationResultsClientAPI interface {
	Get(ctx context.Context, locationName string, operationResultID string) (result healthcareapis.SetObject, err error)
}

var _ OperationResultsClientAPI = (*healthcareapis.OperationResultsClient)(nil)