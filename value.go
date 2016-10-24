package goory

import "fmt"

// Value represents a parameter, the result of a function or a constant
type Value interface {
	Type() Type
	llvm() string
}

// Name reperesents a parameter or return value
type Name struct {
	t    Type
	name string
}

func newName(t Type, name string) Value {
	return Value(Name{t, name})
}

// Name returns the name of the Name
func (n Name) Name() string {
	return n.name
}

// Type returns the type of the Name
func (n Name) Type() Type {
	return n.t
}

func (n Name) llvm() string {
	return "%" + n.name
}

type FunctionValue struct {
	returnType Type
	name       string
}

func newFunctionValue(returnType Type, name string) Value {
	return Value(FunctionValue{returnType, name})
}

// Name returns the name of the FunctionValue
func (f FunctionValue) Name() string {
	return f.name
}

// Type returns the type of the FunctionValue
func (f FunctionValue) Type() Type {
	return f.returnType
}

func (f FunctionValue) llvm() string {
	return "@" + f.name
}

// Constant reprents a constant value
type Constant struct {
	t     Type
	value interface{}
}

// Type returns the type of the constant
func (c Constant) Type() Type {
	return c.t
}

func (c Constant) llvm() string {
	switch c.value.(type) {
	case int8, int16, int32, int64:
		return fmt.Sprintf("%d", c.value)
	case bool:
		return fmt.Sprintf("%t", c.value)
	default:
		panic("Unknown constant type")
	}
}

// ConstInt32 returns a constant value of type int32
func ConstInt32(i int32) Value {
	return Value(Constant{Int32Type, i})
}

// ConstBool returns a constant value of type bool
func ConstBool(b bool) Value {
	return Value(Constant{BoolType, b})
}
