//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armm365securityandcompliance

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

// PrivateLinkServicesForMIPPolicySyncClient contains the methods for the PrivateLinkServicesForMIPPolicySync group.
// Don't use this type directly, use NewPrivateLinkServicesForMIPPolicySyncClient() instead.
type PrivateLinkServicesForMIPPolicySyncClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkServicesForMIPPolicySyncClient creates a new instance of PrivateLinkServicesForMIPPolicySyncClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkServicesForMIPPolicySyncClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkServicesForMIPPolicySyncClient, error) {
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
	client := &PrivateLinkServicesForMIPPolicySyncClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update the metadata of a privateLinkServicesForMIPPolicySync instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// privateLinkServicesForMIPPolicySyncDescription - The service instance metadata.
// options - PrivateLinkServicesForMIPPolicySyncClientBeginCreateOrUpdateOptions contains the optional parameters for the
// PrivateLinkServicesForMIPPolicySyncClient.BeginCreateOrUpdate method.
func (client *PrivateLinkServicesForMIPPolicySyncClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForMIPPolicySyncDescription PrivateLinkServicesForMIPPolicySyncDescription, options *PrivateLinkServicesForMIPPolicySyncClientBeginCreateOrUpdateOptions) (*runtime.Poller[PrivateLinkServicesForMIPPolicySyncClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, privateLinkServicesForMIPPolicySyncDescription, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[PrivateLinkServicesForMIPPolicySyncClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[PrivateLinkServicesForMIPPolicySyncClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update the metadata of a privateLinkServicesForMIPPolicySync instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
func (client *PrivateLinkServicesForMIPPolicySyncClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForMIPPolicySyncDescription PrivateLinkServicesForMIPPolicySyncDescription, options *PrivateLinkServicesForMIPPolicySyncClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, privateLinkServicesForMIPPolicySyncDescription, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateLinkServicesForMIPPolicySyncClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForMIPPolicySyncDescription PrivateLinkServicesForMIPPolicySyncDescription, options *PrivateLinkServicesForMIPPolicySyncClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, privateLinkServicesForMIPPolicySyncDescription)
}

// BeginDelete - Delete a service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// options - PrivateLinkServicesForMIPPolicySyncClientBeginDeleteOptions contains the optional parameters for the PrivateLinkServicesForMIPPolicySyncClient.BeginDelete
// method.
func (client *PrivateLinkServicesForMIPPolicySyncClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForMIPPolicySyncClientBeginDeleteOptions) (*runtime.Poller[PrivateLinkServicesForMIPPolicySyncClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[PrivateLinkServicesForMIPPolicySyncClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[PrivateLinkServicesForMIPPolicySyncClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
func (client *PrivateLinkServicesForMIPPolicySyncClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForMIPPolicySyncClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateLinkServicesForMIPPolicySyncClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForMIPPolicySyncClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the metadata of a privateLinkServicesForMIPPolicySync resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// options - PrivateLinkServicesForMIPPolicySyncClientGetOptions contains the optional parameters for the PrivateLinkServicesForMIPPolicySyncClient.Get
// method.
func (client *PrivateLinkServicesForMIPPolicySyncClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForMIPPolicySyncClientGetOptions) (PrivateLinkServicesForMIPPolicySyncClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return PrivateLinkServicesForMIPPolicySyncClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkServicesForMIPPolicySyncClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkServicesForMIPPolicySyncClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkServicesForMIPPolicySyncClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForMIPPolicySyncClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkServicesForMIPPolicySyncClient) getHandleResponse(resp *http.Response) (PrivateLinkServicesForMIPPolicySyncClientGetResponse, error) {
	result := PrivateLinkServicesForMIPPolicySyncClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkServicesForMIPPolicySyncDescription); err != nil {
		return PrivateLinkServicesForMIPPolicySyncClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all the privateLinkServicesForMIPPolicySync instances in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
// options - PrivateLinkServicesForMIPPolicySyncClientListOptions contains the optional parameters for the PrivateLinkServicesForMIPPolicySyncClient.List
// method.
func (client *PrivateLinkServicesForMIPPolicySyncClient) NewListPager(options *PrivateLinkServicesForMIPPolicySyncClientListOptions) *runtime.Pager[PrivateLinkServicesForMIPPolicySyncClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkServicesForMIPPolicySyncClientListResponse]{
		More: func(page PrivateLinkServicesForMIPPolicySyncClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkServicesForMIPPolicySyncClientListResponse) (PrivateLinkServicesForMIPPolicySyncClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateLinkServicesForMIPPolicySyncClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrivateLinkServicesForMIPPolicySyncClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkServicesForMIPPolicySyncClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PrivateLinkServicesForMIPPolicySyncClient) listCreateRequest(ctx context.Context, options *PrivateLinkServicesForMIPPolicySyncClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateLinkServicesForMIPPolicySyncClient) listHandleResponse(resp *http.Response) (PrivateLinkServicesForMIPPolicySyncClientListResponse, error) {
	result := PrivateLinkServicesForMIPPolicySyncClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkServicesForMIPPolicySyncDescriptionListResult); err != nil {
		return PrivateLinkServicesForMIPPolicySyncClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get all the service instances in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// options - PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupOptions contains the optional parameters for the
// PrivateLinkServicesForMIPPolicySyncClient.ListByResourceGroup method.
func (client *PrivateLinkServicesForMIPPolicySyncClient) NewListByResourceGroupPager(resourceGroupName string, options *PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupOptions) *runtime.Pager[PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupResponse]{
		More: func(page PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupResponse) (PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PrivateLinkServicesForMIPPolicySyncClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync"
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
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PrivateLinkServicesForMIPPolicySyncClient) listByResourceGroupHandleResponse(resp *http.Response) (PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupResponse, error) {
	result := PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkServicesForMIPPolicySyncDescriptionListResult); err != nil {
		return PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update the metadata of a privateLinkServicesForMIPPolicySync instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// servicePatchDescription - The service instance metadata and security metadata.
// options - PrivateLinkServicesForMIPPolicySyncClientBeginUpdateOptions contains the optional parameters for the PrivateLinkServicesForMIPPolicySyncClient.BeginUpdate
// method.
func (client *PrivateLinkServicesForMIPPolicySyncClient) BeginUpdate(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription, options *PrivateLinkServicesForMIPPolicySyncClientBeginUpdateOptions) (*runtime.Poller[PrivateLinkServicesForMIPPolicySyncClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, resourceName, servicePatchDescription, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[PrivateLinkServicesForMIPPolicySyncClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[PrivateLinkServicesForMIPPolicySyncClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update the metadata of a privateLinkServicesForMIPPolicySync instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-25-preview
func (client *PrivateLinkServicesForMIPPolicySyncClient) update(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription, options *PrivateLinkServicesForMIPPolicySyncClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, servicePatchDescription, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *PrivateLinkServicesForMIPPolicySyncClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription, options *PrivateLinkServicesForMIPPolicySyncClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, servicePatchDescription)
}