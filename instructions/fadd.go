package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Float addition
type Fadd struct {
	name string
	lhs  value.Value
	rhs  value.Value
}

func NewFadd(name string, lhs value.Value, rhs value.Value) *Fadd {
	assertEqual(lhs.Type(), rhs.Type())
	if !lhs.Type().Equal(types.NewFloatType()) && !lhs.Type().Equal(types.NewDoubleType()) {
		panic("Types are not float or double types")
	}

	return &Fadd{name, lhs, rhs}
}

func (i *Fadd) String() string {
	return "fadd"
}

func (i *Fadd) IsTerminator() bool {
	return false
}

func (i *Fadd) Type() types.Type {
	return i.lhs.Type()
}

func (i *Fadd) Ident() string {
	return "%" + i.name
}

func (i *Fadd) Llvm() string {
	return fmt.Sprintf("%%%s = fadd %s %s, %s",
		i.name,
		i.Type(),
		i.lhs.Ident(),
		i.rhs.Ident())
}
