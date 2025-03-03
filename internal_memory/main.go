package main

import . "fmt"

/*
when run a go program create in memory 4 segment

	1.Code segment -> all type func store here and constant variable
	2.Data segment -> all global variable store here but not constant variable
	3.Stack -> when run a function create a stack frame
	4.heap ->
*/

/*
go code run in two phases

	1.compile phases
	3.execution phases
*/
var a = 10

const p = 100

func call() {
	add := func(x, y int) {
		z := x + y
		Println(z)
	}
	add(3, 4)
	add(a, p)
}

func main() {
	call()
}
func init() {
	Println("hello i'm init func")
}
