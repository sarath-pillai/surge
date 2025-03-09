package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"surge/fetch"
	"surge/reports"
	"surge/timer"
	"time"
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
	url := flag.String("u", "", "The endpoint IP, URL against which the test needs to be performed")
	concurrency := flag.Int("c", 1, "How many Concurrent requests should be sent")
	body := flag.String("b", "", "HTTP body. This is a file containing the data that needs to be sent")
	method := flag.String("m", "GET", "HTTP method. Currently GET & POST is supported")
	duration := flag.Int("d", 5, "Duration for how long the test should run. Default is 5 seconds")
	contentType := flag.String("ct", "", "contentType")
	consoleOut := flag.Bool("o", false, "Should output be displayed on the console")
	authentication := flag.String("a", "", "Basic authentication in the format username:password")
	reportGeneration := flag.Bool("r", false, "Should CSV report be generated")
	reportFileName := flag.String("f", "", "CSV File name(if not provided it will create a surge-timestamped file)")
	var h headers
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of Surge:\n surge -u 'https://www.example.com'\n")
		flag.PrintDefaults()
	}
	flag.Var(&h, "header", "Header Name and Value")
	flag.Parse()
	if !isFlagPassed("u") {
		fmt.Println("At least One URL needs to be passed. Use -u to pass one")
		os.Exit(1)
	}
	var results []string
	urls := strings.Split(*url, ",")
	fmt.Printf("Running Performance test against %v with concurrency of %d, for a duration of %d seconds\n", *url, *concurrency, *duration)
	for _, u := range urls {
		if *method == "GET" {
			results = timer.ExecuteForDuration(func() []string {
				return fetch.Fetch_get(u, *method, h, *authentication, *concurrency)
			}, time.Duration(*duration)*time.Second)
		}
		if *method == "POST" {
			results = timer.ExecuteForDuration(func() []string {
				return fetch.Fetch_post(u, *method, *contentType, *body, h, *authentication, *concurrency)
			}, time.Duration(*duration)*time.Second)
		}
	}
	min, max, not_ok_status := reports.Stats(results)
	fmt.Printf("Lowest Response Time: %.2fs\nHighest Response Time: %.2fs\nResponses Other than 200: %d\n", min, max, not_ok_status)
	if *consoleOut {
		for _, r := range results {
			fmt.Println(r)
		}
	}
	if *reportGeneration {
		reports.GenerateCSV(*reportFileName, results)
	}

}
