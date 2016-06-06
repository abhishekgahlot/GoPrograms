package main

import (
	"fmt"
)

// when you pass an array it copies it.
func printer(w [3]string) {
	for _, i := range w {
		fmt.Printf("%s", i)
	}
	fmt.Printf("\n")
	w[2] = "blue"
}

//pass the slice they are as referenced not copied so faster, performant
func printer1(w []string) {
	for _, words := range w {
		fmt.Printf("%s\n", words)
	}
}

func main() {
	// arrays
	// words := [3]string{"the", "batman", "ratman"}
	// fmt.Printf("%s\n", words[1])
	// for _, i := range words {
	// 	fmt.Printf("%s\n", i)
	// }

	// fmt.Printf("%d\n", len(words))

	// printer(words)
	// printer(words)

	new_words := []string{"the", "batman", "ratman"}
	// pass as slice is passed as reference
	printer1(new_words)
	printer1(new_words)
	// some examples of slice is [:3], [2:4] ( 4 not included )

	//words := make([]string, 4) // can hold 4 items or ([]string, 0, 4)
	words := make([]string, 0, 4)
	words = append(words, "The")
	words = append(words, "bat")
	words = append(words, "has")
	words = append(words, "woken")
	words = append(words, "up")
	words = append(words, "again")
	fmt.Printf("%d\n", cap(words)) // cap -> capacity of array;

	newWords := make([]string, 4)
	copy(newWords, words) // copy second param into first param of slice
	printer1(newWords)
	printer1(words)
}
