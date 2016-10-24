package goory

import (
	"fmt"
	"reflect"
	"testing"
)

func TestModule(t *testing.T) {
	m := NewModule("testing")

	// Check module name
	if m.Name() != "testing" {
		t.Errorf("Expected module name: %s\nGot: %s", "testing", m.Name())
	}
}

func TestFunction(t *testing.T) {
	m := NewModule("testing")
	f := m.NewFunction("test", Int32Type, Int32Type, Int32Type)

	// Check module
	if !reflect.DeepEqual(f.Module(), m) {
		t.Errorf("Function is not in parent module")
	}

	// Check function name
	if f.Name() != "test" {
		t.Errorf("Expected function name: %s\nGot: %s", "test", f.Name())
	}

	returnType, argTypes := f.Type()
	// Check return type
	if returnType != Int32Type {
		t.Errorf("Expected return type: %s\nGot: %s", Int32Type, returnType)
	}

	// Check argument types
	for _, a := range argTypes {
		if a != Int32Type {
			t.Errorf("Expected argument type: %s\nGot: %s", Int32Type, a)
		}
	}

	// Check parameter values
	v := f.Parameters()
	if len(v) != 2 {
		t.Errorf("Expected 2 parameters\nGot: %d", len(v))
	}
	for i, param := range v {
		name := fmt.Sprintf("%%%d", i)
		if param.llvm() != name {
			t.Errorf("Expected parameter %d to have name: %s\nGot: %s", i+1, name, param.llvm())
		}
	}

	// Check new block can be added
	b := f.AddBlock()
	if b.Name() != "2" {
		t.Errorf("Expected new block to have name: 2, got: %s", b.Name())
	}
}

func TestBlock(t *testing.T) {
	m := NewModule("testing")
	f := m.NewFunction("test", Int32Type, Int32Type, Int32Type)
	b := f.Entry()

	// Check function
	if !reflect.DeepEqual(b.Function(), f) {
		t.Errorf("Block is not in parent function")
	}

	// Check block name
	if b.Name() != "entry" {
		t.Errorf("Expected entry block with name entry\nGot: %s", b.Name())
	}
}

func TestInstruction(t *testing.T) {

	// Returns a new block
	var m *Module
	nb := func() *Block {
		m = NewModule("test")
		f := m.NewFunction("test", Int32Type)
		b := f.Entry()
		return b
	}

	cases := []struct {
		i          *Instruction
		stringName string
		t          Type
		llvm       string
	}{
		{
			i:          nb().Fadd(newName(Float32Type, "left"), newName(Float32Type, "right")),
			stringName: "fadd",
			t:          Float32Type,
			llvm:       "%0 = fadd f32 %left, %right",
		},
		{
			i:          nb().Fsub(newName(Float32Type, "left"), newName(Float32Type, "right")),
			stringName: "fsub",
			t:          Float32Type,
			llvm:       "%0 = fsub f32 %left, %right",
		},
		{
			i:          nb().Fmul(newName(Float32Type, "left"), newName(Float32Type, "right")),
			stringName: "fmul",
			t:          Float32Type,
			llvm:       "%0 = fmul f32 %left, %right",
		},
		{
			i:          nb().Fdiv(newName(Float32Type, "left"), newName(Float32Type, "right")),
			stringName: "fdiv",
			t:          Float32Type,
			llvm:       "%0 = fdiv f32 %left, %right",
		},
		{
			i:          nb().Add(newName(Int32Type, "left"), newName(Int32Type, "right")),
			stringName: "add",
			t:          Int32Type,
			llvm:       "%0 = add i32 %left, %right",
		},
		{
			i:          nb().Sub(newName(Int32Type, "left"), newName(Int32Type, "right")),
			stringName: "sub",
			t:          Int32Type,
			llvm:       "%0 = sub i32 %left, %right",
		},
		{
			i:          nb().Mul(newName(Int32Type, "left"), newName(Int32Type, "right")),
			stringName: "mul",
			t:          Int32Type,
			llvm:       "%0 = mul i32 %left, %right",
		},
		{
			i:          nb().Div(newName(Int32Type, "left"), newName(Int32Type, "right")),
			stringName: "div",
			t:          Int32Type,
			llvm:       "%0 = div i32 %left, %right",
		},
		{
			i:          nb().Ret(newName(Int32Type, "ret")),
			stringName: "ret",
			t:          NilType,
			llvm:       "ret i32 %ret",
		},
		{
			i:          nb().Br(newBlock(nil, "test")),
			stringName: "br",
			t:          NilType,
			llvm:       "br label %test",
		},
		{
			i:          nb().CondBr(newName(BoolType, "cond"), newBlock(nil, "testTrue"), newBlock(nil, "testFalse")),
			stringName: "br",
			t:          NilType,
			llvm:       "br i1 %cond, label %testTrue, label %testFalse",
		},
		{
			i:          nb().Call(newFunction(m, "test", Int32Type, Int32Type, Int32Type), ConstInt32(100), ConstInt32(200)),
			stringName: "call",
			t:          Int32Type,
			llvm:       "call i32 @test(i32 100, i32 200)",
		},
	}

	for _, c := range cases {
		iValue := c.i.Value()

		// Check string
		if c.i.String() != c.stringName {
			t.Errorf("Expected instruction to string to fadd\nGot:%s", c.i.String())
		}

		// Check type
		if iValue.Type() != c.t {
			t.Errorf("Expected instruction: %s, type to be %s\nGot:%s", c.stringName, c.t.String(), iValue.Type().String())
		}

		// Check llvm
		if c.i.llvm() != c.llvm {
			t.Errorf("Expected llvm: %s\nGot: %s", c.llvm, c.i.llvm())
		}
	}

}

