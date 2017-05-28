状态码用来告诉HTTP客户端,HTTP服务器是否产生了预期的Response。HTTP/1.1协议中定义了5类状态码， 状态码由三位数字组成，第一个数字定义了响应的类别

1XX 提示信息 - 表示请求已被成功接收，继续处理
2XX 成功 - 表示请求已被成功接收，理解，接受
3XX 重定向 - 要完成请求必须进行更进一步的处理
4XX 客户端错误 - 请求有语法错误或请求无法实现
5XX 服务器端错误 - 服务器未能实现合法的请求


HTTP协议是无状态的和Connection: keep-alive的区别

无状态是指协议对于事务处理没有记忆能力，服务器不知道客户端是什么状态。从另一方面讲，打开一个服务器上的网页和你之前打开这个服务器上的网页之间没有任何联系。

HTTP是一个无状态的面向连接的协议，无状态不代表HTTP不能保持TCP连接，更不能代表HTTP使用的是UDP协议（面对无连接）。

从HTTP/1.1起，默认都开启了Keep-Alive保持连接特性，简单地说，当一个网页打开完成后，客户端和服务器之间用于传输HTTP数据的TCP连接不会关闭，如果客户端再次访问这个服务器上的网页，会继续使用这一条已经建立的TCP连接。

Keep-Alive不会永久保持连接，它有一个保持时间，可以在不同服务器软件（如Apache）中设置这个时间。




网页优化方面有一项措施是减少HTTP请求次数，就是把尽量多的css和js资源合并在一起，目的是尽量减少网页请求静态资源的次数，提高网页加载速度，同时减缓服务器的压力。

Go语言里面提供了一个完善的net/http包，通过http包可以很方便的就搭建起来一个可以运行的Web服务。同时使用这个包能很简单地对Web的路由，静态文件，模版，cookie等数据进行设置和操作。

```go
package main
import (
    "fmt"
    "net/http"
    "strings"
    "logs"
)

func sayHelloController(w http.ResponseWriter, r *http.Request){
    r.ParseForm() 
    fmt.Println(r.Form) // 
    fmt.Println("path",r.Url.Path)
    fmt.Println("schema",r.Url.Schema)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:",k)
        fmt.Println("val:",strings.join(v,""))
    }
    
    fmt.Fprintf(w,"hello fqc")
}

func main() {
    http.HandlerFunc("/",sayHelloController)
    err := http.ListenAndServe(":9090",nil)
    if err != nil {
        log.Fatal("Listen and Serve error: ",err)
    }

}
```


[使用supervisor管理golang程序进程](http://www.01happy.com/supervisor-golang-daemon/)
