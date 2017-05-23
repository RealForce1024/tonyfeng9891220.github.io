nginx -s start/quit(优雅停止)/stop(立刻停止)/reload/reopen
nginx -v
ps -ef | grep nginx // 查看一个程序是否运行
netstat -tln | grep 81 //查看端口的使用情况
lsof -i:81 // 查看端口属于哪个程序