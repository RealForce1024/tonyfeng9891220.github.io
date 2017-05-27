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















