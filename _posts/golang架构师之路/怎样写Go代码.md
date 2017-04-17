怎样写Go代码?
## 一、下载golang
点击 [官网下载golang sdk](1)    
根据不同系统，官网下载链接会选择相应的平台进行链接跳转，也可手动选择需要的平台安装包。  

     

## 二、安装golang
双击可执行程序一步步next，大多数都可以完成安装。  
除了linux解压tar.gz方式 `tar -zxvf go-versionxx.tar.gz`，以及源码编译方式(使用的频率极少，不分散精力在这，需要时再做)  


## Go代码组织结构详解
**概述**  
* Go开发者通常将所有Go代码保存在单个工作空间中(一般工作区更适合)
* 工作区包含多个代码版本仓库，例如git
* 每个代码仓库包含一个或多个包
* 每个包由单个目录中的一个或多个Go源文件组成
* 包路径确定其导入路径
注意Go的工作空间组织方式不同于那些每个项目都有单独的工作空间并且每个工作空间与版本控制库紧密相关的其他编程环境。比如Java一个项目一个工作空间也是一个版本库（最佳实践）。而Go工作空间则是多个版本控制库，每个版本控制库对应一个项目(单个工作空间多个项目代码库)  
```
java  
project->workspace->git  
1->1->1   

go 
workspace->git->project
1->n->n
```

### 工作区
工作区(workspace)是一个目录层次结构，其根目录有三个目录构成:  
* `src`包含Go源文件
* `pkg`包含包对象
* `bin`包含可执行命令
go工具编译构建`src`源软件包并将生成的二进制文件安装到`pkg`和`bin`目录  
`src`子目录通常包含多个代码版本控制仓库，用于追踪一个或多个源码包的开发。
为了让你了解实践中的工作区是什么样子，演示个例子:
```sh
bin/
    hello                          # command executable
    outyet                         # command executable
pkg/
    linux_amd64/
        github.com/golang/example/
            stringutil.a           # package object
src/
    github.com/golang/example/
        .git/                      # Git repository metadata
	hello/
	    hello.go               # command source
	outyet/
	    main.go                # command source
	    main_test.go           # test source
	stringutil/
	    reverse.go             # package source
	    reverse_test.go        # test source
    golang.org/x/image/
        .git/                      # Git repository metadata
	bmp/
	    reader.go              # package source
	    writer.go              # package source
    ... (many more repositories and packages omitted) ...
```
以上展示了包含两个代码仓库(example和image)的工作区。example仓库包含两个命令(hello和outyet)和一个类库(stringutil)。image仓库包含bmp包和几个其他的。  

典型的工作区包含着由许多包和命令的源码仓库组成。大多数Go开发者将Go源码和依赖保存在单个工作区中。  

命令和库由不同类型的源码包构建。下面将讨论[todo]。
### GOPATH默认环境变量
**GOPATH环境变量特指工作区的路径。**  
默认情况下在系统家目录(${home})下
* Unix $HOME/go
* Plan9 $home/go
* Windows %USERPROFILE%\go (通常是c:\Users\$UserName\go)
如果你想要在非默认工作区的路径工作，那么你需要设置GOPATH环境变量指向那个目录(另一种常见的设置比如是设置GOPATH=$HOME)。**注意GOPATH一定不能是Go SDK的安装目录。**

命令 `go env GOPATH`打印当前的`GOPATH`。  
如果没有设置其他位置的环境变量它将打印默认位置`$HOME/go`。
```sh
$ go env GOPATH
☁  golang架构师之路 [master] ⚡ go env GOPATH                                                                    [master↑1|✚3…
/Users/fqc/github/golang_sidepro
```

为方便起见，将工作区的`bin`子目录添加到系统环境变量`PATH`中:  
```sh
$ export PATH=$PATH:$(go env GOPATH)/bin
```
注意:不是`go env`<font color='red'>`$GOPATH`</font>

另外为了简洁起见，本文档的其他脚本使用`$GOPATH`替代`$(go env GOPATH)`。如果你没有设置`GOPATH`却想让脚本正常运行，可以通过$HOME/go替代$(go env GOPATH)命令或者运行
```sh
export GOPATH=$(go env GOPATH)
```
上述的命令的含义实质还是将$(go env GOPATH)赋值给GOPATH变量，以后使用$GOPATH替代$(go env GOPATH)。  
学习更多的GOPATH环境变量，可以通过`go help gopath`命令。
```
☁  golang_sidepro [master] ⚡ go help gopath                                                                    [master↑2|✚66…
The Go path is used to resolve import statements.
It is implemented by and documented in the go/build package.

The GOPATH environment variable lists places to look for Go code.
On Unix, the value is a colon-separated string.
On Windows, the value is a semicolon-separated string.
On Plan 9, the value is a list.

If the environment variable is unset, GOPATH defaults
to a subdirectory named "go" in the user's home directory
($HOME/go on Unix, %USERPROFILE%\go on Windows),
unless that directory holds a Go distribution.
Run "go env GOPATH" to see the current GOPATH.

......

being checked out for the first time by 'go get': those are always
placed in the main GOPATH, never in a vendor subtree.

See https://golang.org/s/go15vendor for details.
```
注意上述都是在默认情况下的设置。


