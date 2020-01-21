//고루틴(Goroutine) : 함수를 동시에 실행시키는 기능
package main
/*
import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int) {
	r := rand.Intn(100)		//랜덤한 숫자 생성
	time.Sleep(time.Duration(r))		//랜덤한 시간동안 대기
	fmt.Println(n)
}

func main() {
	for i := 0;i < 100; i++ {
		go hello(i)
	}

	fmt.Scanln()
}
*/
//클로저를 고루틴으로 실행
import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)		//CPU를 하나만 사용
	
	s := "Hello, world!"

	for i:=0; i < 100; i++ {
		go func() {
			fmt.Println(s, i)
		}()
	}

	fmt.Scanln()
}
