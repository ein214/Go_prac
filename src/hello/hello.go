package main //Go언어로 작성한 프로그램음 main 패키지부터 실행

import (
	"fmt" //문자열을 출력하기 위해 fmt 패키지 사용
	"calc"
	//"math"
	//"unicode/utf8" //한글, 한자, 일본어 등 2바이트 넘는 문자열 길이 구할 때
)

func main() {
	fmt.Println(calc.Sum(1,2))
	//fmt.Println("Hello, world!")
	/*
	var a float64 = 10.0
	for i := 0; i < 10; i++ {
		a = a - 0.1
	}
	fmt.Println(a)
	//fmt.Println(a == 9.0)

	const epsilon = 1e-14                   //GO언어 머신 엡실론
	fmt.Println(math.Abs(a-9.0) <= epsilon) //true: 연산한 값과 비교할 값의 차이를 구한 뒤 머신 엡실론보다 작거나 같은지 비교
	*/
	/*
		어떤 실수를 가장 가까운 부동소수점 수로 반올림하였을 때 상대 오차는 항상 머신 엡실론 이하 (Go 언어 머신 엡실론 값 = 1e-14)
		따라서 연산한 값과 비교할 값의 차이가 머신 앱실론보다 작거나 같다면 두 실수는 같은 값
		Math.Abs를 쓴 이유는 차이가 음수가 나올 수도 있으므로 절대값으로 만들기 위해

		machine epsilon : 부동소수점 연산에서 반올림을 함으로써 발생하는 오차의 상한.
	*/
	/*
	var num1 uint8 = 3
	var num2 uint8 = 2

	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)
	fmt.Println(num1 << num2) // 3 * 2의 num2제곱
	fmt.Println(num1 >> num2) // 3 * 2의 -num2제곱
	fmt.Println(^num1)

	fmt.Println("문자 연산")

	var s1 string = "한글"
	var s2 string = "한글"
	var s3 string = "Go"

	fmt.Println(s1 == s2)
	fmt.Println(s1 + s3 + s2)
	fmt.Println("안녕하세요" + s2)

	var s4 string = "Hello"
	fmt.Printf("%c", s4[1])
	*/
	/*
		병렬할당(Parallet assignment) : 변수 여러 개에 값을 콤마로 구분하여 대입하는 방법
	*/
	/*
	var x, y int
	var age int

	x, y, age = 10, 20, 5

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(age)

	/*
		Go언어는 반복문이 for문만 있음
		for 초깃값;조건식;변화식 {}
	*/
/*
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//goto keyword
	/*
	var g int = 1

	if g == 1 {
		goto ERROR1
	}

	if g == 2 {
		goto ERROR2
	}

	if g == 3 {
		goto ERROR1
	}

	fmt.Println(g)
	fmt.Println("Success")
	*/
	/*
	s := 4

	switch s {
	case 4:
		fmt.Println("4 이상")
		fallthrough
	case 3:
		fmt.Println("3 이상")
		fallthrough
	case 2:
		fmt.Println("2 이상")
		fallthrough
	case 1:
		fmt.Println("1 이상")
		fallthrough
	case 0:
		fmt.Println("0 이상")

	}

	arr := [5]int{32, 29, 78, 16, 81}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	*/
	// var 맵명 map[string]int 
	/*
	a := map[string]int{"Hello":10, "world":20}
	b := map[string]int{
		"Hello": 10,
		"world": 20,
	}

	

	fmt.Println(a["Hello"])
	fmt.Println(b["world"])

	//배열을 순회할 때는 for반복문에서 range 키워드 사용
	//배열을 대입하면 배열 전체가 복사
	//슬라이스는 길이가 고정되어 있지 않으며 동작으로 크기가 늘어남 배열과는 달리 레퍼런스 타입
	//슬라이스를 복사할 때는 copy 함수 사용
	


	/*
	ERROR1:
	fmt.Println("Error 1")
	return

ERROR2:
	fmt.Println("Error 2")
	return
*/
	
}
