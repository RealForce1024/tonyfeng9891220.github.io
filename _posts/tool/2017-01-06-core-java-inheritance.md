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
private \static \ final \ 构造器 编译器调用都很明确，因此为静态绑定。(static binding)。与此对应的是，`调用的方法依赖于隐式参数的实际类型`，并且在运行时动态绑定。  
每次调用方法都要进行搜索，时间开销特别大。因此虚拟机预先为每个类创建了方法表(method table)，其中列出了所有的方法签名和实际调用方法。 注意如果调用super.f(param)，编译器将对隐式参数超类的方法表进行搜索。  

在运行的时候，e.getSalary()虚拟机解析过程如下：
1. 首先，虚拟机提取出e的实际类型的方法列表(隐式参数类型对应的方法列表)。既可能是Employee、Manager的方法表，也可能是Employee类的其他子类的方法表。  
2. 然后虚拟机搜索定义为getSalary签名的类。  
3. 最后虚拟机调用方法。  

动态绑定有一个非常重要的特性：`无需对现存的代码进行修改，就可以对程序扩展`。    
假设增加一个新类Executive，并且变量e有可能引用这个类的对象，我们不需要对包含调用e.getSalary()的代码进行重新编译。如果e恰好引用一个Executive类的对象，就会自动地调用Executive.getSalary()方法。  

注意：`在覆盖方法的时候，子类方法不能低于超类方法的可见性`。特别是，**如果父类方法是public，子类方法一定要声明为public。**经常会发生这类错误：在声明子类方法的时候，遗漏了public修饰符。此时`编译器会把它解释为试图降低访问权限`。

## 阻止继承: final类和方法
类、方法被修饰为final，将不能被继承、覆盖。  
`final类所有方法自动标识为final。并不包括域`   
注意域也可以被修饰为final。对于final域来说，构造对象之后就不允许改变他们的值了。不过如果将一个类声明为final，只有其中的方法自动地称为final，而不包括域。  

将方法或类声明为final的主要目的是：确保它们不会在子类中改变语义。  

即时编译器在这方面上有很多优化，比传统编译器强很多。  

## 强制类型转换
将一个值存入变量时，编译器将检查是否允许该操作。将一个子类的引用赋值给父类变量，编译器是允许的。但将一个超类的引用赋值给子类变量，必须进行强制类型转换。这样才能够通过运行时的检查。  
如果试图在继承链上进行向下的类型转换，并且"谎报"有关对象的内容，会发生什么情况？
`Manager boss = (Manager)stuff[1];`// error
java运行时系统将会报告这个错，并产生一个ClassCastException异常。如果没有捕获这个异常，那么程序就会终止。 因此，要养成一个良好的程序设计习惯：在进行类型转换之前，先查看以下是否能够成功地转换。这个过程简单地使用instanceof运算符就可以实现。  
```java 
if(stuff[1] instanceof Manager){
    boss = (Manager)stuff[1];
}else{

}
```
如果这个类型转换不可能成功，编译器就不会进行这个转换。例如：
`Date date = (Date)sutff[1];`//编译错误Date并非Employee的子类。  
综上所述：  
只能在继承层次内进行类型转换。  
在将超类转换成子类之前，应该使用instanceof进行检查。

instanceof
x instanceof C //x为null
并不会产生异常，只是返回false。之所以这么处理是因为null没有引用任何对象。当然也不会引用任何C类型的对象。

实际上使用类型转换调整对象的类型并不是一种好的做法。在列举的示例中，大多数情况下并不需要Employee对象转换成Manager对象，这`两个类的对象都能够正确的调用getSalary方法`，这是因为`实现多态的动态绑定机制能够自动地找到相应的方法`。  

只有在使用Manager中特有的方法时才需要进行类型转换。例如，setBonus方法。如果鉴于某种原因，发现需要通过Employee对象调用setBonus方法，那么就应该检查下超类的设计是否合理。重新设计一下超类，并添加setBonus方法才是正确的选择。
注意：只要没有捕获ClassCastException异常，程序就会终止执行。`在一般情况下，应该尽量少用类型转换和instanceof运算符。`  


