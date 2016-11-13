package types

import "fmt"

// Aggregate is a type made up of atomic types
type Aggregate interface {
	String() string
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

func (t ArrayType) String() string {
	return fmt.Sprintf("[%s x %d]", t.baseType.String(), t.count)
}

// structType represents a collection of types
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
