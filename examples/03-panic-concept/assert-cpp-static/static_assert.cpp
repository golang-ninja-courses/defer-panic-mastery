#include <type_traits>

template <class T>
void swap(T& a, T& b) {
    static_assert(std::is_copy_constructible<T>::value, "swap requires copying");
    auto c = b;
    b = a;
    a = c;
}

struct no_copy {
    no_copy(const no_copy&) = delete;
    no_copy() = default;
};

// g++ -std=c++17 static_assert.cpp
int main() {
    int a, b;
    swap(a, b);

    no_copy nc_a, nc_b;
    swap(nc_a, nc_b);
}