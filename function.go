package goory

import "fmt"

// ------------------------
// Function arguments
// ------------------------

type argument struct {
	argType Type
	name    string
}

// Argument represents a function argument
func Argument(argType Type, name string) argument { return argument{argType, name} }
func (v argument) Type() Type                     { return v.argType }
func (v argument) llvm() string                   { return fmt.Sprintf("%s %%%s", v.argType.llvm(), v.name) }
func (v argument) ident() string                  { return "%" + v.name }

// ------------------------
// Function
// ------------------------

// Function is a group of instructions that are executed is a new stack frame
type Function struct {
	module       *Module
	name         string
	functionType FunctionType
	args         []argument
	blocks       []*Block
}

// NewFunction creates a new function with an entry block
func newFunction(module *Module, name string, returnType Type, arguments ...argument) *Function {
	f := Function{
		module:       module,
		name:         name,
		functionType: NewFunctionType(returnType, arguments...),
		args:         arguments,
		blocks:       []*Block{},
	}

	// Create entry block
	b := newBlock(&f, "entry")
	f.blocks = append(f.blocks, b)
	return &f
}

// Module returns the module the function is in
func (f *Function) Module() *Module { return f.module }

// Name returns the function name
func (f *Function) Name() string { return f.name }

// Type returns the function type
func (f *Function) Type() FunctionType { return f.functionType }

// Arguments returns the values of function arguments
func (f *Function) Arguments() []argument { return f.args }

// AddBlock adds a new block to the function
func (f *Function) AddBlock() *Block {
	b := newBlock(f, f.module.nextTempName())
	f.blocks = append(f.blocks, b)
	return b
}

// Entry returns the entry block of the function
func (f *Function) Entry() *Block {
	// Entry will (almost?) always be on top so loop is no cost
	for i, b := range f.blocks {
		if b.name == "entry" {
			return f.blocks[i]
		}
	}

	panic("Function has no entry block, cant find entry block")
}
