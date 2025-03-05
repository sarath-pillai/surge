package fetch

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var client = &http.Client{
	Transport: &http.Transport{},
}

func Fetch_post(u string, ch chan<- string, m string, ct string, b string, h []string, a string) {
	if b == "" {
		req, err := http.NewRequest(m, u, nil)
		if err != nil {
			ch <- fmt.Sprintf("%v", err)
		}
		req.Header.Set("Content-Type", ct)
		if len(h) > 0 {
			for _, v := range h {
				pairs := strings.Split(v, ":")
				req.Header.Set(pairs[0], pairs[1])
			}
		}
		if len(a) > 0 {
			cred := strings.Split(a, ":")
			req.SetBasicAuth(cred[0], cred[1])
		}

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
		if len(h) > 0 {
			for _, v := range h {
				pairs := strings.Split(v, ":")
				req.Header.Set(pairs[0], pairs[1])
			}
		}
		if len(a) > 0 {
			cred := strings.Split(a, ":")
			req.SetBasicAuth(cred[0], cred[1])
		}
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

func Fetch_get(u string, ch chan<- string, m string, h []string, a string) {
	if len(a) > 0 {
		req, err := http.NewRequest(m, u, nil)
		if err != nil {
			ch <- fmt.Sprintf("Error: %v", err)
		}
		for _, v := range h {
			pairs := strings.Split(v, ":")
			req.Header.Set(pairs[0], pairs[1])
		}
		cred := strings.Split(a, ":")
		req.SetBasicAuth(cred[0], cred[1])
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
		ch <- fmt.Sprintf("%.2fs	%7d		%s	%d", elapsed, nbytes, u, response.StatusCode)
	} else {
		req, err := http.NewRequest(m, u, nil)
		if err != nil {
			ch <- fmt.Sprintf("%v", err)
			os.Exit(1)
		}
		start := time.Now()
		response, err := client.Do(req)
		nbytes, err := io.Copy(ioutil.Discard, response.Body)
		elapsed := time.Since(start).Seconds()
		defer response.Body.Close()
		if err != nil {
			ch <- fmt.Sprintf("%v", err)
			os.Exit(1)
		}
		ch <- fmt.Sprintf("%.2fs	%7d		%s	%d", elapsed, nbytes, u, response.StatusCode)
	}
}
