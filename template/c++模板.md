# 模板

参数化编程是模板的起源！

- 模板把函数或类要处理的数据类型参数化，表现为参数的多态性，成为类属
- 模板用于表达逻辑结构相同，但具体数据元素类型不同的数据对象的通用行为



## 函数模板

```c++
template <typename T>
void swap(T &a, T&b) {
  T tmp;
  tmp = a;
  a = b;
  b = tmp;
}

int main(void) {
  int a=2, b=3;
  swap(a, b); // 模板函数1
  //swap<>(a, b); // 模板函数1 complie ok
  //swap<int>(a, b); // 模板函数1 complie ok

  double x=10,y=8;
  swap(x, y); // 模板函数2

  char p='a';
  int data = 23;
	// swap<int>(p,data); // 语法错误：函数模板不提供隐式类型转化！
}
```



- 函数模板和模板函数的辨析：
- 当函数模板和普通函数都符合调用规则的时候，优先使用普通函数
  - 因为普通函数在编译的期间就生成了函数体
  - 而模板函数的生成需要在调用的时候，运行时才会编译？
- 重复模板实例！
  - 相同的模板可能只有前一个生效！
  - 称为：name mangling 命名混淆
  - 函数加命名空间就可以规避
- 函数模板定义在头文件中，譬如 .h 或 .hpp中



## STL分配器 allocator

STL中使用allocator为STL的容器分配内存，其中涉及到内存的数组形式、链表形式、数组+链表形式。

```c++
// TEMPLATE CLASS vector
template <class _Ty,
	class _Alloc = allocator<_Ty>>
  class vector
    : public _Vector_alloc<_Vec_base_types<_Ty, _Alloc>>
  {...
```

- tips：VS中build之后，再点击代码中的 vector，转到定义，可以看到上面的代码。

  

## 参考资料

- [C++ 模板技术与 STL实战开发](https://www.bilibili.com/video/BV1wJ411h7GC?p=1)
- 