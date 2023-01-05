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
	"strconv"
	"strings"
)

// PeriodsClient contains the methods for the BillingPeriods group.
// Don't use this type directly, use NewPeriodsClient() instead.
type PeriodsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPeriodsClient creates a new instance of PeriodsClient with the specified values.
// subscriptionID - The ID that uniquely identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPeriodsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PeriodsClient, error) {
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
	client := &PeriodsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets a named billing period. This is only supported for Azure Web-Direct subscriptions. Other subscription types
// which were not purchased directly through the Azure web portal are not supported
// through this preview API.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-03-01-preview
// billingPeriodName - The name of a BillingPeriod resource.
// options - PeriodsClientGetOptions contains the optional parameters for the PeriodsClient.Get method.
func (client *PeriodsClient) Get(ctx context.Context, billingPeriodName string, options *PeriodsClientGetOptions) (PeriodsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, billingPeriodName, options)
	if err != nil {
		return PeriodsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PeriodsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PeriodsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PeriodsClient) getCreateRequest(ctx context.Context, billingPeriodName string, options *PeriodsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if billingPeriodName == "" {
		return nil, errors.New("parameter billingPeriodName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingPeriodName}", url.PathEscape(billingPeriodName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PeriodsClient) getHandleResponse(resp *http.Response) (PeriodsClientGetResponse, error) {
	result := PeriodsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Period); err != nil {
		return PeriodsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the available billing periods for a subscription in reverse chronological order. This is only supported
// for Azure Web-Direct subscriptions. Other subscription types which were not purchased
// directly through the Azure web portal are not supported through this preview API.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-03-01-preview
// options - PeriodsClientListOptions contains the optional parameters for the PeriodsClient.List method.
func (client *PeriodsClient) NewListPager(options *PeriodsClientListOptions) *runtime.Pager[PeriodsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PeriodsClientListResponse]{
		More: func(page PeriodsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PeriodsClientListResponse) (PeriodsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PeriodsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PeriodsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PeriodsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PeriodsClient) listCreateRequest(ctx context.Context, options *PeriodsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PeriodsClient) listHandleResponse(resp *http.Response) (PeriodsClientListResponse, error) {
	result := PeriodsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PeriodsListResult); err != nil {
		return PeriodsClientListResponse{}, err
	}
	return result, nil
}