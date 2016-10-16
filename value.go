package goory

// Value represents a parameter or the result of a function
type Value struct {
	t    Type
	name string
}

// Name returns the name of the value
func (v *Value) Name() string {
	return v.name
}
