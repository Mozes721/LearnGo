package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func failedUpdate(px *int) {
	x2 := 20
	px = &x2
	fmt.Println(px)
}

func update(px *int) {
	*px = 20
	fmt.Println(px)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// x := 10
	// failedUpdate(&x)
	// fmt.Println(x)
	// update(&x)
	// fmt.Println(x)
	var p Person

	err := json.Unmarshal([]byte(`{"name": "Bob", "age":30}`), &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

}
