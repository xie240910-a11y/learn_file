// 空指针非法访问coredump
#include <iostream>
using namespace std;

void crash() {
    int *p = nullptr;
    *p = 100;   
}

void func2() {
    crash();
}

void func1() {
    func2();
}
int main(int argc, char const *argv[])
{
    cout << "程序开始运行..." << endl;
    func1();
    return 0;
}