// write 0 length bytes to socket, use tcpdump to see if it is sent
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	addr = "localhost:1234"
)

func client() {
	c, err := net.Dial("tcp", addr)
	dieIfError(err)

	i := 0
	byteWritten := 0
	for {
		time.Sleep(time.Second * 5)

		i++
		n, err := c.Write([]byte("")) // zero length
		dieIfError(err)
		byteWritten += n
		fmt.Printf("written: %6d, totoal: %d\n", i, byteWritten)
	}

}

func server() {
	l, err := net.Listen("tcp", addr)
	dieIfError(err)
	fmt.Printf("listen on %s\n", addr)

	for {
		conn, err := l.Accept()
		dieIfError(err)

		// do nothing about this conn
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	b := make([]byte, 1<<10)
	for {
		n, err := conn.Read(b)
		dieIfError(err)
		fmt.Printf("%d, %+v", n, b[:n])
	}

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
