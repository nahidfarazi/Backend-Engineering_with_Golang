package main

import "fmt"

// function input value call parameter (a,b int)
func add(a, b int) {
	c := a + b
	fmt.Println(c)
}
func main() {
	add(2, 3) //function pass value call argument (2,3)
}
