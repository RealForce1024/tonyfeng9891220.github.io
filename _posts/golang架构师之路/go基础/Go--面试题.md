1、写出下面代码输出内容。

```go
package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
```
案例分析:考察golang的异常处理机制和defer机制。panic触发异常，而defer压栈延迟执行，最终panic会将异常向上抛出。
执行结果:

```go
打印后
打印中
打印前
panic: 触发异常

......堆栈异常信息
```

延伸阅读:

- [关于golang的panic recover异常错误处理](http://xiaorui.cc/2016/03/09/%E5%85%B3%E4%BA%8Egolang%E7%9A%84panic-recover%E5%BC%82%E5%B8%B8%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86/)
- [Go语言中使用Defer几个场景](http://developer.51cto.com/art/201306/400489.htm)


2、请指出下面代码的问题并说明原因

```go
type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
}
```
案例解析: go的迭代变量会复用地址，也就说`stu`取地址`&stu`始终是一样的地址，所以循环体中的代码问题在于每个m存储的student实例地址是一样的。 解决方案可以将map中的指针声明修改为结构体类型。

3、下面的代码会输出什么，并说明原因

```go
func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```


