package billing

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

// PermissionsClient is the billing client provides access to billing resources for Azure subscriptions.
type PermissionsClient struct {
	BaseClient
}

// NewPermissionsClient creates an instance of the PermissionsClient client.
func NewPermissionsClient(subscriptionID string) PermissionsClient {
	return NewPermissionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPermissionsClientWithBaseURI creates an instance of the PermissionsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPermissionsClientWithBaseURI(baseURI string, subscriptionID string) PermissionsClient {
	return PermissionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByBillingAccount lists all billing permissions for the caller under a billing account.
// Parameters:
// billingAccountName - billing Account Id.
func (client PermissionsClient) ListByBillingAccount(ctx context.Context, billingAccountName string) (result PermissionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByBillingAccountPreparer(ctx, billingAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByBillingAccount", resp, "Failure sending request")
		return
	}

	result, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByBillingAccount", resp, "Failure responding to request")
		return
	}

	return
}

// ListByBillingAccountPreparer prepares the ListByBillingAccount request.
func (client PermissionsClient) ListByBillingAccountPreparer(ctx context.Context, billingAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/providers/Microsoft.Billing/billingPermissions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingAccountSender sends the ListByBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client PermissionsClient) ListByBillingAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingAccountResponder handles the response to the ListByBillingAccount request. The method always
// closes the http.Response Body.
func (client PermissionsClient) ListByBillingAccountResponder(resp *http.Response) (result PermissionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBillingProfile lists all billingPermissions for the caller has for a billing account.
// Parameters:
// billingAccountName - billing Account Id.
// billingProfileName - billing Profile Id.
func (client PermissionsClient) ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result PermissionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListByBillingProfile")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByBillingProfilePreparer(ctx, billingAccountName, billingProfileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByBillingProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByBillingProfile", resp, "Failure sending request")
		return
	}

	result, err = client.ListByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByBillingProfile", resp, "Failure responding to request")
		return
	}

	return
}

// ListByBillingProfilePreparer prepares the ListByBillingProfile request.
func (client PermissionsClient) ListByBillingProfilePreparer(ctx context.Context, billingAccountName string, billingProfileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/providers/Microsoft.Billing/billingPermissions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingProfileSender sends the ListByBillingProfile request. The method will close the
// http.Response Body if it receives an error.
func (client PermissionsClient) ListByBillingProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingProfileResponder handles the response to the ListByBillingProfile request. The method always
// closes the http.Response Body.
func (client PermissionsClient) ListByBillingProfileResponder(resp *http.Response) (result PermissionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByCustomers lists all billing permissions for the caller under customer.
// Parameters:
// billingAccountName - billing Account Id.
// customerName - customer Id.
func (client PermissionsClient) ListByCustomers(ctx context.Context, billingAccountName string, customerName string) (result PermissionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListByCustomers")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByCustomersPreparer(ctx, billingAccountName, customerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByCustomers", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByCustomersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByCustomers", resp, "Failure sending request")
		return
	}

	result, err = client.ListByCustomersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByCustomers", resp, "Failure responding to request")
		return
	}

	return
}

// ListByCustomersPreparer prepares the ListByCustomers request.
func (client PermissionsClient) ListByCustomersPreparer(ctx context.Context, billingAccountName string, customerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"customerName":       autorest.Encode("path", customerName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/providers/Microsoft.Billing/billingPermissions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByCustomersSender sends the ListByCustomers request. The method will close the
// http.Response Body if it receives an error.
func (client PermissionsClient) ListByCustomersSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByCustomersResponder handles the response to the ListByCustomers request. The method always
// closes the http.Response Body.
func (client PermissionsClient) ListByCustomersResponder(resp *http.Response) (result PermissionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByInvoiceSections lists all billing permissions for the caller under invoice section.
// Parameters:
// billingAccountName - billing Account Id.
// invoiceSectionName - invoiceSection Id.
func (client PermissionsClient) ListByInvoiceSections(ctx context.Context, billingAccountName string, invoiceSectionName string) (result PermissionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListByInvoiceSections")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByInvoiceSectionsPreparer(ctx, billingAccountName, invoiceSectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByInvoiceSections", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByInvoiceSectionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByInvoiceSections", resp, "Failure sending request")
		return
	}

	result, err = client.ListByInvoiceSectionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.PermissionsClient", "ListByInvoiceSections", resp, "Failure responding to request")
		return
	}

	return
}

// ListByInvoiceSectionsPreparer prepares the ListByInvoiceSections request.
func (client PermissionsClient) ListByInvoiceSectionsPreparer(ctx context.Context, billingAccountName string, invoiceSectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Billing/billingPermissions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByInvoiceSectionsSender sends the ListByInvoiceSections request. The method will close the
// http.Response Body if it receives an error.
func (client PermissionsClient) ListByInvoiceSectionsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByInvoiceSectionsResponder handles the response to the ListByInvoiceSections request. The method always
// closes the http.Response Body.
func (client PermissionsClient) ListByInvoiceSectionsResponder(resp *http.Response) (result PermissionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}