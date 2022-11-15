package storsimple

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// BackupScheduleGroupsClient is the client for the BackupScheduleGroups methods of the Storsimple service.
type BackupScheduleGroupsClient struct {
	BaseClient
}

// NewBackupScheduleGroupsClient creates an instance of the BackupScheduleGroupsClient client.
func NewBackupScheduleGroupsClient(subscriptionID string) BackupScheduleGroupsClient {
	return NewBackupScheduleGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackupScheduleGroupsClientWithBaseURI creates an instance of the BackupScheduleGroupsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewBackupScheduleGroupsClientWithBaseURI(baseURI string, subscriptionID string) BackupScheduleGroupsClient {
	return BackupScheduleGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or Updates the backup schedule Group.
// Parameters:
// deviceName - the name of the device.
// scheduleGroupName - the name of the schedule group.
// scheduleGroup - the schedule group to be created
// resourceGroupName - the resource group name
// managerName - the manager name
func (client BackupScheduleGroupsClient) CreateOrUpdate(ctx context.Context, deviceName string, scheduleGroupName string, scheduleGroup BackupScheduleGroup, resourceGroupName string, managerName string) (result BackupScheduleGroupsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupScheduleGroupsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: scheduleGroup,
			Constraints: []validation.Constraint{{Target: "scheduleGroup.BackupScheduleGroupProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "scheduleGroup.BackupScheduleGroupProperties.StartTime", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "scheduleGroup.BackupScheduleGroupProperties.StartTime.Hour", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "scheduleGroup.BackupScheduleGroupProperties.StartTime.Hour", Name: validation.InclusiveMaximum, Rule: int64(23), Chain: nil},
							{Target: "scheduleGroup.BackupScheduleGroupProperties.StartTime.Hour", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
						}},
						{Target: "scheduleGroup.BackupScheduleGroupProperties.StartTime.Minute", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "scheduleGroup.BackupScheduleGroupProperties.StartTime.Minute", Name: validation.InclusiveMaximum, Rule: int64(59), Chain: nil},
								{Target: "scheduleGroup.BackupScheduleGroupProperties.StartTime.Minute", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
							}},
					}},
				}}}},
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupScheduleGroupsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, deviceName, scheduleGroupName, scheduleGroup, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupScheduleGroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupScheduleGroupsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client BackupScheduleGroupsClient) CreateOrUpdatePreparer(ctx context.Context, deviceName string, scheduleGroupName string, scheduleGroup BackupScheduleGroup, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"scheduleGroupName": autorest.Encode("path", scheduleGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups/{scheduleGroupName}", pathParameters),
		autorest.WithJSON(scheduleGroup),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client BackupScheduleGroupsClient) CreateOrUpdateSender(req *http.Request) (future BackupScheduleGroupsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client BackupScheduleGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result BackupScheduleGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the backup schedule group.
// Parameters:
// deviceName - the name of the device.
// scheduleGroupName - the name of the schedule group.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client BackupScheduleGroupsClient) Delete(ctx context.Context, deviceName string, scheduleGroupName string, resourceGroupName string, managerName string) (result BackupScheduleGroupsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupScheduleGroupsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupScheduleGroupsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, deviceName, scheduleGroupName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupScheduleGroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupScheduleGroupsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BackupScheduleGroupsClient) DeletePreparer(ctx context.Context, deviceName string, scheduleGroupName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"scheduleGroupName": autorest.Encode("path", scheduleGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups/{scheduleGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BackupScheduleGroupsClient) DeleteSender(req *http.Request) (future BackupScheduleGroupsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BackupScheduleGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns the properties of the specified backup schedule group name.
// Parameters:
// deviceName - the name of the device.
// scheduleGroupName - the name of the schedule group.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client BackupScheduleGroupsClient) Get(ctx context.Context, deviceName string, scheduleGroupName string, resourceGroupName string, managerName string) (result BackupScheduleGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupScheduleGroupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupScheduleGroupsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, deviceName, scheduleGroupName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupScheduleGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.BackupScheduleGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupScheduleGroupsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client BackupScheduleGroupsClient) GetPreparer(ctx context.Context, deviceName string, scheduleGroupName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"scheduleGroupName": autorest.Encode("path", scheduleGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups/{scheduleGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BackupScheduleGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BackupScheduleGroupsClient) GetResponder(resp *http.Response) (result BackupScheduleGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDevice retrieves all the backup schedule groups in a device.
// Parameters:
// deviceName - the name of the device.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client BackupScheduleGroupsClient) ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result BackupScheduleGroupList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupScheduleGroupsClient.ListByDevice")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupScheduleGroupsClient", "ListByDevice", err.Error())
	}

	req, err := client.ListByDevicePreparer(ctx, deviceName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupScheduleGroupsClient", "ListByDevice", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDeviceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.BackupScheduleGroupsClient", "ListByDevice", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupScheduleGroupsClient", "ListByDevice", resp, "Failure responding to request")
		return
	}

	return
}

// ListByDevicePreparer prepares the ListByDevice request.
func (client BackupScheduleGroupsClient) ListByDevicePreparer(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDeviceSender sends the ListByDevice request. The method will close the
// http.Response Body if it receives an error.
func (client BackupScheduleGroupsClient) ListByDeviceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDeviceResponder handles the response to the ListByDevice request. The method always
// closes the http.Response Body.
func (client BackupScheduleGroupsClient) ListByDeviceResponder(resp *http.Response) (result BackupScheduleGroupList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}