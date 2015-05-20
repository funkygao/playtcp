package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const (
	addr = "localhost:1234"
)

func client() {
	c, err := net.Dial("tcp", addr)
	dieIfError(err)
	c.(*net.TCPConn).SetNoDelay(false) // enable nagles

	i := 0
	byteWritten := 0
	for {
		i++
		fmt.Printf("to write: %6d\n", i)
		n, err := c.Write([]byte(strings.Repeat("X", 10)))
		dieIfError(err)
		byteWritten += n
		fmt.Printf("written: %6d, totoal: %d\n", i, byteWritten)

		time.Sleep(time.Second)
	}

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
	b := make([]byte, 1<<10)
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
