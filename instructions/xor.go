package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Xor struct {
	block value.Value; name string
	lhs  value.Value
	rhs  value.Value
}

func NewXor(block value.Value, name string, lhs, rhs value.Value) *Xor {
	return &Xor{block, name, lhs, rhs}
}

func (i *Xor) Block() value.Value {
	return i.block
}

func (i *Xor) IsTerminator() bool {
	return false
}

func (i *Xor) Type() types.Type {
	return types.NewBoolType()
}

func (i *Xor) Ident() string {
	return "%" + i.name
}

func (i *Xor) Llvm() string {
	return fmt.Sprintf("%s = xor %s %s %s",
		i.Ident(),
		i.lhs.Type().String(),
		i.lhs.Ident(),
		i.rhs.Ident(),
	)
}
