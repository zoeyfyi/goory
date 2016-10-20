package goory

// Function is a group of instructions that are executed is a new stack frame
type Function struct {
	module     *Module
	name       string
	returnType Type
	argTypes   []Type
	args       []Value
	blocks     []Block
}

// NewFunction creates a new function with an entry block
func newFunction(module *Module, name string, returnType Type, argTypes ...Type) Function {
	// Turn types into values
	args := make([]Value, len(argTypes))
	for i, a := range argTypes {
		args[i] = newValue(a, module.nextTempName())
	}

	f := Function{
		module:     module,
		name:       name,
		returnType: returnType,
		argTypes:   argTypes,
		args:       args,
		blocks:     []Block{},
	}

	// Create entry block
	b := newBlock(&f, "entry")
	f.blocks = append(f.blocks, b)
	return f
}

// Module returns the module the function is in
func (f *Function) Module() *Module {
	return f.module
}

// Name returns the function name
func (f *Function) Name() string {
	return f.name
}

// Type returns the function type
func (f *Function) Type() (Type, []Type) {
	return f.returnType, f.argTypes
}

// AddBlock adds a new block to the function
func (f *Function) AddBlock(name string) *Block {
	b := Block{
		name: name,
	}

	f.blocks = append(f.blocks, b)
	return &b
}

// Parameters returns the values of function parameters
func (f *Function) Parameters() []Value {
	return f.args
}

// Entry returns the entry block of the function
func (f *Function) Entry() *Block {
	// Entry will (almost?) always be on top so loop is no cost
	for i, b := range f.blocks {
		if b.name == "entry" {
			return &f.blocks[i]
		}
	}

	panic("Function has no entry block, cant find entry block")
}
