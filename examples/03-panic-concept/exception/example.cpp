#include <iostream>

int Div(int a, int b) {
    return a / b;
}

// g++ example.cpp && ./a.out
int main()
{
    std::cout << "Div(42, 0) = " << Div(42, 0) << std::endl;
    return 0;
}
