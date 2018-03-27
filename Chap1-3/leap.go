package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(" Leap Year OR condition ")
	fmt.Println("The year is 2100, should you leap?")
	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	if leap {
		fmt.Println("Look before you leap!")
	} else {
		fmt.Println("Keep your feet on the ground.")
	}
	fmt.Println(" Torch Not OR condition")
	var haveTorch = true
	var litTorch = false
	if !haveTorch || !litTorch {
		fmt.Println("Nothing to see here.")
	}

	fmt.Println(" The Switch Statement")
	fmt.Println("There is a cavern entrance here and a path to the east.")
	var command = "go inside"
	switch command {
	case "go east":
		fmt.Println("You head further up the mountain.")
	case "enter cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that.")
	}
	fmt.Println(" Countdown Loop")
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Liftoff!")
}
