package main

import (
	"fmt"
	"math/rand"
)

type shuffler interface {
	Len() int
	Swap(i, j int)
}

func shuffle(s shuffler) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i)
		s.Swap(i, j)
	}
}

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}

func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

type stringSlice []string

func (s stringSlice) Len() int {
	return len(s)
}

func (s stringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	is := intSlice{1, 2, 3, 4, 5, 6, 7, 8}
	shuffle(is)
	fmt.Printf("%v\n", is)

	s := stringSlice{"aat", "bat", "cat"}
	shuffle(s)
	fmt.Printf("%v", s)
}
