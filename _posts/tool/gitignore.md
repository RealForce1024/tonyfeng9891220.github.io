---
layout: post
title: git全局配置
category: 工具
tags: git
keywords: git,工具
---
git全局配置

```
[user] 
name = fqc 
email = 337940626@qq.com

[core] 
editor = vim #ubuntu上默认的nano 
[alias] 
ci = commit -a -v 
st = status -s 
br = branch 
ctm = commit -m 
ckt = checkout
throw = reset --hard HEAD 
throwh = reset --hard HEAD^

[color] 
ui = true

[push] 
default = current
```