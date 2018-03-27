package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	year := 2016
	switch month := rand.Intn(12) + 1; month {
	case 2:
		day := rand.Intn(29) + 1
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
	}

	// 	Short declaration is not available for variables declared at the package scope, so era cannot
	//  be declared with era := "AD" at its current location.

	month := rand.Intn(12) + 1
	daysInMonth := 31
	switch month {
	case 2:
		daysInMonth = 29
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day)
}
