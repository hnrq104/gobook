package main

import (
	"fmt"
	"gobook/ch5/links"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int

	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(s string) {
					worklist <- crawl2(s)
				}(link)
			}
		}
	}

}
