package goory

// Value represents a parameter or the result of a function
type Value struct {
	t    Type
	name string
}

func newValue(t Type, name string) Value {
	return Value{t, name}
}

// Name returns the name of the value
func (v *Value) Name() string {
	return v.name
}

// Type returns the type of the value
func (v *Value) Type() Type {
	return v.t
}
