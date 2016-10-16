package goory

const (
	typeInt32 = iota
	typeInt64
	typeFloat32
	typeFloat64
	typeNil
)

// Types
var (
	NilType = Type{
		id: typeNil,
	}

	Float32Type = Type{
		id: typeFloat32,
	}

	Float64Type = Type{
		id: typeFloat64,
	}
)

// Type describes a fundemental type i.e integers, floats etc
type Type struct {
	id int
}

// String returns the type name as a string
func (t Type) String() string {
	switch t.id {
	case typeInt32:
		return "Int32"
	case typeInt64:
		return "Int64"
	case typeFloat32:
		return "Float32"
	case typeFloat64:
		return "Float64"
	default:
		panic("Unknow type id, cannot get string of unknown type")
	}
}

// LLVMType returns the type name in llvm format
func (t Type) LLVMType() string {
	switch t.id {
	case typeInt32:
		return "i32"
	case typeInt64:
		return "i64"
	case typeFloat32:
		return "f32"
	case typeFloat64:
		return "f64"
	default:
		panic("Unknow type id, cannot get llvm type string of unknown type")
	}
}

// FunctionType describes the parameters and return values of a function
type FunctionType struct {
	parameters []Type
	returnType Type
}

// NewFunctionType returns a function type with parameter types and a return type
func NewFunctionType(parameterTypes []Type, returnType Type) *FunctionType {
	return &FunctionType{
		parameters: parameterTypes,
		returnType: returnType,
	}
}

// ParameterTypes returns a slice of a function types parameter types
func (f *FunctionType) ParameterTypes() []Type {
	return f.parameters
}

// ReturnType returns the type of a functions return type
func (f *FunctionType) ReturnType() Type {
	return f.returnType
}
