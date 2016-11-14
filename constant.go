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

func (v constant) Llvm() string {
	return ""
}

func (v constant) Ident() string {
	switch v.value.(type) {
	case float32, float64:
		return fmt.Sprintf("%f", v.value)
	default:
		return fmt.Sprintf("%v", v.value)
	}
}

func (v constant) Type() types.Type {
	return v.constantType
}
