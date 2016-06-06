package main

import (
	"fmt"
	"math/rand"
)

func makeRandoms(c chan int) {
	var i int
	i = 0
	for {
		c <- rand.Intn(1000)
		i += 1
		if i > 1000 {
			break
		}
	}
	close(c)
}

func uniqueId(c chan int) {
	var id int
	id = 0
	counter := 0
	for {
		c <- id
		id += 1
		counter += 1
		if counter > 1000 {
			break
		}
	}
	close(c)
}

func main() {
	// randoms := make(chan int)

	// go makeRandoms(randoms)

	// for n := range randoms {
	// 	fmt.Printf("%d ", n)
	// }

	idChan := make(chan int)
	go uniqueId(idChan)

	for i := range idChan {
		fmt.Printf("%d\n", i)
	}
}
