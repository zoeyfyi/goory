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
	instructions []value.Instruction
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

// Llvm returns the llvm ir representation of the block
func (b *Block) Llvm() string {
	s := "\t" + b.name + ":\n"
	for _, i := range b.instructions {
		s += "\t\t" + i.Llvm() + "\n"
	}

	return s
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

// Add creates a new add instruction.
// lhs and rhs must be integer types of the same size.
// Add returns the result of this instruction with the same type as lhs and rhs.
func (b *Block) Add(lhs value.Value, rhs value.Value) *instructions.Add {
	i := instructions.NewAdd(b, b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Alloca creates a new alloca instruction.
// type is the type you want to allocate space for.
// Call SetNumber to set the number of elements to allocate for (default 1)
// Alloca returns the result of this instruction as a pointer with the same type as base type.
func (b *Block) Alloca(base types.Type) *instructions.Alloca {
	i := instructions.NewAlloca(b, b.nextName(), base)
	b.instructions = append(b.instructions, i)
	return i
}

// And create a new and instruction.
// lhs and rhs must be boolean types.
// And returns the result of the instruction as a boolean type.
func (b *Block) And(lhs, rhs value.Value) *instructions.And {
	i := instructions.NewAnd(b, b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Br creates a new br instruction.
// block must be a block type.
func (b *Block) Br(block value.Value) *instructions.Br {
	i := instructions.NewBr(b, b.nextName(), block)
	b.instructions = append(b.instructions, i)
	return i
}

// Call creates a new call instruction.
// function must be a function type.
// operands must match the types of the function arguments.
// Call returns the result of the instruction with the same type as the function return type.
func (b *Block) Call(function value.Value, operands ...value.Value) *instructions.Call {
	i := instructions.NewCall(b, b.nextName(), function, operands...)
	b.instructions = append(b.instructions, i)
	return i
}

// CondBr creates a new conditional branch instruction.
// condition must be a boolean type.
// trueBlock and falseBlock must be block types.
func (b *Block) CondBr(condition, trueBlock, falseBlock value.Value) *instructions.CondBr {
	i := instructions.NewCondBr(b, b.nextName(), condition, trueBlock, falseBlock)
	b.instructions = append(b.instructions, i)
	return i
}

// Div creates a new integer division instruction.
// lhs and rhs must be integer types.
// Div returns the result of the instruction with the same type as lhs and rhs.
func (b *Block) Div(lhs, rhs value.Value) *instructions.Div {
	i := instructions.NewDiv(b, b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Extractvalue creates a new extract value instruction
// location must be an aggregate type
// position is the index at which to extract the value from location
// Extractvalue returns the type specified in location at index position
func (b *Block) Extractvalue(location value.Value, position value.Value) *instructions.Extractvalue {
	i := instructions.NewExtractvalue(b, b.nextName(), location, position)
	b.instructions = append(b.instructions, i)
	return i
}

// Fadd creates a new float add instruction.
// lhs and rhs must be float or double types.
// Fadd returns the result of the instruction with the same type as lhs and rhs.
func (b *Block) Fadd(lhs value.Value, rhs value.Value) *instructions.Fadd {
	i := instructions.NewFadd(b, b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

var (
	// FloatOeq is a float ordered equal comparison
	FloatOeq = "oeq"

	// FloatOlt is a float ordered less than comparison
	FloatOlt = "olt"

	// Float One is a float ordered not equals comparison
	FloatOne = "one"

	// FloatOgt is a float ordered greater than comparison
	FloatOgt = "ogt"
)

// Fcmp creates a new float add instruction.
// mode controls the behavior of the comparison, see https://godoc.org/github.com/bongo227/goory#pkg-variables.
// lhs and rhs must be float or double types.
// Fcmp returns the result of the instruction as a boolean type.
func (b *Block) Fcmp(mode string, lhs, rhs value.Value) *instructions.Fcmp {
	i := instructions.NewFcmp(b, b.nextName(), mode, lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Fdiv creates a new float add instruction.
// lhs and rhs must be float or double types.
// Fdiv returns the result of the instruction with the same type as lhs and rhs.
func (b *Block) Fdiv(lhs value.Value, rhs value.Value) *instructions.Fdiv {
	i := instructions.NewFdiv(b, b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Fmul creates a new float add instruction.
// lhs and rhs must be float or double types.
// Fmul returns the result of the instruction with the same type as lhs and rhs.
func (b *Block) Fmul(lhs value.Value, rhs value.Value) *instructions.Fmul {
	i := instructions.NewFmul(b, b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Fpext creates a new float extention instruction.
// value must be a float or double type.
// cast must be a type larger than value.
// Fpext returns the result of the instruction with the same type as cast.
func (b *Block) Fpext(value value.Value, cast types.Type) *instructions.Fpext {
	// TODO: add support for fp128 etc
	i := instructions.NewFpext(b, b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Fptosi creates a new float to signed integer type.
// value must be a float of double type.
// cast must be an integer type.
// Fptosi returns the result of the instruction with the same type as cast.
func (b *Block) Fptosi(value value.Value, cast types.Type) *instructions.Fptosi {
	i := instructions.NewFptosi(b, b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Fptrunc creates a new float truncation.
// value must be a float or double type.
// cast must be a type smaller than value.
// Fptrunc returns the result of the instruction with the same type as cast.
func (b *Block) Fptrunc(value value.Value, cast types.Type) *instructions.Fptrunc {
	i := instructions.NewFptrunc(b, b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Fsub creates a new float subtraction instruction.
// lhs and rhs must be float or double types.
// Fsub returns the result of the instruction with the same type as lhs and rhs.
func (b *Block) Fsub(lhs value.Value, rhs value.Value) *instructions.Fsub {
	i := instructions.NewFsub(b, b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

var (
	// IntEq is an integer equals comparison
	IntEq = "eq"

	// IntNe is an integer not equals comparison
	IntNe = "ne"

	// IntUlt is an integer unsigned less than comparison
	IntUlt = "ult"

	// IntSlt is an integer signed less than comparison
	IntSlt = "slt"

	// IntUgt is an integer unsigned greater than comparison
	IntUgt = "ugt"

	// IntSgt is an integer signed greater than comparison
	IntSgt = "sgt"
)

// Insertvalue creates a new insert value instruction.
// location must be an aggregate type the value is to be inserted into.
// value must match the type of the location at index position.
// position spesifys the index in which to insert value in location.
// Insertvalue returns the result of the instruction with the same type as location.
func (b *Block) Insertvalue(location, value value.Value, position value.Value) *instructions.Insertvalue {
	i := instructions.NewInsertvalue(b, b.nextName(), location, value, position)
	b.instructions = append(b.instructions, i)
	return i
}

// Icmp creates a new integer comparison instruction.
// mode controls the behavior of the comparison, see: https://godoc.org/github.com/bongo227/goory#pkg-variables.
// lhs and rhs must be integer types of the the same size.
// Icmp returns the result of the instruction as a boolean type.
func (b *Block) Icmp(mode string, lhs, rhs value.Value) *instructions.Icmp {
	i := instructions.NewIcmp(b, b.nextName(), mode, lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Load creates a new load instruction.
// Load returns the result of the instruction with the same type as the allocation.
func (b *Block) Load(allocation *instructions.Alloca) *instructions.Load {
	i := instructions.NewLoad(b, b.nextName(), allocation)
	b.instructions = append(b.instructions, i)
	return i
}

// Mul creates a new integer multiplication instruction.
// lhs and rhs must be integer types of the the same size.
// Mul returns the result of the instruction with the same type as lhs and rhs.
func (b *Block) Mul(lhs value.Value, rhs value.Value) *instructions.Mul {
	i := instructions.NewMul(b, b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Or creates a new bitwise or instruction.
// lhs and rhs must be boolean types.
// Or returns the result of the instruction as a boolean type.
func (b *Block) Or(lhs value.Value, rhs value.Value) *instructions.Or {
	i := instructions.NewOr(b, b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Phi creates a new phi instruction.
// Phi returns the phi instruction to which you add all incoming blocks.
func (b *Block) Phi() *instructions.Phi {
	i := instructions.NewPhi(b, b.nextName())
	b.instructions = append(b.instructions, i)
	return i
}

// Ret creates a new return instruction.
// value must be the same type as the function return type.
func (b *Block) Ret(value value.Value) *instructions.Ret {
	i := instructions.NewRet(b, value)
	b.instructions = append(b.instructions, i)
	return i
}

// Sitofp creates a new signed integer to float instruction.
// value must be an integer type.
// cast must be a float type.
// Sitofp returns the result of the instruction with the same type as cast.
func (b *Block) Sitofp(value value.Value, cast types.Type) *instructions.Sitofp {
	i := instructions.NewSitofp(b, b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Srem creates a new srem instruction.
// lhs and rhs must be integer types of the same size.
// Srem returns the result of this instruction with the same type as lhs and rhs.
func (b *Block) Srem(lhs value.Value, rhs value.Value) *instructions.Srem {
	i := instructions.NewSrem(b, b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Store creates a new store instruction.
// value must be the same type as the allocation.
// Store retruns the result of the instruction with the same type as lhs and rhs.
func (b *Block) Store(allocation *instructions.Alloca, value value.Value) {
	i := instructions.NewStore(b, allocation, value)
	b.instructions = append(b.instructions, i)
}

// Sub creates a new float sub instruction.
// lhs and rhs must be float or double types.
// Sub retruns the result of the instruction with the same type as lhs and rhs.
func (b *Block) Sub(lhs value.Value, rhs value.Value) *instructions.Sub {
	i := instructions.NewSub(b, b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Trunc creates a new integer truncation instruction.
// value must be an integer type.
// cast must be a integer type smaller than value.
// Trunc returns the result of the instruction with the same type as cast.
func (b *Block) Trunc(value value.Value, cast types.Type) *instructions.Trunc {
	i := instructions.NewTrunc(b, b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Xor creates a new float exclusive or instruction.
// lhs and rhs must be boolean types.
// Xor returns the result of the instruction as a boolean type.
func (b *Block) Xor(lhs value.Value, rhs value.Value) *instructions.Xor {
	i := instructions.NewXor(b, b.nextName(), lhs, rhs)
	b.instructions = append(b.instructions, i)
	return i
}

// Zext creates a new zero extention instruction.
// value must be an integer type.
// cast must be a integer type larger than value.
// Zext returns the result of the instruction with the same type as cast.
func (b *Block) Zext(value value.Value, cast types.Type) *instructions.Zext {
	i := instructions.NewZext(b, b.nextName(), value, cast)
	b.instructions = append(b.instructions, i)
	return i
}

// Cast creates a cast for value to type cast
func (b *Block) Cast(value value.Value, cast types.Type) value.Value {
	atomicCast, isAtomicCast := cast.(types.Atomic)
	if !isAtomicCast {
		panic("Cast is not atomic")
	}

	atomic, isAtomic := value.Type().(types.Atomic)
	if !isAtomic {
		panic("Value is not an atomic type")
	}

	// Value type is the same as cast type so return the value
	if atomic.Equal(atomicCast) {
		return value
	}

	// Cast integer to integer
	if types.IsInteger(atomic) && types.IsInteger(atomicCast) {
		if types.Compare(atomic, atomicCast) == -1 {
			// Cast is bigger so expand value
			i := instructions.NewZext(b, b.nextName(), value, atomicCast)
			b.instructions = append(b.instructions, i)
			return i
		}
		// Cast is smaller so truncate value
		i := instructions.NewTrunc(b, b.nextName(), value, atomicCast)
		b.instructions = append(b.instructions, i)
		return i
	}

	// Cast float to float
	if types.IsFp(atomic) && types.IsFp(atomicCast) {
		if types.Compare(atomic, atomicCast) == -1 {
			// Cast is bigger so expand value
			i := instructions.NewFpext(b, b.nextName(), value, atomicCast)
			b.instructions = append(b.instructions, i)
			return i
		}
		// Cast is smaller so truncate value
		i := instructions.NewFptrunc(b, b.nextName(), value, atomicCast)
		b.instructions = append(b.instructions, i)
		return i
	}

	// Cast integer to float
	if types.IsInteger(atomic) && types.IsFp(atomicCast) {
		i := instructions.NewSitofp(b, b.nextName(), value, atomicCast)
		b.instructions = append(b.instructions, i)
		return i
	}

	// Cast float to integer
	if types.IsFp(atomic) && types.IsInteger(atomicCast) {
		i := instructions.NewFptosi(b, b.nextName(), value, atomicCast)
		b.instructions = append(b.instructions, i)
		return i
	}

	panic("Cannot cast non integer or pointer value")
}
