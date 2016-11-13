package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"
	"github.com/bongo227/goory/value"
)

type Phi struct {
	block    value.Value
	name     string
	incoming []value.Instruction
}

func NewPhi(block value.Value, name string) *Phi {
	return &Phi{block, name, []value.Instruction{}}
}

func (i *Phi) Block() value.Value {
	return i.block
}

func (i *Phi) IsTerminator() bool {
	return false
}

func (i *Phi) Type() types.Type {
	return i.incoming[0].Type()
}

func (i *Phi) Ident() string {
	return "%" + i.name
}

func (i *Phi) AddIncoming(value value.Instruction) {
	i.incoming = append(i.incoming, value)
}

func (i *Phi) Llvm() string {
	s := fmt.Sprintf("%%%s = phi %s ", i.name, i.incoming[0].Type().String())
	for incomingIndex, incoming := range i.incoming {
		s += fmt.Sprintf("[ %s %s ]", incoming.Ident(), incoming.Block().Ident())
		if incomingIndex < len(i.incoming)-1 {
			s += ", "
		}
	}

	return s
}
