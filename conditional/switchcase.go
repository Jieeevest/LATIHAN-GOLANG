package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Intn(4)
	fmt.Println(a)

	switch a {
		case 0:
			fmt.Println("a is 0")
		case 1:
			fmt.Println("a is 1")
		case 2:
			fmt.Println("a is 2")
		default:
			fmt.Println("a is not 0, 1 or 2")
	}
}