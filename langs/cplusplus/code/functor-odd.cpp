// 标准库中的仿函数的奇特模样

template <class T>
struct plus : public binary_function<T, T, T> // binary: 二元的
{
	T operator() (const T& x, const T& y) const {return x + y; }
};

template <class T>
struct minus : public binary_function<T, T, T>
{
	T operator() (const T&x, const T& y) const { return x - y; }
};

template <class T>
struct equal_to : public binary_function<T, T, bool>
{
	bool operator() (const T&x, const T& y) const {return x == y; }
};

template <class T>
struct less : public binary_function<T, T, bool>
{
	bool operator() (const T&x, const T& y) const {return x < y; }
};


// 标准库中，仿函数所使用的奇特的base classes
template <class Arg, class Result>
struct unary_function
{
	typedef Arg argument_type;
	typedef Result result_type;	
};

template <class Arg1, class Arg2, class Result>
struct unary_function
{
	typedef Arg1 first_argument_type;
	typedef Arg2 second_argument_type;
	typedef Result result_type;	
};
