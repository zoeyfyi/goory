package goory

import "fmt"

// Value represents a parameter, the result of a function or a constant
type Value interface {
	Type() Type
	llvm() string
	valueType() // stops instruction satisfying the value interface
}

// Name reperesents a parameter or return value
type Name struct {
	t    Type
	name string
}

func newName(t Type, name string) Value {
	return Value(Name{t, name})
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
	compareEqual = iota
	compareNotEqual
	compareUnsignedGreaterThan
	compareUnsignedGreaterThanOrEqualTo
	compareUnsignedLessThan
	compareUnsignedLessThanOrEqualTo
	compareSignedGreaterThan
	compareSignedGreaterThanOrEqualTo
	compareSignedLessThan
	compareSignedLessThanOrEqualTo
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
	case compareEqual:
		return "eq"
	case compareNotEqual:
		return "ne"
	case compareUnsignedGreaterThan:
		return "ugt"
	case compareUnsignedGreaterThanOrEqualTo:
		return "uge"
	case compareUnsignedLessThan:
		return "ult"
	case compareUnsignedLessThanOrEqualTo:
		return "ule"
	case compareSignedGreaterThan:
		return "sgt"
	case compareSignedGreaterThanOrEqualTo:
		return "sge"
	case compareSignedLessThan:
		return "slt"
	case compareSignedLessThanOrEqualTo:
		return "sle"
	default:
		panic("Unknown compare mode type")
	}
}

// ModeEq
func ModeEq() CompareMode {
	return CompareMode{compareEqual}
}

// ModeNe
func ModeNe() CompareMode {
	return CompareMode{compareNotEqual}
}

// ModeUgt
func ModeUgt() CompareMode {
	return CompareMode{compareUnsignedGreaterThan}
}

// ModeUge
func ModeUge() CompareMode {
	return CompareMode{compareUnsignedGreaterThanOrEqualTo}
}

// ModeUlt
func ModeUlt() CompareMode {
	return CompareMode{compareUnsignedLessThan}
}

// ModeUle
func ModeUle() CompareMode {
	return CompareMode{compareUnsignedLessThanOrEqualTo}
}

// ModeSgt
func ModeSgt() CompareMode {
	return CompareMode{compareSignedGreaterThan}
}

// ModeSge
func ModeSge() CompareMode {
	return CompareMode{compareSignedGreaterThanOrEqualTo}
}

// ModeSlt
func ModeSlt() CompareMode {
	return CompareMode{compareSignedLessThan}
}

// ModeSle
func ModeSle() CompareMode {
	return CompareMode{compareSignedLessThanOrEqualTo}
}
