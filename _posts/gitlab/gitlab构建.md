# gitlab ce install
[官网Omnibus package installation (recommended)](https://about.gitlab.com/installation/#ubuntu?version=ce)
注意
0. 服务器内存至少选择4G以上，根据实际使用，最好是8G以上。
1. 邮箱服务器之后再单独配置，该步骤安装时跳过即可。
2. 最后一步执行安装的时候，指定替换"git@gitlab.com"的服务器地址。或者之后配置也行。ip+port[port可以省略，默认80端口]
也就是真正执行了3条命令。(docker方式更简单，但是有些网络的坑，docker方式更适合快速检验新版本的功能)


```sh
sudo apt-get install -y curl openssh-server ca-certificates
curl -sS https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.deb.sh | sudo bash
sudo EXTERNAL_URL="http://54.222.128.44" apt-get install gitlab-ce
```

## 机型
aws t2.large 2vcpu 8G 50G存储  这些参数都可以动态扩容

## 安装
![](media/15096111888426.jpg)

也可以配置成Ip+端口的形式
![](media/15096121183789.jpg)



## 访问



默认第一次访问时需要修改root账户密码。
![](media/15096123648006.jpg)

![](media/15096114570760.jpg)


#新建工程后关联本地
Command line instructions

## http or ssh?

首选ssh方式
配置方式，很简单，可以参照官网。


```ssh
☁  git-work  git config -l | grep "user.name\|user.email"
user.name=fqc
user.email=feng-qichao@qq.com
☁  git-work  pbcopy < ~/.ssh/id_rsa.pub
☁  git-work  ssh -T git@52.80.26.254
Welcome to GitLab, Administrator!


注意:即使有端口方式，也不需要再单独加。(而docker方式则通不过，需要单独设置)
```
http方式，注意 clone后面的地址方式为http:....
![](media/15096137748661.jpg)

ssh方式，注意git clone 方式为git@...
![](media/15096137852602.jpg)

## Git global setup

```sh
git config -l | grep "user.name\|user.email" ## 查看下当前的配置

git config --global user.name "fqc"
git config --global user.email "feng-qichao@qq.com"
```

## Create a new repository

```sh
git clone http://e81b42d5336f/root/my-git-test.git ==>
##(将e81b..替换成相应的ip:port)  还要注意 ssh 或 http方式

git clone http://54.222.158.102:3000/root/my-git-test.git
cd my-git-test
touch README.md
git add README.md
git commit -m "add README"
git push -u origin master
```

## Existing folder

```sh
cd existing_folder
git init
git remote add origin http://e81b42d5336f/root/my-git-test.git
git add .
git commit -m "Initial commit"
git push -u origin master
```

## Existing Git repository

```sh
cd existing_repo
git remote add origin http://e81b42d5336f/root/my-git-test.git
git push -u origin --all
git push -u origin --tags
```

## gitlab 设置时区
方式一、 修改配置文件
find / -name gitlab.yml

```sh
~$ sudo find / -name gitlab.yml
/opt/gitlab/embedded/service/gitlab-rails/config/gitlab.yml
/var/opt/gitlab/gitlab-rails/etc/gitlab.yml
~$ sudo vim /opt/gitlab/embedded/service/gitlab-rails/config/gitlab.yml
```

![-w300](media/15096153607185.jpg)

方式二、 [官网方式](https://docs.gitlab.com/ce/workflow/timezone.html)


## 设置时间12/24小时格式
发现默认的设置是12小时
[How to change the time format display to 24h](https://gitlab.com/gitlab-org/gitlab-ce/issues/15670)

![](media/15096206610821.jpg)


经过 gitlab-ctl reconfigure之后变成了..所以修改被置回去了。
![-w300](media/15096209532755.jpg)

![](media/15096207985720.jpg)

## ci cd
![](media/15096344125666.jpg)

[getting started with ci](http://52.80.26.254:9999/help/ci/quick_start/README)


```sh
GitLab offers a continuous integration service. If you
add a .gitlab-ci.yml file to the root directory of your repository,
and configure your GitLab project to use a Runner, then each commit or
push, triggers your CI pipeline.
On any push to your repository, GitLab will look for the .gitlab-ci.yml
file and start jobs on Runners according to the contents of the file,
for that commit.
Because .gitlab-ci.yml is in the repository and is version controlled, old
versions still build successfully, forks can easily make use of CI, branches can
have different pipelines and jobs, and you have a single source of truth for CI.
You can read more about the reasons why we are using .gitlab-ci.yml in our
blog about it.

in .gitlab.yml always use spaces, not tabs.

1. Add .gitlab-ci.yml to the root directory of your repository
2. Configure a Runner
```
只需要提供.gitlab-ci.yml就可以跑，但不会执行完成，因为还需要配置真正的执行引擎，gitlab runner

![](media/15096353070040.jpg)


![](media/15096353925778.jpg)


### 配置gitlab runner
![-w300](media/15096355102875.jpg)

![](media/15096354979292.jpg)


[官网gitlab-ci-cd安装](https://about.gitlab.com/features/gitlab-ci-cd/)
[get started gitlab-ci](http://52.80.26.254:9999/help/ci/quick_start/README)
[gitlab CI/CD](https://docs.gitlab.com/ce/ci/quick_start/README.html)
[gitlab-runner安装](https://docs.gitlab.com/runner/install/linux-repository.html)
[gitlab-runner register](https://docs.gitlab.com/runner/register/index.html)
[gitlab 中文官网](https://docs.gitlab.com.cn/runner/register/index.html)

![](media/15096364094906.jpg)

![](media/15096363850100.jpg)

![](media/15096369505369.jpg)


![](media/15096374141111.jpg)



![](media/15096374406004.jpg)


![](media/15096375876990.jpg)

需要先准备好docker环境。
另外gitlab-runner最好是在与gitlab分离的环境上。



安装配置好docker，一切就绪，也可以正常执行了，但是在build项目时候太慢了，需要修改下maven
中央仓库
![](media/15096386196370.jpg)
![](media/15096390908952.jpg)


无计可施，build stage 最后跑了14分钟多，不过有意思的是看到后面的构建了缓存。
![](media/15096398330813.jpg)

开始跑test stage，
![](media/15096399261263.jpg)
居然也跑了12分钟，不过庆幸的是也看到了create chache
![](media/15096410117092.jpg)

这次工程提交到构建完成一共用了26分钟......还好有cache
![](media/15096410677931.jpg)

哈哈，终于只跑了18秒!!!yes!!!!!!
![](media/15096413025798.jpg)

### 文章
[gitlab+runner](https://jicki.me/2016/10/16/GitLab-Runner-CN/) 该作者也写了spring cloud ,k8s的内容挺不错
[搭建自己的 Gitlab CI Runner](https://lutaonan.com/blog/gitlab-ci-runner/)
[使用Gitlab-Runner Docker 构建 node 项目](http://yangblink.com/2016/11/21/%E4%BD%BF%E7%94%A8Gitlab-Runner-Docker-%E6%9E%84%E5%BB%BA-node-%E9%A1%B9%E7%9B%AE/) 
## 注意gitlab-reconfige命令会覆盖或清空掉修改过的配置
# gitlab权限
[关于gitlab若干权限问题](http://blog.smallmuou.xyz/git/2016/03/11/%E5%85%B3%E4%BA%8EGitlab%E8%8B%A5%E5%B9%B2%E6%9D%83%E9%99%90%E9%97%AE%E9%A2%98.html)

# 完美git flow

配合[gitflow](https://www.git-tower.com/learn/git/ebook/cn/command-line/advanced-topics/git-flow) 使得微服务的代码分支管理规范相当完美。

[](![](media/15098093825632.jpg)

`git flow feature start $featureName永远都是从develop分支开始`

[CI/CD Git Flow](https://mritd.me/2017/09/05/git-flow-note/)

## 代码分支保护

![](media/15098156032796.jpg)

![](media/15098155910385.jpg)


![](media/15098104258711.jpg)





![](media/15098155730858.jpg)

# 官方文档
[官网文档汇总](https://docs.gitlab.com/ee/README.html)



![](media/15096221284617.jpg)




# 延伸
[gitlab docker hub](https://hub.docker.com/u/gitlab/)
[docker部署gitlab应用](http://www.jianshu.com/p/05e3bb375f64)
[docker gitlab 如何git clone？](https://segmentfault.com/q/1010000007780951)

[基于docker+gitlabCI搭建私有集成环境](http://blog.kazaff.me/2016/06/15/%E5%9F%BA%E4%BA%8Edocker+gitlabCI%E6%90%AD%E5%BB%BA%E7%A7%81%E6%9C%89%E6%8C%81%E7%BB%AD%E9%9B%86%E6%88%90%E7%8E%AF%E5%A2%83/)
## gitlab container registry
[gitlab container registry](http://www.infoq.com/cn/news/2016/05/gitlab-docker-registry)

[GitLab Container Registry](https://about.gitlab.com/2016/05/23/gitlab-container-registry/)
1. First, ask your system administrator to enable GitLab Container Registry following the administration documentation.让管理员后台开启gitlab容器注册服务
2. After that, you will be allowed to enable Container Registry for your project.
然后你在项目中就可以使用该服务了..!so easy！

## rancher CI/CD
[如何使用GitLab和Rancher构建CI/CD流水线 – Part 2](https://www.cnrancher.com/run-gitlab-rancher-2/)

## 延伸阅读 搭建 ci/cd

[Gitlab 安装使用教程（待更新）](http://wentuotuo.com/2017/10/23/Tool/gitlab-install-and-config/)
[金科代码管理中心使用手册](http://wentuotuo.com/2017/10/24/Tool/jk-gitlab-user-guide/)
[gitlab搭建自己的ci/cd](https://blog.m31271n.com/2017/08/22/%E9%80%9A%E8%BF%87-GitLab-%E6%9E%84%E5%BB%BA%E8%87%AA%E5%B7%B1%E7%9A%84-CI-CD-Pipeline/)
[ci/cd的理论基础](https://blog.m31271n.com/2017/08/21/CI-CD-%E7%9A%84%E7%90%86%E8%AE%BA%E4%B8%8E%E5%AE%9E%E8%B7%B5/)
[通过 GitLab 构建自己的 CI / CD Pipeline](https://blog.m31271n.com/2017/08/22/%E9%80%9A%E8%BF%87-GitLab-%E6%9E%84%E5%BB%BA%E8%87%AA%E5%B7%B1%E7%9A%84-CI-CD-Pipeline/#GitLab-GitLab-CI-GitLab-Container-Registry)


[比较特殊的docker安装gitlab方式.....](http://www.zhishichong.com/article/109155)-->个人觉得多次一举。何不一次搞定？

## 基于gitlab的ci/cd
[基于gitlab与docker的cd](https://www.slahser.com/2016/09/08/%E5%9F%BA%E4%BA%8EGitlab%E4%B8%8EDocker%E7%9A%84CD/)

