// variadic template 数量不定的模板参数

void print()
{

}

template <typename T, typename... Types>
void print (const T& firstArg, const Types&... args)
{
	cout << firstArg << endl;
	print(args...);
}

// usage
{
	print(7.5, "hello", bitset<16>(377), 42);
}

/* output
7.5
hello
0000000101111001 // bitset<>需要重载操作符<<
42
*/