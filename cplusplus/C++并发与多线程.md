# C++并发与多线程



## lock_guard

- 使用方式
  - 严格基于作用域的锁管理类模板
  - 构造时是否枷锁是可选的（不加锁时是假定当前线程已经获得锁的所有权）
  - 析构时自动释放锁
  - 所有权不可转移。？
  - 对象生存期内不允许手动加锁和释放锁

- 实现方式
  - 默认构造函数里锁定互斥量，即调用互斥量的lock函数
  - 析构函数利解锁互斥量，即调用互斥量的unlock函数
- 作用
  - 简化了mutex对象的上锁和解锁操作，方便线程对互斥量上锁
  - 即在某个lock_guard对象的声明周期内，它所管理的锁对象会一直保持上锁状态
  - 而lock_guard的生命周期结束之后，它所管理的锁对象会被解锁



## unique_lock

与lock_guard基本一致，但更灵活

- 所有权可以转移
- 对象生命期内允许手动加锁和释放锁：提供了lock/unlock/try_lock等控制接口
- 在程序抛出异常后，先前已被上锁的mutex对象可以正确进行解锁操作

| std::lock_guard  | 更简单，没有多余的接口，构造函数时拿到锁，析构函数时释放锁，但更省时 |
| ---------------- | ------------------------------------------------------------ |
| std::unique_lock | 更灵活，提供了lock，try_lock, try_lock_for, try_lock_until, unlock等接口 |



## lock选项

- std::adopt_lock选项在lock_guard和unique_lock中的含义相同，需要自己把mutex给lock住，两者都无需再做lock的动作，用于后面对象析构后的自动释放锁

- std::try_to_lock是试图获取锁，不能自己先做lock！它是不阻塞的：

  ```c++
  std::unique_lock(std::mutex) guard(mutex1, std::try_to_lock);
  if (guard.owns_lock()) {
    //拿到锁了
    recvQueue.push_back(i);
  }
  else {
    //没拿到锁
  }
  ```

- std::defer_lock：给出后续的没有加锁的mutex，针对unique_lock做灵活的lock和unlock的自定义操作！在单个业务流程中灵活地随时上锁和解锁！

- unique_lock的try_lock()方法返回bool，不阻塞，锁成功返回true，锁失败返回false。类似std::try_to_lock

- unique_lock的release()方法：返回mutex指针，并释放mutex的所有权！

- 所有权转移：unique_lock的对象可以把自己所拥有的mutex锁对象，释放给另一个unique_lock对象！

  ```c++
  std::unique_lock<std::mutex> guard2(std::move(guard1)); //左值转右值的移动语义,guard1指向空，guard2获得所有权
  
  std::unique_lock<std::mutex> rtn_unique_lock()
  {
    std::unique_lock<std::mutex> tmpguard(m_mutex1);
    return tmpguard; // 从函数返回一个局部的unique_lock对象是可以的。反正这种局部对象tmpguard会导致系统调用unique_lock的移动构造函数，其间会生成临时unique_lock对象！
  }
  ```

  

## 线程死锁

A线程持有了mutexA，在等待mutexB，B线程持有了mutexB，在等待mutextA，导致互相等待，导致线程死锁。



## 线程的坑

- 传递临时对象作为线程参数

- 传递类对象、智能指针做线程参数

- 用成员函数指针做线程参数

  

  ```c++
  void myprint(const int& i, char *pmybuf) {
  	cout << i << endl;	//调试发现，这里i是值传递！和外面的mvar的地址不一样
    cout << pmybuf << endl; //调试发现，这里pmybuf是指针传递！和外面的是mybuf地址一样！线程detach时会不安全！
  }
  
  class A {
    private:
    	int m_i;
    public:
    	A(int i): m_i(i){}
    	A(&A a): m_i(a.m_i) {}
    	~A(){}
    	void thread_work(int m)
    	void operator (int num) {} 	// thread mytobj(std::ref(myobj, 15))
  }
  
  int main() {
  	int mvar = 1;
  	int &mvary = mvar;
  	char mybuf[] = "this is buffer test!";
  
  	thread mytobj(myprint, mvar, mvary); // thread构造函数这里执行了参数的拷贝构造！所以i才是值引用
  	mytobj.detach(); // 异步的子线程和提前结束的主线程，引发线程资源回收的问题！
    
    std::chrono::millseconds dura(2000);
    std::this_thread::sleep_for(dura); //休息2秒
    
    // 类对象
    thread myobj2(myprinxUP, std::ref(objA));
    
    // 智能指针
    unique_ptr<int> mySmartPointer(new int(100));
    thread myobjUP(myprinxUP, std::move(mySmartPointer)); // 注意用thread.join()，不可用detach，否则内存可能泄露
             
    // 类的成员函数
    A myobj(5);
    thread mytobj2(&A::thread_work, &myobj, 15);	// &myobj == std::ref(myobj)
    return 0;
  }
  
  // 优化方法：void myprint(const int i, const string& pmybuf)
  // 使用方法：void myprint(mvar, string(mybuf));
  ```

  - tips: shift+F9查看指定变量名的内存值和地址
  - thread是模板类，如果要传引用的话，就用`std::ref(i)` ！
  - thread的构造函数会调用`std::forward`,把实参的控制权传递给线程对象去了！
  - return之后的临时变量无名，无名临时变量，thread构造函数的右值引用！
  - string()是构造函数，会先于线程执行！
  - 在创建线程的同时，创建**临时对象**来解决：`string(mybuf)`，代码逻辑是先构造，再拷贝构造
  - detach时：int这种参数类型就传值，类类型就用显示类型转换，用引用类型做形参接
  - 隐式类型转换是大坑，用explicit声明为显示类型：className(objectA)
  - 线程的存储空间：主线程先把数据是复制到子线程的存储空间，再由子线程取调用指定的函数去使用数据
  - const是常量左值引用（万能引用），临时变量具有常性
  - T&&类型，universal reference，万能引用，既能绑定右值，又能绑定左值
  - 用类对象做形参，或智能指针做形参，用`std::ref`



