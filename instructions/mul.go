package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Mul struct {
	block value.Value; name string
	lhs  value.Value
	rhs  value.Value
}

func NewMul(block value.Value, name string, lhs value.Value, rhs value.Value) *Mul {
	return &Mul{block, name, lhs, rhs}
}

func (i *Mul) Block() value.Value {
	return i.block
}

func (i *Mul) IsTerminator() bool {
	return false
}

func (i *Mul) Type() types.Type {
	return i.lhs.Type()
}

func (i *Mul) Ident() string {
	return "%" + i.name
}

func (i *Mul) Llvm() string {
	return fmt.Sprintf("%%%s = mul %s %s, %s",
		i.name,
		i.Type(),
		i.lhs.Ident(),
		i.rhs.Ident())
}
