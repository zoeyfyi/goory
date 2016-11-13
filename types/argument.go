package types

import "fmt"

// Arg represents a function argument
type Arg struct {
	argType Type
	name    string
}

func NewArg(argType Type, name string) Arg {
	return Arg{argType, name}
}

func (v Arg) Type() Type {
	return v.argType
}

func (v Arg) Llvm() string {
	return fmt.Sprintf("%s %%%s", v.argType.String(), v.name)
}

func (v Arg) Ident() string {
	return "%" + v.name
}
