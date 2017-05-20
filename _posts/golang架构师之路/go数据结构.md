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

注意:`[...]int{n}`是数组字面值的一中变形，字面值书属于定义范畴

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


### 反正slice和字符串
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

- 反正字符串
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


###  










