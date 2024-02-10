package main

import "fmt"

func main() {
	var arr [5]int

	arr[0] = 100
	fmt.Printf("%d\n", arr[0])

	arr = [...]int{1, 2, 3, 4, 5} // Array literal

	for i, v := range arr {
		fmt.Println(i, v)
		arr[i] += 100
	}
	fmt.Printf("%v\n", arr)
}
