package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var p = flag.Bool("p", false, "Will run populate: Exercise 5.2 ")
var text = flag.Bool("t", false, "Will run printText: Exercise 5.3")

func main() {
	flag.Parse()
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v", err)
		os.Exit(1)
	}

	if *p {
		count := make(map[string]int)
		populate(count, doc)
		for k, v := range count {
			fmt.Printf("count[%q] = %d\n", k, v)
		}
	} else if *text {
		printTextElements(doc)
	} else {
		outline(nil, doc)

	}
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) //push tag
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

// Exercise 5.2
func populate(pop map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		pop[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		populate(pop, c)
	}
}

// Exercise 5.3
func printTextElements(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return
		}
	}

	if n.Type == html.TextNode {
		fmt.Print(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printTextElements(c)
	}
}
