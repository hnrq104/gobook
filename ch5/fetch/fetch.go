package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// changed as to use defer for f.close()
// exercise 5.18
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	defer resp.Body.Close()

	local := path.Base(url)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)

	if err != nil {
		return "", 0, err
	}

	closeThings := func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}

	defer closeThings()

	n, err = io.Copy(f, resp.Body)

	return local, n, err
}

func main() {
	for _, arg := range os.Args[1:] {
		name, nbytes, err := fetch(arg)
		if err != nil {
			fmt.Printf("fetch(%s): %v\n", arg, err)
		} else {
			fmt.Printf("file %s created, %d bytes copied.\n", name, nbytes)
		}
	}
}
