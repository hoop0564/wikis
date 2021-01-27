# Java Tutorial

## 应用领域：

- 银行系统
- 支付系统
- 政企信息系统
- 大数据平台
- 网站后台
- SaaS云
- 手机APP：Android OS
- 云管理系统后台
- 电商系统后台
- 桌面工具

做分布式的复杂应用场景



## 特点

- **Java 语言是体系结构中立的：**

  Java 程序（后缀为 java 的文件）在 Java 平台上被编译为体系结构中立的字节码格式（后缀为 class 的文件），然后可以在实现这个 Java 平台的任何系统中运行。这种途径适合于异构的网络环境和软件的分发。

- **Java 语言是解释型的：**

  Java 程序在 Java 平台上被编译为字节码格式，然后可以在实现这个 Java 平台的任何系统中运行。在运行时，Java 平台中的 Java 解释器对这些字节码进行解释执行，执行过程中需要的类在联接阶段被载入到运行环境中。

- **Java 语言是多线程的：**

  在 Java 语言中，线程是一种特殊的对象，它必须由 **Thread** 类或其子（孙）类来创建。通常有两种方法来创建线程：其一，使用型构为 Thread(Runnable) 的**构造子类**将一个实现了 **Runnable** 接口的对象包装成一个线程，其二，从 Thread 类派生出子类并重写 **run** 方法，使用该子类创建的对象即为线程。值得注意的是 Thread 类已经实现了 Runnable 接口，因此，任何一个线程均有它的 run 方法，而 **run 方法中包含了线程所要运行的代码**。线程的活动由一组方法来控制。Java 语言支持多个线程的同时执行，并提供多线程之间的同步机制（关键字为 synchronized）。

- **Java 语言是动态的：**

  Java 语言的设计目标之一是适应于动态变化的环境。Java 程序需要的类能够动态地被载入到运行环境，也**可以通过网络来载入所需要的类。这也有利于软件的升级**。另外，Java 中的类有一个运行时刻的表示，能进行运行时刻的类型检查。



## 构成

JavaSE

数据库

前端

SSM框架

SpringBoot - 8天，微服务 ？？

SpringCloud - 7天，真的微服务！



## 概念

- JDK：Java Development Kit，包含JRE和JVM

- JRE：Java Runtime Environment

- JVM：Java Virtual Machine

  ![jdk](./images/jdk.png)

JAVA_HOME中有一个src.zip包，里面是java系统库的源代码！



## Java修饰符

Java可以使用修饰符来修饰类中方法和属性。主要有两类修饰符：

- 访问控制修饰符 : default, public , protected, private
- 非访问控制修饰符 : final, abstract, static, synchronized



## Java运行时

**Java既是编译型语言，也是解释型语言！**

![./images/java-run.png](./images/java-run.png)



**Java 源程序与编译型运行区别**

![./images/java-compile.png](./images/java-compile.png)

**变量类型：**

- 局部变量
- 成员变量

- **类变量**：类变量也声明在类中，方法体之外，但必须声明为 **static** 类型。



![IDEA-project-config](./images/IDEA-project-config.png) 



**Java-Doc注释**

```java
/**
* @Description: HelloWorld
* @Author: gzc
*/
public class HelloWorld {
    public static void main(String[] args) {
        // 输出文本
        System.out.println("Hello World");
    }
}

```





## service-manager Java版

- http服务
- 静态文件服务
- http路由
- 调起系统命令
- 发起http请求
- 访问pgsql
- 访问consul