JSON是一种用于发送和接收结构化信息的标准。其他类似标准还有`XML`,`ASN.1`和Google的`Protocol Buffers`，这几种协议特有特色。JSON以简洁性、可读性和流行性等原因，JSON是应用最为广泛的。  

Go语言对上述标准协议都有良好的支持，标准库中encoding/json、encoding/xml、encoding/asn1等包提供支持，Protocol Buffers则由github.com/golang/protobuf包提供支持。有个利好的消息：这些包的API很相似。

encoding/json包是我们此文章的重点

Json是对js各类型的值--数字、字符串、数组、对象--的Unicode文本编码。使用go将各类型的值转为json也是编码，而json转为go类型则是解码。  

基本的json类型 数字，字符串，布尔，这些基本类型可以通过json数组和对象类型递归组合。一个Json数组是有序的值序列。
Json数组<--->Go 数组和切片
Json对象<--->Go map和struct 其中(map[string] key为string类型)

## 编码为json
go切片编码为json数组
使用json.Marshal函数，该函数简单只需要传入一个切片，即可将其编码为json数组，但是不利于阅读，可以使用json.MarshalIdent函数，增加两个参数，指定每行的输出前缀和每个层级的缩进，这样生成的格式非常利于阅读。  

```go
package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool and Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("json marshal failed: %s", err)
	}

	fmt.Printf("%q\n", data)
	fmt.Printf("%s\n", data)
	fmt.Printf("%v\n", data)

	data2, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("json marsharl indent failed: %s", err)
	}
	fmt.Printf("%s\n", data2)
}
```

输出结果:  
```go
"[{\"Title\":\"Casablanca\",\"released\":1942,\"Actors\":[\"Humphrey Bogart\",\"Ingrid Bergman\"]},{\"Title\":\"Cool and Luke\",\"released\":1967,\"color\":true,\"Actors\":[\"Paul Newman\"]},{\"Title\":\"Bullitt\",\"released\":1968,\"color\":true,\"Actors\":[\"Steve McQueen\",\"Jacqueline Bisset\"]}]"

[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool and Luke","released":1967,"color":true,"Actors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}]

[91 123 34 84 105..... 125 93]

[
	{
		"Title": "Casablanca",
		"released": 1942,
		"Actors": [
			"Humphrey Bogart",
			"Ingrid Bergman"
		]
	},
	{
		"Title": "Cool and Luke",
		"released": 1967,
		"color": true,
		"Actors": [
			"Paul Newman"
		]
	},
	{
		"Title": "Bullitt",
		"released": 1968,
		"color": true,
		"Actors": [
			"Steve McQueen",
			"Jacqueline Bisset"
		]
	}
]
```
### json tag
我们看到在Movice有两个成员Year和Color后面追加了结构体Tag，结构体Tag是成员元数据信息，是在编译阶段关联到该成员的元信息字符串:  
```go
Year int `json:"released"`
Color bool `json:"color,omitempty"`
```
结构体Tag需要注意以下几点:  
1. 首先结构体Tag是键值对序列,`key:"value"`形式存在
2. key以包名开头，key为`json`,控制使用`encoding/json`包的编解码的行为，其他包encoding/...也遵守该约定，比如使用`encoding/xml`包，则指定tag为`xml:released`
3. 值`"value"`是字符串字面值，含有`""`，所以结构体Tag的值一般是原生字符串面值的形式指定
4. Tag值中的第一个成员对应的是json字段的名称，比如将Year对应到json对象的released字段，
5. Tag值中的第二个成员则是可选的`omitempty`，其含义为Go语言结构体的成员值为空或零值的时候不生成json字段(Color为bool型，false零值则不会输出)，所以我们看到1942年的Casablanca为`Color`为`false`，果然是黑白电影，没有输出`Color`字段。  
6. tag中为`json:"-"`该字段将不输出
7. 二次编码json
ServerName2字段，json的修饰`string`说明

```json
// ServerName2 的值会进行二次JSON编码
		ServerName  string `json:"serverName"`
		ServerName2 string `json:"serverName2,string"`
		
		{"serverName":"Go \"1.0\" ","serverName2":"\"Go \\\"1.0\\\" \""}
```
8. 第三方simplejson库也非常便利(推荐,由bitly公司(url短链接平台)贡献)

## json解码
编码的逆操作为解码，是将json数据解码为Go语言的数据结构。
通过**定义合适的数据结构**我们可以很轻松的**选择来解码我们感兴趣的数据字段**。  
注意：结构体定义的成员名称一定是json数据中的字段名，并且必须定义为大写。
```go
fmt.Printf("%s\n", data)
var MyFavorites []struct {
    Title string
    Year int //对照
    released int//对照
    Released int
}

if err := json.Unmarshal(data, &MyFavorites); err != nil {
    log.Fatalf("json unmarshal failed:%s", err)
}
fmt.Println(MyFavorites)

// [{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool and Luke","released":1967,"color":true,"Actors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}]
// [{Casablanca 0 0 1942} {Cool and Luke 0 0 1967} {Bullitt 0 0 1968}]

```


复习 匿名结构体
我们看到 `var titles []struct{title string}`这样的声明，其中结构体为匿名类型。  

```go
package main

import "fmt"

func main() {
	p := struct {
		Name string
		Age int
	}{
		"kobe",24}

	fmt.Println(p)
}
// 结构体字面值，需要赋值或被其他语句使用否则
// struct{...} evalued but not used
```

json案例 检索github issues
```go
package main

import (
	"net/url"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
	"os"
	"log"
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Item
}

type Item struct {
	Title         string
	HTMLUrl       string `json:"html_url"`
	RepositoryUrl string
	Number int
	User          *User
}

type User struct {
	Login   string `user_name`
	Id      int
	HTMLUrl string `json:html_url`
}

const IssuesURL string = "https://api.github.com/search/issues/"

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func SearchIssues(items []string) (*IssuesSearchResult,error){
	q := url.QueryEscape(strings.Join(items, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil,err
	}

	if resp.StatusCode!=http.StatusOK {
		resp.Body.Close()
		return nil,fmt.Errorf("search query faield :%s",resp.Status)
	}

	var result IssuesSearchResult
	if err:= json.NewDecoder(resp.Body).Decode(&result);err!=nil{
		resp.Body.Close()
		return nil,err
	}

	resp.Body.Close()
	return &result,nil
}
./run repo:golang/go is:open json decoder
```


go echo vue https://scotch.io/tutorials/create-a-single-page-app-with-go-echo-and-vue
https://scotch.io/bar-talk/build-a-realtime-chat-server-with-go-and-websockets
https://medium.com/@kyawmyintthein


https://github.com/wangyibin/echoswg 


github.com/julienschmidt/httprouter 第三方router 路由，go需要自己实现rest

密码安全
md5+salt  方案
scrypt 专家方案

1）如果你是普通用户，那么我们建议使用LastPass进行密码存储和生成，对不同的网站使用不同的密码；
2）如果你是开发人员， 那么我们强烈建议你采用专家方案(scrypt)进行密码存储。
















