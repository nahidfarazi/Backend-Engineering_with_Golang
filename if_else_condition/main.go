package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Printf("enter your input: ")
	fmt.Scanf("%d", &age)

	if age > 18 && age < 120 {
		fmt.Println("you are eligible for married")
	} else if age > 16 && age < 120 && age < 18 {
		fmt.Println("you are eligible for prem")
	} else if age == 120 || age > 120 {
		fmt.Println("yor are alien")
	} else {
		fmt.Println("you are a kid")
	}

}
