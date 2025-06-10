<?php

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Factory\AppFactory;
use Monolog\Logger;
use Monolog\Level;
use Monolog\Handler\StreamHandler;


require __DIR__ . '/../vendor/autoload.php';

require('dice.php');
//require('otel.php');


$logger = new Logger('dice-server');
$logger->pushHandler(new StreamHandler('php://stdout', Level::Error));

$app = AppFactory::create();

$dice = new Dice();

$handler = function (Request $request, Response $response) use ($logger, $dice) {
    $params = $request->getQueryParams();

    $rolls = 1;
    if(isset($params['rolls'])) {
        $rolls = intval($params['rolls']);
    } 

    $load = $params['load'] ?? '';

    $result = $dice->roll($rolls, $load);
    $response->getBody()->write(json_encode($result));

    return $response;
};

$app->get('/v1/{_}', $handler);
$app->get('/v2/{_}', $handler);

$app->run();