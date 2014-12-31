package expressionParsing

import (
	"fmt"
)

type TypeName string

type TypeDefinition interface {
	IsFunc() bool
	ReturnType() TypeDefinition
	Name() TypeName
}

type FuncTypeDefinition struct {
	Argument   TypeDefinition
	returnType TypeDefinition
	name       string
}

type PrimTypeDefinition struct {
	name TypeName
}

func NewFuncTypeDefinition(arg, retType TypeDefinition) *FuncTypeDefinition {
	return &FuncTypeDefinition{
		Argument:   arg,
		returnType: retType,
	}
}

func NewPrimTypeDefinition(name TypeName) TypeDefinition {
	return PrimTypeDefinition{
		name: name,
	}
}

func (f *FuncTypeDefinition) IsFunc() bool {
	return true
}

func (f *FuncTypeDefinition) IsDefinition() bool {
	return f.Argument == nil
}

func (f *FuncTypeDefinition) ReturnType() TypeDefinition {
	return f.returnType
}

func (f *FuncTypeDefinition) FuncName() string {
	return f.name
}

func (f *FuncTypeDefinition) Name() TypeName {
	return TypeName(fmt.Sprintf("%s->%s", f.Argument.Name(), f.returnType.Name()))
}

func (f PrimTypeDefinition) IsFunc() bool {
	return false
}

func (p PrimTypeDefinition) ReturnType() TypeDefinition {
	return p
}

func (p PrimTypeDefinition) Name() TypeName {
	return p.name
}
