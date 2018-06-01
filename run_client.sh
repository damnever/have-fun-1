#!/bin/bash

run () {
    PROG=$1
    ADDR=$2
    RBUFSZ=$3

    echo "-> run $PROG $ADDR $RBUFSZ"
    case $PROG in
        "go")
            ./client -addr="$ADDR" -rbufsz="$RBUFSZ" ;;
        "py")
            python client.py --addr "$ADDR" --rbufsz "$RBUFSZ" ;;
    esac
}

# normal
run go server:8020 0  # go <-> go
run go server:8021 0  # go <-> py
run py server:8020 0  # py <-> go
run py server:8021 0  # py <-> py
# set SO_RCVBUF
run go server:8020 4194304  # go <-> go
run go server:8021 4194304  # go <-> py
run py server:8020 4194304  # py <-> go
run py server:8021 4194304  # py <-> py
