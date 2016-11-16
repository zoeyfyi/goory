package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Alloca struct {
	block    value.Value
	name     string
	alocType types.Type
	number   value.Value
}

// NewAlloca creates a new Add operation
func NewAlloca(block value.Value, name string, alocType types.Type) *Alloca {
	return &Alloca{block, name, alocType, nil}
}

// SetNumber sets the number of elements to allocate for
func (i *Alloca) SetNumber(number value.Value) {
	i.number = number
}

func (i *Alloca) Block() value.Value {
	return i.block
}

func (i *Alloca) IsTerminator() bool {
	return false
}

func (i *Alloca) Type() types.Type {
	return types.NewPointerType(i.alocType)
}

func (i *Alloca) BaseType() types.Type {
	return i.alocType
}

func (i *Alloca) Ident() string {
	return "%" + i.name
}

func (i *Alloca) Llvm() string {
	if i.number != nil {
		return fmt.Sprintf("%%%s = alloca %s, %s",
			i.name,
			i.alocType.String(),
			i.number.Ident())
	} else {
		return fmt.Sprintf("%%%s = alloca %s",
			i.name,
			i.alocType.String())
	}
}
