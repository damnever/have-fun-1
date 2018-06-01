# -*- coding: utf-8 -*-

import time
import socket
import argparse


parser = argparse.ArgumentParser()
parser.add_argument('--addr', type=str, default='server:8020', help='the server port')  # noqa
parser.add_argument('--rbufsz', type=int, default=0, help='the read buffer size')  # noqa
args = parser.parse_args()


c = socket.socket(socket.AF_INET, socket.SOCK_STREAM, socket.IPPROTO_TCP)
c.setsockopt(socket.IPPROTO_TCP, socket.TCP_NODELAY, 1)
if args.rbufsz > 0:
    c.setsockopt(socket.SOL_SOCKET, socket.SO_RCVBUF, args.rbufsz)

ip, port = args.addr.split(':')
c.connect((ip, int(port)))


n = 5 * (1<<20)  # noqa
total = 0
start = time.time()

for _ in range(10):
    c.send(b'py')
    i = 0
    while i < n:
        i += len(c.recv(4096))

    end = time.time()
    cost = end - start
    print(' {} s'.format(cost))
    start = end
    total += cost

print('[AVG] {} s'.format(total/10))
c.close()
