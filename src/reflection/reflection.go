//리플렉션 : 실행시점에 인터페이스나 구조체 등의 타입 정보를 얻어내거나 결정하는 기능
/*
package main

import (
	"fmt"
	"reflect"
)

type Data struct { //구조체 정의
	a, b int
}

func main() {
	var num int = 1
	fmt.Println(reflect.TypeOf(num)) // 자료형 이름 출력

	var s string = "Hello, world!"
	fmt.Println(reflect.TypeOf(s)) // 자료형 이름 출력

	//var f float32 = 1.3
	//fmt.Println(reflect.TypeOf(f))	// 자료형 이름 출력

	var data Data = Data{1, 2}
	fmt.Println(reflect.TypeOf(data)) // 자료형 이름 출력

	var f float64 = 1.3
	t := reflect.TypeOf(f)  //f의 타입 정보를 t 에 저장
	v := reflect.ValueOf(f) //f 의 값 정보를 v에 저장

	fmt.Println(t.Name())
	fmt.Println(t.Size())
	fmt.Println(t.Kind() == reflect.Float64)

	fmt.Println(t.Kind() == reflect.Int64)

	fmt.Println(v.Type())
	fmt.Println(v.Kind() == reflect.Float64)

	fmt.Println(v.Kind() == reflect.Int64)

	fmt.Println(v.Float())
}
*/

//구조체 태그 가져오기
/*
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `tag1:"이름" tag2:"Name"`
	age  int    `tag2:"나이" tag2:"Age"`
}

func main() {
	p := Person{}

	name, ok := reflect.TypeOf(p).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2"))

	age, ok := reflect.TypeOf(p).FieldByName("age")
	fmt.Println(ok, age.Tag.Get("tag2"), age.Tag.Get("tag2"))


}
*/

//포인터와 인터페이스의 값 가져오기
/*
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a *int = new(int)
	*a = 1

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
	//fmt.Println(reflect.ValueOf(a).Int())			//포인터일 경우에는 이렇게 못가져옴 
	fmt.Println(reflect.ValueOf(a).Elem())
	fmt.Println(reflect.ValueOf(a).Elem().Int())

	var b interface{}
	b = 1

	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.ValueOf(b))
	fmt.Println(reflect.ValueOf(b).Int())
	fmt.Println(reflect.ValueOf(b).Elem())			//인터페이스일 경우에는 Elem을 써서 못가져옴 
}
*/

//리플렉션 이용하여 동적으로 함수 생성 
//reflect,MakeFunc 사용 -
package main

import (
	"fmt"
	"reflect"
)

func h(args []reflect.Value) []reflect.Value {		//매개변수와 리턴값은 반드시 []reflect.Value 를 사용 
	fmt.Println("Hello, world!")
	return nil
}

func main() {
	var hello func()		//함수를 담을 변수 선언

	fn := reflect.ValueOf(&hello).Elem()		//hello의 주소를 넘긴 뒤 Elem으로 값 정보를 가져옴

	v := reflect.MakeFunc(fn.Type(), h)	//h의 함수 정보를 생성

	fn.Set(v)		//hello의 값 정보인 fn에 h의 함수정보 v를 설정하여 v함수를 연결

	hello()
}