## condition_variable wait notify_one notify_all



## aync/future/packaged_task/promise

- std::async是个函数模板，用于启动一个异步线程

- std::future是个类模板

- std::async执行后，会返回一个std::future的类对象，其中的get方法可以获取线程的执行结果

  ```c++
  int myThread() {
    cout << "[begin] thread id is: " << std::this_thread::get_id() << endl;
    sleep(5000);
    cout << "[end] thread id is: " << std::this_thread::get_id() << endl;
    return 123;
  }
  
  std::future<int> result = std::async(myThread); // 绑定关系
  cout << "continuing.." << endl;
  int def = result.get(); // 线程卡在这里，等待thread执行完毕，get返回结果！def=123
  cout << def << endl;
  /*
  continuing..
  [begin] thread id is:1732
  [end] thread id is:1732
  123
  */
  ```

- std::future的对象析构时会自动调用`result.wait()`，导致调用线程等待

- 类方法调用方式：`std::async(&A:myThread, &a, tmpVal);`

- std::async中还可以再传递一个`std::launch::deferred`枚举变量，表示线程函数只有在调用`wait`或`get`时才开始执行，此时是直接在调用线程里同步执行的函数调用，没有产出新的线程！

- `std::packeged_task`打包任务：

  ```c++
  int mythread(int param) {
    ...
  }
  std::packeged_task<int(int)> mypt(mythread); // 打包多个任务
  std::thread t1(std::ref(mypt), 1); // 线程直接开始执行 param=1
  t1.join();
  std::future<int> result = mypt.get_future();
  ```

  ```c++
  std::packaged_task<int(int)> mypt([](int mypar)){
    ...
      return 5;
  }
  
  std::packaged_task<int(int)> mytasks;
  mytasks.push_back(std::move(mypt)); // 移动语义，入进去之后mypt就为空
  auto iter = mytasks.begin();
  mypt2 = std::move(*iter);
  mytasks.erase(iter);
  mypt2(123);
  
  ```

- `std::promise`，类模板

  ```c++
  void mythread(std::promise<int> &tmpp, int calc) {
    int result = calc++;
    tmpp.set_value(result); // 结果保存在这个tmpp对象中
  }
  std::promise<int> myprom;
  std::thread t1(mythread, std::ref(myprom), 1);
  t1.join();
  std::future<int> fu1 = myprom.get_future(); // promise和future绑定，用于获取线程返回值
  auto result = ful.get(); // get只能调用一次，不能调用多次
  ```

  

### future其他成员函数、shared_future、atomic

- std::future_status，通过std::future<T>.wait_for(timeout) 返回的枚举：
  - ready：执行完毕
  - timeout：执行超时
  - deferred：执行被延迟，使用了`std::launch::deferred`

- shared_future是类模板，和future功能相似，但是其get()函数可以执行多次，因为是拷贝复制

- atomic的是否原子和写法有关：

  ```c++
  std::atomic<int> count=0;
  count++;	// ok
  count += 1;	// ok
  count = count + 1; // fail
  ```

- std::async 异步执行任务

  ```c++
  std::future<int> result = std::async(mythread);
  count << result.get() << endl; // 线程可能在此处才开始创建并执行，根据系统繁忙情况吧！
  
  std::future<int> result = std::async(std::launch::deferred, mythread);
  count << result.get() << endl; // 此不创建新线程，是在主线程中执行！
  
  std::future<int> result = std::async(std::launch::async, mythread);
  count << result.get() << endl; // 强制创建新线程
  
  // 位标志
  std::future<int> result = std::async(std::launch::async | std::launch::deferred, mythread);
  count << result.get() << endl; // 可能也可能不创建新线程！
  ```

  std::async 和 std::thread的区别是前者可能不创建新线程，后者肯定创建新线程，但创建线程是有可以失败的！

  **经验：**一个程序里，线程数量不宜超过100-200。

- std::chrono

  ```c++
  std::async(mythread).wait_for(10s);
  
  chrono::seconds operator ""s(unsigned long long _val) {
    return (chrono::seconds(_val))
  }
  ```

  



## 参考资料

- [C++并发与多线程](https://www.bilibili.com/video/BV1Yb411L7ak?p=1)
- [中文版cppreference参考文档](https://github.com/myfreeer/cppreference2mshelp/releases)