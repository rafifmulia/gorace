#!/bin/bash

bombardier -m GET -n 100 -c 4 -l \
  --fasthttp \
  'http://127.0.0.1:8080/race?k=v' > bombardier.out