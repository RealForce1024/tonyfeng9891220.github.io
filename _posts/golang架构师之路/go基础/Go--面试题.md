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

## 边界检查

```go
package main

import "fmt"

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	//greatest := max(4, 7, 9, 123, 543, 23, 435, 53, 125)
	//greatest := max(-1, -2) //如果都是负数将是错误的
	greatest := max2(-1, -2)
	fmt.Println(greatest)
}
func max2(numbers ...int) int {
	var largest int
	for i, v := range numbers {
		if v > largest || i == 0 {
			largest = v
		}
	}
	return largest
}

/*
FYI
For your code to also work with only negative numbers such as

greatest := max(-200 -700)

include this as your range statement
for i, v := range numbers {
	if v > largest || i == 0 {
		largest = v
	}
}

What does that code do?

The first time through the range loop
the index, i, will be zero
so largest will be set to the first number

Originally largest is set to the zero value for an int, which is zero

Zero would be greater than any negative number

if you only have negative numbers
you need largest to be something less than zero

Thanks to Ricardo G for this code improvement!
*/

```

## slice容量
问题:我们知道切片的容量将随着随着长度的增大而自动扩容，那么下面的代码能否正常运行?如果不行，该如何解决?
```go
package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting[3] = "suprabadham"

	fmt.Println(greeting[2])
}

```
## 可变参数的使用

`append（slice []Type,elem ...Type）`

- `...`符号可以解引用slice为可变参数
- 添加类型必须是一致的
- slice是开闭区间

```go
package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 4, 5}
	myOtherSlice := []int{6, 7, 8, 9}
	otherSlice := []int{11,12,13}
	//otherStringSlice:=[]string{"hello","world"}
	mySlice = append(mySlice, myOtherSlice...)
	mySlice = append(mySlice,otherSlice...)
	//mySlice = append(mySlice,otherStringSlice...)

	fmt.Println(mySlice)
}

```
## 去掉周三

```go
package main

import "fmt"

func main() {

	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)

}
```

## 


