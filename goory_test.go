package goory_test

import (
	"fmt"
	"testing"

	"github.com/bongo227/goory"
)

func TestExampleFpadd(t *testing.T) {
	double := goory.DoubleType()
	module := goory.NewModule("test")

	// Function: fpadd
	{
		function := module.NewFunction("fpadd", double)
		a := function.AddArgument(double, "a")
		b := function.AddArgument(double, "b")

		fmt.Println(function.Type().String())

		block := function.Entry()

		result := block.Fadd(a, b)
		block.Ret(result)
	}

	got := module.LLVM()
	expected := `define double @fpadd(double %a, double %b){
	entry:
		%temp0 = fadd double %a, %b
		ret double %temp0
}`

	if got != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, got)
	}
}

func TestExampleMulAdd(t *testing.T) {
	integer := goory.IntType(32)
	module := goory.NewModule("test")

	// Function: mul_add
	{
		function := module.NewFunction("mul_add", integer)
		x := function.AddArgument(integer, "x")
		y := function.AddArgument(integer, "y")
		z := function.AddArgument(integer, "z")

		block := function.Entry()

		result := block.Add(block.Mul(x, y), z)
		block.Ret(result)
	}

	got := module.LLVM()
	expected := `define i32 @mul_add(i32 %x, i32 %y, i32 %z){
	entry:
		%temp0 = mul i32 %x, %y
		%temp1 = add i32 %temp0, %z
		ret i32 %temp1
}`

	if got != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, got)
	}
}

func TestExampleGcd(t *testing.T) {
	integer := goory.IntType(32)
	module := goory.NewModule("test")

	// Function: gcd
	{
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
	}

	got := module.LLVM()
	expected := `define i32 @gcd(i32 %x, i32 %y){
	entry:
		%temp0 = icmp eq i32 %x, %y
		br i1 %temp0, label %temp1, label %temp2
	temp1:
		ret i32 %x
	temp2:
		%temp3 = icmp ult i32 %x, %y
		br i1 %temp3, label %temp4, label %temp7
	temp4:
		%temp5 = sub i32 %y, %x
		%temp6 = call i32 @gcd(i32 %x, i32 %temp5)
		ret i32 %temp6
	temp7:
		%temp8 = sub i32 %x, %y
		%temp9 = call i32 @gcd(i32 %temp8, i32 %y)
		ret i32 %temp9
}`

	if got != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, got)
	}

}

func TestExampleAlloc(t *testing.T) {
	module := goory.NewModule("test")
	integer := goory.IntType(32)
	{
		function := module.NewFunction("allocFun", integer)

		block := function.Entry()

		newInt := block.Alloca(integer)
		block.Store(newInt, goory.Constant(integer, 100))
		block.Ret(block.Load(newInt))
	}

	got := module.LLVM()
	expected := `define i32 @allocFun(){
	entry:
		%temp0 = alloca i32
		store i32 100, i32* %temp0
		%temp1 = load i32, i32* %temp0
		ret i32 %temp1
}`

	if got != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, got)
	}
}

func TestExampleFibonaci(t *testing.T) {
	integer := goory.IntType(32)
	module := goory.NewModule("test")

	// Function: fibonaci
	{
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
	}

	got := module.LLVM()
	expected := `define i32 @fib(i32 %n){
	entry:
		%temp1 = icmp eq i32 %n, 0
		%temp2 = icmp eq i32 %n, 1
		%temp3 = or bool %temp1 %temp2
		br i1 %temp3, label %temp0, label %temp4
	temp0:
		ret i32 %n
	temp4:
		%temp5 = sub i32 %n, 1
		%temp6 = call i32 @fib(i32 %temp5)
		%temp7 = sub i32 %n, 2
		%temp8 = call i32 @fib(i32 %temp7)
		%temp9 = add i32 %temp6, %temp8
		ret i32 %temp9
}`

	if got != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, got)
	}
}

