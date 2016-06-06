package main

import (
	"fmt"
	"os"
)

func main() {
	if numChars, theError := fmt.Printf("Hello, world\n"); theError != nil {
		os.Exit(1)
	} else {
		fmt.Printf("%d", numChars)
	}
}
