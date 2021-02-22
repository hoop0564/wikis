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



### Java修饰符

Java可以使用修饰符来修饰类中方法和属性。主要有两类修饰符：

- 访问控制修饰符 : default, public , protected, private
- 非访问控制修饰符 : final, abstract, static, synchronized



### Java运行时

**Java既是编译型语言，也是解释型语言！**

<img src="./images/java-run.png" alt="./images/java-run.png" style="zoom:50%;" />



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

### 虚函数

Java 中其实没有虚函数的概念，它的普通函数就相当于 C++ 的虚函数，动态绑定是Java的默认行为。如果 Java 中不希望某个函数具有虚函数特性，可以加上 final 关键字变成非虚函数。



### 重写

当子类对象调用重写的方法时，调用的是子类的方法，而不是父类中被重写的方法。

要想调用父类中被重写的方法，则必须使用关键字 super。



## Spring Cloud微服务

Spring Cloud 基于 Spring Boot，为微服务体系开发中的架构问题，提供了**一整套的解决方案**——服务注册与发现，服务消费，服务保护与熔断，网关，分布式调用追踪，分布式配置管理等。

> Spring Boot 是 Spring 的一套快速配置脚手架，使用默认大于配置的理念，用于快速开发单个微服务。



### Spring Cloud 完整技术：

![img](./images/spring-cloud.png)



### Spring Cloud 组件架构：

![img](./images/spring-cloud-framework.png)



**流程：**

- 请求统一通过 API 网关（Zuul）来访问内部服务。
- 网关接收到请求后，从注册中心（Eureka）获取可用服务。
- 由 Ribbon 进行均衡负载后，分发到后端具体实例。
- 微服务之间通过 Feign 进行通信处理业务。
- Hystrix 负责处理服务超时熔断。
- Turbine 监控服务间的调用和熔断相关指标。



### Spring Cloud工具框架

- **Spring Cloud Config 配置中心**，利用 Git 集中管理程序的配置。
- Spring Cloud Netflix 集成众多Netflix的开源软件。
- **Spring Cloud Netflix Eureka 服务中心**（类似于管家的概念，需要什么直接从这里取，就可以了），一个基于 REST 的服务，用于定位服务，以实现云端中间层服务发现和故障转移。
- **Spring Cloud Netflix Hystrix 熔断器**，容错管理工具，旨在通过熔断机制控制服务和第三方库的节点，从而对延迟和故障提供更强大的容错能力。
- **Spring Cloud Netflix Zuul 网关**，是在云平台上提供动态路由，监控，弹性，安全等边缘服务的框架。Web 网站后端所有请求的前门。
- Spring Cloud Netflix Archaius 配置管理 API，包含一系列配置管理API，提供动态类型化属性、线程安全配置操作、轮询框架、回调机制等功能。
- **Spring Cloud Netflix Ribbon 负载均衡**。
- **Spring Cloud Netflix Fegin REST客户端**。
- **Spring Cloud Bus 消息总线**，利用分布式消息将服务和服务实例连接在一起，用于在一个集群中传播状态的变化。
- Spring Cloud for Cloud Foundry 利用 Pivotal Cloudfoundry 集成你的应用程序。
- Spring Cloud Cloud Foundry Service Broker 为建立管理云托管服务的服务代理提供了一个起点。
- **Spring Cloud Cluster 集群工具**，基于 Zookeeper, Redis, Hazelcast, Consul 实现的领导选举和平民状态模式的抽象和实现。
- Spring Cloud Consul 基于 Hashicorp Consul 实现的服务发现和配置管理。
- **Spring Cloud Security 安全控制**，在 Zuul 代理中为 OAuth2 REST 客户端和认证头转发提供负载均衡。
- **Spring Cloud Sleuth 分布式链路监控**，SpringCloud 应用的分布式追踪系统，和 Zipkin，HTrace，ELK 兼容。
- Spring Cloud Data Flow 一个云本地程序和操作模型，组成数据微服务在一个结构化的平台上。
- **Spring Cloud Stream 消息组件**，基于 Redis，Rabbit，Kafka 实现的消息微服务，简单声明模型用以在 Spring Cloud 应用中收发消息。
- Spring Cloud Stream App Starters 基于 Spring Boot 为外部系统提供 Spring 的集成。
- Spring Cloud Task 短生命周期的微服务，为 Spring Booot 应用简单声明添加功能和非功能特性。
- Spring Cloud Task App Starters。
- Spring Cloud Zookeeper 服务发现和配置管理基于 Apache Zookeeper。
- Spring Cloud for Amazon Web Services 快速和亚马逊网络服务集成。
- Spring Cloud Connectors 便于PaaS应用在各种平台上连接到后端像数据库和消息经纪服务。
- Spring Cloud Starters （项目已经终止并且在 Angel.SR2 后的版本和其他项目合并）
- **Spring Cloud CLI 命令行工具**，插件用 Groovy 快速的创建 Spring Cloud 组件应用。



## 参考资料

- [Java 微服务架构选型](https://www.cnblogs.com/zengyjun/p/10309391.html)


