操作步骤
----------------
CI部分
1. 准备环境
宿主机器上
```sh
cd workshop01
./prepare.sh
```
2. 登录虚拟机,启动Gitlab,Jenkins服务 
```sh
vagrant ssh
cd ci
sudo docker-compose up
```

Gitlab  
    http://192.168.33.10/  
    root/abcd1234  
Jenkins   
    http://192.168.33.10:8080/  
    admin/abcd1234

3. 提交代码  
    a. 先删除掉Gitlab仓库中的`hello`工程。(后续脚本都是固定的`hello`命名的工程)  
    b. 新建`helo`工程，设置为public。  
    c. 按照空的`hello`工程提供的readme方案，将本地代码关联到Gitlab仓库。  
      (最核心点在于 git init remote add commit push)

    ```sh
    Git global setup

    git config --global user.name "Administrator"
    git config --global user.email "admin@example.com"

    Create a new repository

    git clone http://a4653c7c23d6/root/hello.git
    cd hello
    touch README.md
    git add README.md
    git commit -m "add README"
    git push -u origin master

    Existing folder

    cd existing_folder
    git init
    git remote add origin http://a4653c7c23d6/root/hello.git
    git add .
    git commit
    git push -u origin master

    Existing Git repository

    cd existing_repo
    git remote add origin http://a4653c7c23d6/root/hello.git
    git push -u origin --all
    git push -u origin --tags
    ```
4. 手动运行web服务   
    a. 使用gradle编译代码    
        虚机中  
    ```sh
    cd ~  
    git clone http://192.168.33.10/root/hello.git  
    cd hello  
    gradle build    
    ```  

    b.java -jar   

    ```sh
    SERVER_PORT=9090 java -jar build/libs/hello-0.0.1-SNAPSHOT.jar
    ```
   访问 192.168.33.10:9090

5. 手动使用docker容器运行web服务       
    a. docker构建镜像  
    ```sh
    cd hello    
    cd src/main/docker    
    cp ../../../build/libs/hello-0.0.1-SNAPSHOT.jar .    
    sudo docker build -t leanms/hello:0.1 .    
    ```

    b. 运行docker容器
    ```sh
    sudo docker run -p 8888:8080 -t leanms/hello:0.1
    ```

    c. 访问web服务  
    http://192.168.33.10:8888 

 6. Jenkins自动化测试
已经关联到了Gitlab，有过`hello`工程    
直接点击构建。  
构建的依据在于Jenkins的Configuration Pipline  script.
注意点:
    1. network需要指定名称，否则默认为父目录_app，例如`--network=ci_hello`，参看`/home/vagrant/hello/src/main/docker/jenkins/docker-compose.yml`文件

    2. git url: "http://gitlab/root/hello.git"不需要改成ip的形式，因为在`~/ci/docker-compose.yml`中`services:gitlab .... jenkins`已经配置。例如：反例证明，在docker-compose中将Gitlab改为gitlab2，然后在Jenkins中执行build将会失败，相应的在pipline中将其改为"http://gitlab2/root/hello.git"将会构建成功。 
    3. 执行过程中 之前可能会有已经启动起来的相同定义的container会有冲突，需要关闭删除下
    ```sh
    sudo docker stop hello_app|| true && sudo docker rm -v hello_app|| true
    ```
    4. pipline script基于grovvy脚本。灵活简洁不冗余。   
    ```sh
    node {
        stage "Checkout"
        git url: "http://gitlab/root/hello.git"

        stage "CheckStyle"
        sh "gradle check --stacktrace"
        archiveCheckstyleResults()

        stage "Build/Analyse/Test"
        sh "git log --format='%H' -n 1 >  src/main/resources/VERSION"
        sh "date >> src/main/resources/VERSION"
        sh "gradle clean build --stacktrace"
        archiveUnitTestResults()

        stage "Generate Docker image"
        sh "pwd"
        sh "cp build/libs/hello-0.0.1-SNAPSHOT.jar src/main/docker"
        sh "cd src/main/docker && docker build -t leanms/hello:0.1 ."

        stage name: "Deploy Docker", concurrency: 1
        sh "docker run -p 8888:8080 -d --network=ci_hello --name hello_app -t leanms/hello:0.1"
        sleep 20

        stage name: "Test APP"
        sh "ping -c 1 hello_app"
        retry(10) {
            sh "netcat -vzw1 hello_app 8080"
            sh "curl http://hello_app:8080"
        }
        sh "curl http://hello_app:8080/version"
        sh "docker rm \$(docker stop \$(docker ps -a -q --filter ancestor=leanms/hello:0.1 --format='{{.ID}}'))"
    }

    def archiveUnitTestResults() {
        step([$class: "JUnitResultArchiver", testResults: "build/**/TEST-*.xml"])
    }

    def archiveCheckstyleResults() {
        step([$class: "CheckStylePublisher",
            canComputeNew: false,
            defaultEncoding: "",
            healthy: "",
            pattern: "build/reports/checkstyle/main.xml",
            unHealthy: ""])
    }

    ```

