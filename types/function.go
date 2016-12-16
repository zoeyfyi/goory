package types

import "fmt"

type Function struct {
	returnType Type
	argTypes   []Type
}

func NewFunction(returnType Type, argTypes ...Type) Function {
	return Function{returnType, argTypes}
}

func (t Function) String() string {
	args := ""
	for i, a := range t.argTypes {
		args += a.String()

		if i < len(t.argTypes)-1 {
			args += ", "
		}
	}

	return fmt.Sprintf("%s (%s)", t.returnType.String(), args)
}

func (t Function) Arguments() []Type {
	return t.argTypes
}

func (t Function) ReturnType() Type {
	return t.returnType
}

func (t Function) Equal(n Type) bool {
	if nf, ok := n.(Function); ok {
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
