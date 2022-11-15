package apimanagement

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

// UserSubscriptionsClient is the apiManagement Client
type UserSubscriptionsClient struct {
	BaseClient
}

// NewUserSubscriptionsClient creates an instance of the UserSubscriptionsClient client.
func NewUserSubscriptionsClient(subscriptionID string) UserSubscriptionsClient {
	return NewUserSubscriptionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUserSubscriptionsClientWithBaseURI creates an instance of the UserSubscriptionsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewUserSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) UserSubscriptionsClient {
	return UserSubscriptionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByUsers lists the collection of subscriptions of the specified user.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// UID - user identifier. Must be unique in the current API Management service instance.
// filter - | Field        | Supported operators    | Supported functions                         |
// |--------------|------------------------|---------------------------------------------|
// | id           | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | name         | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | stateComment | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | userId       | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | productId    | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | state        | eq                     |                                             |
// top - number of records to return.
// skip - number of records to skip.
func (client UserSubscriptionsClient) ListByUsers(ctx context.Context, resourceGroupName string, serviceName string, UID string, filter string, top *int32, skip *int32) (result SubscriptionCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserSubscriptionsClient.ListByUsers")
		defer func() {
			sc := -1
			if result.sc.Response.Response != nil {
				sc = result.sc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: UID,
			Constraints: []validation.Constraint{{Target: "UID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "UID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "UID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.UserSubscriptionsClient", "ListByUsers", err.Error())
	}

	result.fn = client.listByUsersNextResults
	req, err := client.ListByUsersPreparer(ctx, resourceGroupName, serviceName, UID, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionsClient", "ListByUsers", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByUsersSender(req)
	if err != nil {
		result.sc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionsClient", "ListByUsers", resp, "Failure sending request")
		return
	}

	result.sc, err = client.ListByUsersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionsClient", "ListByUsers", resp, "Failure responding to request")
		return
	}
	if result.sc.hasNextLink() && result.sc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByUsersPreparer prepares the ListByUsers request.
func (client UserSubscriptionsClient) ListByUsersPreparer(ctx context.Context, resourceGroupName string, serviceName string, UID string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", UID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{uid}/subscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByUsersSender sends the ListByUsers request. The method will close the
// http.Response Body if it receives an error.
func (client UserSubscriptionsClient) ListByUsersSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByUsersResponder handles the response to the ListByUsers request. The method always
// closes the http.Response Body.
func (client UserSubscriptionsClient) ListByUsersResponder(resp *http.Response) (result SubscriptionCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByUsersNextResults retrieves the next set of results, if any.
func (client UserSubscriptionsClient) listByUsersNextResults(ctx context.Context, lastResults SubscriptionCollection) (result SubscriptionCollection, err error) {
	req, err := lastResults.subscriptionCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionsClient", "listByUsersNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByUsersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionsClient", "listByUsersNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByUsersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionsClient", "listByUsersNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByUsersComplete enumerates all values, automatically crossing page boundaries as required.
func (client UserSubscriptionsClient) ListByUsersComplete(ctx context.Context, resourceGroupName string, serviceName string, UID string, filter string, top *int32, skip *int32) (result SubscriptionCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserSubscriptionsClient.ListByUsers")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByUsers(ctx, resourceGroupName, serviceName, UID, filter, top, skip)
	return
}