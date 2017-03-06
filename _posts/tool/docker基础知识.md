docker基础知识
## docker的出现
货物损坏，运输环节、手续多、效率低下等
轮船运输货物，集装箱带来的变革，我们不需要太过关注集装箱里装的什么。
集装箱规格有统一，提高了装卸效率，转换效率。
产品交付，运维，软件升级

传统软件交付过程当中，
软件发布更新低效，无法敏捷，环境的一致性，迁移的成本高

docker出现后，以上问题有了很大程度上的解决。

## docker是什么?
本身并不是虚拟化技术，而是容器管理工具。
软件工业上的集装箱技术。
## docker可以干什么
- 快速创建环境(开发、测试、生产)

- 整体交付(运行环境+代码)
最大的特点 整体交付 开发运维测试 产品应用
环境一致性保障
更好的完成devops

## docker的安装
内核要求：3.10.0以上。 `uname -r`
安装命令:`curl -sSL https://get.docker.com/ | sh` `which curl`
查看版本:`docker -v`
启动docker: `systemctl start docker`
将docker加入到开机启动:`systemctl docker enable`
确认是否启动:`ps -aux | grep docker`
查看docker是否有安装过镜像或者创建过镜像。
`docker ps -a`

## docker镜像
1. 什么是docker镜像
docker之所以这么火，就是因为提出了镜像的概念。
集装箱未出现之前，码头上可以看到很多的搬运工人。
之后呢，则是更多的塔吊。
搬运模式更加单一且高效，
货物打包之后，可以防止物品之间相互影响。
如果到了另外一个码头，需要准运的时候，只需要将集装箱放到另一艘轮船上即可。
可以保证货物的整体搬迁，并不会损坏里面的货物本身。

docker镜像在IT行业中也像集装箱之于码头的重要变革一样，非常重要的角色。
**docker镜像就是将业务代码和运行环境进行整体的打包。**
2. 如何创建docker镜像
基础镜像从公有仓库直接拉取就行。因为原厂维护，可以得到进行及时的修复和维护。
3. Dockerfile
如果想要定制已有的镜像，可以通过编辑Dockerfile文件，重新build，重新打包成镜像。
推荐使用这种方式。
4. Commit
也可以通过一个镜像启动一个容器，然后进行操作，最终通过commit方式。
但并不推荐该种方式。
虽然说通过commit方式像是通过操作虚拟机一样的方式，但是容器毕竟是容器，它不是虚拟机
希望大家能够适应使用dockerfile的方式生成镜像的习惯。

----

## Docker client
docker 客户端
## Docker server
Docker  daemon 的主要组成部分，接收用户通过docker client发送的请求，并按照相应的
路由规则实现路由分发。
## Docker image
Docker镜像运行之后成为容器。(docker run)
- 启动速度快
- 体积小
    - 磁盘占用空间小
    - 内存消耗小
docker采用了分层技术，比如说构建一个docker镜像，需要基于一个base，基础镜像(父镜像)，不包含父镜像大小，
只是自身的大小。

## Docker Registry
Registry是Docker镜像的中央仓库。(pull/push)


## 克隆docker-training
构建 [centos7\mysql\php-fpm\wordpress] docker镜像
通过以克隆的代码制作成docker镜像，
首先要构建Dockerfile文件
Dockerfile 自动构建docker镜像的配置文件。命令类似shell
通过docker build生成docker镜像。例如有更新的时候想要立即生成镜像，可以通过自动化的平台自动发现git的变化(有git才能谈自动化平台)
写好Dockerfile和了解Dockerfile是非常关键的。(Dockerfile的编写至关重要)

```

#
# DOCKER-VERSION    1.6.2
#
# Dockerizing CentOS7: Dockerfile for building CentOS images
#

FROM centos:centos7.1.1503 #基础镜像
MAINTAINER JOHN,C.Q.Feng <feng-qichao@qq.com> #Dockerfile维护者

ENV TZ "Asia/Shanghai" #环境变量(可有多个)
# ENV TERM xterm
#ENV（environment）设置环境变量，一个Dockerfile中可以写多个。以上例子是：设置docker容器的时区为Shanghai

ADD aliyun-mirror.repo /etc/yum.repos.d/CentOS-Base.repo
ADD aliyun-epel.repo /etc/yum.repos.d/epel.repo

RUN yum install -y curl wget tar bzip2 unzip vim-enhanced passwd sudo yum-utils hostname net-tools rsync man && \
    yum install -y gcc gcc-c++ git make automake cmake patch logrotate python-devel libpng-devel libjpeg-devel && \
    yum install -y --enablerepo=epel pwgen python-pip && \
    yum clean all

RUN pip install supervisor
ADD supervisord.conf /etc/supervisord.conf

RUN mkdir -p /etc/supervisor.conf.d && \
    mkdir -p /var/log/supervisor

EXPOSE 22

ENTRYPOINT ["/usr/bin/supervisord", "-n", "-c", "/etc/supervisord.conf"]
```

`FROM centos:centos7.1.1503`  
衍生自基础镜像 基于父镜像构建其他docker镜像，父镜像：可以通过docker pull 命令获得，也可以自制

`MAINTAINER JOHN,C.Q.Feng <feng-qichao@qq.com>`  
 Dockerfile维护者

`ENV TZ "Asia/Shanghai" #环境变量(可有多个)`
`ENV TERM xterm`  
ENV（environment）设置环境变量，一个Dockerfile中可以写多个。以上例子是：设置docker容器的时区为Shanghai

Dockerfile有两条指令可以拷贝文件  
`ADD aliyun-mirror.repo /etc/yum.repos.d/CentOS-Base.repo`  
`ADD aliyun-epel.repo /etc/yum.repos.d/epel.repo`  














