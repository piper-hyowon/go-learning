package fetchall

// fetchall: URL을 병렬로 반입하고 시간과 크기를 보고

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetchall() {
	start := time.Now()
	ch := make(chan string) // 문자열의 채널 생성
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // 고루틴 시작
		// 각 커맨드라인 인수로 fetch 비동기 호출
	}

	// fetch 결과가 도착하는대로 출력(모든 고루틴이 완료될 때까지 결과 수신)
	for range os.Args[1:] {
		fmt.Println(<-ch) // ch 채널에서 결과 수신하여 출력
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // 에러 메시지 채널에 전송
		return
	}
	// 바이트수와 에러 반환
	nbytes, err := io.Copy(io.Discard, resp.Body) // 본문을 읽고 내용은 io.Discard에 출력해 폐기함
	resp.Body.Close()                             // 리소스 노출 방지
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
