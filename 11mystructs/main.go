package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in Golang")
	// There's no inheritance in golang; No super or parent

	shinichi := User{"John", "mikejohn@go.dev", true, 32}
	fmt.Println(shinichi)
	fmt.Printf("Shinichi details are: %+v\n", shinichi)
	fmt.Printf("Name is %v and email is %v.\n", shinichi.Name, shinichi.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
