package main

//Exercise 8.6

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "verbose")

type rootSize struct {
	nfiles, nbytes int64
	name           string
}

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//Declare a channel for each root
	fileSizes := make([]chan int64, len(roots))
	for i := range fileSizes {
		fileSizes[i] = make(chan int64)
	}

	//Declare root structs
	rs := make([]rootSize, len(roots))
	for i := range rs {
		rs[i].name = roots[i]
	}

	//Declare waitGroup for each root and go walkDir
	var wgs []sync.WaitGroup = make([]sync.WaitGroup, len(roots))
	for i, root := range roots {
		wgs[i].Add(1)
		go walkDir(root, &wgs[i], fileSizes[i])
	}

	//Capture sizes and nfiles for each root dir

	var waitAll sync.WaitGroup
	for i := range roots {

		waitAll.Add(1)
		go func(r int) {
			wgs[r].Wait()
			close(fileSizes[r])
		}(i)

		go func(r int) {
			defer waitAll.Done()
			for size := range fileSizes[r] {
				rs[r].nbytes += size
				rs[r].nfiles++
			}
		}(i)
	}

	//Declare ticker to print progress
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	//Print everything per tick
	if tick != nil {
		go func() {
			for {
				<-tick
				printDiskUsage(rs)
			}
		}()

	}

	waitAll.Wait()
	printDiskUsage(rs)
}

func printDiskUsage(rs []rootSize) {
	for _, r := range rs {
		fmt.Printf("root:%s\tfiles:%d\t%.1fGB\n", r.name, r.nfiles, float64(r.nbytes)/1e9)
	}
}

var sema = make(chan struct{}, 20)

func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()

	sema <- struct{}{} //This will keep at maximum 20 goroutines running

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3: %v\n", err)
		<-sema
		return
	}

	<-sema // this won't allow more than 20 file read at the same time

	for _, entry := range entries {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, wg, fileSizes)
		} else {
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "du3: %v\n", err)
			} else {
				fileSizes <- info.Size()
			}
		}
	}
}
