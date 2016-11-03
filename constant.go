package goory

import "fmt"

type constant struct {
	constantType Type
	value        interface{}
}

// Constant reperesents a literal value
func Constant(constantType Type, value interface{}) Value { return constant{constantType, value} }
func (v constant) llvm() string                           { return "" }
func (v constant) ident() string                          { return fmt.Sprintf("%v", v.value) }
func (v constant) Type() Type                             { return v.constantType }
