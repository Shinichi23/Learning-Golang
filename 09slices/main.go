package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to my slices course")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 954
	highScores[2] = 456
	highScores[3] = 845
	//highScores[4] = 999 doesn't work because we declared 4 integers

	highScores = append(highScores, 555, 666, 321)
	// append work and change thee things and make it work

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	fmt.Println(highScores)

	//how to remove a value from slices base on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}
