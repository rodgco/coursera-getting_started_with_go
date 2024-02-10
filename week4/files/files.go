package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Read file
	data, err := ioutil.ReadFile("text.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

