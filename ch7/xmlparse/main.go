package main

import (
	"encoding/xml"
	"fmt"
	"io"
)

type Node interface{} //Char or *Element
type CharData string
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

//Using a constant is faster and better than doing a type check probably

func parse(r io.Reader) (Node, error) {
	dec := xml.NewDecoder(r)

	var root Element
	stack := []*Element{&root}

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("parse: %v", err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elem := &Element{tok.Name, tok.Attr, nil}
			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, elem)
			stack = append(stack, elem)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			cd := CharData(tok)
			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, cd)
		}
	}
	if len(root.Children) == 0 {
		return nil, fmt.Errorf("parse: empty xml")
	}
	return &root, nil
}
