<?php

declare(strict_types=1);

use OpenTelemetry\Context\Context;
use OpenTelemetry\Context\ContextInterface;
use OpenTelemetry\SDK\Trace\ReadableSpanInterface;
use OpenTelemetry\SDK\Trace\ReadWriteSpanInterface;
use OpenTelemetry\SDK\Trace\SpanExporterInterface;
use OpenTelemetry\SDK\Trace\SpanProcessorInterface;
use OpenTelemetry\SDK\Common\Future\CancellationInterface;
use OpenTelemetry\Context\ContextKeys;
use OpenTelemetry\API\Trace\SpanKind;


class MySpanProcessor implements SpanProcessorInterface
{

    public function __construct()
    {
    }

    public function onStart(ReadWriteSpanInterface $span, ContextInterface $parentContext): void
    {
        $parent = $parentContext->get(ContextKeys::span());
        if (!$parent || !$parent->isRecording()) {
            return; 
        }

        if ($parent->getKind() == SpanKind::KIND_SERVER) {
            $span->setAttribute('service.endpoint', $parent->getName());
        } else {
            $span->setAttribute('service.endpoint', $parent->getAttribute('service.endpoint'));
        }
    }

    public function onEnd(ReadableSpanInterface $span): void
    {
    }

    public function forceFlush(?CancellationInterface $cancellation = null): bool
    {

        return true;
    }

    public function shutdown(?CancellationInterface $cancellation = null): bool
    {
        return true;
    }

    private function flush(Closure $task, string $taskName, bool $propagateResult, ContextInterface $context): bool
    {
        return true;
    }
}
