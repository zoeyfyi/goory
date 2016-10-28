package goory

import "fmt"

// Instruction is an operation that is executed on its operands
type Instruction interface {
	String() string
	IsTerminator() bool
	Type() Type
	Value() Value
	llvm() string
}

func binOpLlvm(instructionName string, value Value, lhs Value, rhs Value) string {
	return fmt.Sprintf("%s = %s %s %s, %s",
		value.llvm(),
		instructionName,
		lhs.Type().LLVMType(),
		lhs.llvm(),
		rhs.llvm())
}

// Float addition
type fadd struct {
	name string
	lhs  Value
	rhs  Value
}

func (i *fadd) String() string     { return "fadd" }
func (i *fadd) IsTerminator() bool { return false }
func (i *fadd) Type() Type         { return i.lhs.Type() }
func (i *fadd) Value() Value       { return Name{i.Type(), i.name} }
func (i *fadd) llvm() string       { return binOpLlvm("fadd", i.Value(), i.lhs, i.rhs) }

// Float subtruction
type fsub struct {
	name string
	lhs  Value
	rhs  Value
}

func (i *fsub) String() string     { return "fsub" }
func (i *fsub) IsTerminator() bool { return false }
func (i *fsub) Type() Type         { return i.lhs.Type() }
func (i *fsub) Value() Value       { return Name{i.Type(), i.name} }
func (i *fsub) llvm() string       { return binOpLlvm("fsub", i.Value(), i.lhs, i.rhs) }

// Float multiplication
type fmul struct {
	name string
	lhs  Value
	rhs  Value
}

func (i *fmul) String() string     { return "fmul" }
func (i *fmul) IsTerminator() bool { return false }
func (i *fmul) Type() Type         { return i.lhs.Type() }
func (i *fmul) Value() Value       { return Name{i.Type(), i.name} }
func (i *fmul) llvm() string       { return binOpLlvm("fmul", i.Value(), i.lhs, i.rhs) }

// Float division
type fdiv struct {
	name string
	lhs  Value
	rhs  Value
}

func (i *fdiv) String() string     { return "fdiv" }
func (i *fdiv) IsTerminator() bool { return false }
func (i *fdiv) Type() Type         { return i.lhs.Type() }
func (i *fdiv) Value() Value       { return Name{i.Type(), i.name} }
func (i *fdiv) llvm() string       { return binOpLlvm("fdiv", i.Value(), i.lhs, i.rhs) }

// Interger addition
type add struct {
	name string
	lhs  Value
	rhs  Value
}

func (i *add) String() string     { return "add" }
func (i *add) IsTerminator() bool { return false }
func (i *add) Type() Type         { return i.lhs.Type() }
func (i *add) Value() Value       { return Name{i.Type(), i.name} }
func (i *add) llvm() string       { return binOpLlvm("add", i.Value(), i.lhs, i.rhs) }

// Interger subtruction
type sub struct {
	name string
	lhs  Value
	rhs  Value
}

func (i *sub) String() string     { return "sub" }
func (i *sub) IsTerminator() bool { return false }
func (i *sub) Type() Type         { return i.lhs.Type() }
func (i *sub) Value() Value       { return Name{i.Type(), i.name} }
func (i *sub) llvm() string       { return binOpLlvm("sub", i.Value(), i.lhs, i.rhs) }

// Interger multiplication
type mul struct {
	name string
	lhs  Value
	rhs  Value
}

func (i *mul) String() string     { return "mul" }
func (i *mul) IsTerminator() bool { return false }
func (i *mul) Type() Type         { return i.lhs.Type() }
func (i *mul) Value() Value       { return Name{i.Type(), i.name} }
func (i *mul) llvm() string       { return binOpLlvm("mul", i.Value(), i.lhs, i.rhs) }

// Interger division
type div struct {
	name string
	lhs  Value
	rhs  Value
}

func (i *div) String() string     { return "div" }
func (i *div) IsTerminator() bool { return false }
func (i *div) Type() Type         { return i.lhs.Type() }
func (i *div) Value() Value       { return Name{i.Type(), i.name} }
func (i *div) llvm() string       { return binOpLlvm("div", i.Value(), i.lhs, i.rhs) }

// Interger truncation
type trunc struct {
	name  string
	value Value
	cast  Type
}

func (i *trunc) String() string     { return "trunc" }
func (i *trunc) IsTerminator() bool { return false }
func (i *trunc) Type() Type         { return i.cast }
func (i *trunc) Value() Value       { return Name{i.Type(), i.name} }
func (i *trunc) llvm() string {
	return fmt.Sprintf("%s = trunc %s %s to %s",
		i.Value().llvm(),
		i.value.Type().LLVMType(),
		i.value.llvm(),
		i.cast.LLVMType())
}

// Interger extension
type zext struct {
	name  string
	value Value
	cast  Type
}

func (i *zext) String() string     { return "zext" }
func (i *zext) IsTerminator() bool { return false }
func (i *zext) Type() Type         { return i.cast }
func (i *zext) Value() Value       { return Name{i.Type(), i.name} }
func (i *zext) llvm() string {
	return fmt.Sprintf("%s = zext %s %s to %s",
		i.Value().llvm(),
		i.value.Type().LLVMType(),
		i.value.llvm(),
		i.cast.LLVMType())
}

// Float truncation
type fptrunc struct {
	name  string
	value Value
	cast  Type
}

func (i *fptrunc) String() string     { return "fptrunc" }
func (i *fptrunc) IsTerminator() bool { return false }
func (i *fptrunc) Type() Type         { return i.cast }
func (i *fptrunc) Value() Value       { return Name{i.Type(), i.name} }
func (i *fptrunc) llvm() string {
	return fmt.Sprintf("%s = fptrunc %s %s to %s",
		i.Value().llvm(),
		i.value.Type().LLVMType(),
		i.value.llvm(),
		i.cast.LLVMType())
}

