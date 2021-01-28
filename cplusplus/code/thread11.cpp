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
    std::thread t1((Fctor()), std::ref(s));
    t1.join();
    cout << "msg after thread: " << s << endl;
    return 0;
}