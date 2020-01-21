# Go_prac
###  가장 빨리 만나는 GO언어 실습내용

- [2019-11-08] Unit 35 동기화 객체

  - 고루틴의 실행 흐름을 제어하는 동기화 객체
  - Mutex: 뮤텍스, 상호배제(mutual exclusion) 여러 스레드(고루틴)에서 공유되는 데이터를 보호할 때 주로 사용
      - sync.Mutex
      - func(m *Mutex) Lock(): 뮤텍스 잠금
      - func(m *Mutex) Unlock(): 뮤텍스 잠금해제
    - RWMutex: 읽기/쓰기 뮤텍스, 읽기와 쓰기 동작을 나누어서 잠금을 걸 수 있음. 쓰기 동작보다 읽기 동작이 많을 때 유리함
      - 읽기 락(Read Lock) : 읽기 락끼리는 서로를 막지 않음. 쓰기락은 막음
      - 쓰기 락(Write Lock) : 쓰기 시도중에 다른곳에서 이전값을 읽으면 안됨. 읽기 쓰기 락 모두 막음
      - sync.RWMutex
      - func (rw *RWMutex) Lock(), func (rw *RWMutex) Unlock() : 쓰기 뮤텍스 잠금, 잠금 해제
      - func (rw *RWMutex) RLock(), func (rw *RWMutex) RUnlock() : 읽기 뮤텍스 잠금, 잠금 해제
    - Cond: 조건 변수 condition variable 대기하고 있는 하나의 객체를 깨울 수도 있고 여러개를 동시에 깨울 수도 있음.
      - sync.Cond
      - func NewCond(l Locker) *Cond : 조건 변수 생성
      - func (c *Cond) Wait() : 고루틴 실행을 멈추고 대기
      - func (c *Cond) Signal() : 대기하고 있는 고루틴을 하나만 깨움
      - func (c *Cond) Broadcast() : 대기하고 있는 모든 고루틴을 깨움
    - Once : 특정 함수를 딱 한번만 실행. 반복문안에서 각종 초기화를 할 때 유용
      - sync.Once()
      - func (*Once) Do(f func())
    - Pool : 고루틴에서 사용할 수 있는 객체 풀, 자주 사용하는 객체를 풀에 보관했다가 다시 사용. 메모리 할당과 해제 횟수를 줄여 성능을 높일 때 사용
      - sync.Pool
      - func(p *Pool) Get( ) interface {} : 풀에 보관된 객체를 가져옴
      - func(p *Pool) Put(x interface{}) : 풀에 객체를 보관
    - WaitGroup : 고루틴이 모두 끝날때까지 기다리는 기능
      - sync.WaitGroup
      - func (wg *WaitGroup) Add(delta int) : 대기 그룹에 고루틴계수 추가
      - func (wg *WaitGroup) Done() : 고루틴이 끝났다는 걸 알려줄 때 사용
      - func (wg *WaitGroup) Wait() : 모든 고루틴이 끝날 때까지 기다림
    - Atomic : 원자적 연산이라고도 하며 더 이상 쪼갤 수 없는 연산, 고루틴, 멀티코어 환경에서 안전하게 값을 연산하는 기능. 여러 스레드, CPU코어에서 같은 변수를 수정할 때 서로 영향을 받지않고 안전하게 연산 가능 
  
- [2019-11-06] Unit 34 채널 

  - 고루틴끼리 데이터를 주고받고, 실행 흐름을 제어하는 기능. 모든 타입을 채널로 사용가능
