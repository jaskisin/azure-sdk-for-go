//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ChangeDataCaptureClient contains the methods for the ChangeDataCapture group.
// Don't use this type directly, use NewChangeDataCaptureClient() instead.
type ChangeDataCaptureClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewChangeDataCaptureClient creates a new instance of ChangeDataCaptureClient with the specified values.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewChangeDataCaptureClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ChangeDataCaptureClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ChangeDataCaptureClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a change data capture resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - changeDataCaptureName - The change data capture name.
//   - changeDataCapture - Change data capture resource definition.
//   - options - ChangeDataCaptureClientCreateOrUpdateOptions contains the optional parameters for the ChangeDataCaptureClient.CreateOrUpdate
//     method.
func (client *ChangeDataCaptureClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, changeDataCapture ChangeDataCaptureResource, options *ChangeDataCaptureClientCreateOrUpdateOptions) (ChangeDataCaptureClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ChangeDataCaptureClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, factoryName, changeDataCaptureName, changeDataCapture, options)
	if err != nil {
		return ChangeDataCaptureClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ChangeDataCaptureClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ChangeDataCaptureClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ChangeDataCaptureClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, changeDataCapture ChangeDataCaptureResource, options *ChangeDataCaptureClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/adfcdcs/{changeDataCaptureName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if changeDataCaptureName == "" {
		return nil, errors.New("parameter changeDataCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{changeDataCaptureName}", url.PathEscape(changeDataCaptureName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if err := runtime.MarshalAsJSON(req, changeDataCapture); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ChangeDataCaptureClient) createOrUpdateHandleResponse(resp *http.Response) (ChangeDataCaptureClientCreateOrUpdateResponse, error) {
	result := ChangeDataCaptureClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ChangeDataCaptureResource); err != nil {
		return ChangeDataCaptureClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a change data capture.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - changeDataCaptureName - The change data capture name.
//   - options - ChangeDataCaptureClientDeleteOptions contains the optional parameters for the ChangeDataCaptureClient.Delete
//     method.
func (client *ChangeDataCaptureClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, options *ChangeDataCaptureClientDeleteOptions) (ChangeDataCaptureClientDeleteResponse, error) {
	var err error
	const operationName = "ChangeDataCaptureClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, changeDataCaptureName, options)
	if err != nil {
		return ChangeDataCaptureClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ChangeDataCaptureClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ChangeDataCaptureClientDeleteResponse{}, err
	}
	return ChangeDataCaptureClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ChangeDataCaptureClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, options *ChangeDataCaptureClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/adfcdcs/{changeDataCaptureName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if changeDataCaptureName == "" {
		return nil, errors.New("parameter changeDataCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{changeDataCaptureName}", url.PathEscape(changeDataCaptureName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a change data capture.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - changeDataCaptureName - The change data capture name.
//   - options - ChangeDataCaptureClientGetOptions contains the optional parameters for the ChangeDataCaptureClient.Get method.
func (client *ChangeDataCaptureClient) Get(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, options *ChangeDataCaptureClientGetOptions) (ChangeDataCaptureClientGetResponse, error) {
	var err error
	const operationName = "ChangeDataCaptureClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, changeDataCaptureName, options)
	if err != nil {
		return ChangeDataCaptureClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ChangeDataCaptureClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ChangeDataCaptureClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ChangeDataCaptureClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, options *ChangeDataCaptureClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/adfcdcs/{changeDataCaptureName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if changeDataCaptureName == "" {
		return nil, errors.New("parameter changeDataCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{changeDataCaptureName}", url.PathEscape(changeDataCaptureName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ChangeDataCaptureClient) getHandleResponse(resp *http.Response) (ChangeDataCaptureClientGetResponse, error) {
	result := ChangeDataCaptureClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ChangeDataCaptureResource); err != nil {
		return ChangeDataCaptureClientGetResponse{}, err
	}
	return result, nil
}

// NewListByFactoryPager - Lists all resources of type change data capture.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - options - ChangeDataCaptureClientListByFactoryOptions contains the optional parameters for the ChangeDataCaptureClient.NewListByFactoryPager
//     method.
func (client *ChangeDataCaptureClient) NewListByFactoryPager(resourceGroupName string, factoryName string, options *ChangeDataCaptureClientListByFactoryOptions) *runtime.Pager[ChangeDataCaptureClientListByFactoryResponse] {
	return runtime.NewPager(runtime.PagingHandler[ChangeDataCaptureClientListByFactoryResponse]{
		More: func(page ChangeDataCaptureClientListByFactoryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ChangeDataCaptureClientListByFactoryResponse) (ChangeDataCaptureClientListByFactoryResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ChangeDataCaptureClient.NewListByFactoryPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByFactoryCreateRequest(ctx, resourceGroupName, factoryName, options)
			}, nil)
			if err != nil {
				return ChangeDataCaptureClientListByFactoryResponse{}, err
			}
			return client.listByFactoryHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByFactoryCreateRequest creates the ListByFactory request.
func (client *ChangeDataCaptureClient) listByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, options *ChangeDataCaptureClientListByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/adfcdcs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFactoryHandleResponse handles the ListByFactory response.
func (client *ChangeDataCaptureClient) listByFactoryHandleResponse(resp *http.Response) (ChangeDataCaptureClientListByFactoryResponse, error) {
	result := ChangeDataCaptureClientListByFactoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ChangeDataCaptureListResponse); err != nil {
		return ChangeDataCaptureClientListByFactoryResponse{}, err
	}
	return result, nil
}

// Start - Starts a change data capture.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - changeDataCaptureName - The change data capture name.
//   - options - ChangeDataCaptureClientStartOptions contains the optional parameters for the ChangeDataCaptureClient.Start method.
func (client *ChangeDataCaptureClient) Start(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, options *ChangeDataCaptureClientStartOptions) (ChangeDataCaptureClientStartResponse, error) {
	var err error
	const operationName = "ChangeDataCaptureClient.Start"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startCreateRequest(ctx, resourceGroupName, factoryName, changeDataCaptureName, options)
	if err != nil {
		return ChangeDataCaptureClientStartResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ChangeDataCaptureClientStartResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ChangeDataCaptureClientStartResponse{}, err
	}
	return ChangeDataCaptureClientStartResponse{}, nil
}

// startCreateRequest creates the Start request.
func (client *ChangeDataCaptureClient) startCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, options *ChangeDataCaptureClientStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/adfcdcs/{changeDataCaptureName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if changeDataCaptureName == "" {
		return nil, errors.New("parameter changeDataCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{changeDataCaptureName}", url.PathEscape(changeDataCaptureName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Status - Gets the current status for the change data capture resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - changeDataCaptureName - The change data capture name.
//   - options - ChangeDataCaptureClientStatusOptions contains the optional parameters for the ChangeDataCaptureClient.Status
//     method.
func (client *ChangeDataCaptureClient) Status(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, options *ChangeDataCaptureClientStatusOptions) (ChangeDataCaptureClientStatusResponse, error) {
	var err error
	const operationName = "ChangeDataCaptureClient.Status"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.statusCreateRequest(ctx, resourceGroupName, factoryName, changeDataCaptureName, options)
	if err != nil {
		return ChangeDataCaptureClientStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ChangeDataCaptureClientStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ChangeDataCaptureClientStatusResponse{}, err
	}
	resp, err := client.statusHandleResponse(httpResp)
	return resp, err
}

// statusCreateRequest creates the Status request.
func (client *ChangeDataCaptureClient) statusCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, options *ChangeDataCaptureClientStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/adfcdcs/{changeDataCaptureName}/status"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if changeDataCaptureName == "" {
		return nil, errors.New("parameter changeDataCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{changeDataCaptureName}", url.PathEscape(changeDataCaptureName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// statusHandleResponse handles the Status response.
func (client *ChangeDataCaptureClient) statusHandleResponse(resp *http.Response) (ChangeDataCaptureClientStatusResponse, error) {
	result := ChangeDataCaptureClientStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return ChangeDataCaptureClientStatusResponse{}, err
	}
	return result, nil
}

// Stop - Stops a change data capture.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - changeDataCaptureName - The change data capture name.
//   - options - ChangeDataCaptureClientStopOptions contains the optional parameters for the ChangeDataCaptureClient.Stop method.
func (client *ChangeDataCaptureClient) Stop(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, options *ChangeDataCaptureClientStopOptions) (ChangeDataCaptureClientStopResponse, error) {
	var err error
	const operationName = "ChangeDataCaptureClient.Stop"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.stopCreateRequest(ctx, resourceGroupName, factoryName, changeDataCaptureName, options)
	if err != nil {
		return ChangeDataCaptureClientStopResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ChangeDataCaptureClientStopResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ChangeDataCaptureClientStopResponse{}, err
	}
	return ChangeDataCaptureClientStopResponse{}, nil
}

// stopCreateRequest creates the Stop request.
func (client *ChangeDataCaptureClient) stopCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, changeDataCaptureName string, options *ChangeDataCaptureClientStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/adfcdcs/{changeDataCaptureName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if changeDataCaptureName == "" {
		return nil, errors.New("parameter changeDataCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{changeDataCaptureName}", url.PathEscape(changeDataCaptureName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
