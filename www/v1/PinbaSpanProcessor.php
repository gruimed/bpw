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


class PinbaSpanProcessor implements SpanProcessorInterface
{
    private array $pinbaTimers = [];

    public function __construct()
    {
    }

    public function onStart(ReadWriteSpanInterface $span, ContextInterface $parentContext): void
    {
        $kind = $span->toSpanData()->getKind();
        if ($kind == SpanKind::KIND_SERVER) {
            return;
        }

        $span_id = $span->toSpanData()->getSpanId();

        $tags = [
            'span' => $span->getName(),
        ];
        $this->pinbaTimers[$span_id] = pinba_timer_start($tags);
    }

    public function onEnd(ReadableSpanInterface $span): void
    {
        $span_id = $span->toSpanData()->getSpanId();
        pinba_timer_stop($this->pinbaTimers[$span_id]);
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
