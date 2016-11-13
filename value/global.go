package value

import (
	"fmt"

	"github.com/bongo227/goory/types"
)

// TODO: should we have a global interface?
type global struct {
	globalType   types.Type
	initialValue Value
	name         string
	constant     bool
}

// Global is a varible global to the whole module
func Global(globalType types.Type, initialValue Value, name string, constant bool) Value {
	return global{globalType, initialValue, name, constant}
}

func (v global) Type() types.Type {
	return v.globalType
}

func (v global) Llvm() string {
	globconst := "global"
	if v.constant {
		globconst = "constant"
	}

	return fmt.Sprintf("@%s = external %s %s %s", v.name, globconst, v.globalType.String(), v.initialValue.Llvm())
}

func (v global) Ident() string {
	return "@" + v.name
}
