package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"
	"strings"
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

//this is boring really,

func handleConn(c net.Conn) {
	defer c.Close()
	const homePath = "/home/hq/"
	currentPath := homePath
	inp := bufio.NewScanner(c)
	for inp.Scan() {
		var resp string

		switch s := inp.Text(); s {
		case "ls":
			files, err := os.ReadDir(currentPath)
			if err != nil {
				resp = fmt.Sprint(err)
				break
			}

			var fileInfo []string
			for _, f := range files {
				fileInfo = append(fileInfo, f.Name())
			}

			resp = strings.Join(fileInfo, "\n")
			resp += "\n"

		case "cd":
			info, err := os.Stat(path.Join(currentPath, s))
			if err != nil {
				resp = fmt.Sprint(err)
				break
			}
			if !info.IsDir() {
				resp = fmt.Sprintf("%s is not a directory\n", s)
				break
			}
			currentPath = path.Join(currentPath, s)
		}

		_, err := io.WriteString(c, resp)
		if err != nil {
			log.Print(err)
			return // client disconnected
		}
	}
}
