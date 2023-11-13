package main

import (
	"flag"
	"fmt"
	"gobook/ch5/links"
	"log"
)

var semaphore = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	semaphore <- struct{}{}
	list, err := links.Extract(url)
	<-semaphore
	if err != nil {
		log.Print(err)
	}
	return list
}

type LinksAndDepth struct {
	links []string
	d     int
}

var depth = flag.Int("depth", 0, "sets maximum depth")

func main() {
	flag.Parse()
	worklist := make(chan LinksAndDepth)
	var n int
	n++
	seen := make(map[string]bool)
	go func() { worklist <- LinksAndDepth{flag.Args(), 0} }()
	for ; n > 0; n-- {
		pages := <-worklist

		if pages.d > *depth {
			n--
			continue
		}
		// d less or equal to depth

		for _, url := range pages.links {
			if seen[url] {
				continue
			}
			//url not seen
			seen[url] = true
			n++
			go func(s string, dep int) {
				worklist <- LinksAndDepth{crawl(s), dep}
			}(url, pages.d+1)

		}

	}

}
