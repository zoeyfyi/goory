package goory

import (
	"github.com/bongo227/goory/types"
)

func IntType(bits int) types.Type { return types.NewIntType(bits) }
func FloatType() types.Type       { return types.NewFloatType() }
func DoubleType() types.Type      { return types.NewDoubleType() }
func BoolType() types.Type        { return types.NewBoolType() }
