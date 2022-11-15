package labservices

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

// LabsClient is the the Managed Labs Client.
type LabsClient struct {
	BaseClient
}

// NewLabsClient creates an instance of the LabsClient client.
func NewLabsClient(subscriptionID string) LabsClient {
	return NewLabsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLabsClientWithBaseURI creates an instance of the LabsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLabsClientWithBaseURI(baseURI string, subscriptionID string) LabsClient {
	return LabsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// AddUsers add users to a lab
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// addUsersPayload - payload for Add Users operation on a Lab.
func (client LabsClient) AddUsers(ctx context.Context, resourceGroupName string, labAccountName string, labName string, addUsersPayload AddUsersPayload) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.AddUsers")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: addUsersPayload,
			Constraints: []validation.Constraint{{Target: "addUsersPayload.EmailAddresses", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.LabsClient", "AddUsers", err.Error())
	}

	req, err := client.AddUsersPreparer(ctx, resourceGroupName, labAccountName, labName, addUsersPayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "AddUsers", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddUsersSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "AddUsers", resp, "Failure sending request")
		return
	}

	result, err = client.AddUsersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "AddUsers", resp, "Failure responding to request")
		return
	}

	return
}

// AddUsersPreparer prepares the AddUsers request.
func (client LabsClient) AddUsersPreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, addUsersPayload AddUsersPayload) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labAccountName":    autorest.Encode("path", labAccountName),
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/addUsers", pathParameters),
		autorest.WithJSON(addUsersPayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddUsersSender sends the AddUsers request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) AddUsersSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// AddUsersResponder handles the response to the AddUsers request. The method always
// closes the http.Response Body.
func (client LabsClient) AddUsersResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate create or replace an existing Lab.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// lab - represents a lab.
func (client LabsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labAccountName string, labName string, lab Lab) (result Lab, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, labAccountName, labName, lab)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client LabsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, lab Lab) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labAccountName":    autorest.Encode("path", labAccountName),
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", pathParameters),
		autorest.WithJSON(lab),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client LabsClient) CreateOrUpdateResponder(resp *http.Response) (result Lab, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete lab. This operation can take a while to complete
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
func (client LabsClient) Delete(ctx context.Context, resourceGroupName string, labAccountName string, labName string) (result LabsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, labAccountName, labName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LabsClient) DeletePreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labAccountName":    autorest.Encode("path", labAccountName),
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) DeleteSender(req *http.Request) (future LabsDeleteFuture, err error) {
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
func (client LabsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get lab
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// expand - specify the $expand query. Example: 'properties($select=maxUsersInLab)'
func (client LabsClient) Get(ctx context.Context, resourceGroupName string, labAccountName string, labName string, expand string) (result Lab, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, labAccountName, labName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client LabsClient) GetPreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labAccountName":    autorest.Encode("path", labAccountName),
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LabsClient) GetResponder(resp *http.Response) (result Lab, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list labs in a given lab account.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// expand - specify the $expand query. Example: 'properties($select=maxUsersInLab)'
// filter - the filter to apply to the operation.
// top - the maximum number of resources to return from the operation.
// orderby - the ordering expression for the results, using OData notation.
func (client LabsClient) List(ctx context.Context, resourceGroupName string, labAccountName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationLabPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.List")
		defer func() {
			sc := -1
			if result.rwcl.Response.Response != nil {
				sc = result.rwcl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, labAccountName, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rwcl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "List", resp, "Failure sending request")
		return
	}

	result.rwcl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rwcl.hasNextLink() && result.rwcl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client LabsClient) ListPreparer(ctx context.Context, resourceGroupName string, labAccountName string, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labAccountName":    autorest.Encode("path", labAccountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LabsClient) ListResponder(resp *http.Response) (result ResponseWithContinuationLab, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client LabsClient) listNextResults(ctx context.Context, lastResults ResponseWithContinuationLab) (result ResponseWithContinuationLab, err error) {
	req, err := lastResults.responseWithContinuationLabPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "labservices.LabsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "labservices.LabsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client LabsClient) ListComplete(ctx context.Context, resourceGroupName string, labAccountName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationLabIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, labAccountName, expand, filter, top, orderby)
	return
}

// Register register to managed lab.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
func (client LabsClient) Register(ctx context.Context, resourceGroupName string, labAccountName string, labName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.Register")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RegisterPreparer(ctx, resourceGroupName, labAccountName, labName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Register", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegisterSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Register", resp, "Failure sending request")
		return
	}

	result, err = client.RegisterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Register", resp, "Failure responding to request")
		return
	}

	return
}

// RegisterPreparer prepares the Register request.
func (client LabsClient) RegisterPreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labAccountName":    autorest.Encode("path", labAccountName),
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/register", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegisterSender sends the Register request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) RegisterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegisterResponder handles the response to the Register request. The method always
// closes the http.Response Body.
func (client LabsClient) RegisterResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update modify properties of labs.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// lab - represents a lab.
func (client LabsClient) Update(ctx context.Context, resourceGroupName string, labAccountName string, labName string, lab LabFragment) (result Lab, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LabsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, labAccountName, labName, lab)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.LabsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client LabsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, lab LabFragment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labAccountName":    autorest.Encode("path", labAccountName),
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", pathParameters),
		autorest.WithJSON(lab),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client LabsClient) UpdateResponder(resp *http.Response) (result Lab, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}