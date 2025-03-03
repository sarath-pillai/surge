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
	concurrency := flag.Int("c", 1, "concurrency")
	body := flag.String("b", "", "body")
	method := flag.String("m", "GET", "method")
	contentType := flag.String("ct", "", "contentType")
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
			if *method == "GET" {
				go fetch.Fetch_get(u, ch)
			}
			if *method == "POST" {
				go fetch.Fetch_post(u, ch, *method, *contentType, *body)
			}
		}
	}
	for range urls {
		for range *concurrency {
			fmt.Println(<-ch)
		}
	}
}
