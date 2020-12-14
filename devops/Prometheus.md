# 系统监控Prometheus



## 目标

监控硬件资源的使用情况，CPU/内存/GPU/消耗资源较多的TOP进程，可以按时间段查看历史数据，数据日志信息可以拷走



## 系统各软件模块介绍

|                           软件名称                           |             说明              | 使用版本 |
| :----------------------------------------------------------: | :---------------------------: | :------: |
|             [Prometheus](https://prometheus.io/)             |    刮取并存储时间序列数据     |  2.22.1  |
|                       Windows Exporter                       | 采集OS的cpu/内存/IO等监控数据 |  0.15.0  |
|                     Nvidia GPU Exporter                      |       采集GPU的监控数据       |    -     |
| [Postgresql Exporter](https://github.com/wrouesnel/postgres_exporter/releases) |   采集Postgresql的监控数据    |  0.8.0   |
|               [Grafana](https://grafana.com/)                |    可视化收集到的时序数据     |  7.3.1   |



## 系统架构图：

![Prometheus architecture](.\assets\architecture.png)



## 安装部署

![](.\assets\OS监控.jpg)

### 部署 Prometheus

1. 下载软件：[官网链接](https://prometheus.io/download/)
2. 配置 `prometheus.yml`：

```yml
...
scrape_configs:
  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
  - job_name: 'Windows'

    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.

    static_configs:
    - targets: ['localhost:9182']
      labels:
        instance: Windows

  - job_name: 'Nvidia SMI Exporter'
    static_configs:
    - targets: ['localhost:9202']
      labels:
        instance: GPU
```

3. 运行软件

```powershell
cd /d D:\softwares\system-monitor\prometheus-2.22.1.windows-amd64
prometheus.exe --web.listen-address="0.0.0.0:29090"
```



### 部署 Grafana

- 下载软件：[官网链接](https://grafana.com/grafana/download?platform=windows)

- 安装完成会自动以服务的方式运行
- 登录后台：http://localhost:3000
- 初始用户名 *admin* 密码 *admin* 



### 配置Prometheus数据源

- 路径: Grafana - Configuration - DataSource
- 选 Prometheus 数据源
- 填入Prometheus的服务地址, 形如: http://10.8.77.74:29090/



### 部署 Windows-Exporter

- 路径: Grafana - Create - Import

- [下载地址](https://github.com/martinlindhe/wmi_exporter/releases)，选择最新msi文件下载安装即可

- 完成安装后会自动创建一个开机自启的服务
- 访问 http://localhost:9182/metrics 会看到数据，说明数据采集器安装成功。
- 添加 Grafana 的 dashboard，ID为：**10467**，[添加方法](https://www.cnblogs.com/guoxiangyue/p/11777227.html)



### 部署 [Nvidia_smi_exporter](https://github.com/phstudy/nvidia_smi_exporter)

- 此应用原是github项目上用go语言为linux平台开发的，但这里改动源码一个地方，重新编译后即可在windows上使用：

  ```go
  //const NVIDIA_SMI_PATH = "/usr/bin/nvidia-smi" 注释后改为下面的：
  const NVIDIA_SMI_PATH = "C:\\Windows\\System32\\nvidia-smi.exe"
  ```

  

- 编译后生成可执行文件 **nvidia_smi_exporter.exe**，将此文件复制到另一台windows电脑上可直接使用，无环境依赖（GO语言的优点）。

- 访问：http://localhost:9202/metrics 会看到数据，说明数据采集器安装成功。

- 添加 Grafana 的 dashboard，ID为：**6387**



### 部署Postgresql Exporter

- [下载地址](https://github.com/wrouesnel/postgres_exporter/releases)

- windows定义环境变量，供其访问

- 添加 Grafana 的 dashboard，ID为：**9628**



## Todo

- 监控数据在哪里？是否可以拷贝走，在另一个环境里可重读

- 监控数据的压缩或只收集有用的监控指标

- UI使用

- Grafana更改监听端口

- MongoDB  Postgresql  RabbitMQ的exporter和dashboard

- GPU exporter注册为windows服务：[golang 编写windows服务](https://blog.csdn.net/weixin_33790053/article/details/92399459)

- 部署在74上

  - Exporter部署在实体机上
  - Prometheus和Grafana都部署在另一台机子上

- docker环境部署

- prometheus.yml配置参数详解！

  

## 参考资料

2. [Prometheus 入门与实践](https://developer.ibm.com/zh/articles/cl-lo-prometheus-getting-started-and-practice/)
3. [Prometheus+Grafana 安装配置](https://www.cnblogs.com/guoxiangyue/p/11772717.html)
4. [Prometheus 监控Windows机器](https://www.cnblogs.com/guoxiangyue/p/11777227.html)
5. [Windows监控：基于Prometheus+Grafana监控CPU、内存、磁盘、网络、GPU信息](https://blog.csdn.net/fly910905/article/details/108275219)
6. [Nvidia SMI Exporter](https://github.com/phstudy/nvidia_smi_exporter)