## 受保护访问 protected
通过之前所学，我们知道类的域标记为private，方法标记为public。`任何声明为private的内容对其他类都是不可见的。注意：这对于子类来说也是完全适用的，`**`子类也不能访问父类的私有域。`**

`子类只能获得访问受保护域的权利`。其他子类是不允许的，这种限制有助于避免滥用受保护机制。  

在实际应用中，要谨慎使用protected属性。假设需要将设计的类提供给其他程序员使用，而在这个类中设置了保护域，由于其他程序员可以对这个类进行再派生出新的子类，并访问其中的受保护域。在这种情况下，`如果需要对这个类的实现进行修改，就必须通知所有使用这个类的程序员`。这违背了OOP提倡的数据封装原则。  
受保护的方法更具有实际意义。如果需要限制某个方法的使用，就可以将其声明为protected。这表明子类(可能很熟悉祖先类)可以得到信任，可以正确的使用这个方法，而其他类则不行。

protected标记的方法示例最好的就是Object类的clone方法。
`protected native Object clone() throws CloneNotSupportedException;`

**事实上，java中的受保护部分对所有子类及同一包中的所有其他类都可见。**

Java用于控制可见性的4个访问修饰符:
private-仅对本类可见
public-对所有类可见
protected-对本包和所有子类可见
默认-对本包可见，不需要修饰符

## Object类 所有类的基类
如果没有明确指出一个类的超类,Object就会使该类的基类。  
Object类中的方法服务因此是必须熟悉掌握的。  
`Object obj = new Employee();`  
Object类型的变量只能用作为各种值的通用持有者。要想对其中的内容进行具体的操作，还需要清楚原始类型，再进行相应的类型转换:
`Employee employee = (Employee)obj;`  
在Java中只有基本类型(primitive types)不是对象，`数值，字符，布尔`类型的值都不是对象。所有的数组类型，不管是基本类型还是对象类型的数组都是扩展于Object类。

### equals方法
在Object类中，这个方法将判断两个对象是否具有相同的引用。如果两个对象具有相同的引用，它们一定是相等的。默认情况合乎情理。然而这种判断对大多数类在实际应用中并不实用，没有太大意义。例如，判断比较两个PrintStream对象是否相等就完全没有意义。 然而经常需要检测两个对象状态的相等性，如果两个对象的状态相等，就认为两个对象是相等的。

```java
class Employee{

    public boolean equals(Object otherObj){
        if(this==otherObj)//默认就是比较引用，看是否是同一对象引用。
            return true;
        if(otherObj==null) 
            return false;//如果不为空，就不用返回，所以判断的依据在于何时返回
        if(getClass()!=otherObj.getClass())
            return false;
        
        Employee other = (Employee)otherObj;
        //方式一 啰嗦 + 属性值可能为空
        //if(this.getName().equals(other.getName())&&
        //    this.getSalary()==other.getSalary()&&
        //   this.getHireDay().equals(other.getHireDay())
        //)
        //return true;
        
        //方式二 属性值可能为空
        //return name.equals(other.name)
        //    && salary==other.salary
        //    && hireDay.equals(other.hireDay)

        //最佳方案 Objects.equals(a,b)方法 如果都为空true，其中一个为空false,
        //都不为空 a.equals(b) 
        return Objects.equals(name,other.name)
            && salary == other.salary
            && Objects.equals(salary,other.salary)
    }
}
```

`在子类中定义equals方法时，首先调用超类的equals方法`，如果检测失败，对象就不可能相等。如果超类中的域相等，那么就需要比较子类中的实例域。  