func TestExampleBubbleSort(t *testing.T) {
	integer := goory.IntType(32)
	intArr := goory.ArrayType(integer, 5)
	module := goory.NewModule("test")
	zero := goory.Constant(integer, 0)

	// Function: bubbleSort
	{
		function := module.NewFunction("bubbleSort", intArr)
		items := function.AddArgument(intArr, "items")

		// entry block
		block := function.Entry()
		i := block.Alloca(integer)
		j := block.Alloca(integer)
		temp := block.Alloca(integer)
		arr := block.Alloca(intArr)
		block.Store(i, zero)
		block.Store(j, zero)
		block.Store(temp, zero)
		block.Store(arr, items)

		// Outer loop check
		outerCheck := function.AddBlock()
		outerResult := outerCheck.Icmp(goory.IntSlt, outerCheck.Load(i), goory.Constant(integer, 4))

		// inner loop check
		innerCheck := function.AddBlock()
		cap := innerCheck.Sub(goory.Constant(integer, 4), innerCheck.Load(i))
		innerResult := innerCheck.Icmp(goory.IntSlt, innerCheck.Load(j), cap)

		// swap check
		swapCheck := function.AddBlock()
		a := swapCheck.Getelementptr(integer, arr, zero, swapCheck.Load(j))
		b := swapCheck.Getelementptr(integer, arr, zero, swapCheck.Add(swapCheck.Load(j), goory.Constant(integer, 1)))
		swapresult := swapCheck.Icmp(goory.IntSgt, swapCheck.Load(a), swapCheck.Load(b))

		outerEnd := function.AddBlock()
		outerEnd.Store(j, zero)
		outerEnd.Store(i, outerEnd.Add(outerEnd.Load(i), goory.Constant(integer, 1)))

		// swap
		swap := function.AddBlock()
		swapA := swap.Getelementptr(integer, arr, zero, swap.Load(j))
		swapB := swap.Getelementptr(integer, arr, zero, swap.Add(swap.Load(j), goory.Constant(integer, 1)))
		swap.Store(temp, swap.Load(swapA))
		swap.Store(swapA, swap.Load(swapB))
		swap.Store(swapB, swap.Load(temp))

		// innerEnd
		innerEnd := function.AddBlock()
		innerEnd.Store(j, innerEnd.Add(innerEnd.Load(j), goory.Constant(integer, 1)))

		// endBlock
		endBlock := function.AddBlock()

		block.Br(outerCheck)
		outerCheck.CondBr(outerResult, innerCheck, endBlock)
		innerCheck.CondBr(innerResult, swapCheck, outerEnd)
		swapCheck.CondBr(swapresult, swap, innerEnd)
		swap.Br(innerEnd)
		innerEnd.Br(innerCheck)
		outerEnd.Br(outerCheck)
		endBlock.Ret(endBlock.Load(arr))
	}

	got := module.LLVM()
	expected := `define [5 x i32] @bubbleSort([5 x i32] %items){
	entry:
		%temp0 = alloca i32
		%temp1 = alloca i32
		%temp2 = alloca i32
		%temp3 = alloca [5 x i32]
		store i32 0, i32* %temp0
		store i32 0, i32* %temp1
		store i32 0, i32* %temp2
		store [5 x i32] %items, [5 x i32]* %temp3
		br label %temp4
	temp4:
		%temp5 = load i32, i32* %temp0
		%temp6 = icmp slt i32 %temp5, 4
		br i1 %temp6, label %temp7, label %temp36
	temp7:
		%temp8 = load i32, i32* %temp0
		%temp9 = sub i32 4, %temp8
		%temp10 = load i32, i32* %temp1
		%temp11 = icmp slt i32 %temp10, %temp9
		br i1 %temp11, label %temp12, label %temp21
	temp12:
		%temp13 = load i32, i32* %temp1
		%temp14 = getelementptr [5 x i32], [5 x i32]* %temp3, i32 0, i32 %temp13
		%temp15 = load i32, i32* %temp1
		%temp16 = add i32 %temp15, 1
		%temp17 = getelementptr [5 x i32], [5 x i32]* %temp3, i32 0, i32 %temp16
		%temp18 = load i32, i32* %temp14
		%temp19 = load i32, i32* %temp17
		%temp20 = icmp sgt i32 %temp18, %temp19
		br i1 %temp20, label %temp24, label %temp33
	temp21:
		store i32 0, i32* %temp1
		%temp22 = load i32, i32* %temp0
		%temp23 = add i32 %temp22, 1
		store i32 %temp23, i32* %temp0
		br label %temp4
	temp24:
		%temp25 = load i32, i32* %temp1
		%temp26 = getelementptr [5 x i32], [5 x i32]* %temp3, i32 0, i32 %temp25
		%temp27 = load i32, i32* %temp1
		%temp28 = add i32 %temp27, 1
		%temp29 = getelementptr [5 x i32], [5 x i32]* %temp3, i32 0, i32 %temp28
		%temp30 = load i32, i32* %temp26
		store i32 %temp30, i32* %temp2
		%temp31 = load i32, i32* %temp29
		store i32 %temp31, i32* %temp26
		%temp32 = load i32, i32* %temp2
		store i32 %temp32, i32* %temp29
		br label %temp33
	temp33:
		%temp34 = load i32, i32* %temp1
		%temp35 = add i32 %temp34, 1
		store i32 %temp35, i32* %temp1
		br label %temp7
	temp36:
		%temp44 = load [5 x i32], [5 x i32]* %temp3
		ret [5 x i32] %temp44
}`

	if got != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, got)
	}
}
