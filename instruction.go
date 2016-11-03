package goory

import "fmt"

// Instruction is an operation that is executed on its operands
type Instruction interface {
	String() string
	IsTerminator() bool
	Type() Type
	llvm() string
	ident() string
}

func binOpLlvm(instructionName string, ident string, lhs Value, rhs Value) string {
	return fmt.Sprintf("%%%s = %s %s %s, %s",
		ident,
		instructionName,
		lhs.Type().llvm(),
		lhs.ident(),
		rhs.ident())
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
func (i *fadd) ident() string      { return "%" + i.name }
func (i *fadd) llvm() string       { return binOpLlvm("fadd", i.name, i.lhs, i.rhs) }

// Float subtruction
type fsub struct {
	name string
	lhs  Value
	rhs  Value
}

func (i *fsub) String() string     { return "fsub" }
func (i *fsub) IsTerminator() bool { return false }
func (i *fsub) Type() Type         { return i.lhs.Type() }
func (i *fsub) ident() string      { return "%" + i.name }
func (i *fsub) llvm() string       { return binOpLlvm("fsub", i.name, i.lhs, i.rhs) }

// Float multiplication
type fmul struct {
	name string
	lhs  Value
	rhs  Value
}

func (i *fmul) String() string     { return "fmul" }
func (i *fmul) IsTerminator() bool { return false }
func (i *fmul) Type() Type         { return i.lhs.Type() }
func (i *fmul) ident() string      { return "%" + i.name }
func (i *fmul) llvm() string       { return binOpLlvm("fmul", i.name, i.lhs, i.rhs) }

// Float division
type fdiv struct {
	name string
	lhs  Value
	rhs  Value
}

func (i *fdiv) String() string     { return "fdiv" }
func (i *fdiv) IsTerminator() bool { return false }
func (i *fdiv) Type() Type         { return i.lhs.Type() }
func (i *fdiv) ident() string      { return "%" + i.name }
func (i *fdiv) llvm() string       { return binOpLlvm("fdiv", i.name, i.lhs, i.rhs) }

// Interger addition
type add struct {
	name string
	lhs  Value
	rhs  Value
}

func (i *add) String() string     { return "add" }
func (i *add) IsTerminator() bool { return false }
func (i *add) Type() Type         { return i.lhs.Type() }
func (i *add) ident() string      { return "%" + i.name }
func (i *add) llvm() string       { return binOpLlvm("add", i.name, i.lhs, i.rhs) }

// Interger subtruction
type sub struct {
	name string
	lhs  Value
	rhs  Value
}

func (i *sub) String() string     { return "sub" }
func (i *sub) IsTerminator() bool { return false }
func (i *sub) Type() Type         { return i.lhs.Type() }
func (i *sub) ident() string      { return "%" + i.name }
func (i *sub) llvm() string       { return binOpLlvm("sub", i.name, i.lhs, i.rhs) }

// Interger multiplication
type mul struct {
	name string
	lhs  Value
	rhs  Value
}

func (i *mul) String() string     { return "mul" }
func (i *mul) IsTerminator() bool { return false }
func (i *mul) Type() Type         { return i.lhs.Type() }
func (i *mul) ident() string      { return "%" + i.name }
func (i *mul) llvm() string       { return binOpLlvm("mul", i.name, i.lhs, i.rhs) }

// Interger division
type div struct {
	name string
	lhs  Value
	rhs  Value
}

func (i *div) String() string     { return "div" }
func (i *div) IsTerminator() bool { return false }
func (i *div) Type() Type         { return i.lhs.Type() }
func (i *div) ident() string      { return "%" + i.name }
func (i *div) llvm() string       { return binOpLlvm("div", i.name, i.lhs, i.rhs) }

// Interger comparisons
type icmp struct {
	name string
	mode intCompareMode
	lhs  Value
	rhs  Value
}

func (i *icmp) String() string     { return "icmp" }
func (i *icmp) IsTerminator() bool { return false }
func (i *icmp) Type() Type         { return i.lhs.Type() }
func (i *icmp) ident() string      { return "%" + i.name }
func (i *icmp) llvm() string {
	return fmt.Sprintf("%%%s = icmp %s %s %s, %s", i.name, i.mode, i.Type().llvm(), i.lhs.ident(), i.rhs.ident())
}

// Float comparisons
type fcmp struct {
	name string
	mode floatCompareMode
	lhs  Value
	rhs  Value
}

func (i *fcmp) String() string     { return "fcmp" }
func (i *fcmp) IsTerminator() bool { return false }
func (i *fcmp) Type() Type         { return i.lhs.Type() }
func (i *fcmp) ident() string      { return "%" + i.name }
func (i *fcmp) llvm() string {
	return fmt.Sprintf("%%%s = fcmp %s %s %s, %s", i.name, i.mode, i.Type().llvm(), i.lhs.ident(), i.rhs.ident())
}

// Interger truncation
type trunc struct {
	name  string
	value Value
	cast  Type
}

func (i *trunc) String() string     { return "trunc" }
func (i *trunc) IsTerminator() bool { return false }
func (i *trunc) Type() Type         { return i.cast }
func (i *trunc) ident() string      { return "%" + i.name }
func (i *trunc) llvm() string {
	return fmt.Sprintf("%s = trunc %s %s to %s",
		i.name,
		i.value.Type().llvm(),
		i.value.llvm(),
		i.cast.llvm())
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
func (i *zext) ident() string      { return "%" + i.name }
func (i *zext) llvm() string {
	return fmt.Sprintf("%s = zext %s %s to %s",
		i.name,
		i.value.Type().llvm(),
		i.value.llvm(),
		i.cast.llvm())
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
func (i *fptrunc) ident() string      { return "%" + i.name }
func (i *fptrunc) llvm() string {
	return fmt.Sprintf("%s = fptrunc %s %s to %s",
		i.name,
		i.value.Type().llvm(),
		i.value.llvm(),
		i.cast.llvm())
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
func (i *fpext) ident() string      { return "%" + i.name }
func (i *fpext) llvm() string {
	return fmt.Sprintf("%s = fpext %s %s to %s",
		i.name,
		i.value.Type().llvm(),
		i.value.llvm(),
		i.cast.llvm())
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
func (i *fptosi) ident() string      { return "%" + i.name }
func (i *fptosi) llvm() string {
	return fmt.Sprintf("%s = fptosi %s %s to %s",
		i.name,
		i.value.Type().llvm(),
		i.value.llvm(),
		i.cast.llvm())
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
func (i *sitofp) ident() string      { return "%" + i.name }
func (i *sitofp) llvm() string {
	return fmt.Sprintf("%s = sitofp %s %s to %s",
		i.name,
		i.value.Type().llvm(),
		i.value.llvm(),
		i.cast.llvm())
}

// Return statement
type ret struct {
	name  string
	value Value
}

func (i *ret) String() string     { return "ret" }
func (i *ret) IsTerminator() bool { return true }
func (i *ret) Type() Type         { return NewVoidType() }
func (i *ret) ident() string      { return "%" + i.name }
func (i *ret) llvm() string {
	return fmt.Sprintf("ret %s %s",
		i.value.Type().llvm(),
		i.value.ident())
}

// Call statement
type call struct {
	name     string
	function *Function
	operands []Value
}

func (i *call) String() string     { return "call" }
func (i *call) IsTerminator() bool { return true }
func (i *call) Type() Type         { return i.function.functionType.returnType }
func (i *call) ident() string      { return "%" + i.name }
func (i *call) llvm() string {
	arguments := ""
	for opIndex, op := range i.operands {
		arguments += fmt.Sprintf("%s %s",
			op.Type().llvm(),
			op.ident())

		if opIndex < len(i.operands)-1 {
			arguments += ", "
		}
	}

	return fmt.Sprintf("%%%s = call %s @%s(%s)",
		i.name,
		i.function.functionType.returnType.llvm(),
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
func (i *br) Type() Type         { return NewVoidType() }
func (i *br) ident() string      { return "%" + i.name }
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
func (i *condBr) Type() Type         { return NewVoidType() }
func (i *condBr) ident() string      { return "%" + i.name }
func (i *condBr) llvm() string {
	return fmt.Sprintf("br i1 %s, label %%%s, label %%%s",
		i.condition.ident(),
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
func (i *none) ident() string      { return "" }
func (i *none) llvm() string       { return "" }
