package main

import (
	"fmt"
)


func simple(a, b int) (int,int){
	var sum int

	sum = a+b
	var g int
	g = 5
	return sum ,g
}
func main() {

	var ferot, sum int
	ferot,_= simple(2,3)
	
	c := 5
	sum = ferot + c
	fmt.Println(sum);
	
}