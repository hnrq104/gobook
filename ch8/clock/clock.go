// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

var p = flag.Uint64("port", 8000, "specifies which port the server will listen to")

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp",
		net.JoinHostPort("localhost", strconv.FormatUint(*p, 10)))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("succesfully connected to port: %d", *p)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format(time.StampMicro)+"\n")
		if err != nil {
			log.Print(err)
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
