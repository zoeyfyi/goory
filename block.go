package goory

import (
	"github.com/bongo227/goory/instructions"
	"github.com/bongo227/goory/types"
	"github.com/bongo227/goory/value"
)

type nextName func() string

// Block is a seqential list of instructions ending with a terminator instruction
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

// Type returns new block type
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

// Add creates a new add instruction.
// lhs and rhs must be integer types of the same size.
// Add returns the result of this instruction with the same type as lhs and rhs.
func (b *Block) Add(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewAdd(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// And create a new and instruction.
// lhs and rhs must be boolean types.
// And returns the result of the instruction as a boolean type.
func (b *Block) And(lhs, rhs value.Value) value.Value {
	i := instructions.NewAnd(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Br creates a new br instruction.
// block must be a block type.
func (b *Block) Br(block value.Value) {
	i := instructions.NewBr(b.nextName(), block)
	b.instructions = append(b.instructions, i)
}

// Call creates a new call instruction.
// function must be a function type.
// operands must match the types of the function arguments.
// Call returns the result of the instruction with the same type as the function return type.
func (b *Block) Call(function value.Value, operands ...value.Value) value.Value {
	i := instructions.NewCall(b.nextName(), function, operands...)
	b.instructions = append(b.instructions, i)
	return i
}

// CondBr creates a new conditional branch instruction.
// condition must be a boolean type.
// trueBlock and falseBlock must be block types.
func (b *Block) CondBr(condition, trueBlock, falseBlock value.Value) {
	i := instructions.NewCondBr(b.nextName(), condition, trueBlock, falseBlock)
	b.instructions = append(b.instructions, i)
}

// Div creates a new integer division instruction.
// lhs and rhs must be integer types.
// Div returns the result of the instruction with the same type as lhs and rhs.
func (b *Block) Div(lhs, rhs value.Value) value.Value {
	i := instructions.NewDiv(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Fadd creates a new float add instruction.
// lhs and rhs must be float or double types.
// Fadd returns the result of the instruction with the same type as lhs and rhs.
func (b *Block) Fadd(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewFadd(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Fcmp creates a new float add instruction.
// mode controls the behavior of the comparison, see https://godoc.org/github.com/bongo227/goory#pkg-variables.
// lhs and rhs must be float or double types.
// Fcmp returns the result of the instruction as a boolean type.
func (b *Block) Fcmp(mode string, lhs, rhs value.Value) value.Value {
	i := instructions.NewFcmp(b.nextName(), mode, lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Fdiv creates a new float add instruction.
// lhs and rhs must be float or double types.
// Fdiv returns the result of the instruction with the same type as lhs and rhs.
func (b *Block) Fdiv(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewFdiv(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Fmul creates a new float add instruction.
// lhs and rhs must be float or double types.
// Fmul returns the result of the instruction with the same type as lhs and rhs.
func (b *Block) Fmul(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewFmul(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Fpext creates a new float extention instruction.
// value must be a float or double type.
// cast must be a type larger than value.
// Fpext returns the result of the instruction with the same type as cast.
func (b *Block) Fpext(value value.Value, cast types.Type) value.Value {
	// TODO: add support for fp128 etc
	i := instructions.NewFpext(b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Fptosi creates a new float to signed integer type.
// value must be a float of double type.
// cast must be an integer type.
// Fptosi returns the result of the instruction with the same type as cast.
func (b *Block) Fptosi(value value.Value, cast types.Type) value.Value {
	i := instructions.NewFptosi(b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Fptrunc creates a new float truncation.
// value must be a float or double type.
// cast must be a type smaller than value.
// Fptrunc returns the result of the instruction with the same type as cast.
func (b *Block) Fptrunc(value value.Value, cast types.Type) value.Value {
	i := instructions.NewFptrunc(b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Fsub creates a new float subtraction instruction.
// lhs and rhs must be float or double types.
// Fsub returns the result of the instruction with the same type as lhs and rhs.
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

// Icmp creates a new integer comparison instruction.
// mode controls the behavior of the comparison, see: https://godoc.org/github.com/bongo227/goory#pkg-variables.
// lhs and rhs must be integer types of the the same size.
// Icmp returns the result of the instruction as a boolean type.
func (b *Block) Icmp(mode string, lhs, rhs value.Value) value.Value {
	i := instructions.NewIcmp(b.nextName(), mode, lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Mul creates a new integer multiplication instruction.
// lhs and rhs must be integer types of the the same size.
// Mul returns the result of the instruction with the same type as lhs and rhs.
func (b *Block) Mul(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewMul(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Or creates a new bitwise or instruction.
// lhs and rhs must be boolean types.
// Or returns the result of the instruction as a boolean type.
func (b *Block) Or(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewOr(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Ret creates a new return instruction.
// value must be the same type as the function return type.
func (b *Block) Ret(value value.Value) {
	i := instructions.NewRet(b.nextName(), value)
	b.instructions = append(b.instructions, i)
}

// Sitofp creates a new signed integer to float instruction.
// value must be an integer type.
// cast must be a float type.
// Sitofp returns the result of the instruction with the same type as cast.
func (b *Block) Sitofp(value value.Value, cast types.Type) value.Value {
	i := instructions.NewSitofp(b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Sub creates a new float sub instruction.
// lhs and rhs must be float or double types.
// Sub retruns the result of the instruction with the same type as lhs and rhs.
func (b *Block) Sub(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewSub(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Trunc creates a new integer truncation instruction.
// value must be an integer type.
// cast must be a integer type smaller than value.
// Trunc returns the result of the instruction with the same type as cast.
func (b *Block) Trunc(value value.Value, cast types.Type) value.Value {
	i := instructions.NewTrunc(b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Xor creates a new float exclusive or instruction.
// lhs and rhs must be boolean types.
// Xor returns the result of the instruction as a boolean type.
func (b *Block) Xor(lhs value.Value, rhs value.Value) value.Value {
	i := instructions.NewXor(b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Zext creates a new zero extention instruction.
// value must be an integer type.
// cast must be a integer type larger than value.
// Zext returns the result of the instruction with the same type as cast.
func (b *Block) Zext(value value.Value, cast types.Type) value.Value {
	i := instructions.NewZext(b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}
