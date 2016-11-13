package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Interger comparisons
type Icmp struct {
	name string
	mode string
	lhs  value.Value
	rhs  value.Value
}

func NewIcmp(name, mode string, lhs, rhs value.Value) *Icmp {
	return &Icmp{name, mode, lhs, rhs}
}

func (i *Icmp) String() string {
	return "icmp"
}

func (i *Icmp) IsTerminator() bool {
	return false
}

func (i *Icmp) Type() types.Type {
	return types.NewBoolType()
}

func (i *Icmp) Ident() string {
	return "%" + i.name
}

func (i *Icmp) Llvm() string {
	return fmt.Sprintf("%%%s = icmp %s %s %s, %s", i.name, i.mode, i.Type().String(), i.lhs.Ident(), i.rhs.Ident())
}
