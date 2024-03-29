package main

import "fmt"

func main() {
	dayMonths := make(map[string]int)
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 31
	dayMonths["Apr"] = 30
	dayMonths["May"] = 31
	dayMonths["Jun"] = 30
	dayMonths["Jul"] = 31
	dayMonths["Aug"] = 31
	dayMonths["Sep"] = 30
	dayMonths["Oct"] = 31
	dayMonths["Nov"] = 30
	dayMonths["Dec"] = 31

	fmt.Printf("%d\n", dayMonths["Jan"])
	fmt.Printf("%d\n", dayMonths["bat"]) // returns 0

	days, ok := dayMonths["bat"]

	if !ok {
		fmt.Printf("Fucked up\n")
	} else {
		fmt.Printf("%d\n", days)
	}

	delete(dayMonths, "Nov") // delete the key from map
	delete(dayMonths, "bat") // wont' complain safe

	for month, days := range dayMonths {
		fmt.Printf("%s has %d\n", month, days)
	}

}
