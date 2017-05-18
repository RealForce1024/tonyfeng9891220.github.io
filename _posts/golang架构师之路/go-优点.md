
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

## go语言使用可变栈


##