# 一、 docker快速入门
## docker国内的一些私有云公司
方便部署和维护
## 公有云对docker的支持

# 具体内容 三部分
1. docker基础
2. docker持续集成
3. docker持续部署 docker集群监控和日志管理
4. docker与mesos(分布式计算框架，调度)

# 初识docker
docker的历史发展
以linux为主，没有再windows上的生产级别应用

docker的主要目标: 通过对应用组件的封装、分发、部署、运行等生命周期的管理，达到应用级别的一次封装，到处运行。

应用组件:不限于web应用，数据库服务，甚至是一个操作系统编译器。

# why docker

1. 环境隔离
2. 更快速的交付部署
3. 更高效的资源利用
4. 更易迁移拓展 (无状态的应用)
5. 更简单的更新管理

# 虚拟化与docker
虚拟化是资源管理技术。 
系统虚拟化
guestos之间完全隔离

容器虚拟化没有guestOs一层
namespace和cgroups

# docker基础操作
目标: 利用容器创建一个可以分发的应用

## 1. 创建模板文件

```sh
FROM ubuntu
MATAINER <feng-qichao@qq.com>

CMD echo "Hello World"
```

##  2. 利用模板文件创建镜像

```
docker build -t="fqc/myubuntu01" .
```
## 3. 上传镜像

- 申请dockerhub 账号
hub.docker.com  存储，分发

执行docker login
```sh
$ docker login
```

- docker push fqc/myubuntu01


## 4. 下载镜像
如果本地没有该镜像，需要从远程仓库下载

可以先删除掉本地已有的镜像(为了演示)
```sh
docker rmi fqc/myubuntu01
```


```sh
docker pull fqc/myubuntu01
```

由于dokcerhub架设在国外，所以国内最好特殊处理下。
## 5. 运行容器

```sh
docker run --name myubuntu01 fqc/myubuntu01
```
## 6. 查看结果




# 二、docker的核心概念和安装
## 安装
操作系统: 64bit linux （centos/ubuntu）
docker版本: 1.12.3（docker rancher版本 本人的考虑和rancher版本的兼容）
内核版本: >=3.10.0
![](media/15110704930739.jpg)


验证安装是否成功
```
sudo docker run hello-world
```
默认docker启动会有socks文件
ps -ef|grep docker


##  docker配置
1. 新建docker组

```sh
sudo groupadd docker
```

2. 将当前用户添加到docker组

```sh
sudo usermod -aG docker mydocker

sudo usermod -aG docker fengqichao
```

3. 退出当前用户重新登录
`exit`
如果没有成功，还可以最终执行下
`service docker restart`
4. 验证docker是否安装成功
`docker run hello-world`

5. 设置docker服务开机自动启动
`sudo chkconfig docker on`

