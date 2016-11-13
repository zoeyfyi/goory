package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Interger subtruction
type Sub struct {
	block value.Value; name string
	lhs  value.Value
	rhs  value.Value
}

func NewSub(block value.Value, name string, lhs, rhs value.Value) *Sub {
	return &Sub{block, name, lhs, rhs}
}

func (i *Sub) Block() value.Value {
	return i.block
}

func (i *Sub) IsTerminator() bool {
	return false
}

func (i *Sub) Type() types.Type {
	return i.lhs.Type()
}

func (i *Sub) Ident() string {
	return "%" + i.name
}

func (i *Sub) Llvm() string {
	return fmt.Sprintf("%%%s = sub %s %s, %s",
		i.name,
		i.Type(),
		i.lhs.Ident(),
		i.rhs.Ident())
}
