package main

import (
	"fmt"
	"sort"
)

// prereqs maps computer science courses to their prerequesites
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languagues",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"network":               {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i, course)
	}
	// mapped version
	fmt.Printf("mapped verion:\n")
	for i, course := range mappedTopoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i, course)
	}

	fmt.Printf("extended mapped verion:\n")
	s, ok := extendedTopoSort(prereqs)
	if ok {
		for i, course := range s {
			fmt.Printf("%d:\t%s\n", i, course)
		}
	}
	fmt.Printf("%t\n", ok)

}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitALL func([]string)
	visitALL = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitALL(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitALL(keys)
	return order
}

// 5.10
func mappedTopoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visit func(string)
	visit = func(n string) {
		if seen[n] {
			return
		}
		seen[n] = true
		for _, key := range m[n] {
			visit(key)
		}
		order = append(order, n)
	}

	for key := range m {
		visit(key)
	}

	return order
}

// reports cycles 5.11
func extendedTopoSort(m map[string][]string) ([]string, bool) {
	var order []string
	seen := make(map[string]bool)
	trees := make(map[string]int)
	tree := 1
	var visit func(string) bool
	visit = func(n string) bool {
		for _, item := range m[n] {
			if !seen[item] {
				seen[item] = true
				trees[item] = tree
				if !visit(item) {
					return false
				}
			} else if trees[item] == trees[n] {
				return false
			}
		}

		order = append(order, n)
		return true
	}

	for key := range m {
		if seen[key] {
			continue
		}

		seen[key] = true
		trees[key] = tree
		if !visit(key) {
			return nil, false
		}
		tree++
	}
	return order, true
}
