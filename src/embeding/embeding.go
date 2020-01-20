package main

import "fmt"

//구조체에서 임베딩을 사용하면 상속과 같은 효과

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello~")
}

type Student struct {
	//p Person			//학생 구조체 안에 사람 구조체를 필드로 가지고있음
	Person
	school string
	grade  int
}

func main() {
	var s Student

	s.Person.greeting()
	s.greeting()
}
