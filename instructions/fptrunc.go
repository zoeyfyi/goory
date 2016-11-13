package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Fptrunc struct {
	name  string
	value value.Value
	cast  types.Type
}

func NewFptrunc(name string, value value.Value, cast types.Type) *Fptrunc {
	return &Fptrunc{name, value, cast}
}

func (i *Fptrunc) String() string {
	return "fptrunc"
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
		i.value.String(),
		i.cast.String())
}
