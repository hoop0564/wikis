# skynet



## 多核编程

- 多线程：统一性强，隔离性差，并发实体是线程
  - 消息队列
  - pipe+epoll/select/poll
  - 资源+锁（信号量、互斥锁、自旋锁、读写锁、原子锁）

- 多进程：统一性差，隔离性强，并发实体是进程
  - 共享内存
  - pipe管道
  - socket
  - 信号
  - 分布式中：CAP
- CSP模型
  - Communicating Sequential Process
  - go语言的goroutine+channel 轻量级的线程
- Actor模型
  - erlang：就是进程
  - skynet：就是轻量级的进程
  - actor就是轻量级的进程



## Actor模型

1. Actor 之间通过消息的方式沟通，用链表的方式构造的消息队列

2. actor 怎么运行的？当actor 消息队列有消息的时候，工作线程取出消息，并在actor环境中运行

3. skynet 中用回调函数实现工作线程取消息并执行回调函数

4. actor 把消息放入全局消息队列

5. 多线程的工作线程从全局消息队列中轮询，取消息消费

6. actor 怎么接收网络数据的？使用epoll：

   1. epoll_create: 创建红黑树，保存我们注册的事件

   2. epoll_ctl: 绑定事件 

   3. epol_event_t: 结构体，放入fd, 放入actorid

      ```c
      static int
      sp_add(int efd, int sock, void *ud) {
      	struct epoll_event ev;
      	ev.events = EPOLLIN;
      	ev.data.ptr = ud; //actiorid
      	if (epoll_ctl(efd, EPOLL_CTL_ADD, sock, &ev) == -1) {
          return 1;
        }
      	return 0;
      }
      ```

   4. epoll_wait

      ```c
      int nevents = epoll_waits(ev[])
      ev[].data.ptr => actorid, fd 
      ```



## 游戏案例

3人猜谜数字，如果人齐了，就开始游戏，中途不可以退出，支持万人同时在线



### 代码结构

- agent.lua : 玩家actor
- hall.lua: 大厅actor
- main.lua: 启动actor
- redis.lua: redis actor
- room.lua: 房间 actor



### 功能需求

- 断线重连
- 满3人开始游戏 匹配队列
- listenfd 网关里面
- connectfd 开不同的actor来处理 - - 轻量级线程！



### 核心代码

- main.lua

```lua
local skynet = require "skynet"
local socket - require "skynet.socket"

local function accept(clientfd, addr)
  	skynet.newservice("agent", clientfd, addr)
end

skynet.start(function()
    local listenfd = socket.listen("0.0.0.0", 8888)
    skynet.uniqueservice("redis")
    skynet.uniqueservice("hall")
    socket.start(listenfd, accept) -- 绑定actor 与 epoll
end)
```

