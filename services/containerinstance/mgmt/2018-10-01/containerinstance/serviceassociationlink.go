package containerinstance

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

// ServiceAssociationLinkClient is the client for the ServiceAssociationLink methods of the Containerinstance service.
type ServiceAssociationLinkClient struct {
	BaseClient
}

// NewServiceAssociationLinkClient creates an instance of the ServiceAssociationLinkClient client.
func NewServiceAssociationLinkClient(subscriptionID string) ServiceAssociationLinkClient {
	return NewServiceAssociationLinkClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceAssociationLinkClientWithBaseURI creates an instance of the ServiceAssociationLinkClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewServiceAssociationLinkClientWithBaseURI(baseURI string, subscriptionID string) ServiceAssociationLinkClient {
	return ServiceAssociationLinkClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete delete the container instance service association link for the subnet. This operation unblocks user from
// deleting subnet.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkName - the name of the virtual network.
// subnetName - the name of the subnet.
func (client ServiceAssociationLinkClient) Delete(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceAssociationLinkClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, virtualNetworkName, subnetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ServiceAssociationLinkClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "containerinstance.ServiceAssociationLinkClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ServiceAssociationLinkClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ServiceAssociationLinkClient) DeletePreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subnetName":         autorest.Encode("path", subnetName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2018-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/providers/Microsoft.ContainerInstance/serviceAssociationLinks/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceAssociationLinkClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ServiceAssociationLinkClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}