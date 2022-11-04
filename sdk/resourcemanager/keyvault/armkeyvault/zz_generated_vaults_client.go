//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

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

// VaultsClient contains the methods for the Vaults group.
// Don't use this type directly, use NewVaultsClient() instead.
type VaultsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVaultsClient creates a new instance of VaultsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVaultsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VaultsClient, error) {
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
	client := &VaultsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckNameAvailability - Checks that the vault name is valid and is not already in use.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// vaultName - The name of the vault.
// options - VaultsClientCheckNameAvailabilityOptions contains the optional parameters for the VaultsClient.CheckNameAvailability
// method.
func (client *VaultsClient) CheckNameAvailability(ctx context.Context, vaultName VaultCheckNameAvailabilityParameters, options *VaultsClientCheckNameAvailabilityOptions) (VaultsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, vaultName, options)
	if err != nil {
		return VaultsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VaultsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VaultsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *VaultsClient) checkNameAvailabilityCreateRequest(ctx context.Context, vaultName VaultCheckNameAvailabilityParameters, options *VaultsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, vaultName)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *VaultsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (VaultsClientCheckNameAvailabilityResponse, error) {
	result := VaultsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResult); err != nil {
		return VaultsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Create or update a key vault in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - The name of the Resource Group to which the server belongs.
// vaultName - Name of the vault
// parameters - Parameters to create or update the vault
// options - VaultsClientBeginCreateOrUpdateOptions contains the optional parameters for the VaultsClient.BeginCreateOrUpdate
// method.
func (client *VaultsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, parameters VaultCreateOrUpdateParameters, options *VaultsClientBeginCreateOrUpdateOptions) (*runtime.Poller[VaultsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, vaultName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VaultsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[VaultsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a key vault in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
func (client *VaultsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, parameters VaultCreateOrUpdateParameters, options *VaultsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vaultName, parameters, options)
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
func (client *VaultsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, parameters VaultCreateOrUpdateParameters, options *VaultsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Delete - Deletes the specified Azure key vault.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - The name of the Resource Group to which the vault belongs.
// vaultName - The name of the vault to delete
// options - VaultsClientDeleteOptions contains the optional parameters for the VaultsClient.Delete method.
func (client *VaultsClient) Delete(ctx context.Context, resourceGroupName string, vaultName string, options *VaultsClientDeleteOptions) (VaultsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return VaultsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VaultsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return VaultsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return VaultsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VaultsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *VaultsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified Azure key vault.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - The name of the Resource Group to which the vault belongs.
// vaultName - The name of the vault.
// options - VaultsClientGetOptions contains the optional parameters for the VaultsClient.Get method.
func (client *VaultsClient) Get(ctx context.Context, resourceGroupName string, vaultName string, options *VaultsClientGetOptions) (VaultsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return VaultsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VaultsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VaultsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VaultsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *VaultsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VaultsClient) getHandleResponse(resp *http.Response) (VaultsClientGetResponse, error) {
	result := VaultsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Vault); err != nil {
		return VaultsClientGetResponse{}, err
	}
	return result, nil
}

// GetDeleted - Gets the deleted Azure key vault.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// vaultName - The name of the vault.
// location - The location of the deleted vault.
// options - VaultsClientGetDeletedOptions contains the optional parameters for the VaultsClient.GetDeleted method.
func (client *VaultsClient) GetDeleted(ctx context.Context, vaultName string, location string, options *VaultsClientGetDeletedOptions) (VaultsClientGetDeletedResponse, error) {
	req, err := client.getDeletedCreateRequest(ctx, vaultName, location, options)
	if err != nil {
		return VaultsClientGetDeletedResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VaultsClientGetDeletedResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VaultsClientGetDeletedResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDeletedHandleResponse(resp)
}

// getDeletedCreateRequest creates the GetDeleted request.
func (client *VaultsClient) getDeletedCreateRequest(ctx context.Context, vaultName string, location string, options *VaultsClientGetDeletedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/locations/{location}/deletedVaults/{vaultName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDeletedHandleResponse handles the GetDeleted response.
func (client *VaultsClient) getDeletedHandleResponse(resp *http.Response) (VaultsClientGetDeletedResponse, error) {
	result := VaultsClientGetDeletedResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedVault); err != nil {
		return VaultsClientGetDeletedResponse{}, err
	}
	return result, nil
}

// NewListPager - The List operation gets information about the vaults associated with the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// options - VaultsClientListOptions contains the optional parameters for the VaultsClient.List method.
func (client *VaultsClient) NewListPager(options *VaultsClientListOptions) *runtime.Pager[VaultsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VaultsClientListResponse]{
		More: func(page VaultsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VaultsClientListResponse) (VaultsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VaultsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VaultsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VaultsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VaultsClient) listCreateRequest(ctx context.Context, options *VaultsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("$filter", "resourceType eq 'Microsoft.KeyVault/vaults'")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2015-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VaultsClient) listHandleResponse(resp *http.Response) (VaultsClientListResponse, error) {
	result := VaultsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceListResult); err != nil {
		return VaultsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - The List operation gets information about the vaults associated with the subscription and
// within the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - The name of the Resource Group to which the vault belongs.
// options - VaultsClientListByResourceGroupOptions contains the optional parameters for the VaultsClient.ListByResourceGroup
// method.
func (client *VaultsClient) NewListByResourceGroupPager(resourceGroupName string, options *VaultsClientListByResourceGroupOptions) *runtime.Pager[VaultsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[VaultsClientListByResourceGroupResponse]{
		More: func(page VaultsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VaultsClientListByResourceGroupResponse) (VaultsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VaultsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VaultsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VaultsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VaultsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VaultsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VaultsClient) listByResourceGroupHandleResponse(resp *http.Response) (VaultsClientListByResourceGroupResponse, error) {
	result := VaultsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultListResult); err != nil {
		return VaultsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - The List operation gets information about the vaults associated with the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// options - VaultsClientListBySubscriptionOptions contains the optional parameters for the VaultsClient.ListBySubscription
// method.
func (client *VaultsClient) NewListBySubscriptionPager(options *VaultsClientListBySubscriptionOptions) *runtime.Pager[VaultsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[VaultsClientListBySubscriptionResponse]{
		More: func(page VaultsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VaultsClientListBySubscriptionResponse) (VaultsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VaultsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VaultsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VaultsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *VaultsClient) listBySubscriptionCreateRequest(ctx context.Context, options *VaultsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/vaults"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *VaultsClient) listBySubscriptionHandleResponse(resp *http.Response) (VaultsClientListBySubscriptionResponse, error) {
	result := VaultsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultListResult); err != nil {
		return VaultsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// NewListDeletedPager - Gets information about the deleted vaults in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// options - VaultsClientListDeletedOptions contains the optional parameters for the VaultsClient.ListDeleted method.
func (client *VaultsClient) NewListDeletedPager(options *VaultsClientListDeletedOptions) *runtime.Pager[VaultsClientListDeletedResponse] {
	return runtime.NewPager(runtime.PagingHandler[VaultsClientListDeletedResponse]{
		More: func(page VaultsClientListDeletedResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VaultsClientListDeletedResponse) (VaultsClientListDeletedResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listDeletedCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VaultsClientListDeletedResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VaultsClientListDeletedResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VaultsClientListDeletedResponse{}, runtime.NewResponseError(resp)
			}
			return client.listDeletedHandleResponse(resp)
		},
	})
}

// listDeletedCreateRequest creates the ListDeleted request.
func (client *VaultsClient) listDeletedCreateRequest(ctx context.Context, options *VaultsClientListDeletedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/deletedVaults"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listDeletedHandleResponse handles the ListDeleted response.
func (client *VaultsClient) listDeletedHandleResponse(resp *http.Response) (VaultsClientListDeletedResponse, error) {
	result := VaultsClientListDeletedResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedVaultListResult); err != nil {
		return VaultsClientListDeletedResponse{}, err
	}
	return result, nil
}

// BeginPurgeDeleted - Permanently deletes the specified vault. aka Purges the deleted Azure key vault.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// vaultName - The name of the soft-deleted vault.
// location - The location of the soft-deleted vault.
// options - VaultsClientBeginPurgeDeletedOptions contains the optional parameters for the VaultsClient.BeginPurgeDeleted
// method.
func (client *VaultsClient) BeginPurgeDeleted(ctx context.Context, vaultName string, location string, options *VaultsClientBeginPurgeDeletedOptions) (*runtime.Poller[VaultsClientPurgeDeletedResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.purgeDeleted(ctx, vaultName, location, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VaultsClientPurgeDeletedResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[VaultsClientPurgeDeletedResponse](options.ResumeToken, client.pl, nil)
	}
}

// PurgeDeleted - Permanently deletes the specified vault. aka Purges the deleted Azure key vault.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
func (client *VaultsClient) purgeDeleted(ctx context.Context, vaultName string, location string, options *VaultsClientBeginPurgeDeletedOptions) (*http.Response, error) {
	req, err := client.purgeDeletedCreateRequest(ctx, vaultName, location, options)
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

// purgeDeletedCreateRequest creates the PurgeDeleted request.
func (client *VaultsClient) purgeDeletedCreateRequest(ctx context.Context, vaultName string, location string, options *VaultsClientBeginPurgeDeletedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/locations/{location}/deletedVaults/{vaultName}/purge"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Update - Update a key vault in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - The name of the Resource Group to which the server belongs.
// vaultName - Name of the vault
// parameters - Parameters to patch the vault
// options - VaultsClientUpdateOptions contains the optional parameters for the VaultsClient.Update method.
func (client *VaultsClient) Update(ctx context.Context, resourceGroupName string, vaultName string, parameters VaultPatchParameters, options *VaultsClientUpdateOptions) (VaultsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vaultName, parameters, options)
	if err != nil {
		return VaultsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VaultsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return VaultsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *VaultsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, parameters VaultPatchParameters, options *VaultsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *VaultsClient) updateHandleResponse(resp *http.Response) (VaultsClientUpdateResponse, error) {
	result := VaultsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Vault); err != nil {
		return VaultsClientUpdateResponse{}, err
	}
	return result, nil
}

// UpdateAccessPolicy - Update access policies in a key vault in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - The name of the Resource Group to which the vault belongs.
// vaultName - Name of the vault
// operationKind - Name of the operation
// parameters - Access policy to merge into the vault
// options - VaultsClientUpdateAccessPolicyOptions contains the optional parameters for the VaultsClient.UpdateAccessPolicy
// method.
func (client *VaultsClient) UpdateAccessPolicy(ctx context.Context, resourceGroupName string, vaultName string, operationKind AccessPolicyUpdateKind, parameters VaultAccessPolicyParameters, options *VaultsClientUpdateAccessPolicyOptions) (VaultsClientUpdateAccessPolicyResponse, error) {
	req, err := client.updateAccessPolicyCreateRequest(ctx, resourceGroupName, vaultName, operationKind, parameters, options)
	if err != nil {
		return VaultsClientUpdateAccessPolicyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VaultsClientUpdateAccessPolicyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return VaultsClientUpdateAccessPolicyResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateAccessPolicyHandleResponse(resp)
}

// updateAccessPolicyCreateRequest creates the UpdateAccessPolicy request.
func (client *VaultsClient) updateAccessPolicyCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, operationKind AccessPolicyUpdateKind, parameters VaultAccessPolicyParameters, options *VaultsClientUpdateAccessPolicyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/accessPolicies/{operationKind}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if operationKind == "" {
		return nil, errors.New("parameter operationKind cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationKind}", url.PathEscape(string(operationKind)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateAccessPolicyHandleResponse handles the UpdateAccessPolicy response.
func (client *VaultsClient) updateAccessPolicyHandleResponse(resp *http.Response) (VaultsClientUpdateAccessPolicyResponse, error) {
	result := VaultsClientUpdateAccessPolicyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultAccessPolicyParameters); err != nil {
		return VaultsClientUpdateAccessPolicyResponse{}, err
	}
	return result, nil
}