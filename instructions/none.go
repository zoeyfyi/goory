package instructions

import "github.com/bongo227/goory/types"

// None instruction
type None struct{}

func NewNone() *None {
	return &None{}
}

func (i *None) IsTerminator() bool {
	return false
}

func (i *None) Type() types.Type {
	return types.VOID
}

func (i *None) Ident() string {
	return ""
}

func (i *None) Llvm() string {
	return ""
}
