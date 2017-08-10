# docker基础操作
命令用法很简单，记住一条即可。
`docker command --help`

## 1. 拉取镜像 docker pull
`docker pull <Image>`
如果没有仓库，则说明是从默认的dockerHub仓库下载，但国内网络环境你懂的。最佳的方式是使用云服务商的dockerHub加速器。阿里云和腾讯云的加速器都非常不错。

## 2. 查看镜像

`docker images`命令列出本地已经下载的镜像

```sh
$ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
ubuntu              latest              14f60031763d        2 weeks ago         120MB
ubuntu              14.04               54333f1de4ed        2 weeks ago         188MB
```

具体的用法及参数

```sh
$ docker images --help

Usage:	docker images [OPTIONS] [REPOSITORY[:TAG]]

List images

Options:
  -a, --all             Show all images (default hides intermediate images)
      --digests         Show digests
  -f, --filter filter   Filter output based on conditions provided
      --format string   Pretty-print images using a Go template
      --help            Print usage
      --no-trunc        Don't truncate output
  -q, --quiet           Only show numeric IDs
```

## 3. 命令行式启动容器 (Ad hoc方式执行容器命令)
所谓ad hoc方式就是 一次性执行完成后即销毁。

格式: `docker run image [command] [arg]`   

```sh
$ docker run ubuntu echo 'hello docker'
hello docker
```

run在新容器中执行命令，如果镜像不存在则到dockerHub中下载。注意指定镜像的方式最好加指定的tag，否则会默认的添加image:latest，没有匹配的，则会远程拉取下载。另一种方式可以通过指定唯一的ImageId即可。

```sh
$ docker run ubuntu echo "hello world"
hello world
$ docker run ubuntu:14.04 echo "hello world"
hello world
```
注意第一个其实默认为`ubuntu:latest`的方式


当命令执行完毕后，该容器就会终止，以上命令只是执行一个命令(ad hoc)，我们可以通过

`docker ps -a`查看所有的容器(运行中的和执行过的)

```sh
$ docker ps -a
CONTAINER ID        IMAGE               COMMAND                 CREATED             STATUS                      PORTS               NAMES
07df14a317da        ubuntu              "echo 'hello docker'"   2 minutes ago       Exited (0) 2 minutes ago
```
我们发现通过`docker run image command arg`命令是一次性启动容器执行命令执行完毕后销毁容器。


## 4. 交互式启动容器(始终运行直到退出)
格式: `docker run -i -t image /bin/bash`
* `-i` --interactive=true|false false是默认  代表:交互式
* `-t` --tty=true|false false是默认   代表:终端
* `-i -t` 可以缩略为 `-it`效果等同
* /bin/bash 指定运行的shell 可默认省略

```sh
docker@fengqichao:~$  docker run -i -t ubuntu /bin/bash
root@eff926ba2186:/# echo 'hello world'
hello world
root@eff926ba2186:/# ps -ef
UID        PID  PPID  C STIME TTY          TIME CMD
root         1     0  0 03:19 ?        00:00:00 /bin/bash
root        13     1  0 03:21 ?        00:00:00 ps -ef
root@eff926ba2186:/# exit
exit
```

## 5. 查看容器

镜像可以理解为类/模板(静态)，而容器则为对象/实例(动态)。
### docker ps

`docker ps [-a]|[-l]`  

* 默认不加任何参数将返回正在运行的容器
* -a 查看所有的容器(已销毁和正运行的)
* -l 最新创建的容器 

注意:`docker ps`命令返回的字段`Containter ID`和`Names`字段均为docker为容器自动分配的。

```sh
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES

$  docker ps -l
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                     PORTS               NAMES
eff926ba2186        ubuntu              "/bin/bash"         5 minutes ago       Exited (0) 3 minutes ago                       laughing_kilby

$ docker ps -a
CONTAINER ID        IMAGE               COMMAND                 CREATED             STATUS                         PORTS               NAMES
eff926ba2186        ubuntu              "/bin/bash"             5 minutes ago       Exited (0) 3 minutes ago                           laughing_kilby
07df14a317da        ubuntu              "echo 'hello docker'"   19 minutes ago      Exited (0) 19 minutes ago                          unruffled_meninsky
```

