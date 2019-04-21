#!/usr/bin/env bash

bombardier -c 250 -n 100000 http://localhost:8082/api-two > .benchmark
