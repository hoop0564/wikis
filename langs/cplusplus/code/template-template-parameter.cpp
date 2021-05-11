// 模板模板参数
template<typename T, // 此处不用typename而用class也行，前者是后出现的关键词，后者是一开始就用的；模板也是c++后来才有的
		 template <typename T>
		 	class Container
		>
class XCLs
{
private:
	Container<T> c;
public:
	...
};

{
	template<typename T>
	using Lst = list<T, allocator<T>>;

//	XCLs<string, list> mylst1; // 错误语句！
	XCLs<string, Lst> mylst2;  // 正确
}