更详细通过命令`docker ps --help`查看

```sh
$ docker ps --help

Usage:	docker ps [OPTIONS]

List containers

Options:
  -a, --all             Show all containers (default shows just running)
  -f, --filter filter   Filter output based on conditions provided
      --format string   Pretty-print containers using a Go template
      --help            Print usage
  -n, --last int        Show n last created containers (includes all states) (default -1)
  -l, --latest          Show the latest created container (includes all states)
      --no-trunc        Don't truncate output
  -q, --quiet           Only display numeric IDs
  -s, --size            Display total file sizes
```

### docker inspect

`docker inspect [containter id] | [name]`    
该命令会自省容器的配置信息  
注意container id可以唯一的标识就行比如前4位唯一标识，只需使用4位即可。
但是name需要全名，是全匹配。

```sh
$ docker ps -l
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                       PORTS               NAMES
bab10e1eb6fc        ubuntu              "/bin/bash"         10 minutes ago      Exited (127) 7 minutes ago                       goofy_hamilton

$ docker inspect bab10e1eb6fc
[
    {
        "Id": "bab10e1eb6fc6a3841c0dbc0bc65ac3f9bfe3074cc145261b38eed0b24f96445",
        "Created": "2017-04-13T10:16:46.373541887Z",
        "Path": "/bin/bash",
        "Args": [],
        "State": {
            "Status": "exited",
            "Running": false,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 0,
            "ExitCode": 127,
            "Error": "",
            "StartedAt": "2017-04-13T10:16:46.575904921Z",
            "FinishedAt": "2017-04-13T10:19:56.042879432Z"
        },
        "Image": "sha256:6a2f32de169d14e6f8a84538eaa28f2629872d7d4f580a303b296c60db36fbd7",
        "ResolvConfPath": "/var/lib/docker/containers/bab10e1eb6fc6a3841c0dbc0bc65ac3f9bfe3074cc145261b38eed0b24f96445/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/bab10e1eb6fc6a3841c0dbc0bc65ac3f9bfe3074cc145261b38eed0b24f96445/hostname",
        "HostsPath": "/var/lib/docker/containers/bab10e1eb6fc6a3841c0dbc0bc65ac3f9bfe3074cc145261b38eed0b24f96445/hosts",
        "LogPath": "/var/lib/docker/containers/bab10e1eb6fc6a3841c0dbc0bc65ac3f9bfe3074cc145261b38eed0b24f96445/bab10e1eb6fc6a3841c0dbc0bc65ac3f9bfe3074cc145261b38eed0b24f96445-json.log",
        "Name": "/goofy_hamilton",
        "RestartCount": 0,
        "Driver": "aufs",
        "MountLabel": "",
        "ProcessLabel": "",
        "AppArmorProfile": "docker-default",
        "ExecIDs": null,
        "HostConfig": {
            "Binds": null,
            "ContainerIDFile": "",
            "LogConfig": {
                "Type": "json-file",
                "Config": {}
            },
            "NetworkMode": "default",
            "PortBindings": {},
            "RestartPolicy": {
                "Name": "no",
                "MaximumRetryCount": 0
            },
            "AutoRemove": false,
            "VolumeDriver": "",
            "VolumesFrom": null,
            "CapAdd": null,
            "CapDrop": null,
            "Dns": [],
            "DnsOptions": [],
            "DnsSearch": [],
            "ExtraHosts": null,
            "GroupAdd": null,
            "IpcMode": "",
            "Cgroup": "",
            "Links": null,
            "OomScoreAdj": 0,
            "PidMode": "",
            "Privileged": false,
            "PublishAllPorts": false,
            "ReadonlyRootfs": false,
            "SecurityOpt": null,
            "UTSMode": "",
            "UsernsMode": "",
            "ShmSize": 67108864,
            "Runtime": "runc",
            "ConsoleSize": [
                0,
                0
            ],
            "Isolation": "",
            "CpuShares": 0,
            "Memory": 0,
            "NanoCpus": 0,
            "CgroupParent": "",
            "BlkioWeight": 0,
            "BlkioWeightDevice": null,
            "BlkioDeviceReadBps": null,
            "BlkioDeviceWriteBps": null,
            "BlkioDeviceReadIOps": null,
            "BlkioDeviceWriteIOps": null,
            "CpuPeriod": 0,
            "CpuQuota": 0,
            "CpuRealtimePeriod": 0,
            "CpuRealtimeRuntime": 0,
            "CpusetCpus": "",
            "CpusetMems": "",
            "Devices": [],
            "DeviceCgroupRules": null,
            "DiskQuota": 0,
            "KernelMemory": 0,
            "MemoryReservation": 0,
            "MemorySwap": 0,
            "MemorySwappiness": -1,
            "OomKillDisable": false,
            "PidsLimit": 0,
            "Ulimits": null,
            "CpuCount": 0,
            "CpuPercent": 0,
            "IOMaximumIOps": 0,
            "IOMaximumBandwidth": 0
        },
        "GraphDriver": {
            "Data": null,
            "Name": "aufs"
        },
        "Mounts": [],
        "Config": {
            "Hostname": "bab10e1eb6fc",
            "Domainname": "",
            "User": "",
            "AttachStdin": true,
            "AttachStdout": true,
            "AttachStderr": true,
            "Tty": true,
            "OpenStdin": true,
            "StdinOnce": true,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "Cmd": [
                "/bin/bash"
            ],
            "Image": "ubuntu",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": {}
        },
        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "f4f485387af4e09ce28a768e6017d75d5f8ad907f1748df6b28639492ba69dfe",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {},
            "SandboxKey": "/var/run/docker/netns/f4f485387af4",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "",
            "Gateway": "",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "",
            "IPPrefixLen": 0,
            "IPv6Gateway": "",
            "MacAddress": "",
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "543bb1bb052c6afd07e436954bfc66188000e72ebde4fba5ad749ca0409c927f",
                    "EndpointID": "",
                    "Gateway": "",
                    "IPAddress": "",
                    "IPPrefixLen": 0,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": ""
                }
            }
        }
    }
]

```

