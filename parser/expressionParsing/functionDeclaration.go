package expressionParsing

type TypeName string

type FunctionDeclaration struct {
	name       string
	returnType TypeName
	argType    TypeName
}

func NewFunctionDeclaration(retType TypeName, argType TypeName) *FunctionDeclaration {
	return &FunctionDeclaration{
		returnType: retType,
		argType:    argType,
	}
}

func NewDefinition(retType TypeName) *FunctionDeclaration {
	return NewFunctionDeclaration(retType, "")
}

func (fd *FunctionDeclaration) SetName(name string) {
	fd.name = name
}

func (fd *FunctionDeclaration) Name() string {
	return fd.name
}

func (fd *FunctionDeclaration) ReturnType() TypeName {
	return fd.returnType
}

func (fd *FunctionDeclaration) ArgumentType() TypeName {
	return fd.argType
}

func (fd *FunctionDeclaration) IsDefinition() bool {
	return fd.argType == ""
}
