package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Getelementptr
type Getelementptr struct {
	block    value.Value
	name     string
	typ      types.Type
	element  value.Pointer
	location []value.Value
}

func NewGetelementptr(block value.Value, name string, typ types.Type, element value.Pointer, location ...value.Value) *Getelementptr {
	return &Getelementptr{block, name, typ, element, location}
}

func (i *Getelementptr) Block() value.Value {
	return i.block
}

func (i *Getelementptr) IsTerminator() bool {
	return false
}

func (i *Getelementptr) Type() types.Type {
	return types.NewPointerType(i.typ)
}

func (i *Getelementptr) BaseType() types.Type {
	return i.typ
}

func (i *Getelementptr) Ident() string {
	return "%" + i.name
}

func (i *Getelementptr) Llvm() string {
	location := ""
	for lIndex, loc := range i.location {
		location += fmt.Sprintf("%s %s", loc.Type().String(), loc.Ident())
		if lIndex < len(i.location)-1 {
			location += ", "
		}
	}

	return fmt.Sprintf("%%%s = getelementptr %s, %s %s, %s",
		i.name,
		i.element.BaseType().String(),
		i.element.Type().String(),
		i.element.Ident(),
		location)
}
