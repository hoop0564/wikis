# Java原理篇



## Questions

- JVM的位置
- JVM的体系结构

- 类加载器
- 双亲委派机制
- 沙箱安全机制
- Native
- 三种JVM
- 栈、堆、堆内存调优



**bean作用域**

- singleton：全局唯一
- prototype：原型，每个对象它都有一个自己的
- ....



**IOC依赖反转**

```java
private UserDao userDao;

// 利用set进行动态实现值得注入！
public void setUserDao(UserDao userDao) {
  this.userDao = userDao;
}
```

- 之前是程序创建对象！控制权在程序员手术！

- 使用了set注入后，程序不再具有主动性，而是变成了被动的接受对象！程序员不用再去管理对象的创建了。

  > DI（Dependency Inversion）是实现IoC（Inversion of Control）的一种方法
  >
  > 采用xml方式配置bean的时候，bean的定义信息是和实现分离的；而采用注解的方式可以把两者合为一体，bean的定义信息直接以注解的形式定义在实现类中，从而达到了零配置的目的。

![image-20210227084933005](basics.assets/class-class.png)



**DDD领域驱动模型**

DDD把模型分为四层：

| 分层                           | 功能         | 举例                           |
| ------------------------------ | ------------ | ------------------------------ |
| UI层                           | 负责界面展示 | 商品列表                       |
| 应用层（Application Layer）    | 负责领域逻辑 | 购买商品                       |
| 领域层（Domain Layer）         | 负责领域逻辑 | 账单、用户、编辑商品、编辑库存 |
| 基建层（Infrastructure Layer） | 负责提供基建 | 持续储存、网络传输             |

> 领域模型应该捕捉“业务规则”或者“领域逻辑”（business rules/ domain logic）
>
> 应用模型则捕捉"应用逻辑"（application logic）

模型属于哪一层，有个粗略的判断方式：如果是一个实体（Entity）和针对实体的增删改查，就属于领域层；如果是一个场景，比如在UI菜单上的选项，就属于应用层。

领域模型只管“合规”，但不管“合理”。譬如在黑名单的客户不允许购买，这个检查通常在应用层做。



> 一个流可以理解为一个数据的序列。输入流表示从一个源读取数据，输出流表示向一个目标写数据。



## maven

IDE都是调用的maven做java项目的依赖管理和编译发布。

对不同的打包环境，可配置pom.xml中的project参数值，指定dev/qa/pre/prod环境的编译选项，对应项目的pom.xml的同级根文件中，也建立系统名称的目录：



### maven打包

**1. 添加profile配置到pom.xml：**

```xml
<profiles>
  <profile>
    <id>dev</id>
    <properties>
      <env>dev</env>
    </properties>
    <!-- 未指定环境时，默认打包dev环境 -->
    <activation>
      <activeByDefatult>true</activeByDefatult>
    </activation>
  </profile>

  <profile>
    <id>product</id>
    <properties>
      <env>product</env>
    </properties>
  </profile>

</profiles>
```

放入`pom.xml`的`dependencies`标签以外



**2. 对应的resources目录建立环境目录：**

![image-20210301080058011](basics.assets/image-20210301080058011.png)



**3. resources的资源文件配置：**

```xml
<resources>
  <resource>
    <directory>src/main/resources/${env}</directory>
  </resource>
  <resource>
    <directory>src/main/java</directory>
    <includes>
      <include>**/*.xml</include>
      <include>**/*.properties</include>
    </includes>
    <filtering>false</filtering>
  </resource>
</resources>
```

放入`pom.xml`的`build`标签中



**4. 执行打包操作**

打开 `Run/Debug/Edit Configuration` 窗口，`Command line`中配置打包命令：

```bash
clean compiler package -Pdev -Dmaven.test.skip-true
```

 