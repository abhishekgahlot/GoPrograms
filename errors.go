package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	errorEmptyString = errors.New("Batman Begins Error")
)

func printer(msg string) error {
	if msg == "" {
		panic(errorEmptyString) //will print out stack trace 
		return errorEmptyString
		return fmt.Errorf("Can't print empty string")
	}
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func main() {
	if err := printer(""); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}
