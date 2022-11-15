package sql

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

// ServerAdvancedThreatProtectionSettingsClient is the the Azure SQL Database management API provides a RESTful set of
// web services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type ServerAdvancedThreatProtectionSettingsClient struct {
	BaseClient
}

// NewServerAdvancedThreatProtectionSettingsClient creates an instance of the
// ServerAdvancedThreatProtectionSettingsClient client.
func NewServerAdvancedThreatProtectionSettingsClient(subscriptionID string) ServerAdvancedThreatProtectionSettingsClient {
	return NewServerAdvancedThreatProtectionSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServerAdvancedThreatProtectionSettingsClientWithBaseURI creates an instance of the
// ServerAdvancedThreatProtectionSettingsClient client using a custom endpoint.  Use this when interacting with an
// Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewServerAdvancedThreatProtectionSettingsClientWithBaseURI(baseURI string, subscriptionID string) ServerAdvancedThreatProtectionSettingsClient {
	return ServerAdvancedThreatProtectionSettingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an Advanced Threat Protection state.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// parameters - the server Advanced Threat Protection state.
func (client ServerAdvancedThreatProtectionSettingsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters ServerAdvancedThreatProtection) (result ServerAdvancedThreatProtectionSettingsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerAdvancedThreatProtectionSettingsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerAdvancedThreatProtectionSettingsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerAdvancedThreatProtectionSettingsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ServerAdvancedThreatProtectionSettingsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, parameters ServerAdvancedThreatProtection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"advancedThreatProtectionName": autorest.Encode("path", "Default"),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"serverName":                   autorest.Encode("path", serverName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advancedThreatProtectionSettings/{advancedThreatProtectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ServerAdvancedThreatProtectionSettingsClient) CreateOrUpdateSender(req *http.Request) (future ServerAdvancedThreatProtectionSettingsCreateOrUpdateFuture, err error) {
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
func (client ServerAdvancedThreatProtectionSettingsClient) CreateOrUpdateResponder(resp *http.Response) (result ServerAdvancedThreatProtection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get a server's Advanced Threat Protection state.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
func (client ServerAdvancedThreatProtectionSettingsClient) Get(ctx context.Context, resourceGroupName string, serverName string) (result ServerAdvancedThreatProtection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerAdvancedThreatProtectionSettingsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerAdvancedThreatProtectionSettingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ServerAdvancedThreatProtectionSettingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerAdvancedThreatProtectionSettingsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServerAdvancedThreatProtectionSettingsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"advancedThreatProtectionName": autorest.Encode("path", "Default"),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"serverName":                   autorest.Encode("path", serverName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advancedThreatProtectionSettings/{advancedThreatProtectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServerAdvancedThreatProtectionSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServerAdvancedThreatProtectionSettingsClient) GetResponder(resp *http.Response) (result ServerAdvancedThreatProtection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer get a list of the server's Advanced Threat Protection states.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
func (client ServerAdvancedThreatProtectionSettingsClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result LogicalServerAdvancedThreatProtectionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerAdvancedThreatProtectionSettingsClient.ListByServer")
		defer func() {
			sc := -1
			if result.lsatplr.Response.Response != nil {
				sc = result.lsatplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByServerNextResults
	req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerAdvancedThreatProtectionSettingsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.lsatplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ServerAdvancedThreatProtectionSettingsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result.lsatplr, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerAdvancedThreatProtectionSettingsClient", "ListByServer", resp, "Failure responding to request")
		return
	}
	if result.lsatplr.hasNextLink() && result.lsatplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client ServerAdvancedThreatProtectionSettingsClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advancedThreatProtectionSettings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client ServerAdvancedThreatProtectionSettingsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client ServerAdvancedThreatProtectionSettingsClient) ListByServerResponder(resp *http.Response) (result LogicalServerAdvancedThreatProtectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServerNextResults retrieves the next set of results, if any.
func (client ServerAdvancedThreatProtectionSettingsClient) listByServerNextResults(ctx context.Context, lastResults LogicalServerAdvancedThreatProtectionListResult) (result LogicalServerAdvancedThreatProtectionListResult, err error) {
	req, err := lastResults.logicalServerAdvancedThreatProtectionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ServerAdvancedThreatProtectionSettingsClient", "listByServerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ServerAdvancedThreatProtectionSettingsClient", "listByServerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerAdvancedThreatProtectionSettingsClient", "listByServerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServerComplete enumerates all values, automatically crossing page boundaries as required.
func (client ServerAdvancedThreatProtectionSettingsClient) ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result LogicalServerAdvancedThreatProtectionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerAdvancedThreatProtectionSettingsClient.ListByServer")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByServer(ctx, resourceGroupName, serverName)
	return
}