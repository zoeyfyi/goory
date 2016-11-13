package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Float multiplication
type Fmul struct {
	block value.Value; name string
	lhs  value.Value
	rhs  value.Value
}

func NewFmul(block value.Value, name string, lhs, rhs value.Value) *Fmul {
	return &Fmul{block, name, lhs, rhs}
}

func (i *Fmul) Block() value.Value {
	return i.block
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
