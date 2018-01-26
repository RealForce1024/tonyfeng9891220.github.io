## web 基于http协议进行表单提交 验证 上传等
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
    fmt.Println(r.Form) // map[url_long:[111 222]]  ?url_long=111&url_long=222
    fmt.Println("path",r.Url.Path)
    fmt.Println("schema",r.Url.Schema)
    fmt.Println(r.Form["url_long"]) // [111 222]
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
要编写一个Web服务器很简单，只要调用http包的两个函数就可以了(nginx,apache等服务器不是必须的)。Go通过简单的几行代码就已经运行起来一个Web服务了，而且这个Web服务内部有支持高并发的特性，Go是如何实现Web高并发的呢?

[使用supervisor管理golang程序进程](http://www.01happy.com/supervisor-golang-daemon/)


使用gogland运行web项目的页面总是显示不出，和gogland配置有关。 直接使用terminal执行go install and run 是没有问题的
```go
GOROOT=/usr/local/go
GOPATH=/Users/fqc/work
/usr/local/go/bin/go build -o "/private/var/folders/dz/9ggdb0ps3r31kxrr545jlffw0000gn/T/build and rungo" /Users/fqc/work/src/run.go
"/private/var/folders/dz/9ggdb0ps3r31kxrr545jlffw0000gn/T/build and rungo"
2017/05/29 15:07:40 http: panic serving [::1]:50581: runtime error: invalid memory address or nil pointer dereference
method: GET
goroutine 35 [running]:
net/http.(*conn).serve.func1(0xc4200b6be0)
	/usr/local/go/src/net/http/server.go:1721 +0xd0
panic(0x1314a20, 0x1505130)
	/usr/local/go/src/runtime/panic.go:489 +0x2cf
html/template.(*Template).escape(0x0, 0x0, 0x0)
	/usr/local/go/src/html/template/template.go:94 +0x38
html/template.(*Template).Execute(0x0, 0x14dd660, 0xc420120000, 0x0, 0x0, 0xc420075230, 0xc420034cc0)
	/usr/local/go/src/html/template/template.go:117 +0x2f
main.login(0x14e1760, 0xc420120000, 0xc420112300)
	/Users/fqc/work/src/run.go:30 +0x46f
net/http.HandlerFunc.ServeHTTP(0x13838b0, 0x14e1760, 0xc420120000, 0xc420112300)
	/usr/local/go/src/net/http/server.go:1942 +0x44
net/http.(*ServeMux).ServeHTTP(0x1511aa0, 0x14e1760, 0xc420120000, 0xc420112300)
	/usr/local/go/src/net/http/server.go:2238 +0x130
net/http.serverHandler.ServeHTTP(0xc4200aed10, 0x14e1760, 0xc420120000, 0xc420112300)
	/usr/local/go/src/net/http/server.go:2568 +0x92
net/http.(*conn).serve(0xc4200b6be0, 0x14e1da0, 0xc420078400)
	/usr/local/go/src/net/http/server.go:1825 +0x612
created by net/http.(*Server).Serve
	/usr/local/go/src/net/http/server.go:2668 +0x2ce
method: GET
2017/05/29 15:07:40 http: panic serving [::1]:50582: runtime error: invalid memory address or nil pointer dereference
goroutine 37 [running]:
net/http.(*conn).serve.func1(0xc4200b6d20)
	/usr/local/go/src/net/http/server.go:1721 +0xd0
panic(0x1314a20, 0x1505130)
	/usr/local/go/src/runtime/panic.go:489 +0x2cf
html/template.(*Template).escape(0x0, 0x0, 0x0)
	/usr/local/go/src/html/template/template.go:94 +0x38
html/template.(*Template).Execute(0x0, 0x14dd660, 0xc42014e000, 0x0, 0x0, 0xc4201480c0, 0xc420034cc0)
	/usr/local/go/src/html/template/template.go:117 +0x2f
main.login(0x14e1760, 0xc42014e000, 0xc420112500)
	/Users/fqc/work/src/run.go:30 +0x46f
net/http.HandlerFunc.ServeHTTP(0x13838b0, 0x14e1760, 0xc42014e000, 0xc420112500)
	/usr/local/go/src/net/http/server.go:1942 +0x44
net/http.(*ServeMux).ServeHTTP(0x1511aa0, 0x14e1760, 0xc42014e000, 0xc420112500)
	/usr/local/go/src/net/http/server.go:2238 +0x130
net/http.serverHandler.ServeHTTP(0xc4200aed10, 0x14e1760, 0xc42014e000, 0xc420112500)
	/usr/local/go/src/net/http/server.go:2568 +0x92
net/http.(*conn).serve(0xc4200b6d20, 0x14e1da0, 0xc4200787c0)
	/usr/local/go/src/net/http/server.go:1825 +0x612
created by net/http.(*Server).Serve
	/usr/local/go/src/net/http/server.go:2668 +0x2ce

```


即使使用命令行的方式运行，也需要特别注意一行代码
`r.ParseForm()`,否则后台解析前台form参数都为空。why? 默认情况下，Handler里面是不会自动解析form的，必须显式的调用r.ParseForm()后，你才能对这个表单数据进行操作。  
>我们输入用户名和密码之后发现在服务器端是不会打印出来任何输出的，为什么呢？默认情况下，Handler里面是不会自动解析form的，必须显式的调用r.ParseForm()后，你才能对这个表单数据进行操作。我们修改一下代码，在fmt.Println("username:", r.Form["username"])之前加一行r.ParseForm(),重新编译，再次测试输入递交，现在是不是在服务器端有输出你的输入的用户名和密码了。

如果修改action的值增加url参数，那么会和相应的表单字段值组成slice。其实在action中的参数会拼接到地址栏的url上。
http://localhost:9090/login?username=fqc
然后form username是我们自填的。 [form filed,url field.... ]
```go
<html>
<head>
    <title></title>
</head>
<body>
<form action="/login?username=fqc" method="post">
    用户名:<input type="text" name="username">
    密码:<input type="password" name="password">
    <input type="submit" value="登陆">
</form>
</body>
</html>
```


```go
v := url.Values{}
v.Set("name", "kobe")
v.Add("name", "james")
v.Add("name", "jordan")
fmt.Println(v.Get("name"))//kobe
fmt.Println(v.Get("name"))//kobe
fmt.Println(v)//map[name:[kobe james jordan]]
```

Tips: Request本身也提供了FormValue()函数来获取用户提交的参数。如r.Form["username"]也可写成r.FormValue("username")。调用**r.FormValue时会自动调用r.ParseForm**，所以不必提前调用。r.FormValue只会返回同名参数中的第一个，若参数不存在则返回空字符串。
注意v.Get()方法 `return vs[0]`所以 上段代码只打印出了'kobe'
```go

// Values maps a string key to a list of values.
// It is typically used for query parameters and form values.
// Unlike in the http.Header map, the keys in a Values map
// are case-sensitive.
type Values map[string][]string

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (v Values) Get(key string) string {
	if v == nil {
		return ""
	}
	vs := v[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}
```

做js注入，首先什么都不做的情况下，chrome浏览器会检测提交数据的合规性
```go
This page isn’t working

Chrome detected unusual code on this page and blocked it to protect your personal information (for example, passwords, phone numbers, and credit cards).
Try visiting the site's homepage.
ERR_BLOCKED_BY_XSS_AUDITOR
```


// username = "<script>alert('hello')</script>" //被注入后，服务器端如果再不转义
```go
//请求的是登陆数据，那么执行登陆的逻辑判断
fmt.Println("username:", r.Form["username"])
fmt.Println("password:", r.Form["password"])
fmt.Println("========================")
fmt.Fprintf(w, "%s", "<script>alert('hello')</script>")
//fmt.Fprintf(w, "%s", strings.Replace(r.Form.Get("username"), "\"", "", 1))
```

Go的`html/template`里面带有下面几个函数可以帮你转义

func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
func HTMLEscapeString(s string) string //转义s之后返回结果字符串
func HTMLEscaper(args ...interface{}) string //支持多个参数一起转义，返回结果字符串

## 防止多次表单提交

```html
<input type="checkbox" name="interest" value="football">足球
<input type="checkbox" name="interest" value="basketball">篮球
<input type="checkbox" name="interest" value="tennis">网球	
用户名:<input type="text" name="username">
密码:<input type="password" name="password">
<input type="hidden" name="token" value="{{.}}">
<input type="submit" value="登陆">
```
```go
func login(w http.ResponseWriter, r *http.Request) {
		fmt.Println("method:", r.Method) //获取请求的方法
		if r.Method == "GET" {
			crutime := time.Now().Unix()
			h := md5.New()
			io.WriteString(h, strconv.FormatInt(crutime, 10))
			token := fmt.Sprintf("%x", h.Sum(nil))

			t, _ := template.ParseFiles("login.gtpl")
			t.Execute(w, token)
		} else {
			//请求的是登录数据，那么执行登录的逻辑判断
			r.ParseForm()
			token := r.Form.Get("token")
			if token != "" {
				//验证token的合法性
			} else {
				//不存在token报错
			}
			fmt.Println("username length:", len(r.Form["username"][0]))
			fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
			fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
			template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
		}
	}
```

普通的驱动支持database/sql 或自定义  支持官方database/sql的应该为首选，因为迁移几乎不需要改代码。
而有些orm框架，比如beego-orm 或 gorm(推荐) 以面向对象的方式进行。

nosql
redis  则推荐使用redigo
mongodb 推荐 mgo


