# -*- coding: utf-8 -*-

import socket
import argparse


parser = argparse.ArgumentParser()
parser.add_argument('--port', type=int, default=8021, help='the server port')
args = parser.parse_args()


s = socket.socket(socket.AF_INET, socket.SOCK_STREAM, socket.IPPROTO_TCP)
s.setsockopt(socket.IPPROTO_TCP, socket.TCP_NODELAY, 1)
s.bind(('0.0.0.0', args.port))
s.listen(5)

n = 5 * (1<<20)  # noqa
data = b'O' * n

while 1:
    conn, _ = s.accept()
    while 1:  # only one client allowed
        try:
            conn.recv(2)
            i = 0
            while i < n:
                i += conn.send(data[i:])
        except socket.error:
            conn.close()
            break

s.close()
