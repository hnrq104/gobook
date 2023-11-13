package main

import (
	"fmt"
	"gobook/ch5/links"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string, 20)
	unSeenLinks := make(chan string)

	for i := 0; i < 20; i++ {
		go func() {
			for url := range unSeenLinks {
				list := crawl(url)
				go func() { worklist <- list }()
			}
		}()
	}

	go func() { worklist <- os.Args[1:] }()
	seen := make(map[string]bool)
	//Main Go Routine checks if items were seen
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unSeenLinks <- link
			}
		}
	}

}
