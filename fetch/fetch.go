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

func Fetch_post(u string, ch chan<- string, m string, ct string, b string, h []string, a string) {
	if b == "" {
		req, err := http.NewRequest(m, u, nil)
		if err != nil {
			ch <- fmt.Sprintf("%v", err)
		}
		req.Header.Set("Content-Type", ct)
		headerSet(h, req)
		authSet(a, req)
		start := time.Now()
		response, err := client.Do(req)
		if err != nil {
			ch <- fmt.Sprintf("Error: %v", err)

		}
		elapsed := time.Since(start).Seconds()
		nbytes, err := io.Copy(io.Discard, response.Body)
		if err != nil {
			ch <- fmt.Sprintf("Error: %v", err)
		}
		defer response.Body.Close()
		ch <- fmt.Sprintf("%.2fs	%7d	%s	%d", elapsed, nbytes, u, response.StatusCode)
	} else {
		data, err := os.Open(b)
		if err != nil {
			ch <- fmt.Sprintf("Error: %v", err)
			log.Fatal(err)
		}
		req, err := http.NewRequest(m, u, data)
		if err != nil {
			ch <- fmt.Sprintf("Error: %v", err)
		}
		req.Header.Set("Content-Type", ct)
		headerSet(h, req)
		authSet(a, req)
		start := time.Now()
		response, err := client.Do(req)
		if err != nil {
			ch <- fmt.Sprintf("Error:- %v", err)
		}
		elapsed := time.Since(start).Seconds()
		nbytes, err := io.Copy(io.Discard, response.Body)
		if err != nil {
			ch <- fmt.Sprintf("Error: %v", err)
		}
		defer response.Body.Close()
		ch <- fmt.Sprintf("%.2fs	%7d	%s	%d", elapsed, nbytes, u, response.StatusCode)
	}
}

func Fetch_get(u string, ch chan<- string, m string, h []string, a string) {
	req, err := http.NewRequest(m, u, nil)
	if err != nil {
		ch <- fmt.Sprintf("Error: %v", err)
	}
	headerSet(h, req)
	authSet(a, req)
	start := time.Now()
	response, err := client.Do(req)
	if err != nil {
		ch <- fmt.Sprintf("Error: %v", err)
	}
	elapsed := time.Since(start).Seconds()
	nbytes, err := io.Copy(io.Discard, response.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error: %v", err)
	}
	defer response.Body.Close()
	ch <- fmt.Sprintf("%.2fs	%7d		%s	%d", elapsed, nbytes, u, response.StatusCode)
}
