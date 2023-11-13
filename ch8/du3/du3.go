package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "verbose")

func main() {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)

	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSizes)
	}

	//waiter
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	//ticker
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1fGB\n", nfiles, float64(nbytes)/1e9)
}

var semaphore = make(chan struct{}, 20)

func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()

	semaphore <- struct{}{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		<-semaphore
		return
	}
	<-semaphore

	for _, entry := range entries {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, wg, fileSizes)
		} else {
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
				return
			}
			fileSizes <- info.Size()
		}
	}

}
