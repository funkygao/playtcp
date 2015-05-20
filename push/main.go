// client write data to server, but server will not read.
// finally, client write buffer got full, and use netstat to see write Q.
// use tcpdump to see if PUSH get ACK.
package main

import (
	"fmt"
	"net"
	"os"
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
		i++
		fmt.Printf("to write: %6d\n", i)
		n, err := c.Write([]byte("hello world!"))
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
		_, err := l.Accept()
		dieIfError(err)

		// do nothing about this conn
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
