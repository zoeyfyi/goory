package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// And represents a bitwise And function
type And struct {
	name string
	lhs  value.Value
	rhs  value.Value
}

// NewAnd creates a refrence to a new And instruction
func NewAnd(name string, lhs value.Value, rhs value.Value) *And {
	assertEqual(lhs.Type(), rhs.Type())
	lhs.Type().Equal(types.NewBoolType())
	return &And{name, lhs, rhs}
}

// String returns the function name 'and'
func (i *And) String() string {
	return "and"
}

// IsTerminator returns false since and is not a block terminator
func (i *And) IsTerminator() bool {
	return false
}

// Type returns the return type of And 'bool'
func (i *And) Type() types.Type {
	return types.NewBoolType()
}

// Ident returns the identifyer used to get the result of this And instruction
func (i *And) Ident() string {
	return "%" + i.name
}

// Llvm returns the llvm ir string of the And instruction
func (i *And) Llvm() string {
	return fmt.Sprintf("%s = and %s %s %s",
		i.Ident(),
		i.lhs.Type().String(),
		i.lhs.Ident(),
		i.rhs.Ident(),
	)
}
