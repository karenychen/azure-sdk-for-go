// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package tracing

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"github.com/Azure/go-amqp"
)

// Inject injects the trace context from ctx into the message
func Inject(ctx context.Context, propagator tracing.Propagator, msg *amqp.Message) {
	propagator.Inject(ctx, MessageCarrierAdapter(msg))
}

// Extract extracts the remote trace context from the message and set it onto ctx
func Extract(ctx context.Context, propagator tracing.Propagator, message *amqp.Message) context.Context {
	if message != nil {
		ctx = propagator.Extract(ctx, MessageCarrierAdapter(message))
	}
	return ctx
}

// messageWrapper implements the TextMapCarrier interface for sender side
type messageWrapper struct {
	message *amqp.Message
}

// MessageCarrierAdapter wraps a Message so that it implements the propagation.TextMapCarrier interface
func MessageCarrierAdapter(message *amqp.Message) tracing.Carrier {
	return &messageWrapper{message: message}
}

func (mw *messageWrapper) Set(key string, value string) {
	if mw.message.ApplicationProperties == nil {
		mw.message.ApplicationProperties = make(map[string]interface{})
	}
	mw.message.ApplicationProperties[key] = value
}

func (mw *messageWrapper) Get(key string) string {
	if mw.message.ApplicationProperties == nil || mw.message.ApplicationProperties[key] == nil {
		return ""
	}
	return mw.message.ApplicationProperties[key].(string)
}

func (mw *messageWrapper) Keys() []string {
	keys := make([]string, 0, len(mw.message.ApplicationProperties))
	for k := range mw.message.ApplicationProperties {
		keys = append(keys, k)
	}
	return keys
}
