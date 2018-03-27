package main

import (
	"fmt"
)

func printHeader() {
	fmt.Printf("%-15v %7v %12v %5v\n", "Spaceline", "Days", "Round-trip", "Price")
	fmt.Printf("=" * 72)
}

func main() {
	printHeader()
}
