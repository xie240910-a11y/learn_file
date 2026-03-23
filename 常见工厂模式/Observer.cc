#include <iostream>
#include <vector>
#include <algorithm>
#include <memory>
using namespace std;
class Observer{
public:
    virtual void update(int value) = 0;
    virtual ~Observer() {}
};

class ConcreteObserver:public Observer{
public:
    void update(int value){
        cout << "收到更新：" << value << endl;
    }
};

class Subject{
private:
    vector<weak_ptr<Observer>> observers;
    int state;
public:
    void attach(weak_ptr<Observer> obs) {
        observers.push_back(obs);
    }
    void setState(int value){
        state = value;
        notify();
    }
    void notify() {
        for(auto it = observers.begin(); it != observers.end();) {
            if(auto obs = it->lock()){
                obs->update(state);
                it++;
            }
            else {
                observers.erase(it);
            }
        }
    }
    void detach(shared_ptr<Observer> obs) {
        observers.erase(
            remove_if(observers.begin(), observers.end(),
                [&](weak_ptr<Observer>& wptr) {
                    return wptr.lock() == obs;
                }),
            observers.end()
        );
    }

};

int main(int argc, char const *argv[])
{
    Subject sub;
    shared_ptr<ConcreteObserver> c1 = make_shared<ConcreteObserver>();
    shared_ptr<ConcreteObserver> c2 = make_shared<ConcreteObserver>();
    sub.attach(c1);
    sub.attach(c2);
    sub.setState(10);

    return 0;
}
