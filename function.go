package goory

// Function is a group of instructions that are executed is a new stack frame
type Function struct {
	name   string
	fType  FunctionType
	blocks []Block
}

// NewFunction creates a new function with an entry block
func NewFunction(name string, fType *FunctionType) (*Function, *Block) {
	b := Block{
		name: "entry",
	}

	return &Function{
		name:   name,
		fType:  *fType,
		blocks: []Block{b},
	}, &b
}

// Name returns the function name
func (f *Function) Name() string {
	return f.name
}

// Type returns the function type
func (f *Function) Type() *FunctionType {
	return &f.fType
}

// AddBlock adds a new block to the function
func (f *Function) AddBlock(name string) *Block {
	b := Block{
		name: name,
	}

	f.blocks = append(f.blocks, b)
	return &b
}

// Entry returns the entry block of the function
func (f *Function) Entry() *Block {
	// Entry will (almost?) always be on top so loop is no cost
	for _, b := range f.blocks {
		if b.name == "entry" {
			return &b
		}
	}

	panic("Function has no entry block, cant find entry block")
}
