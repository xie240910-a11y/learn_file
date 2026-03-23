// double free coredump
#include <iostream>
#include <thread>
#include <unistd.h>

using namespace std;

int *ptr = nullptr;

void threadFunc1() {
    cout << "线程1释放内存" << endl;
    free(ptr);
}

void threadFunc2() {
    cout << "线程2释放内存" << endl;
    free(ptr);
}
int main() {
    ptr = (int*)malloc(sizeof(int));
    *ptr = 10;

    thread t1(threadFunc1);
    thread t2(threadFunc2);

    t1.join();
    t2.join();

    return 0;
}