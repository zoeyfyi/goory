package goory

import (
	"github.com/bongo227/goory/instructions"
	"github.com/bongo227/goory/types"
	"github.com/bongo227/goory/value"
)

type nextName func() string

// Block is a seqential list of instructions
type Block struct {
	function     *Function
	name         string
	instructions []instructions.Instruction
	nextName     nextName
}

func newBlock(function *Function, name string) *Block {
	return &Block{
		function: function,
		name:     name,
		nextName: function.module.nextTempName,
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

// String gets the name of the block
func (b *Block) String() string {
	return b.name
}

func (b *Block) Type() types.Type {
	return types.NewBlockType()
}

// Ident gets the identifyer of the block
func (b *Block) Ident() string {
	return b.name
}

// Terminated returns true if the block is terminated (ends in branch or return)
func (b *Block) Terminated() bool {
	return b.instructions[len(b.instructions)-1].IsTerminator()
}

func (b *Block) IsGlobal() bool {
	return false
}

func (b *Block) IsBlock() bool {
	return true
}

// Add creates a new add instruction
func (b *Block) Add(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewAdd(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// And create a new and instruction
func (b *Block) And(lhs, rhs value.Value) value.Value {
	i := instructions.NewAnd(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Br creates a new br instruction
func (b *Block) Br(block value.Value) value.Value {
	i := instructions.NewBr(b.nextName(), block)
	b.instructions = append(b.instructions, i)
	return i
}

// Call creates a new call instruction
func (b *Block) Call(function value.Value, operands ...value.Value) value.Value {
	i := instructions.NewCall(b.nextName(), function, operands...)
	b.instructions = append(b.instructions, i)
	return i
}

// CondBr creates a new conditional branch instruction
func (b *Block) CondBr(condition, trueBlock, falseBlock value.Value) value.Value {
	i := instructions.NewCondBr(b.nextName(), condition, trueBlock, falseBlock)
	b.instructions = append(b.instructions, i)
	return i
}

// Div creates a new integer division instruction
func (b *Block) Div(lhs, rhs value.Value) value.Value {
	i := instructions.NewDiv(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Fadd creates a new float add instruction
func (b *Block) Fadd(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewFadd(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Fcmp creates a new float add instruction
func (b *Block) Fcmp(mode string, lhs, rhs value.Value) value.Value {
	i := instructions.NewFcmp(b.nextName(), mode, lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Fdiv creates a new float add instruction
func (b *Block) Fdiv(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewFdiv(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Fmul creates a new float add instruction
func (b *Block) Fmul(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewFmul(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Fpext creates a new float add instruction
func (b *Block) Fpext(value value.Value, cast types.Type) value.Value {
	i := instructions.NewFpext(b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Fptosi creates a new float add instruction
func (b *Block) Fptosi(value value.Value, cast types.Type) value.Value {
	i := instructions.NewFptosi(b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Fptrunc creates a new float add instruction
func (b *Block) Fptrunc(value value.Value, cast types.Type) value.Value {
	i := instructions.NewFptrunc(b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Fsub creates a new float add instruction
func (b *Block) Fsub(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewFsub(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

var (
	// IntEq is an integer equals comparison
	IntEq = "eq"

	// IntUlt is an integer unordered less than comparison
	IntUlt = "ult"
)

// Icmp creates a new float add instruction
func (b *Block) Icmp(mode string, lhs, rhs value.Value) value.Value {
	i := instructions.NewIcmp(b.nextName(), mode, lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Mul creates a new float add instruction
func (b *Block) Mul(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewMul(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Or creates a new float add instruction
func (b *Block) Or(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewOr(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Ret creates a new return instruction
func (b *Block) Ret(value value.Value) {
	i := instructions.NewRet(b.nextName(), value)
	b.instructions = append(b.instructions, i)
}

// Sitofp creates a new float add instruction
func (b *Block) Sitofp(value value.Value, cast types.Type) value.Value {
	i := instructions.NewSitofp(b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Sub creates a new float add instruction
func (b *Block) Sub(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewSub(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Trunc creates a new float add instruction
func (b *Block) Trunc(value value.Value, cast types.Type) value.Value {
	i := instructions.NewTrunc(b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Xor creates a new float add instruction
func (b *Block) Xor(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewXor(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Zext creates a new float add instruction
func (b *Block) Zext(value value.Value, cast types.Type) value.Value {
	i := instructions.NewZext(b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// // Cast creates a cast for value to type cast
// func (b *Block) Cast(value Value, cast Atomic) Instruction {
// 	atomic, isAtomic := value.Type().(Atomic)

// 	if !isAtomic {
// 		panic("Value is not an atomic type")
// 	}

// 	// Value type is the same as cast type so return an empty instruction
// 	if atomic == cast {
// 		return &none{value}
// 	}

// 	// Cast integer to integer
// 	if atomic.IsInteger() && cast.IsInteger() {
// 		if Compare(atomic, cast) == 1 {
// 			// Cast is bigger so expand value
// 			i := &zext{b.nextName(), value, cast}
// 			b.instructions = append(b.instructions, i)
// 			return i
// 		}
// 		// Cast is smaller so truncate value
// 		i := &trunc{b.nextName(), value, cast}
// 		b.instructions = append(b.instructions, i)
// 		return i
// 	}

// 	// Cast float to float
// 	if atomic.IsFloat() && cast.IsFloat() {
// 		if Compare(atomic, cast) == 1 {
// 			// Cast is bigger so expand value
// 			i := &fpext{b.nextName(), value, cast}
// 			b.instructions = append(b.instructions, i)
// 			return i
// 		}
// 		// Cast is smaller so truncate value
// 		i := &fptrunc{b.nextName(), value, cast}
// 		b.instructions = append(b.instructions, i)
// 		return i
// 	}

// 	// Cast integer to float
// 	if atomic.IsInteger() && cast.IsFloat() {
// 		i := &sitofp{b.nextName(), value, cast}
// 		b.instructions = append(b.instructions, i)
// 		return i
// 	}

// 	// Cast float to integer
// 	if atomic.IsFloat() && cast.IsInteger() {
// 		i := &fptosi{b.nextName(), value, cast}
// 		b.instructions = append(b.instructions, i)
// 		return i
// 	}

// 	panic("Cannot cast non integer or pointer value")
// }
