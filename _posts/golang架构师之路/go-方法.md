面向对象:
使用方法表达对属性和对应行为的操作，无需直接去操作对象，而是使用方法来操作对象。

```go
const day = 24 * time.Hour
fmt.Println(day.Seconds()) // 86400
```

```go
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
```

go的面向对象两大特点: 封装 组合  

## 方法声明
在函数名之前增加一个变量声明，该参数会将该函数附加到该类型上，即相当于为该类型定义了新的独占方法。  
调用方法面向对象的理解就是向该对象发送消息，从该角度上看，对象即为接一个方法(消息)的接收器，go中习惯称之为receiver。
> 在Go语言中,我们并不会像其它语言那样用this或者self作为接收器; 我们可以任意的选择接收器的名字。由于接收器的名字经常会被使用到,所以保持其在方法间传递时的一致性和简 短性是不错的主意。这里的建议是可以使用其类型的第一个字母,比如这里使用了Point的首字母p。
方法的调用，和方法的声明一样，接收器在前，方法名在后。  
p.Distance表达式叫做选择器，p.X属性获取也是选择器，比较奇怪的是声明一个p的X()方法会编译不通过。  
```go
package main

import (
	"math"
	"fmt"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
}
```

## go可以为任意类型定义方法
Path类型是Path类型的slice
注意`if i > 0`的巧妙设计 每个Path代表一个线段的集合。
```go
type Path []Path
func (p Path) Distance() float64 {
    sum := 0
    for i := range p {
        if i > 0 {
            sum += p[i-1].Distance(p[i])
        }
    }
    return sum
}
```
不知不觉中，我们已经给`[]Path`添加了Distance方法。  
在go中，除了指针和interface外，我们可以很方便的给数值，字符串，slice，map等添加行为方法很方便，这也和其他语言很大的不同之处。  
编译器会根据方法的名字和接收器来决定调用哪一个函数。
不同的类型可以拥有同样的方法名，但同一类型不可有方法名的冲突。  
>方法比之函数的一些好处:方法名可以简短。当我们在包外调用的时候这种好处就会被放大,因为我们可以使用这个短名字,而可以省略掉包的名字,下面是例子:

```go
import "geometry"
perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
fmt.Println(geometry.PathDistance(perim)) // "12", standalone function
fmt.Println(perim.Distance()) // "12", method of geometry.Path
```
在Go中调用包外的函数需要带上包名。

## 基于指针对象的方法
Go中调用一个函数时，会对每一个参数值进行拷贝，**如果一个函数需要更新一个变量，或者函数的其中的一个参数实在太大我们希望能够避免这种默认的拷贝，这种情况下就需要使用到指针了**。对应到更新接收器的对象的方法上来说，当这个接受者变量本身比较大时，我们就可以用其指针而不是对象来声明方法，如下:

```go
func (p *Point) ScaleBy(factor float64) {
    //p.X = p.X * factor
    p.X *= factor
    p.Y *= factor
}

fmt.Printf("%T\n", (*Point).ScaleBy)
// func(*main.Point, float64)
```

>在现实的程序里,一般会约定如果Point这个类有一个指针作为接收器的方法,那么所有Point 的方法都必须有一个指针接收器,即使是那些并不需要这个指针接收器的函数。

只有类型Point和指向它的指针(*Point)才是可能出现在接收器声明里的两种接收器。
不过为了避免歧义，在声明方法时，如果一个类型名本身是指针的话，是不允许出现在接收器中的。

```go
type P *int
func (P) f() {/**/} //compile error:invalid receiver type  
```

接收器为指针的方法调用有以下几种方式:
```go
r := &Point{1,2}
r.ScaleBy(2)
fmt.Println(*r)
```

```go
p := Point{1,2}
pptr := &p
pptr.ScaleBy(2)
fmt.Println(pptr) // &{2 4}
```

```go
p := Point{1, 2}
(&p).ScaleBy(2)
fmt.Println(p) // "{2, 4}"
```

> 不过后面两种方法有些笨拙。幸运的是,go语言本身在这种地方会帮到我们。如果接收器p是 一个Point类型的变量,并且其方法需要一个Point指针作为接收器,我们可以用下面这种简短 的写法:

```go
p.ScaleBy(2)
```

>编译器会隐式地帮我们用&p去调用ScaleBy这个方法。这种简写方法只适用于“变量”,包括 struct里的字段比如p.X,以及array和slice内的元素比如perim[0]。我们不能通过一个无法取到 地址的接收器来调用指针方法,比如临时变量的内存地址就无法获取得到:

```go
Point{1, 2}.ScaleBy(2) // compile error: can't take address of Point literal
```
>我们可以用一个 `*Point` 这样的接收器来调用Point的方法,因为我们可以通过地址来找到 这个变量,只要用解引用符号 * 来取到该变量即可。编译器在这里也会给我们隐式地插入*这个操作符,所以下面这两种写法等价的:
```go
pptr.Distance(q)
(*pptr).Distance(q)
```
总结:
>1. 不管你的method的receiver是指针类型还是非指针类型,都是可以通过指针/非指针类型进行调用的,编译器会帮你做类型转换。 
>2. 在声明一个method的receiver该是指针还是非指针类型时,你需要考虑两方面的内部,第一方面是这个对象本身是不是特别大,如果声明为非指针变量时,调用会产生一次拷 贝;第二方面是如果你用指针类型作为receiver,那么你一定要注意,这种指针类型指向 的始终是一块内存地址,就算你对其进行了拷贝。熟悉C或者C艹的人这里应该很快能明白。

```go
package main

import (
	"fmt"
)

type Point struct {
	X, Y float64
	s    []int  //看其变化
}

func (p Point) change() {
	p.s = []int{1, 2, 3}
}

func main() {
	p := Point{1, 2, []int{}}
	p.change()
	fmt.Println("after change() ", p)
}

// after change()  {1 2 []}
// 传递指针时则  after change()  {1 2 [1,2,3]}
```


## 方法继承
匿名成员可以使用简短的选择器，使得我们可以解决字段类型嵌套的繁琐访问问题，go圣经中解释的很到位了，但是更深一层的理解，匿名成员是成员的继承实现，当拥有了匿名成员，就相当于拥有了该类型内部的导出成员。
而方法也是如此，匿名成员了方法，那么包含该匿名字段的类型也可以调用该方法。  

```go
package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type student struct {
	Human
	school string
}
type employee struct {
	Human
	company string
}

func (h *Human) sayHi(msg string) {
	fmt.Printf("hello,my name is %s, %s\n", h.name, msg)
}

func main() {
	stu := student{Human{"kobe", 12, "1232323"}, "USC"}
	emp := employee{Human{"zhansan", 12, "23232"}, "KFC"}
	stu.sayHi("welcome")
	emp.sayHi("welcome to our company")
}
// hello,my name is kobe, welcome
// hello,my name is zhansan, welcome to our company
```

## 方法重写
```go
package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type student struct {
	Human
	school string
}
type employee struct {
	Human
	company string
}

func (h *Human) sayHi(msg string) {
	fmt.Printf("hello,my name is %s, %s\n", h.name, msg)
}

func (s *student) sayHi(msg string) {
	fmt.Printf("hello,my name is %s ,this is student impl %s\n ", s.name, msg)

}
func main() {
	stu := student{Human{"kobe", 12, "1232323"}, "USC"}
	emp := employee{Human{"zhansan", 12, "23232"}, "KFC"}
	stu.sayHi("welcome")
	emp.sayHi("welcome to our company")
}

```
go语言的面向对象的设计是如此精妙和简约。  





