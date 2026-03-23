#include <iostream>
using namespace std;

int main() {
    int* p = new int(10);  // 分配内存

    cout << *p << endl;

    // ❌ 忘记 delete
    return 0;
}