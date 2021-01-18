# ELK



ElasticStack = ELK + Beats



## Beats

采集一切数据的beats，包括以下：

| types      | cn name     | 说明                                                         |
| ---------- | ----------- | ------------------------------------------------------------ |
| Filebeat   | 日志文件    | 用于监控、收集服务器日志文件，其已取代logstash forwarder     |
| Metricbeat | 服务指标    | 可定期获取外部系统的监控指标信息，其可以监控、收集Apache、haproxy、MongoDB、MySQL、nginx、postgresql、redis、system、zookeeper等服务； |
| Packetbeat | 网络流量    | 网络数据包分析器，收集网络流量信息，嗅探服务器之间的流量，解析应用层协议，并关联到消息的处理，其支持ICMP(v4 & v6)、DNS、HTTP、MySQL、postgresql、redis、mongodb、memcache等协议 |
| Winlogbeat | Win事件日志 | 用于监控、收集windows系统的日志信息                          |
| Heartbeat  | 健康检查    |                                                              |

是elastic公司开源的一款采集系统监控数据的代理agent，是在被监控服务器上以客户端形式运行的数据收集器的统称。可以直接把数据发送给elasticsearch或者通过logstash发送给elasticsearch。



## Elasticseartch

基于java，是个开源分布式搜索引擎，特点有：

1. 分布式
2. 零配置
3. 自动发现
4. 索引自动分片
5. 索引副本机制
6. restful风格接口
7. 多数据源
8. 自动搜索负载等



### 安装

新建elsearch用户，elasticsearch不支持root用户运行：

```bash
useradd elsearch
# 解压安装包
tar -xvf elasticsearch-6.5.4.tar.gz -C /itcast/es/
chown elsearch:elsearch itcast/ -R
su - elsearch
```

### 修改配置文件

使之可以异地访问：

```yml
# vim elasticsearch.yml
network.host: 0.0.0.0

# 如果上面不是localhost或127.0.0.1，ES会认为是生产环境，进而会需要更高的资源需求
# 修改下面两个参数可以配置：初始堆总内存、最大堆总内存
-Xms128m
-Xmx128m
```

### 启动：

```bash
# 前台启动
./elasticsearch
# 后台启动
./elasticsearch -d
```

### 停止：

```bash
# 查看java进程
jps
3386 elasticsearch
# 关闭ES
kill 3286
```



### elasticsearch-head安装

- chrome插件安装：搜索elasticsearch-head，安装后使用

- docker安装

```bash
# 拉去镜像
docker pull mobz/elasticsearch-head:5

# 创建容器
docker create --name elasticsearch-head -p 9100:9100 mobz/elasticsearch-head:5

# 启动容器
docker start elasticsearch-head
```

注意：

由于前后端分离，所以会存在跨域的问题，需要在服务端做cors的配置：

```yml
# vim elasticsearch.yml 添加：
http.cors.enabled: true
http.cors.allow-orgin: "*"
```



### 概念

> 索引

- 索引（index）：是elasticsearch对逻辑数据的逻辑存储，所以它可以分为更小的部分
- 可以把索引看成关系型数据库的表，索引的结构是为快速有效的全文索引准备的，特别是他不存储原始值。
- elasticsearch可以把索引存放在一台机器或者分散到多台服务器上，每个索引由一或多个分片（shard），每个分片可以有多个副本（replica）

> 文档

- 存储在elasticsearch中的主要实体叫文档（document）。用关系型数据库来类比的话，一个文档相当于一个数据库表中的一行记录。
- elasticsearch和MongoDB的文档类似，都可以有不同的结构。但elasticsearch的文档中，相同字段必须有相同类型。
- 文档由多个字段组成，每个字段可能多次出现在一个文档里，这样的字段叫 **多值字段（multivalued）**

- 每个字段的类型，可以是文本、数值、日期等。字段类型也可以是复杂类型，一个字段包含其他子文档或者数组。

> 映射