CD部分
```sh
cd ~/cd 
sh clean.sh
sudo docker-compose up
sh blue_green_upgrade.sh
```



## 开场

精益微服务的持续集成与交付

故事 角色

环境统一准备 vagrant

agenda

破冰
    黄邦伟 喜欢咖啡 新加坡
    吴雪峰 跑步 
    王龙 写代码的
    景韵 devops 
    龚波 户外


    于慧明  看书
    raoke  睡觉 
    陈占文 旅行
    林伟伟 读书

    吕博洋 旅行 
    科之光  旅行 
    石国庆 爬山 
    冯琪超 打球的


精益 微服务 极限编程

大 等待
小 容易push生产

微 对抗单块思维  多个维度最小化

极限编程
scrum 

实战
spring boot
jekins
gitlab
docker
gradle
junit

update push repository

docker\docker compose

讲师、tw
(技术、管理)咨询、外包 技术雷达

tw使命 p1-p3
https://github.com/bahmni 项目

精益 
公司各个业务上的精益...
此次系统、软件架构 集中在代码
瘦身-看板   ：价值流
强调 直通率  而非过度不必要的沟通浪费

- 自动化测试
哪些是不必要的浪费，抽出一些过程， 干..
简单修改代码->上线提交
看瓶颈在哪里  
    测试  等待周期 
    ...
    开发规定..

自动化测试  接口、功能、ui 
开发者的测试与测试人员的测试是否有结合?
开发者和测试者代码互相review
开发者的测试最大化 聚焦 盲点

自己全循环自动化
强调一体化运作

之后会讲 微服务的自测 

- 持续集成
线下传统公司的落地难度
- 持续交付

敏捷 vs 瀑布
微  vs 单体
大数据 soa  not 微
松耦合 vs 独立演进

单一职责 划分的依据  周期、变化  而非类似性
things that change for the same reason stays toghter.

收益
微服务架构、微交付 

返单块 反单一集中式

利益最大化  团队职责

利弊
强化分  边界明确? 如何确定边界?
独立部署
技术多样化

弊
分布式...

分布式计算的
多进程间不好共享，于是自测

动手 
sh prepare.sh
vagrant ssh 

pwd
~
`ci cd temp`  

`cd ci`  
cat docker-compose.yml
查看内部配置 里面 gitlab jenkins

启动docker中的gitlab,jenkins
`sudo docker-compose up`
sudo docker images

宿主机
vim vagrantfile 
private_network  ip的配置
浏览器访问  192.168.33.10
jenkins 192.168.33.10:8080
gitlab账号 jenkins root/admin abcd1234

cd workshop01/hello

gitlab  new project named hello[needed 'hello']
set public 

remove掉原有的
新建后有readme 所有较为方便 
需要修改的是 ip 地址
添加到gitlab即可. done

