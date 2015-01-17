package expressionParsing

import (
	"errors"
	"fmt"
	"github.com/apoydence/gof/parser"
	"regexp"
)

var whiteSpaceRegexp *regexp.Regexp = regexp.MustCompile("\\s+")

type TypeDefinition interface {
	GenerateGo() string
}

type FuncTypeDefinition struct {
	argName string
	argType TypeDefinition
	retType TypeDefinition
	code    string
}

type PrimTypeDefinition struct {
	name string
}

func (pd PrimTypeDefinition) GenerateGo() string {
	return pd.name
}

func (pd PrimTypeDefinition) String() string {
	return pd.name
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

func newFuncTypeDefinition(argName string, argType, retType TypeDefinition) FuncTypeDefinition {
	return FuncTypeDefinition{
		argName: argName,
		argType: argType,
		retType: retType,
	}
}

func ParseTypeDef(code string) (TypeDefinition, error, string) {
	first, rest := getFirstToken(code)
	_, rest2 := getFirstToken(rest)
	third, _ := getFirstToken(rest2)
	if first == "func" {
		argName, rest := getFirstToken(rest)

		if !parser.ValidFunctionName(argName) {
			return nil, errors.New(fmt.Sprintf("%s is not a valid argument name.", argName)), ""
		}

		if len(rest) == 0 {
			return nil, errors.New("Invalid function type definition"), ""
		}
		td, err, rest := ParseTypeDef(rest)
		if err != nil {
			return nil, err, ""
		}
		ret, err, rest := getRetType(rest)
		return newFuncTypeDefinition(argName, td, ret), nil, rest
	} else if third == "->" {
		return ParseTypeDef("func " + code)
	}

	return newPrimTypeDefinition(first), nil, rest
}

func splitWhitespace(line string, n int) []string {
	return whiteSpaceRegexp.Split(line, n)
}

func getFirstToken(line string) (string, string) {
	broken := splitWhitespace(line, 2)
	if len(broken) == 2 {
		return broken[0], broken[1]
	}
	return line, ""
}

func getRetType(code string) (TypeDefinition, error, string) {
	first, rest := getFirstToken(code)
	if first == "->" {
		return ParseTypeDef(rest)
	}

	return nil, errors.New("Invalid function syntax"), ""
}
