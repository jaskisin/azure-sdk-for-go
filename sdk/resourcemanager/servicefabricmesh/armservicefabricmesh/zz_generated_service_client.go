//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmesh

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

// ServiceClient contains the methods for the Service group.
// Don't use this type directly, use NewServiceClient() instead.
type ServiceClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServiceClient creates a new instance of ServiceClient with the specified values.
// subscriptionID - The customer subscription identifier
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServiceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServiceClient, error) {
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
	client := &ServiceClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets the information about the service resource with the given name. The information include the description and
// other properties of the service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-01-preview
// resourceGroupName - Azure resource group name
// applicationResourceName - The identity of the application.
// serviceResourceName - The identity of the service.
// options - ServiceClientGetOptions contains the optional parameters for the ServiceClient.Get method.
func (client *ServiceClient) Get(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string, options *ServiceClientGetOptions) (ServiceClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationResourceName, serviceResourceName, options)
	if err != nil {
		return ServiceClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServiceClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string, options *ServiceClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}/services/{serviceResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationResourceName}", applicationResourceName)
	urlPath = strings.ReplaceAll(urlPath, "{serviceResourceName}", serviceResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceClient) getHandleResponse(resp *http.Response) (ServiceClientGetResponse, error) {
	result := ServiceClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResourceDescription); err != nil {
		return ServiceClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the information about all services of an application resource. The information include the description
// and other properties of the Service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-01-preview
// resourceGroupName - Azure resource group name
// applicationResourceName - The identity of the application.
// options - ServiceClientListOptions contains the optional parameters for the ServiceClient.List method.
func (client *ServiceClient) NewListPager(resourceGroupName string, applicationResourceName string, options *ServiceClientListOptions) *runtime.Pager[ServiceClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServiceClientListResponse]{
		More: func(page ServiceClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServiceClientListResponse) (ServiceClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, applicationResourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServiceClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServiceClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServiceClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ServiceClient) listCreateRequest(ctx context.Context, resourceGroupName string, applicationResourceName string, options *ServiceClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}/services"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationResourceName}", applicationResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceClient) listHandleResponse(resp *http.Response) (ServiceClientListResponse, error) {
	result := ServiceClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResourceDescriptionList); err != nil {
		return ServiceClientListResponse{}, err
	}
	return result, nil
}