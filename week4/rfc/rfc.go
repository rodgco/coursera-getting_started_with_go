package main

import (
	"fmt"
	"net/http"
)

func main() {
	page := "https://rodg.co"
	resp, err := http.Get(page)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println(resp.Status)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(resp.Header.Get("Server"))
	fmt.Println(resp.Header.Get("Date"))
	fmt.Println(resp.Header.Get("Content-Length"))
	fmt.Println(resp.Header.Get("ETag"))
}
