# cgo

unsafe包提供了一些跳过go语言类型安全限制的操作。



## 指针 - unsafe包的灵魂

Go版无类型指针和数值化的指针：

```go
var p unsafe.Pointer = nil		// unsafe
var q uintprt	= uintptr(p)		// builtin
```

C版无类型指针和数值化的指针：

```c
void *p = NULL;
uintptr_t q = (uintptr_t)(p);	// <stdint.h>
```

- `unsafe.Pointer` 是Go指针和C指针转换的中介
- `uintptr`是Go中数值化和指针转换的中介



## unsafe包

```go
type ArbitraryType int
type Pointer *ArbitraryType

func Sizeof(x ArbitraryTypr) uintptr
func Alignof(x ArbitraryType) uintptr

func Offsetof(x ArbitraryType) uintptr
```

- `Pointer`：面向编译器无法保证安全的指针类型转换
- `Sizeof`：值所对应变量在内存中的大小
- `Alignof`：值所对应变量在内存中地址几个字节对齐
- `Offsetof`： 结构体中成员的偏移量



C语言版本：

```c
typedef void* Pointer;
sizeof(type or expression); // C
offsetof(type, member);			// <stddef.h>
alignof(type-id);						// C++ 11
```

- C指针的安全性永远需要自己负责
- `sizeof`是关键字，语义和Go基本一致
- `offsetof`是宏，展开为表达式，语义和Go基本一致
- `alignof` 是新特性，可忽略



## Go字符串和切片的结构

```go
type reflect.StringHeader struct {
  Data uintptr
  Len int
}

type reflect.SliceHeader struct {
  Data uintptr
  Len int
  Cap int
}
```

```c
typedef struct {
  const char *p; 
  GoInt n;
} GoString;

typedef struct {
  void *data; 
  GoInt len; 
  GoInt cap;
} GoSlice;
```

- `reflect` 包定义的结构和CGO生成的C结构是一致的
- `GoString` 和 `Golice` 的头部结构是兼容的



## int32和*C.char相互转换

```go
// int32 => *C.char
var x = int32(9527)
var p *C.char = (*C.char)(unsafe.Pointer(uintptr(x)))

// *C.char => int32
var y *C.char
var q int32 = int32(uintptr(unsafe.Pointer(y)))
```

1. 第一步：int32 => uintptr
2. 第二步：uintptr => unsafe.pointer
3. 第三步：unsafe.Pointer => *C.char
4. 反之亦然



## *X 和 *Y 相互转换

X和Y是两个结构体：

```go
var p *X
var q *Y

q = (*Y)(unsafe.Pointer(p))		// *X => *Y
q = (*X)(unsafe.Pointer(q))		// *Y => *X
```



## []X 和 []Y 相互转换

X和Y是结构相同但名称不相同的`struct`：

<img src="../../images/os/array_X_Y_convert.png" alt="image-20210523114724213" style="zoom:50%;" />

- Go语言中切片，普通数据类型的切片，或结构体的切片，都有 `reflect.SliceHeader` 的指针，可以互转

```go
var p []X
var q []Y	// q = p

pHdr := (*reflect.SliceHeader)(unsafe.Pointer(&p))
qHdr := (*reflect.SliceHeader)(unsafe.Pointer(&q))

pHdr.Data = qHdr.Data
pHdr.Len = qHdr.Len * unsafe.Sizeof(q[0]) / unsafe.Sizeof(p[0])
pHdr.Cap = qHdr.Cap * unsafe.Sizeof(q[0]) / unsafe.Sizeof(p[0])
```

- 所有切片拥有相同的头部 `reflect.SliceHeader`
- 重新构造切片头部即可完成转换



## 示例：float64 数组排序优化

```go
func main() {
  // []float64 强制类型转换为 []int
  var a = []float64{4,2,5,7,2,1,88,1}
  var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
  
  // 以int方式给float64排序
  sort.Ints(b)
  
  // 再打印a的话，就是排序好了的a
}
```

- `float64` 遵循 IEEE754 浮点数标准特性
- 当浮点数有序时，对应的整数也必然是有序的



## Go调用C函数

### Example A

```go
/*
static int add(int a, int b) {
	return a+b;
}
*/
import "C"

func main() {
  C.add(1, 1)
}
```

- `C.add` 通过C虚拟包访问
- 最终会转为 `_Cfunc_add` 名字



### Example B

```go
/*
static int add(int a, int b) {
	return a+b;
}
*/
import "C"

func main() {
  v, err := C.add(1, 1)
  fmt.Println(v, err)
  
  // output:
  // 2 <nil>
}
```

- 任何C函数都可以带两个返回值

- 第二个返回值是 `errno`，对应`error`接口类型



### Example C

```go
/*
#include <errno.h>
static void seterrno(int v) {
	errno = v;
}
*/
import "C"
import "fmt"

func main() {
  _, err := C.seterrno(9527)
  fmt.Println(err)
  
  // output:
  // errno 9527
}
```

- 即使没有返回值，依然可以通过第二个返回值获取errno
- 对应voide类型函数，第一个返回值可以用_占位



### Example D

```go
// static void noreturn() {}

import "C"
import "fmt"

func main() {
  x, _ := C.noreturn()
  fmt.Printf("%#v\n", x)
  
  // output:
  // main._Ctype_void{}
}
```

- 甚至可以获取一个 void 类型函数的返回值
- 返回值类型：`type _Ctype_void [0]byte`



## 导出Go函数

```go
import "C"

//export GoAdd
func GoAdd(a, b C.int) C.int {
  return a + b
}
```

- 可以导出go的私有函数
- 导出C函数名没有名字空间约束，需保证全局没有重名
- main 包的导出函数会在 `_cgo_export.h` ，都是函数声明



不用Go导出的头文件，自己手写 `add.h` 也是可以的，这也是cgo常用的使用方式（默认的有些约束），例如：

add.h:

```c
int c_add(int a, int b);
```

add.c:

```c
#include "add.h"
#include "_cgo_export.h"

int c_add(int a, int b) {
  return GoAdd(a, b)
}
```

- 在C文件中使用 `_cgo_export.h` 头文件
- C文件必须在同一个包，否则会找不到头文件







## 参考资料

- [package unsafe](https://studygolang.com/pkgdoc)

- [深入cgo编程](https://www.bilibili.com/video/BV1rs411M75T?from=search&seid=3978510227066577408)