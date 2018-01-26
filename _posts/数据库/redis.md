# 
[redis面试题](http://www.100mian.com/mianshi/dba/37381.html)

[redis面试题](http://java.jr-jr.com/2015/12/31/redis-vvv/#u9ED1_u540D_u5355_u3001_u5173_u6CE8_u5217_u8868_u3001_u7C89_u4E1D_u5217_u8868_u3001_u53CC_u5411_u5173_u6CE8_u5217_u8868)

& && 
redis-cli shutdown 最优雅方式
kill pid 较为优雅退出
killall redis-server

redis-cli -c

插槽 reshard

# 对key的通用操作

## 三个通配符 * ? []

```sh
127.0.0.1:6379>set site www.hc.com
127.0.0.1:6379> keys sit[e|y] #keys pattern
1) "site"

127.0.0.1:6379> keys s?te
1) "site"
127.0.0.1:6379>
```
## randomKey 
随机返回key，抽奖
## type key
查看键的类型
## exists某键是否存在 

```sh
127.0.0.1:6379> exists age
(integer) 1
127.0.0.1:6379> exists site
(integer) 1
127.0.0.1:6379> exists nnn
(integer) 0
```
## del key

```sh
127.0.0.1:6379> del age
(integer) 1
127.0.0.1:6379> exists age
(integer) 0
127.0.0.1:6379> keys *
1) "name"
2) "site"
```

## rename key

```sh
127.0.0.1:6379> rename site sitewww
OK
127.0.0.1:6379> keys *
1) "sitewww"
2) "name"
127.0.0.1:6379> rename sitewww site
OK
127.0.0.1:6379> keys *
1) "site"
2) "name"
```

## renamenx key

如果使用rename修改，但是key已经存在，则会覆盖，这样将会很危险。  

```sh
127.0.0.1:6379> keys *
1) "site"
2) "name"
127.0.0.1:6379> set site www.baidu.com
OK
127.0.0.1:6379> set search www.google.com
OK
127.0.0.1:6379> keys *
1) "search"
2) "site"
3) "name"
127.0.0.1:6379> rename site search
OK
127.0.0.1:6379> keys *
1) "search"
2) "name"
127.0.0.1:6379>

```

renameNx 如果键已经存在，则不重命名。 

```sh
127.0.0.1:6379> keys *
1) "search"
2) "name"
127.0.0.1:6379> set site www.baidu.com
OK
127.0.0.1:6379> keys *
1) "search"
2) "site"
3) "name"
127.0.0.1:6379> get search
"www.baidu.com"
127.0.0.1:6379> set search www.google.co
OK
127.0.0.1:6379> keys *
1) "search"
2) "site"
3) "name"
127.0.0.1:6379> renamenx site search
(integer) 0
```

## move key

redis 默认有16个空间

```sh
127.0.0.1:6379> keys *
1) "search"
2) "site"
3) "name"
127.0.0.1:6379> move search 1
(integer) 1
127.0.0.1:6379> keys *
1) "site"
2) "name"
127.0.0.1:6379> select 1
OK
127.0.0.1:6379[1]> keys *
1) "search"
```

## ttl key  查看key的生命周期
-1 永久有效
不存在的-2
ttl key

查看key的存活时间，返回的是存活的秒数。

```sh
127.0.0.1:6379> ttl ke
(integer) -2
127.0.0.1:6379> ttl key
(integer) -2
127.0.0.1:6379> ttl site
(integer) -1
127.0.0.1:6379>

```

## expire key 设置key的有效期

pexpire key 设置生命周期的毫秒数
pttl key 以毫秒数返回生命周期

```sh

127.0.0.1:6379> keys *
1) "site"
2) "name"
127.0.0.1:6379> expire site 10
(integer) 1
127.0.0.1:6379> keys *
1) "site"
2) "name"
127.0.0.1:6379> keys *
1) "site"
2) "name"
127.0.0.1:6379> keys *
1) "site"
2) "name"
127.0.0.1:6379> ttl site
(integer) 2
127.0.0.1:6379> ttl site
(integer) 1
127.0.0.1:6379> ttl site
(integer) -2
127.0.0.1:6379> ttl site
(integer) -2
127.0.0.1:6379> ttl site
(integer) -2
127.0.0.1:6379> keys *
1) "name"

```

## Persist key

```sh
127.0.0.1:6379> set site 1000
OK
127.0.0.1:6379> get site
"1000"
127.0.0.1:6379> keys *
1) "site"
2) "name"
127.0.0.1:6379> expire site 10
(integer) 1
127.0.0.1:6379> ttl site
(integer) 8
127.0.0.1:6379> persist site
(integer) 1
127.0.0.1:6379> ttl site
(integer) -1
127.0.0.1:6379> keys *
1) "site"
2) "name"
```

## flushdb
清空库

```sh
127.0.0.1:6379> keys *
1) "site"
2) "name"
127.0.0.1:6379> flushdb
OK

127.0.0.1:6379> keys *
(empty list or set)
127.0.0.1:6379> select 1
OK
127.0.0.1:6379[1]> keys *
1) "search"
127.0.0.1:6379[1]>

```
# 常用数据结构

## string
set key value [ex 秒数] | [px 毫秒数] [nx]|[xx]


## list
## 


## aof
默认不打开，为了数据不丢失，弥补rdb的不足，所以可以采用aof+rdb的方式。

aof如何平衡慢?

appendfsync always  
appendfsync everysec
appendfsync no  

no-appendfsync-on-rewrite
auto-aof-rewrite-percentage 100
auto-aof-rewrite-min-size 

在dump rdb过程中，aof如果停止同步，会不会丢失丢失。



