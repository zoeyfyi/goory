package goory

import "fmt"

// Type defines all llvm types
type Type interface {
	llvm() string
}

// PrimativeType defines a primative llvm type
type PrimativeType interface {
	llvm() string
	Bits() int
}

type intType struct{ bits int }

// IntType is the type of llvm integers
func IntType(bits int) Type    { return IntType{bits} }
func (t intType) llvm() string { return fmt.Sprintf("i%d", t.bits) }
func (t intType) Bits() int    { return t.bits }

type floatType struct{}

// FloatType is the type of llvm single precision floats
func FloatType() Type            { return FloatType{} }
func (t floatType) llvm() string { return "float" }
func (t floatType) Bits() int    { return 32 }

type doubleType struct{}

// DoubleType is the type of llvm double precision floats
func DoubleType() Type            { return DoubleType{} }
func (t DoubleType) llvm() string { return "double" }
func (t DoubleType) Bits() int    { return 64 }

type boolType struct{}

// BoolType is a single bit integer which represents booleans in llvm
func BoolType() Type            { return boolType{} }
func (t BoolType) llvm() string { return "bool" }
func (t BoolType) Bits() int    { return 1 }

type voidType struct{}

// VoidType is a zero bit void/nil/null type in llvm
func VoidType() Type            { return voidType{} }
func (t VoidType) llvm() string { return "void" }
func (t VoidType) Bits() int    { return 0 }

// CompareTypes compares two types a and b
// returns 1 	if a > b
// returns 0 	if a = b
// returns -1 	if a < b
func (a Type) Compare(b Type) int {
	if a.Bits() > b.Bits() {
		return 1
	}

	if a.Bits() == b.Bits() {
		return 0
	}

	return -1
}

type AggregateType interface {
	llvm() string
}

type arrayType struct {
	baseType Type
	count    int
}

// ArrayType is the type of llvm arrays
func ArrayType(baseType Type, count int) Type { return ArrayType{baseType, count} }
func (t ArrayType) llvm() string              { return fmt.Sprintf("[%s x %d]", t.baseType.llvm(), t.count) }

type structType struct {
	types []Type
}

func StructType(types ...Type) Type { return structType{types} }
func (t structType) llvm() string {
	s := "{ "
	for i, structType := range t.types {
		s += structType.llvm()
		if i < len(t.types)-1 {
			s += ", "
		}
	}
	s += " }"
}

type functionType interface {
	returnType Type
	argTypes []Type
}

func FunctionType(returnType Type, argTypes ...Type) Type { return functionType{returnType, argTypes} }