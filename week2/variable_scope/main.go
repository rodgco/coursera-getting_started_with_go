package main

import "fmt"

var x = 4 // package-level scope

func f() {
	var x = 5 // function-level scope
	fmt.Printf("f: %d\n", x)
}

func g() {
	fmt.Printf("g: %d\n", x)
}

func main() {
	f()
	g()
	h()
}

