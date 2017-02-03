package parse

type Node interface {
	LHS() Node
	RHS() Node
	Value() string
}

type NodeNumberLit struct {
	val string
}

type NodeMulti struct {
	lh Node
	rh Node
}

func (num NodeNumberLit) LHS() Node { return num }
func (num NodeNumberLit) RHS() Node { return num }
func (num NodeNumberLit) Value() string { return num.val }

func (multi NodeMulti) LHS() Node { return multi.lh }
func (multi NodeMulti) RHS() Node { return multi.rh }
func (multi NodeMulti) Value() string { return "" }

func NewParser(toknr *Tokenizer) *Parser {
	return &Parser{toknr}
}
