
- 天生并发支持

- 简单 命名 定义 
- 返回值 
- 自带网络编程电池

- 零值初始化机制
**零值初始化机制**可以确保每个声明的变量总是有一个良好定义的值,因此**在Go语言中不存在 未初始化的变量**。这个特性可以简化很多代码,而且可以在没有增加额外工作的前提下确保 边界条件下的合理行为。

```go
var x,y int
var z string
fmt.Printf("%d %d %q", x, y, z)
//0 0 "" 
```
即使未手动显示初始化，但go有零值初始化机制，程序将打印零值，而不是导致错误或产生不可预知的行为

- 简短变量声明
`x, y := 3, 4` // := 变量声明语句  多重变量声明
`m, n = n, m`  //  = 变量赋值操作  多重变量赋值
注意:**简短声明变量对同一域内已经声明过的变量只有赋值操作了，而不是再声明。**

```go
in, err  := os.Open(inFile)
....
out, err := os.Create(outFile)
```

另外，简短声明变量组中必须至少有一个新声明的变量
```go
f, err := os.Open(inFile)
.....
f, err := os.Create(outFile) // compile error : no new variables

// 可以改为
f, err = os.Create(outFile) // ok 
```

注意:是同一域内
```go
_, err := os.Open("/usr/local/a.txt")
fmt.Printf("%v\n", &err)
if _, err := os.Create("/usr/localxx/b.txt"); err != nil {
    fmt.Printf("%v", &err)
}

// 0xc42000e260
// 0xc42000e280
```

```go
_, err := os.Open("/usr/local/a.txt")
fmt.Printf("%v\n", &err)

_, err = os.Create("/usr/localxx/b.txt");
if err != nil {
    fmt.Printf("%v", &err)
}

// 0xc4200701b0
// 0xc4200701b0
```
## go中的i++语句
go的i++是语句而非其他语言中的表达式，所以go中 `j=i++`为非法无效的。  

## os.Args 
os.Args是字符串slice，第一个元素为命令本身，其余的元素则为程序启动时传入的字符串参数

## 字符串
go的字符串为string类型，其`底层类似slice切片,支持切片操作`
### 计算字符串的实际长度
默认情况下字符的长度计算是计算的字符字节的长度，所以遇到中文这样非英文字母的字符，每个字符占有的字节数就不等，因此我们一般想要的字符串长度只是字符的长度，而非字符字节的长度。而`len(str)`计算的正是字符字节的长度。  

由于string类型底层类型slice切片，而go也可以通过len函数访问其长度(字符字节长度)，rune类型是int32的别名，也是go中表示字符的类型，而字符串也可以看做是许多单个字符的数组切片，通过`[]rune(string)`可以转换字符串为[]rune数组。

```go
fmt.Printf("%d\n", len("hello"))
fmt.Printf("%d\n", len("hello中国"))
fmt.Printf("%d\n", len([]rune("hello中国")))
s := "hello中国"
fmt.Printf("%d\n", len([]rune(s)))
fmt.Printf("%v\n",[]rune(s))
fmt.Printf("%q\n",[]rune(s))
fmt.Printf("%v\n",string([]rune(s)))

// 5
// 11
// 7
// 7
// [104 101 108 108 111 20013 22269]
// ['h' 'e' 'l' 'l' 'o' '中' '国']
// hello中国
```


和string类型有密切关联的两个类型 `byte`和`rune`类型  

```go
fmt.Printf("--%d\n", len([]rune("hello中国")))
//fmt.Printf("--%d\n", len([]uint32("hello中国"))) //虽然是别名在编译时通过，但是依然在运行时 cannot convert "hello中国" (type string) to type []uint32

var a byte = 'h'
fmt.Printf("byte 'h': %v\t%[1]T\t%[1]q\n", a) //byte 'h': 104	uint8	'h'
var b rune = 'h'
fmt.Printf("rune 'h': %v\t%[1]T\t%[1]q\n", b) //rune 'h': 104	int32	'h'

var s string = "echo framework"
fmt.Printf("%s\n", s)                  //echo framework
fmt.Printf("%v\t%[1]T\n", s[:4])       //echo	string
fmt.Printf("%s\t%[1]T\t%[1]v\n", s[0]) //%!s(uint8=101)	uint8	101
fmt.Printf("%v\t%[1]T\n", s[0])        //101	uint8
fmt.Printf("%q\n", s[0])               //'e'
```

字符串的元素类型

```go
var s string = "hello world"
```

## go语言使用可变栈