```java
class Manager extends Employee{
    ...
    public boolean equals(Object otherObj){
        if(!super.equals(otherObj))
            return false;
        Manager manager = (Manager)otherObj;
        return bonus==otherObj.bonus;
    }
}
```
相等测试与继承
如果隐式和显示的参数不属于同一类，equals方法该如何处理？该问题很有争议。曾经我们没有说继承的时候，如果类不匹配，equals方法就返回false。但是许多程序员喜欢使用instanceof方法来检测
```java
if(!otherObj instanceof Employee) 
    return false;
```
这样做不但没有解决otherObj是子类的情况，并且还会招致一些麻烦。这就是建议不要使用这种处理方式的原因所在。java语言规范equals方法具有以下特性：  
1. 自反性
2. 对称性
3. 传递性
4. 一致性
5. 对于任意非空引用x，  x.equals(null) 返回false
这些规则似乎合乎情理，从而避免了类库实现者在数据结构定位一个元素时还要考虑调用x.equals(y)还是y.equals(x)的问题。
然而，就对称性来说，当参数不属于同一类的时候就需要仔细思考下。
e.equals(m)
e为Employee,m为Manager  
反过来m.equals(e).也需要返回true。`对称性不允许返回false或者抛出异常`。
说到这里，有些程序员或书籍里还是偏向使用instanceof方法，这里总结下：  
如果子类能够拥有自己的相等概念，则对称性需求将强制采用getClass进行检测。  
如果由超类决定相等的概念，那么使用instanceof进行检测，这样就可以在不同子类的对象之间进行相等的比较。  
在雇员和经理的例子中，只要对应的域相等，就认为两个对象相等。可以使用getClass。  
但是，假设使用雇员的ID作为相等的检测标准，并且这个相等的概念适用于所有的子类，就可以使用instanceof进行检测，并应该将Employee.equals方法声明为final。

检测数组类型，可以使用Arrays.equals(arr1,arr2)
```java
int[] arr1 = {1, 2, 3, 4};
int[] arr2 = {1, 2, 3, 4};
if (Arrays.equals(arr1, arr2)) {//这个工具方法挺有用
    System.out.println("equals");
}
```
List使用equals方法即可。并未有Lists方法。  

以下覆盖父类equals方法，有个明显的错误。  
```java
class Employee{
    public boolean equals(Employee otherObj){
        return Objects.equals(name,otherObj.name)
            && salary == otherObj.salary
            && Objects.equals(hireDay,other.hireDay)
    }
    ....
}
```
`这个方法的显示参数是Employee，其结果并没有覆盖Object类的equals方法，而是定义了一个完全无关的方法。`    
为了避免发生类型错误，可以使用`@Override对覆盖超类的方法进行标记`。  
```java
@Override
public boolean equals(Object otherObj){....}
```
如果出现了错误，并且正在定义一个新的方法，编译器就会给出错误报告。
加入在上面有个明显的错误的方法上添加上@Override，编译器在编译的时候就会给出错误报告，而不是在运行时一点错误没有，但同时却不是我们想要的结果。  


## hash code 散列码
散列码是由对象导出的一个整型值。散列码是没有规律的。如果x和y是两个不同的对象，x.hashCode()和y.hashCode基本不会相同。  
`由于hashCode方法定义在Object类中，因此每个对象都会有个默认的hashCode值，其值为对象的存储地址`。  
```java
String s = "Ok";
StringBuilder sb = new StringBuilder(s);
String t = s;
StringBuilder tb = new StringBuilder(t); 
```
四个对象的hashcode s和tb相等  sb和ts都与之各不相同。
字符串的散列码是由内容导出的，所有s和sb是相同的，而字符串缓冲sb与tb则有着不同的散列码，`这是因为在StringBuilder和StringBuffer类中没有定义hashCode方法，它的散列码是由Object类的默认hashCode方法导出的对象存储值`。  
注意：  
1. 在通过Intellij Idea查看StringBuilder或StringBuffer时，搜索hashCode其设置会自动跳转到父类Object类中
2. 如果重新定义equals方法，就必须重新定义hashCode方法，以便用户可以将对象插入到散列表中。  
3. hashCode方法
```java
Objects.java

static int hashCode(Object obj){
    return obj!=null?obj.hashCode():0;
}
```

## toString()方法
System.out.println() 内部调用toString()方法。  

```java
public void println(Object x) {
        String s = String.valueOf(x);
        synchronized (this) {
            print(s);
            newLine();
        }
    }

public static String valueOf(Object obj) {
        return (obj == null) ? "null" : obj.toString();
}
```
之所以会打印出类名+hashCode，是因为PrintStream类没有覆盖toString方法。  
遗憾的是数组继承了toString而没有重写，我们无法使用toString直观的看到数组内部元素。  
Arrays.toString(arr);

## ArrayList
可变泛型数组
new ArrayList<>(100) // capacity is 100 代表一种潜力，拥有至少存储100个单元的空间(初始化构造后，并没有任何元素)
new int[100] //size is 100 已经分配了空间

