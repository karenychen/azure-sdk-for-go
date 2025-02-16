// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package tracing

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"github.com/Azure/go-amqp"
)

const messagingSystemName = "servicebus"

type Provider = tracing.Provider
type Attribute = tracing.Attribute
type Link = tracing.Link
type Propagator = tracing.Propagator
type Carrier = tracing.Carrier

type Tracer struct {
	tracer      tracing.Tracer
	propagator  tracing.Propagator
	destination string
}

type StartSpanOptions struct {
	Tracer        Tracer
	OperationName MessagingOperationName
	Attributes    []Attribute
}

func NewTracer(provider Provider, moduleName, version, hostName, queueOrTopic, subscription string) Tracer {
	t := Tracer{
		tracer:      provider.NewTracer(moduleName, version),
		propagator:  provider.NewPropagator(),
		destination: queueOrTopic,
	}
	t.tracer.SetAttributes(Attribute{Key: MessagingSystem, Value: messagingSystemName},
		Attribute{Key: DestinationName, Value: queueOrTopic})
	if hostName != "" {
		t.tracer.SetAttributes(Attribute{Key: ServerAddress, Value: hostName})
	}
	if subscription != "" {
		t.tracer.SetAttributes(Attribute{Key: SubscriptionName, Value: subscription})
	}
	return t
}

func (t *Tracer) SpanFromContext(ctx context.Context) Span {
	return t.tracer.SpanFromContext(ctx)
}

func (t *Tracer) LinkFromContext(ctx context.Context, attrs ...Attribute) Link {
	return t.tracer.LinkFromContext(ctx, attrs...)
}

func (t *Tracer) Inject(ctx context.Context, message *amqp.Message) {
	t.propagator.Inject(ctx, messageCarrierAdapter(message))
}

func (t *Tracer) Extract(ctx context.Context, message *amqp.Message) context.Context {
	if message != nil {
		ctx = t.propagator.Extract(ctx, messageCarrierAdapter(message))
	}
	return ctx
}

func StartSpan(ctx context.Context, options *StartSpanOptions) (context.Context, func(error)) {
	if options == nil || options.OperationName == "" {
		return ctx, func(error) {}
	}

	tr := options.Tracer
	operationType := getOperationType(options.OperationName)
	spanKind := getSpanKind(operationType, options.Attributes)
	spanName := fmt.Sprintf("%s %s", options.OperationName, tr.destination)

	attrs := options.Attributes
	if operationType == SettleOperationType {
		attrs = append(attrs, Attribute{Key: DispositionStatus, Value: string(options.OperationName)})
	}

	attrs = append(attrs,
		Attribute{Key: OperationType, Value: string(operationType)},
		Attribute{Key: OperationName, Value: string(options.OperationName)})

	return runtime.StartSpan(ctx, spanName, tr.tracer,
		&runtime.StartSpanOptions{
			Kind:       spanKind,
			Attributes: attrs,
		})
}

//
//// StartSpan creates a span with the specified name and attributes.
//// If no span name is provided, no span is created.
//func StartSpan(ctx context.Context, options *StartSpanOptions) (context.Context, func(error)) {
//	if options == nil || !options.Tracer.Enabled() {
//		return ctx, func(error) {}
//	}
//	operationType := getOperationType(options.OperationName)
//	spanKind := getSpanKind(operationType, options.Attributes)
//	spanName := fmt.Sprintf("%s %s", options.OperationName, queueOrTopic)
//
//	var attrs []Attribute
//	if operationType == SettleOperationType {
//		attrs = append(attrs, Attribute{Key: DispositionStatus, Value: string(options.OperationName)})
//	}
//
//	options.Attributes = append(options.Attributes,
//		Attribute{Key: OperationType, Value: string(operationType)},
//		Attribute{Key: OperationName, Value: string(operationName)})
//
//	return runtime.StartSpan(ctx, spanName, options.Tracer,
//		&runtime.StartSpanOptions{
//			Kind:       spanKind,
//			Attributes: options.Attributes,
//		})
//}
//
//func getEntityPathAttributes(entityPath string) []tracing.Attribute {
//	var attrs []tracing.Attribute
//	queueOrTopic, subscription := splitEntityPath(entityPath)
//	if queueOrTopic != "" {
//		attrs = append(attrs, tracing.Attribute{Key: DestinationName, Value: queueOrTopic})
//	}
//	if subscription != "" {
//		attrs = append(attrs, tracing.Attribute{Key: SubscriptionName, Value: subscription})
//	}
//	return attrs
//}
//
//func splitEntityPath(entityPath string) (string, string) {
//	queueOrTopic := ""
//	subscription := ""
//
//	path := strings.Split(entityPath, "/")
//	if len(path) >= 1 {
//		queueOrTopic = path[0]
//	}
//	if len(path) >= 3 && path[1] == "Subscriptions" {
//		subscription = path[2]
//	}
//	return queueOrTopic, subscription
//}

func getOperationType(operationName MessagingOperationName) MessagingOperationType {
	switch operationName {
	case CreateOperationName:
		return CreateOperationType
	case SendOperationName, ScheduleOperationName, CancelScheduledOperationName:
		return SendOperationType
	case ReceiveOperationName, PeekOperationName, ReceiveDeferredOperationName, RenewMessageLockOperationName,
		AcceptSessionOperationName, GetSessionStateOperationName, SetSessionStateOperationName, RenewSessionLockOperationName:
		return ReceiveOperationType
	case AbandonOperationName, CompleteOperationName, DeferOperationName, DeadLetterOperationName:
		return SettleOperationType
	default:
		return ""
	}
}

func getSpanKind(operationType MessagingOperationType, attrs []Attribute) SpanKind {
	switch operationType {
	case CreateOperationType:
		return SpanKindProducer
	case SendOperationType:
		// return client span if it is a batch operation
		// otherwise return producer span
		for _, attr := range attrs {
			if attr.Key == BatchMessageCount {
				return SpanKindClient
			}
		}
		return SpanKindProducer
	case ReceiveOperationType:
		return SpanKindClient
	case SettleOperationType:
		return SpanKindClient
	default:
		return SpanKindInternal
	}
}
