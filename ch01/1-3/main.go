package main

// 1.3 잠재적으로 비효율적인 버전과 strings.Join을 사용하는 버전의 실행시간 차이를 실험을 통해 측정하라

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// 시간 측정에 time.Since() 사용:
// - time.Since(start)는 time.Now().Sub(start)와 기능적으로 동일한 '헬퍼 함수'
// - 더 간결하고 읽기 쉬움, 일반적으로 사용됨

// time.Since(주어진 시간): 주어진 시간부터 현재까지의 경과 시간을 직접 반환

func main() {
	// 잠재젹으로 비효율적인 버전
	start1 := time.Now()
	s, sep := "", ""
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i] // 문자열 + 연산자 -> 두 값을 결합
		sep = " "
	}
	fmt.Println(s)
	case1 := time.Since(start1)

	// strings.Join 사용 버전
	start2 := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	case2 := time.Since(start2)

	if case1 > case2 {
		fmt.Println("case1 이", case1-case2, "만큼 더 오래 걸린다.")
	} else if case1 < case2 {
		fmt.Println("case2 가", case2-case1, "만큼 더 오래 걸린다.")
	} else {
		fmt.Println("동일")
	}
}
