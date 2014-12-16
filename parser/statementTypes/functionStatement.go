package statementTypes

type FunctionStatement struct {
	name string
}

func NewFunctionStatement(tokens []string) Statement {
	return &FunctionStatement{
		name: fetchName(tokens[1]),
	}
}

func (fs *FunctionStatement) Type() StatementType {
	return FunctionType
}

func (fs *FunctionStatement) Name() string {
	return fs.name
}

func fetchName(name string) string {
	if name == "->" {
		return ""
	}
	return name
}
