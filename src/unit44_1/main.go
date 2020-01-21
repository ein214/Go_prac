// main project main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	file1, _ := os.Create("hello1.txt")        //hello1.txt 파일 생성
	defer file1.Close()                        //main 함수 끝나기 전에 파일을 닫음
	fmt.Fprint(file1, 1, 1.1, "Hello, world!") //값을 그대로 문자열로 만든 뒤 파일에 저장

	file2, _ := os.Create("hello2.txt")
	defer file2.Close()
	fmt.Fprintln(file2, 1, 1.1, "Hello, world!")

	file3, _ := os.Create("hello3.txt")
	defer file3.Close()
	fmt.Fprintf(file3, "%d, %f, %s", 1, 1.1, "hello, world!")
}
