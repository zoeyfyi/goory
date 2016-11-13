package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Fpext struct {
	block value.Value
	name  string
	value value.Value
	cast  types.Type
}

func NewFpext(block value.Value, name string, value value.Value, cast types.Type) *Fpext {
	return &Fpext{block, name, value, cast}
}

func (i *Fpext) Block() value.Value {
	return i.block
}

func (i *Fpext) IsTerminator() bool {
	return false
}

func (i *Fpext) Type() types.Type {
	return i.cast
}

func (i *Fpext) Ident() string {
	return "%" + i.name
}

func (i *Fpext) Llvm() string {
	return fmt.Sprintf("%s = fpext %s %s to %s",
		i.name,
		i.value.Type().String(),
		i.value.Llvm(),
		i.cast.String())
}
