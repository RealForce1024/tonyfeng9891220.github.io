---
layout: post
title: 并发编程实战系列006-synchronized代码块与颗粒度问题 
category: 并发编程
tags: java
keywords: java, synchronized代码块
---  

## synchronized颗粒度问题

**减小锁的颗粒度**
synchronized关键字声明的方法在某些时候是有弊端的。比如A线程调用同步的方法需要执行很长时间，B线程则必须等待较长的时间才能执行。**这种情况可以使用synchronized代码块去优化代码执行时间，也就是减小锁的颗粒度。**

我们之前使用过对象级别，以及类级别的synchronized关键字。但是性能上确实不尽如意。java提供了synchronized代码块，以通过缩小颗粒度的方式提高并发的执行效率。  

## 案例 
synchronized可以使用任意的对象进行加锁，用法比较灵活。  
synchronized代码块使用方式如下示例
```java
public class ObjectLock {
    public void m1() {
        synchronized (this) {//当前对象锁
            System.out.println(Thread.currentThread().getName() + "=>m1()");
            try {
                Thread.sleep(2000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

    public void m2() {
        synchronized (ObjectLock.class) {//类级别锁
            System.out.println(Thread.currentThread().getName() + "=>m2()");
            try {
                Thread.sleep(2000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

    private Object obj = new Object();//任意对象锁

    public void m3() {
        synchronized (obj) {
            System.out.println(Thread.currentThread().getName() + "=>m3()");
            try {
                Thread.sleep(2000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

    public static void main(String[] args) {
        final ObjectLock objectLock = new ObjectLock();
        Thread t1 = new Thread(() -> objectLock.m1(), "t11");
        Thread t2 = new Thread(() -> objectLock.m2(), "t22");
        Thread t3 = new Thread(() -> objectLock.m3(), "t33");

        t1.start();
        t2.start();
        t3.start();
    }
}
```

