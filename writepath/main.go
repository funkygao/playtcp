// start server, then client, within 10s, kill server
// you will see client will panic in 3rd packet instead of 2nd.
// that tells us socket.write() just put to send buffer instead of
// waiting for remote ack.
package main

import (
	"fmt"
    "time"
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
	for {
		i++
		fmt.Printf("sending: %3d\n", i)
        n, err := c.Write([]byte("hello"))
        dieIfError(err)
		fmt.Printf("sent: %3d\n", n)

        time.Sleep(time.Second*10)
	}

}

func server() {
	l, err := net.Listen("tcp", addr)
	dieIfError(err)
	fmt.Printf("listen on %s\n", addr)

    for {
        c, err := l.Accept()
        dieIfError(err)

        b := make([]byte, 100)
        c.Read(b)
        println(string(b))
    }

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
