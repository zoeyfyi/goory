package goory

import "fmt"

// Value are what modules mostly consist with
type Value interface {
	Type() Type
	llvm() string
	ident() string
}

// TODO: should we have a global interface?

type global struct {
	globalType   Type
	initialValue Value
	name         string
	constant     bool
}

// Global is a varible global to the whole module
func Global(globalType Type, initialValue Value, name string, constant bool) Value {
	return global{globalType, initialValue, name, constant}
}

func (v global) Type() Type { return v.globalType }

func (v global) llvm() string {
	globconst := "global"
	if v.constant {
		globconst = "constant"
	}

	return fmt.Sprintf("@%s = external %s %s %s", v.name, globconst, v.globalType.llvm(), v.initialValue.llvm())
}

func (v global) ident() string {
	return "@" + v.name
}
