package main
import (
	"fmt"
	"rsc.io/quote/v4"
)

func main() {
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())

	var c, python, java bool
	fmt.Println(c, python, java)

	var a, b int = 1, 2
	fmt.Println(a, b)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

