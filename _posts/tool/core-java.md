## java基本程序设计结构
- hello world程序
- 结构注意事项
- main方法
- 返回值
- System.exit(0)//非正常退出非0
查看java.lang.System的源代码，我们可以找到System.exit(status)这个方法的说明，代码如下：

```java
    /**
        * Terminates the currently running Java Virtual Machine. The
        * argument serves as a status code; by convention, a nonzero status
        * code indicates abnormal termination.
        * <p>
        * This method calls the <code>exit</code> method in class
        * <code>Runtime</code>. This method never returns normally.
        * <p>
        * The call <code>System.exit(n)</code> is effectively equivalent to
        * the call:
        * <blockquote><pre>
        * Runtime.getRuntime().exit(n)
        * </pre></blockquote>
        *
        * @param      status   exit status.
        * @throws  SecurityException
        *        if a security manager exists and its <code>checkExit</code>
        *        method doesn't allow exit with the specified status.
        * @see        java.lang.Runtime#exit(int)
        */
    public static void exit(int status) {
        Runtime.getRuntime().exit(status);
    }
```
注释中说的很清楚，这个方法是用来结束当前正在运行中的java虚拟机。如何status是非零参数，那么表示是非正常退出。

1. System.exit(0)是将你的整个虚拟机里的内容都停掉了 ，而dispose()只是关闭这个窗口，但是并没有停止整个application exit() 。无论如何，内存都释放了！也就是说连JVM都关闭了，内存里根本不可能还有什么东西
2. System.exit(0)是正常退出程序，而System.exit(1)或者说非0表示非正常退出程序
3. System.exit(status)不管status为何值都会退出程序。和return 相比有以下不同点：return是回到上一层，而System.exit(status)是回到最上层
示例
在一个if-else判断中，
如果我们程序是按照我们预想的执行，到最后我们需要停止程序，那么我们使用System.exit(0)，
而System.exit(1)一般放在catch块中，当捕获到异常，需要停止程序，我们使用System.exit(1)。这个status=1是用来表示这个程序是非正常退出。

- int较为常用，而在特殊情况下标识世界人口，则需要使用long类型

- float用的极少，double类型最为试用，至少从默认类型上就可以看出。另外想想表示普通员工的薪资和高管的薪资上。double才更适用。
- 在金融领域要求无误差，精度极高的场景，必须使用BigDecimal
- 注释的三种 // /* */不允许嵌套  /** */ 
- 浮点类型 F 双精度浮点类型D(默认)
- Unicode字符编码
- char类型不建议使用，尽量使用String类型
```java
    //判断字符是否为字母，比如空格即为no
    char a = ' ';
    if (Character.isJavaIdentifierPart(a)) {
        System.out.println("ok");
    } else {
        System.out.println("no");
    }
```
- java中boolean不能与整型值转换！ 
- 不要在代码中使用$，尽管合法，但其一般运用在java编译器或其他工具生成的名字中
- int i,j//both they are integers 这种风格并不提倡。逐一声明每个变量可提高可读性
- 变量要先声明初始化，并且尽可能的声明靠近在第一次使用最近的地方。(保持良好的编程风格)
- 必须使用final表示常量。只能被赋值一次，不能被修改。命名一般为全大写 final CM_PER_INCH
- 类常量 `static final` 类中的其他方法都可以使用。

- 求商、余数
```java
    float b = 1.2111f;
    System.out.println(0 / b);

    System.out.println(15 / 2);//7
    System.out.println(15 % 2);//1
    System.out.println(15 / 2.);//7.5
    System.out.println(0/1.0);
```


