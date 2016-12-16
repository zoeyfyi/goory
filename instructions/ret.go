package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Ret struct {
	block value.Value
	value value.Value
}

func NewRet(block value.Value, value value.Value) *Ret {
	return &Ret{block, value}
}

func (i *Ret) Block() value.Value {
	return i.block
}

func (i *Ret) IsTerminator() bool {
	return true
}

func (i *Ret) Type() types.Type {
	return types.VOID
}

func (i *Ret) Ident() string {
	return ""
}

func (i *Ret) Llvm() string {
	return fmt.Sprintf("ret %s %s",
		i.value.Type().String(),
		i.value.Ident())
}
