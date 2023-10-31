package main

import (
	"fmt"
	"gobook/ch5/links"
	"io"
	"log"
	"net/http"
	"os"
)

// breadthFirst calls f for each item int the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.

func breadthFirst(f func(string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	links, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return links
}

// NEVER USE THIS! TO DANGEROUS AND IT'S BADLY WRITTEN IT MAKES 2 GET REQUEST FOR EACH PAGE
func modifiedCrawl() func(string) []string {
	const pagesPath = "/home/hq/go/src/gobook/ch5/findlinks3"
	path, err := os.MkdirTemp(pagesPath, "pages*")
	if err != nil {
		log.Fatalf("error creating temporary dir: %v", err)
	}
	folders := make(map[string]string)

	return func(url string) []string {
		fmt.Println(url)
		links, err := links.Extract(url)
		if err != nil {
			log.Print(err)
		}

		//will make another request, it's bad, but ok anyway
		resp, err := http.Get(url)
		if err != nil {
			return nil
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil
		}

		if folders[resp.Request.URL.Host] == "" {
			site, err := os.MkdirTemp(path, resp.Request.URL.Host+"*")
			if err != nil {
				resp.Body.Close()
				return nil
			}
			folders[resp.Request.URL.Host] = site
		}

		file, err := os.CreateTemp(folders[resp.Request.URL.Host], "")
		if err != nil {
			resp.Body.Close()
			return nil
		}

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			resp.Body.Close()
			return nil
		}

		resp.Body.Close()

		return links
	}
}

func main() {
	// Crawl the web, breadth-first
	// Starting from the command line arguments
	breadthFirst(crawl, os.Args[1:])
}