- 所有文档写进索引之前，都会先进行分析，如何将输入的文本分割为词条、哪些词条又会被过滤，这种行为叫做映射（mapping）。一般由用户自己定义规则。

> 文档类型

- 在elasticsearch中，一个索引对象可以存储很多不同用途的对象。例如，一个博客应用程序可以保存文章和评论。？
- 每个文档可以有不同的结构。
- 不同的文档类型不能为相同的属性设置不同的类型！例如，在同一个索引中的所有文档类型中，一个叫title的字段必须具有相同的类型。



中文分词：IK分词器



Beats：轻量级的数据采集器，其中的filebeat和metricbeat最重要

```bash
./filebeat -e -c itcast.yml
```



x server：连接linux系统的windows的UI工具



## Kibana

数据分析的可视化平台。基于nodejs，是一个开源和免费的工具，可以汇总、分析和搜索重要数据日志。

### 数据探索

1. 路径：

   >  导航栏中：Management -> Index Patterns -> Create index pattern：

2. 会自动显示出已有的ES中的索引库。或手动输入做查找：

   > metricbeat-*

3. Time Filter field name 选择 @timestamp，确认创建 【Create index pattern】

4. 导航栏中 Discover，可查看ES中的数据

5. metric-beat的dashboard安装

   1. 在 `metricbeat.yml` 中配置kibana地址：

      ```yml
      setup.kibana:
        host: :"localhost:5601"
      ```

   2. 安装仪表盘到kibana：

      ```bash
      ./metricbeat setup --dashboards
      ```

   3. 在kibana的导航栏Dashboards中可以看到仪表盘数据：

6. file-beat的dashboard安装

   ```bash
   # kibana仪表盘安装
   
   ./filebeat -c itcast-nginx.yml setup
   # 启动filebet
   ./filebeat -e -c itcast-nginx.yml
   ```

   

## logstash

基于java，是个开源的用于手机，分析和存储日志的工具。

logstash的采集工作已经被beats代替掉了，因为前者是java的，需要一个JVM，速度太慢。现在基本通过beats采集。



### 安装启动

```bash
# 检查jdk环境，需要1.8+
java -version

# 解压安装包
tar -xvf logstash-6.5.4.tar.gz

# 第一个logstash示例
bin/logstash -e 'input {stdin {}} output {stdout {}}'
```



### 配置

```yml
input { # 输入
	stdin {...} # 标准输入
}

filter { # 过滤，对数据进行分割、截取等处理
	...
}

output { # 输出
	stdout {...} # 标准输出
}
```

实际中可以通过filebeat收集数据，发送给logstash，也可以用logstash直接采集过滤



### 读取自定义日志

日志结构示例：

```log
2021-01-17 20:21:52|ERROR|读取数据出错|参数: id=1002
```

自定义配置文件：

```yml
# vim itcast-pipeline.conf
input {
	file {
		path => "/itcast/logstash/logs/app.log" # 读取一个日志文件
		start_position => "beginning"
	}
}

filter {
	mutate {
		split => {"message" => "|"}
	}
}

output {
	elasticsearch {
		hosts => ["192.168.1.11:9200","192.168.1.12:9200","192.168.1.13:9200"]
	}
}
```

启动：

```bash
./bin/logstash -f itcast-pipeline.confg
```

logstash捕获到的日志：

```json
{
  "host" => "node01",
  "message" => [
  	[0] "2021-01-17"
		[1] "ERROR"
		[2] "读取数据出错"
		[3] "参数: id=1002"
  ],
	"@version" => "1",
	"@timestamp" => 2021-01-17T20:30:22.294Z,
	"path" => "/itcast/logstash/logs/app.log"
}
```

此时在elasticsearch里可以看到新增了index：logstash-2021.01.17，其中会有数据



## 参考资料

- [Elastic Stack（ELK）从入门到实践](https://www.bilibili.com/video/BV1iJ411c7Az?t=82&p=60)



