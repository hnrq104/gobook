package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

/*	forEachNode calls the functions pre(x) and post(x) for each node
x in the tree rooted at n. Both functions are optional.
pre is called before the children are visited (preorder) and
post is called after (postorder)
*/

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if pre(n) {
			return true
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if forEachNode(c, pre, post) {
			return true
		}
	}
	if post != nil {
		return post(n)
	}

	return false
}

var depth int

func startElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		if len(n.Attr) != 0 {
			fmt.Printf("%*s<%s ", depth*2, "", n.Data)
			for _, at := range n.Attr {
				fmt.Printf("%s=%q ", at.Key, at.Val)
			}
			fmt.Printf(">\n")
		} else {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		}
		depth++
	}
	return false
}

func endElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	if n.Type == html.TextNode {
		for _, s := range strings.Split(n.Data, "\n") {
			fmt.Printf("%*s %s\n", depth*2, "", s)
		}
	}

	if n.Type == html.CommentNode {
		fmt.Printf("<!--%s-->\n", html.UnescapeString(n.Data))
	}
	return false
}

func indentedTags(url string) error {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code: %s", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		resp.Body.Close()
		return fmt.Errorf("parse failed: %v", err)
	}

	forEachNode(doc, startElement, endElement)
	return nil
}

// Exercise 5.8
func ElementById(doc *html.Node, s string) *html.Node { //how would I do it
	if doc.Type == html.ElementNode {
		for _, atr := range doc.Attr {
			if atr.Key == "id" && atr.Val == s {
				return doc
			}
		}
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		if r := ElementById(c, s); r != nil {
			return r
		}
	}
	return nil
}

// following the books question, I think I would need a global variable, as var depth int

func bookElemById(doc *html.Node, s string) *html.Node {
	var found *html.Node
	var idLookUp string
	var findIdNode func(*html.Node) bool

	findIdNode = func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, atr := range n.Attr {
				if atr.Key == "id" && atr.Val == idLookUp {
					found = n
					return true
				}
			}
		}
		return false
	}

	idLookUp = s
	forEachNode(doc, findIdNode, nil)
	return found
}

func main() {
	for _, args := range os.Args[1:] {
		fmt.Printf("reaching url: %s\n", args)
		err := indentedTags(args)
		if err != nil {
			log.Printf("error capturing %s: %v", args, err)
		}
	}
}

func outline2(url string) error {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code: %s", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		resp.Body.Close()
		return fmt.Errorf("parse failed: %v", err)
	}

	depth := 0

	var startElm func(*html.Node) bool
	var endElm func(*html.Node) bool

	startElm = func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s <%s>\n", depth*2, "", n.Data)
			depth++
		}
		return true
	}

	endElm = func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s </%s>\n", depth*2, "", n.Data)
		}
		return true
	}

	forEachNode(doc, startElm, endElm)
	return nil
}
