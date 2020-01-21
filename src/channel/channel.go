//채널(channel) : 고루틴끼리 데이터를 주고받고, 실행 흐름을 제어하는 기능. 모든타입을 채널로 사용할 수 있음. 채널 자체는 값이 아닌 레퍼런스 타입
//make(chan 자료형)

//채널으,ㄹ 사용하기 전에는 반드시 make함수로 공간 할당, 생성하면 동기채널(synchronous channel) 이 행성
//채널을 매개변수로 받는 함수는 반드기 go 키워드를 사용하여 고루틴으로 실행
package main

/*
import "fmt"

func sum(a int, b int, c chan int) {
	c <- a+b		//채널에 a+b를 보냄
}

func main() {
	c := make(chan int)	//int형 채널생성

	go sum(1,2,c)		//sum을 고루틴으로 실행한뒤 채널을 매개변수로 넘겨중

	n := <-c		//채널에서 값을 꺼낸 뒤 n에 대입
	fmt.Println(n)
}
*/

//동기채널
/*
import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)		//bool타입 동기채널 생성
	count := 3					//반복할 횟수 

	go func() {
		for i := 0; i < count; i++ {
			done <- true	//고루틴에 true를 보냄, 값을 꺼낼 때까지 대기
			fmt.Println("고루틴 : ", i)	//반복문의 변수 출력
			time.Sleep(1 * time.Second)		//1초 대기
		}
	}()

	for i := 0; i < count; i++ {
		<-done							//채널에 값이 들어올 때까지 대기, 값을 꺼냄.
		fmt.Println("메인 함수 : ",i)	//반복문의 변수 출력
	}
}

*/

//채널 버퍼링, 채널에 버퍼를 1개 이상 설정하면 비동기 채널 생성 
//make(chan 자료혀으 버퍼_개수)
/*
import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2)		//버퍼가 2개인 비동기 채널 생성
	count := 4

	go func() {
		for i := 0; i < count; i++	{
			done <- true
			fmt.Println("고루틴 : ",i)
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("메인 함수:", i)
	}
}
*/

//range와 close
//이미 닫힌 채널에 값을 보내면 패닉 발생
//채널을 닫으면 range루프가 종료
//채널이 열러있고 값이 들어오지 않는다면 range는 실행되지 않고 계속 대기, 만약 다른 곳에서 채널에 값을 보냈다면 그때부터 range가 계속 반복
/*
import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)		//채널을 닫음
	}()

	for i := range c {	//range를 사용하여 채널에서 값을 꺼냄
		fmt.Println(i)
	}
}
*/

//보내기전용채널과 받기전용채널
/*
import "fmt"

func producer(c chan<- int) {	//보내기 전용 채널
	for i := 0;i < 5; i++ {
		c <- i
	}

	c <- 100	//채널에 값을 보냄 

	//fmt.Println(<-c)
}

func consumer(c <-chan int) {	//받기 전용 채널
	for i := range c {
		fmt.Println(i)
	}
	
	fmt.Println(<-c)		//채널에 값을 꺼냄

	// c <- 1
}

func main() {
	c := make(chan int)

	go producer(c)
	go consumer(c)

	fmt.Scanln()

}
*/

//채널을 리턴값으로 사용
/*
import "fmt"

func sum(a, b int) <-chan int {
	out := make(chan int)		//채널생성
	go func() {
		out <- a+b
	}()

	return out
}

func main() {
	c := sum(1,2)

	fmt.Println(<-c)
}
*/

//채널만 사용하여 값을 더하기 
/*
import "fmt"

func num(a, b int) <-chan int {
	out := make(chan int)
	go func() {
		out <- a
		out <- b
		close(out)
	}()

	return out
}

func sum(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		r := 0
		for i := range c {
			r = r+i
		}
		out <- r
	}()

	return out
}

func main() {
	c := num(1,2)
	out := sum(c)

	fmt.Println(<-out)
}
*/

//셀렉트 사용하기
//select { case <-채널: 코드 }
/*
import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)	//int형 채널 생성
	c2 := make(chan string)	//string 채널 생성

	go func() {
		for {
			c1 <- 10	//c1에 10을 보냄
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Hello, world!"		//채널 c2에 Hello, work 보냄
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case i := <-c1:
				fmt.Println("c1 :",i)
			case s := <-c2:
				fmt.Println("c2 :",s)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
*/

//select 분기문은 채널에 값을 보낼 수도 있음
//case 채널<-값: 코드
import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		for {
			i := <-c1
			fmt.Println("c1 :",i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Hello, world!"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case c1 <- 10:
			case s := <-c2:
				fmt.Println("c2 :", s)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