到本地的代码文件夹中git add commit   push

然后到 虚机里面  home目录中
cd ~
git clone ..  
cd hello
gradle build
SERVER_PORT=9091 java -jar build/libs/hello-0.0.1-SNAPSHOT.jar
访问 192.168.33.10:9090

hr咨询
workshop文件下发

- docker介绍
镜像
not only code , yet have env
以前的状况，从本机到生产环境，需要各种环境配置才能运行。多套呢？之间是否有影响呢？如何隔离
使用docker解决了我们不仅仅提交了代码，连环境也一并提交了呢
`所以我们连环境都交付了!`这是一种全新的软件交付!!docker主要的卖点之一。  

回到虚机 ~
cd 
通过docker build 构建hello

```sh
cd ~
cd hello
cd src/main/docker
cp ../../../build/libs/hello-0.0.1-SNAPSHOT.jar .
sudo docker build -t leanms/hello:0.1 .
sudo docker run -p 8888:8080 -t leanms/hello:0.1
http://192.168.33.10:8888/
```

```sh
sh "sudo docker stop hello_app|| true && sudo docker rm -v hello_app|| true"

cd ~/cd
sudo docker-compose up
sudo ./blue_green_upgrade.sh
```


注意 ci piplie中的task 之前全部运行通过，也无法访问，是由于
TestApp之后删除掉了容器。 每次测试都创建一个容器然后临时跑完就销毁。
如果想通过此处看执行，可以注释掉。意义取决于自己的需求。  

敏捷开发 运维 测试
代码->生产->运维->测试->监控
流水线 输入 产出  最终有个产物出来

- jenkins介绍
pipline  
build now 

强调 虽然写了自动化测试，但是没有形成持续流水线。自动化价值的体现在哪里呢?和代码一样，跑得越多，价值越大。
比如一份代码一月跑一次，跑一次的价值比如拍个脑袋100，一月就只赚100。
流水线的价值就在于天天跑，实时跑，这样将价值最大化。
开发人员提交的代码能够马上的由自动化测试把它拦截下来，能够马上的进行修复。使反馈周期最短。
测试抓取 不断测试...

这里是jenkins2.0，增加pipline强大特性。pipline的好处是所有东西代码化。
maven project
使用pipline  裁剪了没有maven的东西
使用gradle
流水线 自动化
流水线的好处就是之前所做的手动操作都可以放到自动化里面去执行。
代码一提交就能够跑。

gradle vs maven 更多是代码 是对开发者友好 小任务 直接在代码里写
代码风格的控制 checkStyle
代码风格在每个团队中都是控制不住的，诶，这样写是漂流的，刚开始写代码的同学能将需求完成就是阿弥陀佛了，所以在团队级别拉一根线，要有baseLine，不管你写得好坏，但是至少每个写得看起来样子是那样的。
增加checkStyle目的在于编译之前先把代码写得干净一点。然后再做一个编译测试。  
high proity /normal/ ...

测试覆盖率并不代表测试有效性，有时是和业务无关的。

pipline 可以从git中  pipline as code 
这里使用代码贴是有缺陷的，因为没有代码管理。不过jenkins上已经对pipline增强，支持从git svn仓库拉取。  
pipline as code. 流水线是代码的一部分。

jenkins中会将代码验证的。  


注意有个很蠢得是 虽然挂掉了，但是docker也把服务起来了，下次再构建的时候需要先干掉。
用这个`sh "sudo docker stop hello_app|| true && sudo docker rm -v hello_app|| true`

开发项目，流水线板，放在团队边上，每个人一提交，就看看代码是不是正确的，而不是等待测试的是说你这里代码有错，这样有可能就一天半天就过去了，我们希望你做的是不是正确这件事上，你本身就应该保证，除了你自己保证，我还需要在团队级别保证你是正确的，能不能在5分钟之内出来，我团队是接受你的提交的。这样的一条流水线，作为个人你提交的代码我们认为是ok的，就是团队级别是能接受的。  

