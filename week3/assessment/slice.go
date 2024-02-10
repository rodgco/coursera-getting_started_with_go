package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var numbers []int
	numbers = make([]int, 0, 3)

	for {
		var input string 
		fmt.Print("Enter an integer (X to quit): ")
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			if input == "X" {
				break
			}
			fmt.Printf("Invalid integer: %v\n", input)
			continue
		}
		numbers = append(numbers, number)
		sort.Ints(numbers)
		fmt.Printf("Numbers: %v\n", numbers)
	}
}
