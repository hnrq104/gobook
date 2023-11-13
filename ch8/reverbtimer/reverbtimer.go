package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(w io.Writer, shout string, t time.Duration) {
	fmt.Fprintf(w, "\t%s\n", strings.ToUpper(shout))
	time.Sleep(t)
	fmt.Fprintf(w, "\t%s\n", shout)
	time.Sleep(t)
	fmt.Fprintf(w, "\t%s\n", strings.ToLower(shout))
}

func handleConn(c *net.TCPConn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup

	timer := time.NewTimer(5 * time.Second)

	//Listen to timer
	go func() {
		<-timer.C
		c.CloseRead()
	}()

	if input.Scan() {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			echo(c, s, 1*time.Second)
		}(input.Text())
		timer.Reset(5 * time.Second)
	}
	timer.Stop()
	c.CloseRead()
	wg.Wait()
	c.CloseWrite()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handleConn(c.(*net.TCPConn))
	}
}
