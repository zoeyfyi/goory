package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Srem struct {
	block value.Value
	name  string
	lhs   value.Value
	rhs   value.Value
}

// NewSrem creates a new srem operation
func NewSrem(block value.Value, name string, lhs value.Value, rhs value.Value) *Srem {
	assertEqual(lhs.Type(), rhs.Type())
	assertInt(lhs.Type())
	return &Srem{block, name, lhs, rhs}
}

func (i *Srem) Block() value.Value {
	return i.block
}

func (i *Srem) IsTerminator() bool {
	return false
}

func (i *Srem) Type() types.Type {
	return i.lhs.Type()
}

func (i *Srem) Ident() string {
	return "%" + i.name
}

func (i *Srem) Llvm() string {
	return fmt.Sprintf("%%%s = srem %s %s, %s",
		i.name,
		i.Type(),
		i.lhs.Ident(),
		i.rhs.Ident())
}
