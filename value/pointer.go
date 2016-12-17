package value

import "github.com/bongo227/goory/types"

type Pointer interface {
	Type() types.Type
	BaseType() types.Type
	Ident() string
}
