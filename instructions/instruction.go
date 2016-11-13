package instructions

import "github.com/bongo227/goory/types"

// Instruction is an operation that is executed on its operands
type Instruction interface {
	// Returns the llvm name of the instruction
	String() string
	// IsTerminator returns true if the function is a block terminator
	IsTerminator() bool
	// Type returns the type the function returns
	Type() types.Type
	// Llvm returns a string value of the llvm ir it represents
	Llvm() string
	// Ident returns the identifyer used to access the value of the operation
	Ident() string
}

// assertEqual checks all types are the same
func assertEqual(t ...types.Type) {
	if len(t) == 0 || len(t) == 1 {
		return
	}

	for i := 1; i < len(t); i++ {
		if !t[0].Equal(t[i]) {
			panic("Types are not equal")
		}
	}
}

func assertInt(t types.Type) {
	if _, ok := t.(types.IntType); !ok {
		panic("Value is not an integer")
	}
}

func assertFunction(t types.Type) types.FunctionType {
	fType, ok := t.(types.FunctionType)
	if !ok {
		panic("Value is not a function type")
	}

	return fType
}
