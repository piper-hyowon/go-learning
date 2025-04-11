package main

/*
fmt 패키지 함수 정리
1. 표준 출력(콘솔)에 출력
- fmt.Print(), fmt.Println()
- fmt.Printf(): 형식 지정자 사용

2. F로 시작하는 - 지정된 io.Writer(파일 등)에 출력
- fmt.Fprint(), fmt.Fprintln()
- fmt.Fprint(): 지정된 Writer에 형식 지정자 사용

3. S로 시작하는 - 문자열로 반환
- fmt.Sprint(), fmt.Sprintln(), fmt.Srpintf()
*/
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// 1-10 많은 양의 데이터를 생성하는 웹 사이트를 찾고,
// fetchall을 두 번 연속으로 실행하고 결과 시간이 얼마나 달라지는지를 통해 캐시 여부를 조사하라.
// 매번 같은 내용을 받는가? 이 결과를 조사하기 위해 fetchall이 결과를 파일로 출력하게 수정하라.

func main() {
	file, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "파일 생성 오류: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[2:] {
		go fetch(url, ch)
	}

	for range os.Args[2:] {
		fmt.Fprintln(file, <-ch) // 채널에서 받은 값을 파일에 출력
	}
	fmt.Fprintf(file, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // err를 문자열로 변환(ch <- err 은 타입에러!)
		return
	}
	defer resp.Body.Close() // 리소스 획득 후 바로 defer 등록(권장 패턴)
	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while raeding %s: %v", url, err)
	}
	ch <- fmt.Sprintf("%.2fs  %7d %s", time.Since(start).Seconds(), nbytes, url)
}
