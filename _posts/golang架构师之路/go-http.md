## 强大的net包
```go
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    paths := os.Args[1:]
    if len(paths) == 0 {
        return
    }
    for _,path := range paths {
          resp,err := http.Get(path)
          if err!=nil {
              fmt.Fprintf(os.Stderr,"%v\n",err)
              os.Exit(1)
          }
          content,err := ioutil.readAll(resp)
          resp.Close()
          if err!=nil {
              fmt.Fprintf(os.Stderr,"%v\n",err)
              os.Exit(1)
          }

          fmt.printf("%s",content)
    }
}
```

```go
package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"log"
)

func main() {
	paths := os.Args[1:]
	if len(paths) == 0 {
		return
	}
	for _, path := range paths {
		resp, err := http.Get(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		//content, err := ioutil.ReadAll(resp.Body)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatal(err)
		}
		
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

	}
}

```

```go
package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"log"
	"strings"
)

func main() {
	var prefix string = "http://"
	paths := os.Args[1:]
	if len(paths) == 0 {
		return
	}

	for _, path := range paths {
		if !strings.HasPrefix(path,prefix) {
			path = prefix+path
		}

		resp, err := http.Get(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Println(resp.StatusCode)
		//content, err := ioutil.ReadAll(resp.Body)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatal(err)
		}

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

	}
}

```

## 并行抓取url内容

```go
package main

import (
	"time"
	"fmt"
	"os"
	"net/http"
	"io"
	"io/ioutil"
	"strings"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	const prefix = "http://"
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed", time.Since(start).Seconds())
}
func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s:%v", url, err)
		return
	}

	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%v\t%v", seconds, nbytes, url)
}

```

```go
package main

import (
	"time"
	"fmt"
	"os"
	"net/http"
	"io"
	"io/ioutil"
	"strings"
)

func main() {
	start := time.Now()

	const prefix = "http://"
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		fetch(url)
	}

	fmt.Printf("%.2fs elapsed", time.Since(start).Seconds())
}
func fetch(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		//fmt.Sprintf("err:%v\n",err)
		fmt.Printf("err:%v\n", err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}

	seconds := time.Since(start).Seconds()
	fmt.Printf("%.2fs\t%v\t%s\n", seconds, nbytes, url)
}

```

goroutine是一种函数的并发执行方式,而channel是用来在goroutine之间进行参数传递。 main函数本身也运行在一个goroutine中,而go function则表示创建一个新的goroutine,并在 这个新的goroutine中执行这个函数。 


```go
☁  src  go run run.go www.baidu.com www.163.com www.infoq.com
Get http://www.163.com: dial tcp: lookup www.163.com: no such host
0.03s	101950	http://www.baidu.com
3.25s	276081	http://www.infoq.com
3.25s elapsed%                                                                                                

☁  src  go run run.go www.baidu.com www.163.com www.infoq.com
0.05s	102232	http://www.baidu.com
err:Get http://www.163.com: dial tcp: lookup www.163.com: no such host
5.38s	276084	http://www.infoq.com
18.67s elapsed%                                                                                               
```


## 简单的web服务

```go
package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("listen at localhost:9000")
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

```

```go
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 20.
//!+

// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	// mu.Lock()
	if r.URL.Path=="/count"||r.URL.Path=="/favicon.ico"{ // 在chrome浏览器中需要加上，safari不用加也正常
		// return
		goto Label
	}
	mu.Lock()
	count++
	fmt.Println("-->:",count)
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	Label:
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

//!-

```

打印http请求参数
```go
method, url, proto := r.Method, r.URL, r.Proto
fmt.Fprintf(w, "%s %s %s\n", method, url, proto)
for key, value := range r.Header {
	fmt.Fprintf(w, "Header[%q]=%q\n,", key, value)
}

fmt.Fprintf(w,"Host:%q\n",r.Host)
fmt.Fprintf(w,"RemoteAddr=%q\n",r.RemoteAddr)
if err:=r.ParseForm();err!=nil {
	log.Printf("%v\n",err)
}

for key, value := range r.Form {
	fmt.Fprintf(w,"Form[%q]=%q\n",key,value)
}
```