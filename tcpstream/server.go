package main

import (
	"fmt"
	"net"
	"os"

	"github.com/funkygao/golib/color"
)

func main() {
	netListen, err := net.Listen("tcp", ":9988")
	CheckError(err)

	defer netListen.Close()

	Log("listen on :9988")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		buffer = buffer[:]
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		Log(conn.RemoteAddr().String(), n, color.Red(string(buffer[n-1:n])), string(buffer[:n]))
		fmt.Println()
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
