package main

import "fmt"

// add is -> standard or named function
func add(a, b int) {
	fmt.Println(a + b)
}
func main() {
	//below function is an anonymous function
	//also it's an immediately invoke function expression or IIFE
	func(a, b int) {
		c := a + b
		fmt.Println(c)
	}(6, 4)
	add(2, 3)
}
func init() {
	fmt.Println("i am init func")
}

/*


 */
