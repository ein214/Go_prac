package main

import "fmt"

//Go언어에는 클래스가 없는 대신 구조체에 메서드 연결 가능
//func (리시버명 *구조체타입) 함수명() 리턴값_자료형 {}

type RectangleFunc struct {
	width int
	height int
}

func (rect *RectangleFunc) area() int {
	return rect.width * rect.height
}

func structFunc() {
	rect := RectangleFunc{10,20}

	fmt.Println(rect.area())
}

