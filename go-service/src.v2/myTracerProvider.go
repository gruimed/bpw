package main

import (
	"go.opentelemetry.io/otel/trace"
)

type MyTracerProvider struct {
	trace.TracerProvider
}

func (m MyTracerProvider) Tracer(name string, options ...trace.TracerOption) trace.Tracer {
	trcr := m.TracerProvider.Tracer(name, options...)

	return &myTracer{trcr}
}
