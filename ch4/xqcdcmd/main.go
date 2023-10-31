/*Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a request to
https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of
many favorites. Download each URL (once!) and build an offline index. Write a tool xkcd
that, using this index, prints the URL and transcript of each comic that matches a search term
provided on the command line. */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gobook/ch4/xqcd"
	"os"
)

const IndexPath = "/home/hq/go/src/gobook/ch4/xqcd/index.json"

var n = flag.Int("download", 0, "if set, will download \"n\" items from xqcd")

func main() {
	flag.Parse()

	if *n != 0 {
		file, err := os.Create(IndexPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "os.Create(%s): %v", IndexPath, err)
			os.Exit(1)
		}

		if ind, err := xqcd.GetIndex(*n); err != nil {
			fmt.Fprintf(os.Stderr, "xqcd.GetIndex(%d): %v", *n, err)
			os.Exit(1)
		} else if err := xqcd.WriteIndex(file, ind); err != nil {
			fmt.Fprintf(os.Stderr, "xqcd.WriteIndex: %v", err)
			os.Exit(1)
		}
	}

	file, err := os.Open(IndexPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Open(%s): %v", IndexPath, err)
		os.Exit(1)
	}

	var index xqcd.XQCDIndex
	if err := json.NewDecoder(file).Decode(&index); err != nil {
		fmt.Fprintf(os.Stderr, "json.Decode(%s): %v", IndexPath, err)
		os.Exit(1)
	}

	fmt.Printf("Welcome to xqcd searcher, you have %d comics downloaded!\n", index.TotalCount)

	for _, searchterm := range flag.Args() {
		related := xqcd.SearchInIndex(searchterm, &index)
		if len(related) != 0 {
			fmt.Printf("I have found the following comics matching %s!\n", searchterm)
			for _, comic := range related {
				fmt.Printf("Title: %s\n", comic.Title)
				fmt.Printf("Transcript:\n%s\n", comic.Transcript)
				fmt.Printf("Alternative Transcript:\n%s\n", comic.Alt)
				fmt.Printf("URL:%s\n\n", comic.URL)

			}

		} else {
			fmt.Printf("I have found no matching %s :(\n\n", searchterm)
		}

	}

	//

}
