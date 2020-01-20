# Go_prac
###  가장 빨리 만나는 GO언어 실습내용

- UNIT 24까지 

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

    

    