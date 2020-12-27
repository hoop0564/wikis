# go语言学习与实战



## Map

### Map与工厂模式

- map的value可以是一个方法
- 与go的dock type接口方式一起，可以方便的实现单一方法对象的工厂模式



### Map自实现Set

go内置集合中没有Set的实现，可以用 **map[type]bool** 来构造

- 作为key的type为任意类型，value为bool值
- 保证了元素的唯一性
- 基本操作：
  - 添加元素
  - 判断元素是否存在
  - 删除元素
  - 元素个数



## 字符串string

在go语言中：

1. string是数据类型，不是引用或指针类型

2. string是只读的byte slice，len函数可以获取它的长度，长度是byte的个数

3. string的byte数组可以存放任意数据，不知肉眼可读的字符串，二进制数据也都可以

   

   

### Unicode UTF8

1. Unicode 是一种**字符集**（In English：code point）
2. UTF8是unicode的存储实现，转换为字节序列的规则

|     字符      |        中        |
| :-----------: | :--------------: |
|    Unicode    |      0x4E2D      |
|     UTF-8     |     0xE4B8AD     |
| string/[]byte | [0xE4,0xB8,0xAD] |



## struct结构体

### 空结构体的使用场景

- 定义空channel
- 定义只包含一堆接口的结构做方法适配



## Go语言的函数

GO语言中的函数是一等公民。

1. 可以有多个返回值
2. 所有参数都是值传递：slice，map，channel会有传引用的错觉，因为函数调用时复制了指针，但指针指向的同一块内存空间
3. 函数可以作为变量的值
4. 函数可以作为参数和返回值



## 接口和方法

```go

// 第一种方法在实例方法被调用时，实例会被值拷贝一次
// 第一种定义方式在实例对应方法被调用时，实例的成员会进行值复制
func (e Employ) GetString() string {

	return fmt.Sprintf("Id=%d, Name=%s", e.Id, e.Name)
}

// 第二种方法避免了内存拷贝，通常使用
func (e *Employ) GetString2() string {

	return fmt.Sprintf("Id=%d, Name=%s", e.Id, e.Name)
}
```



## Go routine



### Thread vs Routine

Java和Go做比较：

1. 创建时默认的stack大小
   - java5以后的Java thread stack默认是1M
   - goroutine的stack初始化大小是2K
2. 和KSE（Kernal Space Entity）的对应关系
   - Java thread 是 1:1
   - goroutine是 M:N



### select控制程序流

- 多路选择，每个case跟着的是一个阻塞事件，比如channel、timer
- 对case的条件成立的执行是无序的，和switch不同！如果多个case都满足，随机只进入其中一个case！
- 可以实现超时控制！



## channel

### channel的关闭

- 向关闭的channel发送数据，会导致panic
- v, ok <- ch; ok=true/false表示通道正常接收/通道关闭
- 所有的channel接收者都会在channel关闭时，立刻从阻塞等待中返回且ok=false。此广播机制常被利用，进行向多个订阅者同时发送信号。如：进程或协程的退出信号。



### channel实现信号量和互斥锁

```go
type Empty interface{}    // 空接口
type semaphore chan Empty // 信号量
```

- 实现信号量的P操作，就是不断的往channel中放入数据，当channel满时，其他协程就不能再往channel放数据了，而只能阻塞，知道有一个协程释放资源，也就是执行V操作
- V操作就是从channel中取出资源
- 当channel容量**只有一个**时，P和V的数量变为1，同时进行P操作的协程只能有一个，在该协程执行完P操作没有执行V操作时，其他协程只能等待，这就实现了访问临界区资源的mutex功能。



### close(channel)做任务取消



### 只读/只写channel



### 构建对象池

- 因为一些对象例如数据库连接池的创建性能可能较大，需要预先创建

- 对象池需要有获取和归还接口

- 使用buffered channel实现对象池

  > 高可用系统中的一个金句：slow response比quick failure 更糟糕！

```go
runtime.GC() // 主动释放一次GC 会清除sync.pool中缓存的对象
```



## context与任务取消

- 根context：通过context.Background()获得

- 子context：通过context.WithCancel(parentContext)来创建

  ```go
  ctx, cancel := context.WithCancel(context.Background())
  ```

- 当前context被取消时，基于他的子context都会被取消

- 接收取消通知：<-ctx.Done()



## 测试

### Unit Test单元测试

- 表格测试法

  ```go
  func TestSquare(t *testing.T) {
    inputs := [...]int{1,2,3}
    expected := [...]int{1,4,9}
    for i:=0; i<len(inputs); i++ {
      if (squar(inputs[i])!=expected[i]) {
        t.Error("unexpected !")
      }
    }
  }
  ```

  

- 代码覆盖率，且显示tlog

  ```shell
  go test -v -cover
  ```

- 断言

  > github.com/stretchr/testify/assert



### Benchmark性能测试

```bash
# 能看出内存alloc次数
go test -bench=. -benchmem
```

>goos: darwin
>goarch: amd64
>BenchmarkConcatStringByAdd-12            9171664               122 ns/op              16 B/op          4 allocs/op
>BenchmarkConcatStringByBytesBuff-12     20660036                56.2 ns/op            64 B/op          1 allocs/op
>PASS
>ok      _/Users/apple/Documents/wikis/golang/tutorial/go_learning/src/ch35_benchmark    2.922s



### BDD

> Behavior Drive Design 行为驱动开发

#### BDD in Go

 - 项目网站

   > https://github.com/smartystreets/goconvey

- 安装

  > go get -v -u github.com/smartystreets/goconvey/convey

- 启动 WEB UI

  > $GOPATH/bin/goconvey



## 反射-reflect

### 利用反射编写灵活的代码

- 按名字访问结构的成员

  ```go
  reflect.ValueOf(*e).FieldByName("Name")
  ```

- 按名字访问结构的方法

  ```go
  reflect.ValueOf(*e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)}) 
  ```

### 反射优缺点

- 可以构建key-value获取和赋值的万能程序
- 可读性变差，不如直接的set/get好，调试困难，也有性能问题



## Unsafe的不安全编程

- 不适合的场景：无意义的强制类型转换，其值可能有丢失，比如float64和int的unsafe.Pointer转换

  ```go
  i := 10
  f := *(*float64)(unsafe.Pointer(&i))
  ```

  

- 适合的场景：

  - 自定义了一个比如 type MyInt int，后面需要对int类的变量做使用MyInt的方法操作
  - 并发读写中，可以先把数据写到一个buffer内存中，再用atomic一次替换到读内存中，以后使用的读内存块就是最新的了