#include <iostream>
using namespace std;

//Myvector的类模板
template <typename Ty>
class MyVector
{
public:
    //内嵌类型表
    typedef Ty value;
    typedef Ty *viter;

public:
    MyVector(int nLen = 0) : m_nLen(nLen), m_Data(NULL), finish(0)
    {
        if (nLen > 0)
        {
            m_Data = new Ty[nLen];
            start = m_Data;
            end_of_element = nLen;
        }
    }
    ~MyVector()
    {
        delete[] m_Data;
    }

    void push_back(const value &x)
    {
        if (end_of_element != finish)
        {
            *(start + finish) = x;
            ++finish;
        }
        else
        {
            cout << "out of boudary" << endl;
        }
    }

    inline value pop_back()
    {
        --finish;
        return *(start + finish);
    }

    value &operator[](int n)
    {
        if (n > 0 || n <= finish)
        {
            return *(start + n);
        }
        else
        {
            cout << "取值错误" << endl;
        }
    }

protected:
    viter m_Data;       //数组头指针
    int m_nLen;         //数组长度
    viter start;        //数组的起始地址
    int finish;         //数组的满位标志
    int end_of_element; //数组的末尾标识
};

int main()
{
    int x;
    MyVector<int> vec(10);
    vec.push_back(100);
    vec.push_back(200);
    vec.push_back(300);
    x = vec.pop_back();
    cout << "x=" << x << endl;

    cout << vec[0] << endl;
    cout << vec[1] << endl;

    return 0;
}