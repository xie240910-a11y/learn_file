#include <iostream>

class Singleton{
private:
    Singleton() {
        std::cout << "Singleton create...";
    };
    ~Singleton(){};
    Singleton (const Singleton &) = delete;
    Singleton &operator=(const Singleton &) = delete;

    Singleton (Singleton &&) = delete;
    Singleton &operator=(Singleton &&) = delete;
public:
    static Singleton& getInstance(){
        static Singleton instance;
        return instance;
    }
}; 