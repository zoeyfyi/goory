package goory

type nextName func() string

// Block is a seqential list of instructions
type Block struct {
	function     *Function
	name         string
	instructions []Instruction
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

// Terminated returns true if the block is terminated (ends in branch or return)
func (b *Block) Terminated() bool {
	return b.instructions[len(b.instructions)-1].IsTerminator()
}

// Fadd creates a new float addition between lhs and rhs
func (b *Block) Fadd(lhs Value, rhs Value) Instruction {
	i := &fadd{
		name: b.nextName(),
		lhs:  lhs,
		rhs:  rhs,
	}

	b.instructions = append(b.instructions, i)

	return i
}

// Fsub creates a new float subtraction between left and right
func (b *Block) Fsub(lhs Value, rhs Value) Instruction {
	i := &fsub{
		name: b.nextName(),
		lhs:  lhs,
		rhs:  rhs,
	}

	b.instructions = append(b.instructions, i)

	return i
}

// Fmul creates a new float multiplication between left and right
func (b *Block) Fmul(lhs Value, rhs Value) Instruction {
	i := &fmul{
		name: b.nextName(),
		lhs:  lhs,
		rhs:  rhs,
	}

	b.instructions = append(b.instructions, i)

	return i
}

// Fdiv creates a new float division between lhs and rhs
func (b *Block) Fdiv(lhs Value, rhs Value) Instruction {
	i := &fdiv{
		name: b.nextName(),
		lhs:  lhs,
		rhs:  rhs,
	}

	b.instructions = append(b.instructions, i)

	return i
}

// Sub creates a new interger subtraction between left and right
func (b *Block) Sub(lhs Value, rhs Value) Instruction {
	i := &sub{
		name: b.nextName(),
		lhs:  lhs,
		rhs:  rhs,
	}

	b.instructions = append(b.instructions, i)

	return i
}

// Add creates a new interger addition between left and right
func (b *Block) Add(lhs Value, rhs Value) Instruction {
	i := &add{
		name: b.nextName(),
		lhs:  lhs,
		rhs:  rhs,
	}

	b.instructions = append(b.instructions, i)

	return i
}

// Mul creates a new interger multiplication between left and right
func (b *Block) Mul(lhs Value, rhs Value) Instruction {
	i := &mul{
		name: b.nextName(),
		lhs:  lhs,
		rhs:  rhs,
	}

	b.instructions = append(b.instructions, i)
	return i
}

// Div creates a new interger division between left and right
func (b *Block) Div(lhs Value, rhs Value) Instruction {
	i := &div{
		name: b.nextName(),
		lhs:  lhs,
		rhs:  rhs,
	}

	b.instructions = append(b.instructions, i)
	return i
}

type intCompareMode string

const (
	// IntEq is an integer equal comparision
	IntEq intCompareMode = "eq"
	// IntNe is an integer not equal comparision
	IntNe intCompareMode = "ne"
	// IntUgt is an unsigned integer greater than comparision
	IntUgt intCompareMode = "ugt"
	// IntUge is an unsigned integer greater than or equal comparision
	IntUge intCompareMode = "uge"
	// IntUlt is an unsigned integer less than comparision
	IntUlt intCompareMode = "ult"
	// IntUle is an unsigned integer less than or equal comparision
	IntUle intCompareMode = "ule"
	// IntSgt is an signed integer greater than comparision
	IntSgt intCompareMode = "sgt"
	// IntSge is an signed integer greater than or equal comparision
	IntSge intCompareMode = "sge"
	// IntSlt is an signed integer less than comparision
	IntSlt intCompareMode = "slt"
	// IntSle is an unsigned integer less than or equal comparision
	IntSle intCompareMode = "sle"
)

// ICmp compares two integers
func (b *Block) ICmp(mode intCompareMode, lhs Value, rhs Value) Instruction {
	i := &icmp{
		name: b.nextName(),
		mode: mode,
		lhs:  lhs,
		rhs:  rhs,
	}

	b.instructions = append(b.instructions, i)
	return i
}

type floatCompareMode string

const (
	FloatOeq floatCompareMode = "oeq"
	FloatOgt floatCompareMode = "ogt"
	FloatOge floatCompareMode = "oge"
	FloatOlt floatCompareMode = "olt"
	FloatOle floatCompareMode = "ole"
	FloatOne floatCompareMode = "one"
	FloatOrd floatCompareMode = "ord"
	FloatUeq floatCompareMode = "ueq"
	FloatUgt floatCompareMode = "ugt"
	FloatUge floatCompareMode = "uge"
	FloatUlt floatCompareMode = "ult"
	FloatUle floatCompareMode = "ule"
	FloatUne floatCompareMode = "une"
	FloatUno floatCompareMode = "uno"
)

// FCmp compares two floats
func (b *Block) FCmp(mode floatCompareMode, lhs Value, rhs Value) Instruction {
	i := &fcmp{
		name: b.nextName(),
		mode: mode,
		lhs:  lhs,
		rhs:  rhs,
	}

	b.instructions = append(b.instructions, i)
	return i
}

// Ret creates a new return for ret
func (b *Block) Ret(value Value) Instruction {
	i := &ret{
		name:  b.nextName(),
		value: value,
	}

	b.instructions = append(b.instructions, i)

	return i
}

// Call creates a call instruction with the arguments
func (b *Block) Call(function *Function, operands ...Value) Instruction {
	i := &call{
		name:     b.nextName(),
		function: function,
		operands: operands,
	}

	b.instructions = append(b.instructions, i)

	return i
}

// Br creates a branch instruction to block
func (b *Block) Br(block *Block) Instruction {
	i := &br{
		name:  b.nextName(),
		block: block,
	}

	b.instructions = append(b.instructions, i)

	return i
}

// CondBr creates a conditional branch based on the value on condition
func (b *Block) CondBr(condition Value, trueBlock *Block, falseBlock *Block) Instruction {
	i := &condBr{
		name:       b.nextName(),
		condition:  condition,
		trueBlock:  trueBlock,
		falseBlock: falseBlock,
	}

	b.instructions = append(b.instructions, i)

	return i
}

// Cast creates a cast for value to type cast
func (b *Block) Cast(value Value, cast Atomic) Instruction {
	atomic, isAtomic := value.Type().(Atomic)

	if !isAtomic {
		panic("Value is not an atomic type")
	}

	// Value type is the same as cast type so return an empty instruction
	if atomic == cast {
		return &none{value}
	}

	// Cast integer to integer
	if atomic.IsInteger() && cast.IsInteger() {
		if Compare(atomic, cast) == 1 {
			// Cast is bigger so expand value
			i := &zext{b.nextName(), value, cast}
			b.instructions = append(b.instructions, i)
			return i
		}
		// Cast is smaller so truncate value
		i := &trunc{b.nextName(), value, cast}
		b.instructions = append(b.instructions, i)
		return i
	}

	// Cast float to float
	if atomic.IsFloat() && cast.IsFloat() {
		if Compare(atomic, cast) == 1 {
			// Cast is bigger so expand value
			i := &fpext{b.nextName(), value, cast}
			b.instructions = append(b.instructions, i)
			return i
		}
		// Cast is smaller so truncate value
		i := &fptrunc{b.nextName(), value, cast}
		b.instructions = append(b.instructions, i)
		return i
	}

	// Cast integer to float
	if atomic.IsInteger() && cast.IsFloat() {
		i := &sitofp{b.nextName(), value, cast}
		b.instructions = append(b.instructions, i)
		return i
	}

	// Cast float to integer
	if atomic.IsFloat() && cast.IsInteger() {
		i := &fptosi{b.nextName(), value, cast}
		b.instructions = append(b.instructions, i)
		return i
	}

	panic("Cannot cast non integer or pointer value")
}
