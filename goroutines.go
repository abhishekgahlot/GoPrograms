package main

import (
	"fmt"
)

func emit(c chan string) {
	words := []string{"a", "b", "c", "d"}
	for _, i := range words {
		c <- i
	}
	close(c)
}

func main() {
	wordChannel := make(chan string)

	go emit(wordChannel)

	word := <-wordChannel
	// transmit first word
	fmt.Printf("%s\n", word)

	word = <-wordChannel
	// transmit second word
	fmt.Printf("%s\n", word)

	word = <-wordChannel
	// transmit third word
	fmt.Printf("%s\n", word)

	word = <-wordChannel
	// transmit fourth word
	fmt.Printf("%s\n", word)

	word, ok := <-wordChannel
	// transmit fifth word which is empty cause channel is closed.
	fmt.Printf("%s %t\n", word, ok)

	fmt.Printf("nothing above right?\n")
	// for word := range wordChannel {
	// 	fmt.Printf("%s ", word)
	// }

	fmt.Printf("\n")

	// tranmit two now

	go emit(wordChannel)
	go emit(wordChannel)

	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}
}
