```sh
☁  aaaa  ll
total 24
-rw-r--r--  1 fqc  staff    82B Apr 24 12:05 hello.go
-rw-r--r--  1 fqc  staff    82B Apr 24 12:05 hello2.go
-rw-r--r--  1 fqc  staff    82B Apr 24 12:05 hello3.go
☁  aaaa  rm $(ls .|grep hello)
```

## tail
```sh
tail -f -n 3 1.txt
```


## grep or

```sh
☁  .ssh  git config -l | grep "user.name\|user.email"
user.name=fqc
user.email=feng-qichao@qq.com
```
注意 grep 后面正则需要配置


