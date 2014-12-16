package statementTypes

type DeclarationStatement struct {
}

func NewDeclarationStatement() Statement {
	return &DeclarationStatement{}
}

func (ds *DeclarationStatement) Type() StatementType {
	return DeclarationType
}

func (ds *DeclarationStatement) AddLine(tokens []string) {
}
