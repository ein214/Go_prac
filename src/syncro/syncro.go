/*
고루틴의 실행 흐름을 제어하는 동기화 객체
Mutex: 뮤텍스, 상호배제(mutual exclusion) 여러 스레드(고루틴)에서 공유되는 데이터를 보호할 때 주로 사용
RWMutex: 읽기/쓰기 뮤텍스, 읽기와 쓰기 동작을 나누어서 잠금을 걸 수 있음.
Cond: 조건 변수 condition variable 대기하고 있는 하나의 객체를 깨울 수도 있고 여러개를 동시에 깨울 수도 있음.
Once : 특정 함수를 딱 한번만 실행
Pool : 고루틴에서 사용할 수 있는 객체 풀, 자주 사용하는 객체를 풀에 보관했다가 다시 사용
WaitGroup : 고루틴이 모두 끝날때까지 기다리는 기능
Atomic : 원자적 연산이라고도 하며 더 이상 쪼갤 수 없는 연산, 고루틴, 멀티코어 환경에서 안전하게 값을 연산하는 기능
*/

//뮤텍스 : 여러 고루틴이 공유하는 데이터를 보호할 때 사용
//sync패키지에서 제공하는 뮤텍스 구조체와 함수
//- sync.Mutex
//- func(m *Mutex) Lock(): 뮤텍스 잠금
//- func(m *Mutex) Unlock(): 뮤텍스 잠금해제
/*
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())	//모든 CPU 사용

	var data = []int{}		//int형 슬라이스 생성

	go func() {							//고루틴에서
		for i := 0;i < 1000; i++ {		//1000번 반복하면서
			data = append(data, 1)		//data슬라이스에 1을 추가

			runtime.Gosched()			//다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() {							//고루틴에서
		for i := 0;i < 1000;i++ {		//1000번 반복하면서
			data = append(data, 1)		//data슬라이스에 1을 추가

			runtime.Gosched()			//다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	time.Sleep(2 * time.Second)			//2초 대기
	fmt.Println(len(data))				//data슬라이스의 길이 출력

	//출력하면 대략 1800~1990사이의 값이 나오는데., 두 고루틴이 경합을 벌이면서 동시에 data에 접근했기 때문에 append 함수가 정확하게 처리되지 않은 상황.
	//이러한 상황을 경쟁 조건(Race condition) 이라고 함
}
*/

//위 소스 뮤텍스로보호한 것
/*
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())	//모든 CPU사용

	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() {							//고루틴에서
		for i := 0; i < 1000; i++ {		//1000번 반복하면서
			mutex.Lock()				//뮤텍스 잠금, data슬라이스 보호 시작
			data = append(data, 1)		//data슬라이스에 1 추가
			mutex.Unlock()				//뮤텍스 잠금 해제, data슬라이스 보호 종료

			runtime.Gosched()			//다른 고루틴이 CPU.를 사용할 수 있도록 양보
		}
	}()

	go func() {							//고루틴에서
		for i := 0; i < 1000; i++ {		//1000번 반복하면서
			mutex.Lock()				//뮤텍스 잠금, data슬라이스 보호 시작
			data = append(data, 1)		//data슬라이스에 1 추가
			mutex.Unlock()				//뮤텍스 잠금 해제, data슬라이스 보호 종료 , 짝 안맞춰주면 데드락 발생

			runtime.Gosched()			//다른 고루틴이 CPU.를 사용할 수 있도록 양보
		}
	}()


	time.Sleep(2 * time.Second)

	fmt.Println(len(data))
}
*/

//읽기, 쓰기 뮤텍스  : 쓰기 동작보다 읽기 동작이 많을 때 유리함.
//읽기 락(Read Lock) : 읽기 락끼리는 서로를 막지 않음. 쓰기락은 막음
//쓰기 락(Write Lock) : 쓰기 시도중에 다른곳에서 이전값을 읽으면 안됨. 읽기 쓰기 락 모두 막음
//sync.RWMutex
//func (rw *RWMutex) Lock(), func (rw *RWMutex) Unlock() : 쓰기 뮤텍스 잠금, 잠금 해제
//func (rw *RWMutex) RLock(), func (rw *RWMutex) RUnlock() : 읽기 뮤텍스 잠금, 잠금 해제
/*
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data int = 0
	var rwMutex = new(sync.RWMutex)	//읽기, 쓰기 뮤텍스

	go func() {
		for i := 0;i < 3; i++ {
			rwMutex.Lock()
			data += 1
			fmt.Println("write :", data)
			time.Sleep(10 * time.Millisecond)
			rwMutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read1 : ", data)
			time.Sleep(1 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read2 : ", data)
			time.Sleep(2 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)

}
*/

