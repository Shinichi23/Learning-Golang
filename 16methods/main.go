package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in Golang")
	// There's no inheritance in golang; No super or parent

	shinichi := User{"John", "mikejohn@go.dev", true, 32}
	fmt.Println(shinichi)
	fmt.Printf("Shinichi details are: %+v\n", shinichi)
	fmt.Printf("Name is %v and email is %v.\n", shinichi.Name, shinichi.Email)

	shinichi.GetStatus()
	shinichi.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
