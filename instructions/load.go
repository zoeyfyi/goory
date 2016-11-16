package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Load struct {
	block      value.Value
	name       string
	allocation *Alloca
}

// NewLoad creates a new Add operation
func NewLoad(block value.Value, name string, allocation *Alloca) *Load {
	return &Load{block, name, allocation}
}

func (i *Load) Block() value.Value {
	return i.block
}

func (i *Load) IsTerminator() bool {
	return false
}

func (i *Load) Type() types.Type {
	return i.allocation.BaseType()
}

func (i *Load) Ident() string {
	return "%" + i.name
}

func (i *Load) Llvm() string {
	return fmt.Sprintf("%%%s = load %s, %s %s",
		i.name,
		i.Type().String(),
		i.allocation.Type().String(),
		i.allocation.Ident())
}
