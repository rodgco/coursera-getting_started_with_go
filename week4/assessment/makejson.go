package main

import (
	"encoding/json"
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	m := make(map[string]string)

	var name string
	var address string

	in := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	name, err := in.ReadString('\n')
	fmt.Print("Enter address: ")
	address, err = in.ReadString('\n')

	m["name"] = strings.TrimSuffix(name, "\n")
	m["address"] = strings.TrimSuffix(address, "\n")

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
