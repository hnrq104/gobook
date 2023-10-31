// Findlinks prints the links in an HTML read from standart input
package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var r = flag.Bool("r", false, "Use the full recursive visit2")

func main() {
	flag.Parse()
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v", err)
		os.Exit(1)
	}

	var list []string
	if *r {
		list = visit2(nil, doc)
	} else {
		list = visit(nil, doc)
	}

	for _, link := range list {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns results
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

// Exercise 5.2
func visit2(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	links = visit2(links, n.FirstChild)
	links = visit2(links, n.NextSibling)

	return links
}
