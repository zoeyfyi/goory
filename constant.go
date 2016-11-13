package goory

import (
	"fmt"

	"github.com/bongo227/goory/types"
	"github.com/bongo227/goory/value"
)

type constant struct {
	constantType types.Type
	value        interface{}
}

// Constant reperesents a literal value
func Constant(constantType types.Type, value interface{}) value.Value {
	return constant{constantType, value}
}

func (v constant) String() string {
	return ""
}

func (v constant) Ident() string {
	return fmt.Sprintf("%v", v.value)
}

func (v constant) Type() types.Type {
	return v.constantType
}
