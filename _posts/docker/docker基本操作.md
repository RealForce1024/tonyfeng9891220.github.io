# docker基础操作
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
* -i --interactive=true|false false是默认  代表:交互式
* -t --tty=true|false false是默认   代表:终端

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


