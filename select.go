package main

import (
	"fmt"
	"time"
)

func emit(wordChannel chan string, done chan bool) {
	defer close(wordChannel)
	words := []string{"a", "b", "c", "d", "e"}
	i := 0
	t := time.NewTimer(3 * time.Second)

	for {
		select {
		case wordChannel <- words[i]:
			i += 1
			if i == len(words) {
				i = 0
			}
		case <-done:
			fmt.Printf("Channel is done\n")
			//close(done) close any channel
			//done <- true
			//return
		case <-t.C:
			return
		}
	}
}

func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emit(wordCh, doneCh)

	for word := range wordCh {
		fmt.Printf("%s ", word)
	}

	// doneCh <- true
	// <-doneCh

}
