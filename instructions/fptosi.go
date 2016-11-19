package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Float to int
type Fptosi struct {
	block value.Value
	name  string
	value value.Value
	cast  types.Type
}

func NewFptosi(block value.Value, name string, value value.Value, cast types.Type) *Fptosi {
	return &Fptosi{block, name, value, cast}
}

func (i *Fptosi) Block() value.Value {
	return i.block
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
	return fmt.Sprintf("%%%s = fptosi %s %s to %s",
		i.name,
		i.value.Type().String(),
		i.value.Ident(),
		i.cast.String())
}
