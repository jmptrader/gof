package statementTypes

type IfStatement struct {
}

func NewIfStatement() Statement {
	return &IfStatement{}
}

func (is *IfStatement) Type() StatementType {
	return IfType
}

func (is *IfStatement) AddLine(tokens []string) {
}
