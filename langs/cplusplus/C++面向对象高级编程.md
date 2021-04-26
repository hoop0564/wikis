

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

## 参考资料

- [C++面向对象高级编程(上)-基于对象＆面向对象](https://www.bilibili.com/video/BV1Lb4y1R7fs?p=7)

