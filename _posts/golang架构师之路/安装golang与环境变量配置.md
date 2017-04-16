## 一、下载golang
点击 [官网下载golang sdk][1]  
根据不同系统，官网下载链接会选择相应的平台进行链接跳转，也可手动选择需要的平台安装包。
[1]: https://golang.org/doc/install

## 二、安装golang
双击可执行程序一步步next，大多数都可以完成安装。  
除了linux解压tar.gz方式 `tar -zxvf go-versionxx.tar.gz`，以及源码编译方式(使用的频率极少，不分散精力)

## Go代码组织结构详解
### 概述
* Go语言编程者通常将他们的Go代码保存在一个工作空间中
* 一个工作空间包含多个代码版本仓库，例如git
* 每个代码仓库包含一个或多个包
* 每个包由单个目录中的一个或多个Go源文件组成
* 包路径确定其导入路径
注意Go的工作空间组织方式不同于那些每个项目都有单独的工作空间并且每个工作空间与版本控制库紧密相关的其他编程环境。比如Java一个项目一个工作空间也是一个版本库（最佳实践）。而一个Go工作空间则是多个版本控制库，每个版本控制库对应一个项目。  
java project->workspace->->git
1->1->1   
go workspace->git->project
1->n->n

## 三、设置golang环境变量
mac安装后直接会设置好环境变量。
```sh
☁  ~  which go
/usr/local/go/bin/go
```
记住GOPATH并不是golang的安装路径，而是工作空间WorkSpace路径。
GOPATH分两种情况
* 默认不设置 ${home}/go
* 自定义
### 默认golang工作空间${HOME}/go
cd /${home}/go




