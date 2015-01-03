package expressionParsing

import (
	"errors"
	"fmt"
	"regexp"
)

var funcArgTypeRegexp *regexp.Regexp
var funcRetTypeRegexp *regexp.Regexp

func init() {
	funcArgTypeRegexp = regexp.MustCompile("((?P<argName>([a-zA-Z]\\w*))\\s+(?P<argType>([a-zA-Z]\\w*))\\s*->)")
	funcRetTypeRegexp = regexp.MustCompile("(?P<returnType>([a-zA-Z]\\w*))$")
}

type TypeName string

type TypeDefinition interface {
	IsFunc() bool
	ReturnType() TypeDefinition
	Name() TypeName
}

type FuncTypeDefinition struct {
	Argument     TypeDefinition
	returnType   TypeDefinition
	name         string
	ArgumentName string
}

type PrimTypeDefinition struct {
	name TypeName
}

func NewFuncTypeDefinition(argName string, arg, retType TypeDefinition) *FuncTypeDefinition {
	return &FuncTypeDefinition{
		Argument:     arg,
		returnType:   retType,
		ArgumentName: argName,
	}
}

func ParseFuncTypeDefinition(str string) (*FuncTypeDefinition, error) {
	args, ret, err := fetchTypes(str)
	if err != nil {
		return nil, err
	}
	return convertToTypeDef(args, ret, 0).(*FuncTypeDefinition), nil
}

func convertToTypeDef(args []argDesc, retType TypeName, index int) TypeDefinition {
	if index >= len(args) {
		return NewPrimTypeDefinition(retType)
	} else {
		currentArg := args[index]
		a := NewPrimTypeDefinition(currentArg.typeName)
		return NewFuncTypeDefinition(currentArg.name, a, convertToTypeDef(args, retType, index+1))
	}
}

type argDesc struct {
	name     string
	typeName TypeName
}

func fetchTypes(code string) ([]argDesc, TypeName, error) {
	args := make([]argDesc, 0)
	groupIndex := make(map[string]int)
	for i, name := range funcArgTypeRegexp.SubexpNames() {
		groupIndex[name] = i
	}

	argsM := make(map[string]string)
	match := funcArgTypeRegexp.FindAllStringSubmatch(code, -1)

	for _, m := range match {
		name := m[groupIndex["argName"]]
		typeName := m[groupIndex["argType"]]
		if _, ok := argsM[name]; ok {
			return nil, "", errors.New(fmt.Sprintf("The argument name '%s' is used multiple times", name))
		}

		argsM[name] = typeName
		args = append(args, argDesc{name: name, typeName: TypeName(typeName)})
	}

	return args, TypeName(funcRetTypeRegexp.FindString(code)), nil
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
	return TypeName(fmt.Sprintf("%s %s->%s", f.ArgumentName, f.Argument.Name(), f.returnType.Name()))
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
