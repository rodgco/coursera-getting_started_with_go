package main

import "fmt"

func main() {
	var apples int

	fmt.Print("How many apples do you have? ")

	num, err := fmt.Scan(&apples)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("You have %d apples\n", apples)
	fmt.Printf("Number of items successfully scanned: %d\n", num)
}
