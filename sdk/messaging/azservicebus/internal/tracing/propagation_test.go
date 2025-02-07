package tracing

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/tracing/azotel"
	"github.com/Azure/go-amqp"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/propagation"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func Test_TracingMiddleware(t *testing.T) {
	testCases := []struct {
		description     string
		message         *amqp.Message
		carryRemoteSpan bool
	}{
		{
			description:     "nil message, context should not contain remote tracecontext",
			message:         nil,
			carryRemoteSpan: false,
		},
		{
			description: "context should contain remote tracecontext",
			message: &amqp.Message{
				Properties: &amqp.MessageProperties{
					MessageID: "message-id",
				},
			},
			carryRemoteSpan: true,
		},
	}
	otelTP := tracesdk.NewTracerProvider(tracesdk.WithSampler(tracesdk.AlwaysSample()))
	tp := azotel.NewTracingProvider(otelTP, nil)
	tracer := tp.NewTracer("test-tracer", "1.0.0")
	propagator := tp.NewPropagator()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.carryRemoteSpan {
				testRemoteCtx, testRemoteSpan := tracer.Start(context.Background(), "remote-span", nil)
				testRemoteSpan.End()

				p := propagation.TraceContext{}
				p.Inject(testRemoteCtx, MessageCarrierAdapter(tc.message))
			}

			ctx := Extract(context.TODO(), propagator, tc.message)

			// this works because TraceContext.Extract() sets the extracted tracecontext as the remote SpanContext of returned Context
			sp := tracer.SpanFromContext(ctx)
			spCtx := sp.SpanContext()
			require.NotNil(t, spCtx)
			require.EqualValues(t, tc.carryRemoteSpan, spCtx.IsRemote())
		})
	}
}

func Test_TracingMiddlewareLinks(t *testing.T) {
	testCases := []struct {
		description     string
		message         *amqp.Message
		carryRemoteSpan bool
	}{
		{
			description:     "nil message, context should not contain remote tracecontext",
			message:         nil,
			carryRemoteSpan: false,
		},
		{
			description: "context should contain remote tracecontext",
			message: &amqp.Message{
				Properties: &amqp.MessageProperties{
					MessageID: "message-id",
				},
			},
			carryRemoteSpan: true,
		},
	}
	otelTP := tracesdk.NewTracerProvider(tracesdk.WithSampler(tracesdk.AlwaysSample()))
	tp := azotel.NewTracingProvider(otelTP, nil)
	tracer := tp.NewTracer("test-tracer", "1.0.0")
	propagator := tp.NewPropagator()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.carryRemoteSpan {
				testRemoteCtx, testRemoteSpan := tracer.Start(context.Background(), "remote-span", nil)
				testRemoteSpan.End()

				p := propagation.TraceContext{}
				p.Inject(testRemoteCtx, MessageCarrierAdapter(tc.message))
			}

			ctx := Extract(context.TODO(), propagator, tc.message)

			// this works because TraceContext.Extract() sets the extracted tracecontext as the remote SpanContext of returned Context
			sp := tracer.SpanFromContext(ctx)
			spCtx := sp.SpanContext()
			require.NotNil(t, spCtx)
			require.EqualValues(t, tc.carryRemoteSpan, spCtx.IsRemote())

			link := tracer.LinkFromContext(ctx)
			require.NotNil(t, link)
			require.EqualValues(t, spCtx.SpanID(), link.SpanContext.SpanID())
		})
	}
}
