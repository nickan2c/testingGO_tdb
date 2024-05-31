%module hello

// Make mylib_wrap.cxx include this header:
%{
#include "hello.h"
%}

// Make SWIG look into this header:
%include "hello.h"