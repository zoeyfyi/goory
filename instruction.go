package goory

import "fmt"

const (
	instructionFadd = iota
	instructionFsub
	instructionFmul
	instructionFdiv

	instructionRet
	instructionCall
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
func newInstruction(id int, name string, operands ...Value) *Instruction {
	// Get the type the instruction returns
	var t Type
	switch id {
	case instructionFadd, instructionFdiv, instructionFmul, instructionFsub:
		if operands[0].Type() != operands[1].Type() {
			panic("Operators of diffrent types")
		}

		t = operands[0].Type()
	case instructionCall:
		t = operands[0].Type()
	case instructionRet:
		t = NilType
	case instructionBr:
		t = NilType
	case instructionCondBr:
		t = NilType
	default:
		panic("Unkown instruction id")
	}

	return &Instruction{
		id:       id,
		operands: operands,
		value:    newName(t, name),
	}
}

// String returns the string of the instruction type
func (i *Instruction) String() string {
	switch i.id {
	case instructionFadd:
		return "fadd"
	case instructionFdiv:
		return "fdiv"
	case instructionFmul:
		return "fmul"
	case instructionFsub:
		return "fsub"
	case instructionRet:
		return "ret"
	case instructionCall:
		return "call"
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

// Type returns the type of the instruction
func (i *Instruction) Type() Type {
	return i.value.Type()
}

// llvm compiles the instruction to llvm ir
func (i *Instruction) llvm() string {
	switch i.id {
	case instructionFadd:
		return fmt.Sprintf("%s = fadd %s %s, %s", i.Value().llvm(), i.Type().LLVMType(), i.operands[0].llvm(), i.operands[1].llvm())
	case instructionFdiv:
		return fmt.Sprintf("%s = fdiv %s %s, %s", i.Value().llvm(), i.Type().LLVMType(), i.operands[0].llvm(), i.operands[1].llvm())
	case instructionFmul:
		return fmt.Sprintf("%s = fmul %s %s, %s", i.Value().llvm(), i.Type().LLVMType(), i.operands[0].llvm(), i.operands[1].llvm())
	case instructionFsub:
		return fmt.Sprintf("%s = fsub %s %s, %s", i.Value().llvm(), i.Type().LLVMType(), i.operands[0].llvm(), i.operands[1].llvm())

	case instructionRet:
		return fmt.Sprintf("ret %s %s", i.operands[0].Type().LLVMType(), i.operands[0].llvm())
	case instructionCall:
		arguments := ""
		for op := 1; op < len(i.operands); op++ {
			arguments += fmt.Sprintf("%s %s", i.operands[op].Type().LLVMType(), i.operands[op].llvm())
			if op < len(i.operands)-1 {
				arguments += ", "
			}
		}
		return fmt.Sprintf("call %s %s(%s)", i.operands[0].Type().LLVMType(), i.operands[0].llvm(), arguments)
	case instructionBr:
		return fmt.Sprintf("br label %s", i.operands[0].llvm())
	case instructionCondBr:
		return fmt.Sprintf("br i1 %s, label %s, label %s", i.operands[0].llvm(), i.operands[1].llvm(), i.operands[2].llvm())
	default:
		panic("Unkown instruction id")
	}
}
