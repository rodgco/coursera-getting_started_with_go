package main

import (
	"fmt"
	"math"
)

func main() {
	var number float32

	fmt.Print("Enter a floating point number: ")
	_, err := fmt.Scan(&number)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Truncated number:", math.Trunc(float64(number)))
	}
}
