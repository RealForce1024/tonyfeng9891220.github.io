
## 字符串
### 字符串表示

```go
package main

import (
    "fmt"
)

func main() {
    var str1 string = "\\\""

    fmt.Printf("用解释型字符串表示法表示的 %q 所代表的是 %s。\n", str1, `\`)
}
```
字符串表示有两种方式

1. 原生表示法
2. 解释型表示法
后者回车符等特殊符号会被转义。   

### 原生字符串面值

```go
const a = `
		hello world
		jkjkjkj      jkjkjkjkjdksjf                kljlj\\n
		\r\n
		helllo
	`
fmt.Printf("%v", a)

/*
		hello world
		jkjkjkj      jkjkjkjkjdksjf                kljlj\\n
		\r\n
		helllo
*/
```

### 字符串的长度

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, 世界"
	fmt.Printf("%d\n", len(s))
	fmt.Printf("%d\n", utf8.RuneCountInString(s))

	for i, value := range s {
		fmt.Printf("%d\t%q\t%d\n", i, value,value) //%q
		//fmt.Printf("%d\t%s\t%d\n", i, value,value)
	}
}
// 13
// 9
// 0	'h'	104
// 1	'e'	101
// 2	'l'	108
// 3	'l'	108
// 4	'o'	111
// 5	','	44
// 6	' '	32
// 7	'世'	19990
// 10	'界'	30028
```

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, 世界"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i+=size // 步长是字符的size
	}
}

```

```go
package main

import (
	"fmt"
)

func main() {
	s := "hello, 世界"

	for i, val := range s { //val为int32类型 即rune
		fmt.Printf("%d\t%q\t%d%[2]T\n", i, val, val)
	}

	b := "hello"
	for _,v := range b { // int32并未自动降级int8
		fmt.Printf("%v\t%[1]T\n", v)
	}
}
```

### `+`连接字符串 

```go
fmt.Println("hello"+"world")
```

### `join`串联字符串

```go
fmt.Println(strings.Join([]string{"hello", "world", "中国"}, "-"))
```

```go
package main

import (
	"fmt"
)

func main() {
	/*s := "hello, 世界"
	n := 0
	for _,_ := range s { // no new variables on left side of :=
		n++
	}*/

	n := 0
	for _, _ = range "hello, 世界" {
		n++
	}

	fmt.Printf("%v\n", n)

	count := 0
	for range "hello, 中国" {
		count++
	}
	fmt.Printf("%v\n", count)
}

```
### 转换为字符串
string()将数据转换为文本格式。计算机中存储的任何东西本质都是数字0,1。因此自然将65转为对应的文本A。

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 65
	b := string(a)
	fmt.Printf("%v\n", b)

	var c int = 65
	d := strconv.Itoa(c)
	fmt.Printf("%v\n", d)

	f, _ := strconv.Atoi(d)
	fmt.Printf("%v", f)
}
```

### 不能修改的字符串该如何修改
字符串一旦定义不能修改，否则将无法编译通过。

```go
var s string = "hello"
s[0] = "e" //compile error
```

不过可以曲线救国，赋值给string底层类型byte[]
string和byte[]可以互转
```go
var s string = "hello"
var c []byte
//c = s[:] 注意s[:]为string类型  string[n]为byte类型
c = []byte(s)
c[0] = 'e'
s = string(c) // []byte和s是可以互转的
fmt.Printf("%v", s)
```

但是byte不能存储汉字等字符

```go
var s string = "hello 中国"
var c []byte
//c = s[:] 注意s[:]为string类型  string[n]为byte类型
c = []byte(s)
c[0] = 'e'
s = string(c) // []byte和s是可以互转的
fmt.Printf("%v\n", s)

//c[6] = '美' //注意 byte alias as int8 无法存储汉字 constant 32654 overflows byte
//var b byte = '美'
//fmt.Printf("%v\n", b)
var r rune = '美'
fmt.Printf("%v\t%[1]q\n", r)

