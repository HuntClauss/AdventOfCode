#/usr/bin/env bash

BIN=/usr/bin

day=$1
$BIN/cp -r ./template "./$day"
$BIN/mkdir -p ./input/$day
$BIN/touch ./input/$day/test1.in ./input/$day/test2.in ./input/$day/final1.in ./input/$day/final2.in