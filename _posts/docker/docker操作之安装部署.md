## docker安装
官网docker.com get docker

### 检查环境
- linux内核版本>=3.10
`uname -a`
- 存储驱动Device Mapper
`ls -l /sys/class/misc/device-mapper`

如下执行成功则说明具备安装docker的环境。
```sh
$ uname -a
Linux iZbp1f7qdocvjqdy5yu701Z 4.4.0-63-generic #84-Ubuntu SMP Wed Feb 1 17:20:32 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux
$ ls -l /sys/class/misc/device-mapper
lrwxrwxrwx 1 root root 0 Apr 11 07:51 /sys/class/misc/device-mapper -> ../../devices/virtual/misc/device-mapper
```

### 安装
- 官网
[docker 官网](https://docs.docker.com/install/linux/docker-ce/ubuntu/)  
[docker rancher版本对照](https://rancher.com/docs/rancher/v1.6/zh/hosts/ )  
https://rancher.com/docs/rancher/v1.6/en/hosts/#supported-docker-versions

```
$ sudo apt-get install docker-ce=5:18.09.5~3-0~ubuntu-bionic docker-ce-cli=5:18.09.5~3-0~ubuntu-bionic containerd.io
```

dao docker 镜像加速器
```
https://www.daocloud.io/mirror
```

- 阿里云方式
1. `sudo apt-get install -y curl`

2. 阿里云内网 `curl -sSL http://acs-public-mirror.oss-cn-hangzhou.aliyuncs.com/docker-engine/internet | sh -`
 
 这里可能已经更新了，所以可以参照阿里云官方文档操作
 `https://www.alibabacloud.com/help/zh/doc-detail/60742.htm`


3. `docker version`
```
Client:
 Version:      17.04.0-ce
 API version:  1.28
 Go version:   go1.7.5
 Git commit:   4845c56
 Built:        Mon Apr  3 18:07:42 2017
 OS/Arch:      linux/amd64

Server:
 Version:      17.04.0-ce
 API version:  1.28 (minimum version 1.12)
 Go version:   go1.7.5
 Git commit:   4845c56
 Built:        Mon Apr  3 18:07:42 2017
 OS/Arch:      linux/amd64
 Experimental: false
```
4. sudo docker ubuntu echo "hello world"
```java
root@iZbp1f7qdocvjqdy5yu701Z:~# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
hello-world         latest              48b5124b2768        2 months ago        1.84kB
root@iZbp1f7qdocvjqdy5yu701Z:~# docker run ubuntu echo "hello"
Unable to find image 'ubuntu:latest' locally
latest: Pulling from library/ubuntu
c62795f78da9: Pull complete
d4fceeeb758e: Pull complete
5c9125a401ae: Pull complete
0062f774e994: Pull complete
6b33fd031fac: Pull complete
Digest: sha256:c2bbf50d276508d73dd865cda7b4ee9b5243f2648647d21e3a471dd3cc4209a0
Status: Downloaded newer image for ubuntu:latest
hello
root@iZbp1f7qdocvjqdy5yu701Z:~# docker run ubuntu echo "hello"
hello
root@iZbp1f7qdocvjqdy5yu701Z:~#
```
5. sudo adduser docker 

```sh
root@iZbp1f7qdocvjqdy5yu701Z:~# docker run ubuntu echo "hello"
hello
root@iZbp1f7qdocvjqdy5yu701Z:~# adduser docker
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LC_CTYPE = "UTF-8",
	LANG = "en_US.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to a fallback locale ("en_US.UTF-8").
adduser: The user `docker` already exists.

root@iZbp1f7qdocvjqdy5yu701Z:~# su docker
To run a command as administrator (user "root"), use "sudo <command>".
See "man sudo_root" for details.

docker@iZbp1f7qdocvjqdy5yu701Z:/root$ pwd
/root
docker@iZbp1f7qdocvjqdy5yu701Z:/root$ cd ~

docker@iZbp1f7qdocvjqdy5yu701Z:~$ ls
docker@iZbp1f7qdocvjqdy5yu701Z:~$ pwd
/home/docker


```
```sh
docker@iZbp1f7qdocvjqdy5yu701Z:~$ sudo docker run ubuntu echo "hello world"
sudo: unable to resolve host iZbp1f7qdocvjqdy5yu701Z
[sudo] password for docker:
Sorry, try again.
[sudo] password for docker:
hello world

docker@iZbp1f7qdocvjqdy5yu701Z:~$ sudo docker run ubuntu echo "hello world"
hello world

docker@iZbp1f7qdocvjqdy5yu701Z:~$ sudo groupadd docker
groupadd: group 'docker' already exists

docker@iZbp1f7qdocvjqdy5yu701Z:~$ sudo gpasswd -a ${USER} docker
Adding user docker to group docker

docker@iZbp1f7qdocvjqdy5yu701Z:~$ sudo service docker restart
docker@iZbp1f7qdocvjqdy5yu701Z:~$
```
将docker加入到docker用户组，不使用root权限，使用一般用户docker也可以使用sudo
```sh
docker@iZbp1f7qdocvjqdy5yu701Z:~$ docker run ubuntu echo 'hello world'
hello world
docker@iZbp1f7qdocvjqdy5yu701Z:~$

docker@iZbp1f7qdocvjqdy5yu701Z:~$ docker version
Client:
 Version:      17.04.0-ce
 API version:  1.28
 Go version:   go1.7.5
 Git commit:   4845c56
 Built:        Mon Apr  3 18:07:42 2017
 OS/Arch:      linux/amd64

Server:
 Version:      17.04.0-ce
 API version:  1.28 (minimum version 1.12)
 Go version:   go1.7.5
 Git commit:   4845c56
 Built:        Mon Apr  3 18:07:42 2017
 OS/Arch:      linux/amd64
 Experimental: false
```












