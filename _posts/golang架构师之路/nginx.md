nginx -s start/quit(优雅停止)/stop(立刻停止)/reload/reopen
nginx -v
ps -ef | grep nginx // 查看一个程序是否运行
netstat -tln | grep 81 //查看端口的使用情况
lsof -i:81 // 查看端口属于哪个程序

ps -ef | grep port[6379]
netstat -tunpl | grep redis

aop
互联网技术单一拿出来很简单，业务才是王道

cp mkreleasehdr.sh redis-benchmark redis-check-aof redis-check-dump redis-cli redis-server ../bin