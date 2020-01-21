package printf_Test

import "fmt"

func Printf_test() {
	var num1 int = 10
	var num2 float32 = 3.2
	var num3 complex64 = 2.5 + 8.1i
	var s string = "Hello, world!"
	var b bool = true
	var a []int = []int{1,2,3}
	var m map[string]int = map[string]int{"Hello": 1}
	var p *int = new(int)
	type Data struct { a, b int }
	var data Data = Data{1,2}
	var i interface{} = 1

	fmt.Printf("정수 : %d\n", num1)
	fmt.Printf("실수 : %f\n", num2)
	fmt.Printf("복소수 : %f\n", num3)

	fmt.Printf("문자열 : %s\n", s)
	fmt.Printf("불 : %t\n", b)
	fmt.Printf("슬라이스: %v\n", a)
	fmt.Printf("맵 : %v\n", m)
	fmt.Printf("포인터 : %p\n", p)
	fmt.Printf("구조체 : %v\n", data)
	fmt.Printf("인터페이스 : %v\n", i)

	fmt.Printf("%d %f %s\n", num1, num2, s)

}