# Skywalking 导入



## 各软件版本

| 名称                    | 版本   | 用途                     |
| ----------------------- | ------ | ------------------------ |
| Java（JDK）             | 15.0.1 | skywalking启动环境       |
| Skywalking              | 8.1.0  | 监控软件                 |
| Elasticsearch           | 7.9.3  | 分布式搜索和数据分析引擎 |
| SkyAPM.Agent.AspNetCore | 1.0.0  | .net core的监控探针包    |



## 部署配置

- 安装JDK15，配置环境变量：`JAVA_HOME`，此处为：

  ```环境变量
  C:\Program Files\Java\jdk-15.0.1
  ```

- H2的内存存储版本无需配置，是默认启动项配置，可作为开发使用：

  ```powershell
  cd /d E:\skywalking\apache-skywalking-apm-es7-8.1.0\bin
  startup.bat
  ```

- ES7的持久化版本，可作为正式环境的部署使用，配置操作如下：

  - 修改 skywalking的 `config/application.yml` ：

    ```yml
    ...
    storage:
      selector: ${SW_STORAGE:elasticsearch7} #原来是${SW_STORAGE:h2}
    ...
    ```

  - 运行ES7的启动脚本：

    ```powershell
    cd /d E:\skywalking\elasticsearch-7.9.3\bin
    elasticsearch.bat
    ```

  - 等ES7启动成功后，再启动skywalking：

    ```powershell
    cd /d E:\skywalking\apache-skywalking-apm-es7-8.1.0\bin
    startup.bat
    ```



- 登录web后台查看：http://localhost:8080

  

## 项目中导入探针

通过[.net core中承载外部程序集](https://www.cnblogs.com/fuxuyang/p/11819328.html)的方法，在启动时从外部程序集向应用添加增强功能，来实现skywalking探针的集成。

以Mainpage服务为例：

```powershell
cd /d D:\RTCloud\src\Services\MainPage\UIH.RT.Cloud.Mainpage.Web.Host

# 向项目添加指定版本的SkyAPM包
dotnet add package SkyAPM.Agent.AspNetCore -v 1.0.0
#dotnet add UIH.RT.Cloud.Web.Portal.csproj package SkyAPM.Agent.AspNetCore -v 1.0.0

# 使用：dotnet skyapm config [your_service_name] [your_servers] 来产生skywalking的探针配置文件
dotnet skyapm config Mainpage localhost:11800
```

- 修改**skyamp.json**配置文件，以支持连接到skywalking 8.X版本：

  ```json
  {
    "SkyWalking": {
        ...
      "HeaderVersions": [
      "sw8" //修改原来的sw6为sw8
    ],
      "Transport": {
      "ProtocolVersion": "v8", //修改原来的v6为v8
          ...
    }
    }
  }
  ```



- 配置启动依赖程序集：

  ```c#
  ...
  webBuilder
    .UseSetting(WebHostDefaults.HostingStartupAssembliesKey, "SkyAPM.Agent.AspNetCore") // 主机配置承载外部程序集SkyAPM
    .UseStartup<Startup>();
  ...
  ```

