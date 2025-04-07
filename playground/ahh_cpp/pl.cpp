#include <iostream>
using namespace std;
int Add(int a, int b)
{
    return a + b;
}
float Add(float a, float b)
{
    return a + b;
}

template <typename T, typename F>
T GenericAdd(T a, F b)
{
    return a + b;
}

int main()
{
    int a, c;
    float b;
    cout << GenericAdd(a, b) << '\n';
}
