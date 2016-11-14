package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"
	"github.com/bongo227/goory/value"
)

type Extractvalue struct {
	block value.Value; name string
	location value.Value
	position value.Value
}

func NewExtractvalue(block value.Value, name string, location value.Value, position value.Value) *Extractvalue {
	return &Extractvalue{block, name, location, position}
}

func (i *Extractvalue) Block() value.Value {
	return i.block
}

func (i *Extractvalue) IsTerminator() bool {
	return false
}

func (i *Extractvalue) Type() types.Type {
	return i.location.Type().(types.Aggregate).Position(0)
}

func (i *Extractvalue) Ident() string {
	return "%" + i.name
}

func (i *Extractvalue) Llvm() string {
	return fmt.Sprintf("%%%s = extractvalue %s %s, %s",
		i.name,
		i.location.Type().String(),
		i.location.Ident(),
		i.position.Ident())
}
