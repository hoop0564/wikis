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
open  
# 启动nginx，没有报错即为启动成功
nginx
```

验证nginx是否启动成功：http://localhost:8080

nginx配置文件：/usr/local/etc/nginx/nginx.conf

nginx默认网站根目录：/usr/local/var/www

默认的索引文件为：index.html  index.htm



## centos下安装

```bash
[root@c93a66f92342 /]# rpm -ql nginx
/etc/logrotate.d/nginx
/etc/nginx
/etc/nginx/conf.d
/etc/nginx/conf.d/default.conf
/etc/nginx/fastcgi_params
/etc/nginx/koi-utf
/etc/nginx/koi-win
/etc/nginx/mime.types
/etc/nginx/modules
/etc/nginx/nginx.conf
/etc/nginx/scgi_params
/etc/nginx/uwsgi_params
/etc/nginx/win-utf
/etc/sysconfig/nginx
/etc/sysconfig/nginx-debug
/usr/lib/systemd/system/nginx-debug.service
/usr/lib/systemd/system/nginx.service
/usr/lib64/nginx
/usr/lib64/nginx/modules
/usr/libexec/initscripts/legacy-actions/nginx
/usr/libexec/initscripts/legacy-actions/nginx/check-reload
/usr/libexec/initscripts/legacy-actions/nginx/upgrade
/usr/sbin/nginx
/usr/sbin/nginx-debug
/usr/share/doc/nginx-1.18.0
/usr/share/doc/nginx-1.18.0/COPYRIGHT
/usr/share/man/man8/nginx.8.gz
/usr/share/nginx
/usr/share/nginx/html
/usr/share/nginx/html/50x.html
/usr/share/nginx/html/index.html
/var/cache/nginx
/var/log/nginx

[root@c93a66f92342 ~]# netstat -nlpt
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name    
tcp        0      0 0.0.0.0:80              0.0.0.0:*               LISTEN      349/nginx: master p 

[root@c93a66f92342 ~]# nginx -V
nginx version: nginx/1.18.0
built by gcc 4.8.5 20150623 (Red Hat 4.8.5-39) (GCC) 
built with OpenSSL 1.0.2k-fips  26 Jan 2017
TLS SNI support enabled
configure arguments: --prefix=/etc/nginx --sbin-path=/usr/sbin/nginx --modules-path=/usr/lib64/nginx/modules --conf-path=/etc/nginx/nginx.conf --error-log-path=/var/log/nginx/error.log --http-log-path=/var/log/nginx/access.log --pid-path=/var/run/nginx.pid --lock-path=/var/run/nginx.lock --http-client-body-temp-path=/var/cache/nginx/client_temp --http-proxy-temp-path=/var/cache/nginx/proxy_temp --http-fastcgi-temp-path=/var/cache/nginx/fastcgi_temp --http-uwsgi-temp-path=/var/cache/nginx/uwsgi_temp --http-scgi-temp-path=/var/cache/nginx/scgi_temp --user=nginx --group=nginx --with-compat --with-file-aio --with-threads --with-http_addition_module --with-http_auth_request_module --with-http_dav_module --with-http_flv_module --with-http_gunzip_module --with-http_gzip_static_module --with-http_mp4_module --with-http_random_index_module --with-http_realip_module --with-http_secure_link_module --with-http_slice_module --with-http_ssl_module --with-http_stub_status_module --with-http_sub_module --with-http_v2_module --with-mail --with-mail_ssl_module --with-stream --with-stream_realip_module --with-stream_ssl_module --with-stream_ssl_preread_module --with-cc-opt='-O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector-strong --param=ssp-buffer-size=4 -grecord-gcc-switches -m64 -mtune=generic -fPIC' --with-ld-opt='-Wl,-z,relro -Wl,-z,now -pie'

```

 

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



## 配置文件

```bash
# cat /etc/nginx/nginx.conf  

user  nginx;
worker_processes  auto;

error_log  /var/log/nginx/error.log warn;
pid        /var/run/nginx.pid;


events {
    worker_connections  1024;
}


http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;

    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    access_log  /var/log/nginx/access.log  main;

    sendfile        on;
    #tcp_nopush     on;

    keepalive_timeout  65;

    #gzip  on;

    include /etc/nginx/conf.d/*.conf;
}
```



### main主配置

#### 监听端口

root用户可以监听所有端口，普通用户只能监听1024以上的端口，nginx的80端口是master进程监听的root用户启动的。worker进程是nginx用户启动的。

```bash
[root@c93a66f92342 ~]# ps -ef|grep nginx
root       349     1  0 12:51 ?        00:00:00 nginx: master process nginx
nginx      350   349  0 12:51 ?        00:00:00 nginx: worker process
```



#### worker进程数

`worker_processes`配置为auto，即为CPU的核数。可通过指令 `lscpu`查看cpu信息。



#### pid

进程ID是为了运维此进程操作时使用的！

例如执行 `nginx -s reload` 时，就会从 `/var/run/nginx.pid中`读取pid，然后做reload操作。

`systemctl`也是读取文件的pid，来对进程做stop/restart操作的。

如果pid文件访问失败，可能是pid文件不存在或权限不对



#### events{...}

用于定义事件驱动相关配置，该配置与连接的处理密切相关，其中：

```bash
use method; # 定义nginx使用哪种事件驱动类型，在linux中性能最好的是epoll模型
accept_mutex on|off; # 处理新连接的方法，on是指各个worker进程轮流处理，off则会通知所有worker，但只有一个worker进程获得处理连接的权限（惊群现象）。在centos7中将使用 reuseport 会有更好性能。
```



### http配置

http配置段中，可以设置多个server配置，该server就是用来配置虚拟主机的，可以基于IP地址，也可以基于port，生产中更多基于域名的方式来配置虚拟主机。在server配置段中还可以配置多个location字段，该字段用来配置虚拟主机不同uri的响应方式

```bash
# /etc/nginx/conf.d/ip.conf
server {
	listen	127.0.0.1;	# 监听端口
	root	/data/nginx/ip; # web服务根目录
	index	index.html;	
}

# nginx -s reload

# curl http://localhost/
```







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