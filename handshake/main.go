// server will not accept the client connection, use tcpdump to
// see what happens.
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
	for i := 1; i <= 10; i++ {
		_, err := net.Dial("tcp", addr)
		dieIfError(err)
		fmt.Printf("handshaked: %3d\n", i)

		time.Sleep(time.Second)
	}
}

func server() {
	_, err := net.Listen("tcp", addr)
	dieIfError(err)
	fmt.Printf("listen on %s\n", addr)

	select {}
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