// Float extension
type fpext struct {
	name  string
	value Value
	cast  Type
}

func (i *fpext) String() string     { return "fpext" }
func (i *fpext) IsTerminator() bool { return false }
func (i *fpext) Type() Type         { return i.cast }
func (i *fpext) Value() Value       { return Name{i.Type(), i.name} }
func (i *fpext) llvm() string {
	return fmt.Sprintf("%s = fpext %s %s to %s",
		i.Value().llvm(),
		i.value.Type().LLVMType(),
		i.value.llvm(),
		i.cast.LLVMType())
}

// Float to int
type fptosi struct {
	name  string
	value Value
	cast  Type
}

func (i *fptosi) String() string     { return "fptosi" }
func (i *fptosi) IsTerminator() bool { return false }
func (i *fptosi) Type() Type         { return i.cast }
func (i *fptosi) Value() Value       { return Name{i.Type(), i.name} }
func (i *fptosi) llvm() string {
	return fmt.Sprintf("%s = fptosi %s %s to %s",
		i.Value().llvm(),
		i.value.Type().LLVMType(),
		i.value.llvm(),
		i.cast.LLVMType())
}

// Float to int
type sitofp struct {
	name  string
	value Value
	cast  Type
}

func (i *sitofp) String() string     { return "sitofp" }
func (i *sitofp) IsTerminator() bool { return false }
func (i *sitofp) Type() Type         { return i.cast }
func (i *sitofp) Value() Value       { return Name{i.Type(), i.name} }
func (i *sitofp) llvm() string {
	return fmt.Sprintf("%s = sitofp %s %s to %s",
		i.Value().llvm(),
		i.value.Type().LLVMType(),
		i.value.llvm(),
		i.cast.LLVMType())
}

// Interger comparison
type icmp struct {
	name string
	lhs  Value
	rhs  Value
	mode CompareMode
}

func (i *icmp) String() string     { return "icmp" }
func (i *icmp) IsTerminator() bool { return false }
func (i *icmp) Type() Type         { return BoolType }
func (i *icmp) Value() Value       { return Name{i.Type(), i.name} }
func (i *icmp) llvm() string {
	return fmt.Sprintf("%s = icmp %s %s %s, %s",
		i.Value().llvm(),
		i.mode.llvm(),
		i.lhs.Type().LLVMType(),
		i.lhs.llvm(),
		i.rhs.llvm())
}

// Float comparison
type fcmp struct {
	name string
	lhs  Value
	rhs  Value
	mode CompareMode
}

func (i *fcmp) String() string     { return "fcmp" }
func (i *fcmp) IsTerminator() bool { return false }
func (i *fcmp) Type() Type         { return BoolType }
func (i *fcmp) Value() Value       { return Name{i.Type(), i.name} }
func (i *fcmp) llvm() string {
	return fmt.Sprintf("%s = fcmp %s %s %s, %s",
		i.Value().llvm(),
		i.mode.llvm(),
		i.lhs.Type().LLVMType(),
		i.lhs.llvm(),
		i.rhs.llvm())
}

// Return statement
type ret struct {
	name  string
	value Value
}

func (i *ret) String() string     { return "ret" }
func (i *ret) IsTerminator() bool { return true }
func (i *ret) Type() Type         { return NilType }
func (i *ret) Value() Value       { return Name{i.Type(), i.name} }
func (i *ret) llvm() string {
	return fmt.Sprintf("ret %s %s",
		i.value.Type().LLVMType(),
		i.value.llvm())
}

// Call statement
type call struct {
	name     string
	function *Function
	operands []Value
}

func (i *call) String() string     { return "call" }
func (i *call) IsTerminator() bool { return true }
func (i *call) Type() Type         { return i.function.returnType }
func (i *call) Value() Value       { return Name{i.Type(), i.name} }
func (i *call) llvm() string {
	arguments := ""
	for opIndex, op := range i.operands {
		arguments += fmt.Sprintf("%s %s",
			op.Type().LLVMType(),
			op.llvm())

		if opIndex < len(i.operands)-1 {
			arguments += ", "
		}
	}

	return fmt.Sprintf("%s = call %s @%s(%s)",
		i.Value().llvm(),
		i.function.returnType.LLVMType(),
		i.function.name,
		arguments)
}

// Branch statement
type br struct {
	name  string
	block *Block
}

func (i *br) String() string     { return "br" }
func (i *br) IsTerminator() bool { return true }
func (i *br) Type() Type         { return NilType }
func (i *br) Value() Value       { return Name{i.Type(), i.name} }
func (i *br) llvm() string       { return fmt.Sprintf("br label %%%s", i.block.name) }

// Conditional Branch statement
type condBr struct {
	name       string
	condition  Value
	trueBlock  *Block
	falseBlock *Block
}

func (i *condBr) String() string     { return "br" }
func (i *condBr) IsTerminator() bool { return true }
func (i *condBr) Type() Type         { return NilType }
func (i *condBr) Value() Value       { return Name{i.Type(), i.name} }
func (i *condBr) llvm() string {
	return fmt.Sprintf("br i1 %s, label %%%s, label %%%s",
		i.condition.llvm(),
		i.trueBlock.name,
		i.falseBlock.name)
}

// None instruction
type none struct {
	value Value
}

func (i *none) String() string     { return "" }
func (i *none) IsTerminator() bool { return false }
func (i *none) Type() Type         { return i.value.Type() }
func (i *none) Value() Value       { return i.value }
func (i *none) llvm() string       { return "" }
