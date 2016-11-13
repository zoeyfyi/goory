package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Float subtruction
type Fsub struct {
	name string
	lhs  value.Value
	rhs  value.Value
}

func NewFsub(name string, lhs, rhs value.Value) *Fsub {
	return &Fsub{name, lhs, rhs}
}

func (i *Fsub) String() string {
	return "fsub"
}

func (i *Fsub) IsTerminator() bool {
	return false
}

func (i *Fsub) Type() types.Type {
	return i.lhs.Type()
}

func (i *Fsub) Ident() string {
	return "%" + i.name
}

func (i *Fsub) Llvm() string {
	return fmt.Sprintf("%%%s = fsub %s %s, %s",
		i.name,
		i.Type(),
		i.lhs.Ident(),
		i.rhs.Ident())
}
