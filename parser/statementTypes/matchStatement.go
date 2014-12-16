package statementTypes

type MatchStatement struct {
}

func NewMatchStatement(tokens []string) Statement {
	return &MatchStatement{}
}

func (ms *MatchStatement) Type() StatementType {
	return MatchType
}
