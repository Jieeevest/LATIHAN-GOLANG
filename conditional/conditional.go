package main

import (
		"fmt"
		"math/rand"
	)

func checkNumber(number int) {
	if number > 18 {
		fmt.Println("number is greater than 18, value: ", number)
	} else {
		fmt.Println("number is less than 18, value: ", number)
	}
}

func main() {
	//versi 1
	number := rand.Intn(10)
	checkNumber(number)

	//versi 2
	if number := rand.Intn(10); number > 5 {
		fmt.Println("number is greater than 5, value: ", number)
	} else {
		fmt.Println("number is less than 5, value: ", number)
	}

	//versi 3
	if nilai := rand.Intn(100); nilai > 75 {
		fmt.Println("Lulus")
	} else if nilai > 60 {
		fmt.Println("Remedial")
	} else {
		fmt.Println("Tidak Lulus")
	}
}