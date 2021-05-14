// 重载memebr operator new/delete

class Foo {
public: // per-class allocator
	void* operator new(size_t);
	void operator delete(void*, size_t); // size_t is optional
	...
}

{
	Foo* p = new Foo;
	...
	delete p;

	// 第11行等同于：
	{
		// 先执行类重载的new，分配内存
		void* mem = operator new(sizeof(Foo));
		p = static_cast<Foo*>(mem);
		// 构造
		p->Foo::Foo();
	}

	// 第13行等同于：
	{
		p->~Foo();
		operator delete(p);
	}
}