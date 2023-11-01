//This is a educative example, you should NOT panic for expected errors
//but provides a compact illustration of possibilities

package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			//no panic
		case bailout{}:
			//expected panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) //unexpected panic, continue on panicking
		}
	}()

	//bail out of recursion if we find more than one non-empty title
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple titles
			}
			title = n.FirstChild.Data
		}
	}, nil)
	return title, nil
}

// Exercise 5.19

func main() {
	defer func() {
		p := recover()
		fmt.Println(p)
	}()
	PanickedInt(10)
}

func PanickedInt(x int) {
	panic(x)
}
