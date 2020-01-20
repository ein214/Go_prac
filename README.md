# Go_prac
###  가장 빨리 만나는 GO언어 실습내용

- [2019-11-01] UNIT28

  ```go
/*포인터형 변수 var 변수명 *자료형*/
  var numPtr *int

  /*빈 포인터형 변수 메모리 할당 포인터_변수 = new(자료형)*/
  var numPtr *int = new(int)
  
  /*역참조(dereference) *포인터_변수명*/
  *numPtr = 1
  
/*일반 변수에 참조를 사용하면 포인터형 변수에 대입가능 &변수명*/
  var num int = 1
var numPtr * int = &num
  ```
  

  
  
  
- [2019-10-28 ~ 2019-10-31] UNIT 25 - UNIT 27

  - 클로저(Closure) : 함수를 선언 및 정의할 수 있고, 바깥쪽 함수에 선언된 변수에도 접근할 수 있는 변수

  - 지연 호출 : 특정 함수를 현재 함수가 끝나기 직전에 실행

    ```go
    defer 함수명()
    defer 함수명(매개변수)
    ```
  
  - 패닉과 복구 

    - panic(에러_메시지) 로직에 따라 에러로 처리하고 싶을 때 사용
  - recover 함수는 패닉이 발생했을 때 프로그램이 바로 종료되지 않고 예외처리를 할 수 있음. 
      recover함순은 반드시 지연 함수로 사용. 안그러면 프로그램 바로 종료 
  
- [2019-10-17 ~ 2019-10-28] UNIT 24까지 

  - machine epsilon : 부동소수점 연산에서 반올림을 함으로써 발생하는 오차의 상한

  - 병렬할당(Parallet assignment) :  변수 여러개에 값을 콤마로 구분하여 대입

    ```go
    var x, y, int
    var age int
    
    x,y,age = 10,20,5
    ```

    

  - Go언어는 반복문이 for문만 있음
    for 초깃값;조건식;변화식 {}

    ```go
    for i := 0; i < 5; i++ {
    	fmt.Println(i)
    }
    ```

    

    