var runes []rune
runes = []rune(s)
//runes[6] = "美"
runes[6] = '美'
s = string(runes)
fmt.Printf("%v\n", s)
```



## 常量
- 常量声明很像变量，只不过需要使用`const`关键字修饰
- 常量可以使字符，字符串，布尔或数值型
- 常量不能使用`:=`声明
- 常量的值必须在编译期就能够确定!!! 可以在赋值表达式中涉及计算过程，但是所有用于计算的值需要在编译期就能获得。比如`const c1=1/2` ok,但`const c2 = getNumber()`自定义函数在编译期无法获得具体值，因此无法用于常量的赋值。 但内置函数是可以的`len()`。    
常量是定义在程序编译阶段就确定下来的值，而程序在运行时无法改变该值。

```go
package main

import "fmt"

var a string = "hello"

// const a = "hello"
const (
	b = len(a) // 编译错误
	c
)

func main() {

}
```
- 等号右边必须是常量或常量表达式，注意是常量表达式，一般的变量也是编译不通过的
- 常量表达式必须是内置函数
- 反斜杠`\`可以在常量表达式中作为多行的连接符使用。  
- 常量用作枚举

```go
const (
	UNKNOW = 0
	MALE = 1
	FEMALE = 2
)
```

```go
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```

和var()`打包`声明类似，const也可以。  

```go
package main

import (
	"fmt"
)

const (
	A int    = 10
	B string = "ss"
)

func main() {
	fmt.Println(A, B)
}
```

### 常量的初始化规则
- 定义常量组时，如果不给初始值，则表示使用上行表达式。  

```go
package main

import "fmt"

const (
	a = 1
	b = 2
	c
	d
)

func main() {
	fmt.Printf("%v,%v,%v,%v", a, b, c, d)
}
// 1,2,2,2
```

- 每一行声明的变量的个数需要一致

```go
package main

import "fmt"

const (
	a, b = 1, "2"
	//c  编译不通过，需要每一行的变量个数相同
	//c //extra expression in const declaration
	c, d //1, 2
)

func main() {
	fmt.Printf("%v, %v", c, d)
}
```
### 数值型常量  

- 数值型常量是高精度的值。  
- 未指定类型的常量类型取决于上下文
- int可以最大可存储64位整数，有时可能少些取决于系统平台

```go
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
    return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```

### iota

iota是常量的计数器，从0开始，组中每定义一行(非一个)常量都会自动加1

```go
package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

func main() {
	fmt.Printf("%v,%v,%v", a, b, c)
}
// 0,1,2
```

```go
package main

import "fmt"

func main() {
	const(
		a,b = iota,iota
		//c  compile error
		c,d
		e,f
	)

	fmt.Println(c,d,e,f)
}

// 1 1 2 2
```


```go
package main

import "fmt"

func main() {
	const(
		a,b = iota,iota
		c = iota
	)

	fmt.Println(a,b,c)
}
```

- iota在**新的一行**都会自动加1。有了这个规则，就可以简写为

```go
const(
	a = iota
	b
	c
)

const(
	x = iota + 100
	y
	z
)
// 100,101,102
```

使用类型作为常量的枚举值
通过初始化规则与iota可以达到枚举的效果  

```go
type size int // 定义size类型 也是int别名(底层类型为int)

const (
	small size = iota
	middle
	max
)
```

注意:一般情况常量名称最好全大写，但是由于go包的可见性规则，**我们有时只希望自己包中使用，则可以采用
_下划线或cConstName的方式取巧避免与规则冲突。**

- 每遇到一个const，iota就会重置为0

```go
package main

import "fmt"

const (
	a, b = iota, iota //0 0
	c, d //1 1
	e, f // 2 2
)

const aa = iota // 0

func main() {
	fmt.Printf("%v, %v,%v, %v\n", c, d, e, f)
	fmt.Println(aa)
}

```

### 常量iota的值与之定义的顺序有关系，和出现的次数并无关系。

```go
const (
	a, b = iota, iota //0 0
	c, d //1 1
	e, f // 2 2
)
```

```go
package main

import "fmt"

const (
	a = "A"
	b
	c = iota
	d
)

func main() {
	fmt.Printf("%v, %v, %v, %v", a, b, c, d)
}
// A A 2 3
```

```go
package main

