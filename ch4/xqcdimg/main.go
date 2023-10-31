package main

import (
	"flag"
	"fmt"
	"gobook/ch4/xqcd"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const imgPath = "/home/hq/pics/xqcd"

var path = flag.String("path", imgPath, "specifies where image should be downloaded")

func main() {
	flag.Parse()
	log.Printf("creating/using directory at %s", *path)
	err := os.MkdirAll(*path, os.ModeDir)
	if err != nil {
		log.Fatalf("error making directory: %v", err)
	}
	for _, arg := range flag.Args() {
		number, err := strconv.Atoi(arg)
		if err != nil {
			log.Printf("error parsing int from %s\n", arg)
			continue
		}

		comic, err := xqcd.SearchXQCD(number)
		if err != nil {
			log.Printf("error searching for %d: %v", number, err)
			continue
		}

		imgUrl := comic.Img
		resp, err := http.Get(imgUrl)
		if err != nil {
			log.Printf("error getting %s: %v", imgUrl, err)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			log.Printf("unwanted status %s: %s", imgUrl, resp.Status)
			continue
		}

		filename := fmt.Sprintf("%s/%d", *path, number)
		file, err := os.Create(filename)
		if err != nil {
			resp.Body.Close()
			log.Printf("error: %v", err)
			continue
		}

		_, err = io.Copy(file, resp.Body)
		resp.Body.Close()
		file.Close()
		if err != nil {
			log.Printf("error copying: %v", err)
			continue
		}

		log.Printf("xqcd comic %d downloaded succesfully", number)
	}
}
