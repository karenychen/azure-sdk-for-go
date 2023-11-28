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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateEndpointConnectionsForMIPPolicySyncServer is a fake server for instances of the armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClient type.
type PrivateEndpointConnectionsForMIPPolicySyncServer struct {
	// BeginCreateOrUpdate is the fake for method PrivateEndpointConnectionsForMIPPolicySyncClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, properties armm365securityandcompliance.PrivateEndpointConnection, options *armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PrivateEndpointConnectionsForMIPPolicySyncClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClientBeginDeleteOptions) (resp azfake.PollerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateEndpointConnectionsForMIPPolicySyncClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClientGetOptions) (resp azfake.Responder[armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServicePager is the fake for method PrivateEndpointConnectionsForMIPPolicySyncClient.NewListByServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServicePager func(resourceGroupName string, resourceName string, options *armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClientListByServiceOptions) (resp azfake.PagerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClientListByServiceResponse])
}

// NewPrivateEndpointConnectionsForMIPPolicySyncServerTransport creates a new instance of PrivateEndpointConnectionsForMIPPolicySyncServerTransport with the provided implementation.
// The returned PrivateEndpointConnectionsForMIPPolicySyncServerTransport instance is connected to an instance of armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateEndpointConnectionsForMIPPolicySyncServerTransport(srv *PrivateEndpointConnectionsForMIPPolicySyncServer) *PrivateEndpointConnectionsForMIPPolicySyncServerTransport {
	return &PrivateEndpointConnectionsForMIPPolicySyncServerTransport{
		srv:                   srv,
		beginCreateOrUpdate:   newTracker[azfake.PollerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClientCreateOrUpdateResponse]](),
		beginDelete:           newTracker[azfake.PollerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClientDeleteResponse]](),
		newListByServicePager: newTracker[azfake.PagerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClientListByServiceResponse]](),
	}
}

// PrivateEndpointConnectionsForMIPPolicySyncServerTransport connects instances of armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClient to instances of PrivateEndpointConnectionsForMIPPolicySyncServer.
// Don't use this type directly, use NewPrivateEndpointConnectionsForMIPPolicySyncServerTransport instead.
type PrivateEndpointConnectionsForMIPPolicySyncServerTransport struct {
	srv                   *PrivateEndpointConnectionsForMIPPolicySyncServer
	beginCreateOrUpdate   *tracker[azfake.PollerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClientCreateOrUpdateResponse]]
	beginDelete           *tracker[azfake.PollerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClientDeleteResponse]]
	newListByServicePager *tracker[azfake.PagerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClientListByServiceResponse]]
}

// Do implements the policy.Transporter interface for PrivateEndpointConnectionsForMIPPolicySyncServerTransport.
func (p *PrivateEndpointConnectionsForMIPPolicySyncServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateEndpointConnectionsForMIPPolicySyncClient.BeginCreateOrUpdate":
		resp, err = p.dispatchBeginCreateOrUpdate(req)
	case "PrivateEndpointConnectionsForMIPPolicySyncClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PrivateEndpointConnectionsForMIPPolicySyncClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateEndpointConnectionsForMIPPolicySyncClient.NewListByServicePager":
		resp, err = p.dispatchNewListByServicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionsForMIPPolicySyncServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armm365securityandcompliance.PrivateEndpointConnection](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, privateEndpointConnectionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		p.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionsForMIPPolicySyncServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, resourceNameParam, privateEndpointConnectionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		p.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionsForMIPPolicySyncServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, privateEndpointConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionsForMIPPolicySyncServerTransport) dispatchNewListByServicePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServicePager not implemented")}
	}
	newListByServicePager := p.newListByServicePager.get(req)
	if newListByServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByServicePager(resourceGroupNameParam, resourceNameParam, nil)
		newListByServicePager = &resp
		p.newListByServicePager.add(req, newListByServicePager)
		server.PagerResponderInjectNextLinks(newListByServicePager, req, func(page *armm365securityandcompliance.PrivateEndpointConnectionsForMIPPolicySyncClientListByServiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServicePager) {
		p.newListByServicePager.remove(req)
	}
	return resp, nil
}