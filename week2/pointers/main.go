package main

import "fmt"


	
func p1() {
	var x int = 1
	var y int
	var ip *int

	ip = &x
	y = *ip

	x = 2

	fmt.Printf("x: %d\ny: %d\n", x, y)

	*ip = 3
	fmt.Printf("Changed *ip...\nx: %d\ny: %d\n", x, y)
}

func p2() {
	ptr := new(int)
	*ptr = 3

	fmt.Printf("ptr: %d\naddr: %d\n", *ptr, ptr)
}

func main() {
	p1()
	p2()
}

