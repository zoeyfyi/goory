package goory

import "fmt"

const (
	instructionFadd = iota
	instructionFsub
	instructionFmul
	instructionFdiv

	instructionAdd
	instructionSub
	instructionMul
	instructionDiv

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
	case instructionFadd, instructionFdiv, instructionFmul, instructionFsub,
		instructionAdd, instructionSub, instructionMul, instructionDiv:

		// Check their are two operands
		if operands == nil || len(operands) != 2 {
			panic("Expecting two operands for maths instructions")
		}

		// Check the operand types
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
	// Float operations
	case instructionFadd:
		return "fadd"
	case instructionFdiv:
		return "fdiv"
	case instructionFmul:
		return "fmul"
	case instructionFsub:
		return "fsub"
	// Interget operations
	case instructionAdd:
		return "add"
	case instructionSub:
		return "sub"
	case instructionMul:
		return "mul"
	case instructionDiv:
		return "div"
	// Control flow
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
	// Math instructions
	case instructionFadd, instructionFsub, instructionFmul, instructionFdiv,
		instructionAdd, instructionSub, instructionMul, instructionDiv:

		return fmt.Sprintf("%s = %s %s %s, %s", i.Value().llvm(), i.String(), i.Type().LLVMType(), i.operands[0].llvm(), i.operands[1].llvm())

	// Control flow instructions
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
		return fmt.Sprintf("%s = call %s %s(%s)", i.Value().llvm(), i.operands[0].Type().LLVMType(), i.operands[0].llvm(), arguments)
	case instructionBr:
		return fmt.Sprintf("br label %s", i.operands[0].llvm())
	case instructionCondBr:
		return fmt.Sprintf("br i1 %s, label %s, label %s", i.operands[0].llvm(), i.operands[1].llvm(), i.operands[2].llvm())
	default:
		panic("Unkown instruction id")
	}
}
