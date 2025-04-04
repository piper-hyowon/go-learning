package main

// 1.1 echo 프로그램을 수정해 호출한 명령인 osArgs[0]도 같이 출력하라.

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
