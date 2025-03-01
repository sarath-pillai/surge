package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"surge/fetch"
)

func main() {
	url := flag.String("u", "", "url")
	concurrency := flag.Int("c", 100, "concurrency")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of Surge:\n surge -u 'https://www.google.com,https://www.example.com'\n")
	}
	flag.Parse()
	urls := strings.Split(*url, ",")
	if len(urls) < 1 {
		fmt.Println("At least provide one Host/URL to fetch")
		os.Exit(1)
	}
	ch := make(chan string)
	for _, u := range urls {
		for range *concurrency {
			go fetch.Fetch(u, ch)
		}
	}
	for range urls {
		for range *concurrency {
			fmt.Println(<-ch)
		}
	}
}
