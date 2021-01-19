# mongodb



## 分片

分片（sharding）是MongoDB将大型集合分割到不同服务器（集群）上。

和MySQL的分区方案相比，MongoDB几乎能自动完成数据的均衡分片。

MongoDB自带了一个叫做mongos的专有路由进程。mongos就是掌握统一路由的路由器，会将客户端发来的请求准确无误的路由到集群中的一个或者一组服务器上，同时把接收到的响应拼装起来发回到客户端。

分片集群架构节点：

| 组件          | 说明                                                         |
| ------------- | ------------------------------------------------------------ |
| config server | 存储集群所有节点、分片数据路由信息。默认需要配置3个config server节点 |
| mongos        | 提供对外应用访问，所有操作均通过mongos执行。一般由多个mongos节点。数据路由，和客户端打交道的模块。本身没有任何数据，也不知道怎么处理数据，去找config server？？ |
| mongod        | 存储应用数据记录。一般由多个mongod节点，达到数据分片目的。以chunk为单位存数据。 |



## chunk

shard server内部，mongo会把数据分为很多个chunk存储，chunk有两个用途：

1. **Splitting:** 当chunk大小超过配置的chunk size（默认64M）时，MongoDB后台进程会把大chunk切分成小chunk，避免产生过大的chunk
2. **Balancing：** balancer是MongoDB的后台进程，负责chunk的迁移，均衡各个shared server的负载。

如果单位时间存储需求很大，需要配置更大的chunk，以提高性能。

chunk只会分裂，迁移，但不会合并。



chunk大小选择，需根据具体业务：

1. 小的chunk：均衡时迁移速度快，分布更均匀。但分裂频繁，路由节点消耗更多资源。
2. 大的chunk：数据分裂少。数据库移动集中消耗IO。通过设置100-200M。



### shard key

MongoDB中的数据分片是以集合为基本单位的，集合中的数据通过片键（shard key）被分成多部分：

- 片键是在集合中选一个键，用该键的值作为数据拆分的依据
- 片键必须是一个索引
- 片键是每条记录都必须包含的，不可为空值

- **基于范围的分片方式**：自增的片键对写入和数据局均匀分布不友好，总是在一个分片上操作。但对查询很高效！

- **基于哈希的分片方式**：随机片键对数据的均匀分布效果好。但查询时mongos需要对结果进行归并，需要尽量避免这种查询。

  

### 分片集群部署操作

1. 与单独配置可复制集基本一样，多了个启动参数：

   > --shardsvr

2. 需要在数据库和集合中都配置：

   >sh.enableSharding("DbName")
   >
   >sh.shardCollection("Collection1"{片键})

3. 配置文件：mongo.conf，注意不同分片的复制集名称要互不相同：

   ```shell
   fork=true #父子进程方式启动: child process started successfully, parent exiting
   dbpath=/opt/mongo/data/db
   port=27017
   bing_ip=0.0.0.0
   logpath=/opt/mongo/logs/mongodb.log
   logappend=true
   # 复制集名称
   replSet=repl_1
   smallfiles=true
   shardsvr=true
   ```

4. 启动服务：

   ```bash
   ./mongod -f mongo.conf
   ```

5. 使用mongo客户端登录，添加初始化配置：

   ```javascript
   // 配置复制集的var变量
   var rsconf = {
   	_id: 'repl_1' // 和配置文件中的复制集名称相同
     members: [ // 复制集成员
     	{
     		_id: 1,
     		host: "xx.xx.xx.201:27017"
   		},
     	{
     		_id: 2,
     		host: "xx.xx.xx.202:27017"
   		},
     	{
     		_id: 3,
     		host: "xx.xx.xx.203:27017"
   		},
   
     ]
   }
   
   // 加载rsconf配置
   rs.initiate(rsconf);
   ```

   