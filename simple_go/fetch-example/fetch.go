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
	for _, url := range os.Args[1:] {
		start := time.Now()
		hasPrefix := strings.HasPrefix(url, "http://")
		if !hasPrefix {
			url = "http://" + url
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		w, err := io.Copy(os.Stdout, resp.Body)
		s := resp.Status

		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v", url, err)
		}
		fmt.Printf("%v\n", w)
		fmt.Printf("%s\n", s)
		fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	}
}
