package main

import "fmt"
import "printf_Test"

func main() {
	printf_Test.Printf_test()
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

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(m)
	fmt.Println(p)
	fmt.Println(data)
	fmt.Println(i)

	fmt.Println(num1, num2, num3, s, b)
	fmt.Println(p, a, m)
	fmt.Println(data, i)
}