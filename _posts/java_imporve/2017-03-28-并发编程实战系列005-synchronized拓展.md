---
layout: post
title: 并发编程实战系列005-synchronized锁重入和异常释放锁
category: 并发编程
tags: java
keywords: java, synchronized, 异常释放锁, 锁重入 
---

## 并发量真的那么大吗
在一般的并发量上来说，jdk提供的synchronized足够，并且jdk1.8又对synchronized进行了性能优化。而ReentrantLock,信号量等等jdk底层提供的其他锁的处理在一般场景下表现差不多，也就是一般的synchronized足以应对。

## synchronized锁重入
关键字synchronized拥有锁重入的功能，也就是在使用synchronized时，当一个线程得到了对象的锁之后，再次请求此对象时是可以再次得到该对象的锁。
### 案例1
```java
public class SyncDouble1 {
    public synchronized void m1() {
        System.out.println("m1");
        m2();
    }

    public synchronized void m2() {
        System.out.println("m2");
        m3();
    }

    public synchronized void m3() {
        System.out.println("m3");
    }

    public static void main(String[] args) {
        SyncDouble1 syncDouble1 = new SyncDouble1();
        Thread t1 = new Thread(() -> syncDouble1.m1());
        Thread t2 = new Thread(() -> syncDouble1.m1());

        t1.start();
        t2.start();
    }
}
```
### 案例2
```java
public class SyncDubole2 {
    static class Parent {
        public int num = 10;

        public synchronized void subParent() throws InterruptedException {
            num--;
            System.out.println("parent: " + num);
            Thread.sleep(100);
        }
    }

    static class Sub extends Parent {
        public  synchronized void subSub() throws InterruptedException {
            while (num > 0) {
                num--;
                System.out.println("sub: " + num);
                super.subParent();
                Thread.sleep(100);
            }
        }
    }

    public static void main(String[] args) {
        Sub sub = new Sub();
        Thread t1 = new Thread(() -> {
            try {
                sub.subSub();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });
        Thread t2 = new Thread(() -> {
            try {
                sub.subSub();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });

        t1.start();
        t2.start();
    }
}
```
上述两个案例，如果方法没有synchronized的修饰，结果将会混乱。  
案例二则从继承关系上说明使用synchronized一样是线程安全的。  
**`父类和子类都需要同步，否则将会存在线程安全问题。`**

## 出现异常，锁自动释放

对于web应用程序，异常释放锁的情况，如果不及时处理，很可能对你的应用程序业务逻辑产生严重的错误。比如现在执行一个队列任务，很多对象都去在等待第一个对象正确执行完毕再去释放锁，但是对一个对象由于异常的出现，导致业务逻辑没有正常地执行完毕，就释放了锁，那么可想而知后续的对象执行的都是错误的逻辑。所以这点一定要引起注意，在编写代码的时候，一定要考虑周全。  


```java
public class SyncException {
    public int i = 0;

    public synchronized void operate() {
        while (true) {
            try {
                i++;
                Thread.sleep(100);
                System.out.println(Thread.currentThread().getName() + " i=" + i);
                if (i == 10) {
                    Integer.parseInt("a");
                }
                if (i == 20) {
                    Integer.parseInt("b");
                }

            } catch (Exception ex) {
                ex.printStackTrace();
                System.err.println("log info error i=" + i);
                //continue;
                //throw new RuntimeException();
                //throw new InterruptedException;
            }
        }
    }

    public static void main(String[] args) {
        SyncException syncException = new SyncException();
        Thread t1 = new Thread(() -> syncException.operate());
        Thread t2 = new Thread(() -> syncException.operate());

        t1.start();
        t2.start();
    }
}
```


处理方式
1. 记录日志(前后任务间没有关联性，失败不会对其他任务产生业务影响)
2. 终止任务(原子性的任务，具有整体连贯性)

不加synchronized将会很诡异的错误。  
```java
Thread-1 i=2
Thread-0 i=2
Thread-1 i=4
Thread-0 i=4
Thread-1 i=6
Thread-0 i=7
Thread-1 i=8
Thread-0 i=8
Thread-1 i=10
Thread-0 i=10
log info error i=10
log info error i=11
Thread-1 i=12
Thread-0 i=12
Thread-0 i=14
Thread-1 i=14
Thread-0 i=16
Thread-1 i=16
Thread-0 i=18
Thread-1 i=18
Thread-0 i=20
Thread-1 i=21
Thread-0 i=22
Thread-1 i=22
Thread-0 i=24
Thread-1 i=25
```



```
Thread-0 i=1
Thread-0 i=2
Thread-0 i=3
Thread-0 i=4
Thread-0 i=5
Thread-0 i=6
Thread-0 i=7
Thread-0 i=8
Thread-0 i=9
Thread-0 i=10
log info error i=10
Thread-0 i=11
Thread-0 i=12
Thread-0 i=13
Thread-0 i=14
Thread-0 i=15
Thread-0 i=16
Thread-0 i=17
Thread-0 i=18
Thread-0 i=19
Thread-0 i=20
Thread-0 i=21
```

我们可以看到出现异常的时候，第一个线程执行到10抛出异常，线程2立即获得锁继续执行，直到再次遇到异常才终止程序。

```java
Thread-0 i=1
Thread-0 i=2
Thread-0 i=3
Thread-0 i=4
Thread-0 i=5
Thread-0 i=6
Thread-0 i=7
Thread-0 i=8
Thread-0 i=9
Thread-0 i=10
java.lang.NumberFormatException: For input string: "a"
	at java.lang.NumberFormatException.forInputString(NumberFormatException.java:65)
	at java.lang.Integer.parseInt(Integer.java:580)
	at java.lang.Integer.parseInt(Integer.java:615)
	at com.fan.bigdata.thread.thread_exception.SyncException.operate(SyncException.java:16)
	at com.fan.bigdata.thread.thread_exception.SyncException.lambda$main$0(SyncException.java:33)
	at java.lang.Thread.run(Thread.java:745)
log info error i=10
Exception in thread "Thread-0" java.lang.RuntimeException
	at com.fan.bigdata.thread.thread_exception.SyncException.operate(SyncException.java:26)
	at com.fan.bigdata.thread.thread_exception.SyncException.lambda$main$0(SyncException.java:33)
	at java.lang.Thread.run(Thread.java:745)
Thread-1 i=11
Thread-1 i=12
Thread-1 i=13
Thread-1 i=14
Thread-1 i=15
Thread-1 i=16
Thread-1 i=17
Thread-1 i=18
Thread-1 i=19
Thread-1 i=20
java.lang.NumberFormatException: For input string: "b"
	at java.lang.NumberFormatException.forInputString(NumberFormatException.java:65)
	at java.lang.Integer.parseInt(Integer.java:580)
	at java.lang.Integer.parseInt(Integer.java:615)
	at com.fan.bigdata.thread.thread_exception.SyncException.operate(SyncException.java:19)
	at com.fan.bigdata.thread.thread_exception.SyncException.lambda$main$1(SyncException.java:34)
	at java.lang.Thread.run(Thread.java:745)
log info error i=20
Exception in thread "Thread-1" java.lang.RuntimeException
	at com.fan.bigdata.thread.thread_exception.SyncException.operate(SyncException.java:26)
	at com.fan.bigdata.thread.thread_exception.SyncException.lambda$main$1(SyncException.java:34)
	at java.lang.Thread.run(Thread.java:745)
```

## 扩展
plsql存储过程
```sql
begin

    exception
        //记录日志错误信息
end
```
一样的道理，是忽略做错误继续执行，还是终止这一整体任务(由多个子任务组成)。
