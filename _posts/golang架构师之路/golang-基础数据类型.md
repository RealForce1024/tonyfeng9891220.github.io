## go中的i++语句
go的i++是语句而非其他语言中的表达式，所以go中 `j=i++`为非法无效的。  

## os.Args 
os.Args是字符串slice，第一个元素为命令本身，其余的元素则为程序启动时传入的字符串参数

## 字符串
可以结合 《Go基础之一》

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

字符串的元素类型

和string类型有密切关联的两个类型 `byte`和`rune`类型  
string获取字符类似切片获取元素，获取子串同样可以通过切片的方式。不过在获取字符的时候如s[0]，返回的是int8(byte为其别名)类型，输出的是asiic码，可以通过`string(val byteT)`的方式如string(s[0])进行字符串的转换  
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

```go
/*
// Contains reports whether substr is within s.
func Contains(s, substr string) bool {
    return Index(s, substr) >= 0
}

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny(s, chars string) bool {
    return IndexAny(s, chars) >= 0
}

索引>=0 或 不为-1即包含该字符串
*/
s := "hello world"
fmt.Printf("%t\n", strings.Contains(s, "ello")) //true
fmt.Printf("%t\n", strings.Contains(s, ""))     //true 空字符串”“在任何字符串中均存在。
//与Contains略有不同的是ContainsAny返回的是unicode码点（任何文字在Unicode中都对应一个值，这个值称为代码点（code point））
//由于对于空字符的处理方式ContainsAny为false，所以以后首选ContainsAny
fmt.Printf("%t\n", strings.ContainsAny(s, ""))                                                //false
fmt.Printf("%t\n", strings.ContainsAny(s, "he"))                                              //true
fmt.Printf("%t\n", strings.ContainsAny(s, "m"))                                               //false
fmt.Printf("%t\n", strings.ContainsAny(s, "eo"))                                              //true
fmt.Printf("%t\n", strings.ContainsAny(s, "Eo"))                                              //true
fmt.Printf("%t\n", strings.ContainsAny(s, "E"))                                               //false
fmt.Printf("%t\n", strings.ContainsAny(strings.ToLower(s), strings.ToLower("E")))             //false
fmt.Printf("%t\n", strings.ContainsRune(s, 'h')) 
```



