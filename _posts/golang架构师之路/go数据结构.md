<!-- TOC -->

- [go数据结构总览](#go数据结构总览)
- [数组](#数组)
    - [一. 数组声明创建](#一-数组声明创建)
        - [1.声明类型，默认为零值](#1声明类型默认为零值)
        - [2.数组字面值](#2数组字面值)
        - [3. `...` 忽略数组长度定义](#3--忽略数组长度定义)
        - [4. 指定索引方式初始化](#4-指定索引方式初始化)
    - [5. 混合初始化指定值](#5-混合初始化指定值)
    - [二. 相同类型数组的比较](#二-相同类型数组的比较)
    - [低效的数组参数](#低效的数组参数)
    - [更高效的数组指针](#更高效的数组指针)
    - [反转slice和字符串](#反转slice和字符串)
- [切片 slice](#切片-slice)
    - [切片定义](#切片定义)
    - [切片和数组的零值](#切片和数组的零值)
    - [与多个切片共享底层数组](#与多个切片共享底层数组)
    - [反转数组的应用](#反转数组的应用)
    - [append函数](#append函数)
    - [slice之间不能比较](#slice之间不能比较)
    - [slice可以和nil比较(之一)](#slice可以和nil比较之一)
    - [优先使用len(s)==0判断slice是否为空](#优先使用lens0判断slice是否为空)
    - [模拟append函数，理解slice的扩容](#模拟append函数理解slice的扩容)
        - [容量越界](#容量越界)
        - [appendInt](#appendint)
    - [slice的内存使用技巧](#slice的内存使用技巧)
    - [使用slice模拟stack](#使用slice模拟stack)

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

本文将不包括基础类型

## 数组
### 一. 数组声明创建
#### 1.声明类型，默认为零值
数组类型: **`[n]Type`**  
`var arr [5]int` //创建长度为5、int型的数组。数组元素为int类型零值`0`。  
**数组的长度是数组类型的一部分**  

```go
var arr [5]int
fmt.Println(arr[0])
fmt.Println(arr[len(arr)-1])
for i, v := range arr {
  fmt.Printf("%d,%d\n", i, v)
}
```

#### 2.数组字面值

```go
// var arr [5]int = [5]int{1, 2, 3, 3, 4} 代码味道很差
var arr = [5]int{1, 2, 3, 4, 5}
fmt.Printf("%v", arr)
```

#### 3. `...` 忽略数组长度定义
```go
arr := [...]int{3,4}
fmt.Printf("%v", arr)
```

注意:`[...]int{n}`是数组字面值的一中变形，字面值书属于值范畴

```go
func change(arr [...]int)  { // 虽然没有语法错误，但是运行时引发了异常，原因是在数组字面值以外使用了[...]
	for _, value := range arr {
		value++
	}
	fmt.Printf("%v", arr)
}

// command-line-arguments
// src/run.go:25: use of [...] array outside of array literal
```
#### 4. 指定索引方式初始化
除了像前面两小节中的可以数组字面值直接顺序罗列的方式，也可以通过指定下标索引的方式指定值初始化，而数组的长度以最大索引为准，类型则以元素类型为准。未指定索引位置的值为元素类型零值。  

```go
arr := [...]int{1:3,2:4,5:9}
fmt.Printf("%v", arr) // [0 3 4 0 0 9]
```

这里涉及golang的iota和常量的使用，以及数组的索引初始化。(golang代码味道使用) 
```go
type Currency int

const (
  USD Currency = iota
  EUR
  GBP
  RMB
)

symbol := [...]string{USD:"$",EUR:"€",GBP:"£",RMB:"¥"}
fmt.Printf("index: %d, symbol: %s",RMB,symbol[RMB])
```
通过该案例的使用索引使用方式初始化数组，我们发现初始化的顺序无关紧要，索引值我们可以使用常量别名的方式，索引获取到我们需要的值。  

### 5. 混合初始化指定值
字面值和数组一样，可以顺序指定初始化值序列，也可以通过索引和元素值指定，或者两种风格的混合语法初始化

```go
func main() {
	arr := [...]int{1, 2, 5:10}
	info(arr)//1,2,0,0,0,10
}

func info(arr [6]int) {
	fmt.Printf("s=%v", arr)
}
```
注意:info()函数的参数类型[6]int这里是非常大的局限，今后我们可以使用切片slice替代

### 二. 相同类型数组的比较
- 元素类型可以比较的数组之间是可以比较的，不同类型的数组无法比较(编译无法通过)。

```go
arr1 := [3]int{1, 2, 3}
arr2 := [...]int{1, 2, 3}
arr3 := [3]int{1, 2, 4}
fmt.Printf("%t,%t,%t\n", arr1 == arr2, arr1 == arr3, arr2 == arr3)
arr4 := [2]int{1,2} //[2]int与之前的[3]int类型是不同的，所以无法比较，编译也无法通过
//fmt.Println(arr1==arr4) 无法编译通过
fmt.Printf("%v\n", arr4)
```
注意:`%t`打印布尔类型。  

- 实际应用:    
> crypto/sha256包的Sum256函数对一个任意的字节slice类型的数据生 成一个对应的消息摘要。消息摘要有256bit大小,因此对应[32]byte数组类型

```go
package main

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%v\n%v\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println("-------------------")
	fmt.Printf("%x\n%x\n%t\n%T", c1, c2, c1 == c2, c1)
}
```
注意:`%x`称为fmt.Printf()函数的副词参数，其以16进制的格式打印数组或slice的全部的元素。`%t`副词参数用于打印布尔类型的数据，`%T`副词参数用于显示一个值对应的数值类型。    

### 低效的数组参数
由于go的参数调用都是传递的副本，所以对于数组参数，则是传递的数组副本，对于参数的修改只能体现在数组副本上，而不是原先的传入数组。并且传递数组的副本相当低效。

```go
func main() {
	//change([...]int{})
	change([5]int{})
}

func change(arr [5]int)  {
	for _, value := range arr {
		value++
	}
	fmt.Printf("%v", arr)
}
```
### 更高效的数组指针
我们可以通过传递数组的指针的方式，来提高数组参数的效率。    
注意for range中iterm是副本。  

```go
func main() {
	//change([...]int{})
	change(&[5]int{})
}

func change(arr *[5]int) {
	for _, value := range arr { //Oops! item is only a copy of the slice element.
		//value++//此时是只读的，修改是无效的
		value++
	}

	for i := range arr {
		arr[i]++ // 指针数组对应的底层值
	}
	fmt.Printf("%v", arr)
}
```
[go官方github range详解](https://github.com/golang/go/wiki/Range)	  
注：数组虽然高效，但是长度是固定的，没有任何添加或删除数组元素的方法。所以一般除了特殊的场景使用数组外，通常都首选slice切片。  


### 反转slice和字符串
- 反正字符串  

```go
func main() {
	s := []string{"a", "b", "c"}
	myReverse(s)
	fmt.Println(s)
}
func myReverse(s []string) {
	for i := 0; i < int(len(s)/2); i++ {
		j := len(s) - i - 1
		fmt.Println(i, "<=>", j)
		s[i], s[j] = s[j], s[i]
	}
}
```

不同的实现

```go
func main() {
	sli := []string{"hello","world","ok"}
	Reverse(sli)
	fmt.Printf("%v", sli)
}
func Reverse(str []string)  {
	//for i, j := 0, len(str)-1; i < len(str)/2; i, j = i+1, j-1 {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}
```

- 反转字符串  

```go
package main

import "fmt"

func main() {
	fmt.Println(Reverse("hello"))
}
func Reverse(s string) string {
	str := []rune(s)
	//for i, j := 0, len(str)-1; i < len(str)/2; i, j = i+1, j-1 {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}
```

## 切片 slice
切片是动态的数组。其由指针，长度，容量三个属性组成，底层指向数组，因此切片可以指向数组的一部分或全部元素。  
切片长度不超过容量，`len(slice) <= cap(slice)`  

如果切片超出容量则会引发panic，而超出长度则是拓展切片，拓展的切片为新的切片(底层一般会针对底层数组追加新元素进行拓展，所以尽量在一开始确定好合适的容量大小是非常好的，否则有些性能影响)

### 切片定义
切片的定义至少了长度的显示定义，但是切片会隐式定义合适长度的数组。   

```go
var months = []string{1:"January",2:"Febuary"/*....*/12:"December"}
```
注:中间月份代码有所省略
字面值和数组一样，可以顺序指定初始化值序列，也可以通过索引和元素值指定，或者两种风格的混合语法初始化
```go
sli := []int{1,2,5:10}
info(sli) //1,2,0,0,10

func info(s []int) {
	fmt.Printf("len=%v,cap=%v,slice=%v\n", len(s), cap(s),s)
}
//len=6,cap=6,slice=[1 2 0 0 0 10]
```
- 使用内置函数make定义初始化底层数组，并返回视图slice
`make(slice,len,[cap])`

```go
make([]T,len)
make([]T,len,cap) //same as make([]T,cap)[:len]
```

### 切片和数组的零值  
非常重要的是，切片和数组都是**用花括弧包含一系列的初始化元素**，其他语言的使用者使用go时，可能会误以为只是声明了变量类型，但没有初始化，尤其是[]int{} 这种类型的切片值为`[]`，[...]int{}数组值为`[]`
```go
var arr []int
var slice []int
abb := [...]int{}
acc := []int{}
fmt.Printf("%v,%v,%v,%v", arr, slice,abb,acc)
```

### 与多个切片共享底层数组
```go
var Q2 = months[4:7]
var summer = months[6:9]
```
months[6]被切片Q2和summer共同指向   

### 反转数组的应用
可以用于任意长度的slice  
> 一种将slice元素循环向左旋转n个元素的方法是三次调用reverse反转函数,第一次是反转开头 的n个元素,然后是反转剩下的元素,最后是反转整个slice的元素。(如果是向右循环旋转, 则将第三个函数调用移到第一个调用位置就可以了。)   

代码实现:
```go
package main

import "fmt"

func main() {
	arr := [...]int{1,2,3,4,5}
	Reverse(arr[:])
	fmt.Printf("%v\n",arr)

	//左移动两位
	arr2 := []int{1,2,3,4,5}
	Reverse(arr2[2:]) //1,2,3,5,4
	Reverse(arr2[:2]) //2,1,3,5,4
	Reverse(arr2) // 3,4,5,1,2
	fmt.Printf("%v\n",arr2)

	//右移动三位
	arr3 := []int{5,6,7,8,9}
	Reverse(arr3) // 9,8,7,6,5
	Reverse(arr3[:3]) // 7,8,9,6,5
	Reverse(arr3[3:]) // 7,8,9,5,6
	fmt.Printf("%v",arr3)
}
func Reverse(s []int) {
	// for i,j := 0, len(s)-1; i < len(s)/2; i,j = i+1,j-1 {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
```

###  append函数
```go
func main() {
	sli := []int{}
	info(sli)
	sli = append(sli,1,2,3,4)
	info(sli)

}

func info(s []int) {
	fmt.Printf("len=%v,cap=%v,slice=%v\n", len(s), cap(s),s)
}
```

append函数的特殊使用，初始化nil值的slice

```go
//s := []rune //compile error , need . , or literal value ......
var s []rune
info(s)

//s = append(s, "hello world") //can not use string as []rune
//s = append(s, []rune("hello world")) //can not use []rune as type rune，追加元素必须是被追加slice的元素类型
//s1 := append(s, rune("hello world")) // cannot convert "hello world" to type rune. 字符串类型无法转换为rune类型
var runes []rune
for _,r := range "hello 中国" {
	//runes = append(runes, rune(r))
	runes = append(runes,r)
}
fmt.Printf("%q", runes) // ['h' 'e' 'l' 'l' 'o' ' ' '中' '国']


```

当然上述是为了演示append函数的用法及注意事项，上述需求可以更加简便的使用下列转换方式完成。  
```go
fmt.Printf("%q", []rune("hello 中国")) // ['h' 'e' 'l' 'l' 'o' ' ' '中' '国']
```

为了理解rune，我们来看下以下代码

```go
var r rune // int32的别名
//r = "hello" compile error
r = 'h'  
fmt.Printf("%q", r)

var c rune
c = 'h'
fmt.Printf("%q,%T", c, c) // 'h',int32
```

### slice之间不能比较 
与数组不同的是，slice之间不能比较
比如数组比较是通过的
```go
a := [...]int{1, 2, 5:10}
b := [...]int{1, 2, 5:10}
fmt.Printf("%v", a == b)
```

而切片则编译无法通过
```go
a := []int{1, 2, 5: 10}
b := []int{1, 2, 5: 10}
fmt.Printf("%v", a == b) //operation == not defined on []int
```

除了标准库提供了bytes.Equal()方法比较之外，其他的切片都需要我们自己展开来比较。
切片不支持`==`运算，是因为
1. slice是引用类型，即slice的元素是间接引用的，一个slice甚至可以包含自身。	这种情况不适合我们直接比较元素值，虽然可以解决，但并不高效。  
2. 由于slice是引用类型，一个固定的slice值其底层的数组元素值在不同的时间可能会被修改。  

```go
a := []byte{1, 2, 5: 10}
b := []byte{1, 2, 5: 10}
//fmt.Printf("%v", a == b) //operation == not defined on []int
fmt.Printf("%t", bytes.Equal(a,b)) // true
```

手动比较，注意代码的正宗golang味道
```go
a := []int{1, 2, 5: 10}
b := []int{1, 2, 5: 10}
//fmt.Printf("%v", a == b) //operation == not defined on []int
fmt.Printf("%t", compare(a, b))

//func compare(a []int, b []int) bool { taste so so!
func compare(x,y []int) bool { //taste good
	if len(x) != len(y) {
		return false
	}

	// 效率低
	/*for _, va := range a {
		for _, vb := range b {
			if va != vb {
				return false
			}
		}

	}*/
	
	// 这个代码味道一般
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	// taste good
	for i := range x {
        if x[i] != y[i] {
			return false	
		}
	}
	return true
}

```

### slice可以和nil比较(之一)
slice是可以和nil进行比较的。  

```go
func main() {
	var s []int
	info(s)
	s = nil
	info(s)
	//s = []int{nil}//compile error
	s = []int(nil)// nil类型转换为[]int
	info(s)
	s = []int{} //{}初始化了！非nil值的slice
	info(s)

}

func info(s []int) {
	fmt.Printf("len=%v,cap=%v,slice=%v,%v==nil?=%t,%T\n", len(s), cap(s), s, s, s == nil,s)
}

// len=0,cap=0,slice=[],[]==nil?=true,[]int
// len=0,cap=0,slice=[],[]==nil?=true,[]int
// len=0,cap=0,slice=[],[]==nil?=true,[]int
// len=0,cap=0,slice=[],[]==nil?=false,[]int
```
零值slice值为nil，类型[]int

### 优先使用len(s)==0判断slice是否为空
零值slice除了可以和nil进行比较外，其他的和长度为0的slice行为是一样的。除了有特殊说明之外，go语言将nil零值slice和长度为0的slice相等对待。reverse(nil)
```go
func main() {
	var s []int
	info(s)
	reverse(s)

	s0 := []int{}
	info(s)
	reverse(s0)
	
	s1 := []int{1, 2, 3, 4}
	info(s1)
	reverse(s1)
}
func info(s []int) {
	fmt.Printf("len=%v,cap=%v,slice=%v,%v==nil?=%t,%T\n", len(s), cap(s), s, s, s == nil, s)
}
```

### 模拟append函数，理解slice的扩容
slice的扩容依赖于长度和容量的变化，长度要始终不大于容量，一旦长度超过了容量，slice将进行2倍扩容。
我们先来理解下容量
#### 容量越界
举个超过容量越界的例子
```go
m := []int{1,2,3}
n := m[:4] //panic: runtime error: slice bounds out of range
fmt.Printf("%v", n)
```
#### appendInt
这里通过容量不足时进行翻倍拓展slice来模拟内置append函数。  

```go
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := 7
	ret := appendInt(a, b)
	fmt.Printf("ret: %v\n", ret)

	b, c, d := 7, 8, 9
	ret2 := appendInts(a, b, c, d)
	fmt.Printf("ret2: %v\n", ret2)

	ret3 := appendInts(a, b)
	fmt.Printf("ret3: %v\n", ret3)

	var m, n []int
	//var m  []int
	for i := 0; i < 10; i++ {
		//n = appendInts(m, i)
		n = appendInts(m, i)
		fmt.Printf("%d cap=%d %v\n", i, cap(n), n)
		m = n //重点！！！

		/*m = appendInts(m, i)
		fmt.Printf("%d cap=%d %v\n", i, cap(m),m )*/
	}

}
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	fmt.Printf("%v,%v\n", zlen, cap(x))
	if zlen <= cap(x) { //len(x)+1 一定不能超过cap，否则越界
		//z = x[:] 这样容量只能是len(x),zlen自然就大于了cap(x)，越界
		z = x[:zlen] //这里大胆的在基础上加+1 len(x)+1 加了一个元素
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func appendInts(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		//copy(z[len(x):],y)
		copy(z, x)
	}
	//z[:len(x)] = x
	copy(z[len(x):], y)
	return z
}
```

如果长度不超过容量，则直接共享底层数组，进行赋值。否则查过容量，则翻倍拓展容量，并拷贝到新的数组，然后进行赋值。
因此我们无法确定新的slice是否和原始的slice共享底层数组，同样的引申来看，如果共享底层数组的情况，我们也无法确定原先的slice操作是否影响新的slice。

```go
// runes := append(runes,r) //not update!!! 
runes = append(runes,r) //对于变量进行更新
```

变量slice的更新实际应用需要特别注意

```go
var runes []rune
//runes =append(runes,'h') // should use this!!!!!
m := append(runes, 'h')
fmt.Printf("runes: %q\n", runes)
fmt.Printf("m: %q\n", m)

// runes: []
// m: ['h']
```
在go中，对长度，容量，或底层数组变化的操作而更新slice变量是非常有必要的。正确的使用slice，首先要明确尽管对底层数组的元素访问是间接的，但是对于slice对应结构体本身的指针，容量，长度是直接访问的。

要更新上述信息，则需要如`runes = append(runes,r)`这样进行显示地赋值操作。 所以slice不仅仅是引用类型，更是类似struct的聚合类型。  

模拟`[]int`的实际结构
```go
type IntSlice struct {
	ptr *int
	len,cap int
}
```

###slice的内存使用技巧
在原有slice内存空间上返回不包含空字符的字符串。
展示如何避免另辟空间，输入和输出的slice共享底层数组。
缺点，原先的数组将可能会被覆盖。

- 第一版，错误  

```go
s := []string{"hello", "", "world"}
fmt.Printf("%q\n", s)
fmt.Printf("%q\n", nonempty(s))

func nonempty(s []string) []string {
    i := 0 //当想在原有空间上操作，做好有个自控变量
	for _,v := range s {
		if s[i] != "" { //迭代过滤对象错误
			s[i] = v
			i++
		}
	}
	return s
}
```

- bug修复版  

```go
s := []string{"hello", "", "world"}
fmt.Printf("%q\n", s)
fmt.Printf("%q\n", nonempty(s))

func nonempty(s []string) []string {
	i := 0
	for _, v := range s {
		//if s[i] != "" {
		if v != "" {
			s[i] = v
			i++
		}
	}
	//return s //如果返回s，那么i自控变量的意义就没有了
	return s[:i]
}

// [hello  world]
// [hello world]
```

我们看到nonempty函数操作切片s后返回的也是共享的底层数组，但是值已经修改了。所以在使用上不能只想着切片是引用类型，可以进行原值修改，但事实上是需要更新后重新赋值，否则引用的还是原始底层数组的值。
原变量引用未变，如果不更新，则还是原始底层数组。
```go
s := []string{"hello", "", "world"}
fmt.Printf("%q\n", s) // [hello  world] // 原始底层值
fmt.Printf("%q\n", nonempty(s)) // [hello world] 结果值最好在实际应用中赋值保存
fmt.Printf("%q\n", s) // [hello world world] //底层值被修改
```
所以正确的使用 `data = nonempty(s)`

普通版
```go
func nonempty2(s []string) []string {
	out := make([]string, len(s), len(s))
	i := 0
	for _, v := range s {
		if v != "" {
			out[i] = v
			i++
		}
	}
	return out
}
```

小数据量味道好点的，但是大的还是前一版本带索引的访问较好。下面的重构版本还是在于slice的重构，以这种方式重用一个slice，一般都要求为最多为每个输入值产生一个输出值。  
```go
func nonempty3(s []string) []string {
	out := s[:0]
	for _,v := range s {
		if v != "" {
			append(out,v)
		}
	}
	return out
}
```

### 使用slice模拟stack
- 压栈(使用append函数)
```go
var stack []int
stack = append(stack,v) //push  v
```

- 栈顶
```go
top := stack[len(stack)-1]
```

- 弹栈(通过收缩slice)  
```go
stack = stack[:len(stack)-1] 
```

- 删除栈内元素   
通过内置的函数copy将元素依次向前移动一位。  
  
```go
var stack []int
stack = append(stack,1,2,3,4,5) 
fmt.Printf("%v",remove(stack,2)) // 1 2 4 5

func remove(s []int, i int) []int {
	copy(s[i:],s[i+1:])
	return s[:len(s)-1]
}
```

-------
参考资料:  
《The Go Programming Language》


























