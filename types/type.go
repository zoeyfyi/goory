package types

// Type defines all llvm types
type Type interface {
	String() string
	Equal(Type) bool
}

var (
	INT32  = NewIntType(32)
	INT64  = NewIntType(64)
	FLOAT  = NewFloatType()
	DOUBLE = NewDoubleType()
	VOID   = NewVoid()
)

// Compare compares two atomics a and b
// returns 1 	if a > b
// returns 0 	if a = b
// returns -1 	if a < b
func Compare(a Atomic, b Atomic) int {
	if a.Bits() > b.Bits() {
		return 1
	}

	if a.Bits() == b.Bits() {
		return 0
	}

	return -1
}

func IsInteger(a Type) bool {
	_, ok := a.(IntType)
	return ok
}

func IsFp(a Type) bool {
	_, okFloat := a.(FloatType)
	_, okDouble := a.(DoubleType)
	return okFloat || okDouble
}
