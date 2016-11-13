package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Sitofp struct {
	block value.Value
	name  string
	value value.Value
	cast  types.Type
}

func NewSitofp(block value.Value, name string, value value.Value, cast types.Type) *Sitofp {
	return &Sitofp{block, name, value, cast}
}

func (i *Sitofp) Block() value.Value {
	return i.block
}

func (i *Sitofp) IsTerminator() bool {
	return false
}

func (i *Sitofp) Type() types.Type {
	return i.cast
}

func (i *Sitofp) Ident() string {
	return "%" + i.name
}

func (i *Sitofp) Llvm() string {
	return fmt.Sprintf("%s = sitofp %s %s to %s",
		i.name,
		i.value.Type().String(),
		i.value.Llvm(),
		i.cast.String())
}
