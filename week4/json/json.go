package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name string
	age  int
	address string
}

func main() {
	p := Person{"John", 30, "New York"}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Person: %v\n", string(b))
}

