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

- elasticsearch-head是其可视化界面（需要nodejs环境）

  - 修改elasticsearch.yml以支持跨域访问，增加：

    ```yml
    ...
    # 允许使能
    http.cors.enabled: true
    # 所有域
    http.cors.allow-origin: "*"
    ```

  - 重启ES

- 就把ES当做一个数据库，索引就是一个库，文档就是库中的数据



## 参考资料

- [【狂神说Java】ElasticSearch7.6.x最新完整教程](https://www.bilibili.com/video/BV17a4y1x7zq?t=852&p=1)

- [狂神说笔记](https://gitee.com/kuangstudy/openclass)

