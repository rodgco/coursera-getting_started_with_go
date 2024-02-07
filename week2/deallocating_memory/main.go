package main

import "fmt"

var x = 10

func f(n int) {
	fmt.Printf("&n = %p\n\n", &n)
}

func main() {
	fmt.Printf("&x = %p\n", &x)
	f(x)

	var y = 20
	fmt.Printf("&y = %p\n", &y)
	f(y)
}
