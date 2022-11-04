//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

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

// DedicatedCloudNodesClient contains the methods for the DedicatedCloudNodes group.
// Don't use this type directly, use NewDedicatedCloudNodesClient() instead.
type DedicatedCloudNodesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDedicatedCloudNodesClient creates a new instance of DedicatedCloudNodesClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDedicatedCloudNodesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DedicatedCloudNodesClient, error) {
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
	client := &DedicatedCloudNodesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Returns dedicated cloud node by its name
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-04-01
// resourceGroupName - The name of the resource group
// referer - referer url
// dedicatedCloudNodeName - dedicated cloud node name
// dedicatedCloudNodeRequest - Create Dedicated Cloud Node request
// options - DedicatedCloudNodesClientBeginCreateOrUpdateOptions contains the optional parameters for the DedicatedCloudNodesClient.BeginCreateOrUpdate
// method.
func (client *DedicatedCloudNodesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, referer string, dedicatedCloudNodeName string, dedicatedCloudNodeRequest DedicatedCloudNode, options *DedicatedCloudNodesClientBeginCreateOrUpdateOptions) (*runtime.Poller[DedicatedCloudNodesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, referer, dedicatedCloudNodeName, dedicatedCloudNodeRequest, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DedicatedCloudNodesClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DedicatedCloudNodesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Returns dedicated cloud node by its name
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-04-01
func (client *DedicatedCloudNodesClient) createOrUpdate(ctx context.Context, resourceGroupName string, referer string, dedicatedCloudNodeName string, dedicatedCloudNodeRequest DedicatedCloudNode, options *DedicatedCloudNodesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, referer, dedicatedCloudNodeName, dedicatedCloudNodeRequest, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DedicatedCloudNodesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, referer string, dedicatedCloudNodeName string, dedicatedCloudNodeRequest DedicatedCloudNode, options *DedicatedCloudNodesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dedicatedCloudNodeName == "" {
		return nil, errors.New("parameter dedicatedCloudNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedCloudNodeName}", url.PathEscape(dedicatedCloudNodeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Referer"] = []string{referer}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, dedicatedCloudNodeRequest)
}

// Delete - Delete dedicated cloud node
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-04-01
// resourceGroupName - The name of the resource group
// dedicatedCloudNodeName - dedicated cloud node name
// options - DedicatedCloudNodesClientDeleteOptions contains the optional parameters for the DedicatedCloudNodesClient.Delete
// method.
func (client *DedicatedCloudNodesClient) Delete(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string, options *DedicatedCloudNodesClientDeleteOptions) (DedicatedCloudNodesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dedicatedCloudNodeName, options)
	if err != nil {
		return DedicatedCloudNodesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedCloudNodesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return DedicatedCloudNodesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DedicatedCloudNodesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DedicatedCloudNodesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string, options *DedicatedCloudNodesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dedicatedCloudNodeName == "" {
		return nil, errors.New("parameter dedicatedCloudNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedCloudNodeName}", url.PathEscape(dedicatedCloudNodeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns dedicated cloud node
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-04-01
// resourceGroupName - The name of the resource group
// dedicatedCloudNodeName - dedicated cloud node name
// options - DedicatedCloudNodesClientGetOptions contains the optional parameters for the DedicatedCloudNodesClient.Get method.
func (client *DedicatedCloudNodesClient) Get(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string, options *DedicatedCloudNodesClientGetOptions) (DedicatedCloudNodesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, dedicatedCloudNodeName, options)
	if err != nil {
		return DedicatedCloudNodesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedCloudNodesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DedicatedCloudNodesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DedicatedCloudNodesClient) getCreateRequest(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string, options *DedicatedCloudNodesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dedicatedCloudNodeName == "" {
		return nil, errors.New("parameter dedicatedCloudNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedCloudNodeName}", url.PathEscape(dedicatedCloudNodeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DedicatedCloudNodesClient) getHandleResponse(resp *http.Response) (DedicatedCloudNodesClientGetResponse, error) {
	result := DedicatedCloudNodesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedCloudNode); err != nil {
		return DedicatedCloudNodesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Returns list of dedicate cloud nodes within resource group
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-04-01
// resourceGroupName - The name of the resource group
// options - DedicatedCloudNodesClientListByResourceGroupOptions contains the optional parameters for the DedicatedCloudNodesClient.ListByResourceGroup
// method.
func (client *DedicatedCloudNodesClient) NewListByResourceGroupPager(resourceGroupName string, options *DedicatedCloudNodesClientListByResourceGroupOptions) *runtime.Pager[DedicatedCloudNodesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DedicatedCloudNodesClientListByResourceGroupResponse]{
		More: func(page DedicatedCloudNodesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DedicatedCloudNodesClientListByResourceGroupResponse) (DedicatedCloudNodesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DedicatedCloudNodesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DedicatedCloudNodesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DedicatedCloudNodesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DedicatedCloudNodesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DedicatedCloudNodesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DedicatedCloudNodesClient) listByResourceGroupHandleResponse(resp *http.Response) (DedicatedCloudNodesClientListByResourceGroupResponse, error) {
	result := DedicatedCloudNodesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedCloudNodeListResponse); err != nil {
		return DedicatedCloudNodesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Returns list of dedicate cloud nodes within subscription
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-04-01
// options - DedicatedCloudNodesClientListBySubscriptionOptions contains the optional parameters for the DedicatedCloudNodesClient.ListBySubscription
// method.
func (client *DedicatedCloudNodesClient) NewListBySubscriptionPager(options *DedicatedCloudNodesClientListBySubscriptionOptions) *runtime.Pager[DedicatedCloudNodesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[DedicatedCloudNodesClientListBySubscriptionResponse]{
		More: func(page DedicatedCloudNodesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DedicatedCloudNodesClientListBySubscriptionResponse) (DedicatedCloudNodesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DedicatedCloudNodesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DedicatedCloudNodesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DedicatedCloudNodesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DedicatedCloudNodesClient) listBySubscriptionCreateRequest(ctx context.Context, options *DedicatedCloudNodesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DedicatedCloudNodesClient) listBySubscriptionHandleResponse(resp *http.Response) (DedicatedCloudNodesClientListBySubscriptionResponse, error) {
	result := DedicatedCloudNodesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedCloudNodeListResponse); err != nil {
		return DedicatedCloudNodesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Patches dedicated node properties
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-04-01
// resourceGroupName - The name of the resource group
// dedicatedCloudNodeName - dedicated cloud node name
// dedicatedCloudNodeRequest - Patch Dedicated Cloud Node request
// options - DedicatedCloudNodesClientUpdateOptions contains the optional parameters for the DedicatedCloudNodesClient.Update
// method.
func (client *DedicatedCloudNodesClient) Update(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string, dedicatedCloudNodeRequest PatchPayload, options *DedicatedCloudNodesClientUpdateOptions) (DedicatedCloudNodesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dedicatedCloudNodeName, dedicatedCloudNodeRequest, options)
	if err != nil {
		return DedicatedCloudNodesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedCloudNodesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DedicatedCloudNodesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *DedicatedCloudNodesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string, dedicatedCloudNodeRequest PatchPayload, options *DedicatedCloudNodesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dedicatedCloudNodeName == "" {
		return nil, errors.New("parameter dedicatedCloudNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedCloudNodeName}", url.PathEscape(dedicatedCloudNodeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, dedicatedCloudNodeRequest)
}

// updateHandleResponse handles the Update response.
func (client *DedicatedCloudNodesClient) updateHandleResponse(resp *http.Response) (DedicatedCloudNodesClientUpdateResponse, error) {
	result := DedicatedCloudNodesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedCloudNode); err != nil {
		return DedicatedCloudNodesClientUpdateResponse{}, err
	}
	return result, nil
}