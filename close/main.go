// client write data to server, then close socket.
// server sleep before read, what data will it read?
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	addr = ":10234"
)

func client() {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		println(err.Error())
	}

	n, err := c.Write([]byte("hi"))
	log.Println(n, err)
	n, err = c.Write([]byte("hi"))
	log.Println(n, err)
	log.Println("close: err", c.Close())
}

func server() {
	l, err := net.Listen("tcp", addr)
	dieIfError(err)
	log.Printf("listen on %s\n", addr)
	for {
		c, err := l.Accept()
		dieIfError(err)

		go func(c net.Conn) {
			b := make([]byte, 1024)
			log.Println("got client, sleep 5s before read")
			time.Sleep(5 * time.Second)

			for {
				n, err := c.Read(b)
				if err != nil {
					log.Printf("%s: %v\n", c.RemoteAddr(), err)
					break
				}

				log.Println(n, string(b))
			}

		}(c)

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
