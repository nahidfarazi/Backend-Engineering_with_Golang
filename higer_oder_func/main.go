package main

import (
	. "fmt"
)

/* in go have a to type function
1.first oder function
2.higher oder function / first class function
*/
/*
	 higher oder function must be
		any of the following 3
		1.parameter in function -> func name (func name)
		2.return function -> return func name
		3.both
*/

// If any func receive func in a parameter that func call by -> call bank func
func processOp(a, b int, op func(p, q int)) func(x, y int) {
	op(a, b)
	// add(10,20)
	return add
}

/*
	 func call() func(x, y int) {
		return add
	}
*/
func add(x, y int) {
	z := x + y
	Println(z)
}
func main() {
	sum := processOp(2, 3, add)
	sum(3, 4)
}
