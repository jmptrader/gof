package expressionParsing

type ArgumentTypeDefinition struct {
	argType TypeDefinition
}

func NewArgTypeDefinition(argDef TypeDefinition) ArgumentTypeDefinition {
	return ArgumentTypeDefinition{
		argType: argDef,
	}
}

func (atd ArgumentTypeDefinition) GenerateGo() string {
	return atd.argType.GenerateGo()
}
