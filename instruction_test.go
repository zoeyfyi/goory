package goory_test

import (
	"testing"

	. "github.com/bongo227/goory"
	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	Convey("Instruction", t, func() {
		Convey("add", func() {
			module := NewModule("test_add_module")
			a, b := Argument(NewIntType(32), "a"), Argument(NewIntType(32), "b")
			function := module.NewFunction("test_add_function", NewBoolType(), a, b)
			instruction := function.Entry().Add(a, b)

			Convey("should have string value 'add'", func() { So(instruction.String(), ShouldEqual, "add") })
			Convey("should not be a terminator", func() { So(instruction.IsTerminator(), ShouldEqual, false) })
			Convey("should be of type int32", func() { So(instruction.Type(), ShouldResemble, NewIntType(32)) })
			Convey("should produce the correct llvm ir", func() { So(module.LLVM(), ShouldContainSubstring, "%temp0 = add i32 %a, %b") })
		})

		Convey("sub", func() {
			module := NewModule("test_sub_module")
			a, b := Argument(NewIntType(32), "a"), Argument(NewIntType(32), "b")
			function := module.NewFunction("test_sub_function", NewBoolType(), a, b)
			instruction := function.Entry().Sub(a, b)

			Convey("should have string value 'sub'", func() { So(instruction.String(), ShouldEqual, "sub") })
			Convey("should not be a terminator", func() { So(instruction.IsTerminator(), ShouldEqual, false) })
			Convey("should be of type int32", func() { So(instruction.Type(), ShouldResemble, NewIntType(32)) })
			Convey("should produce the correct llvm ir", func() { So(module.LLVM(), ShouldContainSubstring, "%temp0 = sub i32 %a, %b") })
		})

		Convey("mul", func() {
			module := NewModule("test_mul_module")
			a, b := Argument(NewIntType(32), "a"), Argument(NewIntType(32), "b")
			function := module.NewFunction("test_mul_function", NewBoolType(), a, b)
			instruction := function.Entry().Mul(a, b)

			Convey("should have string value 'mul'", func() { So(instruction.String(), ShouldEqual, "mul") })
			Convey("should not be a terminator", func() { So(instruction.IsTerminator(), ShouldEqual, false) })
			Convey("should be of type int32", func() { So(instruction.Type(), ShouldResemble, NewIntType(32)) })
			Convey("should produce the correct llvm ir", func() { So(module.LLVM(), ShouldContainSubstring, "%temp0 = mul i32 %a, %b") })
		})

		Convey("div", func() {
			module := NewModule("test_div_module")
			a, b := Argument(NewIntType(32), "a"), Argument(NewIntType(32), "b")
			function := module.NewFunction("test_div_function", NewBoolType(), a, b)
			instruction := function.Entry().Div(a, b)

			Convey("should have string value 'div'", func() { So(instruction.String(), ShouldEqual, "div") })
			Convey("should not be a terminator", func() { So(instruction.IsTerminator(), ShouldEqual, false) })
			Convey("should be of type int32", func() { So(instruction.Type(), ShouldResemble, NewIntType(32)) })
			Convey("should produce the correct llvm ir", func() { So(module.LLVM(), ShouldContainSubstring, "%temp0 = div i32 %a, %b") })
		})

		Convey("fadd", func() {
			module := NewModule("test_fadd_module")
			a, b := Argument(NewFloatType(), "a"), Argument(NewFloatType(), "b")
			function := module.NewFunction("test_fadd_function", NewBoolType(), a, b)
			instruction := function.Entry().Fadd(a, b)

			Convey("should have string value 'fadd'", func() { So(instruction.String(), ShouldEqual, "fadd") })
			Convey("should not be a terminator", func() { So(instruction.IsTerminator(), ShouldEqual, false) })
			Convey("should be of type float", func() { So(instruction.Type(), ShouldResemble, NewFloatType()) })
			Convey("should produce the correct llvm ir", func() { So(module.LLVM(), ShouldContainSubstring, "%temp0 = fadd float %a, %b") })
		})

		Convey("fsub", func() {
			module := NewModule("test_fsub_module")
			a, b := Argument(NewFloatType(), "a"), Argument(NewFloatType(), "b")
			function := module.NewFunction("test_fsub_function", NewBoolType(), a, b)
			instruction := function.Entry().Fsub(a, b)

			Convey("should have string value 'fsub'", func() { So(instruction.String(), ShouldEqual, "fsub") })
			Convey("should not be a terminator", func() { So(instruction.IsTerminator(), ShouldEqual, false) })
			Convey("should be of type float", func() { So(instruction.Type(), ShouldResemble, NewFloatType()) })
			Convey("should produce the correct llvm ir", func() { So(module.LLVM(), ShouldContainSubstring, "%temp0 = fsub float %a, %b") })
		})

		Convey("fmul", func() {
			module := NewModule("test_fmul_module")
			a, b := Argument(NewFloatType(), "a"), Argument(NewFloatType(), "b")
			function := module.NewFunction("test_fmul_function", NewBoolType(), a, b)
			instruction := function.Entry().Fmul(a, b)

			Convey("should have string value 'fmul'", func() { So(instruction.String(), ShouldEqual, "fmul") })
			Convey("should not be a terminator", func() { So(instruction.IsTerminator(), ShouldEqual, false) })
			Convey("should be of type float", func() { So(instruction.Type(), ShouldResemble, NewFloatType()) })
			Convey("should produce the correct llvm ir", func() { So(module.LLVM(), ShouldContainSubstring, "%temp0 = fmul float %a, %b") })
		})

		Convey("fdiv", func() {
			module := NewModule("test_fdiv_module")
			a, b := Argument(NewFloatType(), "a"), Argument(NewFloatType(), "b")
			function := module.NewFunction("test_fdiv_function", NewBoolType(), a, b)
			instruction := function.Entry().Fdiv(a, b)

			Convey("should have string value 'fdiv'", func() { So(instruction.String(), ShouldEqual, "fdiv") })
			Convey("should not be a terminator", func() { So(instruction.IsTerminator(), ShouldEqual, false) })
			Convey("should be of type float", func() { So(instruction.Type(), ShouldResemble, NewFloatType()) })
			Convey("should produce the correct llvm ir", func() { So(module.LLVM(), ShouldContainSubstring, "%temp0 = fdiv float %a, %b") })
		})

		Convey("fcmp", func() {
			module := NewModule("test_fcmp_module")
			a, b := Argument(NewFloatType(), "a"), Argument(NewFloatType(), "b")
			function := module.NewFunction("test_fcmp_function", NewBoolType(), a, b)
			instruction := function.Entry().FCmp(FloatOeq, a, b)

			Convey("should have string value 'fcmp'", func() { So(instruction.String(), ShouldEqual, "fcmp") })
			Convey("should not be a terminator", func() { So(instruction.IsTerminator(), ShouldEqual, false) })
			Convey("should be of type bool", func() { So(instruction.Type(), ShouldResemble, NewBoolType()) })
			Convey("should produce the correct llvm ir", func() { So(module.LLVM(), ShouldContainSubstring, "%temp0 = fcmp oeq bool %a, %b") })
		})

		Convey("icmp", func() {
			module := NewModule("test_icmp_module")
			a, b := Argument(NewFloatType(), "a"), Argument(NewFloatType(), "b")
			function := module.NewFunction("test_icmp_function", NewBoolType(), a, b)
			instruction := function.Entry().ICmp(IntEq, a, b)

			Convey("should have string value 'icmp'", func() { So(instruction.String(), ShouldEqual, "icmp") })
			Convey("should not be a terminator", func() { So(instruction.IsTerminator(), ShouldEqual, false) })
			Convey("should be of type bool", func() { So(instruction.Type(), ShouldResemble, NewBoolType()) })
			Convey("should produce the correct llvm ir", func() { So(module.LLVM(), ShouldContainSubstring, "%temp0 = icmp eq bool %a, %b") })
		})

		Convey("br", func() {
			module := NewModule("test_br_module")
			function := module.NewFunction("test_br_function", NewBoolType())
			testBlock := function.AddBlock()
			instruction := function.Entry().Br(testBlock)

			Convey("should have string value 'br'", func() { So(instruction.String(), ShouldEqual, "br") })
			Convey("should be a terminator", func() { So(instruction.IsTerminator(), ShouldEqual, true) })
			Convey("should be of type void", func() { So(instruction.Type(), ShouldResemble, NewVoidType()) })
			Convey("should produce the correct llvm ir", func() { So(module.LLVM(), ShouldContainSubstring, "br label %temp0") })
		})

		Convey("conditional br", func() {
			module := NewModule("test_br_module")
			function := module.NewFunction("test_br_function", NewBoolType())
			trueBlock := function.AddBlock()
			falseBlock := function.AddBlock()
			condition := Constant(NewBoolType(), false)
			instruction := function.Entry().CondBr(condition, trueBlock, falseBlock)

			Convey("should have string value 'br'", func() { So(instruction.String(), ShouldEqual, "br") })
			Convey("should be a terminator", func() { So(instruction.IsTerminator(), ShouldEqual, true) })
			Convey("should be of type void", func() { So(instruction.Type(), ShouldResemble, NewVoidType()) })
			Convey("should produce the correct llvm ir", func() { So(module.LLVM(), ShouldContainSubstring, "br i1 false, label %temp0, label %temp1") })
		})

		Convey("call", func() {
			module := NewModule("test_call_module")
			toCall := module.NewFunction("test", NewBoolType(), Argument(NewIntType(32), "a"))
			function := module.NewFunction("test_call_function", NewBoolType())
			instruction := function.Entry().Call(toCall, Constant(NewIntType(32), 14))

			Convey("should have string value 'call'", func() { So(instruction.String(), ShouldEqual, "call") })
			Convey("should not be a terminator", func() { So(instruction.IsTerminator(), ShouldEqual, false) })
			Convey("should have type the same as the function return type", func() { So(instruction.Type(), ShouldResemble, NewBoolType()) })
			Convey("should produce the correct llvm ir", func() { So(module.LLVM(), ShouldContainSubstring, "%temp0 = call bool @test(i32 14)") })
		})

		Convey("ret", func() {
			module := NewModule("test_ret_module")
			a := Argument(NewIntType(32), "a")
			function := module.NewFunction("test_ret_function", NewIntType(32), a)
			instruction := function.Entry().Ret(a)

			Convey("should have string value 'ret'", func() { So(instruction.String(), ShouldEqual, "ret") })
			Convey("should be a terminator", func() { So(instruction.IsTerminator(), ShouldEqual, true) })
			Convey("should be of type void", func() { So(instruction.Type(), ShouldResemble, NewVoidType()) })
			Convey("should produce the correct llvm ir", func() { So(module.LLVM(), ShouldContainSubstring, "ret i32 %a") })
		})

	})
}
