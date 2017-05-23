<!-- TOC -->

    - [go数据结构总览](#go数据结构总览)
- [结构体](#结构体)
    - [结构体声明与初始化](#结构体声明与初始化)
    - [访问成员变量](#访问成员变量)
        - [案例](#案例)
        - [结构体字面值](#结构体字面值)
        - [禁止企图使用未导出的成员](#禁止企图使用未导出的成员)
        - [结构体作为函数参数或返回值](#结构体作为函数参数或返回值)
        - [go中的所有参数都是值拷贝](#go中的所有参数都是值拷贝)
        - [结构体的比较](#结构体的比较)
        - [结构体可比较，可以作为map类型的key](#结构体可比较可以作为map类型的key)
        - [结构体嵌入和匿名成员](#结构体嵌入和匿名成员)
        - [匿名嵌入(匿名成员解决访问繁琐问题)](#匿名嵌入匿名成员解决访问繁琐问题)
        - [结构体字面值无法使用匿名嵌套的便利](#结构体字面值无法使用匿名嵌套的便利)

<!-- /TOC -->

## go数据结构总览
- 基础类型
  - 数字
  - 字符串
  - 布尔
- 复合类型(通过组合简单的基础类型组成复杂数据结构)
  - 数组
  - 结构体
- 引用类型(也属于复合类型，但是变量或状态的间接引用)
  - pointer
  - slice
  - map
  - function
  - channel
- 接口类型

# 结构体
结构体是聚合类型，多种类型的聚集，但不能包括自身结构体类型(指向自身的指针类型可以)。  
## 结构体声明与初始化
`type structName struct {fileName Tye}` 格式
结构体变量和其成员变量的声明  

```go
package main

import (
	"time"
	"fmt"
)

type Employee struct {
	Id        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    float64
	ManagerID int
}

func main() {
	var john Employee
	john.Salary += 5000
	fmt.Println(john.Salary)
}
```

## 访问成员变量
结构体变量点操作符`.`可以访问成员变量
指向结构体**成员变量的指针**通过点操作符`.`也可以访问成员变量

```go
john.Salary -= 5000
position := &john.Position // position *string := &john.Position
*position = "senior" + *position
```

指向结构体变量的指针也可以通过点操作符操作结构体变量。  

```go
var employee *Employee = &john
employee.position += "high"  //等价于(*employee).position += "high" 
```

### 案例  
```go
func main() {
	fmt.Printf("%v\n", EmployeeById(john.Id).Name) //作为参数，之后变量都可以

	id := john.Id
	EmployeeById(id).Position = "suzhou" //如果返回的是结构体类型而不是指针将会编译错误，赋值左边必须可以确定是变量，调用函数返回的是值，并不是一个可取地址的变量

    //如果返回的是结构体类型，可以通过赋值给新的变量来解决
    //john2 := EmployeeById(id)
	//john2.Position = "suzhou"
	//fmt.Printf("%v\n", john2.Position)
}

var john Employee = Employee{
	Id:       1000,
	Name:     "kobe",
	Salary:   10000,
	Position: "beijing",
}

/*func EmployeeById(id int) Employee {
	if id == 1000 {
		return john
	}
	return Employee{}
}*/

func EmployeeById(id int) *Employee {
	if id == 1000 {
		return &john
	}
	return &Employee{}
}
```

### 结构体字面值
两种方式:  
1. 按照成员顺序赋值 (简洁，但是需要记住成员顺序，但是一旦成员顺序调整，就会编译错误，需要相应的修改) 使用场景一般在比较有排列比较规则的结构上，比如坐标`image.Point{X,y}`或枚举`color.RGBA(red,green,blue,alpha)`上。  
2. 成员名:成员值 (优选，与顺序无关，可以包含部分或全部的成员，未显示声明的成员值为其类型的默认值)
注意:1、2两种结构体字面值方式不能混用

```go
package main

import "fmt"

func main() {

	type Point struct {
		X, Y int
	}

	p := Point{1, 2}

	type Student struct {
		Id      int
		Name    string
		Address string
	}

	stu := Student{
		Id:      100,
		Name:    "kobe",
		Address: "beijing",
	}

	fmt.Printf("%v\n%v\n", p, stu)
}

// {1 2}
// {100 kobe beijing}
```

### 禁止企图使用未导出的成员
```go
package p 
type T struct {
    a,b int // a and b are not exported
}

package q
import "p"
m := T{a:1,b:2} //compile error :cannot reference a,b
n := T{1,2}     //compiler error :cannot reference a,b
```

### 结构体作为函数参数或返回值
传入结构和传入结构体的指针作为参数

```go
package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func scale1(p Point, factor int) Point {
	return Point{p.x * factor, p.y * factor}
}

func main() {
	fmt.Println(scale1(Point{1, 2}, 5))  // 5,10
	pp := &Point{1, 2}
	scale2(pp, 5) // 5,10
	fmt.Printf("%v\n",pp ) //&{5,10}
	fmt.Println(pp.x) //5
}

func scale2(p *Point, factor int) {
	p.x = p.x * factor
	p.y = p.y * factor
}
```


将结构体指针作为参数，可以修改底层值。如果考虑效率的话，较大的结构体，通常优先考虑使用指针作为参数。  
```go
package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func scale1(p Point, factor int)  {
	p.X = p.X*factor
	p.Y = p.Y*factor
}

func main() {
	p := Point{1,2}
	scale1(p,5)
	fmt.Printf("%v\n", p) //{1,2}

	pp := &Point{1, 2}
	scale2(pp, 5) // 5,10
	fmt.Printf("%v\n",pp ) //&{5,10}
	fmt.Println(pp.X) //5
}

func scale2(p *Point, factor int) {
	p.X = p.X * factor
	p.Y = p.Y * factor
}
```

### go中的所有参数都是值拷贝
> 如果要在函数体内部修改结构体成员的话，使用指针是必须的。因为在go中，所有的函数参数都是值拷贝传入的，函数参数将不再是函数调用时的原始变量
因此我们在上个案例中看到 `scale1(p,5)` 并没有将结构体的成员扩大5倍，原因是将结构体本身地址的副本传入。  
通常通过指针处理结构体，可以通过以下两种方式创建并初始化结构体指针
- 一  `pp := &{1,2}` 
- 二
 ```go
 pp := new(Point)
 *pp = Point{1,2} 
 ```
&{1,2}可以写在表达式中较为常用，比如函数调用中。

### 结构体的比较
> 如果结构体的全部成员是可以比较的，那么结构体也是可以比较的，那么该两结构体是可以通过`==`或`!=`运算符进行比较的。

```go
p1 := Point{1,2}
p2 := Point{3,4}
p3 := Point{1,2}
fmt.Printf("%t\t%t\n", p1 == p2, p1 == p3) //false true
fmt.Printf("%t\t",p1 > p2) //compile error > not definied on Point
```
> 相等比较运算符==将比较两个结构体的每个成员，因此下面的两个比较的表达式是等价的:

```go
fmt.Println(p1.X==p2.X&&p1.Y==p2.Y) //false
fmt.Println(p1==p2) //false
```
### 结构体可比较，可以作为map类型的key
map类型的key必须是可比较的类型，而结构体满足此条件，因此可以作为map类型的key

```go
type Address struct {
	hostname string
	port int
}

var hits = make(map[Address]int)

hits = map[Address]int{
	Address{"gomaster.me",80}:1,
}

hits[Address{"gomaster.me",80}]++
```

### 结构体嵌入和匿名成员
go语言有种结构体嵌入机制，一个命名的结构体包含另一个结构体类型的匿名成员。可以简化x.d.e.f为x.f，大大便利了嵌套访问的问题。  

```go
type Circle struct {
	X, Y, Radius int
}

type Wheel struct {
	X, Y, Radius, Spokes int// X,Y=>Point    X,Y,Radius =>Circle
}

var w Wheel
w.X = 8
w.Y = 8
w.Radius = 5
w.Spokes = 20
fmt.Println(w)//{8 8 5 20}
```

我们发现X,Y可以抽象为坐标点Point，重构之
```go
type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
    Circle Circle
	Spokes int
}
```

改动后结构体类型变得更加清晰，但按照之前的访问方式变得繁琐起来。

```go
var w Wheel
w.Circle.Center.X = 8
w.Circle.Center.Y = 8
w.Circle.Radius =5
w.Spokes = 20
fmt.Println(w) //{{{8 8} 5} 20}
```
该如何简便的访问呢?

### 匿名嵌入(匿名成员解决访问繁琐问题)
go的一个重要特性是匿名嵌入特性(只声明成员类型，并不指定名称)。
>匿名类型的数据类型必须是命名的类型或指向一个命名的类型的指针。
我们将上面繁琐的访问重构后，Circle和Wheel各有一个匿名成员，Point类型被嵌入到了Circle结构体，Circle类型被嵌入到了Wheel结构体。 
因为匿名嵌入的特性，我们可以直接访问叶子路径而不需要给出完整路径。但即使匿名嵌入，也还是可以通过完整路径访问的，也就是说匿名类型还是有名字的，只是和命名类型相同而已，这些名字在点操作符是可选的。在访问任何匿名成员时，名字是可以省略的。   

```go
type Point struct {
	X, Y int
}

type Circle struct {
	Point	
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

var w Wheel
w.X = 8
w.Y = 8
w.Radius = 5 
w.Spokes = 20
fmt.Println(w) // {{{8 8} 5} 20}
```

### 结构体字面值无法使用匿名嵌套的便利
结构体字面值并没有简短表示匿名成员的语法。  
```go
w = Wheel{8,8,5,20} // compile error: unknown fields
w = Wheel{X:8,Y:8,Radius:5,Spokes:20} //compile error: unknown fields
```

结构体字面值必须遵循类型声明时的结构。  

```go
w1 := Wheel{Circle{Point{8, 8}, 5}, 20}

w2 := Wheel{
	Circle: Circle{
		Point:  Point{8, 8},
		Radius: 5,
	},
	Spokes: 20,
}

fmt.Printf("%v\n%#v\n", w1, w2)
//{{{8 8} 5} 20}	
//main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}
```

注意: `%v`参数 `#`副词可以用来修饰参数 `%#v`表示使用和go语言类似的语法打印值。对于结构体类型来说包含每一个成员的名字

由于嵌套成员有个隐式的成员名字，因此不能有相同类型的嵌套成员，否则将会导致名字冲突。
成员的名字是由其类型隐式决定的，而导出性也是由类型的大小写决定，即所有匿名成员也有可见性的规则约束。但是在包含的结构体访问小写的嵌套成员也是可以访问的。但是包外访问嵌套的不导出成员则不行。  

我们看到被嵌入的类型都是结构体类型，但除了命名的结构体类型，其他被命名的类型都可以当做结构体的匿名成员。  

但为什么要嵌入没有任何子成员类型的匿名成员类型呢？
原因是匿名类型的方法集。  
这里不得不提点运算符，不仅仅可以选择匿名成员嵌套的成员，也可以用于访问匿名成员的方法。外层结构体不仅仅获得了匿名嵌套类型的**所有成员**，也获得了该类型**导出的全部方法**。这个机制可以用于将一个有简单行为的对象组合成又复杂行为的对象。组合是go面向对象编程的核心。  
































