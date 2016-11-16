package types

type Pointer struct {
	base Type
}

// NewPointerType creates a new pointer type based on the base type
func NewPointerType(base Type) Pointer {
	return Pointer{base}
}

func (t Pointer) String() string {
	return t.base.String() + "*"
}

func (t Pointer) Equal(n Type) bool {
	nPtr, ok := n.(Pointer)
	if !ok {
		return false
	}

	return nPtr.base == t.base
}

func (t Pointer) Base() Type {
	return t.base
}
