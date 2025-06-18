#!/bin/bash

go build -o race.bin -race

chmod u+x race.bin

./race.bin