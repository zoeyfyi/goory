package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Ret struct {
	name  string
	value value.Value
}

func NewRet(name string, value value.Value) *Ret {
	return &Ret{name, value}
}

func (i *Ret) String() string {
	return "ret"
}

func (i *Ret) IsTerminator() bool {
	return true
}

func (i *Ret) Type() types.Type {
	return types.NewVoidType()
}

func (i *Ret) Ident() string {
	return "%" + i.name
}

func (i *Ret) Llvm() string {
	return fmt.Sprintf("ret %s %s",
		i.value.Type().String(),
		i.value.Ident())
}