###docker top
可以使用docker+通常的linux命令方式

```sh
$ docker top my-nginx
UID                 PID                 PPID                C                   STIME               TTY                 TIME                CMD
root                2106                2090                0                   Aug09               ?                   00:00:00            nginx: master process nginx -g daemon off;
syslog              2136                2106                0                   Aug09               ?                   00:00:00            nginx: worker process
```
## 6. 删除容器
### 启动时指定删除参数
通过`docker ps -a`我们看到容器终止了但并未从磁盘中删除，如果只是临时启动查看调试，最好是在容器使用完就立即删除。可以在容器启动时指定删除参数 --rm,
`docker run -it --rm ubuntu`



```sh
ubuntu@VM-40-206-ubuntu:~$ docker ps -a
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES

ubuntu@VM-40-206-ubuntu:~$ docker run -it --rm ubuntu

root@ad012ed44bad:/# cat /etc/os-release
NAME="Ubuntu"
VERSION="16.04.2 LTS (Xenial Xerus)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 16.04.2 LTS"
VERSION_ID="16.04"
HOME_URL="http://www.ubuntu.com/"
SUPPORT_URL="http://help.ubuntu.com/"
BUG_REPORT_URL="http://bugs.launchpad.net/ubuntu/"
VERSION_CODENAME=xenial
UBUNTU_CODENAME=xenial

root@ad012ed44bad:/# exit
exit

ubuntu@VM-40-206-ubuntu:~$ docker ps -a
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
```
### 手动删除
`docker rm [containterId][containerName]`
可以一次指定多个id或name进行批量删除。

### 指定范围删除
可以使用-q列出id，-f(filter)指定范围，-a（all）

```sh
ubuntu@VM-40-206-ubuntu:~$ docker rm $(docker ps -a -q)
55786ea74515
3919977d3196
2e560729b00e
ecf071aa33e9
cda92d5f5a3d
ubuntu@VM-40-206-ubuntu:~$ docker ps -a
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
```

## 删除镜像
`docker rmi [options] Image [Image...]`


```sh
ubuntu@VM-40-206-ubuntu:~$ docker rmi --help

Usage:	docker rmi [OPTIONS] IMAGE [IMAGE...]

Remove one or more images

Options:
  -f, --force      Force removal of the image
      --help       Print usage
      --no-prune   Do not delete untagged parents
```

