

# nginx



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