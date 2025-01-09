package main

import "fmt"

func main() {
	age := 12

    if age>18{
        fmt.Println("you are eligible for married");
    }else if age>16{
        fmt.Println("you are eligible for prem");
    }else{
        fmt.Println("you are a kid");
    }
}