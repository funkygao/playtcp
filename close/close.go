// find the 4-way close, handled by os or application
// run this with 'tcpdump -i lo0 -nnN port 10234'
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
	time.Sleep(10 * time.Second)
	log.Println("client start dialing")
	c, err := net.Dial("tcp", addr)
	if err != nil {
		println(err.Error())
	}

	time.Sleep(5 * time.Second)
	log.Println("client start writing")
	c.Write([]byte("hi"))
	log.Println("client closing")
	c.Close()
	log.Println("client left...")
}

func server() {
	log.Println("start listening")
	l, err := net.Listen("tcp", addr)
	dieIfError(err)
	log.Printf("server listen on %s\n", addr)
	for {
		c, err := l.Accept()
		dieIfError(err)

		go func(c net.Conn) {
			b := make([]byte, 1024)
			log.Println("got client, sleep 8s before read")
			time.Sleep(8 * time.Second)

			for {
				n, err := c.Read(b)
				if err != nil {
					log.Printf("server %s: %v\n", c.RemoteAddr(), err)
					break
				}

				log.Println("server recv: ", n, string(b))
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
