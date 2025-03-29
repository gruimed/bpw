<?php

use GuzzleHttp\Client;

class Dice {


    function __construct() {
    }

    public function roll($rolls) {
        $result = [];
        for ($i = 0; $i < $rolls; $i++) {
            $result[] = $this->rollOnce();
        }
        return $result;
    }

    private function rollOnce() {
        $client = new Client();
        $res = $client->request('GET', 'http://echo-service:8088/payload?io_msec=10');
#        for ($i=0; $i<1_000_000; $i++) {
#            $arr[] = 1_000_000 - $i;
#        }

        $result = random_int(1, 6);
        return $result;
    }
}
