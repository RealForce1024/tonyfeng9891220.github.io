## Panic异常
> Go的类型系统会在编译时捕获很多错误,但有些错误只能在运行时检查,如数组访问越界、 空指针引用等。这些运行时错误会引起painc异常。

> 一般而言,当panic异常发生时,程序会中断运行,并立即执行在该goroutine(可以先理解成 线程,在第8章会详细介绍)中被延迟的函数(defer 机制)。随后,程序崩溃并输出日志信 息。日志信息包括panic value和函数调用的堆栈跟踪信息。panic value通常是某种错误信 息。对于每个goroutine,日志信息中都会有与之相对的,发生panic时的函数调用堆栈跟踪信 息。通常,我们不需要再次运行程序去定位问题,日志信息已经提供了足够的诊断依据。因 此,在我们填写问题报告时,一般会将panic异常和日志信息一并记录。

> 不是所有的panic异常都来自运行时,直接调用内置的panic函数也会引发panic异常;panic函 数接受任何值作为参数。当某些不应该发生的场景发生时,我们就应该调用panic。

当程序到达了某条逻辑上不可能到达的路径:
```go
switch s := suit(drawCard()); s {
    case "Spades":
    // ...
    case "Hearts":
    // ...
    case "Diamonds":
    // ...
    case "Clubs":
    // ...
    default:
        panic(fmt.Sprintf("invalid suit %q", s)) // Joker?
}
```

断言函数必须满足的前置条件是个明智的做法，但是这很容易被滥用。除非你能提供更多的错误信息，或者能更快速的发现错误，否则不需要使用断言。编译器会在运行时帮你检查代码。  
```go
func Reset(x *Buffer) {
    if x==nil{
        panic("x is nil")//unnecessary!
    }
    x.elements = nil
}
```

panic类似于其他语言的异常处理，但是应用场景不太一样，由于panic会引起程序的崩溃，所以一般用于特别严重的错误，比如程序逻辑的不一致或错误等
>勤奋的程序员认为 任何崩溃都表明代码中存在漏洞,所以对于大部分漏洞,我们应该使用Go提供的错误机制, 而不是panic,尽量避免程序的崩溃。在健壮的程序中,任何可以预料到的错误,如不正确的输入、错误的配置或是失败的I/O操作都应该被优雅的处理,最好的处理方式,就是使用Go的错误机制。

```go
package main

import "fmt"

func main() {
	f(3)
}

func f(x int)  {
	fmt.Printf("f(%d)\n", x)
	defer fmt.Printf("defer %d\n", x+0/x)
	f(x-1)
}

```

当执行到f(0)时，会引起panic，defer会立即执行
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3
f(0) 引起的panic 异常堆栈信息
```go
panic: runtime error: integer divide by zero
f(3)
f(2)

f(1)
goroutine 1 [running]:
f(0)
main.f(0x0)
defer 1
	/Users/fqc/work/src/run.go:11 +0x1bb
defer 2
main.f(0x1)
defer 3
	/Users/fqc/work/src/run.go:12 +0x180
main.f(0x2)
	/Users/fqc/work/src/run.go:12 +0x180
main.f(0x3)
	/Users/fqc/work/src/run.go:12 +0x180
main.main()
	/Users/fqc/work/src/run.go:6 +0x2a
```

```go
func main() {

defer printStack()
    f(3)
}

func printStack() {
    var buf [4096]byte
    n := runtime.Stack(buf[:], false)
    os.Stdout.Write(buf[:n])
}
```
>将panic机制类比其他语言异常机制的读者可能会惊讶,runtime.Stack为何能输出已经被释放 函数的信息?在Go的panic机制中,延迟函数的调用在释放堆栈信息之前。

Recover捕获异常
当程序因为panic发生崩溃时，我们并不希望程序因此终止，因为很多收尾或资源等问题，我们可能还需要回退或恢复等让其继续正常工作。


[关于golang的panic recover异常错误处理](http://xiaorui.cc/2016/03/09/%E5%85%B3%E4%BA%8Egolang%E7%9A%84panic-recover%E5%BC%82%E5%B8%B8%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86/)

[Go语言中使用Defer几个场景](http://developer.51cto.com/art/201306/400489.htm)










