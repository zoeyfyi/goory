package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Bitwise or
type Or struct {
	block value.Value; name string
	lhs  value.Value
	rhs  value.Value
}

func NewOr(block value.Value, name string, lhs, rhs value.Value) *Or {
	return &Or{block, name, lhs, rhs}
}

func (i *Or) Block() value.Value {
	return i.block
}

func (i *Or) IsTerminator() bool {
	return false
}

func (i *Or) Type() types.Type {
	return types.NewBoolType()
}

func (i *Or) Ident() string {
	return "%" + i.name
}

func (i *Or) Llvm() string {
	return fmt.Sprintf("%s = or %s %s %s",
		i.Ident(),
		i.lhs.Type().String(),
		i.lhs.Ident(),
		i.rhs.Ident(),
	)
}
