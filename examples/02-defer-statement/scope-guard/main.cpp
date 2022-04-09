#include "scope_guard.h"
#include <iostream>

using std::cout;
using std::endl;

// cd examples/02-defer-statement/scope-guard
// g++ -std=c++17 main.cpp && ./a.out

int main() {
  DEFER ( cout << "1" << endl );

  {
    DEFER ( cout << "2" << endl );
  }

  DEFER ( cout << "3" << endl );
}

/*
2
3
1
*/
