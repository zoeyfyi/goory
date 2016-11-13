package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Fptrunc struct {
	block value.Value
	name  string
	value value.Value
	cast  types.Type
}

func NewFptrunc(block value.Value, name string, value value.Value, cast types.Type) *Fptrunc {
	return &Fptrunc{block, name, value, cast}
}

func (i *Fptrunc) Block() value.Value {
	return i.block
}

func (i *Fptrunc) IsTerminator() bool {
	return false
}

func (i *Fptrunc) Type() types.Type {
	return i.cast
}

func (i *Fptrunc) Ident() string {
	return "%" + i.name
}

func (i *Fptrunc) Llvm() string {
	return fmt.Sprintf("%s = fptrunc %s %s to %s",
		i.name,
		i.value.Type().String(),
		i.value.Llvm(),
		i.cast.String())
}
