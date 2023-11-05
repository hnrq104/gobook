// Does the fetching itself
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var c = flag.Bool("c", false, "Exercise 5.5, will count words and images")

func main() {
	flag.Parse()

	if *c {
		for _, arg := range flag.Args() {
			words, images, err := CountWordsAndImages(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Count: %v\n", err)
			}
			fmt.Printf("url:%s words:%d images:%d\n", arg, words, images)
		}
	} else {
		for _, arg := range flag.Args() {
			links, err := findLinks(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			}

			for _, link := range links {
				fmt.Println(link)
			}
		}
	}
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(url string) ([]string, error) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return visit(nil, doc), nil

}

// visit appends to links each link found in n and returns results
// FROM go/src/gobook/ch5/findlinks1
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		if n.Data == "a" || n.Data == "link" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		if n.Data == "script" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

// Exercise 5.5
// CountWordsAndImages does a HTTP GET request
// for the HTML document url and returns the number of
// words and images
func CountWordsAndImages(url string) (words, images int, err error) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("getting %s:%s", resp.Status, url)
		resp.Body.Close()
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing html %s: %v", url, err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
		}
	} else if n.Type == html.TextNode {
		reader := strings.NewReader(n.Data)
		buf := bufio.NewScanner(reader)
		buf.Split(bufio.ScanWords)
		for buf.Scan() {
			words++
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return
}
