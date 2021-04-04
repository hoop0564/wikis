# C++

## 泛型编程

### C++ 泛型版search()函数：

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

1. 使用typename T抽象了数据结构中存储数据的类型。
2. 使用typename Iter，这是不同的数据结构需要自己实现的“迭代器”，这样也就抽象掉了不同类型的数据结构。
3. 然后，我们对数据容器的遍历使用了Iter中的++方法，这是数据容器需要重载的操作符，这样通过操作符重载也就泛型掉了遍历。
4. 在函数的入参上使用了pStart和pEnd来表示遍历的起止。
5. 使用*Iter来取得这个“指针”的内容。这也是通过重载 * 取值操作符来达到的泛型。



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

