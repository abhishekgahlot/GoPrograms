package main

import (
	"fmt"
	"math/rand"
	"time"
)

func work() {
	fmt.Printf("[")
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Printf("]")
}

func worker(sema chan bool) {
	<-sema
	work()
	sema <- true
}

func main() {
	sema := make(chan bool, 10)

	for i := 0; i < 1000; i++ {
		go worker(sema)
	}

	for i := 0; i < 10; i++ {
		sema <- true
	}

	time.Sleep(30 * time.Second)
}