//조건 변수 : 대기하고 있는 객체 하나, 여러 개 동시에 깨울 때 사용
//sync.Cond
//func NewCond(l Locker) *Cond : 조건 변수 생성
//func (c *Cond) Wait() : 고루틴 실행을 멈추고 대기
//func (c *Cond) Signal() : 대기하고 있는 고루틴을 하나만 깨움
//func (c *Cond) Broadcast() : 대기하고 있는 모든 고루틴을 깨움
//대기중인 고루틴을 하나만 깨우는 예제

/*
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)    //뮤텍스 생성
	var cond = sync.NewCond(mutex) //뮤텍스를 이용하여 조건변수 생성

	c := make(chan bool, 3) //비동기 채널 생성

	for i := 0; i < 3; i++ {
		go func(n int) { //고루틴 3개 생성
			mutex.Lock() //뮤텍스 잠금, cond.Wait() 보호시작
			c <- true    //채널 c에 true를 보냄
			fmt.Println("wait begin : ", n)
			cond.Wait() //조건변수 대기
			fmt.Println("wait end : ", n)
			mutex.Unlock() //뮤텍스 잠금 해제, cond.Wait()보호 종료
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c //채널에서 값을 꺼냄, 고루틴 3개가 모두 실행될 때까지 기다림
	}

	/*
		for i := 0; i < 3; i++ {
			mutex.Lock()		//뮤텍스 잠금, cond.Signal() 보호 시작
			fmt.Println("signal : ", i)
			cond.Signal()		//대기하는 고루틴을 하나씩 깨움
			mutex.Unlock()		//뮤텍스 잠금 해제, cond.Signal()보고 종료
		}
*/
/*
	mutex.Lock()
	fmt.Println("broadcast")
	cond.Broadcast()
	mutex.Unlock()

	fmt.Scanln()

}
*/

//Once : 함수를 한 번만 실행. 반복문안에서 각종 초기화를 할 때 유용
//sync.Once()
//func (*Once) Do(f func())
/*
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func hello() {
	fmt.Println("Hello, world!")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	once := new(sync.Once)	//Once 생성

	for i := 0;i < 3; i++ {
		go func(n int) {		//고루틴 3개 생성
			fmt.Println("goroutine: ", n)

			once.Do(hello)		//고루틴은 3개지만 함수는 한번 실행
		}(i)
	}

	fmt.Scanln()

}
*/

//풀 : 객체를 사용한 후 보관해두었다가 다시 사용하는 기눙(캐시 같은 것)
//메모리 할당과 해제 횟수를 줄여 성능을 높일 때 사용
//sync.Pool
//func(p *Pool) Get() interface{} : 풀에 보관된 객체를 가져옴
//func(p *Pool) Put(x interface{}) : 풀에 객체를 보관
/*
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Data struct { //Data 구조체 정의
	tag    string //풀 태그
	buffer []int  //데이터 저장용 슬라이스
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	pool := sync.Pool{ //풀을 할당
		New: func() interface{} { //Get함수를 사용했을 때 호출될 함수 정의
			data := new(Data) //새메모리 할당
			data.tag = "new"
			data.buffer = make([]int, 10) //슬라이스 공간 할당
			return data
		},
	}

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data) //풀에서 *Data 타입으로 데이터 가져옴

			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100)
			}

			fmt.Println(data)

			data.tag = "used"
			pool.Put(data)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data) //풀에서 *Data 타입으로 데이터 가져옴

			n := 0
			for index := range data.buffer {
				data.buffer[index] = n
				n += 2
			}

			fmt.Println(data)

			data.tag = "used"
			pool.Put(data)
		}()
	}

	fmt.Scanln()

}
*/

//대기그룹 : 고루틴이 모두 끝날 때까지 기다릴 때 사용.
//sync.WaitGroup
//func (wg *WaitGroup) Add(delta int) : 대기 그룹에 고루틴계수 추가
//func (wg *WaitGroup) Done() : 고루틴이 끝났다는 걸 알려줄 때 사용
//func (wg *WaitGroup) Wait() : 모든 고루틴이 끝날 때까지 기다림ㄴ
/*
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := new(sync.WaitGroup) //대기그룹 생성

	for i := 0; i < 10; i++ {
		wg.Add(1) //반복할 때마다 wg.Add 함수로 1씩 추가
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("the end")

}
*/

//원자적 연산 - 더 이상 쪼갤 수 없는 연산
//여러 스레드, CPU코어에서 같은 변수를 수정할 때 서로 영향을 받지않고 안전하게 연산 가능 
package main 

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"		//원자적 연산 할때 추가적ㅇ로 필요함
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			//data += 1
			atomic.AddInt32(&data, 1)		//원자적 연산으로 1씩 더함
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//data -= 1
			atomic.AddInt32(&data, -1)		//원자적 연산으로 1씩 뺌
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(data)
}