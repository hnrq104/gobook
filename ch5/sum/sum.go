package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func sum(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func max(first int, values ...int) int {
	num := first
	for _, val := range values {
		if num < val {
			num = val
		}
	}
	return num
}

func min(first int, values ...int) int {
	num := first
	for _, val := range values {
		if num > val {
			num = val
		}
	}
	return num
}

func variadicStringJoin(sep string, values ...string) string {
	start := ""

	for i, s := range values {
		if i != 0 {
			start += sep
		}

		start += s

		if i != len(values)-1 {
			start += sep
		}
	}
	return start

}

func forEachNode(doc *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(doc)
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(doc)
	}

}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	important := make(map[string]bool)
	for _, val := range name {
		important[val] = true
	}
	var elements []*html.Node
	FindTag := func(doc *html.Node) {
		if doc.Type == html.ElementNode {
			if important[doc.Data] {
				elements = append(elements, doc)
			}
		}
	}
	forEachNode(doc, FindTag, nil)
	return elements
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(3, 4, 5, 6, 1, 2))

	var values = []int{1, 2, 3, 4}
	fmt.Println(sum(values...))

	var f func(...int)
	var g func([]int)

	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)

	fmt.Println(max(1, 5, -6, 200000))
	fmt.Println(variadicStringJoin(" "))
	fmt.Println(variadicStringJoin(" ", "henrique", "naju"))

}
