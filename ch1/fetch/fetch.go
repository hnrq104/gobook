//fetch prints the content found at URL.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	//exercise 1.8 add prefix if there isn't one
	for _, arg := range os.Args[1:] {

		url := arg
		if !strings.HasPrefix(arg, "http://") {
			url = "http://" + url
		}
	
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
		fmt.Printf("HTTP status code : %s\n", resp.Status)

	}
}
