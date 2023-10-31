package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Printf("oi")
}
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return err
		}
		log.Printf("server not responding(%v); retrying ...", err)
		time.Sleep(time.Second << tries)
	}

	return fmt.Errorf("could not contact server")

}
