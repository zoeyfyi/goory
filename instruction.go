package goory

import "fmt"

const (
	instructionFadd = iota
	instructionRet
)

// Instruction is an operation that is executed on its operands
type Instruction struct {
	id       int
	operands []Value
	name     string
}

// String returns the string of the instruction type
func (i *Instruction) String() string {
	switch i.id {
	case instructionFadd:
		return "fadd"
	default:
		panic("Unkown instruction id")
	}
}

// Name returns the name the instruction is assigned
func (i *Instruction) Name() string {
	return i.name
}

// SetName sets the name the instruction is assigned
func (i *Instruction) SetName(name string) {
	i.name = name
}

// Type gets the type the instruction returns
func (i *Instruction) Type() Type {
	switch i.id {
	case instructionFadd:
		// Both operands should be the same, so return first
		return i.operands[0].t
	case instructionRet:
		return NilType
	default:
		panic("Unkown instruction id")
	}
}

func (i *Instruction) llvm() string {
	switch i.id {
	case instructionFadd:
		i.Type()
		return fmt.Sprintf("%s = fadd %s %s, %s", i.name, i.Type().LLVMType(), i.operands[0].name, i.operands[1].name)
	case instructionRet:
		return fmt.Sprintf("ret %s %s", i.operands[0].t.LLVMType(), i.operands[0].name)
	default:
		panic("Unkown instruction id")
	}
}
