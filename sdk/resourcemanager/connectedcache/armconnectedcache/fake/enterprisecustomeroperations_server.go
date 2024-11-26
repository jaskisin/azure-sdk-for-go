// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedcache/armconnectedcache"
	"net/http"
	"net/url"
	"regexp"
)

// EnterpriseCustomerOperationsServer is a fake server for instances of the armconnectedcache.EnterpriseCustomerOperationsClient type.
type EnterpriseCustomerOperationsServer struct {
	// BeginCreateOrUpdate is the fake for method EnterpriseCustomerOperationsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, customerResourceName string, resource armconnectedcache.EnterprisePreviewResource, options *armconnectedcache.EnterpriseCustomerOperationsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armconnectedcache.EnterpriseCustomerOperationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method EnterpriseCustomerOperationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, customerResourceName string, options *armconnectedcache.EnterpriseCustomerOperationsClientDeleteOptions) (resp azfake.Responder[armconnectedcache.EnterpriseCustomerOperationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method EnterpriseCustomerOperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, customerResourceName string, options *armconnectedcache.EnterpriseCustomerOperationsClientGetOptions) (resp azfake.Responder[armconnectedcache.EnterpriseCustomerOperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method EnterpriseCustomerOperationsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armconnectedcache.EnterpriseCustomerOperationsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armconnectedcache.EnterpriseCustomerOperationsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method EnterpriseCustomerOperationsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armconnectedcache.EnterpriseCustomerOperationsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armconnectedcache.EnterpriseCustomerOperationsClientListBySubscriptionResponse])

	// Update is the fake for method EnterpriseCustomerOperationsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, customerResourceName string, properties armconnectedcache.PatchResource, options *armconnectedcache.EnterpriseCustomerOperationsClientUpdateOptions) (resp azfake.Responder[armconnectedcache.EnterpriseCustomerOperationsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewEnterpriseCustomerOperationsServerTransport creates a new instance of EnterpriseCustomerOperationsServerTransport with the provided implementation.
// The returned EnterpriseCustomerOperationsServerTransport instance is connected to an instance of armconnectedcache.EnterpriseCustomerOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEnterpriseCustomerOperationsServerTransport(srv *EnterpriseCustomerOperationsServer) *EnterpriseCustomerOperationsServerTransport {
	return &EnterpriseCustomerOperationsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armconnectedcache.EnterpriseCustomerOperationsClientCreateOrUpdateResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armconnectedcache.EnterpriseCustomerOperationsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armconnectedcache.EnterpriseCustomerOperationsClientListBySubscriptionResponse]](),
	}
}

// EnterpriseCustomerOperationsServerTransport connects instances of armconnectedcache.EnterpriseCustomerOperationsClient to instances of EnterpriseCustomerOperationsServer.
// Don't use this type directly, use NewEnterpriseCustomerOperationsServerTransport instead.
type EnterpriseCustomerOperationsServerTransport struct {
	srv                         *EnterpriseCustomerOperationsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armconnectedcache.EnterpriseCustomerOperationsClientCreateOrUpdateResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armconnectedcache.EnterpriseCustomerOperationsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armconnectedcache.EnterpriseCustomerOperationsClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for EnterpriseCustomerOperationsServerTransport.
func (e *EnterpriseCustomerOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *EnterpriseCustomerOperationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if enterpriseCustomerOperationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = enterpriseCustomerOperationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "EnterpriseCustomerOperationsClient.BeginCreateOrUpdate":
				res.resp, res.err = e.dispatchBeginCreateOrUpdate(req)
			case "EnterpriseCustomerOperationsClient.Delete":
				res.resp, res.err = e.dispatchDelete(req)
			case "EnterpriseCustomerOperationsClient.Get":
				res.resp, res.err = e.dispatchGet(req)
			case "EnterpriseCustomerOperationsClient.NewListByResourceGroupPager":
				res.resp, res.err = e.dispatchNewListByResourceGroupPager(req)
			case "EnterpriseCustomerOperationsClient.NewListBySubscriptionPager":
				res.resp, res.err = e.dispatchNewListBySubscriptionPager(req)
			case "EnterpriseCustomerOperationsClient.Update":
				res.resp, res.err = e.dispatchUpdate(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (e *EnterpriseCustomerOperationsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := e.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedCache/enterpriseCustomers/(?P<customerResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armconnectedcache.EnterprisePreviewResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		customerResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("customerResourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, customerResourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		e.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		e.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		e.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (e *EnterpriseCustomerOperationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if e.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedCache/enterpriseCustomers/(?P<customerResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	customerResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("customerResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Delete(req.Context(), resourceGroupNameParam, customerResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnterpriseCustomerOperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedCache/enterpriseCustomers/(?P<customerResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	customerResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("customerResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, customerResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EnterprisePreviewResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnterpriseCustomerOperationsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := e.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedCache/enterpriseCustomers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		e.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armconnectedcache.EnterpriseCustomerOperationsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		e.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (e *EnterpriseCustomerOperationsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := e.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedCache/enterpriseCustomers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := e.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		e.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armconnectedcache.EnterpriseCustomerOperationsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		e.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (e *EnterpriseCustomerOperationsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedCache/enterpriseCustomers/(?P<customerResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armconnectedcache.PatchResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	customerResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("customerResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Update(req.Context(), resourceGroupNameParam, customerResourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EnterprisePreviewResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to EnterpriseCustomerOperationsServerTransport
var enterpriseCustomerOperationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
