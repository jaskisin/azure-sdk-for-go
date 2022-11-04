//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentmanager

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

// StepsClient contains the methods for the Steps group.
// Don't use this type directly, use NewStepsClient() instead.
type StepsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewStepsClient creates a new instance of StepsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewStepsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StepsClient, error) {
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
	client := &StepsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Synchronously creates a new step or updates an existing step.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// stepName - The name of the deployment step.
// options - StepsClientCreateOrUpdateOptions contains the optional parameters for the StepsClient.CreateOrUpdate method.
func (client *StepsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, stepName string, options *StepsClientCreateOrUpdateOptions) (StepsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, stepName, options)
	if err != nil {
		return StepsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StepsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return StepsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *StepsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, stepName string, options *StepsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/steps/{stepName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if stepName == "" {
		return nil, errors.New("parameter stepName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{stepName}", url.PathEscape(stepName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.StepInfo != nil {
		return req, runtime.MarshalAsJSON(req, *options.StepInfo)
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *StepsClient) createOrUpdateHandleResponse(resp *http.Response) (StepsClientCreateOrUpdateResponse, error) {
	result := StepsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StepResource); err != nil {
		return StepsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the step.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// stepName - The name of the deployment step.
// options - StepsClientDeleteOptions contains the optional parameters for the StepsClient.Delete method.
func (client *StepsClient) Delete(ctx context.Context, resourceGroupName string, stepName string, options *StepsClientDeleteOptions) (StepsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, stepName, options)
	if err != nil {
		return StepsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StepsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return StepsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return StepsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StepsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, stepName string, options *StepsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/steps/{stepName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if stepName == "" {
		return nil, errors.New("parameter stepName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{stepName}", url.PathEscape(stepName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the step.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// stepName - The name of the deployment step.
// options - StepsClientGetOptions contains the optional parameters for the StepsClient.Get method.
func (client *StepsClient) Get(ctx context.Context, resourceGroupName string, stepName string, options *StepsClientGetOptions) (StepsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, stepName, options)
	if err != nil {
		return StepsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StepsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StepsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StepsClient) getCreateRequest(ctx context.Context, resourceGroupName string, stepName string, options *StepsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/steps/{stepName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if stepName == "" {
		return nil, errors.New("parameter stepName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{stepName}", url.PathEscape(stepName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StepsClient) getHandleResponse(resp *http.Response) (StepsClientGetResponse, error) {
	result := StepsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StepResource); err != nil {
		return StepsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists the steps in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - StepsClientListOptions contains the optional parameters for the StepsClient.List method.
func (client *StepsClient) List(ctx context.Context, resourceGroupName string, options *StepsClientListOptions) (StepsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return StepsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StepsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StepsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *StepsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *StepsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/steps"
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
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *StepsClient) listHandleResponse(resp *http.Response) (StepsClientListResponse, error) {
	result := StepsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StepResourceArray); err != nil {
		return StepsClientListResponse{}, err
	}
	return result, nil
}