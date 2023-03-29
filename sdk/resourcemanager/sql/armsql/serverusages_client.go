//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql

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

// ServerUsagesClient contains the methods for the ServerUsages group.
// Don't use this type directly, use NewServerUsagesClient() instead.
type ServerUsagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerUsagesClient creates a new instance of ServerUsagesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerUsagesClient, error) {
	cl, err := arm.NewClient(moduleName+".ServerUsagesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerUsagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByServerPager - Returns server usages.
//
// Generated from API version 2014-04-01
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - ServerUsagesClientListByServerOptions contains the optional parameters for the ServerUsagesClient.NewListByServerPager
//     method.
func (client *ServerUsagesClient) NewListByServerPager(resourceGroupName string, serverName string, options *ServerUsagesClientListByServerOptions) *runtime.Pager[ServerUsagesClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerUsagesClientListByServerResponse]{
		More: func(page ServerUsagesClientListByServerResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ServerUsagesClientListByServerResponse) (ServerUsagesClientListByServerResponse, error) {
			req, err := client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			if err != nil {
				return ServerUsagesClientListByServerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServerUsagesClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServerUsagesClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerUsagesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerUsagesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerUsagesClient) listByServerHandleResponse(resp *http.Response) (ServerUsagesClientListByServerResponse, error) {
	result := ServerUsagesClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerUsageListResult); err != nil {
		return ServerUsagesClientListByServerResponse{}, err
	}
	return result, nil
}