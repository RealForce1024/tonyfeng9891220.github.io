```go
for i := 0; i < 10; i++ {
    //fmt.Printf("%v", 1<<i) // shift count type int ,must be unsigned integer
    // 按位左移
    fmt.Printf("%v\t%v\t%v\n", 1<<uint(i), 2<<uint(i), 3<<uint(i))
}
```