package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				var names []string
				for _, arg := range stack {
					names = append(names, arg.Name.Local)
				}
				fmt.Printf("%s: %s\n", strings.Join(names, " "), tok)
			}
		}
	}
}

func containsAll(x []xml.StartElement, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		if x[0].Name.Local == y[0] {
			y = y[1:]
		} else {
			for _, atr := range x[0].Attr {
				if atr.Value == y[0] {
					y = y[1:]
					break
				}
			}
		}

		x = x[1:]
	}
	return false
}
