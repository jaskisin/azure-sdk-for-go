package network

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PrivateDNSZoneGroupsClient is the network Client
type PrivateDNSZoneGroupsClient struct {
	BaseClient
}

// NewPrivateDNSZoneGroupsClient creates an instance of the PrivateDNSZoneGroupsClient client.
func NewPrivateDNSZoneGroupsClient(subscriptionID string) PrivateDNSZoneGroupsClient {
	return NewPrivateDNSZoneGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateDNSZoneGroupsClientWithBaseURI creates an instance of the PrivateDNSZoneGroupsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewPrivateDNSZoneGroupsClientWithBaseURI(baseURI string, subscriptionID string) PrivateDNSZoneGroupsClient {
	return PrivateDNSZoneGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a private dns zone group in the specified private endpoint.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateEndpointName - the name of the private endpoint.
// privateDNSZoneGroupName - the name of the private dns zone group.
// parameters - parameters supplied to the create or update private dns zone group operation.
func (client PrivateDNSZoneGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, parameters PrivateDNSZoneGroup) (result PrivateDNSZoneGroupsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateDNSZoneGroupsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, privateEndpointName, privateDNSZoneGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateDNSZoneGroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateDNSZoneGroupsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PrivateDNSZoneGroupsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, parameters PrivateDNSZoneGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateDnsZoneGroupName": autorest.Encode("path", privateDNSZoneGroupName),
		"privateEndpointName":     autorest.Encode("path", privateEndpointName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateDNSZoneGroupsClient) CreateOrUpdateSender(req *http.Request) (future PrivateDNSZoneGroupsCreateOrUpdateFuture, err error) {
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
func (client PrivateDNSZoneGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result PrivateDNSZoneGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified private dns zone group.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateEndpointName - the name of the private endpoint.
// privateDNSZoneGroupName - the name of the private dns zone group.
func (client PrivateDNSZoneGroupsClient) Delete(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string) (result PrivateDNSZoneGroupsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateDNSZoneGroupsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, privateEndpointName, privateDNSZoneGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateDNSZoneGroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateDNSZoneGroupsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PrivateDNSZoneGroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateDnsZoneGroupName": autorest.Encode("path", privateDNSZoneGroupName),
		"privateEndpointName":     autorest.Encode("path", privateEndpointName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateDNSZoneGroupsClient) DeleteSender(req *http.Request) (future PrivateDNSZoneGroupsDeleteFuture, err error) {
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
func (client PrivateDNSZoneGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the private dns zone group resource by specified private dns zone group name.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateEndpointName - the name of the private endpoint.
// privateDNSZoneGroupName - the name of the private dns zone group.
func (client PrivateDNSZoneGroupsClient) Get(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string) (result PrivateDNSZoneGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateDNSZoneGroupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, privateEndpointName, privateDNSZoneGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateDNSZoneGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PrivateDNSZoneGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateDNSZoneGroupsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateDNSZoneGroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateDnsZoneGroupName": autorest.Encode("path", privateDNSZoneGroupName),
		"privateEndpointName":     autorest.Encode("path", privateEndpointName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateDNSZoneGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateDNSZoneGroupsClient) GetResponder(resp *http.Response) (result PrivateDNSZoneGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all private dns zone groups in a private endpoint.
// Parameters:
// privateEndpointName - the name of the private endpoint.
// resourceGroupName - the name of the resource group.
func (client PrivateDNSZoneGroupsClient) List(ctx context.Context, privateEndpointName string, resourceGroupName string) (result PrivateDNSZoneGroupListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateDNSZoneGroupsClient.List")
		defer func() {
			sc := -1
			if result.pdzglr.Response.Response != nil {
				sc = result.pdzglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, privateEndpointName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateDNSZoneGroupsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pdzglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PrivateDNSZoneGroupsClient", "List", resp, "Failure sending request")
		return
	}

	result.pdzglr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateDNSZoneGroupsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.pdzglr.hasNextLink() && result.pdzglr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client PrivateDNSZoneGroupsClient) ListPreparer(ctx context.Context, privateEndpointName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointName": autorest.Encode("path", privateEndpointName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateDNSZoneGroupsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PrivateDNSZoneGroupsClient) ListResponder(resp *http.Response) (result PrivateDNSZoneGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PrivateDNSZoneGroupsClient) listNextResults(ctx context.Context, lastResults PrivateDNSZoneGroupListResult) (result PrivateDNSZoneGroupListResult, err error) {
	req, err := lastResults.privateDNSZoneGroupListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PrivateDNSZoneGroupsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PrivateDNSZoneGroupsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateDNSZoneGroupsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateDNSZoneGroupsClient) ListComplete(ctx context.Context, privateEndpointName string, resourceGroupName string) (result PrivateDNSZoneGroupListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateDNSZoneGroupsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, privateEndpointName, resourceGroupName)
	return
}