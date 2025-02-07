// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package tracing

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"github.com/stretchr/testify/require"
)

func TestStartSpan(t *testing.T) {
	// no-op when StartSpanOptions is nil
	ctx := context.Background()
	subCtx, _ := StartSpan(ctx, nil)
	require.Equal(t, ctx, subCtx)

	// no-op when StartSpanOptions is empty
	subCtx, _ = StartSpan(ctx, &StartSpanOptions{})
	require.Equal(t, ctx, subCtx)

	// no-op when SpanName is empty
	subCtx, _ = StartSpan(ctx, &StartSpanOptions{SpanName: ""})
	require.Equal(t, ctx, subCtx)

	// creates a span when both tracer and SpanName are set
	tracer := NewSpanValidator(t, SpanMatcher{
		Name: "test",
		Kind: SpanKindInternal,
	}).NewTracer("module", "version")
	subCtx1, endSpan1 := StartSpan(ctx, &StartSpanOptions{Tracer: tracer, SpanName: "test"})
	defer endSpan1(nil)
	require.NotEqual(t, ctx, subCtx1)

	// creates a span when span name has sender prefix
	tracer = NewSpanValidator(t, SpanMatcher{
		Name: "Sender.TestSpan",
		Kind: SpanKindProducer,
	}).NewTracer("module", "version")
	subCtx2, endSpan2 := StartSpan(ctx, &StartSpanOptions{Tracer: tracer, SpanName: "Sender.TestSpan"})
	defer endSpan2(nil)
	require.NotEqual(t, ctx, subCtx2)
}

func TestGetEntityPathAttributes(t *testing.T) {
	entityPath := "queueName"
	expectedAttrs := []tracing.Attribute{
		{Key: DestinationName, Value: entityPath},
	}
	result := GetEntityPathAttributes(entityPath)
	require.ElementsMatch(t, expectedAttrs, result)

	entityPath = "topicName/Subscriptions/subscriptionName"
	expectedAttrs = []tracing.Attribute{
		{Key: DestinationName, Value: "topicName"},
		{Key: SubscriptionName, Value: "subscriptionName"},
	}
	result = GetEntityPathAttributes(entityPath)
	require.ElementsMatch(t, expectedAttrs, result)
}
