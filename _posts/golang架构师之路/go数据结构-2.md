<!-- TOC -->

- [go数据结构总览](#go数据结构总览)
- [map](#map)
    - [声明与初始化](#声明与初始化)
        - [map类型](#map类型)
        - [map的key必须是支持`==`运算符的类型](#map的key必须是支持运算符的类型)
        - [map声明与初始化](#map声明与初始化)
        - [通过下标索引访问map元素](#通过下标索引访问map元素)
        - [删除](#删除)
        - [key不存在map操作也是安全的](#key不存在map操作也是安全的)
        - [map元素的+= ++等操作](#map元素的-等操作)
        - [禁止对map元素取址操作](#禁止对map元素取址操作)
        - [迭代map](#迭代map)
        - [顺序遍历map](#顺序遍历map)

<!-- /TOC -->

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

## map
map是key/value键值对的无序集合。key是唯一的，不会重复，且是可以比较的类型。通过特定的key对value的检索，增加，删除都是常数时间复杂度内完成。

### 声明与初始化
#### map类型
```go
map[k]v
```
k,v不必是同一种类型，但k，v各自的集合都必须属于同一种类型，即声明的类型或隐式推断出的类型(字面值定义)。

#### map的key必须是支持`==`运算符的类型
- 对于key的要求 必须支持`==`运算
map通过key检索值，也就是map通过给定的值测试key是否已经存在，即进行两者的`==`运算比较。浮点数字类型虽然支持`==`操作，但由于精度可能导致的问题以及 NaN与任何浮点数都不等的情况 将浮点数作为key不是那么明智，所以轻易不要选择浮点型作为key。 
- 对于value数据类型没有任何限制，只要value集合类型一致即可

#### map声明与初始化
- make函数创建(声明并初始化)
```go
ages := make(map[string]int)
fmt.Printf("%v\t%t\n",ages,ages==nil)//map[]	false
```

- make(map,n)预设键值对数量
使用make也可以对map的键值对数量进行预设，在初始化时分配大量的内存，可以防止分配过程中频繁分配。  
但对于len(map)并没有什么影响。  

- map字面值
```go
ages := map[string]int{
	"kobe":40,
	"james":32,
}

ages["jordan"] = 52
ages["juice"] = 60
fmt.Printf("%v\n",ages)//map[james:32 jordan:52 juice:60 kobe:40]
```

- 一般定义
```go
var students map[int]string
fmt.Printf("%v\t%[1]t\t%t\n",students,students==nil)//map[]	true
students = map[int]string{
	1: "kobe",
	2: "jordan",
	3: "james",
}
fmt.Printf("%v\n", students) //map[2:jordan 3:james 1:kobe]
```

- 创建空的map表达式 非nil

```go
stu := map[string]string{} //等价于make(map[string]string)
fmt.Printf("%v\t%t\n", stu,stu==nil) //map[]	false
```
注意这里和`var students map[int]string`不同，但等价于`make(map[string]string)`

#### 通过下标索引访问map元素
`map[key]`可访问map的key对应的value值  
```go
ages := map[string]int{
		"kobe":23,
		"james":16,
	}

students := map[int]string{
	1: "kobe",
	2: "james",
	3: "jordan",
}
fmt.Printf("%d\n", ages["kobe"])//23
fmt.Printf("%s\n", students[1]) //kobe
```

#### 删除
通过内置函数`delete(map,key)`删除指定key的元素  

```go
ages := map[string]int{
		"kobe":23,
		"james":16,
	}

fmt.Printf("%d\t%v\n", ages["kobe"],ages)
delete(ages,"james")
fmt.Printf("%v\n", ages)

// 23	map[kobe:23 james:16]
// map[kobe:23]
```
#### key不存在map操作也是安全的
检索，删除操作都是针对map的key进行查找，特殊的一点在于如果key不存在，检索，删除操作也是安全的，对于检索其会返回value对应的零值，删除时，检索不到对应的key，将不会执行删除操作。  

```go
fmt.Prinft("%s",ages["jordan"])
delete(ages,"jordan")
```
#### map元素的+= ++等操作
++ += 等操作同样适合map元素
```go
fmt.Printf("%v\n", ages["kobe"])
ages["kobe"] += 1
ages["kobe"] ++
fmt.Printf("%v\n", ages["kobe"])
```
#### 禁止对map元素取址操作
之前我们讲到可以使用make(map,n)来预设map元素的数量以避免内存频繁动态的分配(map会随着元素数量的增长而重新分配更大的内存空间，从而导致原先的地址无效)。这里引申点在于map元素的地址是可能会发生变化的，所以取址操作没什么意义，go也是不允许编译通过的。
map元素也不是变量，所以&取址操作符也是无意义的。  

```go
_ := &ages["kobe"]

// no new variables on left side of :=
// cannot take the address of ages["kobe"]
```
除了不能取址操作，我们发现`:=`左边必须至少有一个新的变量声明。  

#### 迭代map
for range风格迭代map，每次循环设置k,v变量的值  
```go
ages := map[string]int{
		"kobe":  23,
		"james": 16,
		"jordan":34,
		"allen",25,
}

for k, v := range ages {
	fmt.Printf("%s:%v\n", k, v)
}

// james:16
// jordan:34
// allen:25
// kobe:23
```
map迭代是无序的，这是故意设计的，每次遍历基本都是不同的哈希实现即强制其遍历不依赖具体的哈希函数，使得遍历是随机的。
#### 顺序遍历map
无序是因为key的遍历是随机的，所以要想顺序遍历map，必须先对key进行排序。  

```go
ages := map[string]int{
	"kobe":   23,
	"james":  16,
	"jordan": 34,
	"allen":  25,
}

fmt.Println("随机排列--->")
for k, v := range ages {
	fmt.Printf("%s:%v\n", k, v)
}

fmt.Println("顺序排列--->")
keys := []string{}
for k := range ages {
	//append(keys, k)//Append returns the updated slice.
	// It is therefore necessary to store the result of append,
	// often in the variable holding the slice itself:
	keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
	fmt.Printf("%s:%d\n", k, ages[k])
}

// 随机排列--->
// james:16
// jordan:34
// allen:25
// kobe:23
// 顺序排列--->
// allen:25
// james:16
// jordan:34
// kobe:23
```









-------
参考资料:  
《The Go Programming Language》


