import "fmt"

const (
	a = "A"
	b = iota
	c = "B"
	d = iota
)

func main() {
	fmt.Printf("%v, %v, %v, %v", a, b, c, d)

}
// A 1 B 3
```

```go
package main

import "fmt"

func main() {
	const (
		a = 1
		b
		c = iota
		d
		e = 10
		f
		g = iota
		h
	)

	fmt.Printf("a=%d,b=%d\n", a, b)
	fmt.Printf("c=%d,d=%d\n", c, d)
	fmt.Printf("e=%d,f=%d\n", e, f)
	fmt.Printf("g=%d,h=%d\n", g, h)
}

// a=1,b=1
// c=2,d=3
// e=10,f=10
// g=6,h=7
```





## 函数

1. 函数可以带有0或多个参数。下面的示例中,`add`函数包含两个`int`类型的参数。

```go
package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(40, 60))
}
```

注意**类型在变量的后面**。
可以参考Go的作者之一Rob Pike写的文章[《Go的声明语法》](https://blog.go-zh.org/gos-declaration-syntax)解释为什么使用这种类型声明在后,字段名称在前的方式。
通过阅读，其实最大的好处是在函数式编程中，go语法更便于阅读，将类型放后，函数名中间，func声明在前。

2. 当两个或更多连续的函数参数是同一种类型时，可以省略类型除了最后一个参数。  
例如:
可以将`x int, y int` 简写为`x, y int`

`1.`中的`add`函数参数可以简写为以下:

```go
func add(x, y int) int {
    return x + y
}
```

### 多个返回值
Go函数可以有多个返回值。下面的`swap`函数返回两个字符串

```go
package main

import "fmt" // gogland(IDE) 键入词imp 当使用包组织方式而只有一个包路径时，删除后是可以自动补全包路径的，很智能

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

参数是传递变量副本。不论是值的副本还是指针地址，都变量副本。值的拷贝。当然在指针地址的拷贝传递后盖板会影响到原变量，因为指针指向是原变量。后面会更深刻的阐述。

```go
a, b := "hello", "world"
m, n := swap(a, b)
fmt.Println(a, b) //a,b的值并未交换
fmt.Println(m, n) //返回值变量是a,b变量副本的交换后的值
```


### 命名返回值(naked return)
Go的返回值是可以命名的，它们将被当做变量声明在方法的签名上。

返回值的名称通常被用在标记返回值的含义上。

没有带参数的`return`语句块返回被命名的返回值。这被称作是"赤裸裸的"回报("naked" return)...即"裸"返回或者直接返回语句。  

直接返回语句仅应该用在短函数中，例如下面的示例。但注意:命名返回值用在内容较长的函数中影响可读性。

```go
package main

import "fmt"

func add(x int, y int) (result int) {
	result = x + y
	return
}
func main() {
	fmt.Println(add(2,1))
}
```

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return //naked return
}

