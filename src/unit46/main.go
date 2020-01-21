// main project main.go
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	/*
		fmt.Println(strings.Contains("Hello, world!", "wo"))
		fmt.Println(strings.Contains("Hello, world!", "w o"))
		fmt.Println(strings.Contains("Hello, world!", "ow"))

		fmt.Println(strings.ContainsAny("Hello, world!", "wo"))
		fmt.Println(strings.ContainsAny("Hello, world!", "w o"))
		fmt.Println(strings.ContainsAny("Hello, world!", "ow"))

		fmt.Println(strings.Count("Hello Helium", "He"))		//특정문자열이 몇번 나오는지 구하는 함수

		var r rune
		r = '하'
		fmt.Println(strings.ContainsRune("안녕하세요", r))

		fmt.Println(strings.HasPrefix("Hello, world!", "He"))
		fmt.Println(strings.HasSuffix("Hello, world!", "rld!"))
	*/
	/*
		fmt.Println(strings.Index("Hello, world", "He"))
		fmt.Println(strings.Index("Hello, world", "wor"))
		fmt.Println(strings.Index("Hello, world", "ow"))

		fmt.Println(strings.IndexAny("Hello, world!", "eo"))
		fmt.Println(strings.IndexAny("Hello, world!", "f"))

		var c byte
		c = 'd'
		fmt.Println(strings.IndexByte("Hello, world!", c))
		c = 'f'
		fmt.Println(strings.IndexByte("Hello, world!", c))

		var r rune
		r = '언'
		fmt.Println(strings.IndexRune("고 언어", r))

		f := func(r rune) bool {
			return unicode.Is(unicode.Hangul, r)
		}
		fmt.Println(strings.IndexFunc("Go 언어", f))
		fmt.Println(strings.IndexFunc("Go Language", f))

		fmt.Println(strings.LastIndex("Hello Hello Hello, world!", "Hello"))

		fmt.Println(strings.LastIndexAny("Hello, world", "ol"))

		fmt.Println(strings.LastIndexFunc("Go 언어 안녕", f))
	*/
	/*
		s1 := []string{"Hello,", "world!"}
		fmt.Println(strings.Join(s1, " "))

		s2 := strings.Split("Hello, world!", " ")
		fmt.Println(s2[1])

		s3 := strings.Fields("Hello, world!")	//공백을 기준으로 문자열을 쪼개서 문자열 슬라이스로 저장
		fmt.Println(s3[1])

		f := func(r rune) bool {
			return unicode.Is(unicode.Hangul, r)
		}

		s4 := strings.FieldsFunc("Hello안녕Hello", f)
		fmt.Println(s4)

		fmt.Println(strings.Repeat("Hello", 2))

		fmt.Println(strings.Replace("Hello, world!", "world", "Go", 1))
		fmt.Println(strings.Replace("Hello Hello", "llo", "Go", 2))
	*/

	fmt.Println(strings.Trim("Hello,~~ world!~~", "~~"))
	fmt.Println(strings.Trim("  Hello, world!  ", " "))

	fmt.Println(strings.TrimLeft("~~Hello,~~ world!~~", "~~"))
	fmt.Println(strings.TrimRight("~~Hello,~~ world!~~", "~~"))

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r)
	}

	var s = "안녕 Hello 고 Go 언어"
	fmt.Println(strings.TrimFunc(s, f))
	fmt.Println(strings.TrimLeftFunc(s, f))
	fmt.Println(strings.TrimRightFunc(s, f))

	fmt.Println(strings.TrimSpace("  Hello, world!	  "))

	fmt.Println(strings.TrimSuffix("Hello, world!", "orld!"))
	fmt.Println(strings.TrimSuffix("Hello, world!", "wor"))
}