## 7. 指定名称、端口、后台运行容器

```sh
ubuntu@VM-40-206-ubuntu:~$ docker run --name webserver -d -p 80:80 nginx
Unable to find image 'nginx:latest' locally
latest: Pulling from library/nginx
94ed0c431eb5: Pull complete
9406c100a1c3: Pull complete
aa74daafd50c: Pull complete
Digest: sha256:788fa27763db6d69ad3444e8ba72f947df9e7e163bad7c1f5614f8fd27a311c3
Status: Downloaded newer image for nginx:latest
c8d74c1b40fcd95234568b9a98deb77a3a9fe2c29fac02094b1724cb6550e263

ubuntu@VM-40-206-ubuntu:~$ docker ps -a
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                NAMES
c8d74c1b40fc        nginx               "nginx -g 'daemon ..."   53 seconds ago      Up 52 seconds       0.0.0.0:80->80/tcp   webserver
```

## 8. 停止后台运行的容器
`docker stop <[id]|[name]>`

## 9. 进入容器执行操作

```sh
ubuntu@VM-40-206-ubuntu:~$ docker exec -it webserver bash

root@c8d74c1b40fc:/# echo '<h1>Hello, Docker!</h1>' > /usr/share/nginx/html/index.html
root@c8d74c1b40fc:/# exit
exit

ubuntu@VM-40-206-ubuntu:~$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                NAMES
c8d74c1b40fc        nginx               "nginx -g 'daemon ..."   7 minutes ago       Up 7 minutes        0.0.0.0:80->80/tcp   webserver
```
注意: exec后面的bash不能省略

这种操作生产中一般不用，而是使用Dockerfile来定制镜像
## 10. 查看容器的修改
我们进入到容器中修改了nginx的欢迎页，相当于修改了容器的存储层。

```sh
ubuntu@VM-40-206-ubuntu:~$ docker diff webserver
C /root
A /root/.bash_history
C /run
A /run/nginx.pid
C /usr
C /usr/share
C /usr/share/nginx
C /usr/share/nginx/html
C /usr/share/nginx/html/index.html
C /var
C /var/cache
C /var/cache/nginx
A /var/cache/nginx/client_temp
A /var/cache/nginx/fastcgi_temp
A /var/cache/nginx/proxy_temp
A /var/cache/nginx/scgi_temp
A /var/cache/nginx/uwsgi_temp
```
## 11. 提交容器修改

```sh
ubuntu@VM-40-206-ubuntu:~$ docker commit \
>     --author "gomaster.me" \
>     --message "修改了nginx欢迎页" \
>     webserver \
>     nginx:v2
sha256:2668bc9d4355941ff856b33c1b89afef0f16cb26d20e3dd99c634b1e419ec526

ubuntu@VM-40-206-ubuntu:~$ docker images -a
REPOSITORY          TAG                 IMAGE ID            CREATED              SIZE
nginx               v2                  2668bc9d4355        About a minute ago   107MB
nginx               latest              b8efb18f159b        13 days ago          107MB
ubuntu              latest              14f60031763d        2 weeks ago          120MB
ubuntu              14.04               54333f1de4ed        2 weeks ago          188MB
```

dcoker commit可以提交保留镜像的修改，但是我们看到很多无关的内容也都被加了进来。这属于一种黑箱操作。生产中除非被入侵后作为证据保留提交，一般都使用Dockerfile来定制镜像

## 12. 查看镜像历史
`docker history nginx:v2`


