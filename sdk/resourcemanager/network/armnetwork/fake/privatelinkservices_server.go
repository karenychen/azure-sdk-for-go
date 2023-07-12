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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateLinkServicesServer is a fake server for instances of the armnetwork.PrivateLinkServicesClient type.
type PrivateLinkServicesServer struct {
	// BeginCheckPrivateLinkServiceVisibility is the fake for method PrivateLinkServicesClient.BeginCheckPrivateLinkServiceVisibility
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCheckPrivateLinkServiceVisibility func(ctx context.Context, location string, parameters armnetwork.CheckPrivateLinkServiceVisibilityRequest, options *armnetwork.PrivateLinkServicesClientBeginCheckPrivateLinkServiceVisibilityOptions) (resp azfake.PollerResponder[armnetwork.PrivateLinkServicesClientCheckPrivateLinkServiceVisibilityResponse], errResp azfake.ErrorResponder)

	// BeginCheckPrivateLinkServiceVisibilityByResourceGroup is the fake for method PrivateLinkServicesClient.BeginCheckPrivateLinkServiceVisibilityByResourceGroup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCheckPrivateLinkServiceVisibilityByResourceGroup func(ctx context.Context, location string, resourceGroupName string, parameters armnetwork.CheckPrivateLinkServiceVisibilityRequest, options *armnetwork.PrivateLinkServicesClientBeginCheckPrivateLinkServiceVisibilityByResourceGroupOptions) (resp azfake.PollerResponder[armnetwork.PrivateLinkServicesClientCheckPrivateLinkServiceVisibilityByResourceGroupResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method PrivateLinkServicesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, parameters armnetwork.PrivateLinkService, options *armnetwork.PrivateLinkServicesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.PrivateLinkServicesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PrivateLinkServicesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, serviceName string, options *armnetwork.PrivateLinkServicesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.PrivateLinkServicesClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDeletePrivateEndpointConnection is the fake for method PrivateLinkServicesClient.BeginDeletePrivateEndpointConnection
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDeletePrivateEndpointConnection func(ctx context.Context, resourceGroupName string, serviceName string, peConnectionName string, options *armnetwork.PrivateLinkServicesClientBeginDeletePrivateEndpointConnectionOptions) (resp azfake.PollerResponder[armnetwork.PrivateLinkServicesClientDeletePrivateEndpointConnectionResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateLinkServicesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, options *armnetwork.PrivateLinkServicesClientGetOptions) (resp azfake.Responder[armnetwork.PrivateLinkServicesClientGetResponse], errResp azfake.ErrorResponder)

	// GetPrivateEndpointConnection is the fake for method PrivateLinkServicesClient.GetPrivateEndpointConnection
	// HTTP status codes to indicate success: http.StatusOK
	GetPrivateEndpointConnection func(ctx context.Context, resourceGroupName string, serviceName string, peConnectionName string, options *armnetwork.PrivateLinkServicesClientGetPrivateEndpointConnectionOptions) (resp azfake.Responder[armnetwork.PrivateLinkServicesClientGetPrivateEndpointConnectionResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PrivateLinkServicesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.PrivateLinkServicesClientListOptions) (resp azfake.PagerResponder[armnetwork.PrivateLinkServicesClientListResponse])

	// NewListAutoApprovedPrivateLinkServicesPager is the fake for method PrivateLinkServicesClient.NewListAutoApprovedPrivateLinkServicesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAutoApprovedPrivateLinkServicesPager func(location string, options *armnetwork.PrivateLinkServicesClientListAutoApprovedPrivateLinkServicesOptions) (resp azfake.PagerResponder[armnetwork.PrivateLinkServicesClientListAutoApprovedPrivateLinkServicesResponse])

	// NewListAutoApprovedPrivateLinkServicesByResourceGroupPager is the fake for method PrivateLinkServicesClient.NewListAutoApprovedPrivateLinkServicesByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAutoApprovedPrivateLinkServicesByResourceGroupPager func(location string, resourceGroupName string, options *armnetwork.PrivateLinkServicesClientListAutoApprovedPrivateLinkServicesByResourceGroupOptions) (resp azfake.PagerResponder[armnetwork.PrivateLinkServicesClientListAutoApprovedPrivateLinkServicesByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method PrivateLinkServicesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armnetwork.PrivateLinkServicesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armnetwork.PrivateLinkServicesClientListBySubscriptionResponse])

	// NewListPrivateEndpointConnectionsPager is the fake for method PrivateLinkServicesClient.NewListPrivateEndpointConnectionsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPrivateEndpointConnectionsPager func(resourceGroupName string, serviceName string, options *armnetwork.PrivateLinkServicesClientListPrivateEndpointConnectionsOptions) (resp azfake.PagerResponder[armnetwork.PrivateLinkServicesClientListPrivateEndpointConnectionsResponse])

	// UpdatePrivateEndpointConnection is the fake for method PrivateLinkServicesClient.UpdatePrivateEndpointConnection
	// HTTP status codes to indicate success: http.StatusOK
	UpdatePrivateEndpointConnection func(ctx context.Context, resourceGroupName string, serviceName string, peConnectionName string, parameters armnetwork.PrivateEndpointConnection, options *armnetwork.PrivateLinkServicesClientUpdatePrivateEndpointConnectionOptions) (resp azfake.Responder[armnetwork.PrivateLinkServicesClientUpdatePrivateEndpointConnectionResponse], errResp azfake.ErrorResponder)
}

// NewPrivateLinkServicesServerTransport creates a new instance of PrivateLinkServicesServerTransport with the provided implementation.
// The returned PrivateLinkServicesServerTransport instance is connected to an instance of armnetwork.PrivateLinkServicesClient by way of the
// undefined.Transporter field.
func NewPrivateLinkServicesServerTransport(srv *PrivateLinkServicesServer) *PrivateLinkServicesServerTransport {
	return &PrivateLinkServicesServerTransport{srv: srv}
}

// PrivateLinkServicesServerTransport connects instances of armnetwork.PrivateLinkServicesClient to instances of PrivateLinkServicesServer.
// Don't use this type directly, use NewPrivateLinkServicesServerTransport instead.
type PrivateLinkServicesServerTransport struct {
	srv                                                        *PrivateLinkServicesServer
	beginCheckPrivateLinkServiceVisibility                     *azfake.PollerResponder[armnetwork.PrivateLinkServicesClientCheckPrivateLinkServiceVisibilityResponse]
	beginCheckPrivateLinkServiceVisibilityByResourceGroup      *azfake.PollerResponder[armnetwork.PrivateLinkServicesClientCheckPrivateLinkServiceVisibilityByResourceGroupResponse]
	beginCreateOrUpdate                                        *azfake.PollerResponder[armnetwork.PrivateLinkServicesClientCreateOrUpdateResponse]
	beginDelete                                                *azfake.PollerResponder[armnetwork.PrivateLinkServicesClientDeleteResponse]
	beginDeletePrivateEndpointConnection                       *azfake.PollerResponder[armnetwork.PrivateLinkServicesClientDeletePrivateEndpointConnectionResponse]
	newListPager                                               *azfake.PagerResponder[armnetwork.PrivateLinkServicesClientListResponse]
	newListAutoApprovedPrivateLinkServicesPager                *azfake.PagerResponder[armnetwork.PrivateLinkServicesClientListAutoApprovedPrivateLinkServicesResponse]
	newListAutoApprovedPrivateLinkServicesByResourceGroupPager *azfake.PagerResponder[armnetwork.PrivateLinkServicesClientListAutoApprovedPrivateLinkServicesByResourceGroupResponse]
	newListBySubscriptionPager                                 *azfake.PagerResponder[armnetwork.PrivateLinkServicesClientListBySubscriptionResponse]
	newListPrivateEndpointConnectionsPager                     *azfake.PagerResponder[armnetwork.PrivateLinkServicesClientListPrivateEndpointConnectionsResponse]
}

// Do implements the policy.Transporter interface for PrivateLinkServicesServerTransport.
func (p *PrivateLinkServicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateLinkServicesClient.BeginCheckPrivateLinkServiceVisibility":
		resp, err = p.dispatchBeginCheckPrivateLinkServiceVisibility(req)
	case "PrivateLinkServicesClient.BeginCheckPrivateLinkServiceVisibilityByResourceGroup":
		resp, err = p.dispatchBeginCheckPrivateLinkServiceVisibilityByResourceGroup(req)
	case "PrivateLinkServicesClient.BeginCreateOrUpdate":
		resp, err = p.dispatchBeginCreateOrUpdate(req)
	case "PrivateLinkServicesClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PrivateLinkServicesClient.BeginDeletePrivateEndpointConnection":
		resp, err = p.dispatchBeginDeletePrivateEndpointConnection(req)
	case "PrivateLinkServicesClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateLinkServicesClient.GetPrivateEndpointConnection":
		resp, err = p.dispatchGetPrivateEndpointConnection(req)
	case "PrivateLinkServicesClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "PrivateLinkServicesClient.NewListAutoApprovedPrivateLinkServicesPager":
		resp, err = p.dispatchNewListAutoApprovedPrivateLinkServicesPager(req)
	case "PrivateLinkServicesClient.NewListAutoApprovedPrivateLinkServicesByResourceGroupPager":
		resp, err = p.dispatchNewListAutoApprovedPrivateLinkServicesByResourceGroupPager(req)
	case "PrivateLinkServicesClient.NewListBySubscriptionPager":
		resp, err = p.dispatchNewListBySubscriptionPager(req)
	case "PrivateLinkServicesClient.NewListPrivateEndpointConnectionsPager":
		resp, err = p.dispatchNewListPrivateEndpointConnectionsPager(req)
	case "PrivateLinkServicesClient.UpdatePrivateEndpointConnection":
		resp, err = p.dispatchUpdatePrivateEndpointConnection(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateLinkServicesServerTransport) dispatchBeginCheckPrivateLinkServiceVisibility(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCheckPrivateLinkServiceVisibility == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCheckPrivateLinkServiceVisibility not implemented")}
	}
	if p.beginCheckPrivateLinkServiceVisibility == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/checkPrivateLinkServiceVisibility`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.CheckPrivateLinkServiceVisibilityRequest](req)
		if err != nil {
			return nil, err
		}
		locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCheckPrivateLinkServiceVisibility(req.Context(), locationUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginCheckPrivateLinkServiceVisibility = &respr
	}

	resp, err := server.PollerResponderNext(p.beginCheckPrivateLinkServiceVisibility, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginCheckPrivateLinkServiceVisibility) {
		p.beginCheckPrivateLinkServiceVisibility = nil
	}

	return resp, nil
}

func (p *PrivateLinkServicesServerTransport) dispatchBeginCheckPrivateLinkServiceVisibilityByResourceGroup(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCheckPrivateLinkServiceVisibilityByResourceGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCheckPrivateLinkServiceVisibilityByResourceGroup not implemented")}
	}
	if p.beginCheckPrivateLinkServiceVisibilityByResourceGroup == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/checkPrivateLinkServiceVisibility`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.CheckPrivateLinkServiceVisibilityRequest](req)
		if err != nil {
			return nil, err
		}
		locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCheckPrivateLinkServiceVisibilityByResourceGroup(req.Context(), locationUnescaped, resourceGroupNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginCheckPrivateLinkServiceVisibilityByResourceGroup = &respr
	}

	resp, err := server.PollerResponderNext(p.beginCheckPrivateLinkServiceVisibilityByResourceGroup, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginCheckPrivateLinkServiceVisibilityByResourceGroup) {
		p.beginCheckPrivateLinkServiceVisibilityByResourceGroup = nil
	}

	return resp, nil
}

func (p *PrivateLinkServicesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if p.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/privateLinkServices/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.PrivateLinkService](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, serviceNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(p.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginCreateOrUpdate) {
		p.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (p *PrivateLinkServicesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if p.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/privateLinkServices/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, serviceNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(p.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginDelete) {
		p.beginDelete = nil
	}

	return resp, nil
}

func (p *PrivateLinkServicesServerTransport) dispatchBeginDeletePrivateEndpointConnection(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDeletePrivateEndpointConnection == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeletePrivateEndpointConnection not implemented")}
	}
	if p.beginDeletePrivateEndpointConnection == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/privateLinkServices/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<peConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		peConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("peConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDeletePrivateEndpointConnection(req.Context(), resourceGroupNameUnescaped, serviceNameUnescaped, peConnectionNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginDeletePrivateEndpointConnection = &respr
	}

	resp, err := server.PollerResponderNext(p.beginDeletePrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginDeletePrivateEndpointConnection) {
		p.beginDeletePrivateEndpointConnection = nil
	}

	return resp, nil
}

func (p *PrivateLinkServicesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/privateLinkServices/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armnetwork.PrivateLinkServicesClientGetOptions
	if expandParam != nil {
		options = &armnetwork.PrivateLinkServicesClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameUnescaped, serviceNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateLinkService, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateLinkServicesServerTransport) dispatchGetPrivateEndpointConnection(req *http.Request) (*http.Response, error) {
	if p.srv.GetPrivateEndpointConnection == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetPrivateEndpointConnection not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/privateLinkServices/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<peConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	peConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("peConnectionName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armnetwork.PrivateLinkServicesClientGetPrivateEndpointConnectionOptions
	if expandParam != nil {
		options = &armnetwork.PrivateLinkServicesClientGetPrivateEndpointConnectionOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := p.srv.GetPrivateEndpointConnection(req.Context(), resourceGroupNameUnescaped, serviceNameUnescaped, peConnectionNameUnescaped, options)
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

func (p *PrivateLinkServicesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if p.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/privateLinkServices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(resourceGroupNameUnescaped, nil)
		p.newListPager = &resp
		server.PagerResponderInjectNextLinks(p.newListPager, req, func(page *armnetwork.PrivateLinkServicesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListPager) {
		p.newListPager = nil
	}
	return resp, nil
}

func (p *PrivateLinkServicesServerTransport) dispatchNewListAutoApprovedPrivateLinkServicesPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListAutoApprovedPrivateLinkServicesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAutoApprovedPrivateLinkServicesPager not implemented")}
	}
	if p.newListAutoApprovedPrivateLinkServicesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/autoApprovedPrivateLinkServices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListAutoApprovedPrivateLinkServicesPager(locationUnescaped, nil)
		p.newListAutoApprovedPrivateLinkServicesPager = &resp
		server.PagerResponderInjectNextLinks(p.newListAutoApprovedPrivateLinkServicesPager, req, func(page *armnetwork.PrivateLinkServicesClientListAutoApprovedPrivateLinkServicesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListAutoApprovedPrivateLinkServicesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListAutoApprovedPrivateLinkServicesPager) {
		p.newListAutoApprovedPrivateLinkServicesPager = nil
	}
	return resp, nil
}

func (p *PrivateLinkServicesServerTransport) dispatchNewListAutoApprovedPrivateLinkServicesByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListAutoApprovedPrivateLinkServicesByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAutoApprovedPrivateLinkServicesByResourceGroupPager not implemented")}
	}
	if p.newListAutoApprovedPrivateLinkServicesByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/autoApprovedPrivateLinkServices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListAutoApprovedPrivateLinkServicesByResourceGroupPager(locationUnescaped, resourceGroupNameUnescaped, nil)
		p.newListAutoApprovedPrivateLinkServicesByResourceGroupPager = &resp
		server.PagerResponderInjectNextLinks(p.newListAutoApprovedPrivateLinkServicesByResourceGroupPager, req, func(page *armnetwork.PrivateLinkServicesClientListAutoApprovedPrivateLinkServicesByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListAutoApprovedPrivateLinkServicesByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListAutoApprovedPrivateLinkServicesByResourceGroupPager) {
		p.newListAutoApprovedPrivateLinkServicesByResourceGroupPager = nil
	}
	return resp, nil
}

func (p *PrivateLinkServicesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	if p.newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/privateLinkServices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListBySubscriptionPager(nil)
		p.newListBySubscriptionPager = &resp
		server.PagerResponderInjectNextLinks(p.newListBySubscriptionPager, req, func(page *armnetwork.PrivateLinkServicesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListBySubscriptionPager) {
		p.newListBySubscriptionPager = nil
	}
	return resp, nil
}

func (p *PrivateLinkServicesServerTransport) dispatchNewListPrivateEndpointConnectionsPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPrivateEndpointConnectionsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPrivateEndpointConnectionsPager not implemented")}
	}
	if p.newListPrivateEndpointConnectionsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/privateLinkServices/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPrivateEndpointConnectionsPager(resourceGroupNameUnescaped, serviceNameUnescaped, nil)
		p.newListPrivateEndpointConnectionsPager = &resp
		server.PagerResponderInjectNextLinks(p.newListPrivateEndpointConnectionsPager, req, func(page *armnetwork.PrivateLinkServicesClientListPrivateEndpointConnectionsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListPrivateEndpointConnectionsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListPrivateEndpointConnectionsPager) {
		p.newListPrivateEndpointConnectionsPager = nil
	}
	return resp, nil
}

func (p *PrivateLinkServicesServerTransport) dispatchUpdatePrivateEndpointConnection(req *http.Request) (*http.Response, error) {
	if p.srv.UpdatePrivateEndpointConnection == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdatePrivateEndpointConnection not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/privateLinkServices/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<peConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.PrivateEndpointConnection](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	peConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("peConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.UpdatePrivateEndpointConnection(req.Context(), resourceGroupNameUnescaped, serviceNameUnescaped, peConnectionNameUnescaped, body, nil)
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