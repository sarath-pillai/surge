package fetch

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func Fetch(u string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(u)
	if err != nil {
		ch <- fmt.Sprintf("%v", err)
		os.Exit(1)
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	elapsed := time.Since(start).Seconds()
	if err != nil {
		ch <- fmt.Sprintf("%v", err)
		os.Exit(1)
	}
	ch <- fmt.Sprintf("%.2fs	%7d		%s", elapsed, nbytes, u)
}
