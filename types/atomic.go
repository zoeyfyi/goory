package types

import "fmt"

// Atomic defines an atomic llvm type
type Atomic interface {
	String() string
	Bits() int
	Equal(n Type) bool
}

// intType is the type of llvm integers
type IntType struct {
	bits int
}

// NewIntType creates a new int type bits long
func NewIntType(bits int) Type {
	return IntType{bits}
}

func (t IntType) String() string {
	return fmt.Sprintf("i%d", t.bits)
}

func (t IntType) Bits() int {
	return t.bits
}

func (t IntType) Equal(n Type) bool {
	if nInt, ok := n.(IntType); ok {
		return nInt.Bits() == t.Bits()
	}

	return false
}

// FloatType is the type of llvm single precision floats
type FloatType struct{}

func NewFloatType() Type {
	return FloatType{}
}

func (t FloatType) String() string {
	return "float"
}

func (t FloatType) Bits() int {
	return 32
}

func (t FloatType) Equal(n Type) bool {
	_, ok := n.(FloatType)
	return ok
}

// DoubleType is the type of llvm double precision floats
type DoubleType struct{}

// NewDoubleType creates a new double precision float type
func NewDoubleType() Type {
	return DoubleType{}
}

func (t DoubleType) String() string {
	return "double"
}

func (t DoubleType) Bits() int {
	return 64
}

func (t DoubleType) Equal(n Type) bool {
	_, ok := n.(DoubleType)
	return ok
}

// BoolType is a single bit integer which represents booleans in llvm
type BoolType struct{}

// NewBoolType creates a new bool type
func NewBoolType() Type {
	return BoolType{}
}

func (t BoolType) String() string {
	return "bool"
}

func (t BoolType) Bits() int {
	return 1
}

func (t BoolType) Equal(n Type) bool {
	_, ok := n.(BoolType)
	return ok
}

// VoidType is a zero bit void/nil/null type in llvm
type VoidType struct{}

// NewVoidType creates a new void type
func NewVoidType() Type {
	return VoidType{}
}

func (t VoidType) String() string {
	return "void"
}

func (t VoidType) Bits() int {
	return 0
}

func (t VoidType) Equal(n Type) bool {
	_, ok := n.(VoidType)
	return ok
}
