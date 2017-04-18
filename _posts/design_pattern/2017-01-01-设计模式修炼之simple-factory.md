## 简单工厂模式
角色 
- 客户端 
- 服务器端(接口\实现)
- 简单工厂(解耦调用方式)

### 传统简陋耦合方式 
客户端
```java
public class Client {
    public static void main(String[] args) {
        //客户端知道的太多，与服务实现紧密耦合
        IPrint implA = new ImplA();
        implA.print();
        IPrint implB = new ImplB();
        implB.print();
    }
}
```
接口
```java
public interface IPrint {
    void print();
}
```
服务器端实现
```java
class ImplA implements IPrint {
    @Override
    public void print() {
        System.out.println("ImplA print()");
    }
}

class ImplB implements IPrint {
    @Override
    public void print() {
        System.out.println("ImplB print()");
    }
}
```

### 简单工厂模式
添加工厂角色，使用工厂角色创建具体实现。职能是去除客户端实例化具体实现，解除调用端与实现端的紧耦合。

重点在于添加的工厂创建实例角色的方法
实现方式有以下几种:
注：IPrint,ImplA,ImplB保持不变。所以简单工厂只是改变调用方式进行解耦，并不改变原有接口和实现。

1. 根据不同的参数类型返回具体实例(还是有耦合)
```java
public class Client {
    public static void main(String[] args) {
        //通过工厂进行松耦合，实现与客户端分离
        IPrint IPrint = PrintFactory.getInstance(1);
        IPrint.print();
        //还是有不妥，客户端还是需要知道类型的区别，某种意义上还是没有真正的与服务端分离。
    }
}
```
```java
public class PrintFactory {
    public static IPrint getInstance(int type) {
        if (type == 1) {
            return new ImplA();
        } else {
            return new ImplB();
        }
    }
}
```

2. 根据服务端配置向客户端屏蔽具体实现(松耦合)
```java
public class Client {
    public static void main(String[] args) throws IOException, ClassNotFoundException, InstantiationException, IllegalAccessException {
        //此时可以采用配置文件的方式，客户端不需要传递任何业务类型相关的内容
        IPrint iPrint = PrintFactory.getInstance();
        iPrint.print();
    }
}
```
```java
public class PrintFactory {
    public static IPrint getInstance() throws ClassNotFoundException, IllegalAccessException, InstantiationException {
        Properties properties = PropertiesUtil.propertiesLoad("config.properties");
        String impl = PropertiesUtil.getString(properties, "impl", "com.design.pattern.simple_factory.refactor02.ImplB");
        Class<?> aClass = Class.forName(impl);
        IPrint ip = (IPrint) aClass.newInstance();
        return ip;
    }
}
```
```s
#impl=com.design.pattern.simple_factory.refactor02.ImplA
impl=com.design.pattern.simple_factory.refactor02.ImplB
```

3. 运行时动态决定

```java
public class Client {
    public static void main(String[] args) throws IOException, ClassNotFoundException, InstantiationException, IllegalAccessException {
        //运行时动态决定调用
        for (int i = 0; i < 10; i++) {
            IPrint ip = PrintFactory.getInstance();
            ip.print();
        }
    }
}
```
```java
public class PrintFactory {
    private static int count = 0;

    public static IPrint getInstance() {
        count++;
        if (count <= 5) {
            return new ImplA();
        } else {
            return new ImplB();
        }
    }
}
```