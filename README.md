# goory
[![Build Status](https://travis-ci.org/bongo227/goory.svg?branch=travis)](https://travis-ci.org/bongo227/goory)
[![codecov](https://codecov.io/gh/bongo227/goory/branch/master/graph/badge.svg)](https://codecov.io/gh/bongo227/goory)
[![godocs](https://godoc.org/github.com/bongo227/goory?status.svg)](http://godoc.org/github.com/bongo227/goory)

A go package for compiling syntax trees to LLVM ir and WebAssembly ir

## Example usage
### Adding to floats
```go
double := goory.DoubleType()

module := goory.NewModule("test")

function := module.NewFunction("fpadd", double)
a := function.AddArgument(double, "a")
b := function.AddArgument(double, "b")

block := function.Entry()

result := block.Fadd(a, b)

block.Ret(result)
```

### Gcd algorithum
```go
integer := goory.IntType(32)

module := goory.NewModule("test")

function := module.NewFunction("gcd", integer)
x := function.AddArgument(integer, "x")
y := function.AddArgument(integer, "y")

block := function.Entry()

xyEqual := block.Icmp(goory.IntEq, x, y)

trueBlock := function.AddBlock()
trueBlock.Ret(x)

falseBlock := function.AddBlock()
xyLess := falseBlock.Icmp(goory.IntUlt, x, y)

elseIfBlock := function.AddBlock()
call1 := elseIfBlock.Call(function, x, elseIfBlock.Sub(y, x))
elseIfBlock.Ret(call1)

elseBlock := function.AddBlock()
call2 := elseBlock.Call(function, elseBlock.Sub(x, y), y)
elseBlock.Ret(call2)

block.CondBr(xyEqual, trueBlock, falseBlock)
falseBlock.CondBr(xyLess, elseIfBlock, elseBlock)
```

### Fibonaci numbers
```go
integer := goory.IntType(32)
module := goory.NewModule("test")

function := module.NewFunction("fib", integer)
n := function.AddArgument(integer, "n")

block := function.Entry()

trueBlock := function.AddBlock()
nEqualsZero := block.Icmp(goory.IntEq, n, goory.Constant(integer, 0))
nEqualsOne := block.Icmp(goory.IntEq, n, goory.Constant(integer, 1))
nOneOrZero := block.Or(nEqualsZero, nEqualsOne)
trueBlock.Ret(n)

falseBlock := function.AddBlock()
fib1 := falseBlock.Call(function, falseBlock.Sub(n, goory.Constant(integer, 1)))
fib2 := falseBlock.Call(function, falseBlock.Sub(n, goory.Constant(integer, 2)))
falseBlock.Ret(falseBlock.Add(fib1, fib2))

block.CondBr(nOneOrZero, trueBlock, falseBlock)
```

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
        * and ✔
        * org
        * xor ✔
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
        * trunc .. to ✔
        * zext .. to ✔
        * sext .. to ✔
        * fptrunc .. to ✔
        * fpext .. to ✔
        * fptoui .. to ✔
        * fptosi .. to ✔
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
        