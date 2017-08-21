
## 并发 concurrency
### 数据结构Channel
- Channel是Go程序(Goroutine)的一种高级数据结构。它可作为不同Goroutine之间的桥梁即数据传输通道，在通道内传递的指定类型消息，我们将其称为通道的类型化数据(或类型元素)。  
- Channel定义
通过描述，我们自然能想到一个Channel的定义:  
`chan T`   
chan关键字代表通道 T代表通道内传递数据的类型  
- 在通道内消息的传递是安全的。  
- Goroutine代表Go并发程序，由Go runtime运行时系统调度，依托于内核线程并发执行代码。  
- 通道的初始化
通道类型比较特殊，没有字面值，只能使用make函数初始化构造`值`。  

`make(chan int, 5)`   
第一个参数代表 元素类型为int的通道
第二个参数代表 通道的长度为5，实际代表缓存的长度为5
也就是说上面的通道创建含义为构建一个缓存大小为5，元素类型为int的`通道值`。  

注意:   
  这里着重强调了`值`的概念。当我们在编译的时候，单独声明`chan int`或`make(chan int, 5)`，编译器都无法编译通过。 尤其是在`make(chan int, 5)`编译提示 `make(chan int, 5), evaluated but not used`。这也印证了Go的哲学，去除无用。

通道中缓存即通道暂存的数据为先进先出结构，针对于通道值而言，越早被放入(发送)到通道的元素越早被放出通道(接收)。  
`send -> channel -> receiver`     

- `<-`
符号`<-`为向通道中发送(write/send)或接收(read/receive)数据

声明一个通道并向该通道发送值，然后从该通道中取出值。  

```Go
ch := make(chan int, 1)
ch <- 99
fmt.Println(<-ch) // 99
```

读取通道值可以返回两个值，`val`代表读取的通道元素值，`ok`代表通道的状态。

```Go
ch := make(chan int, 1)
ch <- 99
val, ok := <-ch
if !ok {
	fmt.Println("channel has been closed")
	return
}
fmt.Println(val)
```

第二参数`ok`为bool类型，代表通道值是否有效或通道已关闭。在接收之前或过程中，通道值被关闭了，接收或写操作立即结束并返回一个通道元素类型的零值(这里通道元素类型为int，因此零值则为0)，那么零值容易混淆，我们并不知道0是否正常返回，因此有了第二个布尔返回值代表通道值的状态，我们对于判断就心里有底了。  

关闭通道可以通过程序自动终止或者通过内置函数`close(c chan <- Type)`手动关闭。

关于channel的几点注意事项:
- 通道有效的前提下，直至通道被填充满会阻塞(被放入缓存的数据等于通道长度)，否则都为异步
- 有效通道，接收通道值会在其已空时(没有缓存数据)阻塞
- 在向关闭的通道发送数据将会引起恐慌
- 重复关闭通道会引起恐慌
- 通道类型为引用类型，零值为nil

```Go
package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 99 //非缓冲通道，发送数据会被阻塞，下面的取数据则造成死锁
	val, ok := <-ch // 接收方会一直等到有数据来
	if !ok {
		fmt.Println("channel has been closed")
		return
	}
	fmt.Println(val)
}
// fatal error: all Goroutines are asleep - deadlock!
//
// Goroutine 1 [chan send]:
// main.main()
// 	/Users/fqc/work/src/run.Go:7 +0x7a
```

```Go
package main

import "fmt"

func main() {
	ch := make(chan int,1)
	ch <- 99 //带缓存的通道，直至缓存填满为同步，否则为异步
	val, ok := <-ch // 接收方可以随时取数据，直至缓存数据为空时阻塞
	if !ok {
		fmt.Println("channel has been closed")
		return
	}
	fmt.Println(val)
}
```

总结:  
非缓存通道，发送方在向通道值发送数据的时候会立即阻塞，直到接收方来消费数据。 `make(chan T, n)`，其中n>0
缓存通道，发送方会立即拷贝数据到缓存直到等于缓存长度进入阻塞，接收方可随时取数据，直到缓存数据为空时阻塞。 `make(chan T,0)`，其中n=0

