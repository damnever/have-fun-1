#!/bin/bash

PROG=$1
ADDR=$2
RBUFSZ=$3

case $PROG in
    "go")
        ./client -addr="$ADDR" -rbufsz="$RBUFSZ" ;;
    "py")
        python client.py --addr "$ADDR" --rbufsz "$RBUFSZ" ;;
esac
