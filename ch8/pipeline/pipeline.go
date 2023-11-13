package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	// go func() {
	// 	for x := 0; x < 100; x++ {
	// 		naturals <- x
	// 	}
	// 	close(naturals)
	// }()

	// // Squarer
	// go func() {
	// 	for x := range naturals {
	// 		squares <- x * x
	// 	}
	// 	close(squares)
	// }()

	// // Printer in main go routine
	// for x := range squares {
	// 	fmt.Println(x)
	// }

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)

}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func mirroredQuery() string {
	responses := make(chan string, 3) //THIS HAS TO BE A BUFFERED CHANNEL
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	return <-responses
}

// If it were an unbuffered channel, the two slowest goroutines would
// be stuck trying to send into responses, but they wouldnt be able to, as responses
// would be deallocated to the garbage collector after the return of mirroredQuery

func request(s string) string {
	return s //
}