```sh
ubuntu@VM-40-206-ubuntu:~$ docker history nginx:v2
IMAGE               CREATED             CREATED BY                                      SIZE                COMMENT
2668bc9d4355        3 minutes ago       nginx -g daemon off;                            97B                 修改了nginx欢迎页
b8efb18f159b        13 days ago         /bin/sh -c #(nop)  CMD ["nginx" "-g" "daem...   0B
<missing>           13 days ago         /bin/sh -c #(nop)  STOPSIGNAL [SIGTERM]         0B
<missing>           13 days ago         /bin/sh -c #(nop)  EXPOSE 80/tcp                0B
<missing>           13 days ago         /bin/sh -c ln -sf /dev/stdout /var/log/ngi...   22B
<missing>           13 days ago         /bin/sh -c apt-get update  && apt-get inst...   52.2MB
<missing>           13 days ago         /bin/sh -c #(nop)  ENV NJS_VERSION=1.13.3....   0B
<missing>           13 days ago         /bin/sh -c #(nop)  ENV NGINX_VERSION=1.13....   0B
<missing>           13 days ago         /bin/sh -c #(nop)  MAINTAINER NGINX Docker...   0B
<missing>           2 weeks ago         /bin/sh -c #(nop)  CMD ["bash"]                 0B
<missing>           2 weeks ago         /bin/sh -c #(nop) ADD file:fa8dd9a679f473a...   55.3MB
ubuntu@VM-40-206-ubuntu:~$
```

## 13. 使用定制镜像
根据之前提交的镜像修改，我们可以指定运行定制过的镜像。

```sh
ubuntu@VM-40-206-ubuntu:~$ docker run --name web2 -d -p 81:80 nginx:v2
0f9d91cbf6339270dbd5f79adb4c8316a9a43314e03afa12b60a902dcf2ff62d
```
将容器nginx运行的80端口转发映射到了宿主机的81端口。

注意端口的运行与转发

注意下面的web4，nginx默认启动在80端口，而容器指定的是-p 82:81 虽然没有冲突可以启动，但是服务是访问不了的
```
ubuntu@VM-40-206-ubuntu:~$ docker ps -a
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS                        NAMES
594db566b3c0        nginx:v2            "nginx -g 'daemon ..."   51 seconds ago       Up 50 seconds       80/tcp, 0.0.0.0:82->81/tcp   web4
c0676feff795        nginx:v2            "nginx -g 'daemon ..."   About a minute ago   Created                                          web3
0f9d91cbf633        nginx:v2            "nginx -g 'daemon ..."   4 minutes ago        Up 4 minutes        0.0.0.0:81->80/tcp           web2
c8d74c1b40fc        nginx               "nginx -g 'daemon ..."   24 minutes ago       Up 24 minutes       0.0.0.0:80->80/tcp           webserver
```

## 14. Dockerfile定制镜像
### Dockerfile分类与指令
####1. 镜像基本构建指令
* FROM
* MAINTAINER
* RUN
* EXPOSE

####2. 指定容器运行时运行的命令
* CMD
* ENTRYPOINT

##### CMD指令

用来提供容器运行的默认命令，与run命令类似，都是执行命令。
**区别**
1. RUN指定的命令是在镜像构建过程中执行的，CMD指定的命令是在容器运行中执行的。
2. 当我们使用docker run命令启动一个容器时，指定了一个容器运行时的命令，那么cmd指令中的命令会被覆盖，不会被执行。也就是说cmd指令是运来指定容器运行时的默认行为。

CMD指令有两种模式
1. exec模式
`CMD ["executable","param1","param2"] `
2. shell模式
`CMD command param1 param2 `
3. 与ENTRYPOINT搭配使用模式
作为ENTRYPOINT指令的默认参数
`CMD ["param1","param2"] `


```sh
~/mynginx$ cat Dockerfile

FROM nginx
MAINTAINER gomaster.me@sina.com "xx@qq.com"
RUN echo "Hello Docker!" > /usr/share/nginx/html/index.html
EXPOSE 80
CMD ["/usr/sbin/nginx","-g","daemon off;"]
```
注意CMD的写法 尤其是最后一个参数`dammon off`后的`;`不能省略，否则启动失败。 

启动  
```sh
docker run -d --name my-nginx -p 80:80 nginx
```

可以使用docker+通常的linux命令方式  

```sh
$ docker top my-nginx
UID                 PID                 PPID                C                   STIME               TTY                 TIME                CMD
root                2106                2090                0                   Aug09               ?                   00:00:00            nginx: master process nginx -g daemon off;
syslog              2136                2106                0                   Aug09               ?                   00:00:00            nginx: worker process
```

##### ENTRYPOINT
ENTRYPOINT指令和CMD指令相似，唯一区别在于其不会被run命令中的执行命令覆盖。 