- ==
```java
    if (word == "hello") {
        System.out.println("word=='hello'");
    } else if (word.substring(0, 3) == "hel") {//说明subString返回时产生的是新的字符串对象，源码实现中return new String(value)
        System.out.println("hel=='hel'");
    }
```
- 通常将运算符放到赋值号的左侧
```java
    a = a + b
    等价于
    a += b
    类推
    a *= b;
    a %= b;
```
- 在浮点数算数运算上，jvm的跨平台移植遇到重大挑战。最优性能和最理想结果的冲突。`strictfp`关键字的使用。
- 自增运算符和自减运算符
n++/n-- 后缀形式
++n/--n 前缀形式
当该表达式使用在另外的表达式中，则会出现困扰的地方。
```java
    int m = 7;
    int n = 7;
    int g = 2 * ++m;//g=16 m=8
    int z = 2 * n++;//z=14 n=8
    System.out.println("g=" + g + "m=" + m);
    System.out.println("z=" + z + "n=" + n);
```
建议不要在其他表达式内部使用++,这样编写的代码容易令人困惑,并会产生烦人的bug。

- `==`判断比较对象是否在同一个位置,即是否指向同一引用。

- 与、或按照"短路"方式求值
`&&` `||`
如果第一个操作数已经能够确定表达式，第二个操作数就不必再计算了。
短路运算不仅效率高，而且可以避免一些错误的发生。
例如:
`x!=0&&1/x>x+y`//no division by 0
- `? :`三元操作符(三目运算符)
返回x，y中较小的值
`x<y?x:y`
- 位运算符 // todo
- 数学函数与常量
Math.pow(x,a)//java中没有幂运算，可以借助Math.pow方法
Math.PI
Math.E
`注意：如果希望得到一个完全可测的结果比运行速度更重要的话，那么就应该使用StrictMath类。`它可以实现每个平台得到的结果相同.具体使用的是`fdlibm`类库

- 数值类型转换(二元操作时，需要转换为同一种类型，然后再做计算)
转换时遵循自动向上转换原则,double<-float<-long<-int

```java
    int nx = 123456789;
    float f = nx;
    System.out.println(f);//1.23456792E8

    nx = (int) f;
    System.out.println(nx);//123456792
```
- 强制类型转换
但如果非要向下转换，则为强制类型转换(可能会丢失一些信息)

```java
    double xx = 9.97;
    int yy = (int) xx;
```
通过代码执行结果来看，强制类型转换通过截断浮点类型的小数部分将浮点值转换为整形。
但是如果我们想要获得最接近的整数呢，则需要四舍五入
- 四舍五入round floor函数 
roud函数超过5的向上，且返回值为long类型，所以仍然需要进行(int)强制类型转换
floor函数则是取整
```java
    double x = 9.9881;
    int y = (int) x; //9
    System.out.println(y);
    int c = (int) Math.round(x);//10
    System.out.println(c);
    int d = (int) Math.floor(x);//9
    System.out.println(d);
```
- 强制转换类型表示范围超出
byte与int的互相转换
```java
    //int->byte intVal-256
    System.out.println((byte) 300);//44
    System.out.println((byte) 200);//-56
```

```java
    boolean zz = true;
    i = zz ? 0 : -1;
    System.out.println(i);
```

- +=右结合运算符
```java
    //+=右结合运算符
    //a += b += c;<=> a+=(b+=c)
```        
- 枚举
变量的取值在一个有限的集合内。例如衣服的型号，汽车的颜色，电视机的尺寸，披萨的大小
```java
    enum Color {RED, BULE, YELLOW, Grey}
    Color color = Color.RED;
    System.out.println(color.toString());
```
- String 字符串

1. java字符串本质是Unicode字符序列
"java\u2122"本质就是5个unicode字符 => j、a、v、a、™(\2122) 
```java
    System.out.println("java\u2122");
    java™
```
2. java没有内置的字符串类型，而是标准库提供了一个预定义类String
`每个用双引号括起来的字符串都是String类的一个实例`
```java
    String str = "";//an empty string 
    String greeting = "welcome";
```