func main() {
	x, y := split(99)
	fmt.Println(x, y)
}
```
## 函数值

```go
// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}
func main() {
    f := squares()
    fmt.Println(f()) // "1"
    fmt.Println(f()) // "4"
    fmt.Println(f()) // "9"
    fmt.Println(f()) // "16"
}
```
squares的例子证明，**函数值不仅仅是一串代码，还记录了状态**。在squares中定义的匿名内部函数可以访问和更新squares中的局部变量，这意味着**匿名函数和squares中，存在变量引用**。这就是**函数值属于引用类型**和函数值不可比较的原因。**Go使用闭包（closures）技术实现函数值，Go程序员也把函数值叫做闭包。**

变量的生命周期不由它的作用域决定(闭包):square闭包函数返回后,变量x仍然隐式的存在于f中。

但是需要注意一定要获取到函数值的引用，否则只是相同的值而已。因为只有获得闭包返回变量后才能获得匿名函数的隐式变量。  

```go
func main() {
	fmt.Println(square()())
	fmt.Println(square()())
	fmt.Println(square()())
	fmt.Println(square()())
	fmt.Println("-------")
	f:= square()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func square() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
```

## 函数进阶 函数类型与函数值(闭包)
go语言本质都是**值**类型，函数是**函数值**类型。因此我们可以简单的声明一个函数类型
`var f func(x int, y int) int`  其零值为`nil`

```go
package main

import "fmt"

type myFunc func(x int, y int) int

func main() {
	var f1 func(a string) string //匿名函数类型
	fmt.Println(f1)              // nil

	var f2 myFunc = func(x int, y int) int{return x+y} //其中myFunc也可以省略
	fmt.Printf("%T\n",f2) // myFunc
	fmt.Println(f2(3, 4))

	var f3 = func(x int, y int) int { //函数值（一个函数的实现需要被使用）
		return x + y
	}
	fmt.Println(f3(1, 2)) // 3
}

//<nil>
//main.myFunc
//7
//3
```

接下来更加进一步的将函数表达式(函数变量)进行调用,`本质是匿名函数的调用`

- **函数调用表达式**直接返回结果值
- 匿名函数调用返回值

```go
var result = func(x,y int) int {
		return x+y
	}(11,2)
fmt.Println(result)
```

**函数值之间是不可以比较的，因此也不能作为map的key**
**函数值**不仅仅使我们可以使用**数据参数化函数**，也可以通过**行为**。 (函数可以**值参数化**和**行为参数化**)

```go
package main

import "fmt"

func main() {
	var a = add //var a int = add
	fmt.Println(a(3, 5))

	var b func(x,y int) int
	b = add
	fmt.Println(b(3,1))
}

func add(x, y int) int {
	return x + y
}
```

- 函数行为参数化

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Map(add1, "abcd"))
	fmt.Println(strings.Map(add1, "HAL-9000"))

}

func add1(r rune) rune {
	return r + 1
}
//bcde
//IBM.:111
```

### 匿名函数作为函数值在使用的时候定义

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Map(func(r rune) rune {
		return r + 1
	}, "HAL-9000"))
}
```

匿名函数字面值表达式在使用时在定义，这种定义方式使得**匿名函数可以访问其所在函数的完整词法环境**。这意味着在函数中定义的**内部函数可以引用该函数的变量**。  

go使用闭包技术实现函数值，Go也把函数值称之为闭包。  

```go
package main

import (
	"fmt"
)

/*func square() int {
	var x int
	x++
	return func(x int) int {
		return  x * x
	}(x)
}*/


func square() func() int{
	var x int

	return func() int{
		x++
		fmt.Printf("%v\t%p\n",x,&x)
		return x*x
	}
}
func main() {
	f := square()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	
	/*
	1	0xc420070188
	1
	2	0xc420070188
	4
	3	0xc420070188
	9
	4	0xc420070188
	16
	*/

	//fmt.Println(square()())
	//fmt.Println(square()())
	//fmt.Println(square()())
	//fmt.Println(square()())
	//1	0xc42000e238
	//1
	//1	0xc42000e298
	//1
	//1	0xc42000e2c0
	//1
	//1	0xc42000e2e8
	//1
}

```

>squares的例子证明,函数值不仅仅是一串代码,还记录了状态。在squares中定义的匿名内部函数可以访问和更新squares中的局部变量,这意味着匿名函数和squares中,存在变量引用。这就是函数值属于引用类型和函数值不可比较的原因。Go使用闭包(closures)技术实现函数值,Go程序员也把函数值叫做闭包。

这里再次证明了变量的**作用域和生命周期是两个不同的概念，一个是静态的，一个是动态的。**  


### 函数值的重要特性-记录迭代变量的内存地址

golang中使用`:=`赋值初始化新变量，如果多次则会复用地址。这样在循环中，可以节省很多空间浪费。

```go
package main

import "fmt"

func main() {
	var strs = []string{"hello","world","china"}

	for _, v := range strs{
		fmt.Printf("%v\t%p\n", v,&v)
	}
	fmt.Printf("%v\n", "----------------")
	for _, v := range strs {
		v = v
		fmt.Printf("%v\t%p\n", v, &v)
	}

	fmt.Printf("%v\n", "----------------")
	for _, v := range strs {
		v := v
		fmt.Printf("%v\t%p\n", v, &v)
	}
}

