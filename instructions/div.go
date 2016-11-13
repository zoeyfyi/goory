package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Interger division
type Div struct {
	block value.Value; name string
	lhs  value.Value
	rhs  value.Value
}

// NewDiv creates a new integer division instruction
func NewDiv(block value.Value, name string, lhs value.Value, rhs value.Value) *Div {
	return &Div{block, name, lhs, rhs}
}

func (i *Div) Block() value.Value {
	return i.block
}

func (i *Div) IsTerminator() bool {
	return false
}

func (i *Div) Type() types.Type {
	return i.lhs.Type()
}

func (i *Div) Ident() string {
	return "%" + i.name
}

func (i *Div) Llvm() string {
	return fmt.Sprintf("%%%s = div %s %s, %s",
		i.name,
		i.Type(),
		i.lhs.Ident(),
		i.rhs.Ident())
}
