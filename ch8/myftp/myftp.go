package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", net.JoinHostPort(os.Args[1], os.Args[2]))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	go watch(os.Stdout, conn)
	io.Copy(conn, os.Stdin)
}

func watch(w io.Writer, r io.Reader) {
	_, err := io.Copy(w, r)
	if err != nil {
		log.Print(err)
	}
}
