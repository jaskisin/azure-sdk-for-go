//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

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

// Client contains the methods for the RecoveryServicesBackupClient group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewClient creates a new instance of Client with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
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
	client := &Client{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginBMSPrepareDataMove - Prepares source vault for Data Move operation
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// parameters - Prepare data move request
// options - ClientBeginBMSPrepareDataMoveOptions contains the optional parameters for the Client.BeginBMSPrepareDataMove
// method.
func (client *Client) BeginBMSPrepareDataMove(ctx context.Context, vaultName string, resourceGroupName string, parameters PrepareDataMoveRequest, options *ClientBeginBMSPrepareDataMoveOptions) (*runtime.Poller[ClientBMSPrepareDataMoveResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.bMSPrepareDataMove(ctx, vaultName, resourceGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ClientBMSPrepareDataMoveResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ClientBMSPrepareDataMoveResponse](options.ResumeToken, client.pl, nil)
	}
}

// BMSPrepareDataMove - Prepares source vault for Data Move operation
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
func (client *Client) bMSPrepareDataMove(ctx context.Context, vaultName string, resourceGroupName string, parameters PrepareDataMoveRequest, options *ClientBeginBMSPrepareDataMoveOptions) (*http.Response, error) {
	req, err := client.bmsPrepareDataMoveCreateRequest(ctx, vaultName, resourceGroupName, parameters, options)
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

// bmsPrepareDataMoveCreateRequest creates the BMSPrepareDataMove request.
func (client *Client) bmsPrepareDataMoveCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, parameters PrepareDataMoveRequest, options *ClientBeginBMSPrepareDataMoveOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupstorageconfig/vaultstorageconfig/prepareDataMove"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginBMSTriggerDataMove - Triggers Data Move Operation on target vault
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// parameters - Trigger data move request
// options - ClientBeginBMSTriggerDataMoveOptions contains the optional parameters for the Client.BeginBMSTriggerDataMove
// method.
func (client *Client) BeginBMSTriggerDataMove(ctx context.Context, vaultName string, resourceGroupName string, parameters TriggerDataMoveRequest, options *ClientBeginBMSTriggerDataMoveOptions) (*runtime.Poller[ClientBMSTriggerDataMoveResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.bMSTriggerDataMove(ctx, vaultName, resourceGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ClientBMSTriggerDataMoveResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ClientBMSTriggerDataMoveResponse](options.ResumeToken, client.pl, nil)
	}
}

// BMSTriggerDataMove - Triggers Data Move Operation on target vault
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
func (client *Client) bMSTriggerDataMove(ctx context.Context, vaultName string, resourceGroupName string, parameters TriggerDataMoveRequest, options *ClientBeginBMSTriggerDataMoveOptions) (*http.Response, error) {
	req, err := client.bmsTriggerDataMoveCreateRequest(ctx, vaultName, resourceGroupName, parameters, options)
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

// bmsTriggerDataMoveCreateRequest creates the BMSTriggerDataMove request.
func (client *Client) bmsTriggerDataMoveCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, parameters TriggerDataMoveRequest, options *ClientBeginBMSTriggerDataMoveOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupstorageconfig/vaultstorageconfig/triggerDataMove"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// GetOperationStatus - Fetches operation status for data move operation on vault
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// options - ClientGetOperationStatusOptions contains the optional parameters for the Client.GetOperationStatus method.
func (client *Client) GetOperationStatus(ctx context.Context, vaultName string, resourceGroupName string, operationID string, options *ClientGetOperationStatusOptions) (ClientGetOperationStatusResponse, error) {
	req, err := client.getOperationStatusCreateRequest(ctx, vaultName, resourceGroupName, operationID, options)
	if err != nil {
		return ClientGetOperationStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientGetOperationStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientGetOperationStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOperationStatusHandleResponse(resp)
}

// getOperationStatusCreateRequest creates the GetOperationStatus request.
func (client *Client) getOperationStatusCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, operationID string, options *ClientGetOperationStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupstorageconfig/vaultstorageconfig/operationStatus/{operationId}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getOperationStatusHandleResponse handles the GetOperationStatus response.
func (client *Client) getOperationStatusHandleResponse(resp *http.Response) (ClientGetOperationStatusResponse, error) {
	result := ClientGetOperationStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatus); err != nil {
		return ClientGetOperationStatusResponse{}, err
	}
	return result, nil
}

// BeginMoveRecoveryPoint - Move recovery point from one datastore to another store.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// parameters - Move Resource Across Tiers Request
// options - ClientBeginMoveRecoveryPointOptions contains the optional parameters for the Client.BeginMoveRecoveryPoint method.
func (client *Client) BeginMoveRecoveryPoint(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters MoveRPAcrossTiersRequest, options *ClientBeginMoveRecoveryPointOptions) (*runtime.Poller[ClientMoveRecoveryPointResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.moveRecoveryPoint(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, recoveryPointID, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ClientMoveRecoveryPointResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ClientMoveRecoveryPointResponse](options.ResumeToken, client.pl, nil)
	}
}

// MoveRecoveryPoint - Move recovery point from one datastore to another store.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
func (client *Client) moveRecoveryPoint(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters MoveRPAcrossTiersRequest, options *ClientBeginMoveRecoveryPointOptions) (*http.Response, error) {
	req, err := client.moveRecoveryPointCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, recoveryPointID, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// moveRecoveryPointCreateRequest creates the MoveRecoveryPoint request.
func (client *Client) moveRecoveryPointCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters MoveRPAcrossTiersRequest, options *ClientBeginMoveRecoveryPointOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointId}/move"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	if protectedItemName == "" {
		return nil, errors.New("parameter protectedItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectedItemName}", url.PathEscape(protectedItemName))
	if recoveryPointID == "" {
		return nil, errors.New("parameter recoveryPointID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recoveryPointId}", url.PathEscape(recoveryPointID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}