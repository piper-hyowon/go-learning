package main

// 1.7 io.Copy(dst, src) 함수 호출은 src에서 읽어 dst에 기록한다.
// ioutil.ReadAll대신 이 함수를 사용해 결과 본문을
// 전체 스트림을 저장하는 큰 버퍼 없이 os.Stdout 으로 복사하라.
// io.Copy의 오류 결과를 확인해야한다.

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close() // 코드가 복잡해지면 defer 사용

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)

		}
	}

}
