package main

import (
	"fmt"
)

func printer(msg string) {
	fmt.Printf("%s", msg)
}

// note homogenous type of parameters don't need type decl everytime
func printer1(msg, msg2 string) {
	fmt.Printf("%s, %s", msg, msg2)
}

func printer2(msg string) (string, error) {
	msg += "\n"
	_, err := fmt.Printf(msg)

	return msg, err
}

// can't have multiple dot dot dot and only at the end.
func printer4(msgs ...string) {
	for _, msg := range msgs {
		fmt.Printf("%s\n", msg)
	}
}

// defer is for cleanup of the function.
func printer3(msg string) error {
	defer fmt.Printf("\n")
	defer fmt.Printf("\nMore in queue")
	_, err := fmt.Printf("%s", msg)

	return err
}

func main() {

	printer("New cat\n")
	printer1("Hello world", "Batman")
	msg, err := printer2("BTT")

	if err == nil {
		fmt.Printf("%q\n", msg)
	}

	printer3("Batman Begins")

	printer4("a", "b", "c", "d")
}
