## bufio
bufio库中含io大量的读写操作。 

## 从标准输入读取数据
### Scanner
- NewScanner()
`input := bufio.NewScanner(os.Stdin)`  
- Scanner()
NewScanner方法返回新的Scanner结构，该结构的`Scan()`方法将扫描是否含有新的行内容。
- Text()
input.Text()方法将读取一行内容 
注意终端结束符为ctrl/cmd +d, 注意可能有些终端的快捷键冲突，需要避免。  
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for text, ln := range counts {
		fmt.Printf("%s\t%d\n",text,ln)
	}
}
```

```go
package main

import (
	"fmt"
)

func main() {
	var s string
	//fmt.Scanln(s)
	//fmt.Scanln(&s)
	fmt.Scanf("%s",&s) // 使用指针修改变量底层值
	//fmt.Scanf("%s",s) // 与&s的区别在于vv "" 打印的零值 而通过&s则是修改底层值。  
	//fmt.Scanf("%q",s)
	fmt.Println(s)
	fmt.Printf("%s %[1]q","hello")
}
```

## 从标准文件读入

```go
package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts) // 如果没有指定参数，使用os.Stdin系统标准输入
	} else {
		for _, file := range files {
			file, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}
	for line, n := range counts {
		fmt.Printf("%s\t%d\n", line, n)
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

```



## ioutil

```go
package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		for _, line := range strings.Split(string(data), "\n") {

			counts[line]++
		}
	}

	for line, n := range counts {
		fmt.Printf("%s\t%d\n", line, n)
	}
}

```

## 对重复行输出加上文件名

```go
package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	}

	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			fmt.Fprint(os.Stderr, "%v\n", err)
			continue
		}
		countLines(file, counts)
		file.Close()
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}

	}
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		fileName := f.Name()
		//counts[strings.Join([]string(fileName)[:], text)]++
		counts[fileName+"-->\t"+text]++
	}
}
```










