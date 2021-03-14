// 类成员函数回调 

#include <iostream>
#include <functional>

using namespace std;
using namespace std::placeholders;

typedef std::function<void(int, int)> Fun;

class B {
public:
	void call(int a, Fun f)
	{
		f(a, 2);
	}
};

class Test {
public:
	void callback(int a, int b)
	{
		cout << a << "+" << b << "=" << a + b << endl;
	}

	void bind()
	{
		// bind函数返回一个新的函数对象。可以随便给子函数定几个参数，但是不能多于bind所绑定的原函数的参数个数
		// 或使用boost::bind
		Fun fun = std::bind(&Test::callback, this, _1, _2);
		B b;
		b.call(1, fun);
	}

};

int main()
{
	Test test;
	test.bind();
	return 0;
}

/*
* output:
* 1+2=3
*/
