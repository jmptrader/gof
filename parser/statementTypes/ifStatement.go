package statementTypes

type IfStatement struct {
}

func NewIfStatement(tokens []string) Statement {
	return &IfStatement{}
}

func (is *IfStatement) Type() StatementType {
	return IfType
}
