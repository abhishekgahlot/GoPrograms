package main

import (
	"fmt"
)

func main() {
	for counter := 0; counter < 10; counter++ {
		//fmt.Printf("Hello, World\n")
	}

	// complex for loop
	for i, j := 0, 1; i < 10; i, j = i+1, j*2 {
		fmt.Printf("Hello, World %d\n", j)
	}

	var stop bool
	i := 0
	for !stop {
		fmt.Printf("Hello, World\n")
		i++
		if i == 10 {
			stop = true
		}
	}
}
