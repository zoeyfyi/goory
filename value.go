package goory

import "fmt"

// Value represents a parameter, the result of a function or a constant
type Value interface {
	Type() Type
	llvm() string
}

// Name reperesents a parameter or return value
type name struct {
	name     string
	nameType Type
}

func Name(name string, nameType Type) Value {
	return name{name, nameType}
}

// Name returns the name of the Name
func (n Name) Name() string {
	return n.name
}

func (n Name) valueType() {}

// Type returns the type of the Name
func (n Name) Type() Type {
	return n.t
}

func (n Name) llvm() string {
	return "%" + n.name
}

type FunctionValue struct {
	returnType Type
	name       string
}

func newFunctionValue(returnType Type, name string) Value {
	return Value(FunctionValue{returnType, name})
}

// Name returns the name of the FunctionValue
func (f FunctionValue) Name() string {
	return f.name
}

func (f FunctionValue) valueType() {}

// Type returns the type of the FunctionValue
func (f FunctionValue) Type() Type {
	return f.returnType
}

func (f FunctionValue) llvm() string {
	return "@" + f.name
}

// Constant reprents a constant value
type Constant struct {
	t     Type
	value interface{}
}

// Type returns the type of the constant
func (c Constant) Type() Type {
	return c.t
}

func (c Constant) valueType() {}

func (c Constant) llvm() string {
	switch c.value.(type) {
	case int8, int16, int32, int64:
		return fmt.Sprintf("%d", c.value)
	case float32, float64:
		return fmt.Sprintf("%f", c.value)
	case bool:
		return fmt.Sprintf("%t", c.value)
	default:
		panic("Unknown constant type")
	}
}

// ConstInt32 returns a constant value of type int32
func ConstInt32(i int32) Value {
	return Value(Constant{Int32Type, i})
}

// ConstInt64 returns a constant value of type int64
func ConstInt64(i int64) Value {
	return Value(Constant{Int64Type, i})
}

// ConstFloat32 returns a constant value of type float32
func ConstFloat32(i float32) Value {
	return Value(Constant{Float32Type, i})
}

// ConstFloat64 returns a constant value of type float64
func ConstFloat64(i float64) Value {
	return Value(Constant{Float64Type, i})
}

// ConstBool returns a constant value of type bool
func ConstBool(b bool) Value {
	return Value(Constant{BoolType, b})
}

const (
	// Integer modes
	ieq = iota
	ine
	iugt
	iuge
	iult
	iule
	isgt
	isge
	islt
	isle

	// Float modes
	foeq
	fogt
	foge
	folt
	fole
	fone
	ford
	fueq
	fugt
	fuge
	fult
	fule
	fune
	funo
)

type CompareMode struct {
	value int
}

func (c CompareMode) Type() Type {
	return compareMode
}

func (c CompareMode) valueType() {}

func (c CompareMode) llvm() string {
	switch c.value {
	// Integer modes
	case ieq:
		return "eq"
	case ine:
		return "ne"
	case iugt:
		return "ugt"
	case iuge:
		return "uge"
	case iult:
		return "ult"
	case iule:
		return "ule"
	case isgt:
		return "sgt"
	case isge:
		return "sge"
	case islt:
		return "slt"
	case isle:
		return "sle"
		// Float modes
	case foeq:
		return "oaq"
	case fogt:
		return "ogt"
	case foge:
		return "oge"
	case folt:
		return "olt"
	case fole:
		return "ole"
	case fone:
		return "one"
	case ford:
		return "ord"
	case fueq:
		return "ueq"
	case fugt:
		return "ugt"
	case fuge:
		return "uge"
	case fult:
		return "ult"
	case fule:
		return "ule"
	case fune:
		return "une"
	case funo:
		return "uno"

	default:
		panic("Unknown compare mode type")
	}
}

// IModeEq
func IModeEq() CompareMode {
	return CompareMode{ieq}
}

// IModeNe
func IModeNe() CompareMode {
	return CompareMode{ine}
}

// IModeUgt
func IModeUgt() CompareMode {
	return CompareMode{iugt}
}

// IModeUge
func IModeUge() CompareMode {
	return CompareMode{iuge}
}

// IModeUlt
func IModeUlt() CompareMode {
	return CompareMode{iult}
}

// IModeUle
func IModeUle() CompareMode {
	return CompareMode{iule}
}

// IModeSgt
func IModeSgt() CompareMode {
	return CompareMode{isgt}
}

// IModeSge
func IModeSge() CompareMode {
	return CompareMode{isge}
}

// IModeSlt
func IModeSlt() CompareMode {
	return CompareMode{islt}
}

// IModeSle
func IModeSle() CompareMode {
	return CompareMode{isle}
}

// FModeOeq
func FModeOeq() CompareMode {
	return CompareMode{foeq}
}

// FModeOgt
func FModeOgt() CompareMode {
	return CompareMode{fogt}
}

// FModeOge
func FModeOge() CompareMode {
	return CompareMode{foge}
}

// FModeOlt
func FModeOlt() CompareMode {
	return CompareMode{folt}
}

// FModeOle
func FModeOle() CompareMode {
	return CompareMode{fole}
}

// FModeOne
func FModeOne() CompareMode {
	return CompareMode{fone}
}

// FModeOrd
func FModeOrd() CompareMode {
	return CompareMode{ford}
}

// FModeUeq
func FModeUeq() CompareMode {
	return CompareMode{fueq}
}

// FModeUgt
func FModeUgt() CompareMode {
	return CompareMode{fugt}
}

// FModeUge
func FModeUge() CompareMode {
	return CompareMode{fuge}
}

// FModeUlt
func FModeUlt() CompareMode {
	return CompareMode{fult}
}

// FModeUle
func FModeUle() CompareMode {
	return CompareMode{fule}
}

// FModeUne
func FModeUne() CompareMode {
	return CompareMode{fune}
}

// FModeUno
func FModeUno() CompareMode {
	return CompareMode{funo}
}
