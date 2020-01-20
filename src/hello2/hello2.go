//클로저(Closure) : 함수를 선언 및 정의할 수 있고, 바깥쪽 함수에 선언된 변수에도 접근할 수 있는 함수
//지연 호출 : 특정 함수를 현재 함수가 끝나기 직전에 실행. defer 함수명() or defer 함수명(매개변수)
//패닉과 복구 : panic(에러_메시지) 로직에 따라 에러로 처리하고 싶을 때 사용
//			  recover 함수는 패닉이 발생했을 때 프로그램이 바로 종료되지 않고 예외처리를 할 수 있음. recover 함수는 반드시 지연 함수로 사용. 안그러면 프로그램이 바로 종료되버림. 0
// 변수 := recover()

package main

import (
	"fmt"
)

/*
func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}
*/
/*
func HelloWorld() {
	defer func() {
		fmt.Println("world")
	}()

	func() {
		fmt.Println("Hello")
	}()
}
*/
/*
func ReadHello() {
	file, err := os.Open("Hello.txt")
	defer file.Close() //지연 호출한 file.Close()가 맨 마지막에 호출됨

	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 100)
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))
}
*/
func f() {
	defer func() {
		s := recover()

		fmt.Println(s)
	}()

	panic("Error !!!")
}
func main() {
	f()

	fmt.Println("Hello, world!")

	//ReadHello()
	/*
		hello()
		defer world()		//현재 함수가 끝나기 직전에 호출


		hello()
		hello()
	*/
	/*
		HelloWorld()

		for i := 0; i < 5; i++ {
			defer fmt.Printf("%d", i)
		}
	*/
}

/* 클로저
func calc() func(x int) int {
	a, b := 3, 5
	return func(x int) int {
		return a*x + b	//함수 바깥의 a,b 변수 사용
	}
}

func main() {

	sum := func(a,b int) int {
		return a+b
	}

	r := sum(1,2)

	fmt.Println(r)


	f := calc()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
	fmt.Println(f(5))

}
*/
