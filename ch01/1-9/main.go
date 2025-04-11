package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// 1.9 fetch를 수정해 resp.Status에 있는 HTTP 응답 코드도 같이 출력하라.

const (
	prefix = "https://"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("HTTP Status: %s\n", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		resp.Body.Close()

	}

}
