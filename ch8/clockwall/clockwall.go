package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	for _, arg := range os.Args[1:] {
		var parts = strings.Split(arg, "=")
		if len(parts) != 2 {
			log.Fatalf("argument in wrong format place=port: %s", arg)
		}
		hostport := net.JoinHostPort("localhost", parts[1])
		conn, err := net.Dial("tcp", hostport)
		if err != nil {
			log.Print(err)
			continue
		}
		defer conn.Close()
		go watch(os.Stdout, conn, parts[0])
	}

	time.Sleep(time.Minute)
}

func watch(w io.Writer, r io.Reader, place string) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Fprintf(w, "%s: %s\n", place, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Printf("cant read from %s: %v", place, err)
	}
}
