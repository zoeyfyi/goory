package types

type Void struct{}

// NewVoid creates a new void type
func NewVoid() Void {
	return Void{}
}

// String returns the llvm ir of the void type
func (t Void) String() string {
	return "void"
}

// Equal checks if type a is equal to void
func (t Void) Equal(a Type) bool {
	_, ok := a.(Void)
	return ok
}
