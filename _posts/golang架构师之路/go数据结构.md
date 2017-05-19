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
#### 1.声明类型
数组类型: **`[n]Type`**
`var arr [5]int` //创建长度为5、int型的数组。数组元素为int类型零值`0`。

```go
package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr[0])
	fmt.Println(arr[len(arr)-1])
	for i, v := range arr {
		fmt.Printf("%d,%d\n", i, v)
	}
}
```
#### 2.数组字面值

```go
package main

import "fmt"

func main() {
	// var arr [5]int = [5]int{1, 2, 3, 3, 4} 代码味道很差
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v", arr)
}

```
