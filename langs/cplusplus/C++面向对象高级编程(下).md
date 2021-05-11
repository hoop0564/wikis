# C++面向对象高级编程(下)

## function conversion 转换函数

<img src="../../images/cpp/conversion-function-vector.png" alt="image-20210508083306253" style="zoom:27%;" />

转换函数有两个方向的转换，都可以。上栗是反向转。有偏特化概念

## explicit

明白的，明确的；

告诉编译器不要暗度陈仓自动的做些什么事情

```c++
explict Fraction(int num, int den=1);
```



## pointer-like classes

![image-20210509083112903](../../images/cpp/pointer-like-classes.png)



## specialization 模板特化

特化就是泛化的反面，泛化就是泛型编程，特化就是反之：面对独特的类型做特殊的设计

```c++
// 泛化
template <class Key>
struct hash {};

// 特化
template<>
struct hash<char>
{
	size_t operator()(char x) const { return x; }
};

template<>
struct hash<int>
{
	size_t operator()(int x) const { return x; }
};
```

上面的是全泛化（full generalization）。



模板偏特化（partial generalizaion)

偏特化分为两种：

- 个数上的偏

  ```c++
  template <typename T, typename Alloc=...>
  class vector
  {
  	...
  };
  
  template<typename Alloc=...>
  class vector<bool, Alloc> // T被偏特化为bool，个数的偏
  {
  	...
  };
  ```

  

- 范围上的偏

  ```c++
  template <typename T>
  class C {...}; // 模板泛化
  
  template <typename T>
  class C<T*> {...}; //模板偏特化（范围偏，泛化中多了个特殊的指针型模板）
  
  {
    C<string> obj1;
    C<string*> obj2;
  }
  ```




## 标准库

<img src="../../images/cpp/STL.png" alt="image-20210512065638695" style="zoom:43%;" />

```c++
// 确认C++版本：macro __cpluscplus 此值是如何编译器都要设的值
{
  cout << __cpluscplus;
}
// 201103
```



## variadic template 数量不定的模板参数