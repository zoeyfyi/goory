package goory

import "fmt"

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
func (m *Module) NewFunction(name string, returnType Type, argTypes ...Type) *Function {
	f := newFunction(m, name, returnType, argTypes...)
	m.functions = append(m.functions, f)
	return f
}

// LLVM returns the module as llvm ir
func (m *Module) LLVM() string {
	s := ""

	for _, f := range m.functions {
		// Function definition
		rt, _ := f.Type()

		paramString := ""
		parameters := f.Parameters()
		for i, p := range parameters {
			paramString += fmt.Sprintf("%s %s", p.Type().LLVMType(), p.llvm())
			if i < len(parameters)-1 {
				paramString += ", "
			}
		}

		s += fmt.Sprintf("define %s @%s(%s){\n", rt.LLVMType(), f.name, paramString)

		for _, b := range f.blocks {
			s += "\t" + b.name + ":\n"
			for _, i := range b.instructions {
				s += "\t\t" + i.llvm() + "\n"
			}
		}

		s += "}\n\n"
	}

	return s
}
