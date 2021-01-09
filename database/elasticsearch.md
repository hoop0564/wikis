# ElasticSearch

- 全文搜索引擎，开源 
- java开发的，核心是Lucene.jar包
- 百度、谷歌、GitHub都在用的搜索引擎
- 只支持json数据格式
- Restful风格暴露出来的API
- 对标Solr：
  - solr安装部署麻烦，ES解压即可（需要jdk环境）
  - solr在数据量上来后，性能不行，ES管理GitHub千亿行代码
  - ES天生支持分布式部署
  - solr可以支持xml等各种然并卵的数据格式



## 基础介绍

- 服务端口9200，可到web网页查看当前ES的config：

  ```
  http://localhost:9200/
  ```

- 三剑客ELK之一

  - ElasticSearch
  - Logstatch - 日志收集
  - Kibana - web UI，高级的ES查询工具

- [elasticsearch-head](https://github.com/mobz/elasticsearch-head)是其简单的可视化web界面（需要nodejs环境）
  > npm i && npm i -g grunt && npm run start

  - open web: http://localhost:9100/
  - 添加elasticsearch.yml配置项以支持跨域访问：

    ```yml
    ...
    http.cors.enabled: true
    http.cors.allow-origin: "*"
    ```
    
  - 重启ES

- 就把ES当做一个数据库，索引就是一个库，文档就是库中的数据

- 目录结构

  ```
  bin 启动文件
  config 配置文件
  	log4j2	日志配置文件
  	jvm.options	java 虚拟机相关的配置
  	elasticsearch.yml		elasticsearch的配置文件！默认9200端口！跨域！
  lib		相关jar包
  logs	日志！
  modules	功能模块
  plugins	插件！
  ```



## 特点

- 索引就是一个数据库

- 文档就是库中的数据

- 类型就是一个数据表 - 7.x已经没有了

  



## 参考资料

- [【狂神说Java】ElasticSearch7.6.x最新完整教程](https://www.bilibili.com/video/BV17a4y1x7zq?t=852&p=1)

- [狂神说笔记](https://gitee.com/kuangstudy/openclass)

