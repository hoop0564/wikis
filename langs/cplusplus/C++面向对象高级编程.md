

# C++面向对象高级编程

C++ 98：1.0

C++ 11：2.0

《Effective C++ 中文版 改善程序技术与设计思维的55个有效方法--侯捷》

《STL 源码剖析 --侯捷》



## 拷贝构造 & 拷贝赋值

`String`类:

```c++
class String
{
public:
  String(const char* cstr=0);
  // 以下三个为：Big Three
  String(const String& str);
  String& operator = (const String& str);
  ~String();
  char* get_c_str() const { return m_data};
private;
  char* m_data;
};
```

```c++
inline 
String::String(const char* cstr=0) 
{
  if (cstr) 
  {
    m_data = new char[strlen(cstr)+1];
    strcpy(m_data, cstr);
  }
  else 
  {
    m_data = new char[1];
    *m_data = '\0';
  }
}

String::~String() 
{
  delte [] m_data;
}

inline 
String::String(const String& str)
{
  // 直接取另一个object的private data：兄弟之间互为friend
  m_data = new char[strlen(str.m_data)+1];
  strcpy(m_data, str.m_data);
}

inline
String& String::operator=(const String& str)
{
  if (this === &str) // 检查自我赋值，如果是self assignment，但未检查，下面的代码会出错
    return *this;
  
  delete[] m_data;
  m_data = new char[strlen(str.m_data)+1];
  strcpy(m_data, str.m_data);
  return *this;
}
```





使用 `String`:

```c++
int main() {
  String s1();
  String s2("hello");
  // 拷贝构造
  String s3(s1);
  // 拷贝赋值 重载=符号
  s3 = s2;
}
```

- 如果类中没有写 拷贝构造 和 拷贝赋值 函数，编译器就会给出默认的，机制是一个bit一个bit的忠实的做内存拷贝

- 如果类中有指针型的成员变量，就需要明确给出 拷贝构造 和 拷贝赋值 函数，以免在被使用的时候，出现浅拷贝（理当是深拷贝）

  > class with pointer members must have copy ctor and copy op=



## 单例Singleton

```c++
class A {
public:
  static A& getInstance();
  setup() {...}
private:
  A();
  A(const A& rhs);
  ...
};

A& A::getInstance()
{
  static A a;
  return a;
}

// 使用
A::getInstance().setup();
```



## 复数类complex

> 此示例来自于 `STL` 源码，感谢 侯捷！

```c++
class complex
{
public:
  complex(double r=0, double i=0)
    : re(r), im(i)
    {}
  // return by reference 尽量
  complex& operator += (const complex&);
  double real() const {return re;}
  double imag() const {return im;}
private:
  double re, im;
  
  // 友元
  // do assignment plus 
  // 注意：第一个参数为引用，第二个参数为const引用！
  friend complex& __doapl (complex*, const complex&);
}
```

使用：

```c++
{
  const complex c1(2,1);
  cout << c1.real(); // 如果real()不是const成员函数，此处会报错！
  cout << c1.imag();
}
```

> 传参时，尽量都穿引用，因为引用是广义的指针，传的时候，只占4个字节的空间，更快。
>
> 如果不希望参数被修改，就传：`pass by reference to const`



## 操作符重载

```c++
ostream& operator << (ostream& os, const complex& x)
{
  return os << real(x) << ',' << imag(x);
}
```

<img src="../../images/operator-overloading1.png" alt="image-20210429082846058" style="zoom:50%;" />



<img src="../../images/operator-overloading2.png" alt="image-20210430080653395" style="zoom: 45%;" />

注意上面是全局函数，非成员函数。



## 友元

```c++
inline complex&
  __doapl(complex* ths, const complex& r) 
{
  // 自由取得friend的private成员
  ths->res += r.re;
  ths->im += r.im;;
  return *ths;
}
```

**相同class的各个object互为friend（友元）**

```c++
class complex
{
...
int func(const complex& param) { return param.re + param.im};
}

{
  complex c1(2,1);
  complex c2;
  c2.func(c1); // focus it !
}
```



### return by reference

好处：传递者无需知道接收者是以 `reference`接收

```c++
inline complex&
  __doapl(complex* ths, const complex& r)
{
  ...
  return *ths;
}

inline complex&
  complex::operator += (const complex& r)
{
  return __doapl(this, r);
}

complex c1(2,1);
complex c2(5);
c2 += c1;
c3 += c2 += c1; // 注意：+=的返回值不可为void，需要为complex&
```

<img src="../../images/return-by-reference-prohibit.png" alt="image-20210430081130627" style="zoom:55%;" />



注意：`typename()` 表示要创建的是临时对象，例如 `int(i)`, `complex(1,2)`



## 参考资料

- [C++面向对象高级编程(上)-基于对象＆面向对象](https://www.bilibili.com/video/BV1Lb4y1R7fs?p=7)

