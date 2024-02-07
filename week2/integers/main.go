package main

import (
	"fmt"
	"math"
)

func integerOperators() {
	fmt.Println("Integer Operators")

	var x uint16 = 16

	fmt.Printf("x is %d, %b (binary)\n", x, x)
	fmt.Printf("x plus 3 is %d\n", x+3)
	fmt.Printf("x minus 3 is %d\n", x-3)
	fmt.Printf("x times 3 is %d\n", x*3)
	fmt.Printf("x divided by 3 is %d\n", x/3)
	fmt.Printf("x divided by 3 is %f\n", float64(x)/3)
	fmt.Printf("x mod 3 is %d\n", x%3)
	fmt.Printf("x shifted left by 3 is %d, %b (binary)\n", x<<3, x<<3)
	fmt.Printf("x shifted right by 3 is %d, %b (binary)\n", x>>3, x>>3)
	fmt.Printf("x binary NOT %d, %b (binary)\n", ^x, ^x)
	fmt.Printf("x binary AND 3 is %d, %b (binary)\n", x&3, x&3)
	fmt.Printf("x binary OR 3 is %d, %b (binary)\n", x|3, x|3)
	fmt.Printf("x binary XOR 3 is %d, %b (binary)\n", x^3, x^3)
}

func integerTypeConversion() {
	fmt.Println("Integer Type Conversion")

	var x uint32 = math.MaxUint32
	var y uint64 = math.MaxUint64
	fmt.Printf("x: %d, y: %d\n", x, y)
	x = uint32(y) // y is too big to fit in x, so it's truncated
	y = uint64(x)
	fmt.Printf("x: %d, y: %d\n", x, y)
}

func main() {
	integerOperators()
	fmt.Println()
	integerTypeConversion()
}

