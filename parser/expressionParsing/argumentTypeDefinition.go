package expressionParsing

type ArgumentTypeDefinition struct {
	argType PrimTypeDefinition
}

func NewArgTypeDefinition(argDef PrimTypeDefinition) ArgumentTypeDefinition {
	return ArgumentTypeDefinition{
		argType: argDef,
	}
}

func (atd ArgumentTypeDefinition) GenerateGo() string {
	return atd.argType.GenerateGo()
}
