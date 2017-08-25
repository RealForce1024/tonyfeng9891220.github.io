# handlerFunc或中间件是如何被调用的

```go
echo.Get("/",func(c echo.Context) error {
    return c.String(http.StatusOK,"hello")
})
echo.start(":8080")
```
Get方法的参数一个是路由路径，另外一个是封装的handlerFunc，最终在ServeHTTP方法调用。
![-w450](media/15036373978020.jpg)


## echo api CRUD
模式总结:

- httpHandler: `echo.$httpMethod($path,$handlerFunc)`
- 参数获取: Param, FormValue, 文件
- 快捷访问: curl的使用，注意`$path`路由路径末尾带有"/"将是全匹配，也就是访问时必须带有`/`

## curl json

```sh
☁  ~  curl -s  http://localhost:8080/users | python -m json.tool
{
    "0": {
        "Id": 0,
        "email": "",
        "name": "kobe0"
    },
    "1": {
        "Id": 1,
        "email": "",
        "name": "kobe1"
    },
    "2": {
        "Id": 2,
        "email": "",
        "name": "kobe2"
    }
}
```
当然，也看到了个有趣的,到第10个之后居然看不到了，最大就是第九，向前一翻，原来其按照字典顺序来排列的。  

![-w450](media/15036433021222.jpg)




