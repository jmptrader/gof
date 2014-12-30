package expressionParsing

type TypeName string

type FunctionDeclaration struct {
	returnType TypeName
	argTypes   []TypeName
}

func NewFunctionDeclaration(retType TypeName, argTypes ...TypeName) *FunctionDeclaration {
	return &FunctionDeclaration{
		returnType: retType,
		argTypes:   argTypes,
	}
}

func (fd *FunctionDeclaration) ReturnType() TypeName {
	return fd.returnType
}

func (fd *FunctionDeclaration) ArgumentTypes() []TypeName {
	return fd.argTypes
}
