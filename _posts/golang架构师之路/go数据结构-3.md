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
1. 按照成员顺序赋值 (简洁，但是需要记住成员顺序，但是一旦成员顺序调整，就会编译错误，需要相应的修改) 使用场景一般在比较有排列比较规则的结构上，比如坐标`image.Point{x,y}`或枚举`color.RGBA(red,green,blue,alpha)`上。  
2. 成员名:成员值 (优选，与顺序无关，可以包含部分或全部的成员，未显示声明的成员值为其类型的默认值)
注意:1、2两种结构体字面值方式不能混用

```go
package main

import "fmt"

func main() {

	type Point struct {
		x, y int
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
	x, y int
}

func scale1(p Point, factor int)  {
	p.x = p.x*factor
	p.y = p.y*factor
}

func main() {
	p := Point{1,2}
	scale1(p,5)
	fmt.Printf("%v\n", p) //{1,2}

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
如果结构体的全部成员是可以比较的，那么结构体也是可以比较的。  

```go
p1 := Point{1,2}
p2 := Point{3,4}
p3 := Point{1, 2}
fmt.Printf("%t\t%t\n", p1 == p2, p1 == p3) //false true
fmt.Printf("%t\t",p1 > p2) //compile error > not definied on Point
```




























