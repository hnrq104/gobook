package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		go handleConn(conn)

	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	inp := bufio.NewScanner(c)
	for inp.Scan() {
		_, err := io.WriteString(c, string(inp.Bytes())+"\n")
		if err != nil {
			log.Print(err)
			return // client disconnected
		}
	}
}