```java
ArrayList<Object> arrLst = new ArrayList<>(10);
int[] intArr = new int[10];
System.out.println(Arrays.toString(intArr));
System.out.println(arrLst);

//[0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
//[]
```
注意即使使用ensureCapacity(int n)方法，也基本一样，只是内部分配了一个长度为n的数组。实际上没有太大变化，只是估计上的大小而已。  

数组列表的自动扩容的便利增加了访问元素语法的复杂程度。ArrayList类是被有些人编写且被放在标准库的一个实用类...

数组列表，也可以说是动态数组，一种更方便的数组，对于数据多的时候，增删元素性能差，需要考虑使用链表结构。

动态数组转换为普通数组
```java
List alst = new ArrayList();
alst.add(e);
...
X[] xarr = new X[];
xarr = alst.toArray();
```

java为了兼容性，会在编译时将类型泛型参数去除。类型转换(ArrayList)和(ArrayList<Employee>)将执行相同的运行时检查。  
有编译器检查警告时，可以在保证转换正确的前提下忽略。  

对于泛型数组列表，尖括号中的类型不允许是基本类型，此时需要基本类型的包装类。包装类是final的，不可变不能继承。   
```java
ArrayList<Integer> array = new ArrayList<Integer>();
```
**由于每个值分别包装在对象中，所以ArrayList<Integer>的效率远远低于int[]数组。因此，应该用它构造小型集合，其原因是程序员操作的方便性要比执行效率更加重要。**  

自动装箱  
类型自动转换。
list.add(3)自动转换为list.add(Integer.valueOf(3))  
自动拆箱
`int n = list.get(i)`自动转换为 `int i = list.get(i).intValue();`
在算数运算表达式中也能够自动的装箱和拆箱。可以将自增操作符应用于一个包装器引用:
```java
Integer n = 3; // n.intValue();
n++;
```

一般情况下，我们肯定认为
```java
Integer a = 1000;//自动装箱
Integer b = 1000;
if(a==b) ...// 此时不成立

//如果a,b介于-128~127之间呢?
```
因为检测的是否指向同一片存储区域。    
但在java自动装箱规范中，boolean,char,byte<=127，介于-128~127的short和int将被包装在固定的对象中。 即上述代码中，a=1000,b=100,将==成立。   

最后强调一下，自动装箱和拆箱是编译器认可的，而不是虚拟机。编译器在生成必要的字节码时，插入必要的方法调用，而虚拟机只是执行这些字节码。  

字符串转换为数字，这种api封装在包装类中再合适不过。  
int a = Integer.parseInt("3");

有些人认为包装器类可以用来实现修改数值传参数的方法，然而这时错误的。由于java方法都是值传递，所以不可能编写一个能够增加整形参数值的java方法：
```java
public static void triple(int x){
    x = 3*x;
}
```
如果将int换为Integer呢？
```java
public static void triple(Integer x){ //won't work
    x = 3*x;
}
```
虽然这里使用了引用类型，但是，Integer包装类型是final的，不可变的，不能够使用这些包装器类修改数值参数。  
如果实在想要编写一个修改数值参数值的方法，就需要使用org.omg.CORBA包中定义的持有者(holder)类型，包含IntegerHolder，BooleanHolder等。每个持有者类型都包含一个公有域值，通过它可以访问存储在其中的值。  
```java
public void triple(IntHolder x){
    x.value = 3 * x.value;
}
```

## 可变参数
main方法可以写为 
```java
public static void mian(String...args){
    //一样可以编译通过并执行
}
```
## 枚举
```java
public enum Size{
   SMALL,MEDIUM,LARGE,EXTRA_LARGE; 
}
```
```java
public enum Size{
    SMALL("S"),MEDIUM("M"),LARGE("L"),EXTRA_LARGE("XL");
    private String abbreviation;
    public String getAbbreviation(){
        return abbreviation;
    }
    public Size(String abbreviation){
        this.abbreviation = abbreviation;
    }
}
```
```java
for (Size size : Size.values()) {
    System.out.println(size);
    System.out.println(size.toString());
    System.out.println(size.name());

    Size size2 = Enum.valueOf(Size.class, "MEDIUM");
    Size size1 = Enum.valueOf(Size.class, size.name());
    //...
    System.out.println(size1.ordinal());//下标从0开始
}
Size small = Enum.valueOf(Size.class, "SMALL");
System.out.println(small.toString());
System.out.println(small.getAbbreviation());
```

