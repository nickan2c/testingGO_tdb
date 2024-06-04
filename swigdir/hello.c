#include "hello.h"
// rust driver cgen file
int add(int a, int b) {
    return a + b;
}
// compile this to a .a or .o not a so.
// .a is closer