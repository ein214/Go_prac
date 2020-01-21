// main project main.go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/*
		var s string = "한"
		fmt.Println(len(s))

		var r rune = '한'
		fmt.Println(utf8.RuneLen(r))
	*/

	/*
		var s1 string = "안녕하세요"
		fmt.Println(utf8.RuneCountInString(s1))
	*/

	b := []byte("안녕하세요")

	r, size := utf8.DecodeRune(b)
	fmt.Printf("%c %d\n", r, size)

	r, size = utf8.DecodeRune(b[3:])
	fmt.Printf("%c %d\n", r, size)

	r, size = utf8.DecodeLastRune(b)
	fmt.Printf("%c %d\n", r, size)

	r, size = utf8.DecodeLastRune(b[:len(b)-3])
	fmt.Printf("%c %d\n", r, size)

}
