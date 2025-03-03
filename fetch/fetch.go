package fetch

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func Fetch_post(u string, ch chan<- string, m string, ct string, b string) {
	client := &http.Client{
		Transport: &http.Transport{},
	}
	if b == "" {
		req, err := http.NewRequest(m, u, nil)
		if err != nil {
			ch <- fmt.Sprintf("%v", err)
		}
		req.Header.Set("Content-Type", ct)
		start := time.Now()
		response, err := client.Do(req)
		if err != nil {
			ch <- fmt.Sprintf("Error: %v", err)

		}
		elapsed := time.Since(start).Seconds()
		nbytes, err := io.Copy(ioutil.Discard, response.Body)
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
		start := time.Now()
		response, err := client.Do(req)
		if err != nil {
			ch <- fmt.Sprintf("Error:- %v", err)
		}
		elapsed := time.Since(start).Seconds()
		nbytes, err := io.Copy(ioutil.Discard, response.Body)
		if err != nil {
			ch <- fmt.Sprintf("Error: %v", err)
		}
		defer response.Body.Close()
		ch <- fmt.Sprintf("%.2fs	%7d	%s	%d", elapsed, nbytes, u, response.StatusCode)
	}
}

func Fetch_get(u string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(u)
	if err != nil {
		ch <- fmt.Sprintf("%v", err)
		os.Exit(1)
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	elapsed := time.Since(start).Seconds()
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("%v", err)
		os.Exit(1)
	}
	ch <- fmt.Sprintf("%.2fs	%7d		%s	%d", elapsed, nbytes, u, resp.StatusCode)
}
