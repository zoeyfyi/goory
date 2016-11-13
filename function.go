package goory

import (
	"fmt"

	"github.com/bongo227/goory/types"
)

// Function is a group of instructions that are executed is a new stack frame
type Function struct {
	module     *Module
	name       string
	returnType types.Type
	args       []types.Arg
	blocks     []*Block
}

// AddArgument adds a new parameter to the function
func (f *Function) AddArgument(argType types.Type, name string) types.Arg {
	arg := types.NewArg(argType, name)
	f.args = append(f.args, arg)
	return arg
}

// NewFunction creates a new function with an entry block
func newFunction(module *Module, name string, returnType types.Type) *Function {
	f := Function{
		module:     module,
		name:       name,
		returnType: returnType,
		args:       []types.Arg{},
		blocks:     []*Block{},
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

// Llvm returns the llvm ir representation of the function
func (f *Function) Llvm() string {
	argString := ""
	for i, a := range f.args {
		argString += a.Llvm()
		if i < len(f.args)-1 {
			argString += ", "
		}
	}

	s := fmt.Sprintf("define %s @%s(%s){\n",
		f.Type().(types.FunctionType).ReturnType().String(), f.name, argString)

	for _, b := range f.blocks {
		s += b.Llvm()
	}

	s += "}"
	return s
}

// Ident returns the function identifier
func (f *Function) Ident() string { return "@" + f.name }

// Type returns the function type
func (f *Function) Type() types.Type {
	var args []types.Type
	for _, a := range f.args {
		args = append(args, a.Type())
	}
	return types.NewFunctionType(f.returnType, args...)
}

// Arguments returns the values of function arguments
func (f *Function) Arguments() []types.Arg { return f.args }

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
