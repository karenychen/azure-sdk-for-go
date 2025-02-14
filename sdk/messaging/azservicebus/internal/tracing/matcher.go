// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package tracing

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"github.com/stretchr/testify/require"
)

type Span = tracing.Span
type SpanStatus = tracing.SpanStatus

// NewSpanValidator creates a Provider that verifies a span was created that matches the specified SpanMatcher.
// The returned Provider can be used to create a client with a tracing provider that will validate spans in unit tests.
func NewSpanValidator(t *testing.T, matcher SpanMatcher) Provider {
	return tracing.NewProvider(func(name, version string) Tracer {
		tt := matchingTracer{
			matcher: matcher,
		}

		t.Cleanup(func() {
			require.NotNil(t, tt.match, "didn't find a span with name %s", tt.matcher.Name)
			require.True(t, tt.match.ended, "span wasn't ended")
			if tt.matcher.Kind != 0 {
				require.EqualValues(t, tt.matcher.Kind, tt.match.kind, "span kind values don't match")
			}
			require.EqualValues(t, matcher.Status, tt.match.status, "span status values don't match")
			require.ElementsMatch(t, matcher.Attributes, tt.match.attrs, "span attributes don't match")
			require.ElementsMatch(t, matcher.Links, tt.match.links, "span links don't match")
		})

		return tracing.NewTracer(func(ctx context.Context, spanName string, options *tracing.SpanOptions) (context.Context, Span) {
			kind := tracing.SpanKindInternal
			attrs := []Attribute{}
			links := []Link{}
			if options != nil {
				kind = options.Kind
				attrs = append(attrs, options.Attributes...)
				links = append(links, options.Links...)
			}
			return tt.Start(ctx, spanName, kind, attrs, links)
		}, &tracing.TracerOptions{
			SpanFromContext: func(ctx context.Context) Span {
				return convertSpan(tt.match)
			},
			LinkFromContext: func(ctx context.Context, attrs ...Attribute) Link {
				return tracing.Link{Attributes: attrs}
			},
		})
	}, nil)
}

// SpanMatcher contains the values to match when a span is created.
type SpanMatcher struct {
	Name       string
	Kind       SpanKind
	Status     SpanStatus
	Attributes []Attribute
	Links      []Link
}

type matchingTracer struct {
	matcher SpanMatcher
	match   *span
}

func (mt *matchingTracer) Start(ctx context.Context, spanName string, kind SpanKind, attrs []Attribute, links []Link) (context.Context, Span) {
	if spanName != mt.matcher.Name {
		return ctx, Span{}
	}
	// span name matches our matcher, track it
	mt.match = &span{
		name:  spanName,
		kind:  kind,
		spCtx: SpanContext{},
		attrs: attrs,
		links: links,
	}

	return ctx, convertSpan(mt.match)
}

func convertSpan(sp *span) Span {
	return tracing.NewSpan(tracing.SpanImpl{
		End:           sp.End,
		SpanContext:   sp.SpanContext,
		SetStatus:     sp.SetStatus,
		SetAttributes: sp.SetAttributes,
		AddLink:       sp.AddLink,
	})
}

type span struct {
	name   string
	kind   SpanKind
	status SpanStatus
	spCtx  tracing.SpanContext
	attrs  []Attribute
	links  []Link
	desc   string
	ended  bool
}

func (s *span) End() {
	s.ended = true
}

func (s *span) SpanContext() tracing.SpanContext {
	return s.spCtx
}

func (s *span) SetAttributes(attrs ...Attribute) {
	s.attrs = append(s.attrs, attrs...)
}

func (s *span) AddLink(link Link) {
	s.links = append(s.links, link)
}

func (s *span) SetStatus(code SpanStatus, desc string) {
	s.status = code
	s.desc = desc
	s.ended = true
}
