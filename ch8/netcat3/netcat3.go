package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	tcp := conn.(*net.TCPConn)

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, tcp) //ignoring errors
		log.Println("done")
		tcp.CloseRead()
		done <- struct{}{}
	}()

	mustCopy(tcp, os.Stdin)
	tcp.CloseWrite()
	<-done // wait for background goroutine to end
}

func mustCopy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}
