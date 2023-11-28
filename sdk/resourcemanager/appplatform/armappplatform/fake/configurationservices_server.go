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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ConfigurationServicesServer is a fake server for instances of the armappplatform.ConfigurationServicesClient type.
type ConfigurationServicesServer struct {
	// BeginCreateOrUpdate is the fake for method ConfigurationServicesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string, configurationServiceResource armappplatform.ConfigurationServiceResource, options *armappplatform.ConfigurationServicesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armappplatform.ConfigurationServicesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ConfigurationServicesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string, options *armappplatform.ConfigurationServicesClientBeginDeleteOptions) (resp azfake.PollerResponder[armappplatform.ConfigurationServicesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ConfigurationServicesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string, options *armappplatform.ConfigurationServicesClientGetOptions) (resp azfake.Responder[armappplatform.ConfigurationServicesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ConfigurationServicesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, serviceName string, options *armappplatform.ConfigurationServicesClientListOptions) (resp azfake.PagerResponder[armappplatform.ConfigurationServicesClientListResponse])

	// BeginValidate is the fake for method ConfigurationServicesClient.BeginValidate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginValidate func(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string, settings armappplatform.ConfigurationServiceSettings, options *armappplatform.ConfigurationServicesClientBeginValidateOptions) (resp azfake.PollerResponder[armappplatform.ConfigurationServicesClientValidateResponse], errResp azfake.ErrorResponder)
}

// NewConfigurationServicesServerTransport creates a new instance of ConfigurationServicesServerTransport with the provided implementation.
// The returned ConfigurationServicesServerTransport instance is connected to an instance of armappplatform.ConfigurationServicesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConfigurationServicesServerTransport(srv *ConfigurationServicesServer) *ConfigurationServicesServerTransport {
	return &ConfigurationServicesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armappplatform.ConfigurationServicesClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armappplatform.ConfigurationServicesClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armappplatform.ConfigurationServicesClientListResponse]](),
		beginValidate:       newTracker[azfake.PollerResponder[armappplatform.ConfigurationServicesClientValidateResponse]](),
	}
}

// ConfigurationServicesServerTransport connects instances of armappplatform.ConfigurationServicesClient to instances of ConfigurationServicesServer.
// Don't use this type directly, use NewConfigurationServicesServerTransport instead.
type ConfigurationServicesServerTransport struct {
	srv                 *ConfigurationServicesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armappplatform.ConfigurationServicesClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armappplatform.ConfigurationServicesClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armappplatform.ConfigurationServicesClientListResponse]]
	beginValidate       *tracker[azfake.PollerResponder[armappplatform.ConfigurationServicesClientValidateResponse]]
}

// Do implements the policy.Transporter interface for ConfigurationServicesServerTransport.
func (c *ConfigurationServicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ConfigurationServicesClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "ConfigurationServicesClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "ConfigurationServicesClient.Get":
		resp, err = c.dispatchGet(req)
	case "ConfigurationServicesClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	case "ConfigurationServicesClient.BeginValidate":
		resp, err = c.dispatchBeginValidate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ConfigurationServicesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/configurationServices/(?P<configurationServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappplatform.ConfigurationServiceResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		configurationServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, configurationServiceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		c.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		c.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (c *ConfigurationServicesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/configurationServices/(?P<configurationServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		configurationServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, serviceNameParam, configurationServiceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *ConfigurationServicesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/configurationServices/(?P<configurationServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	configurationServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationServiceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, configurationServiceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConfigurationServiceResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConfigurationServicesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/configurationServices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListPager(resourceGroupNameParam, serviceNameParam, nil)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappplatform.ConfigurationServicesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		c.newListPager.remove(req)
	}
	return resp, nil
}

func (c *ConfigurationServicesServerTransport) dispatchBeginValidate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginValidate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginValidate not implemented")}
	}
	beginValidate := c.beginValidate.get(req)
	if beginValidate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/configurationServices/(?P<configurationServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/validate`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappplatform.ConfigurationServiceSettings](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		configurationServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginValidate(req.Context(), resourceGroupNameParam, serviceNameParam, configurationServiceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginValidate = &respr
		c.beginValidate.add(req, beginValidate)
	}

	resp, err := server.PollerResponderNext(beginValidate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginValidate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginValidate) {
		c.beginValidate.remove(req)
	}

	return resp, nil
}