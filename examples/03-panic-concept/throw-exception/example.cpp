#include <iostream>
#include <stdexcept>

int Div(int a, int b) {
    if (b == 0) {
        throw std::invalid_argument("division by zero");
    }
    return a / b;
}

// g++ example.cpp && ./a.out
int main()
{
    std::cout << "Div(42, 0) = " << Div(42, 0) << std::endl;
    return 0;
}