// hello	0xc4200701b0
// world	0xc4200701b0
// china	0xc4200701b0
// ----------------
// hello	0xc420070200
// world	0xc420070200
// china	0xc420070200
// ----------------
// hello	0xc420070250
// world	0xc420070270
// china	0xc420070290
```

但是我们要说的重点是在于`for range`和闭包的结合使用，会产生的坑效果。

```go
	strs := []string{"hello", "world", "china"}
	for _, v := range strs {
		go func() {
			fmt.Println(v) //闭包 记录迭代变量的内存地址，而不是迭代某一刻的值！！
		}()
	}
	select {}

	//china
	//china
	//china
	//fatal error: all goroutines are asleep - deadlock!
	//
	//	goroutine 1 [select (no cases)]:
	//	main.main()
	//	/Users/fqc/work/src/myecho/main.go:14 +0xfc
	//	exit status 2

```
我们看到程序并未如我们所料，而是打印出了三次china，说明，v的地址是指向了最后的china，内部是函数闭包方式，产生了延迟调用，而v的地址在函数值被调用的时候已经是指向了最后的china。

原来**闭包~~函数值~~中记录迭代变量的内存地址，而不是迭代变量某一刻的值。单次执行看不出什么，但是在迭代中则会等待迭代执行完，才去执行函数值(闭包)**。  (迭代变量<->循环变量好多了)
在go语句和defer语句中也是如此。
所以需要**注意在循环体中将循环变量赋值给新的局部变量非常重要，否则每次获取的都是最后一次迭代值**。  


所谓闭包是指内层函数引用了外层函数中的变量或称为引用了自由变量的函数，其返回值也是一个函数，了解过的语言中有闭包概念的像 js，python，golang 都类似这样。


```go
strs := []string{"hello", "world", "china"}
	for _, v := range strs {
		go func(v string) {
			fmt.Println(v) //闭包函数值记录迭代变量的内存地址，而不是迭代某一刻的值！！
		}(v)
	}
	select {}
```
我们看到结果将会是strs数组的每个值。可以先忽略select{}，目前只是为了防止主routine先退出而看不到结果。



### 函数列表和匿名函数的使用不当

```go
package main

import "fmt"

func main() {
	for _, f := range test() {
		f()
	}
}

func test() []func() {
	var s []func()
	for i := 0; i < 3; i++ {
		s = append(s, func() {
			fmt.Printf("%v,%p\n", i, &i)
		})
	}
	return s
}

//3,0xc42007a050
//3,0xc42007a050
//3,0xc42007a050
```

解决方案
```go
func test() []func() {
	var s []func()
	for i := 0; i < 3; i++ {
		x := i
		/*	s = append(s, func(x int) {
				fmt.Printf("%v,%p\n", x, &x)
			}(x)) //这里就不应该调用啊
			*/
		s = append(s, func() {
			fmt.Printf("%v,%p\n", x, &x)
		})
	}
	return s
}
```


### 使用闭包修改全局变量
```go
package main                   

import (                       
    "fmt"                      
)                              

var x int = 1                  

func main() {                  
    y := func() int {          
        x += 1                 
        return x               
    }()                        
    fmt.Println("main:", x, y)                                                            
} 

// 结果：    main: 2 2
```
### 延迟调用
defer 调用会在当前函数执行结束前才被执行，这些调用被称为延迟调用，
defer 中使用匿名函数依然是一个闭包。


```go
package main

import "fmt"

func main() {
    x, y := 1, 2

    defer func(a int) { 
        fmt.Printf("x:%d,y:%d\n", a, y)  // y 为闭包引用
    }(x)      // 复制 x 的值

    x += 100
    y += 100
    fmt.Println(x, y)
}
```
输出结果：

101 102
x:1,y:102

```
从形式上看，匿名函数都是闭包。闭包的使用非常灵活，上面仅是几个比较简单的示例，不当的使用容易产生难以发现的 bug，当出现意外情况时，首先检查函数的参数，声明可以接收参数的匿名函数，这些类型的闭包问题也就引刃而解了。
```