* exec模式
ENTRYPOINT ["executable","param1","param2"] 
* shell模式
ENTRYPOINT command param1 param2 

注意:使用ENTRYPOINT模式，可以手动 docker run --entrypoint覆盖

##### ENTRYPOINT和CMD组合使用
根据其特点，使用ENTRYPOINT指定执行指令，CMD指定指令默认的参数。
也就是说docker run可以指定CMD覆盖，如果不覆盖则使用Dockerfile中CMD的定义，而ENTRYPOINT则不论如何都会执行。


####3. 目录文件指令
* ADD
* COPY
* VOLUME

##### ADD&&COPY
ADD和COPY都是将文件或目录复制到使用Dockerfile定义的镜像中。支持两个参数 src,dest 来源地址和目标地址。
文件或目录的来源可以是本地地址，也可以是远程的url。
如果是本地地址，必须是构建目录中的相对地址。远程url不推荐使用，更建议使用curl或wget获取文件。
目标路径需要指定镜像中的绝对路径。

区别:
ADD 包含类似tar的解压功能
如果单纯复制文件，Docker推荐使用COPY
`COPY index.html /usr/share/nginx/html/`

```sh
ubuntu@VM-40-206-ubuntu:~/mynginx$ cat Dockerfile
FROM nginx
MAINTAINER gomaster.me@sina.com "xx@qq.com"
#RUN echo "Hello Docker!" > /usr/share/nginx/html/index.html
COPY index.html /usr/share/nginx/html/
EXPOSE 80
CMD ["/usr/sbin/nginx","-g","daemon off;"]
```

####4. 环境设置指令
指定镜像在构建及容器在运行时的环境设置
* WORKDIR
* ENV
* USER

#### 5. 触发器指令
* ONBUILD



### 编辑Dockerfile文件

```sh
~/mynginx$ cat Dockerfile
FROM nginx
RUN echo "Hello Docker!" > /usr/share/nginx/html/index.html
```

### docker build

```sh
~/mynginx$ docker build -t nginx:v3 .
Sending build context to Docker daemon  2.048kB
Step 1/2 : FROM nginx
 ---> b8efb18f159b
Step 2/2 : RUN echo "Hello Docker!" > /usr/share/nginx/html/index.html
 ---> Running in f49677321de0
 ---> 542a777224ae
Removing intermediate container f49677321de0
Successfully built 542a777224ae
Successfully tagged nginx:v3

$ docker run -it -d -p 85:80 --name web2.0 nginx:v3
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                NAMES
10c1bb9d9329        nginx:v3            "nginx -g 'daemon ..."   3 minutes ago       Up 3 minutes        0.0.0.0:85->80/tcp   web2.0
```

注意 `.` 代表的不是指定路径，而是指定上下文路径(docker enginee是cs结构，而build命令是在server端执行的，如何让服务端获得本地文件呢？上下文路径就特别重要)，Dockerfile中命令指定的路径都是上下文路径，也是相对路径。
所以一般将Dockerfile放到项目的根目录或空目录然后将所需文件复制过来，如果有不需要的文件，可以类似的使用如.gitignore的方式使用.dockerignore文件定义忽略的文件。

下面的图片中有什么优点，有什么缺点呢?
![](media/15022902251875.jpg)

