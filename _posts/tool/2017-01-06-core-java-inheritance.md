---
layout: post
title: core-java oop 继承
category: 基础知识
tags: java
keywords: java, 复习
---

继承(inheritance)，面向对象的核心技术之一。  
利用继承，可以基于一个已存在的类构造新的类。继承已存在的类就是复用(继承)这些类的方法和域。在此基础上，还可以添加新的方法和域，以满足新的需求。
```java
public class Manager extends Employee {
    private double bonus;

    public Manager() {
        bonus = 0;
    }

    @Override
    public double getSalary() {
        double salary = super.getSalary();
        return salary + bonus;
    }

    public double getBonus() {
        return bonus;
    }

    public void setBonus(double bonus) {
        this.bonus = bonus;
    }
}
```
子类是无法访问父类私有域的。只能通过访问器。  
java不支持多继承，但是可以通过接口的形式实现多继承的功能。


反射是指在程序运行期间发现更多的类及其属性的能力。功能十分强大。开发软件工具或平台的人(架构设计上层武功..)更应该关注，而编写应用程序的人员用到的几乎很少。  

## 多态(polymorphism)  
一个对象变量可以指示多种实际类型的现象叫做多态。在运行时能够自动地选择调用哪个方法的现象被称为`动态绑定`(dynamic binding)

```java
Employee[] stuff = new Employee[3];
stuff[0] = new Manager();
stuff[1] = new Employee();
stuff[2] = new Employee();

for(Employee e: stuff){//父类引用指向子类实例
   System.out.println("name:" + e.getName()); 
}
```
是否应该设计为继承关系，可以通过`is a`规则。`它表明子类的每个对象也是超类的对象。` 因为是继承嘛，父类有的基本都有了，只不过有些不好访问罢了。另外还可以通过`is a`的置换法则，即父类对象出现的任何地方都可以使用子类对象置换。  
可以将子类对象赋值给父类对象。  

```java
//执行错误，不能转换。
public static void main(String[] args) {
    Manager m = (Manager) new Employee();
    m.getBonus();
}
```
执行了父类的静态代码块，id为随机数输出。
```java
public static void main(String[] args) {
    Employee e = new Manager();
    System.out.println(e.getId());
}
```
**`对象变量是多态的。`**  
一个Employee对象既可以引用一个Employee对象，也可以引用一个Employee类的任何子对象。（置换法则的好处，父类出现的地方，子类都可以替换）
但是注意：
```java
Manager boss = new Manager();
Employee[] stuff = new Employee[3];
stuff[0] = boss;

boss.setBonus(1000); //ok 
stuff[0].setBonus(1000); //error
```
boss和stuf[0]引用同一个对象。但编译器将stuff[0]看成Employee对象。数组的定义是一种数据类型的集合。而这里就是Employee对象的集合。而setBonus并不是Employee类的方法。
不能将一个超类的引用赋值给子类变量。例如：  
Manager m = stuff[i] //error  
原因很清楚: 不是所有的雇员都是经理。如果赋值成功，那么m可能引用了一个不是经理的Employee对象。当在调用m.setBonus(...)时就有可能发生运行时错误。  

另外还有一种非常隐蔽的错误
```java
Manager[] managers = new Manager[10];
Employee[] stuff = managers;

stuff[0] = new Employee();//编译器居然是可以通过的
```

stuff[0] 与 managers[0]引用的是同一个对象，那么就有可能把一个普通雇员擅自归纳经理行列当中了。这是一种非常忌讳发生的情形。 当调用managers[0].setBouns(1000)的时候，将会导致调用一个不存在的实例域，进而扰乱相邻存储空间的内容。  

为了确保不发生类似的错误，所有数组都必须牢记创建它们的元素类型，并负责监督仅将类型兼容的引用存储到数组中。例如：new Manager[10] 视图存储一个Employee类型的引用就会引发ArrayStoreException异常。  

弄清楚方法调用的过程，清楚编译器调用方法的执行流程。  
private \static \ final \ 构造器 编译器调用都很明确，因此为静态绑定。(static binding)。与此对应的是，调用的方法依赖于隐式参数的实际类型，并且在运行时动态绑定。  


