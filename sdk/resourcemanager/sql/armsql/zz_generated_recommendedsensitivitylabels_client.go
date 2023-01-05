//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// RecommendedSensitivityLabelsClient contains the methods for the RecommendedSensitivityLabels group.
// Don't use this type directly, use NewRecommendedSensitivityLabelsClient() instead.
type RecommendedSensitivityLabelsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRecommendedSensitivityLabelsClient creates a new instance of RecommendedSensitivityLabelsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRecommendedSensitivityLabelsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RecommendedSensitivityLabelsClient, error) {
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
	client := &RecommendedSensitivityLabelsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Update - Update recommended sensitivity labels states of a given database using an operations batch.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// databaseName - The name of the database.
// options - RecommendedSensitivityLabelsClientUpdateOptions contains the optional parameters for the RecommendedSensitivityLabelsClient.Update
// method.
func (client *RecommendedSensitivityLabelsClient) Update(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters RecommendedSensitivityLabelUpdateList, options *RecommendedSensitivityLabelsClientUpdateOptions) (RecommendedSensitivityLabelsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, databaseName, parameters, options)
	if err != nil {
		return RecommendedSensitivityLabelsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecommendedSensitivityLabelsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RecommendedSensitivityLabelsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return RecommendedSensitivityLabelsClientUpdateResponse{}, nil
}

// updateCreateRequest creates the Update request.
func (client *RecommendedSensitivityLabelsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters RecommendedSensitivityLabelUpdateList, options *RecommendedSensitivityLabelsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/recommendedSensitivityLabels"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}