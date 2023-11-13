package main

import (
	"log"
	"os"
	"sync"

	"gopl.io/ch8/thumbnail"
)

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Print(err)
		}
	}
}

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // this doesn't wait for them to finish !bad, it won't do anythin
	}
}

// makeThumbnails3 makes thumbnails of the specified files in parallel
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(s string) {
			_, err := thumbnail.ImageFile(s)
			if err != nil {
				log.Print(err)
			}
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<-ch
	}
}

// makeThumbnails4 makes thumbnails of the specified files in parallel
// and returns their errors
func makeThumbnails4(filenames []string) error {
	ch := make(chan error)
	for _, f := range filenames {
		go func(s string) {
			_, err := thumbnail.ImageFile(s)
			ch <- err
		}(f)
	}
	for range filenames {
		if err := <-ch; err != nil {
			return err // incorrect goroutine leak
		}
	}

	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(s string) {
			tfile, err := thumbnail.ImageFile(s)
			ch <- item{tfile, err}
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

// This is the most important example
// makeThumbnails6 makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1)
		//add worker
		go func(s string) {
			defer wg.Done()
			tfile, err := thumbnail.ImageFile(s)
			if err != nil {
				log.Print(err)
				return
			}

			info, _ := os.Stat(tfile)
			sizes <- info.Size()
		}(f)
	}

	//closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	// If closer came before this loop, it would wait forever, since
	// each worker will not call wg.Done() until sizes has transmitted
	// info.Size()
	var total int64
	for size := range sizes {
		total += size
	}

	// if closer came after this loop, the loop would never end,
	// as sizes would never close, so it would be waiting for the next
	// integer

	// see pic in 239

	return total

}
