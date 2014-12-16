package statementTypes

type DeclarationStatement struct {
}

func NewDeclarationStatement(tokens []string) Statement {
	return &DeclarationStatement{}
}

func (ds *DeclarationStatement) Type() StatementType {
	return DeclarationType
}
