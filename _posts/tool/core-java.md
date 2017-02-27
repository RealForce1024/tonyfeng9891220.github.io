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
- java中boolean不能与整型值转换！ 
- 不要在代码中使用$，尽管合法，但其一般运用在java编译器或其他工具生成的名字中
- int i,j//both they are integers 这种风格并不提倡。逐一声明每个变量可提高可读性
- 变量要先声明初始化，并且尽可能的声明靠近在第一次使用最近的地方。(保持良好的编程风格)
- 必须使用final表示常量。只能被赋值一次，不能被修改。命名一般为全大写 final CM_PER_INCH
- 类常量 `static final` 类中的其他方法都可以使用。











