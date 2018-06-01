package main

import (
	"flag"
	"net"
)

var (
	port = flag.String("port", "8020", "the listen port")
)

func main() {
	flag.Parse()
	n := 5 * (1 << 20)
	data := make([]byte, n)
	for i := 0; i < n; i++ {
		data[i] = 'O'
	}

	l, err := net.Listen("tcp", ":"+*port)
	must(err)
	defer l.Close()

	b2 := make([]byte, 2)
	for {
		conn, err := l.Accept()
		must(err)

		go func(conn net.Conn) {
			defer conn.Close()
			// bufio??
			for {
				if _, err := conn.Read(b2); err != nil {
					return
				}
				if _, err = conn.Write(data); err != nil {
					return
				}
			}
		}(conn)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
