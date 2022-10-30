package encodingJson

import (
	"encoding/json"
	"time"
)

type Order struct {
	ID          string    `json:"id"`
	DataOrdered time.Time `json:"data_ordered"`
	CustomerID  string    `json:"customer_id"`
	Items       []Item    `json:"items"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func PersonTo() Person {
	toFile := Person{
		Name: "Fred",
		Age:  40,
	}
	return toFile

}

var o Order

func OrderMarshall() error {

	data := "test order"
	err := json.Unmarshal([]byte(data), &o)
	if err != nil {
		return err
	}
	return nil
}
