package main

import "fmt"

func callBack(a, b int, op func(x, y int)) {
	op(a, b)
}
func add(x, y int) {
	z := x + y
	fmt.Println("i from add func: ", z)
}
func main() {
	callBack(12, 3, add)
}
