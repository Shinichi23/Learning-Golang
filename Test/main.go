package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	r, err := http.Get("https://jsonplaceholder.typicode.com/todos?_limit=5")

	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()
	var todos []Todo
	json.NewDecoder(r.Body).Decode(&todos)

	//fmt.Println(todos)
	fmt.Printf("%#v", todos)

}
