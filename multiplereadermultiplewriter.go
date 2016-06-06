package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func countPrimes(limit int) int {
  // Return if less than 1
  if limit <= 1 {
    return 0
  }

  // Get the sqrt of the limit
  sqrtLimit := int(math.Sqrt(float64(limit)))

  // Create array
  numbers := make([]bool, limit)

  // Set 1 to prime
  numbers[0] = true
  numPrimes := 0

  // Count the number of olds
  if limit%2 == 0 {
    numPrimes = limit / 2
  } else {
    numPrimes = (limit + 1) / 2
  }

  // Loop through odd numbers
  for i := 3; i <= sqrtLimit; i += 2 {
    if !numbers[i] {
      for j := i*i; j < limit; j += i*2 {
              if !numbers[j] {
          numbers[j] = true
          numPrimes -= 1
              }
      }
    }
  }

  return numPrimes
}

func getPage(url int) (int, error) {
	resp, err := numPrimes(int)
	if err != nil {
		return 0, err
	}

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return 0, err
	}

	return len(body), nil
}

func getter(url string, size chan string) {
	length, err := getPage(url)
	if err == nil {
		size <- fmt.Sprintf("%s has %d length", url, length)
	}
}

func main() {
	urls := []int{10000, 20000, 30000, 40000, 50000}

	//pageLen, err := getPage(url)
	size := make(chan string)

	for _, url := range urls {
		go getter(url, size)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-size)
	}
}
