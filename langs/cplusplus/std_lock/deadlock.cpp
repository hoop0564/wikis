#include <iostream>

using namespace std;

// 会出现死锁
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

// 解决死锁方案1
// 先把两个锁都申请到，后面用lock_guard来管理锁的生命周期，adopt_lock表示当前线程已经获得锁，无需再上锁
void deadLockSolution1(mutex &mtA, mutex &mtB)
{
    std::lock(mtA, mtB);
    std::lock_guard<mutex> lock1(mtA, std::adopt_lock);
    cout << "get first mutext A in thread" << this_thread::get_id() << endl;
    this_thread::sleep_for(std::chrono::milliseconds(1));

    std::lock_guard<mutex> lock2(mtB, std::adopt_lock);
    cout << "get first mutext B in thread" << this_thread::get_id() << endl;
    this_thread::sleep_for(std::chrono::milliseconds(1));

    cout << "do sth in thread " << this_thread::get_id() << endl;
}

// 解决死锁方案2
// 用unique_lock构建锁，用defer_lock先不上锁，执行功能代码时在统一lock上锁
// 用unique_lock生命周期来自动解锁mutex
void deadLockSolution2(mutex &mtA, mutex &mtB)
{
    std::unique_lock<std::mutex> lock1(mtA, std::defer_lock);
    cout << "get first mutext A in thread" << this_thread::get_id() << endl;
    this_thread::sleep_for(std::chrono::milliseconds(1));

    std::unique_lock<std::mutex> lock2(mtB, std::defer_lock);
    cout << "get first mutext B in thread" << this_thread::get_id() << endl;
    this_thread::sleep_for(std::chrono::milliseconds(1));

    std::lock(lock1, lock2);
    assert(lock1.owns_lock() == true);

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