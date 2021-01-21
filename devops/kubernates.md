# kubernates

k8s（k和s中间有8个字母）源自于Google的borg资源管理器，10年+容器化基础架构。后使用golang实现的超大规模分布式部署的解决方案。

docker的标准是鲸鱼🐳，k8s的图标是船舵，表示方向盘领航的意思。



## 特点

- 轻量级：消耗资源少
- 开源
- 弹性伸缩：可大可小
- 负载均衡：IPVS



**IPVS：**

IP虚拟服务器（IP Virtual Server），基本上是一种高效的layer-4交换机！是运行在LVS下的提供负载平衡功能的一种技术。（from 章文嵩博士）[reference](https://baike.baidu.com/item/ipvs/5041817?fr=aladdin)

> 当一个TCP连接的初始SYN报文到达时，IPVS就选择一台服务器，将报文转发给它。此后通过查发报文的IP和TCP报文头地址，保证此连接的后继报文被转发到相同的服务器。这样，IPVS不用检查到请求的内容再选择服务器，这就要求后端的服务器组是提供相同的服务，不管请求被送到哪一台服务器，返回结果都应该是一样的。但是在有一些应用中后端的服务器可能功能不一，有的是提供HTML文档的Web服务器，有的是提供图片的Web服务器，有的是提供CGI的Web服务器。这时，就需要基于内容请求分发 (Content-Based Request Distribution)，同时基于内容请求分发可以提高后端服务器上访问的局部性。



**和apache的mesos比较**

也是分布式资源管理框架，Twitter之前使用的，2019年也放弃mesos，转向k8s



**和docker swarm比较**

是docker原厂出品，很轻量，本机只消耗几十MB。但功能相对于k8s太少，比如：滚动更新、回滚等操作，swarm手动实现起来很复杂。也能大规模化，但实现起来还是太费事。

阿里云也在2019年取消swarm，值支持k8s。



> 实践操作！follow the tutorial！



## Pod控制器

Pod控制器是k8s的灵魂！类型有：

- ReplicationController 和ReplicaSet
- Deployment
- DaemonSet
- Job
- CronJob
- StatefulSet
- Horizontal Pod Autoscaling



## 存储

configMap: 专门用于存储配置文件

Secret：存储一些比较重要的数据，比如用户名密码，需要加密的

volume：存一些基本的数据，比如网页文件

PV：是动态的调用过程

实际生产中，需要根据实际情况选择不同的存储方式。



## 调度器

k8s会自动调用容器和pod调度到对应的节点！

也能实现把pod定义到想要的节点运行！



## 集群安全

集群的认证、鉴权、访问控制。需要反复温故知新！



## HELM

相当于linux中的yum安装包管理器



## 运维

CICD构建 With Jenkins

kubeadm源码修改：目的是修改默认1年的证书限制

k8s高可用构建