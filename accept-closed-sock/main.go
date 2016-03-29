package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	addr = "10.1.114.159:10234"
)

func client() {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		println(err.Error())
	}
	log.Printf("client handshaked, sleeping 2s...")

	time.Sleep(time.Second * 2)
	if err = c.Close(); err != nil {
		log.Printf("client: %s", err)
	}
	log.Println("client closed connection")
}

func server() {
	s, err := net.Listen("tcp", addr)
	dieIfError(err)
	log.Printf("server listen on %s\n", addr)

	for {
		log.Println("server sleep 25s")
		time.Sleep(time.Second * 25)
		c, err := s.Accept()
		log.Printf("server accept a new connection %s", c.RemoteAddr().String())
		if err != nil {
			log.Println(err)
		}

		n, err := c.Write([]byte("hi"))
		log.Println("written", n, err)
		n, err = c.Write([]byte("hi"))
		log.Println("written", n, err)
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
