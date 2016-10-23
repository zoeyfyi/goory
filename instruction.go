package goory

import "fmt"

const (
	instructionFadd = iota
	instructionRet
	instructionBr
	instructionCondBr
)

// Instruction is an operation that is executed on its operands
type Instruction struct {
	id       int
	operands []Value
	value    Value
}

// newInstruction creates a new instruction
func newInstruction(id int, name string, operands ...Value) Instruction {
	// Get the type the instruction returns
	var t Type
	switch id {
	case instructionFadd:
		if operands[0].t != operands[1].t {
			panic("Operators of diffrent types")
		}

		t = operands[0].t
	case instructionRet:
		t = NilType
	case instructionBr:
		t = NilType
	case instructionCondBr:
		t = NilType
	default:
		panic("Unkown instruction id")
	}

	return Instruction{
		id:       id,
		operands: operands,
		value:    newValue(t, name),
	}
}

// String returns the string of the instruction type
func (i *Instruction) String() string {
	switch i.id {
	case instructionFadd:
		return "fadd"
	case instructionRet:
		return "ret"
	case instructionBr:
		return "br"
	case instructionCondBr:
		return "br"
	default:
		panic("Unkown instruction id")
	}
}

// Value returns the value of the instruction
func (i *Instruction) Value() Value {
	return i.value
}

// llvm compiles the instruction to llvm ir
func (i *Instruction) llvm() string {
	switch i.id {
	case instructionFadd:
		return fmt.Sprintf("%%%s = fadd %s %%%s, %%%s", i.value.name, i.value.t.LLVMType(), i.operands[0].name, i.operands[1].name)
	case instructionRet:
		return fmt.Sprintf("ret %s %%%s", i.operands[0].t.LLVMType(), i.operands[0].name)
	case instructionBr:
		return fmt.Sprintf("br label %%%s", i.operands[0].name)
	case instructionCondBr:
		return fmt.Sprintf("br i1 %%%s, label %%%s, label %%%s", i.operands[0].name, i.operands[1].name, i.operands[2].name)
	default:
		panic("Unkown instruction id")
	}
}
