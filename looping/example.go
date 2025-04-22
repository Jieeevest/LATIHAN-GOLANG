package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	names := []string{"John", "Paul", "George", "Ringo"}
	for i, name := range names {
		fmt.Printf("%d: %s\n", i, name)
	}

	for _, name := range names {
		fmt.Printf("%s\n", name)
	}

	person := map[string]string{"name": "John", "address": "London"}
	for key, value := range person {
		fmt.Printf("%s: %s\n", key, value)
	}
}