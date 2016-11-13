package value

import "github.com/bongo227/goory/types"

// Instruction is an operation that is executed on its operands
type Instruction interface {
	// Returns the parent block
	Block() Value
	// IsTerminator returns true if the function is a block terminator
	IsTerminator() bool
	// Type returns the type the function returns
	Type() types.Type
	// Llvm returns a string value of the llvm ir it represents
	Llvm() string
	// Ident returns the identifyer used to access the value of the operation
	Ident() string
}
