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

// Terminated returns true if the block is terminated (ends in branch or return)
func (b *Block) Terminated() bool {
	return b.instructions[len(b.instructions)-1].IsTerminator()
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

// Fmul creates a new float multiplication between left and right
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

// Sub creates a new interger subtraction between left and right
func (b *Block) Sub(left Value, right Value) *Instruction {
	i := newInstruction(instructionSub, b.function.module.nextTempName(), left, right)
	b.instructions = append(b.instructions, i)

	return i
}

// Add creates a new interger addition between left and right
func (b *Block) Add(left Value, right Value) *Instruction {
	i := newInstruction(instructionAdd, b.function.module.nextTempName(), left, right)
	b.instructions = append(b.instructions, i)

	return i
}

// Mul creates a new interger multiplication between left and right
func (b *Block) Mul(left Value, right Value) *Instruction {
	i := newInstruction(instructionMul, b.function.module.nextTempName(), left, right)
	b.instructions = append(b.instructions, i)

	return i
}

// Div creates a new interger division between left and right
func (b *Block) Div(left Value, right Value) *Instruction {
	i := newInstruction(instructionDiv, b.function.module.nextTempName(), left, right)
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
