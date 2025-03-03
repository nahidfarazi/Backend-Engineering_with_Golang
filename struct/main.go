package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) userModify(age int) {
	u.Age = age
}

func main() {
	var user1 User
	user1 = User{
		Name: "Rafiqul",
		Age:  22,
	}
	user1.userModify(30)
	fmt.Println("Name: ", user1.Name)
	fmt.Println("Age: ", user1.Age)

	user2 := User{
		Name: "Saiful",
		Age:  21,
	}
	fmt.Println("Name: ", user2.Name)
	fmt.Println("Age: ", user2.Age)

}
