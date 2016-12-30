package types

import "fmt"

// Aggregate is a type made up of atomic types
type Aggregate interface {
	String() string
	Position(int) Type
	Equal(n Type) bool
}

// arrayType is the type of llvm arrays
type ArrayType struct {
	baseType Type
	count    int
}

// NewArrayType creates a new array type
func NewArrayType(baseType Type, count int) Aggregate {
	return ArrayType{baseType, count}
}

// BaseType returns the bast type of the array type
func (t ArrayType) BaseType() Type {
	return t.baseType
}

func (t ArrayType) String() string {
	return fmt.Sprintf("[%d x %s]", t.count, t.baseType.String())
}

// Position returns the type at position n
func (t ArrayType) Position(n int) Type {
	return t.baseType
}

// Equal checks if type n equals type t
func (t ArrayType) Equal(n Type) bool {
	arr, ok := n.(ArrayType)
	if !ok {
		return false
	}

	return t.baseType == arr.baseType && t.count == arr.count
}

// StructType represents a collection of types
type StructType struct {
	types []Type
}

// NewStructType creates a new struct type
func NewStructType(types ...Type) Aggregate {
	return StructType{types}
}

func (t StructType) String() string {
	s := "{ "
	for i, structType := range t.types {
		s += structType.String()
		if i < len(t.types)-1 {
			s += ", "
		}
	}

	return s + " }"
}

// Position returns the type at position n
func (t StructType) Position(n int) Type {
	return t.types[n]
}

// Equal checks if type n equals type t
func (t StructType) Equal(n Type) bool {
	str, ok := n.(StructType)
	if !ok {
		return false
	}

	for i := 0; i < len(t.types); i++ {
		if !t.types[i].Equal(str.types[i]) {
			return false
		}
	}

	return true
}
