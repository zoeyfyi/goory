# goory
[![Build Status](https://travis-ci.org/bongo227/goory.svg?branch=travis)](https://travis-ci.org/bongo227/goory)
[![codecov](https://codecov.io/gh/bongo227/goory/branch/master/graph/badge.svg)](https://codecov.io/gh/bongo227/goory)
[![godocs](https://godoc.org/github.com/bongo227/goory?status.svg)](http://godoc.org/github.com/bongo227/goory)

A go package for compiling syntax trees to LLVM ir and WebAssembly ir

## Todos
* Instructions:
    * Terminator instructions:
        * ret ✔
        * br ✔
        * switch
        * indirectbr
        * invoke
        * resume
        * catchswitch
        * catchret
        * cleanupret
        * unreachable
    * Binary operations:
        * add ✔
        * fadd ✔
        * sub ✔
        * fsub ✔
        * mul ✔
        * fmul ✔
        * udiv
        * sdiv ~
        * fdiv ✔
        * urem
        * srem
        * frem
    * Bitwise binary operations:
        * shl
        * lshr
        * ashr
        * and
        * org
        * xor
    * Vector operations:
        * extractelement
        * insertelement
        * shufflevector
    * Aggregate operations:
        * extractvalue
        * insertvalue
    * Memory access and addressing operations:
        * alloca
        * load
        * store
        * fence
        * cmpxchg
        * atomicrmw
        * getelementptr
    * Conversion operations:
        * trunc .. to ~
        * zext .. to ~
        * sext .. to ~
        * fptrunc .. to ~
        * fpext .. to ~
        * fptoui .. to ~
        * fptosi .. to ~
        * uitofp .. to 
        * sitofp .. to
        * ptrtoint .. to
        * inttoptr .. to
        * bitcast .. to
        * addrspacecast .. to
    * Other operations:
        * icmp ✔
        * fcmp ✔
        * phi 
        * select 
        * call ✔
        * va_arg
        * landingpad
        * catchpad
        * cleanuppad
        