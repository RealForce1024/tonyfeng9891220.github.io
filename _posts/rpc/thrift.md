## 一、安装
**mac平台直接推荐使用方式二**
### 方式一: 官网方式
最好使用相应的平台，比如mac平台就相应的参考  
[参看](https://thrift.apache.org/tutorial/)  

1. 下载地址 [0.9.x系列](http://archive.apache.org/dist/thrift/0.9.3/)  
2. `$ cd $thrift_home`  
`$ ./configure && make` 

可能遇到的问题  
```sh
checking for bison version >= 2.5... no
configure: error: Bison version 2.5 or higher must be installed on the system!
```
[解决方案](https://stackoverflow.com/questions/31805431/how-to-install-bison-on-mac-osx)  

```sh
See here. You can install with brew:
brew install bison
and then use:
brew link bison --force
Don't forget to unlink it if necessary (brew unlink bison).
```

手动安装还可以参考:  
[Mac OS X 下搭建thrift环境(手动安装)](http://www.cnblogs.com/smartloli/p/4220545.html)  

### 方式二: mac平台最简单的brew方式
```sh
☁  ~  sudo brew update
☁  ~  brew install thrift
==> Installing dependencies for thrift: boost
==> Installing thrift dependency: boost
==> Downloading https://homebrew.bintray.com/bottles/boost-1.61.0.el_capitan.bottle.tar.gz
######################################################################## 100.0%
==> Pouring boost-1.61.0.el_capitan.bottle.tar.gz
🍺  /usr/local/Cellar/boost/1.61.0: 12,259 files, 438.4M
==> Installing thrift
==> Downloading https://homebrew.bintray.com/bottles/thrift-0.9.3.el_capitan.bottle.tar.gz
######################################################################## 100.0%
==> Pouring thrift-0.9.3.el_capitan.bottle.tar.gz
==> Caveats
To install Ruby binding:
  gem install thrift

To install PHP extension for e.g. PHP 5.5:
  brew install homebrew/php/php55-thrift
==> Summary
🍺  /usr/local/Cellar/thrift/0.9.3: 94 files, 5.3M
☁  ~  thrift -version
Thrift version 0.9.3
```

总结：
从官网直接下载thrift.tar.gz本地安装，会出错。如果出错，最好的解决方式就是brew在线下载安装。
手动安装步骤，需要先下载安装boost，然后再下载thrift安装。中间编译过程会很漫长，即使mac机器...
所以还是首推thrift自动安装。  


## 二、Thrift Guide
[董西成 thrift介绍](http://dongxicheng.org/search-engine/thrift-guide/)   
[Apache Thrift - 可伸缩的跨语言服务开发框架](https://www.ibm.com/developerworks/cn/java/j-lo-apachethrift/)    
[Thrift RPC 使用指南实战(附golang&PHP代码)](http://blog.csdn.net/liuxinmingcode/article/details/45696237)    
[小探python-thrift通信框架的设计](http://xiaorui.cc/2016/07/24/%E5%B0%8F%E6%8E%A2python-thrift%E9%80%9A%E4%BF%A1%E6%A1%86%E6%9E%B6%E7%9A%84%E8%AE%BE%E8%AE%A1/)  
[Golang开发Thrift接口](http://blog.cyeam.com/golang/2014/07/22/go_thrift)  
[golang:实现thrift的client端协程安全](http://dev.cmcm.com/archives/162)  
[Golang 1.4 net/rpc server源码解析](http://dev.cmcm.com/archives/324)  
[Golang 1.3 sync.Mutex 源码解析](http://dev.cmcm.com/archives/22)  
## hello world
[thrift with go and java](https://my.oschina.net/qinerg/blog/165285)  
[apache官网go thrift](https://thrift.apache.org/tutorial/go)  
[hello thrift with java](http://blog.zhengdong.me/2012/05/10/hello-world-by-thrift-using-java/)  
[Golang RPC 之 Thrift](http://www.jianshu.com/p/a58665a38022)  
[服务架构的一些理解以及TTL](http://kaimingwan.com/post/wei-fu-wu/fu-wu-kuang-jia-de-ji-chong-fu-wu-diao-yong-xing-shi)  
[Golang通过Thrift框架完美实现跨语言调用](http://www.cnblogs.com/shihao/p/3347537.html?utm_source=debugrun&utm_medium=referral)  
[go和python测试thrift](http://www.codexiu.cn/python/blog/1026/)  
## 深入
[开源RPC（gRPC/Thrift）框架性能评测](http://www.eit.name/blog/read.php?566)  很赞的一篇文章  

[几种Go序列化库的性能比较](http://colobu.com/2015/09/28/Golang-Serializer-Benchmark-Comparison/)  
[好几篇关于go的优秀文章及grpc的使用](http://colobu.com/)  
[rpcx框架](https://github.com/smallnest/rpcx)  

[如何实现支持数亿用户的长连消息系统 | Golang高并发案例](http://chuansong.me/n/1641640)  
[golang socket编程](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/08.1.md)  
[golang学习室](https://www.kancloud.cn/digest/batu-go#/catalog)  
[golang thrift 总结一下网络上的一些坑](http://www.cnblogs.com/ka200812/p/5865213.html)  
[golang thrift 源码分析，服务器和客户端究竟是如何工作的](http://www.cnblogs.com/ka200812/p/5868172.html)  
[Golang、Php、Python、Java基于Thrift0.9.1实现跨语言调用](http://idoall.org/blog/post/lion/7)  
[go web评测](http://colobu.com/2017/04/07/go-webframework-benchmark-2017-Spring/)  
[Apache thrift between golang and node.js 台湾某it工作室](http://idanbean.idanbird.net/2016/02/07/apache-thrift-between-golang-and-node-js/)  
[thrift-go（golang）Server端笔记](http://www.cnblogs.com/lijunhao/p/5976733.html)  
[常规rpc通讯过程](http://www.cnblogs.com/lijunhao/p/6137920.html)  
[thrift-missing-guide/](https://diwakergupta.github.io/thrift-missing-guide/)  
[GOLANG使用THRIFT的鉴权和多路复用](http://blog.molibei.com/archives/213)  
[深度解析gRPC以及京东分布式服务框架跨语言实战](http://www.10tiao.com/html/164/201702/2652898208/1.html)  

[go 传送门](http://chuansong.me/search?q=golang)   
很多tcp ，rpc ，100亿次请求的高可用高性能方面的文章
[如果你用Go，不要忘了vet](http://chuansong.me/n/1687532651414)  


[tim yang 微博架构师](https://timyang.net/)  
[Thrift and Protocol Buffers performance in Java[09年]](https://timyang.net/programming/thrift-protocol-buffers-performance-java/)  