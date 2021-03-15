package map_ext

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employ struct {
	Id   int
	Name string
}

// 第一种方法在实例方法被调用时，实例会被值拷贝一次
func (e Employ) GetString() string {
	fmt.Println("test pln")
	fmt.Printf("e name's address is %x\n", unsafe.Pointer(&e.Name)) // 发现指针地址和之前的对象实例不一样
	return fmt.Sprintf("Id=%d, Name=%s", e.Id, e.Name)
}

// 第二种方法避免了内存拷贝，通常使用
func (e *Employ) GetString2() string {
	fmt.Printf("e name's address is %x\n", unsafe.Pointer(&e.Name)) // 指针地址和之前的对象实例一样
	return fmt.Sprintf("Id=%d, Name=%s", e.Id, e.Name)
}

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d exist", n)
	} else {
		t.Logf("%d not exist", n)
	}

	mySet[3] = true

	delete(mySet, 1)
	t.Log(len(mySet))
}

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 初始化的零值为""

	s = "hello"
	// s[1] = '2' // string是不可变的byte slice

	s = "\xE4\xBA\xBB\xFF" // 可以储存任何二进制数据
	t.Log(len(s))          // 4

	s = "中"
	t.Log(len(s)) // 3 是byte数

	c := []rune(s)
	t.Log(len(c))                // 1
	t.Logf("中的unicode %x", c[0]) // 4e2d
	t.Logf("中的utf8 %x", s)       // e4b8ad
}

func TestStringToRun(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s { // range和字符串string配合迭代输出的是rune，而不是byte！
		t.Logf("%[1]c %[1]x %[1]d", c) // [1]都是打印第一个字符
	}
}

func TestStruct(t *testing.T) {
	e := Employ{1, "jack"}
	e1 := Employ{Id: 2, Name: "mike"}
	e2 := new(Employ)
	e2.Id = 3
	e2.Name = "Tom"
	t.Logf("e type is %T, e1 type is %T, e2 type is %T", e, e1, e2) // map_ext.Employ *map_ext.Employ
}
