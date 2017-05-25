JSON是一种用于发送和接收结构化信息的标准。其他类似标准还有`XML`,`ASN.1`和Google的`Protocol Buffers`，这几种协议特有特色。JSON以简洁性、可读性和流行性等原因，JSON是应用最为广泛的。  

Go语言对上述标准协议都有良好的支持，标准库中encoding/json、encoding/xml、encoding/asn1等包提供支持，Protocol Buffers则由github.com/golang/protobuf包提供支持。有个利好的消息：这些包的API很相似。

encoding/json包是我们此文章的重点

Json是对js各类型的值--数字、字符串、数组、对象--的Unicode文本编码。使用go将各类型的值转为json也是编码，而json转为go类型则是解码。  

基本的json类型 数字，字符串，布尔，这些基本类型可以通过json数组和对象类型递归组合。一个Json数组是有序的值序列。
Json数组<--->Go 数组和切片
Json对象<--->Go map和struct 其中(map[string] key为string类型)

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











