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

// Value gets the value of the block
func (b *Block) Value() Value {
	return newValue(NilType, b.name)
}

// Fadd creates a new float addition between left and right
func (b *Block) Fadd(left Value, right Value) Instruction {
	i := newInstruction(instructionFadd, b.function.module.nextTempName(), left, right)
	b.instructions = append(b.instructions, i)

	return i
}

// Ret creates a new return for ret
func (b *Block) Ret(ret Value) Instruction {
	i := newInstruction(instructionRet, b.function.module.nextTempName(), ret)
	b.instructions = append(b.instructions, i)

	return i
}

// Br creates a branch instruction to block
func (b *Block) Br(block Block) Instruction {
	i := newInstruction(instructionBr, b.function.module.nextTempName(), block.Value())
	b.instructions = append(b.instructions, i)

	return i
}

// CondBr creates a conditional branch based on the value on condition
func (b *Block) CondBr(condition Value, trueBlock Block, falseBlock Block) Instruction {
	i := newInstruction(instructionCondBr, b.function.module.nextTempName(), condition, trueBlock.Value(), falseBlock.Value())
	b.instructions = append(b.instructions, i)

	return i
}
