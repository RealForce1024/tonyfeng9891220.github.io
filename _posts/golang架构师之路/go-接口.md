## 接口
如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。

```go
type Men interface {
    sayHi(msg string)
}
```
interface可以被任意的对象实现。  

注意:任意的类型都实现了空的接口`interface{}`，也就是包含0个method的interface。 

## interface值
interface变量的值可以存储实现该接口的任意类型对象。

go的interface实现了"鸭子类型"。

空interface对于描述起不到任何作用，但是空interface在我们需要**存储任意类型的数值**的时候相当有用，它可以存储任意类型的数值。那可以联想到空接口作为参数和返回值是多么任性。  

```go
a := "hello"
b := 1
var c interface{}
c = a
c = b
```


## interface参数
fmt.Println函数都是默认的打印样式，但是我们通过源码看到go提供了Stringer接口只要实现了它的String方法，就将以自定义实现方式打印。

```go
package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}


func (h Human) String() string { //interface stringer 
	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years - ✆ " + h.phone + "❱"
}

func main() {
	bob := Human{"jordan",23,"1222222"}
	fmt.Println(bob)
}
```

fmt print.go 中定义了Stringer接口
```go
type Stringer interface {
    String() string
}
```

注意：
builtin.go
```go
type error interface {
	Error() string
}
```
注意: 实现了error接口的对象(即实现了Error() string的对象),使用fmt输出时,会调用Error()方法,因此 不必再定义String()方法了。
查看fmt包下的print.go源码可以清晰解释为何不需要再定义String()方法了。  
```go
func (p *pp) handleMethods(verb rune) (handled bool) {
	if p.erroring {
		return
	}
	// Is it a Formatter?
	if formatter, ok := p.arg.(Formatter); ok {
		....
	}
	// If we're doing Go syntax and the argument knows how to supply it, take care of it now.
	if p.fmt.sharpV {
        ...
		}
	} else {
		// If a string is acceptable according to the format, see if
		// the value satisfies one of the string-valued interfaces.
		// Println etc. set verb to %v, which is "stringable".
		switch verb {
		case 'v', 's', 'x', 'X', 'q':
			// Is it an error or Stringer?
			// The duplication in the bodies is necessary:
			// setting handled and deferring catchPanic
			// must happen before calling the method.
			switch v := p.arg.(type) {
			case error:
				handled = true
				defer p.catchPanic(p.arg, verb)
				p.fmtString(v.Error(), verb)
				return

			case Stringer:
				handled = true
				defer p.catchPanic(p.arg, verb)
				p.fmtString(v.String(), verb)
				return
			}
		}
	}
	return false
}
```
## interface变量存储的类型
我们知道interface可以存储任意类型的变量，但是我们如何获取一个interface值的实际类型呢?
有以下几种方式
1. Comma-ok断言 
```go
value, ok = element.(T) 
```
element为interface变量，ok我bool型，T为断言类型。
如果element中确实存储了T类型的值，ok返回true，否则返回false。  

实例:

```go
package main

import (
    "fmt"
    "stronv"
)

type Element interface
type List []Element
type Student struct {
    name string
    age int
}

func (s Student) String() string {
    return "(name:"+ s.name + " - age: " + strconv.Itoa(s.age)+"years)"
}

func main() {
    list := make(List,3)
    list[0] = 1
    list[1] = "hello"
    list[2] = Student{name:"kobe",23}

    for i, e := range list {
        if v, ok := e.(int);ok {
           ... 
        }else if v,ok := e.(string); ok {
            ...
        }else if v, ok := e.(Student); ok {
            ...
        }else{
            ...
        }
    }
}
```

使用switch
```go
```

2. 
```go
for i, e := range list {
        switch v := e.(T); v {
            case int:
                //...
            case string: 
                //...
            case Student:
                //...
            default:
        }
    }
```
需要强调的是:element.(type)。switch外面判断一个类型就使用comma-ok语法不能在switch外的任何逻辑里面使用,如果你要在switch外面判断一个类型就使用comma-ok。

## interface 嵌套
类似匿名字段的嵌套，interface也可以嵌套，那么被嵌套的interface的所有方法都可以隐式包含进来。  

```go
type ReadWriter interface {
    Reader
    Writer
}
```














