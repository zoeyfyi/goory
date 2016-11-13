package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Float multiplication
type Fmul struct {
	name string
	lhs  value.Value
	rhs  value.Value
}

func NewFmul(name string, lhs, rhs value.Value) *Fmul {
	return &Fmul{name, lhs, rhs}
}

func (i *Fmul) String() string {
	return "fmul"
}

func (i *Fmul) IsTerminator() bool {
	return false
}

func (i *Fmul) Type() types.Type {
	return i.lhs.Type()
}

func (i *Fmul) Ident() string {
	return "%" + i.name
}

func (i *Fmul) Llvm() string {
	return fmt.Sprintf("%%%s = fmul %s %s, %s",
		i.name,
		i.Type(),
		i.lhs.Ident(),
		i.rhs.Ident())
}
