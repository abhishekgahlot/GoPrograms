package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("maps.go")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	defer f.Close()

	b := make([]byte, 500)
	n, err := f.Read(b)

	//stringVersion := string(b)

	fmt.Printf("%d: %s\n", n, b)
}
