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



## Actor模型的工作方式

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



### actor模型和skynet框架比较

隔离的环境（lua或者内存块） + 回调函数 + 消息队列

- actor定义
  - 用于并行计算
  - actor是最基本的计算单元
  - 基于消息计算
  - actor之间相互隔离

- skynet用框架实现了actor模型
  - 启动多个并发actor
  - actor之间通过消息进行沟通
  - actor拥有私有消息队列，存储有序的消息，mailbox，像邮件列表一样
  - actor通过回调来消耗消息

### skynet中的actor是什么样的结构

```c
struct skynet_context {
  void * instance; // *隔离的环境
  struct skynet_module * mod; // 具体的模块
  void * cb_ud;		// 回调携带的环境
  skynet_cb cb;		// *回调函数
  struct message_queue *queue;	// *消息队列
  FILE * logfile;
  uint64_t cpu_cost;	// cpu计算统计 in microsec
  uint64_t cpu_start;	// in microsec
  char result[32];
  uint32_t handle;		// actor句柄
  int session_id;
  int ref;
  int message_count;
  bool init;
  bool endless;
  bool profile;
  
  CHECKCALLING_DECL
}
```

- C语言actor

  - logger服务：service_logger.c

- lua语言actor

  - lua虚拟机

  - service_snlua.c来加载lua服务

    ```c
    struct snlua {
    	lua_State * L;	// lua虚拟机的指针
      struct skynet_context * ctx;
      size_t mem;
      size_t mem_report; // 内存统计
      size_t mem_limit;
    }; // 隔离的环境
    ```



### lua介绍

- 天然的沙盒环境



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
- 数据保存的问题
- 架构 - actor 拆分：功能进行拆分，热点进行拆分
- 指导思想：**简单可用，逐步优化，忌讳过渡优化**



## 逻辑验证

- 数据库：redis

- 队列的维护

  - 玩家上下线的问题 清除玩家数据
  - 进入游戏 清除玩家数据
  - luatest测试lua脚本

- 提前启动好进程/skynet.newservice，否则临时创建时会有时延，玩家有卡顿

  

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



## 多核开发有哪些解决方案？

## 多进程

- 消息队列解决以下问题，类似zeromq（bind/connect/send/recv） **最终一致性**：
  - 协议问题（拆包粘包处理，protobuffer）
  - 断线重连问题
  - 进程启动顺序问题
  - 负载均衡问题
  - 数据同步问题（监听发布，请求回应 ）
- RPC **强一致性**
  - 两个系统的状态保持一致（银行转账）

- ZoomKeeper 协调管理 分布式系统中的服务 文件系统
  - 配置项管理
  - 集群管理
  - 统一命名服务 相同业务用一个命名
  - 状态同步（zk集群中）
  - 分布式锁 
    - redis setnx(key,1) 如果服务断了 锁就无法释放
    - zk的以目录形式存在的锁，如果服务断了，锁会自动释放