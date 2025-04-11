package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// 1-8 각 인수 URL에 http:// 접두사가 누락된 경우 이를 추가하도록 fetch를 수정하라.
// strings.HasPrefix를 사용할 수도 있다.

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

	}
}
