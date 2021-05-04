# memo



## etcd的raft算法

- 节点会有三个状态：follower，candidate，leader
- 算法中有两个超时设置，一个是选举超时，一个是心跳超时
- 数据的强一致性依靠日志的复制机制
- 初始时，大家都是follwer，这时选举超时机制触发leader选举，超时是随机的一个时间区间150-300ms，最先发起选举的节点，自己就会成为candidate，如果后续的选举得到大多数节点的通过，自己就成为leader；如果同时存在各一半节点数的选举，那么此次选举不算，将重新来一次leader选举；
- 以后所有的数据修改，都提交到leader，开始时先产生一条未提交的修改日志，当leader通知其他的大多数节点都统一修改，就执行这次修改，并提交此次的修改。
- 脑裂
  - 当发生网路分区时，因为节点数是奇数个，多数的节点继续提供服务，少数的节点不可用
  - 到网络分区恢复正常，少数的节点同步多数节点的日志，数据状态实现最终一致

## goroutine 实现机制

- golang语言层面控制的，运行在线程上，队列式的
- 每个协程最多运行10ms



## go的gc

> go语言内置运行时Runtime，抛弃传统的内存分配策略，改为自主分配；使用google的Thread-Cache Malloc算法，自己管理内存池和预分配，不用每次内存分配都需要向进行系统调用。

> 该算法核心思想就是把内存做分级管理，从而降低锁的粒度。它将可用的堆内存采用二级分配（全局内存池、线程内存池）的方式进行管理。每个线程都会自行维护一个独立的内存池，进行内存分配时优先从某个线程中的内存池分配；当内存不足时，才会向全局内存池申请，已避免不同线程对全局内存池的频繁竞争。

- 基本策略：
  - 每次从操作系统申请一块大内存，以减少系统调用
  - 将申请的大内存切分成不同大小的小块，构成链表，供后续使用
  - 为对象分配内存是，只需要从大小合适的链表中提取一个即可
  - 回收对象内存时，将该小块归还给原链表，以便复用
  - 如果闲置内存过多，则尝试归还部分内存给操作系统，降低整体开销。

- go程序启动时，向操作系统申请一块内存空间，切成小块然后自己进行管理。
- 申请到的内存会被分成3个区域，分别为：
  - 512M的span区域
  - 16G的bitmap区域
  - 512G的arena区域
- 这些只是虚拟的地址空间，并不会真正地分配内存
- 内存管理组件，分为3部分组成：
  - cache：每个运行期的工作线程都会绑定一个cache，用于无锁object的分配
  - central：为所有cache提供切分好的的后备span资源
  - heap：管理闲置span，需要时想操作系统申请内存
  - 
  - 分别为并启动多个线程管理，每个线程管理一部分被切割为不同大小的内存片，以后的使用直接向这些线程申请，避免锁粒度的性能消耗，使用完再返回给内存调度器

## go的chann

- chann的数据结构是hchan：

  ```go
  type hchan struct {
  	//channel队列里面总的数据量
  	qcount   uint           // total data in the queue
  	// 循环队列的容量，如果是非缓冲的channel就是0
  	dataqsiz uint           // size of the circular queue
  	// 缓冲队列，数组类型。
  	buf      unsafe.Pointer // points to an array of dataqsiz elements
  	// 元素占用字节的size
  	elemsize uint16
  	// 当前队列关闭标志位，非零表示关闭
  	closed   uint32
  	// 队列里面元素类型
  	elemtype *_type // element type
  	// 队列send索引
  	sendx    uint   // send index
  	// 队列索引
  	recvx    uint   // receive index
  	// 等待channel的G队列。
  	recvq    waitq  // list of recv waiters
  	// 向channel发送数据的G队列。
  	sendq    waitq  // list of send waiters
  
  	// lock protects all fields in hchan, as well as several
  	// fields in sudogs blocked on this channel.
  	//
  	// Do not change another G's status while holding this lock
  	// (in particular, do not ready a G), as this can deadlock
  	// with stack shrinking.
  	// 全局锁
  	lock mutex
  }
  ```

  - 全局的mutex锁
  - 接收协程队列recvq和发送协程队列sendq
  - 数组实现的环形队列circlebuffer，对于有缓冲的channel，sendx和recvx表示读写的缓冲区索引

- 在接收协程接收到新的缓冲消息时，会顺便触发阻塞读协程的重新运行，反之亦然。
- 思考：通过通信来实现共享内存，而不是通过共享内存来实现通信。（CSP）

## mysql索引

- B+树
- 是由二叉树演变而来的N叉数，即子节点数不是2个，是n个
- 叶子节点数目和数据节点的数目一样多，便于范围查找
- 聚簇索引是索引和数据都存储在一起，都在叶子节点，一般主键索引都是聚簇索引
- 非聚簇索引的索引和数据是分开存放的
- 二级索引又名辅助索引，存储的是主键值，而非数据地址，通过二级索引查询时先找到二级索引存储的主键值，然后再通过主键索引查找到存储的数据。唯一索引、普通索引、前缀索引都是二级索引。
- 因为B+树有最左原则，所以复合索引会依赖第一个字段索引排序，每个叶子节点对应的数据是已经排序好的



## SQL的事务隔离级别

- Read Uncommitted: 读未提交，事务做的操作，即使没有提交，其他事务也是可见的，所有会有脏读。脏读即读取了错误的数据，因为可能数据操作会需要undo回滚。
- Read Committed: 读已提交，只能读到已提交的事务，大多数数据库的默认隔离级别，又名不可重复读，因为会出现幻读，幻读即前后两次查询的结果可能不一致，因为这个前后两次间隔了某个事务的操作完成。
- Repeatable Read: 可重复读，mysql的默认隔离级别
- Serializable: 可串行化，事务的最高级别，让所有事务串行执行
- MVCC
  -  Multi-Version Concurrency Control，InnoDB的多版本并发控制，解决了幻读中的幻行问题。
  - InnoDB的MVCC是对每行记录增加两个隐藏的列实现的，一列是行的创建时间，一列是行的删除时间，列值存储的实际是系统版本号
  - 事务的版本号是事务开始的系统版本号，用来和查询到的每行记录的版本号进行比较
  - 事务只查找小于或等于事务版本号的数据行，小于事务版本号的是事务开始前就存在的，等于事务版本号的是事务自身修改过的。
  - 行的删除版本，要么未定义，要么比事务版本号大，以确保事务读取到的数据，在事务开始前未被删除

## mongo事务

- start transaction...commit
- 

## redis



### 事务

- mulit
- watch
- exec
- discard
- 如果没有watch，就会都执行，有语法错误也都顺序执行，不会停下。若有watch，则会都执行，如果无语法错误，则都成功

### 分片

### 备份方式

- RDB：二进制方式

- AOF：记录修改指令

  

## 架构设计的弹力设计

## 数据库异步写

## websocket



## mongo分片

## nodejs中的stream流

## Docker

### 镜像制作

- 将 Dockerfile 置于一个空目录下，或者项目根目录下