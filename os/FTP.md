# FTP

FTP（File Transfer Protocol，文件传输协议） 是 TCP/IP 协议组中的协议之一。

FTP协议包括两个组成部分：其一为FTP服务器，其二为FTP客户端。

其中FTP服务器用来存储文件，用户可以使用FTP客户端通过FTP协议访问位于FTP服务器上的资源。由于FTP传输效率非常高，在网络上传输大的文件时，一般也采用该协议。



**默认情况下FTP协议使用TCP端口中的 20和21这两个端口，其中20用于传输数据，21用于传输控制信息。**



但是，是否使用20作为传输数据的端口与FTP使用的传输模式有关：

- 如果采用主动模式，那么数据传输端口就是20；
- 如果采用被动模式，则具体最终使用哪个端口要服务器端和客户端协商决定。



## 工作方式

FTP支持两种模式：

一种方式叫做Standard (也就是 PORT方式，主动方式)

一种是 Passive(也就是PASV，被动方式)。 

Standard模式 FTP的客户端发送 PORT 命令到FTP服务器。

Passive模式FTP的客户端发送 PASV命令到 FTP Server。



### Port

FTP 客户端首先和FTP服务器的TCP 21端口建立连接，通过这个通道发送命令，客户端需要接收数据的时候在这个通道上发送PORT命令。 PORT命令包含了客户端用什么端口接收数据。在传送数据的时候，服务器端通过自己的TCP 20端口连接至客户端的指定端口发送数据。 FTP server必须和客户端建立一个新的连接用来传送数据。

### Passive

在建立控制通道的时候和Standard模式类似，但建立连接后发送的不是Port命令，而是Pasv命令。FTP服务器收到Pasv命令后，随机打开一个高端端口（[端口号](https://baike.baidu.com/item/端口号)大于1024）并且通知客户端在这个端口上传送数据的请求，客户端连接FTP服务器此端口，通过三次握手建立通道，然后FTP服务器将通过这个端口进行数据的传送。

很多[防火墙](https://baike.baidu.com/item/防火墙)在设置的时候都是不允许接受外部发起的连接的，所以许多位于防火墙后或内网的FTP服务器不支持PASV模式，因为客户端无法穿过防火墙打开FTP服务器的高端端口；而许多内网的客户端不能用PORT模式登陆FTP服务器，因为从服务器的TCP 20无法和内部网络的客户端建立一个新的连接，造成无法工作。



## 传输模式

FTP的传输有两种方式：ASCII传输模式和二进制数据传输模式。



## 传输速度

1000M网络速度是指bit位的速度，理论网络传输速度上限是1000/8=128MB/S

实测案例：

IIS 6的SERVER 配合WINDOWS FTP命令能到40-50MB/S,换一个LINUX下的LFTP客户端就能达到60-65MB/S的性能.

结论：

大多数的千兆FTP传输在30-40MB/S间波动是很正常的.



## 参考资料

- [FTP协议的主动模式和被动模式的区别](https://www.cnblogs.com/rainman/p/11647723.html)

- [FTP协议 - 百度百科](https://baike.baidu.com/item/FTP%E5%8D%8F%E8%AE%AE)

- [网络速率和FTP传输速度关系的问题](https://blog.csdn.net/wwwlh/article/details/5118457)

- [FTP和TCP的文件传输效率对比测试分析](https://blog.csdn.net/derr96677169/article/details/101332793?utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromMachineLearnPai2%7Edefault-1.vipsorttest&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromMachineLearnPai2%7Edefault-1.vipsorttest)

