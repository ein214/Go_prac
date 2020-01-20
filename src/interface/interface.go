package main

import (
	"fmt"
	"strconv"
)

func f1(arg interface{}) {

}

type Any interface{}

func f2(arg Any) {

}

func formatString(arg interface{}) string {
	switch arg.(type) {
	case int:
		i := arg.(int)
		return strconv.Itoa(i)
	case float32:
		f := arg.(float32)
		return strconv.FormatFloat(float64(f), 'f', -1, 32)
	case float64:
		f := arg.(float64)
		return strconv.FormatFloat(f, 'f', -1, 64)
	case string:
		s := arg.(string)
		return s
	default:
		return "Error"
	}
}

func main() {
	fmt.Println(formatString(1))
	fmt.Println(formatString(2.5))
	fmt.Println(formatString("Hello, world!"))
}

//인터페이스 : 메서드 집합, 단 메서드 자체를 구현하지는 않는다
/*
type hello interface {	//인터페이스 정의

}

func main() {
	var h hello		//인터페이스 선언
	fmt.Println(h)	// 빈 인터페이스를 선언했기 때문에 nil값이 나옴
}
*/
/*
type MyInt int 	//int형을 MyInt로 정의

func (i MyInt) Print() {		//MyInt에 Print 메서드 연결
	fmt.Println(i)
}

type Rectangle struct {
	width, height int
}

func (r Rectangle) Print() {
	fmt.Println(r.width, r.height)
}

type Printer interface {
	Print()
}

func main() {
	var i MyInt = 5
	r := Rectangle{10, 20}

	var p Printer

	p = i
	p.Print()

	p = r
	p.Print()

}
*/
