package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var nearDistance = 56000000
	var farDistance = 401000000
	var diffDistance = farDistance - nearDistance
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(diffDistance) + nearDistance

	fmt.Println(num)

}
