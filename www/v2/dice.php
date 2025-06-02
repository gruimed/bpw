<?php

use OpenTelemetry\API\Globals;
use GuzzleHttp\Client;

class Dice {

    private $tracer;

    function __construct() {
        $tracerProvider = Globals::tracerProvider();
        $this->tracer = $tracerProvider->getTracer('otel-php-demo');
    }

    public function roll(int $rolls, string $load) {
        $result = [];
        for ($i = 0; $i < $rolls; $i++) {
            $result[] = $this->rollOnce($load);
        }
        return $result;
    }

    private function rollOnce(string $load) {

        if (strpos($load,'E') !== false) {
            $client = new Client();
            $res = $client->request('GET', 'http://echo-service:8088/payload?io_msec=10');
        }

        if (strpos($load,'C') !== false) {
            for ($i=0; $i<1_000_000; $i++) {
                $arr[] = 1_000_000 - $i;
            }
            sort($arr);
        }
        
        if (strpos($load,'D') !== false) {
            $mysqli = new mysqli('pinba', 'root' , '');
            $res = $mysqli->query('select now()');
            $res->fetch_row();
        }


        $span = $this->tracer
            ->spanBuilder("getRandom")
            ->startSpan();
        $result = random_int(1, 6);
        $span->end();
        return $result;
    }
}
