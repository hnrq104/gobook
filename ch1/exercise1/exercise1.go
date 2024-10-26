package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Program started at: ", start)
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
	finished := time.Now()
	fmt.Println("Program finished at: ", finished)
	fmt.Println("It took :", finished.Nanosecond()-start.Nanosecond(), "nanoseconds")
}
