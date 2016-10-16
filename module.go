package goory

import "fmt"

// Module is a single compilation unit
type Module struct {
	name      string
	functions []Function
}

// Name gets the name of the module
func (m *Module) Name() string {
	return m.name
}

// NewModule creates a new module with a name
func NewModule(name string) *Module {
	return &Module{
		name:      name,
		functions: []Function{},
	}
}

// AddFunction adds a function to the module
func (m *Module) AddFunction(function *Function) {
	m.functions = append(m.functions, *function)
}

// LLVM returns the module as llvm ir
func (m *Module) LLVM() string {
	s := ""

	for _, f := range m.functions {
		// Function definition
		rt := f.Type().ReturnType()
		s += fmt.Sprintf("define %s @%s(){\n", rt.LLVMType(), f.name)

		for _, i := range f.Entry().instructions {
			s += i.llvm() + "\n"
		}

		s += "}\n"
	}

	return s
}
