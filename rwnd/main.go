// use tcpdump the checkout server's win changing
// server will not read data, so the rwnd will becoming less and less
// till zero
package main

import (
	"fmt"
	"net"
	"os"
	//"time"
)

const (
	addr = "localhost:1234"
)

func client() {
	c, err := net.Dial("tcp", addr)
	dieIfError(err)
	i := 0
	for {
		i++
		n, err := c.Write([]byte("hello"))
		dieIfError(err)
		fmt.Printf("%d sent: %3d\n", i, n)

		//time.Sleep(time.Millisecond * 10)
	}

}

func server() {
	l, err := net.Listen("tcp", addr)
	dieIfError(err)
	fmt.Printf("listen on %s\n", addr)

	for {
		_, err := l.Accept()
		dieIfError(err)
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
