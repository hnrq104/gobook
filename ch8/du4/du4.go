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

var done = make(chan struct{})

// como nada é lançado no canal, somente após ele ser fechado
// será possivel recuperar valores (no caso valores 0)
// portanto esse select só vai selecionar o primeiro caso após ele fechar
func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

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

	//closer
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
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
		case <-done:
			for range fileSizes {
				//Do nothing
			}
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

func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()

	if cancelled() {
		return
	}

	for _, entry := range dirents(dir) {
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

var semaphore = make(chan struct{}, 20)

// dirents return the entries of directory dir
func dirents(dir string) []os.DirEntry {
	select {
	case semaphore <- struct{}{}: // acquire token
	case <-done:
		return nil
	}

	defer func() { <-semaphore }()

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries

}
