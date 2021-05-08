// pointer-like class, 关于智能指针

// px -> T
template<class T>
class shared_ptr
{
public:
	T& operator*() const
	{ return *px; }

	T* operator->() const
	{ return px; }

	shared_ptr(T* p) : px(p) {}

private:
	T*		px;
	long*	pn;
	...
};

struct Foo
{
	...
	void method(void) { ... }
};

{
	shared_ptr<Foo> sp(new Foo);

	Foo f(*sp);

	sp->method(); // 此句等同于：px->method(); 注意：此处的箭头符号作用下去以后（解析为->的重载），还可以继续使用！
}