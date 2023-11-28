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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armproviderhub.ClientFactory type.
type ServerFactory struct {
	Server                          Server
	CustomRolloutsServer            CustomRolloutsServer
	DefaultRolloutsServer           DefaultRolloutsServer
	NotificationRegistrationsServer NotificationRegistrationsServer
	OperationsServer                OperationsServer
	ProviderRegistrationsServer     ProviderRegistrationsServer
	ResourceTypeRegistrationsServer ResourceTypeRegistrationsServer
	SKUsServer                      SKUsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armproviderhub.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armproviderhub.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                               *ServerFactory
	trMu                              sync.Mutex
	trServer                          *ServerTransport
	trCustomRolloutsServer            *CustomRolloutsServerTransport
	trDefaultRolloutsServer           *DefaultRolloutsServerTransport
	trNotificationRegistrationsServer *NotificationRegistrationsServerTransport
	trOperationsServer                *OperationsServerTransport
	trProviderRegistrationsServer     *ProviderRegistrationsServerTransport
	trResourceTypeRegistrationsServer *ResourceTypeRegistrationsServerTransport
	trSKUsServer                      *SKUsServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "Client":
		initServer(s, &s.trServer, func() *ServerTransport { return NewServerTransport(&s.srv.Server) })
		resp, err = s.trServer.Do(req)
	case "CustomRolloutsClient":
		initServer(s, &s.trCustomRolloutsServer, func() *CustomRolloutsServerTransport {
			return NewCustomRolloutsServerTransport(&s.srv.CustomRolloutsServer)
		})
		resp, err = s.trCustomRolloutsServer.Do(req)
	case "DefaultRolloutsClient":
		initServer(s, &s.trDefaultRolloutsServer, func() *DefaultRolloutsServerTransport {
			return NewDefaultRolloutsServerTransport(&s.srv.DefaultRolloutsServer)
		})
		resp, err = s.trDefaultRolloutsServer.Do(req)
	case "NotificationRegistrationsClient":
		initServer(s, &s.trNotificationRegistrationsServer, func() *NotificationRegistrationsServerTransport {
			return NewNotificationRegistrationsServerTransport(&s.srv.NotificationRegistrationsServer)
		})
		resp, err = s.trNotificationRegistrationsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "ProviderRegistrationsClient":
		initServer(s, &s.trProviderRegistrationsServer, func() *ProviderRegistrationsServerTransport {
			return NewProviderRegistrationsServerTransport(&s.srv.ProviderRegistrationsServer)
		})
		resp, err = s.trProviderRegistrationsServer.Do(req)
	case "ResourceTypeRegistrationsClient":
		initServer(s, &s.trResourceTypeRegistrationsServer, func() *ResourceTypeRegistrationsServerTransport {
			return NewResourceTypeRegistrationsServerTransport(&s.srv.ResourceTypeRegistrationsServer)
		})
		resp, err = s.trResourceTypeRegistrationsServer.Do(req)
	case "SKUsClient":
		initServer(s, &s.trSKUsServer, func() *SKUsServerTransport { return NewSKUsServerTransport(&s.srv.SKUsServer) })
		resp, err = s.trSKUsServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}