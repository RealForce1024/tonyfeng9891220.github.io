---
layout: post
title: 环境变量
category: 工具
tags: 工具
keywords: 算法,排序,Sort,Algorithm
---
环境变量

~/.bash_profile
```
export JDK_HOME=$(/usr/libexec/java_home)
export SCALA_HOME="/usr/local/share/scala"
export SPARK_HOME="/Users/zhaocongcong/apps/spark-1.6.2-bin-hadoop2.6"
export PYTHONPATH=$SPARK_HOME/python:$SPARK_HOME/python/lib/py4j-0.9-src.zip:$PYTHONPATH
export PATH=$SCALA_HOME/bin:$SPARK_HOME/bin:$SPARK_HOME/sbin:$PYTHONPATH:$PATH
alias wcf='ls -l |grep "^-"|wc -l'
```

~/.zshrc
```
export PATH="$HOME/.jenv/bin:$PATH"
eval "$(jenv init -)"

export AWS_ACCESS_KEY_ID=AKIAOTXKEG4QQCKAYTAQ
export AWS_SECRET_ACCESS_KEY=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
#export JDK_HOME="/Library/Java/JavaVirtualMachines/jdk1.8.0_101.jdk/Contents/Home"
export JDK_HOME=$(/usr/libexec/java_home)
export SCALA_HOME="/usr/local/share/scala"
export SPARK_HOME="/Users/zhaocongcong/apps/spark-1.6.2-bin-hadoop2.6"
export PYTHONPATH=$SPARK_HOME/python:$SPARK_HOME/python/lib/py4j-0.9-src.zip:$PYTHONPATH
export PATH=$SCALA_HOME/bin:$SPARK_HOME/bin:$SPARK_HOME/sbin:$PYTHONPATH:$JDK_HOME:$PATH
```
