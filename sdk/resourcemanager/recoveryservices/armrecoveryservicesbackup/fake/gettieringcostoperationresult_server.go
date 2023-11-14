//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
	"net/http"
	"net/url"
	"regexp"
)

// GetTieringCostOperationResultServer is a fake server for instances of the armrecoveryservicesbackup.GetTieringCostOperationResultClient type.
type GetTieringCostOperationResultServer struct {
	// Get is the fake for method GetTieringCostOperationResultClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *armrecoveryservicesbackup.GetTieringCostOperationResultClientGetOptions) (resp azfake.Responder[armrecoveryservicesbackup.GetTieringCostOperationResultClientGetResponse], errResp azfake.ErrorResponder)
}

// NewGetTieringCostOperationResultServerTransport creates a new instance of GetTieringCostOperationResultServerTransport with the provided implementation.
// The returned GetTieringCostOperationResultServerTransport instance is connected to an instance of armrecoveryservicesbackup.GetTieringCostOperationResultClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGetTieringCostOperationResultServerTransport(srv *GetTieringCostOperationResultServer) *GetTieringCostOperationResultServerTransport {
	return &GetTieringCostOperationResultServerTransport{srv: srv}
}

// GetTieringCostOperationResultServerTransport connects instances of armrecoveryservicesbackup.GetTieringCostOperationResultClient to instances of GetTieringCostOperationResultServer.
// Don't use this type directly, use NewGetTieringCostOperationResultServerTransport instead.
type GetTieringCostOperationResultServerTransport struct {
	srv *GetTieringCostOperationResultServer
}

// Do implements the policy.Transporter interface for GetTieringCostOperationResultServerTransport.
func (g *GetTieringCostOperationResultServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GetTieringCostOperationResultClient.Get":
		resp, err = g.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GetTieringCostOperationResultServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupTieringCost/default/operationResults/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Get(req.Context(), resourceGroupNameParam, vaultNameParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TieringCostInfoClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
