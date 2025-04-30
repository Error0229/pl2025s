#include <iostream>
using namespace std;

template <typename T>
int bar(T f)
{
    int x = 5;
    return f(x);
}

int main()
{
    int x = 20;
    cout << bar([&x](int y) { return x++ + y; }) << endl;
    cout << x << endl;
}
