<?php

use OpenTelemetry\SDK\Sdk;
use OpenTelemetry\SDK\Trace\SpanProcessor\BatchSpanProcessor;
use OpenTelemetry\SDK\Trace\TracerProvider;
use OpenTelemetry\API\Trace\Propagation\TraceContextPropagator;
use OpenTelemetry\API\Common\Time\Clock;
use OpenTelemetry\Contrib\Otlp\SpanExporterFactory;

require __DIR__ . '/../vendor/autoload.php';
//require 'MySpanProcessor.php';

$spanExporter = (new SpanExporterFactory())->create();

$tracerProvider = TracerProvider::builder()
    ->addSpanProcessor(
        new BatchSpanProcessor(
            $spanExporter,
            Clock::getDefault()
        )
    )
//    ->addSpanProcessor(
//        new MySpanProcessor(),
//    )
    ->build();


Sdk::builder()
    ->setTracerProvider($tracerProvider)
    ->setPropagator(TraceContextPropagator::getInstance())
    ->setAutoShutdown(true)
    ->buildAndRegisterGlobal();
