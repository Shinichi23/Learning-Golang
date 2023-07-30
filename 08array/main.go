package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", vegList)
	fmt.Println("Vegy list is: ", len(vegList))

	var cars = [...]string{"Ford", "Fiat", "Kia", "Honda", "Renault", "Peugeot", "Audi", "BMW", "Mercedes"}

	fmt.Println("List of my cars brand: ", cars)
	fmt.Println("List of my cars brand: ", len(cars))

}
