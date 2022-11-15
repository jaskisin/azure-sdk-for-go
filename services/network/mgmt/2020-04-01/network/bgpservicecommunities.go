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

// BgpServiceCommunitiesClient is the network Client
type BgpServiceCommunitiesClient struct {
	BaseClient
}

// NewBgpServiceCommunitiesClient creates an instance of the BgpServiceCommunitiesClient client.
func NewBgpServiceCommunitiesClient(subscriptionID string) BgpServiceCommunitiesClient {
	return NewBgpServiceCommunitiesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBgpServiceCommunitiesClientWithBaseURI creates an instance of the BgpServiceCommunitiesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewBgpServiceCommunitiesClientWithBaseURI(baseURI string, subscriptionID string) BgpServiceCommunitiesClient {
	return BgpServiceCommunitiesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets all the available bgp service communities.
func (client BgpServiceCommunitiesClient) List(ctx context.Context) (result BgpServiceCommunityListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BgpServiceCommunitiesClient.List")
		defer func() {
			sc := -1
			if result.bsclr.Response.Response != nil {
				sc = result.bsclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BgpServiceCommunitiesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.bsclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.BgpServiceCommunitiesClient", "List", resp, "Failure sending request")
		return
	}

	result.bsclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BgpServiceCommunitiesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.bsclr.hasNextLink() && result.bsclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client BgpServiceCommunitiesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/bgpServiceCommunities", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BgpServiceCommunitiesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BgpServiceCommunitiesClient) ListResponder(resp *http.Response) (result BgpServiceCommunityListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client BgpServiceCommunitiesClient) listNextResults(ctx context.Context, lastResults BgpServiceCommunityListResult) (result BgpServiceCommunityListResult, err error) {
	req, err := lastResults.bgpServiceCommunityListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.BgpServiceCommunitiesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.BgpServiceCommunitiesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BgpServiceCommunitiesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client BgpServiceCommunitiesClient) ListComplete(ctx context.Context) (result BgpServiceCommunityListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BgpServiceCommunitiesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}