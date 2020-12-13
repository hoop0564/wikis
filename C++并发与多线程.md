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



## unique_guard

与lock_guard基本一致，但更灵活

- 所有权可以转移

- 对象生命期内允许手动加锁和释放锁：提供了lock/unlock/try_lock等控制接口

- 在程序抛出异常后，先前已被上锁的mutex对象可以正确进行解锁操作

  | std::lock_guard  | 更简单，没有多余的接口，构造函数时拿到锁，析构函数时释放锁，但更省时 |
  | ---------------- | ------------------------------------------------------------ |
  | std::unique_lock | 更灵活，提供了lock，try_lock, try_lock_for, try_lock_until, unlock等接口 |

  

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



## 参考资料

- [C++并发与多线程](https://www.bilibili.com/video/BV1Yb411L7ak?p=1)

  