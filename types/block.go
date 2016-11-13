package types

type Block struct{}

// NewBlockType creates a new block type
func NewBlockType() Type {
	return Block{}
}

func (b Block) String() string {
	return "BlockType"
}

func (b Block) Equal(n Type) bool {
	_, ok := n.(Block)
	return ok
}
