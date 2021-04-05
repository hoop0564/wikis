# C++

## C++ 语言

1980 年，AT&T 贝尔实验室的 Bjarne Stroustrup 创建的 C++ 语言横空出世，它既可以全面兼容 C 语言，又巧妙揉和了一些面向对象的编程理念。



C++ 很大程度就是用来解决 C 语言中的各种问题和各种不方便的。比如：

- 用引用来解决指针的问题。
- 用 namespace 来解决名字空间冲突的问题。
- 通过 try-catch 来解决检查返回值编程的问题。
- 用 class 来解决对象的创建、复制、销毁的问题，从而可以达到在结构体嵌套时可以深度复制的内存安全问题。
- 通过重载操作符来达到操作上的泛型。（比如用>>操作符消除printf()的数据类型不够泛型的问题。）
- 通过模板 template 和虚函数的多态以及运行时识别来达到更高层次的泛型和多态。
- 用 RAII、智能指针的方式，解决了 C 语言中因为需要释放资源而出现的那些非常 ugly 也很容易出错的代码的问题。
- 用 STL 解决了 C 语言中算法和数据结构的 N 多种坑。



C++ 是如何有效解决程序泛型问题的：

**第一，它通过类的方式来解决。**

类里面会有构造函数、析构函数表示这个类的分配和释放。还有它的拷贝构造函数，表示了对内存的复制。还有重载操作符，像我们要去比较大于、等于、不等于。这样可以让一个用户自定义的数据类型和内建的那些数据类型就很一致了。

**第二，通过模板达到类型和算法的妥协。**

模板的特化会根据使用者的类型在编译时期生成那个模板的代码。模板可以通过一个虚拟类型来做类型绑定，这样不会导致类型转换时的问题。模板很好地取代了 C 时代宏定义带来的问题。

**第三，通过虚函数和运行时类型识别。**

- 虚函数带来的多态在语义上可以支持“同一类”的类型泛型。
- 运行时类型识别技术可以做到在泛型时对具体类型的特殊处理。

这样一来，就可以写出基于抽象接口的泛型。



## 泛型编程

一个良好的泛型编程需要解决如下几个泛型编程的问题：

- 算法的泛型；
- 类型的泛型；
- 数据结构（数据容器）的泛型。



为了解决泛型的问题，我们需要动用以下几个 C++ 的技术。

- 使用模板技术来抽象类型，这样可以写出类型无关的数据结构（数据容器）。
- 使用一个迭代器来遍历或是操作数据结构内的元素。



### C++ 泛型版search()函数：

C语言版：

```c
int search(void* a, size_t size, void* target, 
  size_t elem_size, int(*cmpFn)(void*, void*) )
{
  for(int i=0; i<size; i++) {
    if ( cmpFn (a + elem_size * i, target) == 0 ) {
      return i;
    }
  }
  return -1;
}
```

C++版：

```c++
template<typename T, typename Iter>
Iter search(Iter pStart, Iter pEnd, T target) 
{
  for(Iter p = pStart; p != pEnd; p++) {
    if ( *p == target ) 
      return p;
  }
  return NULL;
}
```

在 C++ 的泛型版本中，我们可以看到：

1. 使用typename T抽象了数据结构中存储数据的**类型**。
2. 使用typename Iter，这是不同的数据结构需要自己实现的“**迭代器**”，这样也就抽象掉了不同类型的数据结构。
3. 然后，我们对数据容器的遍历使用了Iter中的++方法，这是数据容器需要**重载的操作符**，这样通过操作符重载也就泛型掉了遍历。
4. 在函数的入参上使用了pStart和pEnd来表示**遍历的起止**。
5. 使用*Iter来取得这个“指针”的内容。这也是通过**重载 * 取值操作符**来达到的泛型。



### C++ STL 中的find()函数的代码。

```c++
template<class InputIterator, class T>
  InputIterator find (InputIterator first, InputIterator last, const T& val)
{
  while (first!=last) {
    if (*first==val) return first;
    ++first;
  }
  return last;
}
```



## Reduce

把整个迭代器值给你一个 operation：

```c++
template<class Iter, class T, class Op>
T reduce(Iter start, Iter end, T init, Op op) {
  T result = init;
  while (start != end) {
    result = op(result, *start);
    start ++;
  }
  return result;
}  
```

在 C++ STL 中，与我的这个 reduce 函数对应的函数名叫 accumulate()，其实际代码有两个版本。

```c++
template<class InputIt, class T>
T accumulate(InputIt first, InputIt last, T init)
{
  for (; first != last; ++first) {
    init = init + *first;
  }
  return init;
}
```

第二个版本，更为抽象，因为需要传入一个“二元操作函数”——BinaryOperation op来做 accumulate。

```c++
template<class InputIt, class T, class BinaryOperation>
T accumulate(InputIt first, InputIt last, T init, BinaryOperation op)
{
  for (; first != last; ++ first) {
    init = op(init, *first);
  }
  return init;
}
```

实际使用：

```c++
struct Employee {
  string name;
  string id;
  int vacation;
  double salary；
};

double sum_salaries = 
  reduce(staff.begin(), staff.end(), 0,0, 
        [](double s, Employee e)
         {return s + e.salary;});

double max_salary = 
  redule(staff.begin(), staff.end(), 0,0,
        [](double s, Employee e)
         {return s > e.salary ? s : e.salary;})
```
