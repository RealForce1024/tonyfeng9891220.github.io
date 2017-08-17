## 理论 重要！


很有启发的一片文章 [不可错过的「持续集成」进阶指南](https://zhuanlan.zhihu.com/p/23264046)
[谈谈持续集成，持续交付，持续部署之间的区别](http://blog.flow.ci/cicd_difference/)
[应用Docker进行持续交付：用技术改变交付路程](https://yq.aliyun.com/articles/54783)
[架構師觀點: 你需要什麼樣的 CI / CD ?](http://columns.chicken-house.net/2017/08/05/what-cicd-do-you-need/)
看了该篇文章后，选择gitlab+docker/k8s是目前我决定使用构建CI/CD的平台。

使用gitlab和k8s集成，对于目前的我来说没有必要。成本太高。大规模团队才适用。  

- [Demo: CI/CD with GitLab](https://www.youtube.com/watch?v=1iXFbchozdY)
- [Containers, Schedulers and GitLab CI](https://www.youtube.com/watch?v=d-9awBxEbvQ)

这位印度工程师的案例非常好！但他的发音...
- [CI/CD using Docker and Gitlab](https://www.youtube.com/watch?v=TPn5pL2yTmo)
发现需要自行开启gitlab registry服务

[GitLab容器注册服务已集成于Docker容器](http://www.infoq.com/cn/news/2016/05/gitlab-docker-registry?utm_campaign=infoq_content&utm_source=infoq&utm_medium=feed&utm_term=global)

[GitLab University: Docker](https://www.youtube.com/watch?v=ugOrCcbdHko)
[Demo: Idea to Production](https://www.youtube.com/watch?v=pY4IbEXxxGY)
[GitLab University: Basics of Git & GitLab](https://www.youtube.com/watch?v=03wb9FvO4Ak)

[devops核心 持续交付](http://www.jianshu.com/p/5b433bc5ddf6)
[DevOps系列｜为什么说Docker吊打了传统持续交付！](https://www.easyops.cn/news/362)
[持续集成系统的演进之路](http://jolestar.com/ci-teamcity-vs-jenkins/)
[https://fir.im/ 团队的ci产品 商业](http://docs.flow.ci/zh/ios_quick_start.html)  
[flow.ci](https://flow.ci/?d=1502815891634)
[论坛讨论](https://segmentfault.com/q/1010000003784336)
[](http://cosven.me/blogs/9)
[深入浅出Docker（四）：Docker的集成测试部署之道](http://www.bkjia.com/Linux/936347.html)
[小团队玩不转的测试](http://blog.kazaff.me/2016/08/18/%E5%B0%8F%E5%9B%A2%E9%98%9F%E7%8E%A9%E4%B8%8D%E8%BD%AC%E7%9A%84%E6%B5%8B%E8%AF%95/)
[如何将一个核心银行系统装入几个容器？来自DOES EU 17大会的观点](http://www.infoq.com/cn/news/2017/08/containers-core-banking) [youtube 大会视频](https://www.youtube.com/watch?v=6FFFrqjybnE)

[开源组件搭配Docker、MESOS、MARATHON，不要太配哦 | 又拍云企业容器](http://weibo.com/ttarticle/p/show?id=2309404050097963984969&sudaref=www.google.com&retcode=6102)

[微服务下架构下的开发与部署](http://weibo.com/ttarticle/p/show?id=2309351002704130532656034300#related)

[docker的应用场景](https://www.zhihu.com/question/22969309)
[游戏运维的最佳实践：搜狐畅游自动化运维之旅](http://dockone.io/article/2547)

[JAVA后端工作流推荐五--Gitlab Runner对Gradle构建的SpringBoot项目进行持续集成--理论篇](https://blog.dxscx.com/2017/01/09/gitlab-runner/)
[JAVA后端工作流推荐六--Gitlab Runner在Gradle构建的SpringBoot项目中的应用--实战篇](https://blog.dxscx.com/2017/01/09/gitlab-runner-gradle/)
阿里云容器服务[基于Jenkins和Docker搭建持续交付流水线](https://www.alibabacloud.com/zh/getting-started/projects/setup-jenkins-based-continuous-delivery-pipeline-with-docker)

## 视频
[Grab 在 gopherChina 上的演讲风格十分有趣，把持续集成和部署做到了极致，正如他们所说，“持续集成，我们是认真的”](https://www.v2ex.com/t/272343)
可以查查去年的视频

## 实战训练
[基于Docker的DevOps实战培训(GitLab+Jenkins)](http://docs.devopshub.cn/udad-devops-docker-hols/index.html)
[DevOpsHub 文档中心](http://docs.devopshub.cn/)
[基于Docker的Gitlab服务器搭建手动搭建 Gitlab 个人仓库，基于Docker，建议使用工具 docker-compose](http://qii404.me/2017/04/17/docker-gitlab.html)
[Jmeter 一次模拟简单秒杀场景的实践 Docker + Nodejs + Kafka + Redis + MySQL](http://www.itnose.net/detail/6712076.html)
[docker问答系列 应用代码是应该挂载宿主目录还是放入镜像内？](http://chuansong.me/n/1459002051928)

```sh
应用代码是应该挂载宿主目录还是放入镜像内？ 

两种方法都可以。

如果代码变动非常频繁，比如开发阶段，代码几乎每几分钟就需要变动调试，这种情况可以使用 --volume 挂载宿主目录的办法。这样不用每次构建新镜像，直接再次运行就可以加载最新代码，甚至有些工具可以观察文件变化从而动态加载，这样可以提高开发效率。

如果代码没有那么频繁变动，比如发布阶段，这种情况，应该将构建好的应用放入镜像。一般来说是使用 CI/CD 工具，如 Jenkins, Drone.io, Gitlab CI 等，进行构建、测试、制作镜像、发布镜像、以及分步发布上线。

对于配置文件也是同样的道理，如果是频繁变更的配置，可以挂载宿主，或者动态配置文件可以使用卷。但是对于并非频繁变更的配置文件，应该将其纳入版本控制中，走 CI/CD 流程进行部署。

需要注意的一点是，绑定宿主目录虽然方便，但是不利于集群部署，因为集群部署前还需要确保集群各个节点同步存在所挂载的目录及其内容。因此集群部署更倾向于将应用打入镜像，方便部署。

```

[使用Docker Image跑Gitlab](https://www.bbsmax.com/A/gVdnm37N5W/)

[搭建本地私有Docker仓库](http://blog.kazaff.me/2016/06/16/%E6%90%AD%E5%BB%BA%E6%9C%AC%E5%9C%B0%E7%A7%81%E6%9C%89docker%E4%BB%93%E5%BA%93/)
[尝试持续集成--第一版](http://blog.kazaff.me/2016/06/16/%E5%B0%9D%E8%AF%95%E6%8C%81%E7%BB%AD%E9%9B%86%E6%88%90--%E7%AC%AC%E4%B8%80%E7%89%88/)
[使用gitlab runner做持续集成测试](http://www.51testing.com/html/20/n-3719320.html)
[基于docker+gitlabCI搭建私有集成环境](http://blog.kazaff.me/2016/06/15/%E5%9F%BA%E4%BA%8Edocker+gitlabCI%E6%90%AD%E5%BB%BA%E7%A7%81%E6%9C%89%E6%8C%81%E7%BB%AD%E9%9B%86%E6%88%90%E7%8E%AF%E5%A2%83/)
## 资料
[持续集成](http://www.cnblogs.com/99fu/p/6042744.html)
[我所了解的几种持续集成方案](http://www.jianshu.com/p/e3c5fdc84416)
[麻袋理财基于Docker的容器实践：互联网金融征信项目的微服务化之旅](https://yq.aliyun.com/articles/59951)
## gitlab-ci和jenkins的比较
[Gitlab-ci 自動測試](http://phorum.study-area.org/index.php?topic=71620.0)
[gitlab vs jenkins](https://gxnotes.com/article/92531.html)

[持续集成环境选择：Jenkins VS gitlab-ci](http://blog.csdn.net/xinluke/article/details/53982150)

```
Jenkins

Jenkins作为老牌的持续集成框架，在这么多年的发展中，积累很多优秀的plugin工具，对进行持续集成工作带来很大的便利。

gitlab-ci

gitlab-ci作为gitlab提供的一个持续集成的套件，完美和gitlab进行集成，gitlab-ci已经集成进gitlab服务器中，在使用的时候只需要安装配置gitlab-runner即可。 
gitlab-runner基本上提供了一个可以进行编译的环境，负责从gitlab中拉取代码，根据工程中配置的gitlab-ci.yml，执行相应的命令进行编译。

jenkins VS gitlab-runner

gitlab-runner配置简单，很容易与gitlab集成。当新建一个项目的时候，不需要配置webhook回调地址，也不需要同时在jenkins新建这个项目的编译配置，只需在工程中配置gitlab-ci.yml文件，就可以让这个工程可以进行编译。
gitlab-runner没有web页面，但编译的过程直接就在gitlab中可以看到，不需要像jenkins进入web控制台查看编译过程。
gitlab-runner仅仅是提供了一个编译的环境而已，全部的编译都通过shell脚本命令进行。当然，jenkins也可以是全部的编译都通过shell脚本命令进行。
jenkins的好处就是编译服务和代码仓库分离，而且编译配置文件不需要在工程中配置，如果团队有开发、测试、配置管理员、运维、实施等完整的人员配置，那就采用jenkins，这样职责分明。不仅仅如此，jenkins依靠它丰富的插件，可以配置很多gitlab-ci不存在的功能，比如说看编译状况统计等。如果团队是互联网类型，讲究的是敏捷开发，那么开发=devOps，肯定是采用最便捷的开发方式，推荐gitlab-ci。
如果有些敏感的配置文件不方便存放在工程中（例如nexus上传jar的账户和密码或者是其他配置的账户密码）,都可以在服务器中配置即可。
gitlab-ci对于编译需要的环境，比如jdk，maven都需要自行配置。在jenkins中，对于编译需要的环境，比如jdk，maven都可以在Web控制台安装即可。当然，jenkins也是可以自行配置的（有时候通过控制台配置下载不下来）。
-
总结

在使用过两者后，个人觉得gitlab-ci更简单易用，如果有gitlab-ci达不到的要求，可以考虑使用jenkins。
```

## docker+gitlab-ci搭建
[使用Gitlab和Gitlab CI做持续集成（理论篇）](https://my.oschina.net/donhui/blog/717930)

[Gitlab-CI + Docker 构建完全容器化持续集成方案](https://my.oschina.net/u/2400083/blog/818222)

[Docker搭建Gitlab CI 全过程详解](https://my.oschina.net/u/1396253/blog/181498)

[gitlab-ci and docker](https://ronmi.github.io/post/docker/gitlab-ci-docker/)
[在 Docker 里构造 Meteor 持续集成环境](http://blog.csdn.net/pinxue/article/details/45500277)

[](https://www.bbsmax.com/A/gVdnm37N5W/)
[CI持续集成系统环境--Gitlab+Gerrit+Jenkins完整对接](http://www.cnblogs.com/kevingrace/p/5651447.html)


## 其他
[Gitlab：为什么我们坚持用云？](http://www.yunweipai.com/archives/18870.html)

  [为 Gitlab 和 Jenkins 添加 InfluxDB 支持](http://blog.fleeto.us/content/wei-gitlab-he-jenkins-tian-jia-influxdb-zhi-chi)

[k8s内容较多 伪架构师](http://blog.fleeto.us/)

[如何基于接口文档生成模拟数据](http://blog.kazaff.me/2016/09/21/%E5%A6%82%E4%BD%95%E5%9F%BA%E4%BA%8E%E6%8E%A5%E5%8F%A3%E6%96%87%E6%A1%A3%E7%94%9F%E6%88%90%E6%A8%A1%E6%8B%9F%E6%95%B0%E6%8D%AE/)

[3 天烧脑式基于Docker的CI/CD实战训练营 | 北京站 宣传贴 不过目录可以学习下](http://dockone.io/article/2526)

[gitlab docker 运维升级](http://www.jianshu.com/p/f836c3b867f8)



