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
### 数组声明创建
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


