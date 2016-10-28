package goory

const (
	// Integers
	typeInt = iota
	typeInt8
	typeInt16
	typeInt32
	typeInt64

	// Floats
	typeFloat
	typeFloat32
	typeFloat64

	typeBool
	typeNil
	typeCompareMode
)

// Types
var (
	NilType = Type{
		id: typeNil,
	}

	IntType = Type{
		id: typeInt,
	}

	Int8Type = Type{
		id: typeInt8,
	}

	Int16Type = Type{
		id: typeInt16,
	}

	Int32Type = Type{
		id: typeInt32,
	}

	Int64Type = Type{
		id: typeInt64,
	}

	FloatType = Type{
		id: typeFloat,
	}

	Float32Type = Type{
		id: typeFloat32,
	}

	Float64Type = Type{
		id: typeFloat64,
	}

	BoolType = Type{
		id: typeBool,
	}

	compareMode = Type{
		id: typeCompareMode,
	}
)

// Type describes a fundemental type i.e integers, floats etc
type Type struct {
	id int
}

// String returns the type name as a string
func (t Type) String() string {
	switch t.id {
	case typeInt:
		return "Int"
	case typeInt8:
		return "Int8"
	case typeInt16:
		return "Int16"
	case typeInt32:
		return "Int32"
	case typeInt64:
		return "Int64"
	case typeFloat:
		return "Float"
	case typeFloat32:
		return "Float32"
	case typeFloat64:
		return "Float64"
	case typeBool:
		return "Bool"
	case typeNil:
		return "Nil"
	default:
		panic("Unknow type id, cannot get string of unknown type")
	}
}

// LLVMType returns the type name in llvm format
func (t Type) LLVMType() string {
	switch t.id {
	case typeInt:
		return "i64" // TODO: Check what type this should be
	case typeInt8:
		return "i8"
	case typeInt16:
		return "i16"
	case typeInt32:
		return "i32"
	case typeInt64:
		return "i64"
	case typeFloat:
		return "double" // TODO: Check what type this should be
	case typeFloat32:
		return "float"
	case typeFloat64:
		return "double"
	case typeBool:
		return "i1"
	case typeNil:
		return "null"
	default:
		panic("Unknow type id, cannot get llvm type string of unknown type")
	}
}
