package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Store struct {
	block      value.Value
	allocation *Alloca
	value      value.Value
}

// NewStore creates a new Add operation
func NewStore(block value.Value, allocation *Alloca, value value.Value) *Store {
	return &Store{block, allocation, value}
}

func (i *Store) Block() value.Value {
	return i.block
}

func (i *Store) IsTerminator() bool {
	return false
}

func (i *Store) Type() types.Type {
	return types.VOID
}

func (i *Store) Ident() string {
	return ""
}

func (i *Store) Llvm() string {
	return fmt.Sprintf("store %s %s, %s %s",
		i.value.Type().String(),
		i.value.Ident(),
		i.allocation.Type().String(),
		i.allocation.Ident())
}