[配置docker服务开机自动启动](http://openskill.cn/article/530)
## docker核心概念-镜像
### 概念
镜像是一个只读的模板
镜像可以用来创建docker容器

###给镜像打标签
docker tag imageId iamge:tag

```
docker tag 2f3d1k... fqc/myubuntu01:test
```
## docker核心概念-容器
从镜像创建的运行实例
启动/开始/停止/删除
每个容器都是相互隔离的、保证安全的平台

**可以把容器看做是一个简易版的linux和运行在其中的应用程序。**
注意: 镜像是只读的，容器在启动的时候创建一层可写层作为最上层。


## docker核心概念-仓库
**仓库是集中存放镜像文件的场所**

有时候会把仓库和仓库注册服务器(Registry)混为一谈，并不严格区分。 实际上，**仓库注册服务器往往存放着多个仓库，每个仓库又包含了多个镜像，每个镜像有不同的标签(tag)**。


通过docker registry可以管理多个仓库，通过仓库可以管理镜像的不同标签版本，类似git。

# 三、 docker镜像的操作
## 01. 获取镜像

**镜像是docker运行容器的前提。**

`docker pull <域名>/<namespace>/<repo>:<tag>`


```sh
docker pull ubuntu 
```
id 镜像层的id，根据aufs来的

不加域名是dockerhub官方下载，命名空间 docker官方的可以省略，tag省略则为latest

## 02. 查看镜像列表

`docker images`

## 03. 查看镜像详细信息及格式化 过滤技巧 -f
`docker inspect <image id>`

格式化 过滤信息
`docker inspect -f {{.Os}} 2fa9.....`
docker inspect命令返回的是一个json的格式消息，如果我们只要其中的一项或几项内容时，可以通过-f参数来指定。 Image_id可以使用前几位代替，保障id输入是唯一的即可。

## 04.docker search <image_name>
`docker search nginx`

## 05. 删除镜像
### `docker rmi <image>:<tag>`

### 注意 
**当同一个镜像拥有多个标签，docker rmi只是删除该镜像的多个标签中的指定标签而已，而不影响镜像文件**。

**如果只有一个tag，那么将会删除该镜像（被物理删除）**。

**当有该镜像创建的容器存在时，镜像文件默认是无法被删除的。**

有运行的容器在使用该镜像的时候，默认删除不了。
`docker run ubuntu echo "hello world"`

### 强制删除镜像(不建议)
`docker rmi -f 1fa9....`

如果使用该镜像的容器比较多，将会打乱容器和镜像之间的管理，强烈建议不要强制删除。很难再显示查找出来容器使用了哪个镜像。
注意: **强制删除镜像不会删除容器**
删除了镜像，容器只会显示imageId了，不会再显示容器名称。

## 06. 创建镜像
### 创建镜像的方法有三种: 
#### 1. **基于**已有的镜像的**容器**创建
#### 2. 基于本地模板的导入
#### 3. 基于Dockerfile创建

基于已有镜像的容器创建
`docker commit <options> <container_id> <repository:tag>`

参数说明:
-a, --author 作者信息
-m, --message 提交信息
-p, --pause=true 提交时暂停容器运行


### 基于已有镜像的容器创建

```sh
docker pull ubuntu
docker run -it ubuntu bash  #将容器的标准输出绑定到终端 并交互式执行bash
touch test.txt
echo "hello" > test.txt
cat test.txt
exit
```
极简的镜像，没有vim，进行手工安装，之后我们会讲解 如何合理使用docker镜像。



```sh
dokcer ps -a
docker commit -a 'fqc' -m "add test.txt" 容器id fqc/ubuntu_test
docker iamges
```

## 07. 迁出镜像和载入镜像

### 迁出
`docker save - o <image>.tar <image>:<tag>`

可以使用docker save命令来迁出镜像，其中image可以为标签或id

参数说明: 
-o: 镜像存储压缩后的文件名称
### 载入镜像
`docker load --input <image>.tar `  
或者  
`docker load < <image>.tar`  
使用docker rmi命令可以删除镜像，其中image可以为标签或id

这将导入镜像及相关的元数据信息(包括标签等),可以使用docker images命令进行查看。


验证执行容器看是否有改变的内容

```sh
docker run -it imageId bash
cat test.txt
```
## 08. 上传镜像
`docker push <域名>/<namespace>/<repo>:<tag>`

可以使用 docker push命令上传镜像到仓库，默认上传到DockeHub官方仓库(需要登录，否则提示authentication required)。


## 总结 
镜像是容器的前提，也是容器的资源，积累好的镜像及好的制作方法，在使用容器时，才能事半功倍。
# 四、容器的常用操作
1. 创建容器
2. 终止容器
3. 进入容器
4. 删除容器
5. 导入和导出容器

## 1. 创建容器
docker的容器十分轻量级，用户可以随时创建和删除容器

### 新建容器 docker create
`docker create --name test_create -it ubuntu`

说明: docker create命令创建的容器处于停止状态（status 为 created），可以使用`docker start`命令启动它。
### 新建并启动容器 docker run
#### useage
`docker run ubuntu /bin/echo "hello world"`

说明:  等价于 docker create && docker start

####  docker run背后运行的内容

1. 检查本地是否有存在的指定的镜像，不存在就从公有仓库下载
2. 利用本地镜像创建并启动一个容器
3. 分配一个文件系统，并在只读的镜像层外面挂载一层可读写层
4. 从宿主机配置的网桥接口桥接一个虚拟接口到容器中去
5. 从地址池配置一个ip地址给容器
6. 执行用户的指定的用户程序
7. 执行完毕后容器被终止



```sh
docker run -it -d --name test_network ubuntu bash
docker exec -it test_network bash
cat /etc/hosts
```


执行完成后容器被终止
```sh
docker run -it --name test_finish ubuntu echo "hello world"
docker ps -a
```

### 交互式运行 -i -t 参数的含义


```sh
docker run -i -t ubuntu /bin/bash
```

-i: 让容器的标准输入保持打开
-t: 让docker分配一个伪终端并绑定到容器的标准输入上

在交互模式下，用户可以通过所创建的终端来输入命令，**exit命令退出容器。(类似ssh)**
**退出后，容器自动处于终止状态。**
### 后台守护式运行 -d参数的含义
更多的时候，需要让docker容器运行在后台以守护的形式运行。用户可以通过添加-d参数来实现。


```sh
docker run -d ubuntu /bin/sh -c "while true; do echo hello world; sleep 1; done"
```
### 日志
补充: 查看日志 `docker logs [-f] <container_id>`
注意: `-f` 参数滚动查看(`--follow         Follow log output`)  

```
docker logs  --help
Usage:	docker logs [OPTIONS] CONTAINER

Fetch the logs of a container

Options:
      --details        Show extra details provided to logs
  -f, --follow         Follow log output
      --help           Print usage
      --since string   Show logs since timestamp
      --tail string    Number of lines to show from the end of the logs (default "all")
  -t, --timestamps     Show timestamps
```

![](media/15110778340203.jpg)



## 2. 终止容器

```sh
docker stop <container_id>
```

注意: **当容器中的应用终结时，容器也会自动停止。**
查看终止的容器: `docker ps -a`
查看运行的容器: `docker ps`
重启容器: `docker restart <container_id>`

## 3. 进入容器
使用`-d`参数时，容器启动后会进入后台，用户无法查看到容器中的信息。
`docker exec <options> <container_id> <command>`

Exec可以直接在容器内部运行命令。

进入容器: `docker exec -i -t <container_id> bash`
之后再执行exit的时候，并不会影响容器的运行。


```sh
docker run -it ubuntu bash
$ exit
exit
```


```sh
docker run -d --name test_it ubuntu /bin/bash echo "hello world"
docker exec -it test_it bash 
$ ps -ef 
$ exit
docker pes # 依旧运行
```

## 4. 删除容器
删除终止状态的容器
`docker rm `

删除正在运行的容器
`docker rm -f` 
## 5. 导入和导出容器
### docker export 导出容器
使用`docker export`将一个已经创建的容器导出到一个文件中，不过容器是否已经处于运行状态。
`docker export <container_id> > file.tar`
或
`docker export test_id > test.tar`

### docker import 导入容器
使用docker import将之前导出的文件导入，**成为镜像**

cat test.tar | docker import - fqc/testimport:latest


###总结及建议:
 容器是直接提供拥有服务的组件，也是docker实现快速的启停和高效服务性能的基础。


在生产环境中，因为容器自身的轻量级的特性，推荐在使用容器前，引入高可用的特性，例如HA_proxy，当容器出现故障的时候，可以快速切换到其他容器，还可以实现容器的自动重启。掌握一些shell，还可以写一些容器自动化的脚本，可以涉及到一些镜像的服务，服务的管理。



## 五、 仓库
### 公有仓库
[dockerhub](dockerhub.com)

注意search、pull是不需要登录的，而push则是需要登录校验的
```sh
☁  ~  docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username (fengqc): fengqc
Password:
Login Succeeded
☁  ~
```
### 私有仓库


## 六、数据管理
### 将宿主目录挂载到容器内部目录

本地目录和容器目录建立了映射关系，双方的改变都会影响到到对方

数据卷绑定了固定的目录，和容器内的目录有映射关系，当容器内部或外部改变了目录的内容时，该目录就会被改变。

有何作用？生产中一般怎么使用？
如果我们生产中挂载的是数据库mysql的服务，我们可以把容器内部产生的数据和mysql本身这个应用分开，当容器挂了的时候，数据还在宿主机上存留，当再启动的时候，还可以把该目录挂载回去，来启动mysql，数据是不会丢失的。


### 数据卷容器
数据卷容器用于用户**需要在容器间共享一些持续更新的数据**，数据卷容器**专门提供数据卷供其他容器挂载使用**。


### 利用数据卷迁移容器

## 七、docker 网络

### 容器对外服务
-p  -P

docker port my_mysql 3306

### 容器间相互通信
容器的连接(link)是除了端口映射外的另一种可以与容器中应用进行交互的方式。
使用 --link 参数可以让容器之间安全的进行交互。

目的是在于不想暴露给其他ip或宿主机，只想给指定的服务应用。
通过docker自己的网卡实现通信

## 八、 dockerfile

dockerfile是容器非常核心的部分基础的部分。 通过dockerfile可以固化镜像的构建过程，通过dockerfile可以实现自动化构建。


### . Dockerfile核心指令集


利用Dockerfile创建镜像

Run为镜像的操作指令 (中间操作的过程 安装些软件或变更些文件)
CMD为容器启动时的指令

```sh

```

expose 暴露端口  

```sh
EXPOSE 80 443 # 容器暴露多个端口
```

env设置环境变量


```sh
ENV NGINX_BIN_PATH /usr/sbin/
```

add src dest
将复制指定的src到容器的dest。 其中src可以是dockerfile所在目录的一个相对路径；也可以是url，也可以是tar文件(自动解压)

```sh
add hello.txt /hello.txt
```

volume 将内部的某些盘显示的挂载出来


```sh
VOLUME /  # 我们可以使用-v命令将 将容器的目录通过宿主机管理
```

USER 指定容器运行时的用户名，后续的Run也会指定该用户。

WorkDir 说白了就是切换工作目录，后续的Run CMD ENTRYPOINT都会在此目录下执行

copy src dest
复制本地主机的src到容器的dest，目标不存在时，会自动创建。
当使用本地目录为源目录时，推荐使用copy
其实本地时和add是等价的


CMD 
ENTRYPOINT 

如果命令带参数，一般推荐使用CMD命令。 cmd命令可以在通过docker启动时指定命令来覆盖，易于调试。
entrypoint命令不会被启动命令覆盖，不太容易调试
他俩在dockerfile中如果分别有多条，都会被覆盖掉，只保留最后一条执行。(它俩都会各自执行最后一条的，)

3. Dockerfile最佳实践

### 错误定位
可以根据出错所在步骤，后台运行该容器进入后，手动执行下一步出错的层

使用官方docker镜像 下载自己喜欢的软件，需要 apt-get update && apt-get install xx. 
重要的是apt-get update 需要先执行以下。
centos yum update && yum install xxx 一样

netstat -nlp|grep 80
curl 没有浏览器的时候 我们就用curl啦
### 最佳实战分享


## docker之持续集成
git + jenkins + docker 持续集成平台
docker在持续集成中的作用: docker提供代码编译、打包、测试的相关环境。  
优势1. 环境可以是任意版本 2.节省空间 3.环境相对隔离