### 其他方式build镜像
### git url
`docker build https://github.com/twang2218/gitlab-ce-zh.git\#:gitlab-9.4.3`
[1.8bug](https://github.com/moby/moby/issues/33686)
gitlab尽量单独部署一台机器，4g以上内存，低配机器安装都是个问题呢。
![-w500](media/15022664509003.jpg)

## 查看unhealthy的容器状态

```sh
ubuntu@VM-40-206-ubuntu:~$ docker ps
CONTAINER ID        IMAGE                          COMMAND                  CREATED             STATUS                   PORTS                                   NAMES
ed235fa04001        twang2218/gitlab-ce-zh:9.4.3   "/assets/wrapper"        2 hours ago         Up 2 hours (unhealthy)   22/tcp, 443/tcp, 0.0.0.0:3000->80/tcp   inspiring_mcclintock
5d3c556c20ee        nginx:v3                       "nginx -g 'daemon ..."   3 hours ago         Up 3 hours               0.0.0.0:8888->80/tcp                    web-nginxv3
ubuntu@VM-40-206-ubuntu:~$ docker inspect --format '{{json .State.Health}}' inspiring_mcclintock | python -m json.tool
{
    "FailingStreak": 136,
    "Log": [
        {
            "End": "2017-08-09T18:18:54.473248048+08:00",
            "ExitCode": -1,
            "Output": "rpc error: code = 2 desc = containerd: container not found",
            "Start": "2017-08-09T18:18:54.467226404+08:00"
        },
        {
            "End": "2017-08-09T18:19:54.478155603+08:00",
            "ExitCode": -1,
            "Output": "rpc error: code = 2 desc = containerd: container not found",
            "Start": "2017-08-09T18:19:54.473353039+08:00"
        },
        {
            "End": "2017-08-09T18:20:54.485160036+08:00",
            "ExitCode": -1,
            "Output": "rpc error: code = 2 desc = containerd: container not found",
            "Start": "2017-08-09T18:20:54.478275134+08:00"
        },
        {
            "End": "2017-08-09T18:21:54.490044376+08:00",
            "ExitCode": -1,
            "Output": "rpc error: code = 2 desc = containerd: container not found",
            "Start": "2017-08-09T18:21:54.485308758+08:00"
        },
        {
            "End": "2017-08-09T18:22:54.497147517+08:00",
            "ExitCode": -1,
            "Output": "rpc error: code = 2 desc = containerd: container not found",
            "Start": "2017-08-09T18:22:54.490164636+08:00"
        }
    ],
    "Status": "unhealthy"
}
```
## gitlab中国社区镜像
`docker run -d -p 3000:80 twang2218/gitlab-ce-zh:9.4.3 `   
[一个很不错的gitlab社区版本](https://github.com/twang2218/gitlab-ce-zh)

## DockerHub加速器
[腾讯云docker等加速服务](https://github.com/tencentyun/qcloud-documents/blob/master/product/%E8%AE%A1%E7%AE%97%E4%B8%8E%E7%BD%91%E7%BB%9C/%E4%BA%91%E6%9C%8D%E5%8A%A1%E5%99%A8/Linux%E7%B3%BB%E7%BB%9F%E4%BA%91%E6%9C%8D%E5%8A%A1%E5%99%A8%E8%BF%90%E7%BB%B4%E6%89%8B%E5%86%8C/%E4%BD%BF%E7%94%A8%E8%85%BE%E8%AE%AF%E4%BA%91%E8%BD%AF%E4%BB%B6%E6%BA%90%E5%8A%A0%E9%80%9F%E8%BD%AF%E4%BB%B6%E5%8C%85%E4%B8%8B%E8%BD%BD%E5%92%8C%E6%9B%B4%E6%96%B0.md)  
[腾讯各镜像](https://market.qcloud.com/categories/67)  
[腾讯云dockerHub加速器](https://www.qcloud.com/document/product/457/7207)  
[docker普通用户不使用sudo的方法](http://www.cnblogs.com/ksir16/p/6530587.html)  

* DaoCloud

[daocloud docker镜像加速器](https://www.daocloud.io/mirror#accelerator-doc)
```
curl -sSL https://get.daocloud.io/daotools/set_mirror.sh | sh -s http://bbfa5e62.m.daocloud.io Copy

```
>该脚本可以将 --registry-mirror 加入到你的 Docker 配置文件 /etc/default/docker 中。适用于 Ubuntu14.04、Debian、CentOS6 、CentOS7、Fedora、Arch Linux、openSUSE Leap 42.1，其他版本可能有细微不同。更多详情请访问文档。

* 阿里云
[阿里云Docker 镜像加速器](https://yq.aliyun.com/articles/29941)

## 一些问题

### 建立docker用户组
1. 建立docker组:  
`$ sudo groupadd docker`
2. 将当前用户加入docker组:  
`$ sudo usermod -aG docker $USER`

### why daemon off with nginx on docker
[docker运行nginx为什么要使用 daemon off](https://segmentfault.com/a/1190000009583997)

### docker网络入门 端口转发
[Docker网络原则入门：EXPOSE，-p，-P，-link](http://dockone.io/article/455)


