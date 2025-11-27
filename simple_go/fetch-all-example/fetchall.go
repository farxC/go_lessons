package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	hasPrefix := strings.HasPrefix(url, "http://")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	var filename string
	if !hasPrefix {
		filename = "file_" + url
		url = "http://" + url
	} else {
		filename = "file_" + url[7:]
	}
	filename = strings.ReplaceAll(filename, "/", "_")

	fmt.Println(filename)
	resp, err := client.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2.fs %7d %s", secs, nbytes, url)

}
