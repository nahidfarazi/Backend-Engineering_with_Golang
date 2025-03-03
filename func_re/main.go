package main

import "fmt"

func ReturnFunc(x, y, z int) int {

	p := x * y * z
	return p
}
func ArgumentFunc(a, b int) {
	z := a + b
	fmt.Println(z)
}

func main() {
	res := ReturnFunc(2, 3, 4)
	ArgumentFunc(12, 3)
	fmt.Println(res)

}
