package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("Is this a digit? %t\n", unicode.IsDigit('1'))
}
