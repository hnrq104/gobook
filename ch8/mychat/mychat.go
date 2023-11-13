package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	adr  string
	name string
	ch   chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}
		case newcomer := <-entering:
			log.Printf("%s @%s has arrived", newcomer.name, newcomer.adr)
			clients[newcomer] = true
			newcomer.ch <- "Welcome " + newcomer.name + "!\n"
			for cli := range clients {
				newcomer.ch <- cli.name
			}
		case leaver := <-leaving:
			log.Printf("%s @%s has left", leaver.name, leaver.adr)
			delete(clients, leaver)
			for cli := range clients {
				leaver.ch <- cli.name
			}
			close(leaver.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string, 10)
	go clientWriter(conn, ch)

	adr := conn.RemoteAddr().String()

	var ticker = time.NewTimer(5 * time.Minute)

	go func() {
		<-ticker.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	ch <- "Insert your desired username below:"
	input.Scan()
	name := input.Text()

	cli := client{adr, name, ch}

	entering <- cli
	messages <- adr + " " + name + " has arrived!"

	for input.Scan() {
		messages <- name + ": " + input.Text()
		ticker.Reset(5 * time.Minute)
	} //ignoring erros

	leaving <- cli
	messages <- adr + " " + name + " has left :("
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
