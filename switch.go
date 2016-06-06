package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("Hello, World!")
	// switch doesn't have break it has fallthrough
	switch {
	case err != nil:
		os.Exit(1)
	case n != 14: // n != 0
		fmt.Printf("Wrong number of chars %d", n)
		fallthrough
	default:
		fmt.Printf("everything good")
	}

	fmt.Printf("\n")

	atoz := "the quick brown fox jumps over the lazy dog"
	vowels := 0
	consonants := 0
	zeds := 0

	for _, r := range atoz {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		case 'z':
			zeds+=2
			fallthrough
		default:
			consonants++
		}
	}

	fmt.Printf("Vowels %d, consonants %d zeds %d", vowels, consonants, zeds)

}
