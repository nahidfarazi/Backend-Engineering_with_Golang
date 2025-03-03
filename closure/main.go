package main

import "fmt"

const a = 100

var p = 10

func outer() func() {
	money := 100
	age := 21
	fmt.Println("Age =", age)
	show := func() {
		money = money + a + p
		fmt.Println(money)
	}
	return show
}
func call() {
	incr1 := outer()
	incr1()
	incr1()
	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	call()
}
func init() {
	fmt.Println("== Bank ==")
}

// 1.code segment -> a, outer,outer.anonymous,call, main,init
// 2, data segment -> p
// 3.stack memory
// 4. heap (gc) ->  money show = adress
