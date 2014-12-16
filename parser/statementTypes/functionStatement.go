package statementTypes

type FunctionStatement struct {
	name string
}

func NewFunctionStatement() Statement {
	return &FunctionStatement{}
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

func (ds *FunctionStatement) AddLine(tokens []string) {
}
