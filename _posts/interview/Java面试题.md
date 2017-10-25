[Java 面试宝典](http://wiki.jikexueyuan.com/project/java-interview-bible/)
[Java工程师面试题整理 社招版](https://zhuanlan.zhihu.com/p/21551758)
[2016Java面试题整理 github java interview](https://github.com/it-interview/easy-java)
[Java面试题 全方位 包含sql 框架等 场景面试题](https://www.jfox.info/)
[最强求职攻略:java程序员如何通过阿里、百度的招聘面试](https://www.jfox.info/%E6%9C%80%E5%BC%BA%E6%B1%82%E8%81%8C%E6%94%BB%E7%95%A5java%E7%A8%8B%E5%BA%8F%E5%91%98%E5%A6%82%E4%BD%95%E9%80%9A%E8%BF%87%E9%98%BF%E9%87%8C%E7%99%BE%E5%BA%A6%E7%9A%84%E6%8B%9B%E8%81%98%E9%9D%A2.html)
[牛客网 常考java面试题](https://www.nowcoder.com/ta/review-java)
[Java面试题整理](https://dongchuan.gitbooks.io/java-interview-question/content/sql/one_to_multi.html)


--整理

继承和形参

```java
public class PrintWhoAmI {
    public static void main(String[] args) {
        A a = new B();
        test(a); // please print result
    }

    public static void test(A a) {
        System.out.println("test A");
        a.whoAmI();
    }

    public static void test(B b) {
        System.out.println("test B");
        b.whoAmI();
    }
}

class A {
    public void whoAmI() {
        System.out.println("I'm A");
    }
}

class B extends A {
    @Override
    public void whoAmI() {
        System.out.println("I'm B");
    }
}
```

解析: 
该题主要考察继承和形参。继承中"父类引用指向子类实现"，子类调用负责具体实现。而形参则由父类引用决定。所以`A a = new B()`中形参有`a`决定，所以编译器会调用`test(A a)`，而方法执行中`a.whoAmI()`则由具体的`new B()`决定。 
因此将会输出

```java
test A
I'm B
```

---


当执行  new Son(); 时，会输出什么？

```java
public class Father {  
    private String name="FATHER";  
    public Father(){  
        whoAmI();  
        tellName(name);  
    }  
    public void whoAmI(){  
        System.out.println("Father says, I am " + name);  
    }  
    public void tellName(String name){  
        System.out.println("Father's name is " + name);  
    }  
}  
  
public class Son extends Father {  
    private String name="SON";  
    public Son(){  
        whoAmI();  
        tellName(name);  
    }  
    public void whoAmI(){  
        System.out.println("Son says, I am " + name);  
    }  
    public void tellName(String name){  
        System.out.println("Son's name is " + name);  
    }  
}
```

 
最终结果如下：

```java
Son says, I am null
Son's name is FATHER
Son says, I am SON
Son's name is SON
```

解析: 
1.创建Son的时候先创建Father，而Father中的whoAmI是已经被Son覆盖了，因此这里打印的name这个field是Son中的field，而此时还没有构造Son，因此Son中的name的值是null(这里即使是写了String name = "SON"也是没有用的，因为父类没有构造结束之前，这里是不会被执行的)。
2.Father在执行tellName的时候，传递的参数name是Father自身的name这个field值，这个值是已经被赋值为"FATHER"的，因此会打印出“Son's name is FATHER”
3.Father构造完毕后开始构造Son，这里的打印结果就可以按照常规方式来解释了。 

---

java中的equals,hashCode的区别和联系?

```java
1. public boolean equals(Object) 和 int hashCode() 这两个方法都定义在Object顶级类当中，Java中的每个对象都同时含有这两个方法。equals方法是比较对象是否相等，而hashCode则是返回对象的散列值，在Object类中的默认实现是“将该对象的内部地址转换成一个整数返回”。
2. 在对象比较的情况下，一般需要同时重写equals和hashCode方法，同时重写hashCode的原因在于散列值的计算是由equals逻辑中定义的字段值而来。否则equals方法也一般不会起效。
3. 在构造散列表的时候，由于哈希碰撞问题以及元素重复问题，需要注意hashCode和equals方法的重写问题。判断对象equals之前是需要先判断hashCode的。
4. 在equals逻辑中不要直接调用属性字段值，在orm框架中会有延迟加载问题，需要调用getter方法
5. a、如果两个对象equals，Java运行时环境会认为他们的hashcode一定相等。 
   b、如果两个对象不equals，他们的hashcode有可能相等。 
   c、如果两个对象hashcode相等，他们不一定equals。 
   d、如果两个对象hashcode不相等，他们一定不equals。 
   总之一句话 如果对象equals相等，那么需要重写hashCode确保相等，反之则不一定。
```

[ Java 中的 ==, equals 与 hashCode 的区别与联系](http://blog.csdn.net/justloveyou_/article/details/52464440)
[Java hashCode() 和 equals()的若干问题解答](http://www.cnblogs.com/skywang12345/p/3324958.html)
[由==到equals再到hashCode方法](http://tracylihui.github.io/2015/09/29/java/%E7%94%B1==%E5%88%B0equals%E5%86%8D%E5%88%B0hashCode%E6%96%B9%E6%B3%95/)

[如何“记住” equals 和 == 的区别？](https://www.zhihu.com/question/26872848/answer/34364603)
[从一道面试题彻底搞懂hashCode与equals的作用与区别及应当注意的细节](http://blog.csdn.net/lijiecao0226/article/details/24609559)

```
作者：leeon
链接：https://www.zhihu.com/question/26872848/answer/34364603
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

又是这个问题哈， @金波 的回答是我看到最喜欢的，一句话很清晰的比较了两个方法。Java 语言里的 equals方法其实是交给开发者去覆写的，让开发者自己去定义满足什么条件的两个Object是equal的。所以我们不能单纯的说equals到底比较的是什么。你想知道一个类的equals方法是什么意思就是要去看定义。Java中默认的 equals方法实现如下：public boolean equals(Object obj) {
    return (this == obj);
}
而String类则覆写了这个方法,直观的讲就是比较字符是不是都相同。public boolean equals(Object anObject) {
    if (this == anObject) {
        return true;
    }
    if (anObject instanceof String) {
        String anotherString = (String)anObject;
        int n = count;
        if (n == anotherString.count) {
            char v1[] = value;
            char v2[] = anotherString.value;
            int i = offset;
            int j = anotherString.offset;
            while (n-- != 0) {
                if (v1[i++] != v2[j++])
                    return false;
            }
            return true;
        }
    }
    return false;
}
equals如何比较并不重要，但是不理解equals存在的目的就容易踩坑。比如这里面的一个例子http://atleeon.com/code/2013/11/29/java-equals-hashcode/
```

[HASHCODE和HASHMAP、HASHTABLE](http://www.debugrun.com/a/WXl4bmk.html)

```java
HashCode的作用

首先，想要明白hashCode的作用，你必须要先知道Java中的集合。
　　总的来说，Java中的集合（Collection）有两类，一类是List，再有一类是Set。你知道它们的区别吗？前者集合内的元素是有序的，元素可以重复；后者元素无序，但元素不可重复。那么这里就有一个比较严重的问题了：要想保证元素不重复，可两个元素是否重复应该依据什么来判断呢？这就是Object.equals方法了。但是，如果每增加一个元素就检查一次，那么当元素很多时，后添加到集合中的元素比较的次数就非常多了。也就是说，如果集合中现在已经有1000个元素，那么第1001个元素加入集合时，它就要调用1000次equals方法。这显然会大大降低效率。
    于是，Java采用了哈希表的原理。哈希（Hash）实际上是个人名，由于他提出一哈希算法的概念，所以就以他的名字命名了。哈希算法也称为散列算法，是将数据依特定算法直接指定到一个地址上。如果详细讲解哈希算法，那需要更多的文章篇幅，我在这里就不介绍了。初学者可以这样理解，hashCode方法实际上返回的就是对象存储的物理地址（PS：这是一种算法，数据结构里面有提到。在某一个地址上（对应一个哈希值，该值并不特指内存地址），存储的是一个链表。在put一个新值时，根据该新值计算出哈希值，找到相应的位置，发现该位置已经蹲了一个，则新值就链接到旧值的下面，由旧值指向（next）它（也可能是倒过来指。。。）。可以参考HashMap）。
    这样一来，当集合要添加新的元素时，先调用这个元素的hashCode方法，就一下子能定位到它应该放置的物理位置上。如果这个位置上没有元素，它就可以直接存储在这个位置上，不用再进行任何比较了；如果这个位置上已经有元素了，就调用它的equals方法与新元素进行比较，相同的话就不存了，不相同就散列其它的地址。所以这里存在一个冲突解决的问题。这样一来实际调用equals方法的次数就大大降低了，几乎只需要一两次。
    所以，Java对于eqauls方法和hashCode方法是这样规定的：
1、如果两个对象相同，那么它们的hashCode值一定要相同；
2、如果两个对象的hashCode相同，它们并不一定相同
    上面说的对象相同指的是用eqauls方法比较。
    你当然可以不按要求去做了，但你会发现，相同的对象可以出现在Set集合中。同时，增加新元素的效率会大大下降。
```

