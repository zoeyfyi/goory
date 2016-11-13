package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Float division
type Fdiv struct {
	name string
	lhs  value.Value
	rhs  value.Value
}

func NewFdiv(name string, lhs value.Value, rhs value.Value) *Fdiv {
	return &Fdiv{name, lhs, rhs}
}

func (i *Fdiv) String() string {
	return "fdiv"
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
