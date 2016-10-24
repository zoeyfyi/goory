package goory

// Block is a seqential list of instructions
type Block struct {
	function     *Function
	name         string
	instructions []*Instruction
}

func newBlock(function *Function, name string) *Block {
	return &Block{
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
	return newName(NilType, b.name)
}

// Fadd creates a new float addition between left and right
func (b *Block) Fadd(left Value, right Value) *Instruction {
	i := newInstruction(instructionFadd, b.function.module.nextTempName(), left, right)
	b.instructions = append(b.instructions, i)

	return i
}

// Fdiv creates a new float division between left and right
func (b *Block) Fdiv(left Value, right Value) *Instruction {
	i := newInstruction(instructionFdiv, b.function.module.nextTempName(), left, right)
	b.instructions = append(b.instructions, i)

	return i
}

// Fmul creates a new float addition between left and right
func (b *Block) Fmul(left Value, right Value) *Instruction {
	i := newInstruction(instructionFmul, b.function.module.nextTempName(), left, right)
	b.instructions = append(b.instructions, i)

	return i
}

// Fsub creates a new float subtraction between left and right
func (b *Block) Fsub(left Value, right Value) *Instruction {
	i := newInstruction(instructionFsub, b.function.module.nextTempName(), left, right)
	b.instructions = append(b.instructions, i)

	return i
}

// Ret creates a new return for ret
func (b *Block) Ret(ret Value) *Instruction {
	i := newInstruction(instructionRet, b.function.module.nextTempName(), ret)
	b.instructions = append(b.instructions, i)

	return i
}

// Call creates a call instruction with the arguments
func (b *Block) Call(function *Function, arguments ...Value) *Instruction {
	operands := []Value{function.value}
	operands = append(operands, arguments...)
	i := newInstruction(instructionCall, b.function.module.nextTempName(), operands...)
	b.instructions = append(b.instructions, i)

	return i
}

// Br creates a branch instruction to block
func (b *Block) Br(block *Block) *Instruction {
	i := newInstruction(instructionBr, b.function.module.nextTempName(), block.Value())
	b.instructions = append(b.instructions, i)

	return i
}

// CondBr creates a conditional branch based on the value on condition
func (b *Block) CondBr(condition Value, trueBlock *Block, falseBlock *Block) *Instruction {
	i := newInstruction(instructionCondBr, b.function.module.nextTempName(), condition, trueBlock.Value(), falseBlock.Value())
	b.instructions = append(b.instructions, i)

	return i
}
