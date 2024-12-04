//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armworkloadssapvirtualinstance

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

// SapDatabaseInstancesClient contains the methods for the SapDatabaseInstances group.
// Don't use this type directly, use NewSapDatabaseInstancesClient() instead.
type SapDatabaseInstancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSapDatabaseInstancesClient creates a new instance of SapDatabaseInstancesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSapDatabaseInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SapDatabaseInstancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SapDatabaseInstancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates the Database resource corresponding to the Virtual Instance for SAP solutions resource.
// This will be used by service only. PUT by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapVirtualInstanceName - The name of the Virtual Instances for SAP solutions resource
//   - databaseInstanceName - Database resource name string modeled as parameter for auto generation to work correctly.
//   - resource - Request body of Database resource of a SAP system.
//   - options - SapDatabaseInstancesClientBeginCreateOptions contains the optional parameters for the SapDatabaseInstancesClient.BeginCreate
//     method.
func (client *SapDatabaseInstancesClient) BeginCreate(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, resource SAPDatabaseInstance, options *SapDatabaseInstancesClientBeginCreateOptions) (*runtime.Poller[SapDatabaseInstancesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, sapVirtualInstanceName, databaseInstanceName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SapDatabaseInstancesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SapDatabaseInstancesClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates the Database resource corresponding to the Virtual Instance for SAP solutions resource.
// This will be used by service only. PUT by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
func (client *SapDatabaseInstancesClient) create(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, resource SAPDatabaseInstance, options *SapDatabaseInstancesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "SapDatabaseInstancesClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, databaseInstanceName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *SapDatabaseInstancesClient) createCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, resource SAPDatabaseInstance, options *SapDatabaseInstancesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/databaseInstances/{databaseInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if databaseInstanceName == "" {
		return nil, errors.New("parameter databaseInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseInstanceName}", url.PathEscape(databaseInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the Database resource corresponding to a Virtual Instance for SAP solutions resource.
// This will be used by service only. Delete by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapVirtualInstanceName - The name of the Virtual Instances for SAP solutions resource
//   - databaseInstanceName - Database resource name string modeled as parameter for auto generation to work correctly.
//   - options - SapDatabaseInstancesClientBeginDeleteOptions contains the optional parameters for the SapDatabaseInstancesClient.BeginDelete
//     method.
func (client *SapDatabaseInstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, options *SapDatabaseInstancesClientBeginDeleteOptions) (*runtime.Poller[SapDatabaseInstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sapVirtualInstanceName, databaseInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SapDatabaseInstancesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SapDatabaseInstancesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the Database resource corresponding to a Virtual Instance for SAP solutions resource.
// This will be used by service only. Delete by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
func (client *SapDatabaseInstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, options *SapDatabaseInstancesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SapDatabaseInstancesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, databaseInstanceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SapDatabaseInstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, options *SapDatabaseInstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/databaseInstances/{databaseInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if databaseInstanceName == "" {
		return nil, errors.New("parameter databaseInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseInstanceName}", url.PathEscape(databaseInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the SAP Database Instance resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapVirtualInstanceName - The name of the Virtual Instances for SAP solutions resource
//   - databaseInstanceName - Database resource name string modeled as parameter for auto generation to work correctly.
//   - options - SapDatabaseInstancesClientGetOptions contains the optional parameters for the SapDatabaseInstancesClient.Get
//     method.
func (client *SapDatabaseInstancesClient) Get(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, options *SapDatabaseInstancesClientGetOptions) (SapDatabaseInstancesClientGetResponse, error) {
	var err error
	const operationName = "SapDatabaseInstancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, databaseInstanceName, options)
	if err != nil {
		return SapDatabaseInstancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SapDatabaseInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SapDatabaseInstancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SapDatabaseInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, options *SapDatabaseInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/databaseInstances/{databaseInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if databaseInstanceName == "" {
		return nil, errors.New("parameter databaseInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseInstanceName}", url.PathEscape(databaseInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SapDatabaseInstancesClient) getHandleResponse(resp *http.Response) (SapDatabaseInstancesClientGetResponse, error) {
	result := SapDatabaseInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPDatabaseInstance); err != nil {
		return SapDatabaseInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the Database resources associated with a Virtual Instance for SAP solutions resource.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapVirtualInstanceName - The name of the Virtual Instances for SAP solutions resource
//   - options - SapDatabaseInstancesClientListOptions contains the optional parameters for the SapDatabaseInstancesClient.NewListPager
//     method.
func (client *SapDatabaseInstancesClient) NewListPager(resourceGroupName string, sapVirtualInstanceName string, options *SapDatabaseInstancesClientListOptions) *runtime.Pager[SapDatabaseInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SapDatabaseInstancesClientListResponse]{
		More: func(page SapDatabaseInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SapDatabaseInstancesClientListResponse) (SapDatabaseInstancesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SapDatabaseInstancesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, options)
			}, nil)
			if err != nil {
				return SapDatabaseInstancesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SapDatabaseInstancesClient) listCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, options *SapDatabaseInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/databaseInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SapDatabaseInstancesClient) listHandleResponse(resp *http.Response) (SapDatabaseInstancesClientListResponse, error) {
	result := SapDatabaseInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPDatabaseInstanceListResult); err != nil {
		return SapDatabaseInstancesClientListResponse{}, err
	}
	return result, nil
}

// BeginStart - Starts the database instance of the SAP system.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapVirtualInstanceName - The name of the Virtual Instances for SAP solutions resource
//   - databaseInstanceName - Database resource name string modeled as parameter for auto generation to work correctly.
//   - options - SapDatabaseInstancesClientBeginStartOptions contains the optional parameters for the SapDatabaseInstancesClient.BeginStart
//     method.
func (client *SapDatabaseInstancesClient) BeginStart(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, options *SapDatabaseInstancesClientBeginStartOptions) (*runtime.Poller[SapDatabaseInstancesClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, sapVirtualInstanceName, databaseInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SapDatabaseInstancesClientStartResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SapDatabaseInstancesClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Start - Starts the database instance of the SAP system.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
func (client *SapDatabaseInstancesClient) start(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, options *SapDatabaseInstancesClientBeginStartOptions) (*http.Response, error) {
	var err error
	const operationName = "SapDatabaseInstancesClient.BeginStart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, databaseInstanceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// startCreateRequest creates the Start request.
func (client *SapDatabaseInstancesClient) startCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, options *SapDatabaseInstancesClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/databaseInstances/{databaseInstanceName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if databaseInstanceName == "" {
		return nil, errors.New("parameter databaseInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseInstanceName}", url.PathEscape(databaseInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// BeginStop - Stops the database instance of the SAP system.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapVirtualInstanceName - The name of the Virtual Instances for SAP solutions resource
//   - databaseInstanceName - Database resource name string modeled as parameter for auto generation to work correctly.
//   - options - SapDatabaseInstancesClientBeginStopOptions contains the optional parameters for the SapDatabaseInstancesClient.BeginStop
//     method.
func (client *SapDatabaseInstancesClient) BeginStop(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, options *SapDatabaseInstancesClientBeginStopOptions) (*runtime.Poller[SapDatabaseInstancesClientStopResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stop(ctx, resourceGroupName, sapVirtualInstanceName, databaseInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SapDatabaseInstancesClientStopResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SapDatabaseInstancesClientStopResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Stop - Stops the database instance of the SAP system.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
func (client *SapDatabaseInstancesClient) stop(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, options *SapDatabaseInstancesClientBeginStopOptions) (*http.Response, error) {
	var err error
	const operationName = "SapDatabaseInstancesClient.BeginStop"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.stopCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, databaseInstanceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// stopCreateRequest creates the Stop request.
func (client *SapDatabaseInstancesClient) stopCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, options *SapDatabaseInstancesClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/databaseInstances/{databaseInstanceName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if databaseInstanceName == "" {
		return nil, errors.New("parameter databaseInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseInstanceName}", url.PathEscape(databaseInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// Update - Updates the Database resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sapVirtualInstanceName - The name of the Virtual Instances for SAP solutions resource
//   - databaseInstanceName - Database resource name string modeled as parameter for auto generation to work correctly.
//   - properties - Database resource update request body.
//   - options - SapDatabaseInstancesClientUpdateOptions contains the optional parameters for the SapDatabaseInstancesClient.Update
//     method.
func (client *SapDatabaseInstancesClient) Update(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, properties UpdateSAPDatabaseInstanceRequest, options *SapDatabaseInstancesClientUpdateOptions) (SapDatabaseInstancesClientUpdateResponse, error) {
	var err error
	const operationName = "SapDatabaseInstancesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, databaseInstanceName, properties, options)
	if err != nil {
		return SapDatabaseInstancesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SapDatabaseInstancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SapDatabaseInstancesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *SapDatabaseInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, databaseInstanceName string, properties UpdateSAPDatabaseInstanceRequest, options *SapDatabaseInstancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/databaseInstances/{databaseInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if databaseInstanceName == "" {
		return nil, errors.New("parameter databaseInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseInstanceName}", url.PathEscape(databaseInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *SapDatabaseInstancesClient) updateHandleResponse(resp *http.Response) (SapDatabaseInstancesClientUpdateResponse, error) {
	result := SapDatabaseInstancesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPDatabaseInstance); err != nil {
		return SapDatabaseInstancesClientUpdateResponse{}, err
	}
	return result, nil
}
