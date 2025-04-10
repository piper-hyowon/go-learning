package main

import (
	"bufio"
	"fmt"
	"os"
)

// dup2: 표준 입력 or 파일 목록에서 텍스트를 읽고, 두 번이상 나타나는 각 줄의 카운트와 텍스트 출력

// 1.4 dup2를 수정해 중복된 줄이 있는 파일명을 모두 출력하라.
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		if hasDuplicates(counts) {
			fmt.Println("표준 입력에서 중복")
		}

	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			// 맵 초기화
			for k := range counts {
				delete(counts, k)
			}

			countLines(f, counts)
			f.Close()

			if hasDuplicates(counts) {
				fmt.Println(arg)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// input.Err() 에서의 잠재적 오류는 무시
}

func hasDuplicates(counts map[string]int) bool {
	for _, n := range counts {
		if n > 1 {
			return true
		}
	}
	return false

}
