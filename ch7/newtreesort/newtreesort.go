package newtreesort

import (
	"fmt"
	"strings"
)

type Tree struct {
	value       int
	left, right *Tree
}

//Exercise 7.3

func (t *Tree) String() string {
	if t == nil {
		return ""
	}
	return t.left.String() + " " + fmt.Sprint(t.value) + " " + t.right.String()
}

func (t *Tree) ToString() string {

	depth := 0
	lines := make([]string, 0)

	var TreeLine func(*Tree)
	TreeLine = func(t *Tree) {

		lines = append(lines, fmt.Sprintf("%*s%v", 2*depth, "", t.value))
		depth++
		if t.left != nil {
			TreeLine(t.left)
		}
		if t.right != nil {
			TreeLine(t.right)
		}
		depth--
	}
	TreeLine(t)

	return strings.Join(lines, "\n")
}

// Sort sorts values in place.

func Sort(values []int) {
	var root *Tree
	for _, v := range values {
		root = Add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the element of t in order
// and returns the resulting slice

func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func Add(t *Tree, value int) *Tree {
	if t == nil {
		t = new(Tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = Add(t.left, value)
	} else {
		t.right = Add(t.right, value)
	}
	return t
}
