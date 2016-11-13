package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Branch statement
type Br struct {
	name  string
	block value.Value
}

func NewBr(name string, block value.Value) *Br {
	block.Type().Equal(types.NewBlockType())
	return &Br{name, block}
}

func (i *Br) String() string {
	return "br"
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
	return fmt.Sprintf("br label %%%s", i.block.Ident())
}