## 反射
能够分析类的能力的程序称为反射(reflective),反射的功能及其强大。  
反射库(reflection library)提供了一个极其丰富且精心设计的工具集，以便编写能够动态操纵java代码的程序。特别是在设计或运行时添加新的类时，能够快速的应用开发工具动态地查询新添加类的能力。  
```java
String className = "java.util.Date";
try {
    Class clazz = Class.forName(className);
    System.out.println(clazz.toString());
    Date date = new Date();
    System.out.println(date.getClass());
} catch (ClassNotFoundException e) {
    e.printStackTrace();
}
```
反射是一种功能强大且复杂的机制。使用它的人员主要是工具构造者，而不是应用程序员(当然想要提升还是需要掌握的)。  

### Class类
`在程序运行期间，Java运行时系统始终为所有的对象维护一个被称为`**`运行时的类型标识。`**    
`这个信息跟踪着每个对象所属的类。虚拟机利用运行时类型信息选择相应的方法执行`。  
我们可以通过专门的Java类访问上述信息。保存这些信息的类被称为Class，不要混淆哦。  
Object类中的getClass()方法会返回一个Class类型的实例。(说到Object类，我们应该想到继承，任何类都是Object类的子类)  

```java
Employee  e = new Employee();
Class clazz = e.getClass();
```
获取Class类对象的方式有两种
- Class.forName("xxx");
- Object.getClass();
Class对象代表类的对象，int.class 代表int的类型的对象,Employee.class

`一个Class对象实际上表示的是一个类型，而这个类型未必是一种类。`例如，int不是类，但int.class是一个Class类型的对象。

虚拟机为每个对象管理一个Class对象。因此可以利用==运算符实现两个类对象比较的操作。例如：  
if(e.getClass()==Employee.class){...}

又有Class对象包含了类的运行信息，所以还可以通过其newInstance()方法快速的创建一个类的对象实例。例如：  
`e.getClass().newInstance();//该方法调用默认的构造器(没有参数的构造器)初始化创建新的对象。如果这个类没有默认的构造器，就会抛出异常`。从这点上来看，最好给每个类显示的提供默认构造。因为即使有了其他构造，而没有默认构造，也可能会造成日后日用上的异常。  
将forName与newInstance配合使用  
```java
String s = "java.util.Date";
Object o = Class.forName(s).newInstance();  
```
问题是如果想要使用有参数的构造器呢？Class.newInstance()自然就不满足。  
而需要使用Constructor类的newInstance方法

## 异常
抛出异常比终止程序灵活的多。因为可以提供一个"捕获"异常的处理器(handler)对异常进行处理。 

如果没有提供异常处理，程序将会终止，并打印出堆栈信息。

异常有两种类型:`未检查异常`和`已检查异常`。  
- 已检查异常：`编译器将会检查是否提供了处理器`。有很多异常为已检查
- 未检查异常：也有很多常见的异常，例如null引用，都属于未检查异常。`编译器不会查看是否为这些错误提供了处理器。`     
`应该精心地编写代码来避免这些错误的发生，而不要将精力花在编写异常处理上。`    

并非所有的错误都是可以避免的。`如果竭尽全力还是发生了异常，编译器就要求提供一个处理器`。Class.forName方法就是一个抛出已检查异常的例子。

将可以抛出异常的一个或多个方法调用代码放在try块中，然后在catch子句中提供处理器代码。  
```java
try{
    //... 可能发生异常的代码块
}catch(Exception ex){
    //... 代码异常处理块
}finally{
    //... 
}
```
继承结构中，Throwable是Exception类的超类。该类有个printStackTrace方法可以打印出栈的轨迹。    
`对于已检查异常，只需要提供一个异常处理器。可以很容易地发现抛出已检查异常的方法。如果调用了一个抛出已检查异常的方法，而又灭有提供处理器，编译器将会给出错误报告。`







