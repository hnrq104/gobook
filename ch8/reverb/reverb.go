package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	//Exercise 8.4
	var wg sync.WaitGroup
	tcp := c.(*net.TCPConn)

	for input.Scan() {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			echo(tcp, text, 1*time.Second)

		}(input.Text())
	}
	tcp.CloseRead()

	//note ignoring potential erros from input.Err
	wg.Wait()
	tcp.CloseWrite()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