채널 자체는 값이 아닌 레퍼런스 타입
    
  - 고루틴A(채널에 값을 보내고 꺼냄) <-> 채널 <-> 고루틴B(채널에 값을 보내고 꺼냄)
  
  - **make(chan 자료형)**
  
    - 채널을 사용하기 전에 반드시 **make**함수로 공간을 할당해야하고, 이렇게 생성하면 동기 채널(synchronous channel)이 생성
  
    - 채널을 매개변수로 받는 함수는 반드시 **go**키워드를 사용하여 고루틴으로 실행해야함.
  
    - 함수에서 채널을 매개변수로 받을 때는 **변수형 chan 자료형** 
  
    - 채널에 값을 보낼 때는 **채널<-값**, 가져올때는 **<-채널**
  
      
  
    ```go
    func sum(a int, b int, c chan int) {
        c <- a+b	//채널에 a와 b의 합을 보냄
    }
    func main() {
        c := make(chan int)		//int 형 채널 생성
        go sum(1,2,c)	//sum을 고루틴으로 실행, 채널을 매개변수로 넘겨줌
        
        n := <-c
        fmt.Println(n)	//채널에서 값을 꺼내와서 n에 대입
    }
    
    ```
  
  - **make(chan 자료형, 버퍼_개수)**
  
    - 채널에 버퍼를 1개 이상 설정하면 ==비동기 채널(asynchronous channel)== 생성
  
      비동기 채널은 보내는 쪽에서 버퍼가 가득 차면 실행을 멈추고 대기하며 받는 쪽에서는 버퍼에 값이 없으면 대기 
  
  - **range**키워드와 **close**함수
  
    - range는 채널에 값이 몇개나 들어올지 모르기 때문에 값이 들어올때마다 계속 꺼내기 위해 사용 
    - 이미 닫힌 채널에 값을 보내면 패닉이 발생 -> 채널을 닫으면 range 루프가 종료 -> 채널이 열려있고 값이 들어오지않는다면 range는 실행되지 않고 계속 대기. 만약 다른 곳에서 채널에 값을 보냈다면 그때부터 range가 계속 반복
  
    ```go
    	func main() {
    		c := make(chan int)
    		
    		go func() {
    			for i := 0;i < 5; i++ {
    				c <- i
    			}
    			close(c)
    		}()
    		
    		for i := range c {
    			fmt.Println(i)
    		}
    	}
    ```
  
  - 보내기 전용 채널과 받기 전용 채널
  
    - 보내기 전용 채널 **chan<- 자료형** : 값을 보낼 수만 있고 가져오려고 하면 컴파일 에러 발생
    - 받기 전용 채널 **<-chan 자료형** : range 키워드 또는 <-채널 형식으로 값을 꺼낼수만 있고 보내려고 하면 컴파일 에러 발생 
  
  - select 분기문 : 여러 채널을 손쉽게 사용
  
    - **select { case <- 채널: 코드 }**
    - switch 분기문과 비슷하지만 select 키워드 뒤에 검사할 변수를 따로 지정하지 않으며 각 채널에 값이 들어오면 해당 case가 실행
    - default 케이스를 지정할 수 있으며 case에 지정된 채널에 값이 들어오지 않았을 때 즉시 실행. 단 default에 적절한 처리를 하지않으면 CPU 코어를 모두 점유하기때문에 주의해야함.
  
- 

- [2019-11-05] Unit 33 고루틴 

  - 고루틴(Goroutine) : 함수를 동시에 실행시키는 기능 
go 함수명
    
  ```go
    func hello() {
        fmt.Println("Hello, world!")
    }
    
    func main() {
        go hello()		//함수를 고루틴으로 실행
    }
    ```
    
    
  
- [2019-11-04] Unit 31 ~ 32 구조체 임베딩 / 인터페이스

  - Go언어는 클래스, 상속을 제공하지 않음 -> 구조체에서 임베딩을 사용하면 상속과 같은 효과

  - 인터페이스는 메서드 집합. 단, 메서드 자체를 구현하지는 않음
    type 인터페이스명 interface {  }
  
  - 메서드를 가지는 인터페이스 정의
  
    type 인터페이스명 interface { 메서드 }
  
  
  
- [2019-11-04] UNIT29

  - Go언어에서 구조체 사용법 type 구조체명 struct { }

    ```go
    type Rect struct {
    	width int
    	height int
    }
    ```
    
  - Rect  타입으로 인스턴스 생성 var 변수명 구조체_타입
  
    ```go
    var rect Rect
    ```
  
  - 포인터에 메모리 공간 할당 구조체_포인터 = new(구조체_타입)
  
    ```go
    var rect1 *Rect
    rect1 = new(Rect)
    ```
  
    
  
- [2019-11-01] UNIT28

  - 포인터형 변수 var 변수명 *자료형

    ```go
    var numPtr *int
    ```
  
  - 빈 포인터형 변수 메모리 할당 포인터_변수 = new(자료형)
  
    ```go
    var numPtr *int = new(int)
    ```
  
  - 역참조(dereference) *포인터_변수명
  
    ```go
    *numPtr = 1
    ```
  
  - 일반 변수에 참조를 사용하면 포인터형 변수에 대입가능 &변수명
  
    ```go
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

    

    
