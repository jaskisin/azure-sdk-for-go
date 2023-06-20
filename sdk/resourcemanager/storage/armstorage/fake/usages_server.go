//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"net/http"
	"net/url"
	"regexp"
)

// UsagesServer is a fake server for instances of the armstorage.UsagesClient type.
type UsagesServer struct {
	// NewListByLocationPager is the fake for method UsagesClient.NewListByLocationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByLocationPager func(location string, options *armstorage.UsagesClientListByLocationOptions) (resp azfake.PagerResponder[armstorage.UsagesClientListByLocationResponse])
}

// NewUsagesServerTransport creates a new instance of UsagesServerTransport with the provided implementation.
// The returned UsagesServerTransport instance is connected to an instance of armstorage.UsagesClient by way of the
// undefined.Transporter field.
func NewUsagesServerTransport(srv *UsagesServer) *UsagesServerTransport {
	return &UsagesServerTransport{srv: srv}
}

// UsagesServerTransport connects instances of armstorage.UsagesClient to instances of UsagesServer.
// Don't use this type directly, use NewUsagesServerTransport instead.
type UsagesServerTransport struct {
	srv                    *UsagesServer
	newListByLocationPager *azfake.PagerResponder[armstorage.UsagesClientListByLocationResponse]
}

// Do implements the policy.Transporter interface for UsagesServerTransport.
func (u *UsagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "UsagesClient.NewListByLocationPager":
		resp, err = u.dispatchNewListByLocationPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *UsagesServerTransport) dispatchNewListByLocationPager(req *http.Request) (*http.Response, error) {
	if u.srv.NewListByLocationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByLocationPager not implemented")}
	}
	if u.newListByLocationPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/usages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := u.srv.NewListByLocationPager(locationUnescaped, nil)
		u.newListByLocationPager = &resp
	}
	resp, err := server.PagerResponderNext(u.newListByLocationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(u.newListByLocationPager) {
		u.newListByLocationPager = nil
	}
	return resp, nil
}