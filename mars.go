// My weight loss program.

package main

import "fmt"

// Main Program is where go starts

func main() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(230.4 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(63 * 365 / 687)
	fmt.Print(" years old.\n")
	// Using Printf gives better control over output
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 230.4*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 63*365/687)
	// Using Println puts a return on the line
	fmt.Println("If used Print I would need a return character. With Println I don't")
	// Multiple format verbs can be used in the with Print f
	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 154.0)
	// Printf can also help aligning text
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

}
