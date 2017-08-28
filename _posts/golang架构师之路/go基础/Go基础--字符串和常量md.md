
## 字符串
### 字符串表示

```go
package main

import (
	"fmt"
)

func main() {
	var str1 string = "\\\""
	fmt.Printf("%v\n",str1)

	fmt.Printf("用解释型字符串表示法表示的 %q 所代表的是 %s。\n", str1, `\"`)
	fmt.Printf("用解释型字符串表示法表示的 %q 所代表的是 %s。\n", str1,"\\")
	fmt.Printf("用解释型字符串表示法表示的 %q 所代表的是 %s。\n", str1,`\\`)
}
// \"
//用解释型字符串表示法表示的 "\\\"" 所代表的是 \"。
//用解释型字符串表示法表示的 "\\\"" 所代表的是 \。
//用解释型字符串表示法表示的 "\\\"" 所代表的是 \\。
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

//0	h
//1	e
//2	l
//3	l
//4	o
//5	,
//6
//7	世
//10	界

```

```go
package main

import (
	"fmt"
)

func main() {
	s := "hello, 世界"

	for i, val := range s { //val为int32类型 即rune
		fmt.Printf("%d\t%q\t%d%t%[2]T\n", i, val, val)
	}

	b := "hello"
	for _,v := range b { // int32并未自动降级int8
		fmt.Printf("%v\t%[1]T\n", v)
	}
}


//0	'h'	104	int32
//1	'e'	101	int32
//2	'l'	108	int32
//3	'l'	108	int32
//4	'o'	111	int32
//5	','	44	int32
//6	' '	32	int32
//7	'世'	19990	int32
//10	'界'	30028	int32
//104	int32
//101	int32
//108	int32
//108	int32
//111	int32
```
###字符串连接
#### `+`连接字符串 

```go
fmt.Println("hello"+"world")
```

#### `join`串联字符串

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

//9
//9
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
	fmt.Printf("%v\n", b) //A

	var c int = 65
	d := strconv.Itoa(c)
	fmt.Printf("%v\t%s\t%q\n", d,d,d)

	f, _ := strconv.Atoi(d)
	fmt.Printf("%v", f)
}


//A
//65	65	"65"
//65
```

### 不能修改的字符串该如何修改
字符串一旦定义不能修改，否则将无法编译通过。

```go
var s string = "hello"
s[0] = "e" //compile error
```

不过可以曲线救国，赋值给string底层类型byte[]
string和[]byte可以互转
```go
var s string = "hello"
var c []byte
//c = s[:] 注意s[:]为string类型  string[n]为byte类型
c = []byte(s)
c[0] = 'e'
s = string(c) // []byte和s是可以互转的
fmt.Printf("%v", s)
```
注意：是字符串内容元素不能修改，但不是说字符串变量的值不能修改。

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


//eello 中国
//32654	'美'
//eello 美国
```

## 常量

- 常量声明很像变量，只不过需要使用`const`关键字修饰
- 常量可以使字符，字符串，布尔或数值型
- 常量不能使用`:=`声明
- 常量的值必须在编译期就能够确定!!! 可以在赋值表达式中涉及计算过程，但是所有用于计算的值需要在编译期就能获得。比如`const c1=1/2` ok,但`const c2 = getNumber()`自定义函数在编译期无法获得具体值，因此无法用于常量的赋值。 但内置函数是可以,比如`len()`。    
**常量是定义在程序编译阶段就确定下来的值，而程序在运行时无法改变该值。**

```go
package main

import "fmt"

var a string = "hello"

const (
	b = len("hello")
	c //省略的话，直接和b的声明类型，值等相等
	//d = len(b) // 编译错误 这里虽然是内置函数,但是引用给了变量的方式是无法用于常量的
)

func main() {
	fmt.Println(c) //5
}
```

- 等号右边必须是常量或常量表达式，**注意是常量表达式，一般的变量也是编译不通过的**
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

iota是常量的计数器，从0开始，组中每定义一行(非一个)常量都会自动加1（按行+1）

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
   fmt.Println(a,b)
	fmt.Println(c,d,e,f)
}
// 0 0
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
//0 0 1
```

- iota在**新的一行**都会自动加1。有了这个规则，就可以简写为

```go
const(
	a = iota
	b
	c
)

const(
	x = iota + 100//0+100
	y //100+1
	z //101+1
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


