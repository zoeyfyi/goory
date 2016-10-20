package goory

import "fmt"

// Module is a single compilation unit
type Module struct {
	name      string
	functions []Function
	tempCount int
}

// NewModule creates a new module with a name
func NewModule(name string) Module {
	return Module{
		name:      name,
		functions: []Function{},
		tempCount: 0,
	}
}

// Name gets the name of the module
func (m *Module) Name() string {
	return m.name
}

// nextTempName increments a temp counter to produce temp names
func (m *Module) nextTempName() string {
	s := fmt.Sprintf("%d", m.tempCount)
	m.tempCount++
	return s
}

// NewFunction adds a new function to module
func (m *Module) NewFunction(name string, returnType Type, argTypes ...Type) Function {
	return newFunction(m, name, returnType, argTypes...)
}

// LLVM returns the module as llvm ir
func (m *Module) LLVM() string {
	s := ""

	for _, f := range m.functions {
		// Function definition
		rt, _ := f.Type()
		s += fmt.Sprintf("define %s @%s(){\n", rt.LLVMType(), f.name)

		for _, i := range f.Entry().instructions {
			s += i.llvm() + "\n"
		}

		s += "}\n"
	}

	return s
}
