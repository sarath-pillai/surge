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
	start := time.Now()
	if b == "" {
		resp, err := http.Post(u, ct, nil)
		if err != nil {
			ch <- fmt.Sprintf("%v", err)
		}
		nbytes, err := io.Copy(ioutil.Discard, resp.Body)
		elapsed := time.Since(start).Seconds()
		resp.Body.Close()
		if err != nil {
			ch <- fmt.Sprintf("%v", err)
			os.Exit(1)
		}
		ch <- fmt.Sprintf("%.2fs	%7d		%s		%d", elapsed, nbytes, u, resp.StatusCode)
	} else {
		data, err := os.Open(b)
		if err != nil {
			ch <- fmt.Sprintf("%v", err)
			log.Fatal(err)
		}
		resp, err := http.Post(u, ct, data)
		if err != nil {
			ch <- fmt.Sprintf("%v", err)
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
