package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Sitofp struct {
	name  string
	value value.Value
	cast  types.Type
}

func NewSitofp(name string, value value.Value, cast types.Type) *Sitofp {
	return &Sitofp{name, value, cast}
}

func (i *Sitofp) String() string {
	return "sitofp"
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
		i.value.String(),
		i.cast.String())
}
