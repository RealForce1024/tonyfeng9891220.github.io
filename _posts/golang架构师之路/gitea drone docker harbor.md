
 
###  1. 申请centos-7 虚机 
 
```sh
ssh -i "aws-test01.pem" centos@ec2-52-80-111-246.cn-north-1.compute.amazonaws.com.cn
```

### 建立数据库

```sh
CREATE DATABASE IF NOT EXISTS gitea_db
  DEFAULT CHARSET utf8
  COLLATE utf8_general_ci;
```
###  2. 安装golang环境 
 
```sh
yum clean all
yum update
wget https://storage.googleapis.com/golang/go1.8.5.linux-amd64.tar.gz
tar -xvf go1.8.5.linux-amd64.tar.gz
sudo mv go /usr/local
export GOROOT=/usr/local/go
export GOPATH=$HOME/Apps/app1
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
go version
go env
```


网络原因可能导致不好下载golang，可以取巧翻墙再上传
![](media/15152260114293.jpg)

###  3. 安装git
 
 
 ![](media/15152976566484.jpg)


```sh
sudo yum install -y git
```

### 4. centos7 install mysql5.7

```sh
curl -LO http://dev.mysql.com/get/mysql57-community-release-el7-11.noarch.rpm

1.配置 yum 源

去 MySQL 官网下载 YUM 的 RPM 安装包，http://dev.mysql.com/downloads/repo/yum/

下载 mysql 源安装包

$ curl -LO http://dev.mysql.com/get/mysql57-community-release-el7-11.noarch.rpm
安装 mysql 源

$ sudo yum localinstall mysql57-community-release-el7-11.noarch.rpm
检查 yum 源是否安装成功

$ sudo yum repolist enabled | grep "mysql.*-community.*"
mysql-connectors-community           MySQL Connectors Community              21
mysql-tools-community                MySQL Tools Community                   38
mysql57-community                    MySQL 5.7 Community Server             130
如上所示，找到了 mysql 的安装包

2.安装

$ sudo yum install mysql-community-server
3.启动

安装服务

$ sudo systemctl enable mysqld
启动服务

$ sudo systemctl start mysqld
查看服务状态

$ sudo systemctl status mysqld
4.修改 root 默认密码

MySQL 5.7 启动后，在 /var/log/mysqld.log 文件中给 root 生成了一个默认密码。通过下面的方式找到 root 默认密码，然后登录 mysql 进行修改：

$ grep 'temporary password' /var/log/mysqld.log
[Note] A temporary password is generated for root@localhost: **********
登录 MySQL 并修改密码

$ mysql -u root -p
Enter password: 
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY 'MyNewPass4!';
注意：MySQL 5.7 默认安装了密码安全检查插件（validate_password），默认密码检查策略要求密码必须包含：大小写字母、数字和特殊符号，并且长度不能少于 8 位。

通过 MySQL 环境变量可以查看密码策略的相关信息：

mysql> SHOW VARIABLES LIKE 'validate_password%';
+--------------------------------------+--------+
| Variable_name                        | Value  |
+--------------------------------------+--------+
| validate_password_check_user_name    | OFF    |
| validate_password_dictionary_file    |        |
| validate_password_length             | 8      |
| validate_password_mixed_case_count   | 1      |
| validate_password_number_count       | 1      |
| validate_password_policy             | MEDIUM |
| validate_password_special_char_count | 1      |
+--------------------------------------+--------+
7 rows in set (0.01 sec)
具体修改，参见 http://dev.mysql.com/doc/refman/5.7/en/validate-password-options-variables.html#sysvar_validate_password_policy

指定密码校验策略

$ sudo vi /etc/my.cnf

[mysqld]
# 添加如下键值对, 0=LOW, 1=MEDIUM, 2=STRONG
validate_password_policy=0
禁用密码策略

$ sudo vi /etc/my.cnf
	
[mysqld]
# 禁用密码校验策略
validate_password = off
重启 MySQL 服务，使配置生效

$ sudo systemctl restart mysqld
5.添加远程登录用户

MySQL 默认只允许 root 帐户在本地登录，如果要在其它机器上连接 MySQL，必须修改 root 允许远程连接，或者添加一个允许远程连接的帐户，为了安全起见，本例添加一个新的帐户：

mysql> GRANT ALL PRIVILEGES ON *.* TO 'admin'@'%' IDENTIFIED BY 'secret' WITH GRANT OPTION;
6.配置默认编码为 utf8

MySQL 默认为 latin1, 一般修改为 UTF-8

$ vi /etc/my.cnf
[mysqld]
# 在myslqd下添加如下键值对
character_set_server=utf8
init_connect='SET NAMES utf8'
重启 MySQL 服务，使配置生效

$ sudo systemctl restart mysqld
查看字符集

mysql> SHOW VARIABLES LIKE 'character%';
+--------------------------+----------------------------+
| Variable_name            | Value                      |
+--------------------------+----------------------------+
| character_set_client     | utf8                       |
| character_set_connection | utf8                       |
| character_set_database   | utf8                       |
| character_set_filesystem | binary                     |
| character_set_results    | utf8                       |
| character_set_server     | utf8                       |
| character_set_system     | utf8                       |
| character_sets_dir       | /usr/share/mysql/charsets/ |
+--------------------------+----------------------------+
8 rows in set (0.00 sec
7.开启端口

$ sudo firewall-cmd --zone=public --add-port=3306/tcp --permanent
$ sudo firewall-cmd --reload
参考资料
Using the MySQL Yum Repository
MySQL 5.7 安装与配置（YUM）
```
[centos7 install mysql5.7](http://qizhanming.com/blog/2017/05/10/centos-7-yum-install-mysql-57)

[centos7 主从复制](http://qizhanming.com/blog/2017/06/20/how-to-config-mysql-57-master-slave-replication-on-centos-7)

[CentOS7安装配置mysql5.7](http://blog.csdn.net/jssg_tzw/article/details/68944693)
[CentOS7下安装MySQL5.7安装与配置（YUM）](http://www.centoscn.com/mysql/2016/0626/7537.html)
[（笔记）CentOS 7 安装与卸载MySQL 5.7跳坑](https://www.jianshu.com/p/e54ff5283f18)
[(笔记备份)CentOS 7 安装与卸载 MySQL 5.7](https://micorochio.github.io/2017/01/16/mark-CentOS-7-install-mysql-5-7/)

![](media/15152336546119.jpg)
aopjLCkl/9uz



create db

```sh
CREATE DATABASE IF NOT EXISTS gitea_db
  DEFAULT CHARSET utf8
  COLLATE utf8_general_ci;
```

![](media/15152344342366.jpg)

 
### 5. 安装 gitea

```sh
wget -O gitea https://dl.gitea.io/gitea/1.0.1/gitea-1.0.1-linux-amd64
chmod +x gitea
```
经测试有时还是本地下载后上传速度快...

![](media/15152266083433.jpg)


设置开机启动
```
sudo vim /etc/systemd/system/gitea.service
sudo systemctl daemon-reload
sudo systemctl enable gitea
sudo systemctl restart gitea
```

```
[fengqichao@host-10-150-26-82 ~]$ sudo vim /etc/systemd/system/gitea.service        
[Unit]
Description=Gitea (Git with a cup of tea)
After=syslog.target
After=network.target
#After=mysqld.service
#After=postgresql.service
#After=memcached.service
#After=redis.service

[Service]
# Modify these two values and uncomment them if you have
# repos with lots of files and get an HTTP error 500 because
# of that
###
#LimitMEMLOCK=infinity
#LimitNOFILE=65535
RestartSec=2s
Type=simple
User=fengqichao
Group=fengqichao
WorkingDirectory=/home/fengqichao/gitea
ExecStart=/home/fengqichao/gitea web
Restart=always
Environment=USER=fengqichao HOME=/home/fengqichao/git

[Install]
WantedBy=multi-user.target
                                      
```
6. 运行


```sh
./gitea web
```
![](media/15152267354409.jpg)


![](media/15152293407472.jpg)


![](media/15152293629349.jpg)


![](media/15152293789836.jpg)








https://gist.github.com/joffilyfe/1a99250cb74bb75e29cbe8d6ca8ceedb


![](media/15152352013193.jpg)





![](media/15152353771862.jpg)


82服务器上指定了mysql也ok
![](media/15152359774366.jpg)


![](media/15152360563697.jpg)

数据库库中果然有！
![](media/15152362873771.jpg)


1. 面临数据库迁移的问题
![](media/15152369444633.jpg)


方案1. 打通网络，直接导入
方案2. 将大家的代码都全部合并到master，然后手动的上传到该地址仓库，然后建立自己的分支开发，注意修改gitconfig 中的ip地址



对比gitea 1. 简单且易维护 2.性能高 3.占用资源少(比较看重)
http://10.150.26.83:3000/repo/migrate

https://blog.wu-boy.com/drone-devops/
https://www.udemy.com/devops-oneday/?couponCode=DRONE-DEVOPS
一天學會 DevOps 自動化測試及部署
![](media/15152373607932.jpg)


![](media/15152377336638.jpg)

https://gist.github.com/joffilyfe/1a99250cb74bb75e29cbe8d6ca8ceedb

https://blog.wu-boy.com/2017/09/why-i-choose-drone-as-ci-cd-tool/

https://rootsongjc.gitbooks.io/kubernetes-handbook/content/practice/jenkins-ci-cd.html


解决yum因为多个python版本，使用了高版本而无法使用的问题
http://blog.csdn.net/ei__nino/article/details/8495295

登录之后,,,尼玛6000...
![](media/15153054009037.jpg)


## 系统开机服务

### 开机服务

```sh
[fengqichao@host-10-150-26-83 ~]$ sudo vim /etc/systemd/system/gitea.service     
[Unit]
Description=Gitea (Git with a cup of tea)
After=syslog.target
After=network.target
#After=mysqld.service
#After=postgresql.service
#After=memcached.service
#After=redis.service

[Service]
# Modify these two values and uncomment them if you have
# repos with lots of files and get an HTTP error 500 because
# of that
###
#LimitMEMLOCK=infinity
#LimitNOFILE=65535
RestartSec=2s
Type=simple
User=fengqichao
Group=fengqichao
WorkingDirectory=/home/fengqichao
ExecStart=/home/fengqichao/gitea web
#ExecStart=/bin/bash /home/fengqichao/start-gitea.sh
Restart=always
Environment=USER=fengqichao HOME=/home/fengqichao

[Install]
WantedBy=multi-user.target
```



### 特别蛋疼的bug解决

![](media/15152993861035.jpg)


![](media/15152985276892.jpg)


```sh
sudo cp  /usr/local/git/bin/git  /usr/bin/
```

![](media/15152990966318.jpg)



http://www.zslin.com/web/article/detail/9




## 安装 supervisor
最终解决了python多版本 ，yum因为python版本切换的问题，以及supervisor 手动安装的问题，最终还是重新使用pip安装上了supervisor
https://blog.fazero.me/2016/12/16/supervisor-usage/

![](media/15152445067787.jpg)

## 
mkdir -p /home/fengqichao/git/gitea/log/supervisor
sudo vim /etc/supervisor/supervisord.conf


82让人真的很蛋疼...需要运维来重置了...

83服务器上的supervisor ok的





83上重新走了一遍
![](media/15152474093870.jpg)


![](media/15152474209042.jpg)


![](media/15152474370477.jpg)

22:04 开始立即安装

30s左右之后立即跳转了 哈哈~
![](media/15152474959664.jpg)



![](media/15152497523192.jpg)

![](media/15152503113951.jpg)

sudo systemctl daemon-reload
sudo systemctl restart gitea.service 
sudo systemctl status gitea.service


![](media/15152515514264.jpg)

![](media/15152525068238.jpg)

![](media/15152527557213.jpg)



```sh
[fengqichao@host-10-150-26-83 ~]$ sudo systemctl daemon-reload          
[fengqichao@host-10-150-26-83 ~]$ sudo systemctl restart gitea.service 
[fengqichao@host-10-150-26-83 ~]$ sudo systemctl enable gitea          
[fengqichao@host-10-150-26-83 ~]$ sudo systemctl status gitea.service -l
‚óè gitea.service - Gitea (Git with a cup of tea)
   Loaded: loaded (/etc/systemd/system/gitea.service; enabled; vendor preset: disabled)
   Active: activating (auto-restart) since Sun 2018-01-07 00:05:43 CST; 763ms ago
  Process: 5828 ExecStart=/bin/bash /home/fengqichao/start-gitea.sh (code=exited, status=0/SUCCESS)
 Main PID: 5828 (code=exited, status=0/SUCCESS)
```

![](media/15152929849967.jpg)
使用啥的都是使用了建立的mysql，一切都ok。


![](media/15152963028754.jpg)

![](media/15152963147288.jpg)






---------------
## ubuntu上安装mysql
[Install MySQL on Ubuntu 14.04](https://linode.com/docs/databases/mysql/install-mysql-on-ubuntu-14-04/)
[在Ubuntu中
安装MySQL](http://blog.fens.me/linux-mysql-install/)


```sh
sudo apt-get install mysql-server
```

git golang 和centos完全一样

秒执行



systemd的使用t
https://linux.cn/article-3719-1.html


## gitlab迁移到 gitea的方案

方案1. 打通网络
方案2. 将大家的代码都全部合并到master，然后手动的上传到该地址仓库，然后建立自己的分支开发，注意修改gitconfig 中的ip地址



对比gitea 1. 简单且易维护 2.性能高 3.占用资源少(比较看重)
http://10.150.26.82:3000/repo/migrate


## readme
官网
gitea: https://docs.gitea.io
drone: https://drone.io
关键词: drone ci cd / drone持续集成

**持续集成鼓励开发团队尽早测试并将其更改集中到共享代码库，以最大程度地减少集成冲突。 通过消除部署或发布方式上的障碍，持续交付从此基础上构建。 通过部署自动通过测试套件的每个构建，进一步扩展连续部署。
虽然上述术语主要涉及策略和实践，但软件工具在允许组织实现这些目标方面发挥重要作用。 CI / CD软件可以帮助团队通过一系列的阶段自动推进新的变化，以减少反馈时间和消除过程中的摩擦。**
[CI / CD工具比较：Jenkins，GitLab CI，Buildbot，Drone和大厅](https://www.howtoing.com/ci-cd-tools-comparison-jenkins-gitlab-ci-buildbot-drone-and-concourse)

[安裝 Drone 0.5 自動測試平台並與 Github 連結](https://yami.io/drone/)


```
若要達到上述這種效果就需事先寫好「單元測試」，而 Drone 扮演的角色就是自動執行單元測試，而且這個執行的環境是獨立的，不會真正干擾到你的主機。

Drone 亦能在測試成功時自動幫你部署到正式的伺服器，或是執行任何額外的腳本。倘若你聽過 Travis CI，那麼也許你就會對這些事情感到不陌生。
```


[Drone](https://yeasy.gitbooks.io/docker_practice/content/cases/ci/drone.html)

[体验基于gogs+Drone搭建的CI/CD平台](https://www.jianshu.com/p/15506f46f75a)

[Docker + Drone CI/CD 实践](https://segmentfault.com/a/1190000012066735)

[使用Rancher和DroneCI建立超高速Docker CI/CD流水线](https://segmentfault.com/a/1190000010761792)

[單元測試簡介以及在 Docker 上部署 Drone 並連結至 GitHub](https://blog.birkhoff.me/unit-test-intro-and-install-drone/)

[k8s与CICD--将drone部署到kubernetes中，实现agent动态收缩](https://studygolang.com/articles/11968?fr=sidebar)
[flow.ci + Github + Slack 一步步搭建 Python 自动化持续集成](https://segmentfault.com/a/1190000005909578)
https://github.com/rootsongjc/kubernetes-handbook/blob/master/practice/drone-ci-cd.md

https://www.jianshu.com/p/15506f46f75a


[基于docker的高可用方案](https://www.jianshu.com/p/62e5ac2ede2b?utm_campaign=maleskine&utm_content=note&utm_medium=seo_notes&utm_source=recommendation)

https://www.jianshu.com/p/ac8020d9c473

[基于rancher的ci cd](https://segmentfault.com/a/1190000007808836)

![](media/15153063705145.jpg)


**持续集成，交付和部署软件是旨在使您的流程可靠和可重复的复杂自动化系统。 从上面的描述可以看出，关于自动化测试和发布如何最大程度地实现，有很多不同的想法，重点放在方程的不同部分。 没有一个工具可以满足每个项目的需求，但是通过这么多高品质的开源解决方案，您很有可能找到一个符合团队需求的系统。**


[如何在Ubuntu 16.04上设置Drone的连续集成管道](https://www.howtoing.com/how-to-set-up-continuous-integration-pipelines-with-drone-on-ubuntu-16-04)


## 实践

install docker  docker-compose 

```sh
wget https://bootstrap.pypa.io/get-pip.py
sudo python get-pip.py
sudo pip install -U docker-compose
docker-compose version
```

docker-compose补全命令
```sh
curl -L https://raw.githubusercontent.com/docker/compose/1.8.0/contrib/completion/bash/docker-compose > /etc/bash_completion.d/docker-compose
```

![](media/15153158678653.jpg)




```sh
version: '2'

services:
  drone-server:
    image: drone/drone:0.8

    ports:
      - 80:8000
      - 9000
    volumes:
      - /var/lib/drone:/var/lib/drone/
    restart: always
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=54.223.251.31
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=29960276873586de9d08
      - DRONE_GITHUB_SECRET=75522ea3dfaf87f8f0b653d5474d72c38461b047
      - DRONE_SECRET=Onto-Tear-Level-English-9

  drone-agent:
    image: drone/agent:0.8

    command: agent
    restart: always
    depends_on:
      - drone-server
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - DRONE_SERVER=drone-server:9000
      - DRONE_SECRET=Onto-Tear-Level-English-9
```

54.223.251.31


![](media/15153183880410.jpg)

![](media/15153184059809.jpg)

还是要做好加速器
![](media/15153187029560.jpg)



```sh
[centos@ip-172-31-4-17 drone]$ ll
total 4
-rw-rw-r--. 1 centos centos 731 Jan  7 09:50 docker-compose.yml
[centos@ip-172-31-4-17 drone]$ curl -sSL https://get.daocloud.io/daotools/set_mirror.sh | sh -s http://bbfa5e62.m.daocloud.io
docker version >= 1.12
{"registry-mirrors": ["http://bbfa5e62.m.daocloud.io"]}
Success.
You need to restart docker to take effect: sudo systemctl restart docker
[centos@ip-172-31-4-17 drone]$ sudo systemctl restart docker
[centos@ip-172-31-4-17 drone]$ sudo docker-compose up
Pulling drone-server (drone/drone:0.8)...
0.8: Pulling from drone/drone
297640bfa2ef: Downloading [>                                                  ]  1.787kB/155.1297640bfa2ef: Downloading [==================================================>]  155.1kB/155.1297640bfa2ef: Extracting [==========>                                        ]  32.77kB/155.1k297640bfa2ef: Extracting [==================================================>]  155.1kB/155.1k297640bfa2ef: Extracting [==================================================>]  155.1kB/155.1k297640bfa2ef: Pull complete
e3ff4294cfa6: Downloading [>                                                  ]  106.1kB/9.15Me3ff4294cfa6: Downloading [===>                                               ]  707.4kB/9.15Me3ff4294cfa6: Downloading [=======>                                           ]  1.316MB/9.15Me3ff4294cfa6: Downloading [==========>                                        ]  1.919MB/9.15Me3ff4294cfa6: Downloading [=============>                                     ]   2.41MB/9.15Me3ff4294cfa6: Downloading [=============>                                     ]  2.503MB/9.15Me3ff4294cfa6: Downloading [====================>                              ]  3.831MB/9.15Me3ff4294cfa6: Downloading [========================>                          ]  4.532MB/9.15Me3ff4294cfa6: Downloading [============================>                      ]  5.163MB/9.15Me3ff4294cfa6: Downloading [===============================>                   ]  5.734MB/9.15Me3ff4294cfa6: Downloading [==================================>                ]  6.364MB/9.15Me3ff4294cfa6: Downloading [=====================================>             ]  6.895MB/9.15Me3ff4294cfa6: Downloading [========================================>          ]  7.406MB/9.15Me3ff4294cfa6: Downloading [===========================================>       ]  7.894MB/9.15Me3ff4294cfa6: Downloading [=============================================>     ]  8.298MB/9.15Me3ff4294cfa6: Downloading [===============================================>   ]  8.708MB/9.15Me3ff4294cfa6: Pull complete
Digest: sha256:ff0da5b034a04c0b09e1936ce5b3c8081d82f08124d6440d21eb813c2a92c234
Status: Downloaded newer image for drone/drone:0.8
Pulling drone-agent (drone/agent:0.8)...
0.8: Pulling from drone/agent
297640bfa2ef: Already exists
aef610f97352: Pull complete
Digest: sha256:1f0d0406bf86c850c3bed3c7bda445d412424598397b256ac4a7b1e017da5dc7
Creating drone_drone-server_1 ... done
Creating drone_drone-server_1 ...
Creating drone_drone-agent_1  ... done
Attaching to drone_drone-server_1, drone_drone-agent_1
drone-server_1  | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
drone-server_1  |  - using env:	export GIN_MODE=release
drone-server_1  |  - using code:	gin.SetMode(gin.ReleaseMode)
drone-server_1  |
drone-server_1  | [GIN-debug] GET    /logout                   --> github.com/drone/drone/server.GetLogout (12 handlers)
drone-server_1  | [GIN-debug] GET    /login                    --> github.com/drone/drone/server.HandleLogin (12 handlers)
drone-server_1  | [GIN-debug] GET    /api/user                 --> github.com/drone/drone/server.GetSelf (13 handlers)
drone-server_1  | [GIN-debug] GET    /api/user/feed            --> github.com/drone/drone/server.GetFeed (13 handlers)
drone-server_1  | [GIN-debug] GET    /api/user/repos           --> github.com/drone/drone/server.GetRepos (13 handlers)
drone-server_1  | [GIN-debug] POST   /api/user/token           --> github.com/drone/drone/server.PostToken (13 handlers)
drone-server_1  | [GIN-debug] DELETE /api/user/token           --> github.com/drone/drone/server.DeleteToken (13 handlers)
drone-server_1  | [GIN-debug] GET    /api/users                --> github.com/drone/drone/server.GetUsers (13 handlers)
drone-server_1  | [GIN-debug] POST   /api/users                --> github.com/drone/drone/server.PostUser (13 handlers)
drone-server_1  | [GIN-debug] GET    /api/users/:login         --> github.com/drone/drone/server.GetUser (13 handlers)
drone-server_1  | [GIN-debug] PATCH  /api/users/:login         --> github.com/drone/drone/server.PatchUser (13 handlers)
drone-server_1  | [GIN-debug] DELETE /api/users/:login         --> github.com/drone/drone/server.DeleteUser (13 handlers)
drone-server_1  | [GIN-debug] POST   /api/repos/:owner/:name   --> github.com/drone/drone/server.PostRepo (16 handlers)
drone-server_1  | [GIN-debug] GET    /api/repos/:owner/:name   --> github.com/drone/drone/server.GetRepo (15 handlers)
drone-server_1  | [GIN-debug] GET    /api/repos/:owner/:name/builds --> github.com/drone/drone/server.GetBuilds (15 handlers)
drone-server_1  | [GIN-debug] GET    /api/repos/:owner/:name/builds/:number --> github.com/drone/drone/server.GetBuild (15 handlers)
drone-server_1  | [GIN-debug] GET    /api/repos/:owner/:name/logs/:number/:pid --> github.com/drone/drone/server.GetProcLogs (15 handlers)
drone-server_1  | [GIN-debug] GET    /api/repos/:owner/:name/logs/:number/:pid/:proc --> github.com/drone/drone/server.GetBuildLogs (15 handlers)
drone-server_1  | [GIN-debug] GET    /api/repos/:owner/:name/files/:number --> github.com/drone/drone/server.FileList (15 handlers)
drone-server_1  | [GIN-debug] GET    /api/repos/:owner/:name/files/:number/:proc/*file --> github.com/drone/drone/server.FileGet (15 handlers)
drone-server_1  | [GIN-debug] GET    /api/repos/:owner/:name/secrets --> github.com/drone/drone/server.GetSecretList (16 handlers)
drone-agent_1   | {"time":"2018-01-07T09:51:36Z","level":"debug","message":"request next execution"}
drone-server_1  | [GIN-debug] POST   /api/repos/:owner/:name/secrets --> github.com/drone/drone/server.PostSecret (16 handlers)
drone-server_1  | [GIN-debug] GET    /api/repos/:owner/:name/secrets/:secret --> github.com/drone/drone/server.GetSecret (16 handlers)
drone-server_1  | [GIN-debug] PATCH  /api/repos/:owner/:name/secrets/:secret --> github.com/drone/drone/server.PatchSecret (16 handlers)
drone-server_1  | [GIN-debug] DELETE /api/repos/:owner/:name/secrets/:secret --> github.com/drone/drone/server.DeleteSecret (16 handlers)
drone-server_1  | [GIN-debug] GET    /api/repos/:owner/:name/registry --> github.com/drone/drone/server.GetRegistryList (16 handlers)
drone-server_1  | [GIN-debug] POST   /api/repos/:owner/:name/registry --> github.com/drone/drone/server.PostRegistry (16 handlers)
drone-server_1  | [GIN-debug] GET    /api/repos/:owner/:name/registry/:registry --> github.com/drone/drone/server.GetRegistry (16 handlers)
drone-server_1  | [GIN-debug] PATCH  /api/repos/:owner/:name/registry/:registry --> github.com/drone/drone/server.PatchRegistry (16 handlers)
drone-server_1  | [GIN-debug] DELETE /api/repos/:owner/:name/registry/:registry --> github.com/drone/drone/server.DeleteRegistry (16 handlers)
drone-server_1  | [GIN-debug] PATCH  /api/repos/:owner/:name   --> github.com/drone/drone/server.PatchRepo (16 handlers)
drone-server_1  | [GIN-debug] DELETE /api/repos/:owner/:name   --> github.com/drone/drone/server.DeleteRepo (16 handlers)
drone-server_1  | [GIN-debug] POST   /api/repos/:owner/:name/chown --> github.com/drone/drone/server.ChownRepo (16 handlers)
drone-server_1  | [GIN-debug] POST   /api/repos/:owner/:name/repair --> github.com/drone/drone/server.RepairRepo (16 handlers)
drone-server_1  | [GIN-debug] POST   /api/repos/:owner/:name/move --> github.com/drone/drone/server.MoveRepo (16 handlers)
drone-server_1  | [GIN-debug] POST   /api/repos/:owner/:name/builds/:number --> github.com/drone/drone/server.PostBuild (16 handlers)
drone-server_1  | [GIN-debug] DELETE /api/repos/:owner/:name/builds/:number --> github.com/drone/drone/server.ZombieKill (16 handlers)
drone-server_1  | [GIN-debug] POST   /api/repos/:owner/:name/builds/:number/approve --> github.com/drone/drone/server.PostApproval (16 handlers)
drone-server_1  | [GIN-debug] POST   /api/repos/:owner/:name/builds/:number/decline --> github.com/drone/drone/server.PostDecline (16 handlers)
drone-server_1  | [GIN-debug] DELETE /api/repos/:owner/:name/builds/:number/:job --> github.com/drone/drone/server.DeleteBuild (16 handlers)
drone-server_1  | [GIN-debug] GET    /api/badges/:owner/:name/status.svg --> github.com/drone/drone/server.GetBadge (12 handlers)
drone-server_1  | [GIN-debug] GET    /api/badges/:owner/:name/cc.xml --> github.com/drone/drone/server.GetCC (12 handlers)
drone-server_1  | [GIN-debug] POST   /hook                     --> github.com/drone/drone/server.PostHook (12 handlers)
drone-server_1  | [GIN-debug] POST   /api/hook                 --> github.com/drone/drone/server.PostHook (12 handlers)
drone-server_1  | [GIN-debug] GET    /stream/events            --> github.com/drone/drone/server.EventStreamSSE (12 handlers)
drone-server_1  | [GIN-debug] GET    /stream/logs/:owner/:name/:build/:number --> github.com/drone/drone/server.LogStreamSSE (15 handlers)
drone-server_1  | [GIN-debug] GET    /api/info/queue           --> github.com/drone/drone/server.GetQueueInfo (13 handlers)
drone-server_1  | [GIN-debug] GET    /authorize                --> github.com/drone/drone/server.HandleAuth (12 handlers)
drone-server_1  | [GIN-debug] POST   /authorize                --> github.com/drone/drone/server.HandleAuth (12 handlers)
drone-server_1  | [GIN-debug] POST   /authorize/token          --> github.com/drone/drone/server.GetLoginToken (12 handlers)
drone-server_1  | [GIN-debug] GET    /api/builds               --> github.com/drone/drone/server.GetBuildQueue (13 handlers)
drone-server_1  | [GIN-debug] GET    /api/debug/pprof/         --> github.com/drone/drone/server/debug.IndexHandler.func1 (13 handlers)
drone-server_1  | [GIN-debug] GET    /api/debug/pprof/heap     --> github.com/drone/drone/server/debug.HeapHandler.func1 (13 handlers)
drone-server_1  | [GIN-debug] GET    /api/debug/pprof/goroutine --> github.com/drone/drone/server/debug.GoroutineHandler.func1 (13 handlers)
drone-server_1  | [GIN-debug] GET    /api/debug/pprof/block    --> github.com/drone/drone/server/debug.BlockHandler.func1 (13 handlers)
drone-server_1  | [GIN-debug] GET    /api/debug/pprof/threadcreate --> github.com/drone/drone/server/debug.ThreadCreateHandler.func1 (13 handlers)
drone-server_1  | [GIN-debug] GET    /api/debug/pprof/cmdline  --> github.com/drone/drone/server/debug.CmdlineHandler.func1 (13 handlers)
drone-server_1  | [GIN-debug] GET    /api/debug/pprof/profile  --> github.com/drone/drone/server/debug.ProfileHandler.func1 (13 handlers)
drone-server_1  | [GIN-debug] GET    /api/debug/pprof/symbol   --> github.com/drone/drone/server/debug.SymbolHandler.func1 (13 handlers)
drone-server_1  | [GIN-debug] POST   /api/debug/pprof/symbol   --> github.com/drone/drone/server/debug.SymbolHandler.func1 (13 handlers)
drone-server_1  | [GIN-debug] GET    /api/debug/pprof/trace    --> github.com/drone/drone/server/debug.TraceHandler.func1 (13 handlers)
drone-server_1  | [GIN-debug] GET    /metrics                  --> github.com/drone/drone/server/metrics.PromHandler.func1 (13 handlers)
drone-server_1  | [GIN-debug] GET    /version                  --> github.com/drone/drone/server.Version (12 handlers)
drone-server_1  | [GIN-debug] GET    /healthz                  --> github.com/drone/drone/server.Health (12 handlers)
Connection to ec2-54-223-251-31.cn-north-1.compute.amazonaws.com.cn closed by remote host.
Connection to ec2-54-223-251-31.cn-north-1.compute.amazonaws.com.cn closed.
☁  aws  ssh -i "aws-test01.pem" centos@ec2-54-223-251-31.cn-north-1.compute.amazonaws.com.cn

Last login: Sun Jan  7 09:49:20 2018 from 23.99.107.46
-bash: warning: setlocale: LC_CTYPE: cannot change locale (UTF-8): No such file or directory
[centos@ip-172-31-4-17 ~]$
```


![](media/15153189231681.jpg)

![](media/15153190411794.jpg)

访问: http://54.223.251.31/ 

![](media/15153192619894.jpg)



```sh
https://github.com/login?client_id=29960276873586de9d08&return_to=%2Flogin%2Foauth%2Fauthorize%3Fclient_id%3D29960276873586de9d08%26redirect_uri%3Dhttp%253A%252F%252F54.223.251.31%252Fauthorize%26response_type%3Dcode%26scope%3Drepo%2Brepo%253Astatus%2Buser%253Aemail%2Bread%253Aorg%26state%3Ddrone
```
![](media/15153193482261.jpg)

![](media/15153195522267.jpg)

然后看到
![](media/15153195787011.jpg)

可能是这里的问题
![](media/15153198790770.jpg)

哈哈~果然搞定!
![](media/15153199435038.jpg)

![](media/15153200088127.jpg)


![](media/15153201137402.jpg)


![](media/15153201419098.jpg)


![](media/15153202323279.jpg)

![](media/15153202517085.jpg)

![](media/15153202937091.jpg)


![](media/15153203161355.jpg)




真的很赞
![](media/15153217990574.jpg)


![](media/15153217722531.jpg)

![](media/15153218582766.jpg)

一下就跑了三个环境
![](media/15153218745688.jpg)

![](media/15153219070867.jpg)


![](media/15153219829424.jpg)

![](media/15153219918243.jpg)


![](media/15153219994587.jpg)



新增java项目，注意手动同步以下
![](media/15153222901737.jpg)

![](media/15153223137406.jpg)

![](media/15153223348764.jpg)


![](media/15153225333326.jpg)


![](media/15153225535983.jpg)



java1.8

```sh
sudo mv jdk1.8.0_152 /usr/local
export JAVA_HOME=/usr/local/jdk1.8.0_152
export PATH=$GOPATH/bin:$GOROOT/bin:$JAVA_HOME/bin:$PATH
```

在服务器上安装gradle速度还是可以的，但是jdk还是手动上传吧



```sh
sudo mv apache-maven-3.5.2 /usr/local
export MAVEN_HOME=/usr/local/apache-maven-3.5.2
export PATH=$JAVA_HOME/bin:$MAVEN_HOME/bin:$PATH
```

![](media/15153799424014.jpg)




# gradle to maven
https://notes.wanghao.work/2017-08-21-Gradle%E3%80%81Maven%E9%A1%B9%E7%9B%AE%E7%9B%B8%E4%BA%92%E8%BD%AC%E6%8D%A2.html




http://blog.csdn.net/lsgqjh/article/details/72597786

http://download.csdn.net/download/happyzwh/9950314

http://blog.csdn.net/zstack_org/article/details/53665575


镜像搞定，一切就绪了
![](media/15153813716056.jpg)


![](media/15153821440867.jpg)


## 成功
![](media/15153907383424.jpg)

虽然喜悦，但是在跑另外一个的
![](media/15153907639921.jpg)

不过这次仅仅用了一半的时间
![](media/15153909352553.jpg)

这样是不错的
![](media/15153909664413.jpg)


 再最后验证下，是否重用本地依赖缓存
 
 
 ![](media/15153910692850.jpg)

 
 
 
 解决 maven cache的问题
 
 [drone volume cache plugin](https://github.com/Drillster/drone-volume-cache)
 
 [Cache directories between builds #143](https://github.com/drone/drone/issues/143)
 [Cache Docker images #43 Closed](https://github.com/drone/drone/issues/43)
 
 [Add install section to drone.yml #691](https://github.com/drone/drone/issues/691)
 
 ![](media/15153925648248.jpg)

 
 
 [](https://github.com/Drillster/drone-volume-cache/blob/master/DOCS.md)
 [](https://github.com/Drillster/drone-volume-cache/blob/master/DOCS.md)
 
 ![](media/15153938843982.jpg)

 
 
 
 https://gist.github.com/danielepolencic/2b43329495d018dc6bfe790a79b559d4
 [ ERROR: Insufficient privileges to use volumes](https://discourse.drone.io/t/solved-error-insufficient-privileges-to-use-volumes/260)
 http://readme.drone.io/admin/user-admins/
 ![](media/15153949283131.jpg)

 
 
 
 ![](media/15153952993368.jpg)

 
 ![](media/15153958127610.jpg)

 J解决缓存问题
[Implement proper Caching](https://github.com/drone/drone/issues/147)


```sh
docker run --rm \
  -e PLUGIN_REBUILD=true \
  -e PLUGIN_MOUNT=".m2" \
  -e DRONE_REPO_OWNER="foo" \
  -e DRONE_REPO_NAME="bar" \
  -e DRONE_JOB_NUMBER=0 \
  -v $(pwd):$(pwd) \
  -v /tmp/cache:/cache \
  -w $(pwd) \
  drillster/drone-volume-cache
```


酷炫的页面..
https://notes.wanghao.work/2017-08-21-Gradle%E3%80%81Maven%E9%A1%B9%E7%9B%AE%E7%9B%B8%E4%BA%92%E8%BD%AC%E6%8D%A2.html


线上坑
$path=$mavenjkjk:path  这个错误的命令执行若让其生效后会很蛋疼...各种命令找不到



