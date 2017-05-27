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








