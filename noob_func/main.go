package main

import "fmt"

func main() {
	// variable  expression or function expression

	//a := 10 // expression
	//add(2,3) in local scope you can't invoke func above the func expression
	add := func(a, b int) {
		c := a + b
		fmt.Println(c)
	}
	add(2, 3)
}

func init() {
	fmt.Println("i'm init function")
}
