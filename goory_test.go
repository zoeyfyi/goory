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
	m := NewModule("testing")
	f := m.NewFunction("test", Int32Type, Float32Type, Float32Type)
	b := f.Entry()
	i := b.Fadd(f.Parameters()[0], f.Parameters()[1])

	// Check name
	if i.Name() != "2" {
		// 0 and 1 for parameters so instruction should be 2
		t.Errorf("Expected instruction to have name of 2")
	}

	// Check string
	if i.String() != "fadd" {
		t.Errorf("Expected instruction to string to fadd\nGot:%s", i.String())
	}

	// Check type
	if i.Type() != Float32Type {
		t.Errorf("Expected instruction type to be float32\nGot:%s", i.Type().String())
	}

	// Check llvm
	llvm := "%2 = fadd f32 %0, %1"
	if i.llvm() != llvm {
		t.Errorf("Expected llvm: %s\nGot: %s", llvm, i.llvm())
	}
}
