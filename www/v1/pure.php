<?php

/*
$arr = [];

for ($i = 0; $i < 1_000_000 ; $i++) {
    $arr[$i] = 1_000_000 - $i;
}

//sort($arr);


*/



$vector = new \Ds\Vector();
$vector->allocate(1_000_000);

for ($i = 0; $i < 1_000_000 ; $i++) {
    $vector->insert($i, 1_000_000 - $i);
}


//$vector->sort();
echo rand(1,6);