package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"surge/fetch"
	"surge/reports"
)

type headers []string

func (h *headers) String() string {
	return fmt.Sprintf("%s", *h)
}

func (h *headers) Set(value string) error {
	*h = append(*h, value)
	return nil
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	url := flag.String("u", "", "url")
	concurrency := flag.Int("c", 1, "concurrency")
	body := flag.String("b", "", "HTTP body")
	method := flag.String("m", "GET", "HTTP method")
	contentType := flag.String("ct", "", "contentType")
	authentication := flag.String("a", "", "Basic authentication in the format username:password")
	var h headers
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of Surge:\n surge -u 'https://www.google.com,https://www.example.com'\n")
		flag.PrintDefaults()
	}
	flag.Var(&h, "header", "Header Name and Value")
	flag.Parse()
	if !isFlagPassed("u") {
		fmt.Println("At least One URL needs to be passed. Use -u to pass one")
		os.Exit(1)
	}
	urls := strings.Split(*url, ",")
	ch := make(chan string)
	fmt.Printf("Running Performance test against %v with concurrency of %d\n", *url, *concurrency)
	for _, u := range urls {
		for range *concurrency {
			if *method == "GET" {
				go fetch.Fetch_get(u, ch, *method, h, *authentication)
			}
			if *method == "POST" {
				go fetch.Fetch_post(u, ch, *method, *contentType, *body, h, *authentication)
			}
		}
	}
	var results []string
	for range urls {
		for range *concurrency {
			results = append(results, <-ch)
		}
	}
	min, max, not_ok_status := reports.Stats(results)
	fmt.Printf("Lowest Response Time: %.2fs\nHighest Response Time: %.2fs\nResponses Other than 200: %d\n", min, max, not_ok_status)
	for _, r := range results {
		fmt.Println(r)
	}
}
