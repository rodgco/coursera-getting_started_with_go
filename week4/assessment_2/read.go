package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var filename string
	var names []Name

	fmt.Print("Enter the name of the file: ")
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	names = make([]Name, 0, 10)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		name := strings.SplitN(line, " ", 2)
		names = append(names, Name{fname: name[0], lname: name[1]})
	}

	for _, name := range names {
		fmt.Println(name.fname, name.lname)
	}	
}
