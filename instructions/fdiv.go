package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Float division
type Fdiv struct {
	block value.Value; name string
	lhs  value.Value
	rhs  value.Value
}

func NewFdiv(block value.Value, name string, lhs value.Value, rhs value.Value) *Fdiv {
	return &Fdiv{block, name, lhs, rhs}
}

func (i *Fdiv) Block() value.Value {
	return i.block
}

func (i *Fdiv) IsTerminator() bool {
	return false
}

func (i *Fdiv) Type() types.Type {
	return i.lhs.Type()
}

func (i *Fdiv) Ident() string {
	return "%" + i.name
}

func (i *Fdiv) Llvm() string {
	return fmt.Sprintf("%%%s = fdiv %s %s, %s",
		i.name,
		i.Type(),
		i.lhs.Ident(),
		i.rhs.Ident())
}
