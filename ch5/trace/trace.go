package main

import (
	"log"
	"time"
)

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	// lots of work
	time.Sleep(5 * time.Second)
}

func double(x int) (result int) {
	g := func() { log.Printf("double(%d) = %d", x, result) }
	defer g()
	return x + x
}

func main() {
	// bigSlowOperation()
	double(5)
}
