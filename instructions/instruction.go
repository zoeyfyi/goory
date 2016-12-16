package instructions

import "github.com/bongo227/goory/types"

// assertEqual checks all types are the same
func assertEqual(t ...types.Type) {
	if len(t) == 0 || len(t) == 1 {
		return
	}

	for i := 1; i < len(t); i++ {
		if !t[0].Equal(t[i]) {
			panic("Types are not equal")
		}
	}
}

func assertInt(t types.Type) {
	if _, ok := t.(types.IntType); !ok {
		panic("Value is not an integer")
	}
}

func assertFunction(t types.Type) types.Function {
	fType, ok := t.(types.Function)
	if !ok {
		panic("Value is not a function type")
	}

	return fType
}
