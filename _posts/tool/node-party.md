![Alt text](./1474780340838.png)

13:30-14:30 ng2
14:50-15:50 flux
16:10-17:00 杨疯疯

中秋ng2 final  

江志成 （雪狼） 
同步？ git...


## 大纲
演示 入门
亮点 why ng2？
案例 early bird 分享
周边 新技术，新领域
杂谈 开发组与中文社区简介



## 演示 （代码讲解，log看github源码）
idea 
ts写的
angular cli 辅助开发 初始化 
引入bootstrap  chekout. reversion 右击 git log中的

```
npm run g c user-list 生成四个文件 .
```
ts  @component ...类似Java注解
html
scss 全部局部化的样式 只会影响到自己的组件  父 ，子，兄弟都不会影响

*ngFor = "let item of items"  click    idea支持赞 跳转双向...
路由绑定

userModel  类似 c#的感觉 ，很简单 作者一样..都是微软的 ts

**step**
.... 
填充最简单的master-...
加modole
提取服务
改成observable    | async  管道 v2   v1 过滤器
提取内容组件

```
npm run g s user-api
```

all in github 。。。。。 but in time



## 亮点  
### 工程化
大项目 团队协作   
复杂度 解耦  依赖注入 

### 现代化
v1 09年 ng1目前看还是比较丑的... pwa上说的.... gde社区在研究pwa

### 集成化
v1 各种风格  选择搭配

v2 内置   cli   pwa  universial  ... 官方组织  其他组织等  ngplus  material 等

### 模块化
不仅是语言语法的特性 import
ng2本身的特性 

数据流 抽象成 接口   连接 redux nodejs ... 使用di进行优化 等等 

引入 
AoT ：瘦身，提速   60k <
服务端渲染: 提速 seo
ts ：提高质量    优点多多 js的时代 webstorm可能把库中的内容都改了，但是ts是静态的，与IDE结合相当好
webworker：提速
更好的路由、表单  ng1的路由相当烂(当然本身也很强)  ng2的路由是最好的.......
模型数据表单  （实现动态表单）  各种校验 啥的 都可以完成。

### 简化了什么
内置指令 -改为绑定  
测试中不再需要魔术
更简单的依赖注入  使用类型而不再使用名字
cli  :  大幅度简化开发
官方的style guide 建议大家看下


### 案例分析
ifish

前ng2  冯杨。。快速上手
后台纯Java
nvgd3？ 图标 2天
微站论坛  开源社区 文章打包到js中 打开很慢  但一路不需要流量了


## 新技术 新领域
AoT编译
服务端渲染
nativeScript 整合  直接打包原生应用  开发环境支持好   运行安装脚本即可 所需依赖都包含  但beta版  ，较为（reactiveNative好些）
RXJs整合
PWA支持  内部也做了写壳程序？

ng中文社区
数独
一周大概新增上千人

官方开发组
2.5年
55+18+8 各个版本
5651 提交数量 0920

## 开发资源
chrome插件 ： augury
idea ws idea vsc（轻量） 等
start  angularclass ng2-webpack-starter
cli: ng cli   自己可以定制，官方的目前蛋疼，追求完美
@types: npm i --save-dev @type/*
组件库 ： primeng、 Material 2(目前早起开发阶段，消费类项目)
资源大全（english版）: wx.angular.cn 
大家可以共同维护

学习资源： 官网（墙外）angular.io
中文官网 ： angular.cn
中文社区公众号： angular中文社区
知乎： “angular 2” 话题
github 
官方源码
angular-bbs/*: ng中文社区微站论坛
greengerong/rebirth 破浪的ng2版个人博客源码 涉及较多
ng-book2 :中文版正在翻译  图灵将出版




川 ：
0. 经历
1. 造轮子 发明轮子

