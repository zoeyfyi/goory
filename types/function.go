package types

type FunctionType struct {
	returnType Type
	argTypes   []Type
}

func NewFunctionType(returnType Type, argTypes ...Type) Type {
	return FunctionType{returnType, argTypes}
}

func (t FunctionType) String() string {
	return "function type"
}

func (t FunctionType) Arguments() []Type {
	return t.argTypes
}

func (t FunctionType) ReturnType() Type {
	return t.returnType
}

func (t FunctionType) Equal(n Type) bool {
	if nf, ok := n.(FunctionType); ok {
		if !nf.returnType.Equal(t.returnType) {
			return false
		}

		if len(t.argTypes) != len(nf.argTypes) {
			return false
		}

		for i := 0; i < len(t.argTypes); i++ {
			if !t.argTypes[i].Equal(nf.argTypes[i]) {
				return false
			}
		}

		return true
	}

	return false
}
