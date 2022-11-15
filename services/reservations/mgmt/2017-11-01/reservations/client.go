// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
//
// Package reservations implements the Azure ARM Reservations service API version 2017-11-01.
//
// This API describe Azure Reservation
package reservations

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

const (
	// DefaultBaseURI is the default URI used for the service Reservations
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Reservations.
type BaseClient struct {
	autorest.Client
	BaseURI string
}

// New creates an instance of the BaseClient client.
func New() BaseClient {
	return NewWithBaseURI(DefaultBaseURI)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string) BaseClient {
	return BaseClient{
		Client:  autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI: baseURI,
	}
}

// GetAppliedReservationList get applicable `Reservation`s that are applied to this subscription.
// Parameters:
// subscriptionID - id of the subscription
func (client BaseClient) GetAppliedReservationList(ctx context.Context, subscriptionID string) (result AppliedReservations, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetAppliedReservationList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAppliedReservationListPreparer(ctx, subscriptionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.BaseClient", "GetAppliedReservationList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAppliedReservationListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "reservations.BaseClient", "GetAppliedReservationList", resp, "Failure sending request")
		return
	}

	result, err = client.GetAppliedReservationListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.BaseClient", "GetAppliedReservationList", resp, "Failure responding to request")
		return
	}

	return
}

// GetAppliedReservationListPreparer prepares the GetAppliedReservationList request.
func (client BaseClient) GetAppliedReservationListPreparer(ctx context.Context, subscriptionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/appliedReservations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAppliedReservationListSender sends the GetAppliedReservationList request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetAppliedReservationListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetAppliedReservationListResponder handles the response to the GetAppliedReservationList request. The method always
// closes the http.Response Body.
func (client BaseClient) GetAppliedReservationListResponder(resp *http.Response) (result AppliedReservations, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetCatalog sends the get catalog request.
// Parameters:
// subscriptionID - id of the subscription
func (client BaseClient) GetCatalog(ctx context.Context, subscriptionID string) (result ListCatalog, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetCatalog")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetCatalogPreparer(ctx, subscriptionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.BaseClient", "GetCatalog", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCatalogSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "reservations.BaseClient", "GetCatalog", resp, "Failure sending request")
		return
	}

	result, err = client.GetCatalogResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.BaseClient", "GetCatalog", resp, "Failure responding to request")
		return
	}

	return
}

// GetCatalogPreparer prepares the GetCatalog request.
func (client BaseClient) GetCatalogPreparer(ctx context.Context, subscriptionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/catalogs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetCatalogSender sends the GetCatalog request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetCatalogSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetCatalogResponder handles the response to the GetCatalog request. The method always
// closes the http.Response Body.
func (client BaseClient) GetCatalogResponder(resp *http.Response) (result ListCatalog, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}