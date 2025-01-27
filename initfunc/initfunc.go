package main

import "fmt"
var a=10
func main() {
	fmt.Println("Main : ",a)
}
func init(){
	fmt.Println("init : ",a)
	a = 20
}