#include <iostream>

using namespace std;

void deadLock(mutex &mtA, mutex &mtB)
{
    lock_guard<mutex> lock1(mtA);
    cout << "get first mutext A in thread" << this_thread::get_id() << endl;
    this_thread::sleep_for(std::chrono::milliseconds(1));

    lock_guard<mutex> lock2(mtB);
    cout << "get first mutext B in thread" << this_thread::get_id() << endl;
    this_thread::sleep_for(std::chrono::milliseconds(1));

    cout << "do sth in thread " << this_thread::get_id() << endl;
}

int main()
{
    mutex mt1, mt2;
    std::thread t1([&] { deadLock(mt1, mt2) });
    std::thread t2([&] { deadLock(mt2, mt2) });

    t1.join();
    t2.join();
}