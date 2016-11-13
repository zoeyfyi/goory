package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Float to int
type Fptosi struct {
	name  string
	value value.Value
	cast  types.Type
}

func NewFptosi(name string, value value.Value, cast types.Type) *Fptosi {
	return &Fptosi{name, value, cast}
}

func (i *Fptosi) String() string {
	return "fptosi"
}

func (i *Fptosi) IsTerminator() bool {
	return false
}

func (i *Fptosi) Type() types.Type {
	return i.cast
}

func (i *Fptosi) Ident() string {
	return "%" + i.name
}

func (i *Fptosi) Llvm() string {
	return fmt.Sprintf("%s = fptosi %s %s to %s",
		i.name,
		i.value.Type().String(),
		i.value.String(),
		i.cast.String())
}
