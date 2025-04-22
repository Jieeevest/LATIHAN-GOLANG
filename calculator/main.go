package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("\n=== Calculator ===")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			Addition()
		case 2:
			Subtraction()
		case 3:
			Multiplication()
		case 4:
			Division()
		case 5:
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func Addition() {
	var num1, num2 int
	fmt.Println("\n=== Addition ===")
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	fmt.Println("Result:", num1+num2)
}

func Subtraction() {
	var num1, num2 int
	fmt.Println("\n=== Subtraction ===")
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	fmt.Println("Result:", num1-num2)
}

func Multiplication() {
	var num1, num2 int
	fmt.Println("\n=== Multiplication ===")
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	fmt.Println("Result:", num1*num2)
}

func Division() {
	var num1, num2 int
	fmt.Println("\n=== Division ===")
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	if num2 == 0 {
		fmt.Println("Error: Cannot divide by zero")
		return
	}
	fmt.Println("Result:", num1/num2)
}
