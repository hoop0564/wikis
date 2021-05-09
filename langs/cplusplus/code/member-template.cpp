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