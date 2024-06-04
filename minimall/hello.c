//#include "hello.h"
// rust driver cgen file
extern "C"{
int add(int a, int b) {
    return a + b;
}
}