3. 截取子串
subString函数可以对字符串进行截取，获取想要的字符串
```java
    String greeting = "Hello";
    String substring = greeting.substring(0, 3);
    System.out.println(substring);
```
String substring(int startIdx,int endIdx) 
参数为[)闭开区间(不包含endIdx下标),这样有个好处，可以方便计算子串的长度endIdx-startIdx.
`subString(a,b),子串长度为b-a`
比如"hello".substring(0,3),"`hello`的子串`hex`的长度3-0

4. 连接
字符串之间允许使用"+"进行连接。
字符串与非字符串进行拼接时，后者都会先被转成字符串。(任何一个字符串都能够被转成字符串)
```java
    int age = 13;
    String rating = "PG"+age;//PG13
```

5. 不可变字符串
java中未提供修改字符串的方法，即字符串是不可变的。
那么 String greeting = "Hello";
修改为Help!,无法通过修改lo为p!的方式进行修改。
可以通过提取与拼接的方式
String greeting = "hello";
greeting = greeting.substring(0,3)+"p!";

`由于不能修改java字符串中的字符,所以在java文档中将String类对象称为不可变字符串。`
greeting.index(0)='G';//编译不会通过

不能改变其值，但是可以其引用。是变量指向另一个地址。
而"Hello"还是不变。

通过拼接的方式效率确实不高，但是不可变字符串却有另一个优点:**`编译器可以让字符串共享`**。  
**`Java的设计者认为共享带来的高效率远远胜过于提取、拼接字符串所带来的低效率。`**  
**而且实践中往往并不经常进行字符串的修改，而是字符串的比较**
例外：将源自于文件或键盘的单个字符或较短的字符串汇集成字符串。Java提供了StringBuilder

http://www.jianshu.com/p/b04fbfe000e4

## equals
检测字符串是否相等  
s.equals(t)//s,t字符串变量
"Hello".equals(greeting)//也可以是字符串常量  
"Hello".equalsIgnoreCase(greeting)  
**一定不能使用==比较字符串相等**  
**==只能确定两个字符串是否在同一个位置上**
```java
    String greeting = "hello";
    if (greeting == ("hello")) {
        System.out.println("true");
    } else if (greeting.substring(0, 3) == "hel") {
        System.out.println("not ==");
    }
```

如果虚拟机始终共享将相同的字符串共享，就可以使用==运算符检测是否相等。  
但`实际上只有字符串常量是共享的`。  
而`+`或`substring`等操作产生的结果并不是共享的
所以就不能够使用`==`进行字符串的比较，因为其可能是不共享的，因此就不会在同一个位置上。  
严禁使用`==`进行字符串的比较
除了equals方法外，也可使用compareTo方法
s.compareTo(greeting)  
equals方法较为清晰，首选。

- 检查字符串不为null且非空  
`首先要检查对象不为null，因为如果在一个null值上调用方法，会出错`
```java
    String xy = "";
    if (xy!=null&&xy.length()!=0) {
        System.out.println("不为空");
    }
```
- 代码点与代码单元
char类型是由一个采用UTF-16编码表示unicode代码点的代码单元。  
大多数的常用unicode字符使用一个代码单元就可以表示。`而辅助字符串需要一对代码单元表示`。  
因此应该尽量避免使用char类型，以避免一些低级错误。  

代码点实际上就是指unicode表示的字符位置号
A 65
a 97
- 字符串的API 使用频率很高
    length()
    ```java
    int length = "xyz".length();
    System.out.println(length);
    ```

    replace()
    ```java
     String lo = "Help!".replace("p!", "lo");
     System.out.println(lo);//hello
    ```

    codeAtPoint(int idx)
    ```java
    int i1 = "aBcd".codePointAt(1);//索引为1位置的字符的unicode索引号 B->98
    System.out.println(i1);
    ```

    indexOf(....)
    ```java
    int i2 = "axyz".indexOf(97); //从第一个unicode编码位置97代表的字符在该字符串中是否存在flag,flag(存在)?0:-1
    System.out.println(i2);
    ```
    自然有lastIndexOf(...)的重载





- 爬虫、搜索、大数据、架构









