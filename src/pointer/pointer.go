package main

import "fmt"

func hello(n *int) {
	*n = 2	//매개변수 n에 2를 대입
}
func main() {
	//var numPtr *int		//포인터형 변수를 선언, nil(NULL값)로 초기화됨
	//var numPtr *int = new(int)		//빈 포인터형 변수는 바로 사용할 수 없으므로 new함수로 메모리 할당
									//new 함수는 지정한 자료형의 크기에 맞는 메모리 공간을 할당. Go언어는 메모리를 관리해주는 가비지 컬렉션을 지원하므로 메모리를 할당한 뒤 해제하지 않아도 됩니다. 
	
	//*포인터_변수명
	//var numPtr *int = new(int)	//공간할당
	//*numPtr = 1					//역참조로 포인터형 변수에 값을 대입

	//&변수명 : 일반 변수에 참조(레퍼런스)를 사용하면 포인터형 변수에 대입할 수 있음.
	//var num int = 1
	//var numPtr *int = &num		//참조로 num변수의 메모리 주소를 구하여 numPtr 포인터 변수에 대입

	//fmt.Println(numPtr)
	//fmt.Println(&num)

	var n int = 1

	hello(&n) 

	fmt.Println(n)
}

