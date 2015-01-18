package expressionParsing

import (
	"fmt"
)

type FuncTypeDefinition struct {
	argName string
	argType TypeDefinition
	retType TypeDefinition
	code    string
}

func (fd FuncTypeDefinition) GenerateGo() string {
	return fmt.Sprintf("func (%s %s) %s", fd.argName, fd.argType.GenerateGo(), fd.retType.GenerateGo())
}

func (fd FuncTypeDefinition) String() string {
	return fd.code
}

func (fd FuncTypeDefinition) ReturnType() TypeDefinition {
	return fd.retType
}

func (fd FuncTypeDefinition) ArgumentType() TypeDefinition {
	return fd.argType
}

func (fd FuncTypeDefinition) ArgumentName() string {
	return fd.argName
}

func newPrimTypeDefinition(name string) PrimTypeDefinition {
	return PrimTypeDefinition{
		name: name,
	}
}
