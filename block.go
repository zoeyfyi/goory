package goory

// Block is a seqential list of instructions
type Block struct {
	function     *Function
	name         string
	instructions []Instruction
}

func newBlock(function *Function, name string) Block {
	return Block{
		function: function,
		name:     name,
	}
}

// Function returns the parent function
func (b *Block) Function() *Function {
	return b.function
}

// Name gets the name of the block
func (b *Block) Name() string {
	return b.name
}

// Fadd creates a new float addition between left and right
func (b *Block) Fadd(left Value, right Value) Instruction {
	i := Instruction{
		id:       instructionFadd,
		operands: []Value{left, right},
		name:     b.function.module.nextTempName(),
	}

	b.instructions = append(b.instructions, i)

	return i
}

// Ret creates a new return for ret
func (b *Block) Ret(ret Value) Instruction {
	i := Instruction{
		id:       instructionRet,
		operands: []Value{ret},
		name:     b.function.module.nextTempName(),
	}

	b.instructions = append(b.instructions, i)

	return i
}
