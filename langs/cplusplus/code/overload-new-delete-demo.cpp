// .h
class Foo
{
public:
	int _id; 	// 4bytes
	long _data;	// 4bytes
	string _str;	// 4bytes Foo的对象占用12个字节！

public:
	Foo() : _id(0) {
		cout << "default ctor.this=" << this << " id=" << _id << endl;
	}
	Foo(int i): _id(i) {
		cout << "ctor.this=" << this << " id=" << _id << endl;		
	}
	// virtual
	~Foo() {
		cout << "dtor.this=" << this << " id=" << _id << endl;
	}

	static void* operator new(size_t size);
	static void operator delete(void* pdead, size_t size);
	static void* operator new[](size_t size);
	static void operator delete[](void* pdead, size_t size);
};


// .cpp
void* Foo::operator new(size_t size) {
	Foo* p = (Foo*)malloc(size);
	cout << ...;
	return p;
}

void Foo::operator delete(void* pdead, size_t size) {
	cout << ...;
	free(pdead);
}

void* Foo::operator new[](size_t size) {
	Foo* p = (Foo*)malloc(size);
	cout << ...;
	return p;
}

void Foo::operator delete[](void* pdead, size_t size) {
	cout << ...;
	free(pdead);
}
