#!/bin/bash

set -ex

./server -port=8020 &
GO=$!
python server.py --port 8021 &
PY=$!

wait $GO
wait $PY