```sh
☁  golang架构师之路 [master] ⚡ cd $GOPATH                                                                       [master↑1|✚3…
☁  golang_sidepro [master] ⚡ ll                                                                                [master↑2|✚66…
total 48
-rw-r--r--   1 fqc  staff    19K Jun  1  2016 LICENSE
-rw-r--r--   1 fqc  staff    56B Jun  1  2016 README.md
drwxr-xr-x   9 fqc  staff   306B Apr 16 22:39 bin
drwxr-xr-x   3 fqc  staff   102B Apr 14 15:20 pkg
drwxr-xr-x  20 fqc  staff   680B Apr 14 14:58 src
```


### 自定义GOPATH环境变量
GOPATH可以是你操作系统的任意目录(除了Go SDK目录)。在Unix示例中我们将它设置到`$HOME/work`。注意GOPATH一定不能和GO SDK的安装目录。另一个常用设置是GOPATH=$HOME
* Unix系统
    * Bash
    * Zsh
* Windows系统

#### Bash 
编辑`.bash_profile`，添加如下一行:
```sh
export GOPATH=$HOME/work
```
保存并退出编辑器。然后刷新`~/.bash_profile`。
```sh
source ~/.bash_profile
```
#### Zsh
编辑`~/.zshrc`文件，添加如下一行:
```sh
export GOPATH=$HOME/work
```
保存并退出编辑器。然后刷新`~/.zshrc`

### Windows
工作空间你可以选择自己喜欢的，但是我们将采用`C:work`作为示例演示。注意GOPATH不能和你的`GO SDK`安装路径一样
1. 创建`C:\work`目录
2. 选择`开始`->`控制面板`->`系统和安全`->`系统`->`高级系统设置`->`环境变量`
3. 创建用户变量
4. 变量名一栏输入`GOPATH`
5. 变量值一栏输入`C:\work`
6. 点击完成


总结:
记住GOPATH并不是golang的安装路径，而是工作空间WorkSpace路径。
GOPATH分两种情况
* 默认不设置 ${home}/go
* 自定义

查看go sdk的安装路径
```sh
☁  ~  which go
/usr/local/go/bin/go
```

## 导入路径 Import path
`import path`是唯一标识包的字符串。包的导入路径对应其在工作区或远程代码库的位置(如下所述)。  

标准库的代码包使用较短的导入路径即可，例如"fmt"或"net/http"。对于你自己写的包，必须选择不太可能与标准库或未来添加的外部库相冲突的一个基本路径。  

注意在你可以成功构建你的代码之前不要发布代码到远程仓库。组织好你的代码就像某天你将发布它一样是个好的习惯。实际上你可以选择任意的路径，只要它在针对工作区，标准库或更大的Go生态圈是为唯一的。  

我们将采用github.com/user作为项目根路径。在工作区存放源码的文件夹下创建该目录：
```sh
mkdir -p $GOPATH/src/github.com/user
```
### 第一个go程序
为了编译和运行一个简单的程序，首先我们需要选择包路径(我们这里采用github.com/user/hello)，在工作区创建相应的目录:
```sh
mkdir $GOPATH/src/github.com/user/hello
```
然后，在刚创建的目录里新建文件`hello.go`，包含如下代码:
```go
package main
import "fmt"
func main(){
    fmt.Print("hello world.\n")
}
```
现在你可以使用go工具构建并安装该程序:
```sh
go intall github.com/user/hello
```
注意你可以在系统的任意位置运行上述命令。`go tool`根据配置的环境变量`GOPATH`下寻找github.com/user/hello包。

错误示例:
```sh
☁  ~  go install $GOPATH/src/github.com/user/hello
can't load package: package /Users/fqc/work/src/github.com/user/hello: import "/Users/fqc/work/src/github.com/user/hello": cannot import absolute path
☁  ~  go install $GOPATH/src/github.com/user/hello/hello.go
go install: no install location for .go files listed on command line (GOBIN not set)
```
你也可以忽略包路径如果你在包路径下执行`go install`:
```sh
$ cd $GOPATH/src/github.com/user/hello
$ go install
```
上述命令构建`hello` command，生成可执行二进制文件。然后将二进制文件安装到工作区的bin目录为hello(或在Windows下，hello.exe)。在我们的示例中，将为`$GOPATH/bin/hello`，也就是`$HOME/work/bin/hello`。

go tool只会在发生错误时打印输出，因此如果这些命令不产生输入，则说明它们已成功执行。  

你可以通过输入下面的全路径运行程序:
```sh
$ $GOPATH/bin/hello
Hello, wolrd.
```
或者你可以更加聪明便捷的方式，将$GOPATH/bin将入到PATH中，以后只需要输入二进制名称即可:
```sh
$ hello
Hello, world.
```

如果你正在使用源码控制系统，现在是个初始化仓库的绝佳时机，添加文件并提交第一次修改.另外，这个步骤是可选的:你不是非要使用代码版本控制来写Go代码(但是最好用上)。  
```sh
$ cd $GOPATH/src/github.com/user/hello
$ git init
Initialized empty Git repository in /home/user/work/src/github.com/user/hello/.git/
$ git add hello.go
$ git commit -m "initial commit"
[master (root-commit) 0b4507d] initial commit
 1 file changed, 1 insertion(+)
  create mode 100644 hello.go
```
最后就是关联远程仓库推送代码啦，读者可以自行完成。  

参考 
[golang文档][1]  

[1]:https://golang.org/doc/code.html  


[SettingGOPATH][2]

[2]:https://github.com/golang/go/wiki/SettingGOPATH
