/**
 * 实现同步写文件的
 */
#include <iostream>
#include <mutex>
#include <fstream>
#include <string>

using namespace std;

class LogFile
{
private:
    std::mutex m_mutex;
    std::ofstream f;

public:
    LogFile();
    ~LogFile();
    void shared_print(string key, int val)
    {
        std::lock_guard<std::mutex> locker(m_mutex);
        f << "KEY: " << key << " VAL: " << val << endl;
    }
};

LogFile::LogFile()
{
    f.open("log.txt");
}

LogFile::~LogFile()
{
    f.close();
}
