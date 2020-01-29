package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile(
		"test.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	//strings 패키지로 읽어서 io패키지 내 copy로 쓰기 인스턴스에 복사 후 데이터 저장
	s := "test Hello world!"
	r := strings.NewReader(s)

	w := bufio.NewWriter(file)
	io.Copy(w, r)
	w.Flush()

	//os.Stdout으로 바로 문자열 읽기
	r2 := strings.NewReader(s)
	io.Copy(os.Stdout, r2)

	file2, err := os.OpenFile(
		"test2.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file2.Close()

	w2 := bufio.NewWriter(file2)
	r4 := bufio.NewReader(file)

	rw := bufio.NewReadWriter(r4, w2)
	file.Seek(0, os.SEEK_SET)
	data, _, _ := rw.ReadLine()

	fmt.Fprintf(w2, "%s %T", string(data), data)
	w2.Flush()
}
