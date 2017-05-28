>goroutine是Go并行设计的核心。goroutine说到底其实就是协程,但是它比线程更小,十几个goroutine可能体现在底层就是五六个线程,Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB),当然会根据相应的数据伸缩。也正因为如此,可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。

goroutine是go runtime管理的一个**线程管理器**。  
通过`go`关键字实现goroutine。通过`go`关键字 后跟函数(必须跟函数)，就启动了一个goroutine。

```go
go hello(a,b,c)
```

```go
package main
import(
    "fmt"
    "runtime"
)

func say(msg string) {
    for i:=0; i<5; i++ {
        runtime.Gosched()
        fmt.Println(msg)
        // time.sleep(1*time.Second)
    }
}

func main() {
    go say("world") //开一个新的Goroutines执行
    say("hello") //当前Goroutines执行
}
//每次输出的结果不同，但是hello每次都可以保证输出5次，不论顺序。而world极端情况没有输出,是因为在另外一个goroutines中输出。  
```
>可以看到go关键字很方便的就实现了并发编程。 上面的多个goroutine运行在同一个进程里面,共享内存数据,不过设计上我们要遵循:不要通过共享来通信,而要通过通信来共享。

## go routine
goroutine运行在相同的地址空间,因此访问共享内存必须做好同步。那么goroutine之间如何进行数据的通 信呢,Go提供了一个很好的通信机制channel。channel可以与Unix shell 中的双向管道做类比:可以通过 它发送或者接收值。这些值只能是特定的类型:channel类型。定义一个channel时,也需要定义发送到 channel的值的类型。注意,**必须使用make创建channel**

```go
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface)
```
channel通过<-接收和发送数据，接收和发送取决于chan的的位置
ch <- v //发送v到channel ch
v := <-ch // 从ch中接收数据，并赋值给v


```go
package main

import "fmt"

func sum(a []int, ch chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	ch <- sum
}

func main() {
	ch := make(chan int)
	s := []int{1, -4, -9, 20, 4, 9}
	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)
	//x, y := <-ch, ch
	x, y := <-ch, <-ch
	fmt.Printf("(%d)+(%d)=%d\n", x, y, x+y)
}
```
>默认情况下,channel接收和发送数据都是阻塞的,除非另一端已经准备好,这样就使得Goroutines同步变的更加的简单,而不需要显式的lock。所谓阻塞,也就是如果读取(value := <-ch)它将会被阻塞,直到有数据接收。其次,任何发送(ch<-5)将会被阻塞,直到数据被读出。无缓冲channel是在多个goroutine之 间同步很棒的工具。



```go
package main
import "fmt"
func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2

    fmt.Println(<- ch)
    fmt.Println(<- ch)
}

```
将buff值修改为0，1都会
```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/Users/fqc/work/src/run.go:8 +0x9c
```


使用range读取channel数据，for range channel专门读取channel的数据，无需关注索引。

```go
package main

import "fmt"

func fibonacci(n int, ch chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	ch := make(chan int,10)
	go fibonacci(cap(ch), ch)
	for  v := range ch {
		fmt.Println(v)
	}
}

```

>for i := range c能够不断的读取channel里面的数据,直到该channel被显式的关闭。上面代码我们看 到可以显式的关闭channel,生产者通过内置函数 close 关闭channel。关闭channel之后就无法再发送任 何数据了,在消费方可以通过语法 v, ok := <-ch 测试channel是否被关闭。如果ok返回false,那么说channel已经没有任何数据并且已经被关闭。
>记住应该在生产者的地方关闭channel,而不是消费的地方去关闭它,这样容易引起panic,另外记住一点的就是channel不像文件之类的,不需要经常去关闭,只有当你确实没有任何发送数据了,或者你想显式的结束range循环之类的
如果去除close(ch)，也将会dead lock
----------------
reference   
《build web app with go》