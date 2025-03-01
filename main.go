package main

import (
	"fmt"
	"os"
	"surge/fetch"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("At least provide one Host/URL to fetch")
		os.Exit(1)
	}
	ch := make(chan string)
	for _, url := range args {
		go fetch.Fetch(url, ch)
	}
	for range args {
		fmt.Println(<-ch)
	}
}
