# Ansible
## install

**任何管理系统受益于被管理的机器在主控机附近运行.如果在云中运行,可以考虑在使用云中的一台机器来运行Ansible.**

mac(主控机)安装即可，远程集群无需安装

```sh
sudo pip install ansible
```

ubuntu


```sh
sudo apt-get install software-properties-common
sudo apt-add-repository ppa:ansible/ansible
sudo apt-get update
sudo apt-get install ansible
```

## 验证
`ansible --version`

## 添加一台机器


```sh
sudo vim /etc/ansible/hosts
$add ip
```


```sh
$ ssh-keygen
$ ssh-copy-id remote-ip
$ ssh remote-ip
```

```sh
vim /etc/hosts
ip AreaDNSName
```

```sh
ansile all -m ping
ansible aliyun -m ping
```

`ansible all -m ping ` 
1. ansible命令主体  ansible/ansible-playbook
2. pattern 目标机器或机器群 all指代所有 具体ip或名称 或 正则
3. -m 指定使用的模块 这里使用ping模块
4. 参数 没有则不写

`ansible all -a 'ls' `

![](media/15021998310598.jpg)


## ansible 命令详解
![](media/15022006401620.jpg)
其中 -l 实际是限定于某台机器或某组或符合某正则的机器

![](media/15022012267016.jpg)
![](media/15022013012125.jpg)


![](media/15022013712766.jpg)


![](media/15022014198416.jpg)


![](media/15022015116167.jpg)

![](media/15022015348986.jpg)

![](media/15022015856016.jpg)

![](media/15022018226642.jpg)



![](media/15022017739553.jpg)


![](media/15022018671413.jpg)


![](media/15022018913835.jpg)

![](media/15022019253408.jpg)


![](media/15022019596830.jpg)


![](media/15022021029277.jpg)


## 注意
1. 
![](media/15021858553748.jpg)

2. 如果ssh-copy-id 出现permision denied的问题，不要忘记检查该用户是否允许登录。在腾讯云的使用中，发现默认情况下root是被禁止ssh远程登录的。[腾讯云ubuntu设置root用户登录](http://bbs.qcloud.com/thread-11554-1-1.html)。 对于aws估计也会有该问题，要注意！！


