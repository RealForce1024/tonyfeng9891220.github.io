# 信号量

## 让当前线程等待任务线程执行完毕

```go
package main

import (
	"fmt"
)

//var exit chan struct{} //引发panic: close of nil channel

var exit = make(chan byte)
func main() {
	
	go func() {
		fmt.Println("hello world")
		close(exit)
	}()
	<-exit
}
```

另外一种最简单也是最不稳定的是使用time.Sleep主观时间控制，在生产中不建议使用，因为go并发任务是由调度器来调度执行的，时间并不受控。

## 等待多个任务执行完毕


```go
package main

import (
	"fmt"
)

var exit = make(chan int)

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("hello world:", i)			<-exit //阻塞
		}()
	}
	close(exit)//发起信号，golang在关闭信道之前，会完成接收任务。
}
```
上述代码表面上看似可以解决，实际上有致命的错误
1. i将会产生同步问题
2. 打印执行次数不固定 也是因为主goroutine执行完毕，并未等待子任务执行完。可以将exit等去掉，使用time.Sleep验证此想法

```go
for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("hello world:")
		}()
	}
time.Sleep(time.Second)
```

所以需要使用其他方案

```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("hello world:", i)
			wg.Done()
		}()
	}
	wg.Wait()
}
```

这种方案可以解决等待多个任务完成的问题了，但是i的值我们看到还是各种不稳定，设置重复，我们需要通过goroutine的延迟执行，立即计算并复制参数的特性，记录下来i的值，以免被计算。

所以，必须使用传参的方式

```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mux sync.Mutex

func main() {
	mux.Lock()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("groutine-%d done\n",id)
		}(i)
	}
	wg.Wait()//等待归零，解除阻塞
}
```

## 多处使用Wait，都能接收到通知

```go
package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go func() {
		wg.Wait()
		fmt.Println("wait exit")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("done")
	}()
	wg.Wait()
	fmt.Println("main exit")
}

//done
//wait exit
//main exit
```

# 将任意函数并发执行

```go
package main

import (
	"fmt"
	"go_commons"
	"runtime"
	"sync"
)

func main() {
	defer go_commons.TraceTime()()
	//count()
	n := runtime.NumCPU()
	//test(n) //13.x ms
	//test2(n)
	test3(n, count)
}
func test2(n int) {
	/*
	go func() {
		for i := 0; i < n; i++ {
			count()
		}
	}()
	*/

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count()
		}()
	}
	wg.Wait()
}

func test3(core int, fn func()) {
	var wg sync.WaitGroup
	for i := 0; i < core; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//count()
			fn()
		}()
	}
	wg.Wait()
}

func test(n int) {
	for i := 0; i < n; i++ {
		count()
	}
}

func count() {
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum += i
	}
	fmt.Println(sum)
}

```

# go channel的正确使用姿势

```go
package main

import (
	"fmt"
	"time"
)

type Addr struct {
	City, District string
}

type Person struct {
	Name string
	Age  int
	Addr
}

type PersonHandler interface {
	Batch(<-chan Person) <-chan Person
	Handle(person *Person)
}

type PersonHandlerImpl struct{}

func (handler PersonHandlerImpl) Batch(origs <-chan Person) <-chan Person {
	//dest := make(<-chan Person, 100)
	dest := make(chan Person, 100)
	go func() {
		for p := range origs {
			handler.Handle(&p)
			dest <- p
		}
		fmt.Println("all people has been handled")
		close(dest)
	}()
	return dest
}

func (handler PersonHandlerImpl) Handle(person *Person) {
	if person.District == "haidian" {
		person.District = "changping"
	}
}

var persons []Person
var personTotal int = 200

func init() {
	for i := 0; i < personTotal; i++ {
		persons = append(persons, Person{Name: fmt.Sprintf("P%v", i), Age: 28, Addr: Addr{City: "bj", District: "haidian"}})
	}
}

func getHandler() PersonHandler {
	return PersonHandlerImpl{}
}
func main() {
	defer trace()()
	time.Sleep(time.Second)
	handler := getHandler()
	origs := make(chan Person, 100)
	fetchPersons(origs)
	dest := handler.Batch(origs)
	sign := save(dest)
	<-sign
}
func trace() func() {
	now:=time.Now()
	return func() {
		duration := time.Since(now)
		fmt.Println("运行时间:", duration)
	}
}

// fectch 将向通道中写入
func fetchPersons(origs chan<- Person) {
	go func() {
		for _, p := range persons {
			origs <- p
		}
		fmt.Println("all person has been fetched.")
		close(origs)
	}()
}

// save 从通道中读取
func save(dest <-chan Person) <-chan byte {
	sign := make(chan byte, 1)

	go func() {
		/*for p,ok := range dest {
			if !ok {
				break
			}
			savePerson(p)
		}*/

		for {
			p, ok := <-dest
			if !ok {
				fmt.Println("all people saved")
				sign <- 0
				break
			}
			savePerson(p)
		}
	}()

	//sign <- 1 //这里会同步执行
	return sign
}

func savePerson(p Person) bool {
	fmt.Println(p)
	return true
}
```

