package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"
	"github.com/bongo227/goory/value"
)

type Phi struct {
	block    value.Value
	name     string
	incoming []phiNode
}

type phiNode struct {
	value value.Value
	block value.Value
}

func NewPhi(block value.Value, name string) *Phi {
	return &Phi{block, name, []phiNode{}}
}

func (i *Phi) Block() value.Value {
	return i.block
}

func (i *Phi) IsTerminator() bool {
	return false
}

func (i *Phi) Type() types.Type {
	return i.incoming[0].value.Type()
}

func (i *Phi) Ident() string {
	return "%" + i.name
}

func (i *Phi) AddIncoming(value value.Instruction) {
	i.incoming = append(i.incoming, phiNode{value, value.Block()})
}

func (i *Phi) AddIncomingConst(value value.Value, block value.Value) {
	i.incoming = append(i.incoming, phiNode{value, block})
}

func (i *Phi) Llvm() string {
	s := fmt.Sprintf("%%%s = phi %s ", 
		 i.name, 
		 i.Type().String())
	
	for incomingIndex, incoming := range i.incoming {
		s += fmt.Sprintf("[ %s %s ]", incoming.value.Ident(), incoming.block.Ident())
		if incomingIndex < len(i.incoming)-1 {
			s += ", "
		}
	}

	return s
}
