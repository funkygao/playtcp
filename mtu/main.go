// client keeps connect to server, but server will not accept.
// it is used to check the backlog feature.
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	addr        = "localhost:1234"
	loopbackMtu = 16384
	loopbackMss = loopbackMtu - 40
)

func client() {
	c, err := net.Dial("tcp", addr)
	dieIfError(err)

	data := strings.Repeat("X", loopbackMss+1)
	fmt.Printf("data len: %d\n", len(data))
	n, err := c.Write([]byte(data))
	dieIfError(err)
	fmt.Printf("written: %d\n", n)
}

func server() {
	l, err := net.Listen("tcp", addr)
	dieIfError(err)
	fmt.Printf("listen on %s\n", addr)

	for {
		conn, err := l.Accept()
		dieIfError(err)

		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	b := make([]byte, loopbackMtu*10)
	n, err := conn.Read(b)
	dieIfError(err)
	fmt.Printf("read: %d\n", n)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	switch os.Args[1] {
	case "server":
		server()

	case "client":
		client()

	default:
		usage()
	}
}

func usage() {
	fmt.Printf("Usage: %s <server|client>\n", os.Args[0])
	os.Exit(1)
}

func dieIfError(err error) {
	if err != nil {
		panic(err)
	}
}
