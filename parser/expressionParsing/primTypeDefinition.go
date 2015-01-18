package expressionParsing

type PrimTypeDefinition struct {
	name string
}

func (pd PrimTypeDefinition) GenerateGo() string {
	return pd.name
}

func (pd PrimTypeDefinition) String() string {
	return pd.name
}