## 单向通道与多项通道  
除了按照有无缓存划分通道的种类，还可以通过通道的方向划分为单向通道和双向通道，而双向通道是默认的。  
单向通道即数据只能按照一个方向进行传输，按照发送者和接受者的数据流方向的不同，可以分为接收通道和发送通道。  
```Go
type Sender chan<- int // 发送者通道
type Receiver <-chan int // 接受者通道
```
注意
1. 发送者通道 `chan<-` 箭头**指向**通道
2. 接受者通道 `<-chan` 箭头**来自**通道
类型Receiver,Sender代表接收/发送通道类型，chan关键字后跟随的箭头符号代表了数据的流向

###
```Go
package main
import (
	"fmt"
	"time"
)

func main() {
	Go Run()
	time.Sleep(2 * time.Second)
}

func Run() {
	fmt.Println("Go concurrency")
}
```

使用匿名函数
```Go
package main

import (
	"fmt"
	"time"
)

func main() {
	Go func() {
		fmt.Println("hello concurrency")
	}()
	time.Sleep(2 * time.Second)
}

func Run() {
	fmt.Println("Go concurrency")
}
```

### 更优雅的通信，而非不靠谱的线程睡眠
Go可以使用通信机制解决共享内存带来的苦恼。    
- Channel是Goroutine通信的桥梁，大都是阻塞同步的   
- 使用make创建，close关闭
- 可使用for range 对Channel进行迭代不断操作
- 引用类型
- 可以设置缓存大小，未填满前不会阻塞
- 可以设置单向、双向通道

channel读消息会阻塞同步的，直到通道中有消息写入，channel读取到才会继续执行。  

```Go
package main

import "fmt"

func main() {
	c := make(chan bool)
	Go func() {
		fmt.Println("Goroutine")
		c <- true // 存 发送消息
	}()
	<-c //取 接收消息 main执行到这里会阻塞，直到匿名函数中存入了true，channel读取到才会继续执行。
}

```

迭代chaanel时需要明确的正确的执行关闭了channel，否则会造成死锁
```Go
package main

import "fmt"

func main() {
	c := make(chan bool)
	Go func() {
		fmt.Println("Goroutine")
		c <- true
		close(c)
	}()
	for val := range c {
		fmt.Println(val)
	}
}

// Goroutine
// true
```

没有正确关闭channel导致死锁

```Go
...
//close(c)
...

// fatal error: all Goroutines are asleep - deadlock!
// Goroutine
//
// true
// Goroutine 1 [chan receive]:
// main.main()
// 	/Users/fqc/work/src/run.Go:12 +0x85
```

使用make函数都是双向通道channel，可存可取
单向通道，只能存或取，一般用在参数传递上，目的在于防止误读误写操作。  

未设置缓存大小时，缓存为零值，那么它将是阻塞同步的。设置缓存而没有存满的时候该通道为异步的，不会发生阻塞。  
有缓存和无缓存的区别  

```Go
package main

import "fmt"

func main() {
	c := make(chan bool)
	Go func() {
		fmt.Println("Goroutine")
		<-c
	}()
	c <- true
}
// Goroutine
```

`c := make(chan bool, 1)`  
将不会输出。   
```Go
c := make(chan bool)
Go func() {
	fmt.Println("Goroutine")
	<-c
}()
c <- true // 无缓存的时候是阻塞的，里面的内容需要被写完或消息被发送完。因此有缓存的时候还可以被写入
```

有缓存的时候，异步，都会向下执行
诀窍:首先认为是瀑布执行，然后分析关键点 1.看读写顺序 2.看有无缓存
阻塞或异步的时候，读始终需要让写先执行
阻塞的时候，读取的时候需要有消息，否则一直阻塞
异步的时候，读更快嘛...
有缓存的时候爱读不读，没缓存的时候需要等待写入玩 读出来，这和实际开发中联系 缓存可以不去读，但无缓存需要强制去读
有缓存是异步的，无缓存是同步阻塞的。

如何让有缓存的时候，也想让其同步，该如何实现?

