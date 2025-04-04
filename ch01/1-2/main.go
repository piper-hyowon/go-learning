package main

// 1.2 echo 프로그램을 수정해 각 인자의 인덱스와 값을 한 줄에 하나씩 출력하라.

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args[0:] {
		fmt.Println(i, v)
	}
}

// 실행 예시:
// go run main.go --version 3.0 --username hyowon
//
// 출력 결과:
// 0 /path/to/executable
// 1 --version
// 2 3.0
// 3 --username
// 4 hyowon
//
// 주요 개념:
// - os.Args: 명령행 인수를 담고 있는 문자열 슬라이스
// - os.Args[0]: 프로그램의 실행 경로
// - os.Args[1:]: 사용자가 입력한 명령행 인수들
