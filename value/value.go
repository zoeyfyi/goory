package value

import "github.com/bongo227/goory/types"

// Value are what modules mostly consist with
type Value interface {
	Type() types.Type
	String() string
	Ident() string
}
