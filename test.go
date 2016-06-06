package main

import (
	"fmt"
)

func worker(bat chan int) {
	some := []int{1, 2, 4, 5, 6, 7, 8, 9}
	for _, i := range some {
		bat <- i
	}
	defer close(bat)
}

func main() {
	somechan := make(chan int)
	go worker(somechan)

	for i := range somechan {
		fmt.Printf("%d", i)
	}

}
