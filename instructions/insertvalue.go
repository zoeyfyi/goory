package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"
	"github.com/bongo227/goory/value"
)

type Insertvalue struct {
	block value.Value; name string
	location value.Value
	value    value.Value
	position int
}

func NewInsertvalue(block value.Value, name string, location, value value.Value, position int) *Insertvalue {
	return &Insertvalue{block, name, location, value, position}
}

func (i *Insertvalue) Block() value.Value {
	return i.block
}

func (i *Insertvalue) IsTerminator() bool {
	return false
}

func (i *Insertvalue) Type() types.Type {
	return i.location.Type()
}

func (i *Insertvalue) Ident() string {
	return "%" + i.name
}

func (i *Insertvalue) Llvm() string {
	return fmt.Sprintf("%%%s = insertvalue %s %s, %s %s, %d",
		i.name,
		i.location.Type().String(),
		i.location.Ident(),
		i.value.Type().String(),
		i.value.Ident(),
		i.position)
}
