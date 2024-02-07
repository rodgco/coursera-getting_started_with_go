package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func findian(s string) string {
	s = strings.ToLower(s)

	if strings.HasPrefix(s, "i") && strings.Contains(s, "a") && strings.HasSuffix(s, "n") {
		return "Found!"
	} else {
		return "Not Found!"
	}
}

func test() {
	texts := []string{"ian", "Ian", "iuiygaygn", "I d skd a efju N", "ihhhhhn", "ina", "xian"}

	for _, s := range texts {
		fmt.Printf("%s: %s\n", s, findian(s))
	}
}

func main() {
	//test()

	var input string
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a string: ")
	input, err:= in.ReadString('\n')

	input = strings.TrimSuffix(input, "\n")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(findian(input))
}
