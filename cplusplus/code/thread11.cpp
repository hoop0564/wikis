#include <string>
#include <iostream>
#include <thread>

using namespace std;

void function_1(string &msg)
{
    cout << "msg in thread: " << msg << endl;
    msg = "world";
}

class Fctor
{
public:
    void operator()(string &msg)
    {
        cout << "class msg in thread: " << msg << endl;
        msg = "world!!";
    }
};

int main()
{
    string s = "hello";
    // std::thread t1(function_1, std::ref(s));
    // std::thread t1((Fctor()), s); //仍然是值传递！
    std::thread t1((Fctor()), std::ref(s)); //是引用传递
    // std::thread t1((Fctor()), std::move(s)); //移动语义，主线程的s被掏空，内存被转移到t1线程中
    // std::thread t2 = std::move(t1);
    t1.join();
    cout << "msg after thread: " << s << endl;
    return 0;
}