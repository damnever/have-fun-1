package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"time"
)

var (
	addr   = flag.String("addr", "server:8021", "the server address")
	rbufsz = flag.Int("rbufsz", 0, "the tcp read buffer size")
)

func main() {
	flag.Parse()
	var n int64 = 5 * (1 << 20)

	// bufio??
	conn, err := net.Dial("tcp", *addr)
	must(err)
	defer conn.Close()
	if *rbufsz > 0 {
		tcpconn, _ := conn.(*net.TCPConn)
		must(tcpconn.SetReadBuffer(*rbufsz))
	}

	var total time.Duration
	start := time.Now()
	for i := 0; i < 10; i++ {
		_, err = conn.Write([]byte("go"))
		must(err)
		_, err = io.CopyN(ioutil.Discard, conn, n)
		must(err)

		end := time.Now()
		cost := end.Sub(start)
		fmt.Printf(" %v\n", cost)
		start = end
		total += cost
	}

	fmt.Printf("[AVG] %v\n", total/time.Duration(10))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
