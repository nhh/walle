#!/usr/bin/env bash

bombardier -c 15 -n 10000 http://localhost:8082/api-two > .benchmark
