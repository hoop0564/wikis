

# nginx

nginx是异步框架的web服务器，也可以用作反向代理、负载平衡以及作为缓存服务器。nginx是目前互联网公司web服务器的主流计数，用于处理高并发甚至海量并发的网站数据。分为开源的社区版和闭源的商业版，**Tengine**就是淘宝在nginx基础上进行二次开发，以获取更高的稳定性和并发能力，经历了双十一的技术洗礼，足以证明其稳定性和高性能。



## 主要特性

- **高并发、高性能**：单台普通的服务器可以轻松处理上万并发连接（最多建议3w）

- **模块化设计，非常好的扩展性**：通过加载、卸载某个模块以实现相应的功能，

- **热部署、热更新**：支持配置文件的热更新，版本热升级、动态加载模块、日志热更换

  > reload操作的执行流程：nginx会新创建一个新的nginx进程，使用了新的配置文件；老的nginx停止listen服务，且处理完所有当前的请求之后，再graceful stop。

- **内存低消耗：据统计**，1w个keep-alive连接模式下的非活动连接，仅消耗内存2.5M

  > 非活动连接: 建立了连接，open状态，但此时并没有发生请求的发送和回应

- **配置、维护简单**：nginx的配置非常简单，运维非常友好



## 基本功能

- **web服务器**: 最基本的功能，也是非常重要的功能之一
- **反向代理服务器**：http协议的反向代理服务器，是生产环境中最常用的功能
- **FastCGI(php)、uWSGI(python) 代理服务器**：此时请求不是http协议，是跟后端服务相关的协议
- **TCP/UDP代理服务器**：也即作为四层调度器

- **Mail邮件代理服务器**：几乎不怎么使用了现在。



## 基础架构

如下图，`nginx`为`master/workers`架构，一个`master`主进程，负责管理和维护多个`worker`进程，真正接收并处理用户请求的其实是`worker`进程，`master`不对用户请求进行处理。即`master`主进程负责分析并加载配置文件，管理`worker`进程，接收用户信号传递以及平滑升级等功能。

`nginx`具有强大的缓存功能，其中`cache loader`负责载入缓存对象，`cache manager`负责管理缓存对象。



![image-20210215193356998](./pictures/nginx-framework.png)





## mac下安装

```bash
# 常规操作
brew update
# 检查要安装的软件是否存在
brew search nginx
# 查看nginx的相关信息
brew info nginx
# 安装nginx
brew install nginx
# 查看nginx安装目录
open /usr/local/etc/nginx
# 查看nginx可执行文件目录
open /usr/local/Cellar/nginx
# 启动nginx，没有报错即为启动成功
nginx
```

验证nginx是否启动成功：http://localhost:8080

nginx配置文件：/usr/local/etc/nginx/nginx.conf

nginx默认网站根目录：/usr/local/var/www

默认的索引文件为：index.html  index.htm



## nginx的源码安装和重新编译

官网下载解压后：

```bash
# 重新编译nginx
./configure --prefix=/usr/local/nginx --with-http_stub_status_module
make
sudo make install

  # 输出nginx的相关目录：
  nginx path prefix: "/usr/local/nginx"
  nginx binary file: "/usr/local/nginx/sbin/nginx"
  nginx modules path: "/usr/local/nginx/modules"
  nginx configuration prefix: "/usr/local/nginx/conf"
  nginx configuration file: "/usr/local/nginx/conf/nginx.conf"
  nginx pid file: "/usr/local/nginx/logs/nginx.pid"
  nginx error log file: "/usr/local/nginx/logs/error.log"
  nginx http access log file: "/usr/local/nginx/logs/access.log"
  nginx http client request body temporary files: "client_body_temp"
  nginx http proxy temporary files: "proxy_temp"
  nginx http fastcgi temporary files: "fastcgi_temp"
  nginx http uwsgi temporary files: "uwsgi_temp"
  nginx http scgi temporary files: "scgi_temp"
  
# 启动
sudo /usr/local/nginx/sbin/nginx 

./nginx -V #查询版本信息

suod vim /usr/local/nginx/conf/nginx.conf
        # 增加配置开启状态查询 才能查询到指标数据
        location /nginx-status {
            stub_status on; 
            access_log off;
        }   

```

安装包1M左右，C语言编写，5w并发。

支持的负载均衡方式：轮询、权重、IP hash、动静分离（静态资源和非静态需要后台做业务处理的）



## DDos攻击

分布式拒绝访问，Distributed deny of service

1. 原理
   - 多台不同主机发送假请求到服务端，使服务器处理不过来，导致正常的用户请求无法响应
2. 攻击方式
   - 攻击网络带宽，网络中的待处理包是有上限的。
   - TCP的握手信息连接表的数目是有上限的，Sync Flood攻击
   - web服务就发送很多的恶意的http请求，耗尽目标网站的资源，达到DDOS的目的，也叫CC攻击，挑战黑洞
3. 防御方式
   - 备用网站继续提供服务或发出公告
   - 看有无相同特征，如IP、User-Agent，然后进行拦截，可以使用：
     - 硬件防火墙，软件前面加上上，效果最好，但价格最贵
     - 软件防火墙，系统一般自带
     - web服务器软件，例如nginx
   - 如果没有特征，只能带宽扩容，较低ddos的攻击危害
   - 使用CDN，使得用户可以就近访问到CDN的资源，主站减少压力，所有请求先到CDN，如果CDN没有，从CDN上访问主站，但注意不要泄露主站的地址，但CDN只能放置静态资源。