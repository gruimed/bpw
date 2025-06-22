package main

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type parentSpanDataKey struct{}

type ParentSpanData struct {
	Name       string
	Kind       string
	Attributes []attribute.KeyValue
}

type myTracer struct{ trace.Tracer }

func (wt *myTracer) Start(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {

	cfg := trace.NewSpanStartConfig(opts...)

	if isSet := ctx.Value(parentSpanDataKey{}); isSet == nil { // No data is set yet.
		data := &ParentSpanData{
			Name: name,
			Kind: cfg.SpanKind().String(),
			Attributes: []attribute.KeyValue{
				{
					Key:   "tracer.span.name",
					Value: attribute.StringValue(name),
				},
				{
					Key:   "tracer.span.kind",
					Value: attribute.StringValue(cfg.SpanKind().String()),
				},
			},
		}

		serviceEndpoint := "not-set"

		if oldData, ok := isSet.(*ParentSpanData); ok {
			set := attribute.NewSet(oldData.Attributes...)
			val, _ := set.Value("service.endpoint")
			serviceEndpoint = val.AsString()
		}

		// Pay attention that otel libs can mix span kinds.
		if cfg.SpanKind() == trace.SpanKindServer {
			serviceEndpoint = name
		}

		data.Attributes = append(
			data.Attributes,
			attribute.KeyValue{
				Key:   "service.endpoint",
				Value: attribute.StringValue(serviceEndpoint),
			},
		)

		ctx = context.WithValue(ctx, parentSpanDataKey{}, data)
	}

	return wt.Tracer.Start(ctx, name, opts...)
}
