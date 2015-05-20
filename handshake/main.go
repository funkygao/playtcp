// client keeps connect to server, but server will not accept.
// it is used to check the backlog feature.
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
	i := 0
	for {
		_, err := net.Dial("tcp", addr)
		dieIfError(err)
		i++
		fmt.Printf("handshaked: %3d\n", i)
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
