package main

import (
	"fmt"
)

func showCloth() {
	fmt.Println("1.Zara")
	fmt.Println("2.Mans World")
	fmt.Println("3.Nike \n\n")

}
func electricItem() {
	fmt.Println("1.Mobile")
	fmt.Println("2.Computer")
	fmt.Println("3.Sound Box \n\n")
}
func groceriesItem() {
	fmt.Println("1.Fruits")
	fmt.Println("2.Vegetables")
	fmt.Println("3.Meats\n\n")
}
func clothItem() {
	var option string
	fmt.Println("show")
	fmt.Printf("Enter your option : ")
	fmt.Scanln(&option)
	switch option {
	case "show":
		showCloth()
	}
}

func allItem(itemOption int) {
	fmt.Println("1.cloth")
	fmt.Println("2.electronic")
	fmt.Println("3.groceries")
	fmt.Println("<----------------------->")
	fmt.Println()
	fmt.Printf("Enter your option : ")
	fmt.Scanln(&itemOption)
	switch itemOption {
	case 1:
		clothItem()
	case 2:
		electricItem()
	case 3:
		groceriesItem()
	}
}

func main() {
	for {
		var allOption int
		allItem(allOption)
	}

}
