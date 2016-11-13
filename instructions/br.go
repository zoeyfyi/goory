package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Branch statement
type Br struct {
	block  value.Value
	name   string
	branch value.Value
}

func NewBr(block value.Value, name string, branch value.Value) *Br {
	block.Type().Equal(types.NewBlockType())
	return &Br{block, name, branch}
}

func (i *Br) Block() value.Value {
	return i.block
}

func (i *Br) IsTerminator() bool {
	return true
}

func (i *Br) Type() types.Type {
	return types.NewVoidType()
}

func (i *Br) Ident() string {
	return "%" + i.name
}

func (i *Br) Llvm() string {
	return fmt.Sprintf("br label %%%s", i.branch.Ident())
}
