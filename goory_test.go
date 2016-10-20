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
	if !reflect.DeepEqual(f.Module(), &m) {
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
		name := fmt.Sprintf("%d", i)
		if param.Name() != name {
			t.Errorf("Expected parameter %d to have name: %s\nGot: %s", i+1, name, param.Name())
		}
	}
}

func TestBlock(t *testing.T) {
	m := NewModule("testing")
	f := m.NewFunction("test", Int32Type, Int32Type, Int32Type)
	b := f.Entry()

	// Check function
	if !reflect.DeepEqual(b.Function(), &f) {
		t.Errorf("Block is not in parent function")
	}

	// Check block name
	if b.Name() != "entry" {
		t.Errorf("Expected entry block with name entry\nGot: %s", b.Name())
	}
}

func TestInstruction(t *testing.T) {

	// Returns a new block
	nb := func() *Block {
		m := NewModule("testing")
		f := m.NewFunction("test", Int32Type)
		b := f.Entry()
		return b
	}

	cases := []struct {
		i          Instruction
		stringName string
		t          Type
		llvm       string
	}{
		{
			i:          nb().Fadd(newValue(Float32Type, "left"), newValue(Float32Type, "right")),
			stringName: "fadd",
			t:          Float32Type,
			llvm:       "%0 = fadd f32 %left, %right",
		},
		{
			i:          nb().Ret(newValue(Int32Type, "ret")),
			stringName: "ret",
			t:          NilType,
			llvm:       "ret i32 %ret",
		},
	}

	for _, c := range cases {
		iValue := c.i.Value()

		// Check name
		if iValue.Name() != "0" {
			// 0 and 1 for parameters so instruction should be 2
			t.Errorf("Expected instruction to have name of 0")
		}

		// Check string
		if c.i.String() != c.stringName {
			t.Errorf("Expected instruction to string to fadd\nGot:%s", c.i.String())
		}

		// Check type
		if iValue.Type() != c.t {
			t.Errorf("Expected instruction type to be float32\nGot:%s", iValue.Type().String())
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

func TestLLVMCompile(t *testing.T) {
	var cases []struct {
		m    Module
		llvm string
	}

	addCase := func(m Module, llvm string) {
		cases = append(cases, struct {
			m    Module
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
	%2 = fadd f32 %0, %1
	ret f32 %2
}`)
	}

	for _, c := range cases {
		if c.m.LLVM() != c.llvm {
			t.Errorf("Expected:\n%s\nGot:\n%s", c.llvm, c.m.LLVM())
		}
	}
}
