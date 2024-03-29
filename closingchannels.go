package main

import (
	"fmt"
	"time"
)

func printer(msg string, goCh chan bool) {
	<-goCh
	fmt.Printf("%s\n", msg)
}

func main() {
	goCh := make(chan bool)

	for i := 0; i < 10; i++ {
		go printer(fmt.Sprintf("%d", i), goCh)
	}

	time.Sleep(5 * time.Second)
	close(goCh)
	time.Sleep(5 * time.Second)

}
