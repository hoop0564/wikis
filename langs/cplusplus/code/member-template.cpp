// member-template.cpp 成员模板
// 本身已经是模板的类，里面存在模板成员函数，就叫成员模板！

template <class T1, class T2>

struct pair
{
	typedef T1 first_type;
	typedef T2 second_type;

	T1 first;
	T2 second;

	pair() : first(T1()), second(T2()){}
	pair(const T1& a, const T2& b) : first(a), second(b) {}

	template <class U1, class U2>
	pair(const pair<U1, U2>& p) : first(p.first), second(p.second){};
};


// 实际案例，标准库中共有很多此设计手法
{
	class Base1 {}; // 鱼类
	class Derived1: public Base1 {}; // 鲫鱼

	class Base2 {}; // 鸟类
	class Derived2: public Base2 {}; // 麻雀

	{
		pair<Derived1, Derived2> p; // T1, T2
		pair<Base1, Base2> p2(p); 
		// 上句等价于:
		pair<Base1, Base2> p2(pair<Derived1, Derived2>());
	}

}


{
	template <typename _Tp>
	class shared_ptr: public __shared_ptr<_Tp>
	{
		...
		template<typename _Tp1>
		explicit shared_ptr(_Tp1* __p) : __shared_ptr<_Tp1>(__p) {}
		...
	}

	{
		Base1* ptr = new Derived1; // up-cast
		shared_ptr<Base1> sptr(new Derived1); // 模拟up-cast, _Tp是Base1，_Tp1是Derived1
	}
}
