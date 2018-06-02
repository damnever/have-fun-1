package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	addr   = flag.String("addr", "server:8021", "the server address")
	rbufsz = flag.Int("rbufsz", 0, "the tcp read buffer size")
)

func main() {
	flag.Parse()
	var n int64 = 5 * (1 << 20)

	// https://github.com/golang/go/issues/25701
	// Dialer.Control
	s, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	must(err)
	must(syscall.SetsockoptInt(s, syscall.IPPROTO_TCP, syscall.TCP_NODELAY, 1))
	if *rbufsz > 0 {
		must(syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_RCVBUF, *rbufsz))
	}
	must(syscall.Connect(s, parseAddr()))
	f := os.NewFile(uintptr(s), "tcp")
	conn, err := net.FileConn(f)
	must(err)
	f.Close()
	defer conn.Close()

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

func parseAddr() *syscall.SockaddrInet4 {
	parts := strings.Split(*addr, ":")
	port, err := strconv.Atoi(parts[1])
	must(err)
	var ip [4]byte
	copy(ip[:], net.ParseIP(parts[0]).To4())
	return &syscall.SockaddrInet4{
		Port: port,
		Addr: ip,
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
