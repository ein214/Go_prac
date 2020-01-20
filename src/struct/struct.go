//type 구조체명 struct {} : 여러 변수를 담을 수 있는 구조체 
//var 변수명 구조체_타입 : 생성한 구조체 타입으로 인스턴스를 생성하는 것 
//구조체_포인터 = new(구조체타입) : 포인터에 메모리공간을 할당하는 것 
package main

import "fmt"
type Rectangle struct {
	width int
	height int
}

//다른 언어의 생성자(contructor) 흉내
/*
func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height}		//구조체 인스턴스를 생성한 뒤 포인터를 리턴 
}
*/

//사각형의 넓이를 구하는 함수 
func rectangleArea(rect *Rectangle) int {		//매개변수로 구조체 포인터를 받음
	return rect.width * rect.height
}

func rectangleScaleA(rect *Rectangle, factor int) {		//매개변수 구조체 포인터
	rect.width = rect.width * factor		//포인터이므로 원래값이 변경
	rect.height = rect.height * factor		//포인터이므로 원래값이 변경
}

func rectangleScaleB(rect Rectangle, factor int) {		//매개변수로 구조체 인스턴스를 받음 
	rect.width = rect.width * factor		//값이 복사된거라 원래값에는 영향없음
	rect.height = rect.height * factor		//값이 복사된거라 원래값에는 영향없음
}

func main() {
	/*
	var rect1 Rectangle		//구조체 인스턴스 생성
	var rect2 *Rectangle = new(Rectangle)	//구조체 포인터 선언 후 메모리 할당

	rect1.height = 20	//구조체 인스턴스의 필드에 접근할 때 . 사용
	rect2.height = 62	//구조체 포인터에 접근할 때도 . 사용

	fmt.Println(rect1)
	fmt.Println(rect2)		//구조체 포인터이기 때문에 앞에 &가 붙음
	*/

	//rect := NewRectangle(20, 10)

	//fmt.Println(rect)

	/*
	rect := Rectangle{30, 30}
	area := rectangleArea(&rect)

	fmt.Println(area)
	*/

	rect1 := Rectangle{30,30}
	rectangleScaleA(&rect1, 10)		//구조체의 포인터를 넘김
	fmt.Println(rect1)

	rect2 := Rectangle{30,30}
	rectangleScaleB(rect2, 10)
	fmt.Println(rect2)


}