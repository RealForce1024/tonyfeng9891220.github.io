# etcd

## 安装
目前社区最新为v3.X(rpc)，但主流还是v2.X(http+json)  
[各版本下载地址](https://github.com/coreos/etcd/tags)   
最新的2.X系列为v2.3.8

1. 下载 
wget https://github.com/coreos/etcd/releases/download/v2.3.8/etcd-v2.3.8-linux-amd64.tar.gz 
2. 解压
tar -zxvf etcd-v2.3.8-linux-amd64.tar.gz -C /usr/local/
3. 配置环境变量
vim /etc/profile  
```sh
export ETCD=/usr/local/etcd
export PATH=$...:$ETCD:$PATH
```
4. 验证
任意路径下执行命令`etcd --version`  

```sh
etcd Version: 2.3.8
Git SHA: 7e4fc7e
Go Version: go1.7.5
Go OS/Arch: linux/amd64
```

## 认识etcd
### 目录结构
```sh
$ cd $ETCD 
$ ls
Documentation  etcd  etcdctl  README-etcdctl.md  README.md
```
Documention和*.md都是文档说明，很重要，但是我们现在需要快速掌握etcd的使用，可以先忽略，可以回头在过来查。  
所以剩下的etcd和etcdctl两个简单的二进制执行命令是最为重要的啦。

可以通过帮助命令查看etcd支持的参数 
`etcd --help 或etcd -h`,我们可以看到
- etcd [flags] 开启etcd的服务
- etcd --version 查看etcd的版本
- etcd -h | --help 命令帮助
- 以及其他更具体的参数

```sh
root@iZbp1gbi9f8xfa097dfgn6Z:/usr/local/etcd# etcd --help
usage: etcd [flags]
       start an etcd server

       etcd --version
       show the version of etcd

       etcd -h | --help
       show the help information about etcd


member flags:

        --name 'default'
                human-readable name for this member.
        --data-dir '${name}.etcd'
                path to the data directory.
        --wal-dir ''
                path to the dedicated wal directory.
        --snapshot-count '10000'
                number of committed transactions to trigger a snapshot to disk.
        --heartbeat-interval '100'
                time (in milliseconds) of a heartbeat interval.
        --election-timeout '1000'
                time (in milliseconds) for an election to timeout. See tuning documentation for details.
        --listen-peer-urls 'http://localhost:2380,http://localhost:7001'
                list of URLs to listen on for peer traffic.
        --listen-client-urls 'http://localhost:2379,http://localhost:4001'
                list of URLs to listen on for client traffic.
        --max-snapshots '5'
                maximum number of snapshot files to retain (0 is unlimited).
        --max-wals '5'
                maximum number of wal files to retain (0 is unlimited).
        -cors ''
                comma-separated whitelist of origins for CORS (cross-origin resource sharing).


clustering flags:

        --initial-advertise-peer-urls 'http://localhost:2380,http://localhost:7001'
                list of this member's peer URLs to advertise to the rest of the cluster.
        --initial-cluster 'default=http://localhost:2380,default=http://localhost:7001'
                initial cluster configuration for bootstrapping.
        --initial-cluster-state 'new'
                initial cluster state ('new' or 'existing').
        --initial-cluster-token 'etcd-cluster'
                initial cluster token for the etcd cluster during bootstrap.
                Specifying this can protect you from unintended cross-cluster interaction when running multiple clusters.
        --advertise-client-urls 'http://localhost:2379,http://localhost:4001'
                list of this member's client URLs to advertise to the public.
                The client URLs advertised should be accessible to machines that talk to etcd cluster. etcd client libraries
parse these URLs to connect to the cluster.
        --discovery ''
                discovery URL used to bootstrap the cluster.
        --discovery-fallback 'proxy'
                expected behavior ('exit' or 'proxy') when discovery services fails.
        --discovery-proxy ''
                HTTP proxy to use for traffic to discovery service.
        --discovery-srv ''
                dns srv domain used to bootstrap the cluster.
        --strict-reconfig-check
                reject reconfiguration requests that would cause quorum loss.

proxy flags:

        --proxy 'off'
                proxy mode setting ('off', 'readonly' or 'on').
        --proxy-failure-wait 5000
                time (in milliseconds) an endpoint will be held in a failed state.
        --proxy-refresh-interval 30000
                time (in milliseconds) of the endpoints refresh interval.
        --proxy-dial-timeout 1000
                time (in milliseconds) for a dial to timeout.
        --proxy-write-timeout 5000
                time (in milliseconds) for a write to timeout.
        --proxy-read-timeout 0
                time (in milliseconds) for a read to timeout.


security flags:

        --ca-file '' [DEPRECATED]
                path to the client server TLS CA file. '-ca-file ca.crt' could be replaced by '-trusted-ca-file ca.crt -client-cert-auth' and etcd will perform the same.
        --cert-file ''
                path to the client server TLS cert file.
        --key-file ''
                path to the client server TLS key file.
        --client-cert-auth 'false'
                enable client cert authentication.
        --trusted-ca-file ''
                path to the client server TLS trusted CA key file.
        --peer-ca-file '' [DEPRECATED]
                path to the peer server TLS CA file. '-peer-ca-file ca.crt' could be replaced by '-peer-trusted-ca-file ca.crt -peer-client-cert-auth' and etcd will perform the same.
        --peer-cert-file ''
                path to the peer server TLS cert file.
        --peer-key-file ''
                path to the peer server TLS key file.
        --peer-client-cert-auth 'false'
                enable peer client cert authentication.
        --peer-trusted-ca-file ''
                path to the peer server TLS trusted CA file.

logging flags

        --debug 'false'
                enable debug-level logging for etcd.
        --log-package-levels ''
                specify a particular log level for each etcd package (eg: 'etcdmain=CRITICAL,etcdserver=DEBUG').

unsafe flags:

Please be CAUTIOUS when using unsafe flags because it will break the guarantees
given by the consensus protocol.

        --force-new-cluster 'false'
                force to create a new one-member cluster.


experimental flags:

        --experimental-v3demo 'false'
                enable experimental v3 demo API.
        --experimental-auto-compaction-retention '0'
                auto compaction retention in hour. 0 means disable auto compaction.
        --experimental-gRPC-addr '127.0.0.1:2378'
                gRPC address for experimental v3 demo API.

profiling flags:
        --enable-pprof 'false'
                Enable runtime profiling data via HTTP server. Address is at client URL + "/debug/pprof"
```

## 启动etcd

运行 etcd(直接执行`etcd`命令)将默认组建一个两个节点的集群。数据库服务端默认监听在 2379 和 4001 端口，etcd 实例监听在 2380 和 7001 端口。显示类似如下的信息：

```sh
$ etcd 

root@iZbp1gbi9f8xfa097dfgn6Z:/usr/local/etcd# etcd
2017-06-02 14:30:45.285585 I | etcdmain: etcd Version: 2.3.8
2017-06-02 14:30:45.285708 I | etcdmain: Git SHA: 7e4fc7e
2017-06-02 14:30:45.285749 I | etcdmain: Go Version: go1.7.5
2017-06-02 14:30:45.285787 I | etcdmain: Go OS/Arch: linux/amd64
2017-06-02 14:30:45.285824 I | etcdmain: setting maximum number of CPUs to 1, total number of available CPUs is 1
2017-06-02 14:30:45.285857 W | etcdmain: no data-dir provided, using default data-dir ./default.etcd
2017-06-02 14:30:45.286116 I | etcdmain: listening for peers on http://localhost:2380
2017-06-02 14:30:45.286230 I | etcdmain: listening for peers on http://localhost:7001
2017-06-02 14:30:45.286315 I | etcdmain: listening for client requests on http://localhost:2379
2017-06-02 14:30:45.286429 I | etcdmain: listening for client requests on http://localhost:4001
2017-06-02 14:30:45.286647 I | etcdserver: name = default
2017-06-02 14:30:45.286684 I | etcdserver: data dir = default.etcd
2017-06-02 14:30:45.286705 I | etcdserver: member dir = default.etcd/member
2017-06-02 14:30:45.286725 I | etcdserver: heartbeat = 100ms
2017-06-02 14:30:45.286745 I | etcdserver: election = 1000ms
2017-06-02 14:30:45.286770 I | etcdserver: snapshot count = 10000
2017-06-02 14:30:45.286796 I | etcdserver: advertise client URLs = http://localhost:2379,http://localhost:4001
2017-06-02 14:30:45.286819 I | etcdserver: initial advertise peer URLs = http://localhost:2380,http://localhost:7001
2017-06-02 14:30:45.286850 I | etcdserver: initial cluster = default=http://localhost:2380,default=http://localhost:7001
2017-06-02 14:30:45.290037 I | etcdserver: starting member ce2a822cea30bfca in cluster 7e27652122e8b2ae
2017-06-02 14:30:45.290125 I | raft: ce2a822cea30bfca became follower at term 0
2017-06-02 14:30:45.290159 I | raft: newRaft ce2a822cea30bfca [peers: [], term: 0, commit: 0, applied: 0, lastindex: 0, lastterm: 0]
2017-06-02 14:30:45.290184 I | raft: ce2a822cea30bfca became follower at term 1
2017-06-02 14:30:45.290313 I | etcdserver: starting server... [version: 2.3.8, cluster version: to_be_decided]
2017-06-02 14:30:45.290534 E | etcdmain: failed to notify systemd for readiness: No socket
2017-06-02 14:30:45.290567 E | etcdmain: forgot to set Type=notify in systemd service file?
2017-06-02 14:30:45.293285 N | etcdserver: added local member ce2a822cea30bfca [http://localhost:2380 http://localhost:7001] to cluster 7e27652122e8b2ae
2017-06-02 14:30:45.690424 I | raft: ce2a822cea30bfca is starting a new election at term 1
2017-06-02 14:30:45.690587 I | raft: ce2a822cea30bfca became candidate at term 2
2017-06-02 14:30:45.690691 I | raft: ce2a822cea30bfca received vote from ce2a822cea30bfca at term 2
2017-06-02 14:30:45.690752 I | raft: ce2a822cea30bfca became leader at term 2
2017-06-02 14:30:45.690801 I | raft: raft.node: ce2a822cea30bfca elected leader ce2a822cea30bfca at term 2
2017-06-02 14:30:45.691105 I | etcdserver: setting up the initial cluster version to 2.3
2017-06-02 14:30:45.693493 N | etcdserver: set the initial cluster version to 2.3
2017-06-02 14:30:45.693603 I | etcdserver: published {Name:default ClientURLs:[http://localhost:2379 http://localhost:4001]} to cluster 7e27652122e8b2ae
```

运行`etcdctl`客户端命令进行测试，进行键值对的设置和获取 `testkey:"hello world"`:
可以先让`etcd &` 后台运行，或者托管supervisor中
查看etcd的进程，可以用来etcd的停止。
`ps -ef | grep etcd`
`lsof -i :2380`

```sh
$ etcdctl set testkey "hello etcd"
hello etcd
$ etcdctl get testkey
hello etcd
$ etcdctl get testkey2
Error:  100: Key not found (/testkey2) [6]
```
看到上面的键值对的存取成功，说明etcd的服务已经启动成功！ yeap~

etcd支持restful的api，所以可以使用http的方式调用，
我们看到启动信息中的这几行

```sh
etcdmain: listening for client requests on http://localhost:2379
etcdmain: listening for client requests on http://localhost:4001
etcdserver: advertise client URLs = http://localhost:2379,http://localhost:4001
```

`curl -L http://localhost:4001/v2/keys/testkey`

```sh
$ curl -L http://localhost:4001/v2/get/testkey
404 page not found
$ curl -L http://localhost:4001/v2/keys/testkey
{"action":"get","node":{"key":"/testkey","value":"hello etcd","modifiedIndex":6,"createdIndex":6}}
$ curl -L http://localhost:2379/v2/keys/testkey
{"action":"get","node":{"key":"/testkey","value":"hello etcd","modifiedIndex":6,"createdIndex":6}}
$ curl -L http://localhost:2380/v2/keys/testkey
404 page not found
```


--------------------------------------
参考资料

[docker从入门到实践](https://www.gitbook.com/book/yeasy/docker_practice/details)
[etcd学习笔记](https://skyao.gitbooks.io/learning-etcd3/content/installation/)
大神 敖小剑https://www.gitbook.com/@skyao
grpc etcd3 微服务 spring netty 
[在 Docker 上运行一个 RESTful 风格的微服务](https://segmentfault.com/a/1190000002930500)
[运行在 Docker 上的微服务 - 服务发现与注册](https://segmentfault.com/a/1190000002943994)

[使用Etcd和Haproxy做Docker服务发现](https://yushuangqi.com/blog/2016/shi-yong--etcd-he--haproxy-zuo--docker-fu-wu-fa-xian.html)




 