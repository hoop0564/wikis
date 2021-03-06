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
2. 大的chunk：数据分裂少。数据库移动集中消耗IO。通常设置100-200M。



## shard key

MongoDB中的数据分片是以集合为基本单位的，集合中的数据通过片键（shard key）被分成多部分：

- 片键是在集合中选一个键，用该键的值作为数据拆分的依据
- 片键必须是一个索引
- 片键是每条记录都必须包含的，不可为空值

- **基于范围的分片方式**：自增的片键对写入和数据均匀分布不友好，总是在一个分片上操作。但对查询很高效！

- **基于哈希的分片方式**：随机片键对数据的均匀分布效果好。但查询时mongos需要对结果进行归并，需要尽量避免这种查询。

  

## 集群部署操作

### 环境信息

>os: centos
>
>mongodb: 4.x
>
>3台虚拟机：xx.201/202/203
>
>集群环境
>
>2个分片复制集:
>
>shard1（xx.201:27017 xx.202:27017 xx.203:27017 ）
>
>shard2（xx.201:27018 xx.202:27018 xx.203:27018 ）
>
>1个config复制集:
>
>（xx.201:28018 xx.202:28018 xx.203:28018 ）
>
>1个mongos节点



### 分片复制集配置

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

5. 使用mongo客户端登录，添加初始化配置（201、202、203都需要分别配置和使用）：

   ```javascript
   // 配置复制集的var变量
   var rsconf = {
   	_id: 'repl_1' // 和配置文件中的复制集名称相同
     members: [ // 复制集成员
     	{
     		_id: 1,
     		host: "192.168.1.201:27017"
   		},
     	{
     		_id: 2,
     		host: "192.168.1.202:27017"
   		},
     	{
     		_id: 3,
     		host: "192.168.1.203:27017"
   		},
   
     ]
   }
   
   // 加载rsconf配置
   rs.initiate(rsconf);
   ```



### 搭建config-server节点复制集

​	创建config配置文件：mongo-cfg.conf （201、202、203都需要分别配置和使用）

```yml
systemlog:
	destination: file # 文件类型的日志
	path: /opt/mongo/mongo-cfg/logs/mongodb.log
	logAppend: true # 追加
storage:
	journal:
		enabled: true
  dbPath: /opt/mongo/mongo-cfg/data # 数据存储位置
  directoryPerDB: true # 是否一个库一个文件夹
  wiredTiger: # 引擎配置
  	engineConfig:
  		cacheSizeGB: 1 # 最大使用的cache
  		directoryForIndexes: true
    collectionConfig:
    	blockCompression: zlib # 表压缩配置
    indexConfig:
    	prefixCompression: true
net:
	bindIp: 192.168.1.201
	port: 28018
replication:
	oplogSizeMB: 2048
	replSetName: configReplSet # 配置节点的复制集名称
sharding:
	clusterRole: configsvr # 告诉这是config server
processManagement: # 后台进程是fork
	fork: true
   
```

1. 启动配置复制集

   ```bash
   ./mongod -f /opt/mongo/mongo-cfg.confg
   
   ```

2. 客户端登录复制集

   ```bash
   # 
   ./mongo -host 192.168.1.201 -port 28018
   ```

3. 初始化命令：

   ```javascript
   // 在任意一台执行此命令
   rs.initiate(
     _id: "configReplSet",
     configsvr: true,
     members: [
     	{_id: 0, host: "192.168.1.201:28018"},
     	{_id: 1, host: "192.168.1.202:28018"},
     	{_id: 2, host: "192.168.1.203:28018"}
     ]
   )
   ```

   在mongo客户端的命令行窗口，等待数秒，复制集会选举出primary：

   >configReplSet:OTHER>
   >configReplSet:SECONDARY>
   >...
   >configReplSet:PRIMARY>

   

### mongos节点配置

- mongo配置文件：mongos.conf

  ```yml
  systemlog:
  	destination: file
  	path: /opt/mongo/mongos/log/mongos.log
  	logAppend: true
  net:
  	bindIp: 192.168.1.201
  	port: 28017 # mongos的服务监听端口
  sharding:
  	configDB: configReplSet/test201:28018,test202:28018,test203:28018 # testXXX是主机名称
  processManagement:
  	fork: true
  ```

  ```bash
  cat /etc/hosts
  ...
  192.168.1.201 test201
  192.168.1.202 test202
  192.168.1.203 test203
  ```

- 启动服务

  ```bash
  ./monogs -config mongos.conf
  ```

- 登录mongos节点

  ```bash
  ./mongo 192.168.1.201:28017
  ```

  

- 登录后客户端中执行命令，添加集群中的分片节点
  ```javascript
  // 切换到admin数据库才能操作
  use admin;
  
  // 添加shard1复制集
  db.runCommand( {
    addshard: "yidian_repl/192.168.1.201:27017,192.168.1.202:27017,192.168.1.203:27017",
    name: "shard1"
  })
  
  // 添加shard2复制集
  db.runCommand( {
    addshard: "yidian_repl2/192.168.1.201:27018,192.168.1.202:27018,192.168.1.203:27018",
    name: "shard2"
  })
  
  // 查询当前的分片信息
  db.runCommand({listshards: 1})
  
  // 查看当前sharding status
  sh.status()
  ```



## 测试

- 开启数据库分片部署

  ```javascript
  // 测试数据库开启分片
  db.runCommand({enablesharding: "testdb"})
  ```

- 创建分片的键（id）

  ```javascript
  // 测试时是空集合 并会自动在片键id上创建索引
  db.runCommand({shardcollection: "testdb.users", key: {id: 1}})
  
  // 如果是已经存在的表 需要声明索引
  use testdb
  db.users.ensureIndex({id: 1})
  ```

- 添加测试数据

  ```javascript
  let arr = []
  for (let i = 0; i < 5000000; i++) {
    let uid = i;
    let name = "nick-" + i;
    arr.push({id, name})
  }
  db.users.insertMany(arr)
  
  // 查看分片数据的状态
  sh.status()
  ...
  {id: 0} --> shard1 ..
  {id: 250000} --> shard1 ..
  {id: 375001} --> shard2 ..
  {id: 500003} --> shard2 ..
  ...
  
  ```

- 删除分片

  ```javascript
  db.runCommand({removeShared: "shard2"})
  ```

  



## 参考资料

- [mongo集群搭建](https://www.bilibili.com/video/BV1p4411J7sq?t=364)