func TestType(t *testing.T) {
	cases := []struct {
		t          Type
		llvm       string
		stringType string
	}{
		{
			t:          Int32Type,
			llvm:       "i32",
			stringType: "Int32",
		},
		{
			t:          Int64Type,
			llvm:       "i64",
			stringType: "Int64",
		},
		{
			t:          Float32Type,
			llvm:       "f32",
			stringType: "Float32",
		},
		{
			t:          Float64Type,
			llvm:       "f64",
			stringType: "Float64",
		},
		{
			t:          NilType,
			llvm:       "null",
			stringType: "Nil",
		},
		{
			t:          BoolType,
			llvm:       "i1",
			stringType: "Bool",
		},
	}

	for _, c := range cases {
		if c.t.LLVMType() != c.llvm {
			t.Errorf("Expected llvm type: %q\nGot: %q", c.llvm, c.t.LLVMType())
		}

		if c.t.String() != c.stringType {
			t.Errorf("Expected string type: %q\nGot: %q", c.stringType, c.t.String())
		}
	}
}

func TestValues(t *testing.T) {
	var (
		typeName          = 100
		typeFunctionValue = 101
		typeConstant      = 102
	)

	cases := []struct {
		v         Value
		valueType int
		name      string
		t         Type
		llvm      string
	}{
		{
			v:         newName(Int32Type, "test"),
			valueType: typeName,
			name:      "test",
			t:         Int32Type,
			llvm:      "%test",
		},
		{
			v:         newFunctionValue(Int32Type, "test"),
			valueType: typeFunctionValue,
			name:      "test",
			t:         Int32Type,
			llvm:      "@test",
		},
		{
			v:         ConstInt32(1002),
			valueType: typeConstant,
			t:         Int32Type,
			llvm:      "1002",
		},
	}

	for _, c := range cases {
		// Check type
		if c.v.Type() != c.t {
			t.Errorf("Expected type: %s, got type: %s", c.t.String(), c.v.Type().String())
		}

		// Check llvm
		if c.v.llvm() != c.llvm {
			t.Errorf("Expected llvm: %s, got: %s", c.llvm, c.v.llvm())
		}

		// Check value casts
		switch c.valueType {
		case typeName:
			name, ok := c.v.(Name)
			if !ok {
				t.Errorf("Expected value %+s to be of type name", c.v)
			}
			if c.name != name.Name() {
				t.Errorf("Expected name: %s, got name: %s", c.name, name.Name())
			}
		case typeConstant:
			_, ok := c.v.(Constant)
			if !ok {
				t.Errorf("Expected value %+s to be of type constant", c.v)
			}
		case typeFunctionValue:
			functionValue, ok := c.v.(FunctionValue)
			if !ok {
				t.Errorf("Expected value %+s to be of type functionValue", c.v)
			}
			if c.name != functionValue.Name() {
				t.Errorf("Expected name: %s, got name: %s", c.name, functionValue.Name())
			}
		}
	}
}

func TestLLVMCompile(t *testing.T) {
	var cases []struct {
		m    *Module
		llvm string
	}

	addCase := func(m *Module, llvm string) {
		cases = append(cases, struct {
			m    *Module
			llvm string
		}{m, llvm})
	}

	// Add two floats function
	{
		m := NewModule("testing")
		f := m.NewFunction("addFloats", Float32Type, Float32Type, Float32Type)
		b := f.Entry()
		add := b.Fadd(f.Parameters()[0], f.Parameters()[1])
		b.Ret(add.Value())

		addCase(m, `define f32 @addFloats(f32 %0, f32 %1){
	entry:
		%2 = fadd f32 %0, %1
		ret f32 %2
}`)
	}

	//if true function
	{
		m := NewModule("testing")
		f := m.NewFunction("ifs", Int32Type)
		b := f.Entry()

		ifTrue, ifFlase := f.AddBlock(), f.AddBlock()
		ifTrue.Ret(ConstInt32(100))
		ifFlase.Ret(ConstInt32(200))
		trueValue := ConstBool(true)
		b.CondBr(trueValue, ifTrue, ifFlase)

		addCase(m, `define i32 @ifs(){
	entry:
		br i1 true, label %0, label %1
	0:
		ret i32 100
	1:
		ret i32 200
}`)
	}

	for _, c := range cases {
		if c.m.LLVM() != c.llvm {
			t.Errorf("Expected:\n%s\nGot:\n%s", c.llvm, c.m.LLVM())
		}
	}
}
