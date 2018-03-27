// How long does it take to get to Mars?
package main

import "fmt"

func main() {
	const lightspeed = 299729 // km/s
	const dayHours = 24       // hours / day
	var distance = 56000000   // km
	var speed = 100800        // km/h

	fmt.Println(distance/lightspeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightspeed, "seconds")

	distance = 96300000 // kms
	fmt.Println((distance/speed)/dayHours, "days")

	// Page 18 Section 2.7

	var days = 28
	distance = 56000000
	speed = distance / (dayHours + days)
	fmt.Printf("To get to Malacandra (56,000,000 km) in 28 days you need to go %v km/h\n", speed)

}
