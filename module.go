package goory

import (
	"fmt"

	"github.com/bongo227/goory/types"
)

// Module is a single compilation unit
type Module struct {
	name      string
	functions []*Function
	tempCount int
}

// NewModule creates a new module with a name
func NewModule(name string) *Module {
	return &Module{
		name:      name,
		functions: []*Function{},
		tempCount: 0,
	}
}

// Name gets the name of the module
func (m *Module) Name() string {
	return m.name
}

// nextTempName increments a temp counter to produce temp names
func (m *Module) nextTempName() string {
	s := fmt.Sprintf("temp%d", m.tempCount)
	m.tempCount++
	return s
}

// NewFunction adds a new function to module
func (m *Module) NewFunction(name string, returnType types.Type) *Function {
	f := newFunction(m, name, returnType)
	m.functions = append(m.functions, f)
	return f
}

// LLVM returns the module as llvm ir
func (m *Module) LLVM() string {
	s := ""

	for findex, f := range m.functions {
		argString := ""
		for i, a := range f.args {
			argString += a.String()
			if i < len(f.args)-1 {
				argString += ", "
			}
		}

		s += fmt.Sprintf("define %s @%s(%s){\n",
			f.Type().String(), f.name, argString)

		for _, b := range f.blocks {
			s += "\t" + b.name + ":\n"
			for _, i := range b.instructions {
				s += "\t\t" + i.Llvm() + "\n"
			}
		}

		if findex == len(m.functions)-1 {
			s += "}"
		} else {
			s += "}\n\n"
		}
	}

	return s
}
