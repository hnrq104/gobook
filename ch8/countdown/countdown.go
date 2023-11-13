package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // reads a single byte
		abort <- struct{}{}
	}()
	// ch := time.Tick(1 * time.Second)

	// for countdown := 10; countdown >= 0; countdown-- {
	// 	fmt.Println(countdown)
	// 	<-ch
	// }

	// select {
	// case <-time.After(5 * time.Second):
	// 	fmt.Println("Launch started!")
	// case <-abort:
	// 	fmt.Println("Mission aborted!")
	// }
	//launch

	fmt.Println("Commencing countdown, press return to abort mission!")
	tick := time.NewTicker(time.Second)

	for countdown := 5; countdown >= 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick.C:

		case <-abort:
			fmt.Println("MISSION ABORTED!")
			break
		}
	}
	tick.Stop()
	fmt.Println("Good luck on your trip commander!")

}
