package goory

import "fmt"

// Block is a seqential list of instructions
type Block struct {
	name         string
	instructions []Instruction
	tempName     int
}

func (b *Block) nextTempName() string {
	s := fmt.Sprintf("%d", b.tempName)
	b.tempName++
	return s
}

// Name gets the name of the block
func (b *Block) Name() string {
	return b.name
}

// Fadd creates a new float addition between left and right
func (b *Block) Fadd(left Value, right Value) (*Instruction, Type) {
	i := Instruction{
		id:       instructionFadd,
		operands: []Value{left, right},
		name:     b.nextTempName(),
	}

	b.instructions = append(b.instructions, i)

	return &i, i.Type()
}

// Ret creates a new return for ret
func (b *Block) Ret(ret Value) *Instruction {
	i := Instruction{
		id:       instructionRet,
		operands: []Value{ret},
		name:     b.nextTempName(),
	}

	b.instructions = append(b.instructions, i)

	return &i
}
