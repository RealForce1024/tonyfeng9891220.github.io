面向对象:
使用方法表达对属性和对应行为的操作，无需直接去操作对象，而是使用方法来操作对象。

```go
const day = 24 * time.Hour
fmt.Println(day.Seconds()) // 86400
```

```go
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
```

go的面向对象两大特点: 封装 组合  

