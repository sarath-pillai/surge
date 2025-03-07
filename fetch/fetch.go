package fetch

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var client = &http.Client{
	Transport: &http.Transport{},
}

func headerSet(h []string, req *http.Request) {
	if len(h) > 0 {
		for _, v := range h {
			hs := strings.SplitN(v, ":", 2)
			if len(hs) == 2 {
				req.Header.Set(strings.TrimSpace(hs[0]), strings.TrimSpace(hs[1]))
			}
		}
	}
}

func authSet(a string, req *http.Request) {
	if len(a) > 0 {
		as := strings.SplitN(a, ":", 2)
		if len(as) == 2 {
			req.SetBasicAuth(strings.TrimSpace(as[0]), strings.TrimSpace(as[1]))
		}
	}
}

func makeRequest(r *http.Request, ch chan<- string) {
	start := time.Now()
	response, err := client.Do(r)
	if err != nil {
		ch <- fmt.Sprintf("Error: %v", err)
		return

	}
	elapsed := time.Since(start).Seconds()
	nbytes, err := io.Copy(io.Discard, response.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error: %v", err)
		return
	}
	defer response.Body.Close()
	ch <- fmt.Sprintf("%.2fs	%7d	%d", elapsed, nbytes, response.StatusCode)
}

func Fetch_post(u string, m string, ct string, b string, h []string, a string, c int) []string {
	var ResultsTotal []string
	if b == "" {
		req, err := http.NewRequest(m, u, nil)
		if err != nil {
			fmt.Sprintf("%v", err)
		}
		req.Header.Set("Content-Type", ct)
		headerSet(h, req)
		authSet(a, req)
		ch := make(chan string, c)
		go func() {
			for range c {
				go makeRequest(req, ch)
			}
			close(ch)
		}()
		for result := range ch {
			ResultsTotal = append(ResultsTotal, result)
		}
	} else {
		data, err := os.Open(b)
		ch := make(chan string, c)
		if err != nil {
			fmt.Sprintf("Error: %v", err)
			log.Fatal(err)
		}
		req, err := http.NewRequest(m, u, data)
		if err != nil {
			fmt.Sprintf("%v", err)
		}
		req.Header.Set("Content-Type", ct)
		headerSet(h, req)
		authSet(a, req)
		go func() {
			for range c {
				makeRequest(req, ch)
			}
			close(ch) // Closing the channel after sending
		}()
		for result := range ch {
			ResultsTotal = append(ResultsTotal, result)
		}
	}
	return ResultsTotal
}

func Fetch_get(u string, m string, h []string, a string, c int) []string {
	var ResultsTotal []string
	req, err := http.NewRequest(m, u, nil)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	headerSet(h, req)
	authSet(a, req)
	ch := make(chan string)
	for range c {
		go makeRequest(req, ch)
	}
	for range c {
		ResultsTotal = append(ResultsTotal, <-ch)
	}
	return ResultsTotal
}
