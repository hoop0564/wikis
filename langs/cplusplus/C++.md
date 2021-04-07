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



### 泛型的本质

要了解泛型的本质，就需要了解类型的本质。

- 类型是对内存的一种抽象。
- 不同的类型，会有不同的内存布局和内存分配的策略。
- 不同的类型，有不同的操作。所以，对于特定的类型，也有特定的一组操作。

所以，要做到泛型，我们需要做下面的事情：

- 标准化掉类型的内存分配、释放和访问。
- 标准化掉类型的操作。比如：比较操作，I/O 操作，复制操作……
- 标准化掉数据容器的操作。比如：查找算法、过滤算法、聚合算法……标准化掉类型上特有的操作。
- 需要有标准化的接口来回调不同类型的具体操作……



所以，C++ 动用了非常繁多和复杂的技术来达到泛型编程的目标。

- 通过类中的构造、析构、拷贝构造，重载赋值操作符，标准化（隐藏）了类型的内存分配、释放和复制的操作。
- 通过重载操作符，可以标准化类型的比较等操作。
- 通过 iostream，标准化了类型的输入、输出控制。
- 通过模板技术（包括模板的特化），来为不同的类型生成类型专属的代码。
- 通过迭代器来标准化数据容器的遍历操作。
- 通过面向对象的接口依赖（虚函数技术），来标准化了特定类型在特定算法上的操作。
- 通过函数式（函数对象），来标准化对于不同类型的特定操作。

> **屏蔽掉数据和操作数据的细节，让算法更为通用，让编程者更多地关注算法的结构，而不是在算法中处理不同的数据类型。**



### Reduce

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



## 函数式编程

核心思想是将**运算过程**尽量写成一系列嵌套的函数调用，关注的是**做什么**而不是怎么做，因而被称为**声明式编程**。

函数式编程，它的理念就来自于数学中的代数。

```mathematica
f(x)=5x^2+4x+3
g(x)=2f(x)+5=10x^2+8x+11
h(x)=f(x)+g(x)=15x^2+12x+14

# 斐波拉契数列
f(x)=f(x-1)+f(x-2)
```

对于函数式编程来说，它只关心定义**输入数据和输出数据相关的关系，数学表达式里面其实是在做一种映射（mapping），输入的数据和输出的数据关系是什么样的，是用函数来定义的。**



### 特征

- stateless：函数不维护任何状态。函数式编程的核心精神是 stateless，简而言之就是它不能存在状态，打个比方，你给我数据我处理完扔出来。里面的数据是不变的。
- immutable：输入数据是不能动的，动了输入数据就有危险，所以要返回新的数据集。



### 优势

- 没有状态就没有伤害。
- 并行执行无伤害。
- Copy-Paste 重构代码无伤害。
- 函数的执行没有顺序上的问题。



### 柯里化（Curring）

柯里化，Currying，将一个函数的多个参数分解成多个函数， 然后将函数多层封装起来，每层函数都返回一个函数去接收下一个参数，这可以简化函数的多个参数。在 C++ 中，这很像 STL 中的 bind1st 或是 bind2nd。

```javascript
// 普通的add函数
function add(x, y) {
    return x + y
}

// Currying后
function curryingAdd(x) {
    return function (y) {
        return x + y
    }
}

add(1, 2)           // 3
curryingAdd(1)(2)   // 3
```

技术实践：

```javascript
// 实现一个add方法，使计算结果能够满足如下预期：
add(1)(2)(3) = 6;
add(1, 2, 3)(4) = 10;
add(1)(2)(3)(4)(5) = 15;

function add() {
    // 第一次执行时，定义一个数组专门用来存储所有的参数
    var _args = Array.prototype.slice.call(arguments);

    // 在内部声明一个函数，利用闭包的特性保存_args并收集所有的参数值
    var _adder = function() {
        _args.push(...arguments);
        return _adder;
    };

    // 利用toString隐式转换的特性，当最后执行时隐式转换，并计算最终的值返回
    _adder.toString = function () {
        return _args.reduce(function (a, b) {
            return a + b;
        });
    }
    return _adder;
}

add(1)(2)(3)                // 6
add(1, 2, 3)(4)             // 10
add(1)(2)(3)(4)(5)          // 15
add(2, 6)(1)                // 9

```

把一个字符串数组中的字符串都转成小写:

python版：

```python
# 函数式
def toUpper(item):
  return item.upper()
 
upper_name = map(toUpper, ["Can", "You", "Hear", "Me"])

print upper_name
# 输出 ['CAN', 'YOU', 'HEAR', 'ME']
```

C++版：

```c++
string s = "hello";
transform(s.begin(), s.end(), back_inserter(out), ::toupper);
```

**map 和 reduce 不关心源输入数据，它们只是控制，并不是业务。控制是描述怎么干，而业务是描述要干什么。**



### 函数式的 pipeline 模式

```python

def even_filter(nums):
    for num in nums:
        if num % 2 == 0:
            yield num
def multiply_by_three(nums):
    for num in nums:
        yield num * 3
def convert_to_string(nums):
    for num in nums:
        yield 'The Number: %s' % num

nums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
pipeline = convert_to_string(multiply_by_three(even_filter(nums)))
for num in pipeline:
    print num
# 输出：
# The Number: 6
# The Number: 12
# The Number: 18
# The Number: 24
# The Number: 30        
```



Python 的关键字 yield，它是一个类似 return 的关键字，只是这个函数返回的是 Generator（生成器）。所谓生成器，指的是 yield 返回的是一个可迭代的对象，并没有真正的执行函数。也就是说，只有其返回的迭代对象被迭代时，yield 函数才会真正运行，运行到 yield 语句时就会停住，然后等下一次的迭代。（ yield 是个比较诡异的关键字）这就是 lazy evaluation（懒惰加载）。



## 参考资料

- [编程范式游记](https://time.geekbang.org/column/article/2711)

- [详解JS函数柯里化](https://www.jianshu.com/p/2975c25e4d71)