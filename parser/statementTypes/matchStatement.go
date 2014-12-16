package statementTypes

type MatchStatement struct {
}

func NewMatchStatement() Statement {
	return &MatchStatement{}
}

func (ms *MatchStatement) Type() StatementType {
	return MatchType
}

func (ms *MatchStatement) AddLine(tokens []string) {
}
