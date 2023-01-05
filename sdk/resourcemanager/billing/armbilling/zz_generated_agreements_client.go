//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AgreementsClient contains the methods for the Agreements group.
// Don't use this type directly, use NewAgreementsClient() instead.
type AgreementsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewAgreementsClient creates a new instance of AgreementsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAgreementsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AgreementsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AgreementsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Get - Gets an agreement by ID.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// agreementName - The ID that uniquely identifies an agreement.
// options - AgreementsClientGetOptions contains the optional parameters for the AgreementsClient.Get method.
func (client *AgreementsClient) Get(ctx context.Context, billingAccountName string, agreementName string, options *AgreementsClientGetOptions) (AgreementsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, billingAccountName, agreementName, options)
	if err != nil {
		return AgreementsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AgreementsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgreementsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AgreementsClient) getCreateRequest(ctx context.Context, billingAccountName string, agreementName string, options *AgreementsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/agreements/{agreementName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if agreementName == "" {
		return nil, errors.New("parameter agreementName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agreementName}", url.PathEscape(agreementName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AgreementsClient) getHandleResponse(resp *http.Response) (AgreementsClientGetResponse, error) {
	result := AgreementsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Agreement); err != nil {
		return AgreementsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByBillingAccountPager - Lists the agreements for a billing account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// options - AgreementsClientListByBillingAccountOptions contains the optional parameters for the AgreementsClient.ListByBillingAccount
// method.
func (client *AgreementsClient) NewListByBillingAccountPager(billingAccountName string, options *AgreementsClientListByBillingAccountOptions) *runtime.Pager[AgreementsClientListByBillingAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[AgreementsClientListByBillingAccountResponse]{
		More: func(page AgreementsClientListByBillingAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AgreementsClientListByBillingAccountResponse) (AgreementsClientListByBillingAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByBillingAccountCreateRequest(ctx, billingAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AgreementsClientListByBillingAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AgreementsClientListByBillingAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AgreementsClientListByBillingAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByBillingAccountHandleResponse(resp)
		},
	})
}

// listByBillingAccountCreateRequest creates the ListByBillingAccount request.
func (client *AgreementsClient) listByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, options *AgreementsClientListByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/agreements"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingAccountHandleResponse handles the ListByBillingAccount response.
func (client *AgreementsClient) listByBillingAccountHandleResponse(resp *http.Response) (AgreementsClientListByBillingAccountResponse, error) {
	result := AgreementsClientListByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgreementListResult); err != nil {
		return AgreementsClientListByBillingAccountResponse{}, err
	}
	return result, nil
}