大坑注意：  
network -> ci 
我们通过docker-compose.yml启动起来的，其实没有指定docker network的name。默认没指定名称的话，将会默认父级文件夹的名字。

sudo docker network ls
defalut `parentFilePath`


持续集成的目的 将代码达到可部署的版本。  
持续集成到最后出来一个可部署的包，但是还未端到端，还没到线上。
接下来是持续交付
很多时候说 不停机交付部署 
那不停机交付都有哪些方法?  
蓝绿部署测试 vs 灰度发布(分级发布 逐步扩散 高级)
区别是什么?  

灰度一套环境，按照机器或者用户逐渐的放纵流量或者关停服务  
蓝绿部署 则是两套相同的环境，一套backup，一套线上的。比如先更新蓝环境(一般蓝环境是冷的)，更新完蓝环境再把绿环境切下来编程蓝环境。
灰度更加高级，其实是分级发布，带有逐步测试，逐步扩散的区别。
如果只是简单的升级，直接蓝绿部署即可。但另外有A-B测试，不要和蓝绿部署混淆，不是一个概念。

hello_app 需上线 更新版本
要做两个，是需要解决导流的问题，这里使用nginx，反向代理loadBalance 流量切换用。
A  B  
微服务  数据库实例还是schema，但在微服务中更多的强调是schema，而非是实例。有时候会有数据中心，就是说你这个数据库是由数据中心运维保障的，我们要求其多实例，但是资源上或license上是不能我们控制的。所以最低的要求是schema，pr或rquest上过来，需要改代码，该我们的软件，软件是由logic(算法)+数据结构。数据结构很大程度上是体现在schema上的，即表结构上的。 所以为什么单体架构很痛苦的一点就是：我们的数据库是在一个表结构里面，很多个应用功能都在用它，然后每个请求过来，多多少少改一点数据库，然后就会向后影响很大，搞得大家都不敢动这个数据库表了.只能够不断往数据库中塞字段塞字段，塞到最后几十个或者上百个column的表。有时候谁都不知道这是干什么用的，大家拍拍脑袋说好是没用的，但是大家没有人敢去删，因为你都不知道哪里在用。所以微服务强调单数据库或独立数据库。

有无状态，啥时候掐掉流量很重要。  
今天这样讲，是由于我们有开发设计上的假设：
1. 数据库变更要小 
2. 向下兼容
3. rollback 

虚机中
cd ~/cd
more docker-compose.yml
三个服务，蓝、绿(同样的imgage) hello_app、nginx
有个小知识点:nginx要等两个服务都起来了，才启动，根据
```sh
depends_on:
        cd_hello_blue:
              condition: service_healthy
        cd_hello_green:
              condition: service_healthy
      ports:
      ..
```
sudo docker-compose up

切换脚本 blue_green_upgrade.sh
这次很暴力的切换，下次将换一种方式(考虑到很多台服务器的时候)。这次是假设只要两个实例。 hello 端到端 

很多时候是不需要那么过度设计，都要上高大上的东西。
前天吴老师在devops大会上听到阿里部署一个应用要一个万节点，什么概念，比如最差一台服务器的并发量为1000，也就说同时并发在千万级别，也就说用户量至少在上亿级别。很多时候，企业级应用都是部署在企业里，可能就1000个人，也就两个节点就够了。






- tdd讲解

jenkins中的配置 等多少秒 构建
H/1... report

蓝绿部署的  
ci  cd 

do not break? hurt? 不要影响别的代码或服务

微服务的落地研讨
1. 落地有何困难？
2. 技术栈的选型 / 交付的选型(今天的关注) /运维的技术栈
3. 微交付 微测试 集成测试弱化  unit测试要强化
4. 遗留系统..

现状考虑
精益坊要求
反馈

建模能力需提升






























roll haking 黑客增长





