// [C++11]std::packaged_task介绍及使用
// https://blog.csdn.net/godmaycry/article/details/72868559

#include <iostream>
#include <future>
#include <chrono>
#include <functional>

using namespace std;

/*
当保存的数据类型是可调对象时（如函数、lambda表达式等），使用std::packaged_task比std::promise更简洁。语法糖。
*/

int Test_Fun(int a, int b, int &c) 
{
	std::this_thread::sleep_for(std::chrono::seconds(5));
	c = a + b + 230;
	return c;
}

int main()
{
	std::packaged_task<int(int, int, int&)> pt1(Test_Fun);
	std::future<int> fu1 = pt1.get_future();
	
	int c=0;
	std::thread t1(std::move(pt1), 1, 2, std::ref(c));

	int iResult = fu1.get();
	cout << "iResult:" << iResult << endl;
	cout << "c:" << c << endl;

	return 0;
}