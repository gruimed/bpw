package main

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"log"
)

type mySpanProcessor struct {
}

func (s2 mySpanProcessor) OnStart(ctx context.Context, current sdkTrace.ReadWriteSpan) {
	log.Println("mySpanProcessor start")

	parent := trace.SpanFromContext(ctx)
	/*if !parent.IsRecording() {
		return
	}*/

	attrs := []attribute.KeyValue{
		{
			Key:   "sp.parent.SpanID",
			Value: attribute.StringValue(parent.SpanContext().SpanID().String()),
		},
		{
			Key:   "sp.parent.span.is-recording",
			Value: attribute.BoolValue(parent.IsRecording()),
		},
		{
			Key:   "sp.current.SpanID",
			Value: attribute.StringValue(current.SpanContext().SpanID().String()),
		},
	}

	if parentData := ctx.Value(parentSpanDataKey{}); parentData != nil {
		data, ok := parentData.(*ParentSpanData)
		if !ok {
			log.Println("mySpanProcessor parentData can't be type casted")
			attrs = append(attrs, attribute.KeyValue{
				Key:   "sp.parent.no-typecast.SpanID",
				Value: attribute.StringValue(current.SpanContext().SpanID().String()),
			})
		}
		attrs = append(attrs, data.Attributes...)
	} else {
		attrs = append(attrs, attribute.KeyValue{
			Key:   "sp.parent.no-data.SpanID",
			Value: attribute.StringValue(current.SpanContext().SpanID().String()),
		})
	}

	current.SetAttributes(attrs...)

	log.Println("mySpanProcessor stop")
}

func (s2 mySpanProcessor) OnEnd(s sdkTrace.ReadOnlySpan) {

}

func (s2 mySpanProcessor) Shutdown(ctx context.Context) error {
	return nil
}

func (s2 mySpanProcessor) ForceFlush(ctx context.Context) error {
	return nil
}
