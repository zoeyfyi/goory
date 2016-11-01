package goory

import "fmt"

// Type defines all llvm types
type Type interface {
	llvm() string
}

// ------------------------
// Atomic Types
// ------------------------

// Atomic defines an atomic llvm type
type Atomic interface {
	llvm() string
	Bits() int
	IsInteger() bool
	IsFloat() bool
	IsDouble() bool
}

type intType struct{ bits int }

// IntType is the type of llvm integers
func NewIntType(bits int) Type    { return intType{bits} }
func (t intType) llvm() string    { return fmt.Sprintf("i%d", t.bits) }
func (t intType) Bits() int       { return t.bits }
func (t intType) IsInteger() bool { return true }
func (t intType) IsFloat() bool   { return false }
func (t intType) IsDouble() bool  { return false }

type floatType struct{}

// FloatType is the type of llvm single precision floats
func NewFloatType() Type            { return floatType{} }
func (t floatType) llvm() string    { return "float" }
func (t floatType) Bits() int       { return 32 }
func (t floatType) IsInteger() bool { return false }
func (t floatType) IsFloat() bool   { return true }
func (t floatType) IsDouble() bool  { return false }

type doubleType struct{}

// DoubleType is the type of llvm double precision floats
func NewDoubleType() Type            { return doubleType{} }
func (t doubleType) llvm() string    { return "double" }
func (t doubleType) Bits() int       { return 64 }
func (t doubleType) IsInteger() bool { return false }
func (t doubleType) IsFloat() bool   { return false }
func (t doubleType) IsDouble() bool  { return true }

type boolType struct{}

// BoolType is a single bit integer which represents booleans in llvm
func NewBoolType() Type            { return boolType{} }
func (t boolType) llvm() string    { return "bool" }
func (t boolType) Bits() int       { return 1 }
func (t boolType) IsInteger() bool { return false }
func (t boolType) IsFloat() bool   { return false }
func (t boolType) IsDouble() bool  { return false }

type voidType struct{}

// VoidType is a zero bit void/nil/null type in llvm
func NewVoidType() Type            { return voidType{} }
func (t voidType) llvm() string    { return "void" }
func (t voidType) Bits() int       { return 0 }
func (t voidType) IsInteger() bool { return false }
func (t voidType) IsFloat() bool   { return false }
func (t voidType) IsDouble() bool  { return false }

// Compare compares two atomics a and b
// returns 1 	if a > b
// returns 0 	if a = b
// returns -1 	if a < b
func Compare(a Atomic, b Atomic) int {
	if a.Bits() > b.Bits() {
		return 1
	}

	if a.Bits() == b.Bits() {
		return 0
	}

	return -1
}

func IsInteger(a Atomic) bool {
	_, ok := a.(intType)
	return ok
}

func IsFloat(a Atomic) bool {
	_, ok := a.(floatType)
	return ok
}

func IsDouble(a Atomic) bool {
	_, ok := a.(doubleType)
	return ok
}

// ------------------------
// Aggeregate Types
// ------------------------

// Aggregate is a type made up of atomic types
type Aggregate interface {
	llvm() string
}

type arrayType struct {
	baseType Type
	count    int
}

// ArrayType is the type of llvm arrays
func NewArrayType(baseType Type, count int) Aggregate { return arrayType{baseType, count} }
func (t arrayType) llvm() string                      { return fmt.Sprintf("[%s x %d]", t.baseType.llvm(), t.count) }

type structType struct {
	types []Type
}

func NewStructType(types ...Type) Aggregate { return structType{types} }
func (t structType) llvm() string {
	s := "{ "
	for i, structType := range t.types {
		s += structType.llvm()
		if i < len(t.types)-1 {
			s += ", "
		}
	}

	return s + " }"
}

// ------------------------
// Function Type
// ------------------------

type FunctionType struct {
	returnType Type
	argTypes   []argument
}

// FunctionType is the type of a function
func NewFunctionType(returnType Type, argTypes ...argument) FunctionType {
	return FunctionType{returnType, argTypes}
}
